// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package ec2_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSEC2DHCPOptionsDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::EC2::DHCPOptions", "awscc_ec2_dhcp_optionses", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyDataSourceConfig(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(fmt.Sprintf("data.%s", td.ResourceName), "ids.#"),
			),
		},
	})
}
