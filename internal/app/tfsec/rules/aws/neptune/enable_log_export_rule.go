package neptune

import (
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/rules/aws/neptune"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		BadExample: []string{`
 resource "aws_neptune_cluster" "bad_example" {
   cluster_identifier                  = "neptune-cluster-demo"
   engine                              = "neptune"
   backup_retention_period             = 5
   preferred_backup_window             = "07:00-09:00"
   skip_final_snapshot                 = true
   iam_database_authentication_enabled = true
   apply_immediately                   = true
   enable_cloudwatch_logs_exports      = []
 }
 `},
		GoodExample: []string{`
 resource "aws_neptune_cluster" "good_example" {
   cluster_identifier                  = "neptune-cluster-demo"
   engine                              = "neptune"
   backup_retention_period             = 5
   preferred_backup_window             = "07:00-09:00"
   skip_final_snapshot                 = true
   iam_database_authentication_enabled = true
   apply_immediately                   = true
   enable_cloudwatch_logs_exports      = ["audit"]
 }
 `},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/neptune_cluster#enable_cloudwatch_logs_exports",
		},
		RequiredTypes: []string{
			"resource",
		},
		RequiredLabels: []string{
			"aws_neptune_cluster",
		},
		Base: neptune.CheckEnableLogExport,
		CheckTerraform: func(resourceBlock block.Block, _ block.Module) (results rules.Results) {
			if enableCloudwatchLogsExportsAttr := resourceBlock.GetAttribute("enable_cloudwatch_logs_exports"); enableCloudwatchLogsExportsAttr.IsNil() {
				results.Add("Resource uses default value for enable_cloudwatch_logs_exports", resourceBlock)
			} else if enableCloudwatchLogsExportsAttr.NotContains("audit") {
				results.Add("Resource should have audit in enable_cloudwatch_logs_exports", enableCloudwatchLogsExportsAttr)
			}
			return results
		},
	})
}
