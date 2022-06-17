// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package wafv2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_wafv2_ip_set", iPSetDataSourceType)
}

// iPSetDataSourceType returns the Terraform awscc_wafv2_ip_set data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::WAFv2::IPSet resource type.
func iPSetDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"addresses": {
			// Property: Addresses
			// CloudFormation resource type schema:
			// {
			//   "description": "List of IPAddresses.",
			//   "items": {
			//     "description": "IP address",
			//     "maxLength": 50,
			//     "minLength": 1,
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "List of IPAddresses.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "ARN of the WAF entity.",
			//   "type": "string"
			// }
			Description: "ARN of the WAF entity.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "Description of the entity.",
			//   "pattern": "^[a-zA-Z0-9=:#@/\\-,.][a-zA-Z0-9+=:#@/\\-,.\\s]+[a-zA-Z0-9+=:#@/\\-,.]{1,256}$",
			//   "type": "string"
			// }
			Description: "Description of the entity.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ip_address_version": {
			// Property: IPAddressVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "Type of addresses in the IPSet, use IPV4 for IPV4 IP addresses, IPV6 for IPV6 address.",
			//   "enum": [
			//     "IPV4",
			//     "IPV6"
			//   ],
			//   "type": "string"
			// }
			Description: "Type of addresses in the IPSet, use IPV4 for IPV4 IP addresses, IPV6 for IPV6 address.",
			Type:        types.StringType,
			Computed:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Id of the IPSet",
			//   "pattern": "^[0-9a-f]{8}-(?:[0-9a-f]{4}-){3}[0-9a-f]{12}$",
			//   "type": "string"
			// }
			Description: "Id of the IPSet",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the IPSet.",
			//   "pattern": "^[0-9A-Za-z_-]{1,128}$",
			//   "type": "string"
			// }
			Description: "Name of the IPSet.",
			Type:        types.StringType,
			Computed:    true,
		},
		"scope": {
			// Property: Scope
			// CloudFormation resource type schema:
			// {
			//   "description": "Use CLOUDFRONT for CloudFront IPSet, use REGIONAL for Application Load Balancer and API Gateway.",
			//   "enum": [
			//     "CLOUDFRONT",
			//     "REGIONAL"
			//   ],
			//   "type": "string"
			// }
			Description: "Use CLOUDFRONT for CloudFront IPSet, use REGIONAL for Application Load Balancer and API Gateway.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::WAFv2::IPSet",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::WAFv2::IPSet").WithTerraformTypeName("awscc_wafv2_ip_set")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"addresses":          "Addresses",
		"arn":                "Arn",
		"description":        "Description",
		"id":                 "Id",
		"ip_address_version": "IPAddressVersion",
		"key":                "Key",
		"name":               "Name",
		"scope":              "Scope",
		"tags":               "Tags",
		"value":              "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
