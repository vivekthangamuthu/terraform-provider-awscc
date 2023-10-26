---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_entityresolution_id_mapping_workflow Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  IdMappingWorkflow defined in AWS Entity Resolution service
---

# awscc_entityresolution_id_mapping_workflow (Resource)

IdMappingWorkflow defined in AWS Entity Resolution service



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id_mapping_techniques` (Attributes) (see [below for nested schema](#nestedatt--id_mapping_techniques))
- `input_source_config` (Attributes List) (see [below for nested schema](#nestedatt--input_source_config))
- `output_source_config` (Attributes List) (see [below for nested schema](#nestedatt--output_source_config))
- `role_arn` (String)
- `workflow_name` (String) The name of the IdMappingWorkflow

### Optional

- `description` (String) The description of the IdMappingWorkflow
- `tags` (Attributes Set) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `created_at` (String) The time of this IdMappingWorkflow got created
- `id` (String) Uniquely identifies the resource.
- `updated_at` (String) The time of this IdMappingWorkflow got last updated at
- `workflow_arn` (String) The default IdMappingWorkflow arn

<a id="nestedatt--id_mapping_techniques"></a>
### Nested Schema for `id_mapping_techniques`

Optional:

- `id_mapping_type` (String)
- `provider_properties` (Attributes) (see [below for nested schema](#nestedatt--id_mapping_techniques--provider_properties))

<a id="nestedatt--id_mapping_techniques--provider_properties"></a>
### Nested Schema for `id_mapping_techniques.provider_properties`

Required:

- `provider_service_arn` (String) Arn of the Provider Service being used.

Optional:

- `intermediate_source_configuration` (Attributes) (see [below for nested schema](#nestedatt--id_mapping_techniques--provider_properties--intermediate_source_configuration))
- `provider_configuration` (Map of String) Additional Provider configuration that would be required for the provider service. The Configuration must be in JSON string format

<a id="nestedatt--id_mapping_techniques--provider_properties--intermediate_source_configuration"></a>
### Nested Schema for `id_mapping_techniques.provider_properties.intermediate_source_configuration`

Required:

- `intermediate_s3_path` (String) The s3 path that would be used to stage the intermediate data being generated during workflow execution.




<a id="nestedatt--input_source_config"></a>
### Nested Schema for `input_source_config`

Required:

- `input_source_arn` (String) An Glue table ARN for the input source table
- `schema_arn` (String) The SchemaMapping arn associated with the Schema


<a id="nestedatt--output_source_config"></a>
### Nested Schema for `output_source_config`

Required:

- `output_s3_path` (String) The S3 path to which Entity Resolution will write the output table

Optional:

- `kms_arn` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_entityresolution_id_mapping_workflow.example <resource ID>
```