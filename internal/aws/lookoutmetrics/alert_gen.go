// Code generated by generators/resource/main.go; DO NOT EDIT.

package lookoutmetrics

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
	registry.AddResourceTypeFactory("aws_lookoutmetrics_alert", alertResourceType)
}

// alertResourceType returns the Terraform aws_lookoutmetrics_alert resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::LookoutMetrics::Alert resource type.
func alertResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"action": {
			// Property: Action
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "LambdaConfiguration": {
			         "additionalProperties": false,
			         "description": "Configuration options for a Lambda alert action.",
			         "properties": {
			           "LambdaArn": {
			             "maxLength": 256,
			             "pattern": "",
			             "$ref": "#/definitions/Arn",
			             "type": "string"
			           },
			           "RoleArn": {
			             "maxLength": 256,
			             "pattern": "",
			             "$ref": "#/definitions/Arn",
			             "type": "string"
			           }
			         },
			         "$ref": "#/definitions/LambdaConfiguration",
			         "required": [
			           "RoleArn",
			           "LambdaArn"
			         ],
			         "type": "object"
			       },
			       "SNSConfiguration": {
			         "additionalProperties": false,
			         "description": "Configuration options for an SNS alert action.",
			         "properties": {
			           "RoleArn": {
			             "maxLength": 256,
			             "pattern": "",
			             "$ref": "#/definitions/Arn",
			             "type": "string"
			           },
			           "SnsTopicArn": {
			             "maxLength": 256,
			             "pattern": "",
			             "$ref": "#/definitions/Arn",
			             "type": "string"
			           }
			         },
			         "$ref": "#/definitions/SNSConfiguration",
			         "required": [
			           "RoleArn",
			           "SnsTopicArn"
			         ],
			         "type": "object"
			       }
			     },
			     "$ref": "#/definitions/Action",
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"lambda_configuration": {
						// Property: LambdaConfiguration
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "description": "Configuration options for a Lambda alert action.",
						     "properties": {
						       "LambdaArn": {
						         "maxLength": 256,
						         "pattern": "",
						         "$ref": "#/definitions/Arn",
						         "type": "string"
						       },
						       "RoleArn": {
						         "maxLength": 256,
						         "pattern": "",
						         "$ref": "#/definitions/Arn",
						         "type": "string"
						       }
						     },
						     "$ref": "#/definitions/LambdaConfiguration",
						     "required": [
						       "RoleArn",
						       "LambdaArn"
						     ],
						     "type": "object"
						   }
						*/
						Description: "Configuration options for a Lambda alert action.",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"lambda_arn": {
									// Property: LambdaArn
									// CloudFormation resource type schema:
									/*
									   {
									     "maxLength": 256,
									     "pattern": "",
									     "$ref": "#/definitions/Arn",
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
								"role_arn": {
									// Property: RoleArn
									// CloudFormation resource type schema:
									/*
									   {
									     "maxLength": 256,
									     "pattern": "",
									     "$ref": "#/definitions/Arn",
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
							},
						),
						Optional: true,
					},
					"sns_configuration": {
						// Property: SNSConfiguration
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "description": "Configuration options for an SNS alert action.",
						     "properties": {
						       "RoleArn": {
						         "maxLength": 256,
						         "pattern": "",
						         "$ref": "#/definitions/Arn",
						         "type": "string"
						       },
						       "SnsTopicArn": {
						         "maxLength": 256,
						         "pattern": "",
						         "$ref": "#/definitions/Arn",
						         "type": "string"
						       }
						     },
						     "$ref": "#/definitions/SNSConfiguration",
						     "required": [
						       "RoleArn",
						       "SnsTopicArn"
						     ],
						     "type": "object"
						   }
						*/
						Description: "Configuration options for an SNS alert action.",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"role_arn": {
									// Property: RoleArn
									// CloudFormation resource type schema:
									/*
									   {
									     "maxLength": 256,
									     "pattern": "",
									     "$ref": "#/definitions/Arn",
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
								"sns_topic_arn": {
									// Property: SnsTopicArn
									// CloudFormation resource type schema:
									/*
									   {
									     "maxLength": 256,
									     "pattern": "",
									     "$ref": "#/definitions/Arn",
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Required: true,
			// Action is a force-new attribute.
		},
		"alert_description": {
			// Property: AlertDescription
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A description for the alert.",
			     "maxLength": 256,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "A description for the alert.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// AlertDescription is a force-new attribute.
		},
		"alert_name": {
			// Property: AlertName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The name of the alert. If not provided, a name is generated automatically.",
			     "maxLength": 63,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The name of the alert. If not provided, a name is generated automatically.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// AlertName is a force-new attribute.
		},
		"alert_sensitivity_threshold": {
			// Property: AlertSensitivityThreshold
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "A number between 0 and 100 (inclusive) that tunes the sensitivity of the alert.",
			     "type": "integer"
			   }
			*/
			Description: "A number between 0 and 100 (inclusive) that tunes the sensitivity of the alert.",
			Type:        types.NumberType,
			Required:    true,
			// AlertSensitivityThreshold is a force-new attribute.
		},
		"anomaly_detector_arn": {
			// Property: AnomalyDetectorArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon resource name (ARN) of the Anomaly Detector to alert.",
			     "maxLength": 256,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The Amazon resource name (ARN) of the Anomaly Detector to alert.",
			Type:        types.StringType,
			Required:    true,
			// AnomalyDetectorArn is a force-new attribute.
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "maxLength": 256,
			     "pattern": "",
			     "$ref": "#/definitions/Arn",
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::LookoutMetrics::Alert",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::LookoutMetrics::Alert").WithTerraformTypeName("aws_lookoutmetrics_alert").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_lookoutmetrics_alert", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
