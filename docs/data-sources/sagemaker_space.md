---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_sagemaker_space Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::SageMaker::Space
---

# awscc_sagemaker_space (Data Source)

Data Source schema for AWS::SageMaker::Space



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `domain_id` (String) The ID of the associated Domain.
- `space_arn` (String) The space Amazon Resource Name (ARN).
- `space_name` (String) A name for the Space.
- `space_settings` (Attributes) A collection of settings. (see [below for nested schema](#nestedatt--space_settings))
- `tags` (Attributes List) A list of tags to apply to the space. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--space_settings"></a>
### Nested Schema for `space_settings`

Read-Only:

- `jupyter_server_app_settings` (Attributes) The Jupyter server's app settings. (see [below for nested schema](#nestedatt--space_settings--jupyter_server_app_settings))
- `kernel_gateway_app_settings` (Attributes) The kernel gateway app settings. (see [below for nested schema](#nestedatt--space_settings--kernel_gateway_app_settings))

<a id="nestedatt--space_settings--jupyter_server_app_settings"></a>
### Nested Schema for `space_settings.jupyter_server_app_settings`

Read-Only:

- `default_resource_spec` (Attributes) (see [below for nested schema](#nestedatt--space_settings--jupyter_server_app_settings--default_resource_spec))

<a id="nestedatt--space_settings--jupyter_server_app_settings--default_resource_spec"></a>
### Nested Schema for `space_settings.jupyter_server_app_settings.default_resource_spec`

Read-Only:

- `instance_type` (String) The instance type that the image version runs on.
- `sage_maker_image_arn` (String) The ARN of the SageMaker image that the image version belongs to.
- `sage_maker_image_version_arn` (String) The ARN of the image version created on the instance.



<a id="nestedatt--space_settings--kernel_gateway_app_settings"></a>
### Nested Schema for `space_settings.kernel_gateway_app_settings`

Read-Only:

- `custom_images` (Attributes List) A list of custom SageMaker images that are configured to run as a KernelGateway app. (see [below for nested schema](#nestedatt--space_settings--kernel_gateway_app_settings--custom_images))
- `default_resource_spec` (Attributes) The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the KernelGateway app. (see [below for nested schema](#nestedatt--space_settings--kernel_gateway_app_settings--default_resource_spec))

<a id="nestedatt--space_settings--kernel_gateway_app_settings--custom_images"></a>
### Nested Schema for `space_settings.kernel_gateway_app_settings.custom_images`

Read-Only:

- `app_image_config_name` (String) The Name of the AppImageConfig.
- `image_name` (String) The name of the CustomImage. Must be unique to your account.
- `image_version_number` (Number) The version number of the CustomImage.


<a id="nestedatt--space_settings--kernel_gateway_app_settings--default_resource_spec"></a>
### Nested Schema for `space_settings.kernel_gateway_app_settings.default_resource_spec`

Read-Only:

- `instance_type` (String) The instance type that the image version runs on.
- `sage_maker_image_arn` (String) The ARN of the SageMaker image that the image version belongs to.
- `sage_maker_image_version_arn` (String) The ARN of the image version created on the instance.




<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)

