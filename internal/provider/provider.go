// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func NewinfraapiProvider() provider.Provider {
	return &infraapiProvider{}
}

var _ provider.Provider = (*infraapiProvider)(nil)

type infraapiProvider struct{}

func (p *infraapiProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "infraapi"
}

func (p *infraapiProvider) Schema(context.Context, provider.SchemaRequest, *provider.SchemaResponse) {
}

func (p *infraapiProvider) Configure(context.Context, provider.ConfigureRequest, *provider.ConfigureResponse) {
}

func (p *infraapiProvider) Resources(context.Context) []func() resource.Resource {
	return nil
}

func (p *infraapiProvider) DataSources(context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewinfraapiEndpointDataSource,
	}
}
