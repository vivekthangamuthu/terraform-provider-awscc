// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package networkmanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_networkmanager_link", linkDataSource)
}

// linkDataSource returns the Terraform awscc_networkmanager_link data source.
// This Terraform data source corresponds to the CloudFormation AWS::NetworkManager::Link resource.
func linkDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Bandwidth
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The Bandwidth for the link.",
		//	  "properties": {
		//	    "DownloadSpeed": {
		//	      "description": "Download speed in Mbps.",
		//	      "type": "integer"
		//	    },
		//	    "UploadSpeed": {
		//	      "description": "Upload speed in Mbps.",
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"bandwidth": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DownloadSpeed
				"download_speed": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "Download speed in Mbps.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: UploadSpeed
				"upload_speed": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "Upload speed in Mbps.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The Bandwidth for the link.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date and time that the device was created.",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date and time that the device was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the link.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the link.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GlobalNetworkId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the global network.",
		//	  "type": "string"
		//	}
		"global_network_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the global network.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LinkArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the link.",
		//	  "type": "string"
		//	}
		"link_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the link.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LinkId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the link.",
		//	  "type": "string"
		//	}
		"link_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the link.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Provider
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The provider of the link.",
		//	  "type": "string"
		//	}
		"provider_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The provider of the link.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SiteId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the site",
		//	  "type": "string"
		//	}
		"site_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the site",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The state of the link.",
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The state of the link.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags for the link.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a link resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
			Description: "The tags for the link.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of the link.",
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of the link.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::NetworkManager::Link",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkManager::Link").WithTerraformTypeName("awscc_networkmanager_link")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bandwidth":         "Bandwidth",
		"created_at":        "CreatedAt",
		"description":       "Description",
		"download_speed":    "DownloadSpeed",
		"global_network_id": "GlobalNetworkId",
		"key":               "Key",
		"link_arn":          "LinkArn",
		"link_id":           "LinkId",
		"provider_name":     "Provider",
		"site_id":           "SiteId",
		"state":             "State",
		"tags":              "Tags",
		"type":              "Type",
		"upload_speed":      "UploadSpeed",
		"value":             "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
