// Code generated by generators/resource/main.go; DO NOT EDIT.

package logs_test

import (
	"fmt"

	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

type logGroupTest struct{}

func TestAccAWSLogsLogGroup_basic(t *testing.T) {
	data := acctest.NewTestData(t, "AWS::Logs::LogGroup", "aws_logs_log_group", "test")
	r := logGroupTest{}

	data.ResourceTest(t, []resource.TestStep{
		{
			Config: r.basic(data),

			Check: resource.ComposeTestCheckFunc(
				data.CheckExistsInAWS(),
			),
		},
	})
}

func TestAccAWSLogsLogGroup_disappears(t *testing.T) {
	data := acctest.NewTestData(t, "AWS::Logs::LogGroup", "aws_logs_log_group", "test")
	r := logGroupTest{}

	data.ResourceTest(t, []resource.TestStep{
		data.DisappearsStep(acctest.DisappearsStepData{
			Config: r.basic,
		}),
	})
}

func (r logGroupTest) basic(data acctest.TestData) string {
	return fmt.Sprintf(`
resource %[1]q %[2]q {
  provider = cloudapi
}
`, data.TerraformResourceType, data.ResourceLabel)
}
