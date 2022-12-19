// Code generated by generators/resource/main.go; DO NOT EDIT.

package opensearchserverless_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSOpenSearchServerlessVpcEndpoint_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::OpenSearchServerless::VpcEndpoint", "awscc_opensearchserverless_vpc_endpoint", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}