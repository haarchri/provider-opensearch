/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	elasticsearchcomponenttemplate "github.com/haarchri/provider-opensearch/config/elasticsearchcomponenttemplate"
	elasticsearchindextemplate "github.com/haarchri/provider-opensearch/config/elasticsearchindextemplate"
	opensearchismpolicy "github.com/haarchri/provider-opensearch/config/opensearchismpolicy"
	opensearchismpolicymapping "github.com/haarchri/provider-opensearch/config/opensearchismpolicymapping"
	opensearchrole "github.com/haarchri/provider-opensearch/config/opensearchrole"
	opensearchrolemapping "github.com/haarchri/provider-opensearch/config/opensearchrolemapping"
)

const (
	resourcePrefix = "opensearch"
	modulePath     = "github.com/haarchri/provider-opensearch"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithShortName("opensearch"),
		ujconfig.WithRootGroup("opensearch.crossplane.io"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		opensearchrole.Configure,
		elasticsearchindextemplate.Configure,
		opensearchismpolicy.Configure,
		opensearchismpolicymapping.Configure,
		opensearchrolemapping.Configure,
		elasticsearchcomponenttemplate.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
