---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_connect_user Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Connect::User
---

# awscc_connect_user (Resource)

Resource Type definition for AWS::Connect::User



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **instance_arn** (String) The identifier of the Amazon Connect instance.
- **phone_config** (Attributes) The phone settings for the user. (see [below for nested schema](#nestedatt--phone_config))
- **routing_profile_arn** (String) The identifier of the routing profile for the user.
- **security_profile_arns** (Set of String) One or more security profile arns for the user
- **username** (String) The user name for the account.

### Optional

- **directory_user_id** (String) The identifier of the user account in the directory used for identity management.
- **hierarchy_group_arn** (String) The identifier of the hierarchy group for the user.
- **identity_info** (Attributes) The information about the identity of the user. (see [below for nested schema](#nestedatt--identity_info))
- **password** (String) The password for the user account. A password is required if you are using Amazon Connect for identity management. Otherwise, it is an error to include a password.
- **tags** (Attributes Set) One or more tags. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **id** (String) Uniquely identifies the resource.
- **user_arn** (String) The Amazon Resource Name (ARN) for the user.

<a id="nestedatt--phone_config"></a>
### Nested Schema for `phone_config`

Required:

- **after_contact_work_time_limit** (Number) The After Call Work (ACW) timeout setting, in seconds.
- **auto_accept** (Boolean) The Auto accept setting.
- **desk_phone_number** (String) The phone number for the user's desk phone.
- **phone_type** (String) The phone type.


<a id="nestedatt--identity_info"></a>
### Nested Schema for `identity_info`

Optional:

- **email** (String) The email address. If you are using SAML for identity management and include this parameter, an error is returned.
- **first_name** (String) The first name. This is required if you are using Amazon Connect or SAML for identity management.
- **last_name** (String) The last name. This is required if you are using Amazon Connect or SAML for identity management.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- **value** (String) The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_connect_user.example <resource ID>
```