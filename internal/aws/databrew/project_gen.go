// Code generated by generators/resource/main.go; DO NOT EDIT.

package databrew

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
	registry.AddResourceTypeFactory("aws_databrew_project", projectResourceType)
}

// projectResourceType returns the Terraform aws_databrew_project resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::DataBrew::Project resource type.
func projectResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"dataset_name": {
			// Property: DatasetName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Dataset name",
			     "maxLength": 255,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "Dataset name",
			Type:        types.StringType,
			Required:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Project name",
			     "maxLength": 255,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "Project name",
			Type:        types.StringType,
			Required:    true,
			// Name is a force-new attribute.
		},
		"recipe_name": {
			// Property: RecipeName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Recipe name",
			     "maxLength": 255,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "Recipe name",
			Type:        types.StringType,
			Required:    true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Role arn",
			     "type": "string"
			   }
			*/
			Description: "Role arn",
			Type:        types.StringType,
			Required:    true,
		},
		"sample": {
			// Property: Sample
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "Size": {
			         "description": "Sample size",
			         "type": "integer"
			       },
			       "Type": {
			         "description": "Sample type",
			         "enum": [
			           "FIRST_N",
			           "LAST_N",
			           "RANDOM"
			         ],
			         "type": "string"
			       }
			     },
			     "$ref": "#/definitions/Sample",
			     "required": [
			       "Type"
			     ],
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"size": {
						// Property: Size
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Sample size",
						     "type": "integer"
						   }
						*/
						Description: "Sample size",
						Type:        types.NumberType,
						Optional:    true,
					},
					"type": {
						// Property: Type
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Sample type",
						     "enum": [
						       "FIRST_N",
						       "LAST_N",
						       "RANDOM"
						     ],
						     "type": "string"
						   }
						*/
						Description: "Sample type",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Optional: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "insertionOrder": false,
			     "items": {
			       "additionalProperties": false,
			       "description": "A key-value pair to associate with a resource.",
			       "properties": {
			         "Key": {
			           "maxLength": 128,
			           "minLength": 1,
			           "type": "string"
			         },
			         "Value": {
			           "maxLength": 256,
			           "minLength": 0,
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
			// Multiset.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 128,
						     "minLength": 1,
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
						     "maxLength": 256,
						     "minLength": 0,
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
			Computed: true,
			// Tags is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::DataBrew::Project.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataBrew::Project").WithTerraformTypeName("aws_databrew_project").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_databrew_project", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
