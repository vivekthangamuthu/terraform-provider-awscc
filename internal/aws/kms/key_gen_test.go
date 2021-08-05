// Code generated by generators/resource/main.go; DO NOT EDIT.

package kms_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

func TestAccAWSKMSKey_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::KMS::Key", "aws_kms_key", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
