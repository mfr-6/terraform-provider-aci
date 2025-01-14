
resource "aci_annotation" "full_example_application_epg" {
  parent_dn = aci_application_epg.example.id
  key       = "test_key"
  value     = "test_value"
}

resource "aci_annotation" "full_example_tenant" {
  parent_dn = aci_tenant.example.id
  key       = "test_key"
  value     = "test_value"
}
