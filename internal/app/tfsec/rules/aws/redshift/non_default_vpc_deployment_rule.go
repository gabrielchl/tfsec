package redshift

import (
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/rules/aws/redshift"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		LegacyID: "AWS087",
		BadExample: []string{`
 resource "aws_redshift_cluster" "bad_example" {
 	cluster_identifier = "tf-redshift-cluster"
 	database_name      = "mydb"
 	master_username    = "foo"
 	master_password    = "Mustbe8characters"
 	node_type          = "dc1.large"
 	cluster_type       = "single-node"
 }
 `},
		GoodExample: []string{`
 resource "aws_redshift_cluster" "good_example" {
 	cluster_identifier = "tf-redshift-cluster"
 	database_name      = "mydb"
 	master_username    = "foo"
 	master_password    = "Mustbe8characters"
 	node_type          = "dc1.large"
 	cluster_type       = "single-node"
 
 	cluster_subnet_group_name = "redshift_subnet"
 }
 `},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/redshift_cluster#cluster_subnet_group_name",
		},
		RequiredTypes:  []string{"resource"},
		RequiredLabels: []string{"aws_redshift_cluster"},
		Base:           redshift.CheckUsesVPC,
		CheckTerraform: func(resourceBlock block.Block, _ block.Module) (results rules.Results) {
			if resourceBlock.MissingChild("cluster_subnet_group_name") {
				results.Add("Resource is being deployed outside of a VPC", resourceBlock)
			}
			return results
		},
	})
}
