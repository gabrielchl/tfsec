---
title: S3 Bucket does not have logging enabled.
---

### Default Severity: <span class="severity high">high</span>

### Explanation


Buckets should have logging enabled so that access can be audited. 


### Possible Impact
There is no way to determine the access to this bucket

### Suggested Resolution
Add a logging block to the resource to enable access logging


### Insecure Example

The following example will fail the aws-s3-no-public-access-with-acl check.
```terraform

resource "aws_s3_bucket" "bad_example" {
	acl = "public-read"
}

```



### Secure Example

The following example will pass the aws-s3-no-public-access-with-acl check.
```terraform

resource "aws_s3_bucket" "good_example" {
	acl = "private"
}

```



### Links


- [https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket){:target="_blank" rel="nofollow noreferrer noopener"}

- [https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerLogs.html](https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerLogs.html){:target="_blank" rel="nofollow noreferrer noopener"}



