---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_route53profiles_profile Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Route53Profiles::Profile
---

# awscc_route53profiles_profile (Data Source)

Data Source schema for AWS::Route53Profiles::Profile



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String) The Amazon Resource Name (ARN) of the resolver profile.
- `client_token` (String) The id of the creator request
- `name` (String) The name of the profile.
- `profile_id` (String) The ID of the profile.
- `tags` (Attributes List) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
