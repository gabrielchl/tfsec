package compute

import (
	"github.com/aquasecurity/defsec/rules/digitalocean/compute"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		LegacyID: "DIG001",
		BadExample: []string{`
 resource "digitalocean_firewall" "bad_example" {
 	name = "only-22-80-and-443"
   
 	droplet_ids = [digitalocean_droplet.web.id]
   
 	inbound_rule {
 	  protocol         = "tcp"
 	  port_range       = "22"
 	  source_addresses = ["0.0.0.0/0", "::/0"]
 	}
 }
 `},
		GoodExample: []string{`
 resource "digitalocean_firewall" "good_example" {
 	name = "only-22-80-and-443"
   
 	droplet_ids = [digitalocean_droplet.web.id]
   
 	inbound_rule {
 	  protocol         = "tcp"
 	  port_range       = "22"
 	  source_addresses = ["192.168.1.0/24", "2002:1:2::/48"]
 	}
 }
 `},
		Links: []string{
			"https://registry.terraform.io/providers/digitalocean/digitalocean/latest/docs/resources/firewall",
		},
		RequiredTypes:  []string{"resource"},
		RequiredLabels: []string{"digitalocean_firewall"},
		Base:           compute.CheckNoPublicIngress,
	})
}
