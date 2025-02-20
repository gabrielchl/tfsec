package compute

import (
	"github.com/aquasecurity/defsec/rules/oracle/compute"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		LegacyID: "OCI001",
		BadExample: []string{`
 resource "opc_compute_ip_address_reservation" "bad_example" {
 	name            = "my-ip-address"
 	ip_address_pool = "public-ippool"
   }
 `},
		GoodExample: []string{`
 resource "opc_compute_ip_address_reservation" "good_example" {
 	name            = "my-ip-address"
 	ip_address_pool = "cloud-ippool"
   }
 `},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/opc/latest/docs/resources/opc_compute_ip_address_reservation",
			"https://registry.terraform.io/providers/hashicorp/opc/latest/docs/resources/opc_compute_instance",
		},
		RequiredTypes:  []string{"resource"},
		RequiredLabels: []string{"opc_compute_ip_address_reservation"},
		Base:           compute.CheckNoPublicIp,
	})
}
