// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package datazone

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/datazone"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory:  newDataSourceEnvironmentBlueprint,
			TypeName: "aws_datazone_environment_blueprint",
			Name:     "Environment Blueprint",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory:  newResourceAssetType,
			TypeName: "aws_datazone_asset_type",
			Name:     "Asset Type",
		},
		{
			Factory:  newResourceDomain,
			TypeName: "aws_datazone_domain",
			Name:     "Domain",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  newResourceEnvironment,
			TypeName: "aws_datazone_environment",
			Name:     "Environment",
		},
		{
			Factory:  newResourceEnvironmentBlueprintConfiguration,
			TypeName: "aws_datazone_environment_blueprint_configuration",
			Name:     "Environment Blueprint Configuration",
		},
		{
			Factory:  newResourceEnvironmentProfile,
			TypeName: "aws_datazone_environment_profile",
			Name:     "Environment Profile",
		},
		{
			Factory:  newResourceFormType,
			TypeName: "aws_datazone_form_type",
			Name:     "Form Type",
		},
		{
			Factory:  newResourceGlossary,
			TypeName: "aws_datazone_glossary",
			Name:     "Glossary",
		},
		{
			Factory:  newResourceGlossaryTerm,
			TypeName: "aws_datazone_glossary_term",
			Name:     "Glossary Term",
		},
		{
			Factory:  newResourceProject,
			TypeName: "aws_datazone_project",
			Name:     "Project",
		},
		{
			Factory:  newResourceUserProfile,
			TypeName: "aws_datazone_user_profile",
			Name:     "User Profile",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{}
}

func (p *servicePackage) ServicePackageName() string {
	return names.DataZone
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*datazone.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*datazone.Options){
		datazone.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return datazone.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*datazone.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*datazone.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *datazone.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*datazone.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
