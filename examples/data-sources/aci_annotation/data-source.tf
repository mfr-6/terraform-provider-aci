
data "aci_annotation" "example_application_epg" {
  parent_dn = aci_application_epg.example.id
  key       = "test_key"
}

data "aci_annotation" "example_tenant" {
  parent_dn = aci_tenant.example.id
  key       = "test_key"
}
