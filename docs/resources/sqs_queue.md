---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_sqs_queue Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::SQS::Queue
---

# awscc_sqs_queue (Resource)

Resource Type definition for AWS::SQS::Queue



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `content_based_deduplication` (Boolean) For first-in-first-out (FIFO) queues, specifies whether to enable content-based deduplication. During the deduplication interval, Amazon SQS treats messages that are sent with identical content as duplicates and delivers only one copy of the message.
- `deduplication_scope` (String) Specifies whether message deduplication occurs at the message group or queue level. Valid values are messageGroup and queue.
- `delay_seconds` (Number) The time in seconds for which the delivery of all messages in the queue is delayed. You can specify an integer value of 0 to 900 (15 minutes). The default value is 0.
- `fifo_queue` (Boolean) If set to true, creates a FIFO queue. If you don't specify this property, Amazon SQS creates a standard queue.
- `fifo_throughput_limit` (String) Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are perQueue and perMessageGroupId. The perMessageGroupId value is allowed only when the value for DeduplicationScope is messageGroup.
- `kms_data_key_reuse_period_seconds` (Number) The length of time in seconds for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. The value must be an integer between 60 (1 minute) and 86,400 (24 hours). The default is 300 (5 minutes).
- `kms_master_key_id` (String) The ID of an AWS managed customer master key (CMK) for Amazon SQS or a custom CMK. To use the AWS managed CMK for Amazon SQS, specify the (default) alias alias/aws/sqs.
- `maximum_message_size` (Number) The limit of how many bytes that a message can contain before Amazon SQS rejects it. You can specify an integer value from 1,024 bytes (1 KiB) to 262,144 bytes (256 KiB). The default value is 262,144 (256 KiB).
- `message_retention_period` (Number) The number of seconds that Amazon SQS retains a message. You can specify an integer value from 60 seconds (1 minute) to 1,209,600 seconds (14 days). The default value is 345,600 seconds (4 days).
- `queue_name` (String) A name for the queue. To create a FIFO queue, the name of your FIFO queue must end with the .fifo suffix.
- `receive_message_wait_time_seconds` (Number) Specifies the duration, in seconds, that the ReceiveMessage action call waits until a message is in the queue in order to include it in the response, rather than returning an empty response if a message isn't yet available. You can specify an integer from 1 to 20. Short polling is used as the default or when you specify 0 for this property.
- `redrive_allow_policy` (String) The string that includes the parameters for the permissions for the dead-letter queue redrive permission and which source queues can specify dead-letter queues as a JSON object.
- `redrive_policy` (String) A string that includes the parameters for the dead-letter queue functionality (redrive policy) of the source queue.
- `tags` (Attributes List) The tags that you attach to this queue. (see [below for nested schema](#nestedatt--tags))
- `visibility_timeout` (Number) The length of time during which a message will be unavailable after a message is delivered from the queue. This blocks other components from receiving the same message and gives the initial component time to process and delete the message from the queue. Values must be from 0 to 43,200 seconds (12 hours). If you don't specify a value, AWS CloudFormation uses the default value of 30 seconds.

### Read-Only

- `arn` (String) Amazon Resource Name (ARN) of the queue.
- `id` (String) Uniquely identifies the resource.
- `queue_url` (String) URL of the source queue.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_sqs_queue.example <resource ID>
```