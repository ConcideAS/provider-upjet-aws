// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	cachepolicy "github.com/upbound/provider-aws/v2/internal/controller/cluster/cloudfront/cachepolicy"
	connectiongroup "github.com/upbound/provider-aws/v2/internal/controller/cluster/cloudfront/connectiongroup"
	distribution "github.com/upbound/provider-aws/v2/internal/controller/cluster/cloudfront/distribution"
	distributiontenant "github.com/upbound/provider-aws/v2/internal/controller/cluster/cloudfront/distributiontenant"
	fieldlevelencryptionconfig "github.com/upbound/provider-aws/v2/internal/controller/cluster/cloudfront/fieldlevelencryptionconfig"
	fieldlevelencryptionprofile "github.com/upbound/provider-aws/v2/internal/controller/cluster/cloudfront/fieldlevelencryptionprofile"
	function "github.com/upbound/provider-aws/v2/internal/controller/cluster/cloudfront/function"
	keygroup "github.com/upbound/provider-aws/v2/internal/controller/cluster/cloudfront/keygroup"
	monitoringsubscription "github.com/upbound/provider-aws/v2/internal/controller/cluster/cloudfront/monitoringsubscription"
	multitenantdistribution "github.com/upbound/provider-aws/v2/internal/controller/cluster/cloudfront/multitenantdistribution"
	originaccesscontrol "github.com/upbound/provider-aws/v2/internal/controller/cluster/cloudfront/originaccesscontrol"
	originaccessidentity "github.com/upbound/provider-aws/v2/internal/controller/cluster/cloudfront/originaccessidentity"
	originrequestpolicy "github.com/upbound/provider-aws/v2/internal/controller/cluster/cloudfront/originrequestpolicy"
	publickey "github.com/upbound/provider-aws/v2/internal/controller/cluster/cloudfront/publickey"
	realtimelogconfig "github.com/upbound/provider-aws/v2/internal/controller/cluster/cloudfront/realtimelogconfig"
	responseheaderspolicy "github.com/upbound/provider-aws/v2/internal/controller/cluster/cloudfront/responseheaderspolicy"
	vpcorigin "github.com/upbound/provider-aws/v2/internal/controller/cluster/cloudfront/vpcorigin"
)

// Setup_cloudfront creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudfront(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cachepolicy.Setup,
		connectiongroup.Setup,
		distribution.Setup,
		distributiontenant.Setup,
		fieldlevelencryptionconfig.Setup,
		fieldlevelencryptionprofile.Setup,
		function.Setup,
		keygroup.Setup,
		monitoringsubscription.Setup,
		multitenantdistribution.Setup,
		originaccesscontrol.Setup,
		originaccessidentity.Setup,
		originrequestpolicy.Setup,
		publickey.Setup,
		realtimelogconfig.Setup,
		responseheaderspolicy.Setup,
		vpcorigin.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_cloudfront creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_cloudfront(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cachepolicy.SetupGated,
		connectiongroup.SetupGated,
		distribution.SetupGated,
		distributiontenant.SetupGated,
		fieldlevelencryptionconfig.SetupGated,
		fieldlevelencryptionprofile.SetupGated,
		function.SetupGated,
		keygroup.SetupGated,
		monitoringsubscription.SetupGated,
		multitenantdistribution.SetupGated,
		originaccesscontrol.SetupGated,
		originaccessidentity.SetupGated,
		originrequestpolicy.SetupGated,
		publickey.SetupGated,
		realtimelogconfig.SetupGated,
		responseheaderspolicy.SetupGated,
		vpcorigin.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_cloudfront registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_cloudfront(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		cachepolicy.SetupWebhookWithManager,
		connectiongroup.SetupWebhookWithManager,
		distribution.SetupWebhookWithManager,
		distributiontenant.SetupWebhookWithManager,
		fieldlevelencryptionconfig.SetupWebhookWithManager,
		fieldlevelencryptionprofile.SetupWebhookWithManager,
		function.SetupWebhookWithManager,
		keygroup.SetupWebhookWithManager,
		monitoringsubscription.SetupWebhookWithManager,
		multitenantdistribution.SetupWebhookWithManager,
		originaccesscontrol.SetupWebhookWithManager,
		originaccessidentity.SetupWebhookWithManager,
		originrequestpolicy.SetupWebhookWithManager,
		publickey.SetupWebhookWithManager,
		realtimelogconfig.SetupWebhookWithManager,
		responseheaderspolicy.SetupWebhookWithManager,
		vpcorigin.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
