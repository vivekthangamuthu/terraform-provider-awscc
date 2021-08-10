// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53recoverycontrol

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
	registry.AddResourceTypeFactory("aws_route53recoverycontrol_cluster", clusterResourceType)
}

// clusterResourceType returns the Terraform aws_route53recoverycontrol_cluster resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Route53RecoveryControl::Cluster resource type.
func clusterResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"cluster_arn": {
			// Property: ClusterArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the cluster.",
			     "maxLength": 2048,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the cluster.",
			Type:        types.StringType,
			Computed:    true,
		},
		"cluster_endpoints": {
			// Property: ClusterEndpoints
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Endpoints for the cluster.",
			     "insertionOrder": false,
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "Endpoint": {
			           "maxLength": 128,
			           "minLength": 1,
			           "type": "string"
			         },
			         "Region": {
			           "maxLength": 32,
			           "minLength": 1,
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/ClusterEndpoint",
			       "type": "object"
			     },
			     "type": "array"
			   }
			*/
			Description: "Endpoints for the cluster.",
			// Multiset.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"endpoint": {
						// Property: Endpoint
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 128,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"region": {
						// Property: Region
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 32,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Name of a Cluster. You can use any non-white space character in the name",
			     "maxLength": 64,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "Name of a Cluster. You can use any non-white space character in the name",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// Name is a force-new attribute.
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Deployment status of a resource. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.",
			     "enum": [
			       "PENDING",
			       "DEPLOYED",
			       "PENDING_DELETION"
			     ],
			     "type": "string"
			   }
			*/
			Description: "Deployment status of a resource. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "AWS Route53 Recovery Control Cluster resource schema",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53RecoveryControl::Cluster").WithTerraformTypeName("aws_route53recoverycontrol_cluster").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_route53recoverycontrol_cluster", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
