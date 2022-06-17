// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_lightsail_load_balancer", loadBalancerDataSourceType)
}

// loadBalancerDataSourceType returns the Terraform awscc_lightsail_load_balancer data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Lightsail::LoadBalancer resource type.
func loadBalancerDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"attached_instances": {
			// Property: AttachedInstances
			// CloudFormation resource type schema:
			// {
			//   "description": "The names of the instances attached to the load balancer.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "The names of the instances attached to the load balancer.",
			Type:        types.SetType{ElemType: types.StringType},
			Computed:    true,
		},
		"health_check_path": {
			// Property: HealthCheckPath
			// CloudFormation resource type schema:
			// {
			//   "description": "The path you provided to perform the load balancer health check. If you didn't specify a health check path, Lightsail uses the root path of your website (e.g., \"/\").",
			//   "type": "string"
			// }
			Description: "The path you provided to perform the load balancer health check. If you didn't specify a health check path, Lightsail uses the root path of your website (e.g., \"/\").",
			Type:        types.StringType,
			Computed:    true,
		},
		"instance_port": {
			// Property: InstancePort
			// CloudFormation resource type schema:
			// {
			//   "description": "The instance port where you're creating your load balancer.",
			//   "type": "integer"
			// }
			Description: "The instance port where you're creating your load balancer.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"ip_address_type": {
			// Property: IpAddressType
			// CloudFormation resource type schema:
			// {
			//   "description": "The IP address type for the load balancer. The possible values are ipv4 for IPv4 only, and dualstack for IPv4 and IPv6. The default value is dualstack.",
			//   "type": "string"
			// }
			Description: "The IP address type for the load balancer. The possible values are ipv4 for IPv4 only, and dualstack for IPv4 and IPv6. The default value is dualstack.",
			Type:        types.StringType,
			Computed:    true,
		},
		"load_balancer_arn": {
			// Property: LoadBalancerArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"load_balancer_name": {
			// Property: LoadBalancerName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of your load balancer.",
			//   "pattern": "\\w[\\w\\-]*\\w",
			//   "type": "string"
			// }
			Description: "The name of your load balancer.",
			Type:        types.StringType,
			Computed:    true,
		},
		"session_stickiness_enabled": {
			// Property: SessionStickinessEnabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Configuration option to enable session stickiness.",
			//   "type": "boolean"
			// }
			Description: "Configuration option to enable session stickiness.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"session_stickiness_lb_cookie_duration_seconds": {
			// Property: SessionStickinessLBCookieDurationSeconds
			// CloudFormation resource type schema:
			// {
			//   "description": "Configuration option to adjust session stickiness cookie duration parameter.",
			//   "type": "string"
			// }
			Description: "Configuration option to adjust session stickiness cookie duration parameter.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"tls_policy_name": {
			// Property: TlsPolicyName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the TLS policy to apply to the load balancer.",
			//   "type": "string"
			// }
			Description: "The name of the TLS policy to apply to the load balancer.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Lightsail::LoadBalancer",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lightsail::LoadBalancer").WithTerraformTypeName("awscc_lightsail_load_balancer")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"attached_instances":         "AttachedInstances",
		"health_check_path":          "HealthCheckPath",
		"instance_port":              "InstancePort",
		"ip_address_type":            "IpAddressType",
		"key":                        "Key",
		"load_balancer_arn":          "LoadBalancerArn",
		"load_balancer_name":         "LoadBalancerName",
		"session_stickiness_enabled": "SessionStickinessEnabled",
		"session_stickiness_lb_cookie_duration_seconds": "SessionStickinessLBCookieDurationSeconds",
		"tags":            "Tags",
		"tls_policy_name": "TlsPolicyName",
		"value":           "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
