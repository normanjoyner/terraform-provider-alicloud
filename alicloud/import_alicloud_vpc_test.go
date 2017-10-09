package alicloud

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccAlicloudVpc_importBasic(t *testing.T) {
	resourceName := "alicloud_vpc.foo"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVpcConfig,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}