// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package redshiftserverless

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_redshiftserverless_workgroup", workgroupDataSource)
}

// workgroupDataSource returns the Terraform awscc_redshiftserverless_workgroup data source.
// This Terraform data source corresponds to the CloudFormation AWS::RedshiftServerless::Workgroup resource.
func workgroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"base_capacity": {
			// Property: BaseCapacity
			// CloudFormation resource type schema:
			// {
			//   "type": "integer"
			// }
			Type:     types.Int64Type,
			Computed: true,
		},
		"config_parameters": {
			// Property: ConfigParameters
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "ParameterKey": {
			//         "maxLength": 255,
			//         "minLength": 0,
			//         "type": "string"
			//       },
			//       "ParameterValue": {
			//         "maxLength": 15000,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"parameter_key": {
						// Property: ParameterKey
						Type:     types.StringType,
						Computed: true,
					},
					"parameter_value": {
						// Property: ParameterValue
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"enhanced_vpc_routing": {
			// Property: EnhancedVpcRouting
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Computed: true,
		},
		"namespace_name": {
			// Property: NamespaceName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 64,
			//   "minLength": 3,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"publicly_accessible": {
			// Property: PubliclyAccessible
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Computed: true,
		},
		"security_group_ids": {
			// Property: SecurityGroupIds
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "maxLength": 255,
			//     "minLength": 0,
			//     "pattern": "^sg-[0-9a-fA-F]{8,}$",
			//     "type": "string"
			//   },
			//   "maxItems": 32,
			//   "minItems": 1,
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"subnet_ids": {
			// Property: SubnetIds
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "maxLength": 255,
			//     "minLength": 0,
			//     "pattern": "^subnet-[0-9a-fA-F]{8,}$",
			//     "type": "string"
			//   },
			//   "maxItems": 32,
			//   "minItems": 1,
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
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
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "minItems": 0,
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
		"workgroup": {
			// Property: Workgroup
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "BaseCapacity": {
			//       "type": "integer"
			//     },
			//     "ConfigParameters": {
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "ParameterKey": {
			//             "maxLength": 255,
			//             "minLength": 0,
			//             "type": "string"
			//           },
			//           "ParameterValue": {
			//             "maxLength": 15000,
			//             "minLength": 0,
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "CreationDate": {
			//       "type": "string"
			//     },
			//     "Endpoint": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Address": {
			//           "type": "string"
			//         },
			//         "Port": {
			//           "type": "integer"
			//         },
			//         "VpcEndpoints": {
			//           "insertionOrder": false,
			//           "items": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "NetworkInterfaces": {
			//                 "insertionOrder": false,
			//                 "items": {
			//                   "additionalProperties": false,
			//                   "properties": {
			//                     "AvailabilityZone": {
			//                       "type": "string"
			//                     },
			//                     "NetworkInterfaceId": {
			//                       "type": "string"
			//                     },
			//                     "PrivateIpAddress": {
			//                       "type": "string"
			//                     },
			//                     "SubnetId": {
			//                       "type": "string"
			//                     }
			//                   },
			//                   "type": "object"
			//                 },
			//                 "type": "array"
			//               },
			//               "VpcEndpointId": {
			//                 "type": "string"
			//               },
			//               "VpcId": {
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "type": "array"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "EnhancedVpcRouting": {
			//       "type": "boolean"
			//     },
			//     "NamespaceName": {
			//       "maxLength": 64,
			//       "minLength": 3,
			//       "pattern": "^[a-z0-9-]+$",
			//       "type": "string"
			//     },
			//     "PubliclyAccessible": {
			//       "type": "boolean"
			//     },
			//     "SecurityGroupIds": {
			//       "insertionOrder": false,
			//       "items": {
			//         "maxLength": 255,
			//         "minLength": 0,
			//         "pattern": "^sg-[0-9a-fA-F]{8,}$",
			//         "type": "string"
			//       },
			//       "type": "array"
			//     },
			//     "Status": {
			//       "enum": [
			//         "CREATING",
			//         "AVAILABLE",
			//         "MODIFYING",
			//         "DELETING"
			//       ],
			//       "type": "string"
			//     },
			//     "SubnetIds": {
			//       "insertionOrder": false,
			//       "items": {
			//         "maxLength": 255,
			//         "minLength": 0,
			//         "pattern": "^subnet-[0-9a-fA-F]{8,}$",
			//         "type": "string"
			//       },
			//       "type": "array"
			//     },
			//     "WorkgroupArn": {
			//       "type": "string"
			//     },
			//     "WorkgroupId": {
			//       "type": "string"
			//     },
			//     "WorkgroupName": {
			//       "maxLength": 64,
			//       "minLength": 3,
			//       "pattern": "^[a-z0-9-]*$",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"base_capacity": {
						// Property: BaseCapacity
						Type:     types.Int64Type,
						Computed: true,
					},
					"config_parameters": {
						// Property: ConfigParameters
						Attributes: tfsdk.SetNestedAttributes(
							map[string]tfsdk.Attribute{
								"parameter_key": {
									// Property: ParameterKey
									Type:     types.StringType,
									Computed: true,
								},
								"parameter_value": {
									// Property: ParameterValue
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"creation_date": {
						// Property: CreationDate
						Type:     types.StringType,
						Computed: true,
					},
					"endpoint": {
						// Property: Endpoint
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"address": {
									// Property: Address
									Type:     types.StringType,
									Computed: true,
								},
								"port": {
									// Property: Port
									Type:     types.Int64Type,
									Computed: true,
								},
								"vpc_endpoints": {
									// Property: VpcEndpoints
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"network_interfaces": {
												// Property: NetworkInterfaces
												Attributes: tfsdk.ListNestedAttributes(
													map[string]tfsdk.Attribute{
														"availability_zone": {
															// Property: AvailabilityZone
															Type:     types.StringType,
															Computed: true,
														},
														"network_interface_id": {
															// Property: NetworkInterfaceId
															Type:     types.StringType,
															Computed: true,
														},
														"private_ip_address": {
															// Property: PrivateIpAddress
															Type:     types.StringType,
															Computed: true,
														},
														"subnet_id": {
															// Property: SubnetId
															Type:     types.StringType,
															Computed: true,
														},
													},
												),
												Computed: true,
											},
											"vpc_endpoint_id": {
												// Property: VpcEndpointId
												Type:     types.StringType,
												Computed: true,
											},
											"vpc_id": {
												// Property: VpcId
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"enhanced_vpc_routing": {
						// Property: EnhancedVpcRouting
						Type:     types.BoolType,
						Computed: true,
					},
					"namespace_name": {
						// Property: NamespaceName
						Type:     types.StringType,
						Computed: true,
					},
					"publicly_accessible": {
						// Property: PubliclyAccessible
						Type:     types.BoolType,
						Computed: true,
					},
					"security_group_ids": {
						// Property: SecurityGroupIds
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
					},
					"status": {
						// Property: Status
						Type:     types.StringType,
						Computed: true,
					},
					"subnet_ids": {
						// Property: SubnetIds
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
					},
					"workgroup_arn": {
						// Property: WorkgroupArn
						Type:     types.StringType,
						Computed: true,
					},
					"workgroup_id": {
						// Property: WorkgroupId
						Type:     types.StringType,
						Computed: true,
					},
					"workgroup_name": {
						// Property: WorkgroupName
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"workgroup_name": {
			// Property: WorkgroupName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 64,
			//   "minLength": 3,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::RedshiftServerless::Workgroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RedshiftServerless::Workgroup").WithTerraformTypeName("awscc_redshiftserverless_workgroup")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"address":              "Address",
		"availability_zone":    "AvailabilityZone",
		"base_capacity":        "BaseCapacity",
		"config_parameters":    "ConfigParameters",
		"creation_date":        "CreationDate",
		"endpoint":             "Endpoint",
		"enhanced_vpc_routing": "EnhancedVpcRouting",
		"key":                  "Key",
		"namespace_name":       "NamespaceName",
		"network_interface_id": "NetworkInterfaceId",
		"network_interfaces":   "NetworkInterfaces",
		"parameter_key":        "ParameterKey",
		"parameter_value":      "ParameterValue",
		"port":                 "Port",
		"private_ip_address":   "PrivateIpAddress",
		"publicly_accessible":  "PubliclyAccessible",
		"security_group_ids":   "SecurityGroupIds",
		"status":               "Status",
		"subnet_id":            "SubnetId",
		"subnet_ids":           "SubnetIds",
		"tags":                 "Tags",
		"value":                "Value",
		"vpc_endpoint_id":      "VpcEndpointId",
		"vpc_endpoints":        "VpcEndpoints",
		"vpc_id":               "VpcId",
		"workgroup":            "Workgroup",
		"workgroup_arn":        "WorkgroupArn",
		"workgroup_id":         "WorkgroupId",
		"workgroup_name":       "WorkgroupName",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}