// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ce

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ce_anomaly_monitor", anomalyMonitorDataSourceType)
}

// anomalyMonitorDataSourceType returns the Terraform awscc_ce_anomaly_monitor data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::CE::AnomalyMonitor resource type.
func anomalyMonitorDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"creation_date": {
			// Property: CreationDate
			// CloudFormation resource type schema:
			// {
			//   "description": "The date when the monitor was created. ",
			//   "maxLength": 40,
			//   "minLength": 0,
			//   "pattern": "(\\d{4}-\\d{2}-\\d{2})(T\\d{2}:\\d{2}:\\d{2}Z)?",
			//   "type": "string"
			// }
			Description: "The date when the monitor was created. ",
			Type:        types.StringType,
			Computed:    true,
		},
		"dimensional_value_count": {
			// Property: DimensionalValueCount
			// CloudFormation resource type schema:
			// {
			//   "description": "The value for evaluated dimensions.",
			//   "minimum": 0,
			//   "type": "integer"
			// }
			Description: "The value for evaluated dimensions.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"last_evaluated_date": {
			// Property: LastEvaluatedDate
			// CloudFormation resource type schema:
			// {
			//   "description": "The date when the monitor last evaluated for anomalies.",
			//   "maxLength": 40,
			//   "minLength": 0,
			//   "pattern": "(\\d{4}-\\d{2}-\\d{2})(T\\d{2}:\\d{2}:\\d{2}Z)?",
			//   "type": "string"
			// }
			Description: "The date when the monitor last evaluated for anomalies.",
			Type:        types.StringType,
			Computed:    true,
		},
		"last_updated_date": {
			// Property: LastUpdatedDate
			// CloudFormation resource type schema:
			// {
			//   "description": "The date when the monitor was last updated.",
			//   "maxLength": 40,
			//   "minLength": 0,
			//   "pattern": "(\\d{4}-\\d{2}-\\d{2})(T\\d{2}:\\d{2}:\\d{2}Z)?",
			//   "type": "string"
			// }
			Description: "The date when the monitor was last updated.",
			Type:        types.StringType,
			Computed:    true,
		},
		"monitor_arn": {
			// Property: MonitorArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Monitor ARN",
			//   "pattern": "^arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:[-a-zA-Z0-9/:_]+$",
			//   "type": "string"
			// }
			Description: "Monitor ARN",
			Type:        types.StringType,
			Computed:    true,
		},
		"monitor_dimension": {
			// Property: MonitorDimension
			// CloudFormation resource type schema:
			// {
			//   "description": "The dimensions to evaluate",
			//   "enum": [
			//     "SERVICE"
			//   ],
			//   "type": "string"
			// }
			Description: "The dimensions to evaluate",
			Type:        types.StringType,
			Computed:    true,
		},
		"monitor_name": {
			// Property: MonitorName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the monitor.",
			//   "maxLength": 1024,
			//   "minLength": 0,
			//   "pattern": "[\\S\\s]*",
			//   "type": "string"
			// }
			Description: "The name of the monitor.",
			Type:        types.StringType,
			Computed:    true,
		},
		"monitor_specification": {
			// Property: MonitorSpecification
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"monitor_type": {
			// Property: MonitorType
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "DIMENSIONAL",
			//     "CUSTOM"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"resource_tags": {
			// Property: ResourceTags
			// CloudFormation resource type schema:
			// {
			//   "description": "Tags to assign to monitor.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name for the tag.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag.",
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
			Description: "Tags to assign to monitor.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name for the tag.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag.",
						Type:        types.StringType,
						Computed:    true,
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
		Description: "Data Source schema for AWS::CE::AnomalyMonitor",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CE::AnomalyMonitor").WithTerraformTypeName("awscc_ce_anomaly_monitor")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"creation_date":           "CreationDate",
		"dimensional_value_count": "DimensionalValueCount",
		"key":                     "Key",
		"last_evaluated_date":     "LastEvaluatedDate",
		"last_updated_date":       "LastUpdatedDate",
		"monitor_arn":             "MonitorArn",
		"monitor_dimension":       "MonitorDimension",
		"monitor_name":            "MonitorName",
		"monitor_specification":   "MonitorSpecification",
		"monitor_type":            "MonitorType",
		"resource_tags":           "ResourceTags",
		"value":                   "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
