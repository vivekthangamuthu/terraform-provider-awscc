// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ivschat

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ivschat_logging_configuration", loggingConfigurationDataSource)
}

// loggingConfigurationDataSource returns the Terraform awscc_ivschat_logging_configuration data source.
// This Terraform data source corresponds to the CloudFormation AWS::IVSChat::LoggingConfiguration resource.
func loggingConfigurationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "LoggingConfiguration ARN is automatically generated on creation and assigned as the unique identifier.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws:ivschat:[a-z0-9-]+:[0-9]+:logging-configuration/[a-zA-Z0-9-]+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "LoggingConfiguration ARN is automatically generated on creation and assigned as the unique identifier.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DestinationConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Destination configuration for IVS Chat logging.",
		//	  "properties": {
		//	    "CloudWatchLogs": {
		//	      "additionalProperties": false,
		//	      "description": "CloudWatch destination configuration for IVS Chat logging.",
		//	      "properties": {
		//	        "LogGroupName": {
		//	          "description": "Name of the Amazon CloudWatch Logs log group where chat activity will be logged.",
		//	          "maxLength": 512,
		//	          "minLength": 1,
		//	          "pattern": "^[\\.\\-_/#A-Za-z0-9]+$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "LogGroupName"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Firehose": {
		//	      "additionalProperties": false,
		//	      "description": "Kinesis Firehose destination configuration for IVS Chat logging.",
		//	      "properties": {
		//	        "DeliveryStreamName": {
		//	          "description": "Name of the Amazon Kinesis Firehose delivery stream where chat activity will be logged.",
		//	          "maxLength": 64,
		//	          "minLength": 1,
		//	          "pattern": "^[a-zA-Z0-9_.-]+$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "DeliveryStreamName"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "S3": {
		//	      "additionalProperties": false,
		//	      "description": "S3 destination configuration for IVS Chat logging.",
		//	      "properties": {
		//	        "BucketName": {
		//	          "description": "Name of the Amazon S3 bucket where chat activity will be logged.",
		//	          "maxLength": 63,
		//	          "minLength": 3,
		//	          "pattern": "^[a-z0-9-.]+$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "BucketName"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"destination_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CloudWatchLogs
				"cloudwatch_logs": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: LogGroupName
						"log_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Name of the Amazon CloudWatch Logs log group where chat activity will be logged.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "CloudWatch destination configuration for IVS Chat logging.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Firehose
				"firehose": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DeliveryStreamName
						"delivery_stream_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Name of the Amazon Kinesis Firehose delivery stream where chat activity will be logged.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Kinesis Firehose destination configuration for IVS Chat logging.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: S3
				"s3": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: BucketName
						"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Name of the Amazon S3 bucket where chat activity will be logged.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "S3 destination configuration for IVS Chat logging.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Destination configuration for IVS Chat logging.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The system-generated ID of the logging configuration.",
		//	  "maxLength": 12,
		//	  "minLength": 12,
		//	  "pattern": "^[a-zA-Z0-9]+$",
		//	  "type": "string"
		//	}
		"logging_configuration_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The system-generated ID of the logging configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the logging configuration. The value does not need to be unique.",
		//	  "maxLength": 128,
		//	  "minLength": 0,
		//	  "pattern": "^[a-zA-Z0-9-_]*$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the logging configuration. The value does not need to be unique.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The state of the logging configuration. When the state is ACTIVE, the configuration is ready to log chat content.",
		//	  "enum": [
		//	    "CREATING",
		//	    "CREATE_FAILED",
		//	    "DELETING",
		//	    "DELETE_FAILED",
		//	    "UPDATING",
		//	    "UPDATING_FAILED",
		//	    "ACTIVE"
		//	  ],
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The state of the logging configuration. When the state is ACTIVE, the configuration is ready to log chat content.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IVSChat::LoggingConfiguration",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IVSChat::LoggingConfiguration").WithTerraformTypeName("awscc_ivschat_logging_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                       "Arn",
		"bucket_name":               "BucketName",
		"cloudwatch_logs":           "CloudWatchLogs",
		"delivery_stream_name":      "DeliveryStreamName",
		"destination_configuration": "DestinationConfiguration",
		"firehose":                  "Firehose",
		"key":                       "Key",
		"log_group_name":            "LogGroupName",
		"logging_configuration_id":  "Id",
		"name":                      "Name",
		"s3":                        "S3",
		"state":                     "State",
		"tags":                      "Tags",
		"value":                     "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
