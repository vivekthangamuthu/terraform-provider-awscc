// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package timestream

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_timestream_database", databaseDataSourceType)
}

// databaseDataSourceType returns the Terraform awscc_timestream_database data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Timestream::Database resource type.
func databaseDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"database_name": {
			// Property: DatabaseName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name for the database. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the database name.",
			//   "pattern": "^[a-zA-Z0-9_.-]{3,256}$",
			//   "type": "string"
			// }
			Description: "The name for the database. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the database name.",
			Type:        types.StringType,
			Computed:    true,
		},
		"kms_key_id": {
			// Property: KmsKeyId
			// CloudFormation resource type schema:
			// {
			//   "description": "The KMS key for the database. If the KMS key is not specified, the database will be encrypted with a Timestream managed KMS key located in your account.",
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The KMS key for the database. If the KMS key is not specified, the database will be encrypted with a Timestream managed KMS key located in your account.",
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
			//     "description": "You can use the Resource Tags property to apply tags to resources, which can help you identify and categorize those resources.",
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
			//   "maxItems": 200,
			//   "type": "array"
			// }
			Description: "An array of key-value pairs to apply to this resource.",
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
		Description: "Data Source schema for AWS::Timestream::Database",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Timestream::Database").WithTerraformTypeName("awscc_timestream_database")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":           "Arn",
		"database_name": "DatabaseName",
		"key":           "Key",
		"kms_key_id":    "KmsKeyId",
		"tags":          "Tags",
		"value":         "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
