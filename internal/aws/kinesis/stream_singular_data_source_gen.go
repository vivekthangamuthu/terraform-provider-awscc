// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package kinesis

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_kinesis_stream", streamDataSource)
}

// streamDataSource returns the Terraform awscc_kinesis_stream data source.
// This Terraform data source corresponds to the CloudFormation AWS::Kinesis::Stream resource.
func streamDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon resource name (ARN) of the Kinesis stream",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon resource name (ARN) of the Kinesis stream",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Kinesis stream.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9_.-]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Kinesis stream.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RetentionPeriodHours
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of hours for the data records that are stored in shards to remain accessible.",
		//	  "minimum": 24,
		//	  "type": "integer"
		//	}
		"retention_period_hours": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of hours for the data records that are stored in shards to remain accessible.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ShardCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of shards that the stream uses. Required when StreamMode = PROVISIONED is passed.",
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"shard_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of shards that the stream uses. Required when StreamMode = PROVISIONED is passed.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StreamEncryption
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream.",
		//	  "properties": {
		//	    "EncryptionType": {
		//	      "description": "The encryption type to use. The only valid value is KMS. ",
		//	      "enum": [
		//	        "KMS"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "KeyId": {
		//	      "anyOf": [
		//	        {},
		//	        {}
		//	      ],
		//	      "description": "The GUID for the customer-managed AWS KMS key to use for encryption. This value can be a globally unique identifier, a fully specified Amazon Resource Name (ARN) to either an alias or a key, or an alias name prefixed by \"alias/\".You can also use a master key owned by Kinesis Data Streams by specifying the alias aws/kinesis.",
		//	      "maxLength": 2048,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "EncryptionType",
		//	    "KeyId"
		//	  ],
		//	  "type": "object"
		//	}
		"stream_encryption": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EncryptionType
				"encryption_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The encryption type to use. The only valid value is KMS. ",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: KeyId
				"key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The GUID for the customer-managed AWS KMS key to use for encryption. This value can be a globally unique identifier, a fully specified Amazon Resource Name (ARN) to either an alias or a key, or an alias name prefixed by \"alias/\".You can also use a master key owned by Kinesis Data Streams by specifying the alias aws/kinesis.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StreamModeDetails
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "default": {
		//	    "StreamMode": "PROVISIONED"
		//	  },
		//	  "description": "The mode in which the stream is running.",
		//	  "properties": {
		//	    "StreamMode": {
		//	      "description": "The mode of the stream",
		//	      "enum": [
		//	        "ON_DEMAND",
		//	        "PROVISIONED"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "StreamMode"
		//	  ],
		//	  "type": "object"
		//	}
		"stream_mode_details": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: StreamMode
				"stream_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The mode of the stream",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The mode in which the stream is running.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An arbitrary set of tags (key–value pairs) to associate with the Kinesis stream.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "An arbitrary set of tags (key-value pairs) to associate with the Kinesis stream.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 255,
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
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An arbitrary set of tags (key–value pairs) to associate with the Kinesis stream.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Kinesis::Stream",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Kinesis::Stream").WithTerraformTypeName("awscc_kinesis_stream")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                    "Arn",
		"encryption_type":        "EncryptionType",
		"key":                    "Key",
		"key_id":                 "KeyId",
		"name":                   "Name",
		"retention_period_hours": "RetentionPeriodHours",
		"shard_count":            "ShardCount",
		"stream_encryption":      "StreamEncryption",
		"stream_mode":            "StreamMode",
		"stream_mode_details":    "StreamModeDetails",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
