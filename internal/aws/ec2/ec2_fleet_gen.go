// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_ec2_ec2_fleet", eC2FleetResourceType)
}

// eC2FleetResourceType returns the Terraform aws_ec2_ec2_fleet resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EC2::EC2Fleet resource type.
func eC2FleetResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"context": {
			// Property: Context
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Optional: true,
		},
		"excess_capacity_termination_policy": {
			// Property: ExcessCapacityTerminationPolicy
			// CloudFormation resource type schema:
			/*
			   {
			     "enum": [
			       "termination",
			       "no-termination"
			     ],
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Optional: true,
		},
		"fleet_id": {
			// Property: FleetId
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"launch_template_configs": {
			// Property: LaunchTemplateConfigs
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "LaunchTemplateSpecification": {
			           "additionalProperties": false,
			           "properties": {
			             "LaunchTemplateId": {
			               "type": "string"
			             },
			             "LaunchTemplateName": {
			               "maxLength": 128,
			               "minLength": 3,
			               "pattern": "",
			               "type": "string"
			             },
			             "Version": {
			               "type": "string"
			             }
			           },
			           "$ref": "#/definitions/FleetLaunchTemplateSpecificationRequest",
			           "type": "object"
			         },
			         "Overrides": {
			           "items": {
			             "additionalProperties": false,
			             "properties": {
			               "AvailabilityZone": {
			                 "type": "string"
			               },
			               "InstanceType": {
			                 "type": "string"
			               },
			               "MaxPrice": {
			                 "type": "string"
			               },
			               "Placement": {
			                 "additionalProperties": false,
			                 "properties": {
			                   "Affinity": {
			                     "type": "string"
			                   },
			                   "AvailabilityZone": {
			                     "type": "string"
			                   },
			                   "GroupName": {
			                     "type": "string"
			                   },
			                   "HostId": {
			                     "type": "string"
			                   },
			                   "HostResourceGroupArn": {
			                     "type": "string"
			                   },
			                   "PartitionNumber": {
			                     "type": "integer"
			                   },
			                   "SpreadDomain": {
			                     "type": "string"
			                   },
			                   "Tenancy": {
			                     "type": "string"
			                   }
			                 },
			                 "$ref": "#/definitions/Placement",
			                 "type": "object"
			               },
			               "Priority": {
			                 "type": "number"
			               },
			               "SubnetId": {
			                 "type": "string"
			               },
			               "WeightedCapacity": {
			                 "type": "number"
			               }
			             },
			             "$ref": "#/definitions/FleetLaunchTemplateOverridesRequest",
			             "type": "object"
			           },
			           "type": "array",
			           "uniqueItems": false
			         }
			       },
			       "$ref": "#/definitions/FleetLaunchTemplateConfigRequest",
			       "type": "object"
			     },
			     "maxItems": 50,
			     "type": "array",
			     "uniqueItems": false
			   }
			*/
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"launch_template_specification": {
						// Property: LaunchTemplateSpecification
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "LaunchTemplateId": {
						         "type": "string"
						       },
						       "LaunchTemplateName": {
						         "maxLength": 128,
						         "minLength": 3,
						         "pattern": "",
						         "type": "string"
						       },
						       "Version": {
						         "type": "string"
						       }
						     },
						     "$ref": "#/definitions/FleetLaunchTemplateSpecificationRequest",
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"launch_template_id": {
									// Property: LaunchTemplateId
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"launch_template_name": {
									// Property: LaunchTemplateName
									// CloudFormation resource type schema:
									/*
									   {
									     "maxLength": 128,
									     "minLength": 3,
									     "pattern": "",
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"version": {
									// Property: Version
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"overrides": {
						// Property: Overrides
						// CloudFormation resource type schema:
						/*
						   {
						     "items": {
						       "additionalProperties": false,
						       "properties": {
						         "AvailabilityZone": {
						           "type": "string"
						         },
						         "InstanceType": {
						           "type": "string"
						         },
						         "MaxPrice": {
						           "type": "string"
						         },
						         "Placement": {
						           "additionalProperties": false,
						           "properties": {
						             "Affinity": {
						               "type": "string"
						             },
						             "AvailabilityZone": {
						               "type": "string"
						             },
						             "GroupName": {
						               "type": "string"
						             },
						             "HostId": {
						               "type": "string"
						             },
						             "HostResourceGroupArn": {
						               "type": "string"
						             },
						             "PartitionNumber": {
						               "type": "integer"
						             },
						             "SpreadDomain": {
						               "type": "string"
						             },
						             "Tenancy": {
						               "type": "string"
						             }
						           },
						           "$ref": "#/definitions/Placement",
						           "type": "object"
						         },
						         "Priority": {
						           "type": "number"
						         },
						         "SubnetId": {
						           "type": "string"
						         },
						         "WeightedCapacity": {
						           "type": "number"
						         }
						       },
						       "$ref": "#/definitions/FleetLaunchTemplateOverridesRequest",
						       "type": "object"
						     },
						     "type": "array",
						     "uniqueItems": false
						   }
						*/
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"availability_zone": {
									// Property: AvailabilityZone
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"instance_type": {
									// Property: InstanceType
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"max_price": {
									// Property: MaxPrice
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"placement": {
									// Property: Placement
									// CloudFormation resource type schema:
									/*
									   {
									     "additionalProperties": false,
									     "properties": {
									       "Affinity": {
									         "type": "string"
									       },
									       "AvailabilityZone": {
									         "type": "string"
									       },
									       "GroupName": {
									         "type": "string"
									       },
									       "HostId": {
									         "type": "string"
									       },
									       "HostResourceGroupArn": {
									         "type": "string"
									       },
									       "PartitionNumber": {
									         "type": "integer"
									       },
									       "SpreadDomain": {
									         "type": "string"
									       },
									       "Tenancy": {
									         "type": "string"
									       }
									     },
									     "$ref": "#/definitions/Placement",
									     "type": "object"
									   }
									*/
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"affinity": {
												// Property: Affinity
												// CloudFormation resource type schema:
												/*
												   {
												     "type": "string"
												   }
												*/
												Type:     types.StringType,
												Optional: true,
											},
											"availability_zone": {
												// Property: AvailabilityZone
												// CloudFormation resource type schema:
												/*
												   {
												     "type": "string"
												   }
												*/
												Type:     types.StringType,
												Optional: true,
											},
											"group_name": {
												// Property: GroupName
												// CloudFormation resource type schema:
												/*
												   {
												     "type": "string"
												   }
												*/
												Type:     types.StringType,
												Optional: true,
											},
											"host_id": {
												// Property: HostId
												// CloudFormation resource type schema:
												/*
												   {
												     "type": "string"
												   }
												*/
												Type:     types.StringType,
												Optional: true,
											},
											"host_resource_group_arn": {
												// Property: HostResourceGroupArn
												// CloudFormation resource type schema:
												/*
												   {
												     "type": "string"
												   }
												*/
												Type:     types.StringType,
												Optional: true,
											},
											"partition_number": {
												// Property: PartitionNumber
												// CloudFormation resource type schema:
												/*
												   {
												     "type": "integer"
												   }
												*/
												Type:     types.NumberType,
												Optional: true,
											},
											"spread_domain": {
												// Property: SpreadDomain
												// CloudFormation resource type schema:
												/*
												   {
												     "type": "string"
												   }
												*/
												Type:     types.StringType,
												Optional: true,
											},
											"tenancy": {
												// Property: Tenancy
												// CloudFormation resource type schema:
												/*
												   {
												     "type": "string"
												   }
												*/
												Type:     types.StringType,
												Optional: true,
											},
										},
									),
									Optional: true,
								},
								"priority": {
									// Property: Priority
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "number"
									   }
									*/
									Type:     types.NumberType,
									Optional: true,
								},
								"subnet_id": {
									// Property: SubnetId
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"weighted_capacity": {
									// Property: WeightedCapacity
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "number"
									   }
									*/
									Type:     types.NumberType,
									Optional: true,
								},
							},
							schema.ListNestedAttributesOptions{},
						),
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{
					MaxItems: 50,
				},
			),
			Required: true,
			// LaunchTemplateConfigs is a force-new attribute.
		},
		"on_demand_options": {
			// Property: OnDemandOptions
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "AllocationStrategy": {
			         "type": "string"
			       },
			       "CapacityReservationOptions": {
			         "additionalProperties": false,
			         "properties": {
			           "UsageStrategy": {
			             "enum": [
			               "use-capacity-reservations-first"
			             ],
			             "type": "string"
			           }
			         },
			         "$ref": "#/definitions/CapacityReservationOptionsRequest",
			         "type": "object"
			       },
			       "MaxTotalPrice": {
			         "type": "string"
			       },
			       "MinTargetCapacity": {
			         "type": "integer"
			       },
			       "SingleAvailabilityZone": {
			         "type": "boolean"
			       },
			       "SingleInstanceType": {
			         "type": "boolean"
			       }
			     },
			     "$ref": "#/definitions/OnDemandOptionsRequest",
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"allocation_strategy": {
						// Property: AllocationStrategy
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"capacity_reservation_options": {
						// Property: CapacityReservationOptions
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "UsageStrategy": {
						         "enum": [
						           "use-capacity-reservations-first"
						         ],
						         "type": "string"
						       }
						     },
						     "$ref": "#/definitions/CapacityReservationOptionsRequest",
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"usage_strategy": {
									// Property: UsageStrategy
									// CloudFormation resource type schema:
									/*
									   {
									     "enum": [
									       "use-capacity-reservations-first"
									     ],
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"max_total_price": {
						// Property: MaxTotalPrice
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"min_target_capacity": {
						// Property: MinTargetCapacity
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "integer"
						   }
						*/
						Type:     types.NumberType,
						Optional: true,
					},
					"single_availability_zone": {
						// Property: SingleAvailabilityZone
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "boolean"
						   }
						*/
						Type:     types.BoolType,
						Optional: true,
					},
					"single_instance_type": {
						// Property: SingleInstanceType
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "boolean"
						   }
						*/
						Type:     types.BoolType,
						Optional: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			// OnDemandOptions is a force-new attribute.
		},
		"replace_unhealthy_instances": {
			// Property: ReplaceUnhealthyInstances
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "boolean"
			   }
			*/
			Type:     types.BoolType,
			Optional: true,
			Computed: true,
			// ReplaceUnhealthyInstances is a force-new attribute.
		},
		"spot_options": {
			// Property: SpotOptions
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "AllocationStrategy": {
			         "enum": [
			           "lowestPrice",
			           "diversified",
			           "capacityOptimized"
			         ],
			         "type": "string"
			       },
			       "InstanceInterruptionBehavior": {
			         "enum": [
			           "hibernate",
			           "stop",
			           "terminate"
			         ],
			         "type": "string"
			       },
			       "InstancePoolsToUseCount": {
			         "type": "integer"
			       },
			       "MaxTotalPrice": {
			         "type": "string"
			       },
			       "MinTargetCapacity": {
			         "type": "integer"
			       },
			       "SingleAvailabilityZone": {
			         "type": "boolean"
			       },
			       "SingleInstanceType": {
			         "type": "boolean"
			       }
			     },
			     "$ref": "#/definitions/SpotOptionsRequest",
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"allocation_strategy": {
						// Property: AllocationStrategy
						// CloudFormation resource type schema:
						/*
						   {
						     "enum": [
						       "lowestPrice",
						       "diversified",
						       "capacityOptimized"
						     ],
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"instance_interruption_behavior": {
						// Property: InstanceInterruptionBehavior
						// CloudFormation resource type schema:
						/*
						   {
						     "enum": [
						       "hibernate",
						       "stop",
						       "terminate"
						     ],
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"instance_pools_to_use_count": {
						// Property: InstancePoolsToUseCount
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "integer"
						   }
						*/
						Type:     types.NumberType,
						Optional: true,
					},
					"max_total_price": {
						// Property: MaxTotalPrice
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"min_target_capacity": {
						// Property: MinTargetCapacity
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "integer"
						   }
						*/
						Type:     types.NumberType,
						Optional: true,
					},
					"single_availability_zone": {
						// Property: SingleAvailabilityZone
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "boolean"
						   }
						*/
						Type:     types.BoolType,
						Optional: true,
					},
					"single_instance_type": {
						// Property: SingleInstanceType
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "boolean"
						   }
						*/
						Type:     types.BoolType,
						Optional: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			// SpotOptions is a force-new attribute.
		},
		"tag_specifications": {
			// Property: TagSpecifications
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "ResourceType": {
			           "enum": [
			             "client-vpn-endpoint",
			             "customer-gateway",
			             "dedicated-host",
			             "dhcp-options",
			             "egress-only-internet-gateway",
			             "elastic-gpu",
			             "elastic-ip",
			             "export-image-task",
			             "export-instance-task",
			             "fleet",
			             "fpga-image",
			             "host-reservation",
			             "image",
			             "import-image-task",
			             "import-snapshot-task",
			             "instance",
			             "internet-gateway",
			             "key-pair",
			             "launch-template",
			             "local-gateway-route-table-vpc-association",
			             "natgateway",
			             "network-acl",
			             "network-insights-analysis",
			             "network-insights-path",
			             "network-interface",
			             "placement-group",
			             "reserved-instances",
			             "route-table",
			             "security-group",
			             "snapshot",
			             "spot-fleet-request",
			             "spot-instances-request",
			             "subnet",
			             "traffic-mirror-filter",
			             "traffic-mirror-session",
			             "traffic-mirror-target",
			             "transit-gateway",
			             "transit-gateway-attachment",
			             "transit-gateway-connect-peer",
			             "transit-gateway-multicast-domain",
			             "transit-gateway-route-table",
			             "volume",
			             "vpc",
			             "vpc-flow-log",
			             "vpc-peering-connection",
			             "vpn-connection",
			             "vpn-gateway"
			           ],
			           "type": "string"
			         },
			         "Tags": {
			           "items": {
			             "additionalProperties": false,
			             "properties": {
			               "Key": {
			                 "type": "string"
			               },
			               "Value": {
			                 "type": "string"
			               }
			             },
			             "$ref": "#/definitions/Tag",
			             "required": [
			               "Value",
			               "Key"
			             ],
			             "type": "object"
			           },
			           "type": "array",
			           "uniqueItems": false
			         }
			       },
			       "$ref": "#/definitions/TagSpecification",
			       "type": "object"
			     },
			     "type": "array",
			     "uniqueItems": false
			   }
			*/
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"resource_type": {
						// Property: ResourceType
						// CloudFormation resource type schema:
						/*
						   {
						     "enum": [
						       "client-vpn-endpoint",
						       "customer-gateway",
						       "dedicated-host",
						       "dhcp-options",
						       "egress-only-internet-gateway",
						       "elastic-gpu",
						       "elastic-ip",
						       "export-image-task",
						       "export-instance-task",
						       "fleet",
						       "fpga-image",
						       "host-reservation",
						       "image",
						       "import-image-task",
						       "import-snapshot-task",
						       "instance",
						       "internet-gateway",
						       "key-pair",
						       "launch-template",
						       "local-gateway-route-table-vpc-association",
						       "natgateway",
						       "network-acl",
						       "network-insights-analysis",
						       "network-insights-path",
						       "network-interface",
						       "placement-group",
						       "reserved-instances",
						       "route-table",
						       "security-group",
						       "snapshot",
						       "spot-fleet-request",
						       "spot-instances-request",
						       "subnet",
						       "traffic-mirror-filter",
						       "traffic-mirror-session",
						       "traffic-mirror-target",
						       "transit-gateway",
						       "transit-gateway-attachment",
						       "transit-gateway-connect-peer",
						       "transit-gateway-multicast-domain",
						       "transit-gateway-route-table",
						       "volume",
						       "vpc",
						       "vpc-flow-log",
						       "vpc-peering-connection",
						       "vpn-connection",
						       "vpn-gateway"
						     ],
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"tags": {
						// Property: Tags
						// CloudFormation resource type schema:
						/*
						   {
						     "items": {
						       "additionalProperties": false,
						       "properties": {
						         "Key": {
						           "type": "string"
						         },
						         "Value": {
						           "type": "string"
						         }
						       },
						       "$ref": "#/definitions/Tag",
						       "required": [
						         "Value",
						         "Key"
						       ],
						       "type": "object"
						     },
						     "type": "array",
						     "uniqueItems": false
						   }
						*/
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"key": {
									// Property: Key
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
								"value": {
									// Property: Value
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
							},
							schema.ListNestedAttributesOptions{},
						),
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
			Computed: true,
			// TagSpecifications is a force-new attribute.
		},
		"target_capacity_specification": {
			// Property: TargetCapacitySpecification
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "DefaultTargetCapacityType": {
			         "enum": [
			           "on-demand",
			           "spot"
			         ],
			         "type": "string"
			       },
			       "OnDemandTargetCapacity": {
			         "type": "integer"
			       },
			       "SpotTargetCapacity": {
			         "type": "integer"
			       },
			       "TotalTargetCapacity": {
			         "type": "integer"
			       }
			     },
			     "$ref": "#/definitions/TargetCapacitySpecificationRequest",
			     "required": [
			       "TotalTargetCapacity"
			     ],
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"default_target_capacity_type": {
						// Property: DefaultTargetCapacityType
						// CloudFormation resource type schema:
						/*
						   {
						     "enum": [
						       "on-demand",
						       "spot"
						     ],
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"on_demand_target_capacity": {
						// Property: OnDemandTargetCapacity
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "integer"
						   }
						*/
						Type:     types.NumberType,
						Optional: true,
					},
					"spot_target_capacity": {
						// Property: SpotTargetCapacity
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "integer"
						   }
						*/
						Type:     types.NumberType,
						Optional: true,
					},
					"total_target_capacity": {
						// Property: TotalTargetCapacity
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "integer"
						   }
						*/
						Type:     types.NumberType,
						Required: true,
					},
				},
			),
			Required: true,
		},
		"terminate_instances_with_expiration": {
			// Property: TerminateInstancesWithExpiration
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "boolean"
			   }
			*/
			Type:     types.BoolType,
			Optional: true,
			Computed: true,
			// TerminateInstancesWithExpiration is a force-new attribute.
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			/*
			   {
			     "enum": [
			       "maintain",
			       "request",
			       "instant"
			     ],
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			// Type is a force-new attribute.
		},
		"valid_from": {
			// Property: ValidFrom
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			// ValidFrom is a force-new attribute.
		},
		"valid_until": {
			// Property: ValidUntil
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			// ValidUntil is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::EC2::EC2Fleet",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::EC2Fleet").WithTerraformTypeName("aws_ec2_ec2_fleet").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_ec2_ec2_fleet", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
