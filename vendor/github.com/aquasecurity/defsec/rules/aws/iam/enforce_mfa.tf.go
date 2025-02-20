package iam

var terraformEnforceMfaGoodExamples = []string{
        `
data aws_caller_identity current {}

resource aws_iam_group support {
  name =  "support"
}

module enforce_mfa {
  source  = "terraform-module/enforce-mfa/aws"
  version = "0.12.0"

  policy_name                     = "managed-mfa-enforce"
  account_id                      = data.aws_caller_identity.current.id
  groups                          = [aws_iam_group.support.name]
  manage_own_signing_certificates  = true
  manage_own_ssh_public_keys      = true
  manage_own_git_credentials      = true
}
`,
}

var terraformEnforceMfaBadExamples = []string{
        `
data aws_caller_identity current {}

resource aws_iam_group support {
  name =  "support"
}

resource aws_iam_group developers {
  name =  "developers"
}
`,`
data aws_caller_identity current {}

resource aws_iam_group support {
  name =  "support"
}

resource aws_iam_group developers {
  name =  "developers"
}

module enforce_mfa {
  source  = "terraform-module/enforce-mfa/aws"
  version = "0.12.0"

  policy_name                     = "managed-mfa-enforce"
  account_id                      = data.aws_caller_identity.current.id
  groups                          = [aws_iam_group.support.name]
  manage_own_signing_certificates  = true
  manage_own_ssh_public_keys      = true
  manage_own_git_credentials      = true
}
`,
}

var terraformEnforceMfaLinks = []string{
        `https://registry.terraform.io/modules/terraform-module/enforce-mfa/aws/latest`,`https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_passwords_account-policy.html#password-policy-details`,
}

var terraformEnforceMfaRemediationMarkdown = ``
