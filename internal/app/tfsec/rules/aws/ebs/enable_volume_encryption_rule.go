package ebs

import (
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/rules/aws/ebs"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		BadExample: []string{`
 resource "aws_ebs_volume" "bad_example" {
   availability_zone = "us-west-2a"
   size              = 40
 
   tags = {
     Name = "HelloWorld"
   }
   encrypted = false
 }
 `},
		GoodExample: []string{`
 resource "aws_ebs_volume" "good_example" {
   availability_zone = "us-west-2a"
   size              = 40
 
   tags = {
     Name = "HelloWorld"
   }
   encrypted = true
 }
 `},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ebs_volume#encrypted",
		},
		RequiredTypes: []string{
			"resource",
		},
		RequiredLabels: []string{
			"aws_ebs_volume",
		},
		Base: ebs.CheckEnableVolumeEncryption,
		CheckTerraform: func(resourceBlock block.Block, _ block.Module) (results rules.Results) {
			if encryptedAttr := resourceBlock.GetAttribute("encrypted"); encryptedAttr.IsNil() { // alert on use of default value
				results.Add("Resource uses default value for encrypted", resourceBlock)
			} else if encryptedAttr.IsFalse() {
				results.Add("Resource does not have encrypted set to true", encryptedAttr)
			}
			return results
		},
	})
}
