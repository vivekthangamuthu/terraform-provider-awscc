---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_sagemaker_image_version Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::SageMaker::ImageVersion
---

# awscc_sagemaker_image_version (Data Source)

Data Source schema for AWS::SageMaker::ImageVersion



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **base_image** (String) The registry path of the container image on which this image version is based.
- **container_image** (String) The registry path of the container image that contains this image version.
- **image_arn** (String) The Amazon Resource Name (ARN) of the parent image.
- **image_name** (String) The name of the image this version belongs to.
- **image_version_arn** (String) The Amazon Resource Name (ARN) of the image version.
- **version** (Number) The version number of the image version.

