// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package bedrock

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_bedrock_knowledge_base", knowledgeBaseDataSource)
}

// knowledgeBaseDataSource returns the Terraform awscc_bedrock_knowledge_base data source.
// This Terraform data source corresponds to the CloudFormation AWS::Bedrock::KnowledgeBase resource.
func knowledgeBaseDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time at which the knowledge base was created.",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time at which the knowledge base was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Description of the Resource.",
		//	  "maxLength": 200,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Description of the Resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FailureReasons
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of reasons that the API operation on the knowledge base failed.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "Failure Reason for Error.",
		//	    "maxLength": 2048,
		//	    "type": "string"
		//	  },
		//	  "maxItems": 2048,
		//	  "type": "array"
		//	}
		"failure_reasons": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of reasons that the API operation on the knowledge base failed.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KnowledgeBaseArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the knowledge base.",
		//	  "maxLength": 128,
		//	  "minLength": 0,
		//	  "pattern": "^arn:aws(|-cn|-us-gov):bedrock:[a-zA-Z0-9-]*:[0-9]{12}:knowledge-base/[0-9a-zA-Z]+$",
		//	  "type": "string"
		//	}
		"knowledge_base_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the knowledge base.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KnowledgeBaseConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Contains details about the embeddings model used for the knowledge base.",
		//	  "properties": {
		//	    "Type": {
		//	      "description": "The type of a knowledge base.",
		//	      "enum": [
		//	        "VECTOR"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "VectorKnowledgeBaseConfiguration": {
		//	      "additionalProperties": false,
		//	      "description": "Contains details about the model used to create vector embeddings for the knowledge base.",
		//	      "properties": {
		//	        "EmbeddingModelArn": {
		//	          "description": "The ARN of the model used to create vector embeddings for the knowledge base.",
		//	          "maxLength": 1011,
		//	          "minLength": 20,
		//	          "pattern": "^arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:(([0-9]{12}:custom-model/[a-z0-9-]{1,63}[.]{1}[a-z0-9-]{1,63}/[a-z0-9]{12})|(:foundation-model/[a-z0-9-]{1,63}[.]{1}[a-z0-9-]{1,63}))$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "EmbeddingModelArn"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "Type",
		//	    "VectorKnowledgeBaseConfiguration"
		//	  ],
		//	  "type": "object"
		//	}
		"knowledge_base_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The type of a knowledge base.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: VectorKnowledgeBaseConfiguration
				"vector_knowledge_base_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: EmbeddingModelArn
						"embedding_model_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ARN of the model used to create vector embeddings for the knowledge base.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Contains details about the model used to create vector embeddings for the knowledge base.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Contains details about the embeddings model used for the knowledge base.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KnowledgeBaseId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique identifier of the knowledge base.",
		//	  "pattern": "^[0-9a-zA-Z]{10}$",
		//	  "type": "string"
		//	}
		"knowledge_base_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique identifier of the knowledge base.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the knowledge base.",
		//	  "pattern": "^([0-9a-zA-Z][_-]?){1,100}$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the knowledge base.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the IAM role with permissions to invoke API operations on the knowledge base. The ARN must begin with AmazonBedrockExecutionRoleForKnowledgeBase_",
		//	  "maxLength": 2048,
		//	  "pattern": "^arn:aws(-[^:]+)?:iam::([0-9]{12})?:role/.+$",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the IAM role with permissions to invoke API operations on the knowledge base. The ARN must begin with AmazonBedrockExecutionRoleForKnowledgeBase_",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of a knowledge base.",
		//	  "enum": [
		//	    "CREATING",
		//	    "ACTIVE",
		//	    "DELETING",
		//	    "UPDATING",
		//	    "FAILED"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of a knowledge base.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StorageConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The vector store service in which the knowledge base is stored.",
		//	  "oneOf": [
		//	    {
		//	      "required": [
		//	        "OpensearchServerlessConfiguration"
		//	      ]
		//	    },
		//	    {
		//	      "required": [
		//	        "PineconeConfiguration"
		//	      ]
		//	    },
		//	    {
		//	      "required": [
		//	        "RdsConfiguration"
		//	      ]
		//	    }
		//	  ],
		//	  "properties": {
		//	    "OpensearchServerlessConfiguration": {
		//	      "additionalProperties": false,
		//	      "description": "Contains the storage configuration of the knowledge base in Amazon OpenSearch Service.",
		//	      "properties": {
		//	        "CollectionArn": {
		//	          "description": "The ARN of the OpenSearch Service vector store.",
		//	          "maxLength": 2048,
		//	          "pattern": "^arn:aws:aoss:[a-z]{2}(-gov)?-[a-z]+-\\d{1}:\\d{12}:collection/[a-z0-9-]{3,32}$",
		//	          "type": "string"
		//	        },
		//	        "FieldMapping": {
		//	          "additionalProperties": false,
		//	          "description": "A mapping of Bedrock Knowledge Base fields to OpenSearch Serverless field names",
		//	          "properties": {
		//	            "MetadataField": {
		//	              "description": "The name of the field in which Amazon Bedrock stores metadata about the vector store.",
		//	              "maxLength": 2048,
		//	              "pattern": "^.*$",
		//	              "type": "string"
		//	            },
		//	            "TextField": {
		//	              "description": "The name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.",
		//	              "maxLength": 2048,
		//	              "pattern": "^.*$",
		//	              "type": "string"
		//	            },
		//	            "VectorField": {
		//	              "description": "The name of the field in which Amazon Bedrock stores the vector embeddings for your data sources.",
		//	              "maxLength": 2048,
		//	              "pattern": "^.*$",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "MetadataField",
		//	            "TextField",
		//	            "VectorField"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "VectorIndexName": {
		//	          "description": "The name of the vector store.",
		//	          "maxLength": 2048,
		//	          "pattern": "^.*$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "CollectionArn",
		//	        "FieldMapping",
		//	        "VectorIndexName"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "PineconeConfiguration": {
		//	      "additionalProperties": false,
		//	      "description": "Contains the storage configuration of the knowledge base in Pinecone.",
		//	      "properties": {
		//	        "ConnectionString": {
		//	          "description": "The endpoint URL for your index management page.",
		//	          "maxLength": 2048,
		//	          "pattern": "^.*$",
		//	          "type": "string"
		//	        },
		//	        "CredentialsSecretArn": {
		//	          "description": "The ARN of the secret that you created in AWS Secrets Manager that is linked to your Pinecone API key.",
		//	          "pattern": "^arn:aws(|-cn|-us-gov):secretsmanager:[a-z0-9-]{1,20}:([0-9]{12}|):secret:[a-zA-Z0-9!/_+=.@-]{1,512}$",
		//	          "type": "string"
		//	        },
		//	        "FieldMapping": {
		//	          "additionalProperties": false,
		//	          "description": "Contains the names of the fields to which to map information about the vector store.",
		//	          "properties": {
		//	            "MetadataField": {
		//	              "description": "The name of the field in which Amazon Bedrock stores metadata about the vector store.",
		//	              "maxLength": 2048,
		//	              "pattern": "^.*$",
		//	              "type": "string"
		//	            },
		//	            "TextField": {
		//	              "description": "The name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.",
		//	              "maxLength": 2048,
		//	              "pattern": "^.*$",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "MetadataField",
		//	            "TextField"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "Namespace": {
		//	          "description": "The namespace to be used to write new data to your database.",
		//	          "maxLength": 2048,
		//	          "pattern": "^.*$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "ConnectionString",
		//	        "CredentialsSecretArn",
		//	        "FieldMapping"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "RdsConfiguration": {
		//	      "additionalProperties": false,
		//	      "description": "Contains details about the storage configuration of the knowledge base in Amazon RDS. For more information, see Create a vector index in Amazon RDS.",
		//	      "properties": {
		//	        "CredentialsSecretArn": {
		//	          "description": "The ARN of the secret that you created in AWS Secrets Manager that is linked to your Amazon RDS database.",
		//	          "pattern": "^arn:aws(|-cn|-us-gov):secretsmanager:[a-z0-9-]{1,20}:([0-9]{12}|):secret:[a-zA-Z0-9!/_+=.@-]{1,512}$",
		//	          "type": "string"
		//	        },
		//	        "DatabaseName": {
		//	          "description": "The name of your Amazon RDS database.",
		//	          "maxLength": 63,
		//	          "pattern": "^[a-zA-Z0-9_\\-]+$",
		//	          "type": "string"
		//	        },
		//	        "FieldMapping": {
		//	          "additionalProperties": false,
		//	          "description": "Contains the names of the fields to which to map information about the vector store.",
		//	          "properties": {
		//	            "MetadataField": {
		//	              "description": "The name of the field in which Amazon Bedrock stores metadata about the vector store.",
		//	              "maxLength": 63,
		//	              "pattern": "^[a-zA-Z0-9_\\-]+$",
		//	              "type": "string"
		//	            },
		//	            "PrimaryKeyField": {
		//	              "description": "The name of the field in which Amazon Bedrock stores the ID for each entry.",
		//	              "maxLength": 63,
		//	              "pattern": "^[a-zA-Z0-9_\\-]+$",
		//	              "type": "string"
		//	            },
		//	            "TextField": {
		//	              "description": "The name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.",
		//	              "maxLength": 63,
		//	              "pattern": "^[a-zA-Z0-9_\\-]+$",
		//	              "type": "string"
		//	            },
		//	            "VectorField": {
		//	              "description": "The name of the field in which Amazon Bedrock stores the vector embeddings for your data sources.",
		//	              "maxLength": 63,
		//	              "pattern": "^[a-zA-Z0-9_\\-]+$",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "MetadataField",
		//	            "PrimaryKeyField",
		//	            "TextField",
		//	            "VectorField"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "ResourceArn": {
		//	          "description": "The ARN of the vector store.",
		//	          "pattern": "^arn:aws(|-cn|-us-gov):rds:[a-zA-Z0-9-]*:[0-9]{12}:cluster:[a-zA-Z0-9-]{1,63}$",
		//	          "type": "string"
		//	        },
		//	        "TableName": {
		//	          "description": "The name of the table in the database.",
		//	          "maxLength": 63,
		//	          "pattern": "^[a-zA-Z0-9_\\.\\-]+$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "CredentialsSecretArn",
		//	        "DatabaseName",
		//	        "FieldMapping",
		//	        "ResourceArn",
		//	        "TableName"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Type": {
		//	      "description": "The storage type of a knowledge base.",
		//	      "enum": [
		//	        "OPENSEARCH_SERVERLESS",
		//	        "PINECONE",
		//	        "RDS"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Type"
		//	  ],
		//	  "type": "object"
		//	}
		"storage_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: OpensearchServerlessConfiguration
				"opensearch_serverless_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CollectionArn
						"collection_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ARN of the OpenSearch Service vector store.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: FieldMapping
						"field_mapping": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: MetadataField
								"metadata_field": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The name of the field in which Amazon Bedrock stores metadata about the vector store.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: TextField
								"text_field": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: VectorField
								"vector_field": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The name of the field in which Amazon Bedrock stores the vector embeddings for your data sources.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "A mapping of Bedrock Knowledge Base fields to OpenSearch Serverless field names",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: VectorIndexName
						"vector_index_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The name of the vector store.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Contains the storage configuration of the knowledge base in Amazon OpenSearch Service.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: PineconeConfiguration
				"pinecone_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ConnectionString
						"connection_string": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The endpoint URL for your index management page.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: CredentialsSecretArn
						"credentials_secret_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ARN of the secret that you created in AWS Secrets Manager that is linked to your Pinecone API key.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: FieldMapping
						"field_mapping": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: MetadataField
								"metadata_field": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The name of the field in which Amazon Bedrock stores metadata about the vector store.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: TextField
								"text_field": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Contains the names of the fields to which to map information about the vector store.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Namespace
						"namespace": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The namespace to be used to write new data to your database.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Contains the storage configuration of the knowledge base in Pinecone.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RdsConfiguration
				"rds_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CredentialsSecretArn
						"credentials_secret_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ARN of the secret that you created in AWS Secrets Manager that is linked to your Amazon RDS database.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: DatabaseName
						"database_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The name of your Amazon RDS database.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: FieldMapping
						"field_mapping": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: MetadataField
								"metadata_field": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The name of the field in which Amazon Bedrock stores metadata about the vector store.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: PrimaryKeyField
								"primary_key_field": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The name of the field in which Amazon Bedrock stores the ID for each entry.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: TextField
								"text_field": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: VectorField
								"vector_field": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The name of the field in which Amazon Bedrock stores the vector embeddings for your data sources.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Contains the names of the fields to which to map information about the vector store.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: ResourceArn
						"resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ARN of the vector store.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: TableName
						"table_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The name of the table in the database.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Contains details about the storage configuration of the knowledge base in Amazon RDS. For more information, see Create a vector index in Amazon RDS.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The storage type of a knowledge base.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The vector store service in which the knowledge base is stored.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A map of tag keys and values",
		//	  "patternProperties": {
		//	    "": {
		//	      "description": "Value of a tag",
		//	      "maxLength": 256,
		//	      "minLength": 0,
		//	      "pattern": "^[a-zA-Z0-9\\s._:/=+@-]*$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A map of tag keys and values",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time at which the knowledge base was last updated.",
		//	  "type": "string"
		//	}
		"updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time at which the knowledge base was last updated.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Bedrock::KnowledgeBase",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Bedrock::KnowledgeBase").WithTerraformTypeName("awscc_bedrock_knowledge_base")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"collection_arn":                      "CollectionArn",
		"connection_string":                   "ConnectionString",
		"created_at":                          "CreatedAt",
		"credentials_secret_arn":              "CredentialsSecretArn",
		"database_name":                       "DatabaseName",
		"description":                         "Description",
		"embedding_model_arn":                 "EmbeddingModelArn",
		"failure_reasons":                     "FailureReasons",
		"field_mapping":                       "FieldMapping",
		"knowledge_base_arn":                  "KnowledgeBaseArn",
		"knowledge_base_configuration":        "KnowledgeBaseConfiguration",
		"knowledge_base_id":                   "KnowledgeBaseId",
		"metadata_field":                      "MetadataField",
		"name":                                "Name",
		"namespace":                           "Namespace",
		"opensearch_serverless_configuration": "OpensearchServerlessConfiguration",
		"pinecone_configuration":              "PineconeConfiguration",
		"primary_key_field":                   "PrimaryKeyField",
		"rds_configuration":                   "RdsConfiguration",
		"resource_arn":                        "ResourceArn",
		"role_arn":                            "RoleArn",
		"status":                              "Status",
		"storage_configuration":               "StorageConfiguration",
		"table_name":                          "TableName",
		"tags":                                "Tags",
		"text_field":                          "TextField",
		"type":                                "Type",
		"updated_at":                          "UpdatedAt",
		"vector_field":                        "VectorField",
		"vector_index_name":                   "VectorIndexName",
		"vector_knowledge_base_configuration": "VectorKnowledgeBaseConfiguration",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}