---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_vpn_gateway Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Schema for EC2 VPN Gateway
---

# awscc_ec2_vpn_gateway (Resource)

Schema for EC2 VPN Gateway



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `type` (String) The type of VPN connection the virtual private gateway supports.

### Optional

- `amazon_side_asn` (Number) The private Autonomous System Number (ASN) for the Amazon side of a BGP session.
- `tags` (Attributes List) Any tags assigned to the virtual private gateway. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `vpn_gateway_id` (String) VPN Gateway ID generated by service

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_ec2_vpn_gateway.example <resource ID>
```