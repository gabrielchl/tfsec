package elasticache

import (
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/rules/aws/elasticache"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		BadExample: []string{`
resource "aws_security_group" "bar" {
	name = "security-group"
}

resource "aws_elasticache_security_group" "bad_example" {
	name = "elasticache-security-group"
	security_group_names = [aws_security_group.bar.name]
	description = ""
}
		`},
		GoodExample: []string{`
resource "aws_security_group" "bar" {
	name = "security-group"
}

resource "aws_elasticache_security_group" "good_example" {
	name = "elasticache-security-group"
	security_group_names = [aws_security_group.bar.name]
	description = "something"
}
	`},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elasticache_security_group#description",
		},
		RequiredTypes: []string{
			"resource",
		},
		RequiredLabels: []string{
			"aws_elasticache_security_group",
		},
		Base: elasticache.CheckAddDescriptionForSecurityGroup,
		CheckTerraform: func(resourceBlock block.Block, _ block.Module) (results rules.Results) {
			if descriptionAttr := resourceBlock.GetAttribute("description"); descriptionAttr.IsNil() { // alert on use of default value
				results.Add("Resource uses default value for description", resourceBlock)
			} else if descriptionAttr.IsEmpty() {
				results.Add("Resource has description set to ", descriptionAttr)
			}
			return results
		},
	})
}
