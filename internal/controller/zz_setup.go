/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	template "github.com/haarchri/provider-opensearch/internal/controller/component/template"
	templateindex "github.com/haarchri/provider-opensearch/internal/controller/index/template"
	ismpolicy "github.com/haarchri/provider-opensearch/internal/controller/opensearch/ismpolicy"
	ismpolicymapping "github.com/haarchri/provider-opensearch/internal/controller/opensearch/ismpolicymapping"
	role "github.com/haarchri/provider-opensearch/internal/controller/opensearch/role"
	rolesmapping "github.com/haarchri/provider-opensearch/internal/controller/opensearch/rolesmapping"
	providerconfig "github.com/haarchri/provider-opensearch/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		template.Setup,
		templateindex.Setup,
		ismpolicy.Setup,
		ismpolicymapping.Setup,
		role.Setup,
		rolesmapping.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
