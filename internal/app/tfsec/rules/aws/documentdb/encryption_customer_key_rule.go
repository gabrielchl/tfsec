package documentdb

import (
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/rules/aws/documentdb"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		BadExample: []string{`
 resource "aws_docdb_cluster" "docdb" {
   cluster_identifier      = "my-docdb-cluster"
   engine                  = "docdb"
   master_username         = "foo"
   master_password         = "mustbeeightchars"
   backup_retention_period = 5
   preferred_backup_window = "07:00-09:00"
   skip_final_snapshot     = true
 }
 `},
		GoodExample: []string{`
 resource "aws_kms_key" "docdb_encryption" {
 	enable_key_rotation = true
 }
 			
 resource "aws_docdb_cluster" "docdb" {
   cluster_identifier      = "my-docdb-cluster"
   engine                  = "docdb"
   master_username         = "foo"
   master_password         = "mustbeeightchars"
   backup_retention_period = 5
   preferred_backup_window = "07:00-09:00"
   skip_final_snapshot     = true
   kms_key_id 			  = aws_kms_key.docdb_encryption.arn
 }
 `},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/docdb_cluster#kms_key_id",
		},
		RequiredTypes: []string{
			"resource",
		},
		RequiredLabels: []string{
			"aws_docdb_cluster",
			"aws_docdb_cluster_instance",
		},
		Base: documentdb.CheckEncryptionCustomerKey,
		CheckTerraform: func(resourceBlock block.Block, module block.Module) (results rules.Results) {

			if resourceBlock.MissingChild("kms_key_id") {
				results.Add("Resource does not use CMK", resourceBlock)
				return
			}

			kmsKeyAttr := resourceBlock.GetAttribute("kms_key_id")
			if kmsKeyAttr.IsDataBlockReference() {
				kmsData, err := module.GetReferencedBlock(kmsKeyAttr, resourceBlock)
				if err != nil {
					return
				}
				keyIdAttr := kmsData.GetAttribute("key_id")
				if keyIdAttr.IsNotNil() && keyIdAttr.StartsWith("alias/aws/") {
					results.Add("Resource explicitly uses the default CMK", keyIdAttr)
				}
			}

			return results
		},
	})
}
