# live: the opt-in suite against real AWS

This suite is never run in CI. It is run by hand, with James's approval each time, against a real
AWS account. It creates and destroys real resources through Cloud Control and IAM, and every value
configured is checked against what AWS reports back.

## What it needs

A dedicated AWS account, and a profile in `~/.aws/config` / `~/.aws/credentials` for an IAM user or
role in that account. **Never root keys** — `guard` in `live_test.go` refuses a caller identity ARN
ending `:root` before making any Cloud Control call.

The identity needs a policy allowing:

- the Cloud Control actions, which IAM names under `cloudformation:` (not `cloudcontrol:`):
  `cloudformation:CreateResource`, `cloudformation:GetResource`, `cloudformation:UpdateResource`,
  `cloudformation:DeleteResource`, `cloudformation:ListResources`,
  `cloudformation:GetResourceRequestStatus`, `cloudformation:ListResourceRequests`,
  `cloudformation:CancelResourceRequest`
- `sts:GetCallerIdentity`
- the two EC2 reads discovery uses to say which resources AWS itself owns: `ec2:DescribeVpcs` and
  `ec2:DescribeSubnets`. **Both are already covered by the `ec2:Describe*` below**, so a policy built from this
  list needs no change; they are named because a narrower policy that spells out each action would need them,
  and because without them discovery still works — it marks nothing and says so on stderr.
- the handler permissions for the four types `TestTheLifecycleAgainstRealAWS` exercises:
  `ec2:CreateVpc`, `ec2:DeleteVpc`, `ec2:ModifyVpcAttribute`, `ec2:CreateSubnet`, `ec2:DeleteSubnet`,
  `ec2:CreateSecurityGroup`, `ec2:DeleteSecurityGroup`, `ec2:AuthorizeSecurityGroupIngress`,
  `ec2:RevokeSecurityGroupIngress`, `ec2:AuthorizeSecurityGroupEgress`, `ec2:RevokeSecurityGroupEgress`,
  `ec2:Describe*`, `ec2:CreateTags`, `ec2:DeleteTags`, `iam:CreateRole`, `iam:DeleteRole`,
  `iam:GetRole`, `iam:UpdateRole`, `iam:ListRoles`, `iam:TagRole`, `iam:UntagRole`,
  `iam:ListRolePolicies`, `iam:ListAttachedRolePolicies`
- the handler permissions for the three more types `TestStorageAndContainersAgainstRealAWS` exercises. These are
  taken as-is from each type's `handlers.*.permissions` in `schemas/CloudformationSchema.zip`
  (`aws-s3-bucket.json`, `aws-ecr-repository.json`, `aws-ecs-cluster.json`), covering every property the handler
  can touch, not only the narrow set this test configures — pare this down to what a real run actually needs once
  it has been exercised once, since AWS names the missing action in `AccessDenied`:
  - S3 bucket (`aws-s3-bucket.json`, create/read/update/delete/list handlers, deduplicated):
    `s3:CreateBucket`, `s3:PutBucketTagging`, `s3:TagResource`, `s3:PutBucketAbac`,
    `s3:PutAnalyticsConfiguration`, `s3:PutEncryptionConfiguration`, `s3:PutBucketCORS`,
    `s3:PutInventoryConfiguration`, `s3:PutLifecycleConfiguration`, `s3:PutMetricsConfiguration`,
    `s3:PutBucketNotification`, `s3:PutBucketReplication`, `s3:PutBucketWebsite`,
    `s3:PutAccelerateConfiguration`, `s3:PutBucketPublicAccessBlock`, `s3:PutReplicationConfiguration`,
    `s3:PutObjectAcl`, `s3:PutBucketObjectLockConfiguration`, `s3:GetBucketAcl`, `s3:ListBucket`,
    `iam:PassRole`, `s3:DeleteObject`, `s3:PutBucketLogging`, `s3:PutBucketVersioning`,
    `s3:PutObjectLockConfiguration`, `s3:PutBucketOwnershipControls`, `s3:PutIntelligentTieringConfiguration`,
    `s3:GetBucketMetadataTableConfiguration`, `s3:CreateBucketMetadataTableConfiguration`,
    `s3tables:CreateNamespace`, `s3tables:CreateTable`, `s3tables:CreateTableBucket`, `s3tables:GetTable`,
    `s3tables:PutTableBucketPolicy`, `s3tables:PutTableEncryption`, `s3tables:PutTablePolicy`,
    `s3tables:GetTableMetadataLocation`, `s3tables:UpdateTableMetadataLocation`,
    `s3:GetAccelerateConfiguration`, `s3:GetLifecycleConfiguration`, `s3:GetBucketPublicAccessBlock`,
    `s3:GetAnalyticsConfiguration`, `s3:GetBucketCORS`, `s3:GetEncryptionConfiguration`,
    `s3:GetInventoryConfiguration`, `s3:GetBucketLogging`, `s3:GetMetricsConfiguration`,
    `s3:GetBucketNotification`, `s3:GetBucketVersioning`, `s3:GetReplicationConfiguration`,
    `s3:GetBucketWebsite`, `s3:GetBucketObjectLockConfiguration`, `s3:GetBucketTagging`,
    `s3:ListTagsForResource`, `s3:GetBucketAbac`, `s3:GetBucketOwnershipControls`,
    `s3:GetIntelligentTieringConfiguration`, `s3:PutBucketAcl`, `s3:UntagResource`,
    `s3:DeleteBucketMetadataTableConfiguration`, `s3:UpdateBucketMetadataJournalTableConfiguration`,
    `s3:UpdateBucketMetadataInventoryTableConfiguration`, `s3:UpdateBucketMetadataAnnotationTableConfiguration`,
    `s3:DeleteBucketWebsite`, `s3:DeleteBucketAnalyticsConfiguration`, `s3:DeleteBucketCors`,
    `s3:DeleteBucketMetricsConfiguration`, `s3:DeleteBucketEncryption`, `s3:DeleteBucketLifecycle`,
    `s3:DeleteBucketReplication`, `s3:DeleteBucket`, `s3:ListAllMyBuckets`
  - ECR repository (`aws-ecr-repository.json`): `ecr:CreateRepository`, `ecr:PutLifecyclePolicy`,
    `ecr:SetRepositoryPolicy`, `ecr:TagResource`, `kms:DescribeKey`, `kms:CreateGrant`, `kms:RetireGrant`,
    `ecr:DescribeRepositories`, `ecr:GetLifecyclePolicy`, `ecr:GetRepositoryPolicy`, `ecr:ListTagsForResource`,
    `ecr:UntagResource`, `ecr:DeleteLifecyclePolicy`, `ecr:DeleteRepositoryPolicy`,
    `ecr:PutImageScanningConfiguration`, `ecr:PutImageTagMutability`, `ecr:DeleteRepository`
  - ECS cluster (`aws-ecs-cluster.json`): `ecs:CreateCluster`, `ecs:DescribeClusters`,
    `iam:CreateServiceLinkedRole`, `ecs:TagResource`, `kms:DescribeKey`, `ecs:PutAccountSettingDefault`,
    `ecs:UntagResource`, `ecs:PutAccountSetting`, `ecs:ListTagsForResource`, `ecs:UpdateCluster`,
    `ecs:UpdateClusterSettings`, `ecs:PutClusterCapacityProviders`, `ecs:DeleteCluster`, `ecs:ListClusters`

A missing permission shows up as `AccessDenied` naming the action; add that action and try again.

## Variables

| Variable | Required | Meaning |
| --- | --- | --- |
| `INFRENA_AWS_LIVE_PROFILE` | yes | the AWS CLI profile to use |
| `INFRENA_AWS_LIVE_ACCOUNT` | yes | the account ID the profile must resolve to; the suite refuses to run against any other account |
| `INFRENA_AWS_LIVE_REGION` | no (default `us-east-1`) | the region to create resources in |

Do not put a real account ID in this file or in any committed file. Set it in your shell.

## What it creates

`TestTheLifecycleAgainstRealAWS` creates a VPC (`10.99.0.0/16`), a subnet (`10.99.1.0/24`), a
security group, and an IAM role named `infrena-live-<unix time>`, all tagged
`infrena-live-run: <unix time>`. It updates each, discovers and imports the VPC, then deletes
everything it created, in dependency order. It takes a few minutes to run.

It refuses to run against the wrong account (checked with `sts:GetCallerIdentity` before any
Cloud Control call) or with root credentials.

`TestStorageAndContainersAgainstRealAWS` creates an S3 bucket, an ECR repository and an ECS cluster, all named
`infrena-live-<unix time>` and tagged `infrena-live-run: <unix time>`, to cover nested-value reconciliation the
first four types don't exercise:

- the bucket with `VersioningConfiguration` (`status: Enabled`) and a `LifecycleConfiguration` with one rule
  (`status: Enabled`, `expiration_in_days: 1`), both written with snake_case keys; updated by changing the
  expiration to 3 days
- the repository with `ImageTagMutability: MUTABLE`, `ImageScanningConfiguration` (`scan_on_push: true`), and a
  `LifecyclePolicy` whose `lifecycle_policy_text` is JSON text spaced unlike a minified document; updated by
  flipping mutability to `IMMUTABLE`
- the cluster with one `ClusterSettings` entry (`containerInsights` left `disabled` throughout, since enabling
  Container Insights is not free); updated by adding a tag

Nothing is ever put in the bucket or the repository — it stays free-tier, and an empty bucket/repository is
required to delete it. It deletes everything it created, in dependency order, and is a separate test function
from `TestTheLifecycleAgainstRealAWS` so a run can target either alone with `-run`.

## Running it

```bash
INFRENA_AWS_LIVE_PROFILE=infrena-live INFRENA_AWS_LIVE_ACCOUNT=111111111111 \
  go test -tags live -count=1 -v -timeout 30m ./live/

# or just one of the two suites:
INFRENA_AWS_LIVE_PROFILE=infrena-live INFRENA_AWS_LIVE_ACCOUNT=111111111111 \
  go test -tags live -count=1 -v -timeout 30m -run TestStorageAndContainersAgainstRealAWS ./live/
```

## Cleaning up after a crashed run

If a run is interrupted before its cleanup runs, `TestSweepLeftovers` finds and deletes anything
this suite tagged more than an hour ago, across all seven types (buckets and repositories included — safe
without checking for emptiness, since this suite never puts objects or images in them):

```bash
INFRENA_AWS_LIVE_PROFILE=infrena-live INFRENA_AWS_LIVE_ACCOUNT=111111111111 \
  go test -tags live -count=1 -v -run TestSweepLeftovers ./live/
```

This suite never runs in CI: `.github/workflows/ci.yml` builds with `go vet ./...` and
`go test -count=1 ./...`, neither of which includes the `live` build tag.
