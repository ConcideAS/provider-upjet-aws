// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package cloudfront

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-aws/v2/config/cluster/common"
)

// Configure adds configurations for the cloudfront group.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("aws_cloudfront_distribution", func(r *config.Resource) {
		r.UseAsync = true
		delete(r.References, "origin.domain_name")
		// Temporary edit until example generation pipeline fix
		for i, exp := range r.MetaResource.Examples {
			if exp.Name == "s3_distribution" {
				r.MetaResource.Examples = append(r.MetaResource.Examples[:i], r.MetaResource.Examples[i+1:]...)
				break
			}
		}
	})

	// Setting the field as sensitive to be able to pass the content from a k8s secret
	p.AddResourceConfigurator("aws_cloudfront_function", func(r *config.Resource) {
		r.TerraformResource.Schema["code"].Sensitive = true
	})

	// Setting the field as sensitive to be able to pass the content from a k8s secret
	p.AddResourceConfigurator("aws_cloudfront_public_key", func(r *config.Resource) {
		r.TerraformResource.Schema["encoded_key"].Sensitive = true
	})

	p.AddResourceConfigurator("aws_cloudfront_key_group", func(r *config.Resource) {
		r.References["items"] = config.Reference{
			TerraformName:     "aws_cloudfront_public_key",
			RefFieldName:      "ItemRefs",
			SelectorFieldName: "ItemSelector",
		}
	})

	p.AddResourceConfigurator("aws_cloudfront_realtime_log_config", func(r *config.Resource) {
		r.References["endpoint.kinesis_stream_config.stream_arn"] = config.Reference{
			TerraformName: "aws_kinesis_stream",
			Extractor:     common.PathTerraformIDExtractor,
		}
	})

	p.AddResourceConfigurator("aws_cloudfront_vpc_origin", func(r *config.Resource) {
		r.AddSingletonListConversion("vpc_origin_endpoint_config", "vpcOriginEndpointConfig")
		r.AddSingletonListConversion("vpc_origin_endpoint_config[*].origin_ssl_protocols", "vpcOriginEndpointConfig[*].originSslProtocols")
	})
	p.AddResourceConfigurator("aws_cloudfront_connection_group", func(r *config.Resource) {
		// No UseAsync. upjet's async path for terraform-plugin-framework
		// resources runs Create against a deep copy of the managed resource
		// (external_async_tfpluginfw.go:215,246 - the copy exists to avoid a
		// data race, upjet#472) and discards the ExternalCreation it returns.
		// The sync Create is what calls setExternalName(mg, state), so on the
		// async path the external name is written to an object that is thrown
		// away and never reaches the API server.
		//
		// The resource is then unrecoverable: Crossplane records
		// external-create-succeeded, the external name stays empty, Observe
		// resolves the stub id and finds nothing, and Crossplane refuses to
		// create again rather than leak a second resource. Reproduced twice on
		// aws-eu-central-1-dev, leaving two orphaned distributions.
		//
		// These creates return as soon as CloudFront allocates an id - the
		// propagation to Deployed is not waited on - so a synchronous Create
		// does not hold a reconcile worker for long.
	})

	p.AddResourceConfigurator("aws_cloudfront_distribution_tenant", func(r *config.Resource) {
		// No UseAsync. upjet's async path for terraform-plugin-framework
		// resources runs Create against a deep copy of the managed resource
		// (external_async_tfpluginfw.go:215,246 - the copy exists to avoid a
		// data race, upjet#472) and discards the ExternalCreation it returns.
		// The sync Create is what calls setExternalName(mg, state), so on the
		// async path the external name is written to an object that is thrown
		// away and never reaches the API server.
		//
		// The resource is then unrecoverable: Crossplane records
		// external-create-succeeded, the external name stays empty, Observe
		// resolves the stub id and finds nothing, and Crossplane refuses to
		// create again rather than leak a second resource. Reproduced twice on
		// aws-eu-central-1-dev, leaving two orphaned distributions.
		//
		// These creates return as soon as CloudFront allocates an id - the
		// propagation to Deployed is not waited on - so a synchronous Create
		// does not hold a reconcile worker for long.
		// connection_group_id has no *Ref of its own, so saga-gitops wrapped this
		// resource in a provider-kubernetes Object and patched the id in from the
		// ConnectionGroup's status. That wrapper is what loses the external name:
		// two controllers then write the same object, and the annotation write
		// loses the race with the Object's re-apply -
		//   Cannot initialize managed resource ... the object has been modified;
		//   please apply your changes to the latest version and try again
		// A native reference removes the wrapper, and with it the conflict.
		r.References["connection_group_id"] = config.Reference{
			TerraformName: "aws_cloudfront_connection_group",
		}
		r.AddSingletonListConversion("customizations", "customizations")
		r.AddSingletonListConversion("customizations[*].certificate", "customizations[*].certificate")
		r.AddSingletonListConversion("customizations[*].geo_restriction", "customizations[*].geoRestriction")
		r.AddSingletonListConversion("customizations[*].web_acl", "customizations[*].webAcl")
		r.AddSingletonListConversion("managed_certificate_request", "managedCertificateRequest")
	})

	p.AddResourceConfigurator("aws_cloudfront_multitenant_distribution", func(r *config.Resource) {
		// No UseAsync. upjet's async path for terraform-plugin-framework
		// resources runs Create against a deep copy of the managed resource
		// (external_async_tfpluginfw.go:215,246 - the copy exists to avoid a
		// data race, upjet#472) and discards the ExternalCreation it returns.
		// The sync Create is what calls setExternalName(mg, state), so on the
		// async path the external name is written to an object that is thrown
		// away and never reaches the API server.
		//
		// The resource is then unrecoverable: Crossplane records
		// external-create-succeeded, the external name stays empty, Observe
		// resolves the stub id and finds nothing, and Crossplane refuses to
		// create again rather than leak a second resource. Reproduced twice on
		// aws-eu-central-1-dev, leaving two orphaned distributions.
		//
		// These creates return as soon as CloudFront allocates an id - the
		// propagation to Deployed is not waited on - so a synchronous Create
		// does not hold a reconcile worker for long.
		// Same reason as connection_group_id on the tenant: without this the
		// vpc origin id can only be patched in by an Object wrapper, and that
		// wrapper is what clobbers the external-name annotation.
		r.References["origin.vpc_origin_config.vpc_origin_id"] = config.Reference{
			TerraformName: "aws_cloudfront_vpc_origin",
		}
		r.AddSingletonListConversion("cache_behavior[*].allowed_methods", "cacheBehavior[*].allowedMethods")
		r.AddSingletonListConversion("cache_behavior[*].trusted_key_groups", "cacheBehavior[*].trustedKeyGroups")
		r.AddSingletonListConversion("default_cache_behavior", "defaultCacheBehavior")
		r.AddSingletonListConversion("default_cache_behavior[*].allowed_methods", "defaultCacheBehavior[*].allowedMethods")
		r.AddSingletonListConversion("default_cache_behavior[*].trusted_key_groups", "defaultCacheBehavior[*].trustedKeyGroups")
		r.AddSingletonListConversion("origin[*].custom_origin_config", "origin[*].customOriginConfig")
		r.AddSingletonListConversion("origin[*].custom_origin_config[*].origin_mtls_config", "origin[*].customOriginConfig[*].originMtlsConfig")
		r.AddSingletonListConversion("origin[*].origin_shield", "origin[*].originShield")
		r.AddSingletonListConversion("origin[*].vpc_origin_config", "origin[*].vpcOriginConfig")
		r.AddSingletonListConversion("origin_group[*].failover_criteria", "originGroup[*].failoverCriteria")
		r.AddSingletonListConversion("restrictions", "restrictions")
		r.AddSingletonListConversion("restrictions[*].geo_restriction", "restrictions[*].geoRestriction")
		r.AddSingletonListConversion("tenant_config", "tenantConfig")
		r.AddSingletonListConversion("tenant_config[*].parameter_definition[*].definition", "tenantConfig[*].parameterDefinition[*].definition")
		r.AddSingletonListConversion("tenant_config[*].parameter_definition[*].definition[*].string_schema", "tenantConfig[*].parameterDefinition[*].definition[*].stringSchema")
		r.AddSingletonListConversion("viewer_certificate", "viewerCertificate")
	})

}
