// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceL3extRsRedistributePolWithL3extOut(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigL3extRsRedistributePolMinDependencyWithL3extOut,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "route_control_profile_name", "test_tn_rtctrl_profile_name"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "source", "direct"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigL3extRsRedistributePolAllDependencyWithL3extOut,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "route_control_profile_name", "test_tn_rtctrl_profile_name"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "source", "direct"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigL3extRsRedistributePolMinDependencyWithL3extOut,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "route_control_profile_name", "test_tn_rtctrl_profile_name"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "source", "direct"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigL3extRsRedistributePolResetDependencyWithL3extOut,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "route_control_profile_name", "test_tn_rtctrl_profile_name"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "source", "direct"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_l3out_redistribute_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "route_control_profile_name", "test_tn_rtctrl_profile_name"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "source", "direct"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Update with children
			{
				Config:             testConfigL3extRsRedistributePolChildrenDependencyWithL3extOut,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "route_control_profile_name", "test_tn_rtctrl_profile_name"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "source", "direct"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotations.0.key", "annotations_1"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotations.1.key", "annotations_2"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotations.1.value", "value_2"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_l3out_redistribute_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "route_control_profile_name", "test_tn_rtctrl_profile_name"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "source", "direct"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotations.0.key", "annotations_1"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotations.1.key", "annotations_2"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotations.1.value", "value_2"),
				),
			},
			// Update with children removed from config
			{
				Config:             testConfigL3extRsRedistributePolChildrenRemoveFromConfigDependencyWithL3extOut,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotations.0.key", "annotations_1"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotations.1.key", "annotations_2"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotations.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigL3extRsRedistributePolChildrenRemoveOneDependencyWithL3extOut,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotations.0.key", "annotations_2"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotations.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotations.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigL3extRsRedistributePolChildrenRemoveAllDependencyWithL3extOut,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_redistribute_policy.test", "annotations.#", "0"),
				),
			},
		},
	})
}

const testConfigL3extRsRedistributePolMinDependencyWithL3extOut = testConfigL3extOutMinDependencyWithFvTenant + `
resource "aci_l3out_redistribute_policy" "test" {
  parent_dn = aci_l3_outside.test.id
  route_control_profile_name = "test_tn_rtctrl_profile_name"
  source = "direct"
}
`

const testConfigL3extRsRedistributePolAllDependencyWithL3extOut = testConfigL3extOutMinDependencyWithFvTenant + `
resource "aci_l3out_redistribute_policy" "test" {
  parent_dn = aci_l3_outside.test.id
  route_control_profile_name = "test_tn_rtctrl_profile_name"
  annotation = "annotation"
  source = "direct"
}
`

const testConfigL3extRsRedistributePolResetDependencyWithL3extOut = testConfigL3extOutMinDependencyWithFvTenant + `
resource "aci_l3out_redistribute_policy" "test" {
  parent_dn = aci_l3_outside.test.id
  route_control_profile_name = "test_tn_rtctrl_profile_name"
  annotation = "orchestrator:terraform"
  source = "direct"
}
`
const testConfigL3extRsRedistributePolChildrenDependencyWithL3extOut = testConfigL3extOutMinDependencyWithFvTenant + `
resource "aci_l3out_redistribute_policy" "test" {
  parent_dn = aci_l3_outside.test.id
  route_control_profile_name = "test_tn_rtctrl_profile_name"
  source = "direct"
  annotations = [
	{
	  key = "annotations_1"
	  value = "value_1"
	},
	{
	  key = "annotations_2"
	  value = "value_2"
	},
  ]
}
`

const testConfigL3extRsRedistributePolChildrenRemoveFromConfigDependencyWithL3extOut = testConfigL3extOutMinDependencyWithFvTenant + `
resource "aci_l3out_redistribute_policy" "test" {
  parent_dn = aci_l3_outside.test.id
  route_control_profile_name = "test_tn_rtctrl_profile_name"
  source = "direct"
}
`

const testConfigL3extRsRedistributePolChildrenRemoveOneDependencyWithL3extOut = testConfigL3extOutMinDependencyWithFvTenant + `
resource "aci_l3out_redistribute_policy" "test" {
  parent_dn = aci_l3_outside.test.id
  route_control_profile_name = "test_tn_rtctrl_profile_name"
  source = "direct"
  annotations = [ 
	{
	  key = "annotations_2"
	  value = "value_2"
	},
  ]
}
`

const testConfigL3extRsRedistributePolChildrenRemoveAllDependencyWithL3extOut = testConfigL3extOutMinDependencyWithFvTenant + `
resource "aci_l3out_redistribute_policy" "test" {
  parent_dn = aci_l3_outside.test.id
  route_control_profile_name = "test_tn_rtctrl_profile_name"
  source = "direct"
  annotations = []
}
`
