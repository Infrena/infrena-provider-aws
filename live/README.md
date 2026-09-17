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
- the handler permissions for the three RDS types `TestDatabasesAgainstRealAWS` exercises, taken the same way
  from each type's `handlers.*.permissions` in `schemas/CloudformationSchema.zip` (`aws-rds-dbinstance.json`,
  `aws-rds-dbsubnetgroup.json`, `aws-rds-dbparametergroup.json`):
  - DB instance (`aws-rds-dbinstance.json`, create/read/update/delete/list handlers, deduplicated):
    `ec2:DescribeAccountAttributes`, `ec2:DescribeAvailabilityZones`, `ec2:DescribeInternetGateways`,
    `ec2:DescribeSecurityGroups`, `ec2:DescribeSubnets`, `ec2:DescribeVpcAttribute`, `ec2:DescribeVpcs`,
    `iam:CreateServiceLinkedRole`, `iam:GetRole`, `iam:ListRoles`, `iam:PassRole`, `kms:CreateGrant`,
    `kms:DescribeKey`, `rds:AddRoleToDBInstance`, `rds:AddTagsToResource`, `rds:CreateDBInstance`,
    `rds:CreateDBInstanceReadReplica`, `rds:CreateDBSnapshot`, `rds:DeleteDBInstance`,
    `rds:DescribeDBClusterSnapshots`, `rds:DescribeDBClusters`, `rds:DescribeDBEngineVersions`,
    `rds:DescribeDBInstanceAutomatedBackups`, `rds:DescribeDBInstances`, `rds:DescribeDBParameterGroups`,
    `rds:DescribeDBSnapshots`, `rds:DescribeEvents`, `rds:ModifyDBInstance`, `rds:PromoteReadReplica`,
    `rds:RebootDBInstance`, `rds:RemoveRoleFromDBInstance`, `rds:RemoveTagsFromResource`,
    `rds:RestoreDBInstanceFromDBSnapshot`, `rds:RestoreDBInstanceToPointInTime`,
    `rds:StartDBInstanceAutomatedBackupsReplication`, `rds:StopDBInstanceAutomatedBackupsReplication`,
    `secretsmanager:CreateSecret`, `secretsmanager:TagResource`
  - DB subnet group (`aws-rds-dbsubnetgroup.json`): `iam:CreateServiceLinkedRole`, `rds:AddTagsToResource`,
    `rds:CreateDBSubnetGroup`, `rds:DeleteDBSubnetGroup`, `rds:DescribeDBSubnetGroups`, `rds:ListTagsForResource`,
    `rds:ModifyDBSubnetGroup`, `rds:RemoveTagsFromResource`
  - DB parameter group (`aws-rds-dbparametergroup.json`): `iam:CreateServiceLinkedRole`, `rds:AddTagsToResource`,
    `rds:CreateDBParameterGroup`, `rds:DeleteDBParameterGroup`, `rds:DescribeDBParameterGroups`,
    `rds:DescribeDBParameters`, `rds:DescribeEngineDefaultParameters`, `rds:ListTagsForResource`,
    `rds:ModifyDBParameterGroup`, `rds:RemoveTagsFromResource`, `rds:ResetDBParameterGroup`
- the handler permissions for the three ELBv2 types `TestLoadBalancersAgainstRealAWS` exercises, taken the same way
  from each type's `handlers.*.permissions` in `schemas/CloudformationSchema.zip`
  (`aws-elasticloadbalancingv2-loadbalancer.json`, `aws-elasticloadbalancingv2-targetgroup.json`,
  `aws-elasticloadbalancingv2-listener.json`). The VPC, subnets and security group it also creates need nothing
  beyond the `ec2:` actions already listed above:
  - load balancer: `elasticloadbalancing:AddTags`, `elasticloadbalancing:CreateLoadBalancer`,
    `elasticloadbalancing:DeleteLoadBalancer`, `elasticloadbalancing:DescribeCapacityReservation`,
    `elasticloadbalancing:DescribeLoadBalancerAttributes`, `elasticloadbalancing:DescribeLoadBalancers`,
    `elasticloadbalancing:DescribeTags`, `elasticloadbalancing:ModifyCapacityReservation`,
    `elasticloadbalancing:ModifyIpPools`, `elasticloadbalancing:ModifyLoadBalancerAttributes`,
    `elasticloadbalancing:RemoveTags`, `elasticloadbalancing:SetIpAddressType`,
    `elasticloadbalancing:SetSecurityGroups`, `elasticloadbalancing:SetSubnets`, plus the one `ec2:` action its
    handlers list, `ec2:DescribeIpamPools` (**already covered by the `ec2:Describe*` above**)
  - target group: `elasticloadbalancing:AddTags`, `elasticloadbalancing:CreateTargetGroup`,
    `elasticloadbalancing:DeleteTargetGroup`, `elasticloadbalancing:DeregisterTargets`,
    `elasticloadbalancing:DescribeTags`, `elasticloadbalancing:DescribeTargetGroupAttributes`,
    `elasticloadbalancing:DescribeTargetGroups`, `elasticloadbalancing:DescribeTargetHealth`,
    `elasticloadbalancing:ModifyTargetGroup`, `elasticloadbalancing:ModifyTargetGroupAttributes`,
    `elasticloadbalancing:RegisterTargets`, `elasticloadbalancing:RemoveTags`, plus the one `ec2:` action its
    handlers list, `ec2:DescribeVpcs` (**already covered by the `ec2:Describe*` above**)
  - listener: `elasticloadbalancing:AddTags`, `elasticloadbalancing:CreateListener`,
    `elasticloadbalancing:DeleteListener`, `elasticloadbalancing:DescribeListenerAttributes`,
    `elasticloadbalancing:DescribeListeners`, `elasticloadbalancing:DescribeTags`,
    `elasticloadbalancing:ModifyListener`, `elasticloadbalancing:ModifyListenerAttributes`,
    `elasticloadbalancing:RemoveTags`. Its handlers list no `ec2:` action; they do list
    `cognito-idp:DescribeUserPoolClient`, which only an `authenticate-cognito` default action needs and this test
    never configures — leave it out unless a run says otherwise.

- the handler permissions for the two Route 53 types `TestDNSAgainstRealAWS` exercises, taken the same way from
  each type's `handlers.*.permissions` in `schemas/CloudformationSchema.zip` (`aws-route53-hostedzone.json`,
  `aws-route53-recordset.json`). The VPC it also creates needs nothing beyond the `ec2:` actions already listed
  above:
  - hosted zone (create/read/update/delete/list handlers, deduplicated): `route53:CreateHostedZone`,
    `route53:CreateQueryLoggingConfig`, `route53:ChangeTagsForResource`, `route53:GetChange`,
    `route53:GetHostedZone`, `route53:UpdateHostedZoneFeatures`, `route53:AssociateVPCWithHostedZone`,
    `route53:ListTagsForResource`, `route53:ListQueryLoggingConfigs`, `route53:UpdateHostedZoneComment`,
    `route53:DisassociateVPCFromHostedZone`, `route53:DeleteQueryLoggingConfig`, `route53:DeleteHostedZone`,
    `route53:ListHostedZones`, plus the one `ec2:` action its handlers list, `ec2:DescribeVpcs` (**already
    covered by the `ec2:Describe*` above**)
  - record set: `route53:ListHostedZones`, `route53:GetChange`, `route53:ChangeResourceRecordSets`,
    `route53:ListResourceRecordSets`, `route53:GetHostedZone`

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

`TestDatabasesAgainstRealAWS` creates a VPC (`10.98.0.0/16`) with two subnets in different Availability Zones
(`10.98.1.0/24`, `10.98.2.0/24` — a DB subnet group must span at least two AZs), a DB subnet group, a DB
parameter group, and a DB instance, all named `infrena-live-<unix time>` and tagged
`infrena-live-run: <unix time>`:

- the DB instance is the smallest free-tier-eligible shape: engine `postgres`, `db.t3.micro`, 20 GiB `gp3`
  storage, single-AZ, no public access, `BackupRetentionPeriod: 0`, and a `MasterUserPassword` generated fresh
  for the run. **This is free-tier eligible but NOT free** if the account's AWS Free Tier allowance is already
  used up elsewhere (by another instance, another account under the same organization's Free Tier, or because
  the account is past its 12-month Free Tier window) — RDS bills for `db.t3.micro` and `gp3` storage like any
  other instance once free tier no longer applies.
- it updates the DB instance's tags — cheap, immediate (RDS's `AddTagsToResource`, not `ModifyDBInstance`), and
  deliberately not the instance class or storage, which would take AWS a long time or replace the instance
- `AWS::RDS::DBInstance` has no `SkipFinalSnapshot`-like property in its schema, so there is nothing to set to
  avoid a final snapshot on delete; Cloud Control's `DeleteResource` for this type has no way to be told to take
  one either

It deletes everything it created, in dependency order (instance, then subnet group, then parameter group, then
the subnets, then the VPC), and is a separate test function so a run can target it alone with `-run`.

`TestLoadBalancersAgainstRealAWS` creates a VPC (`10.97.0.0/16`) with two subnets in different Availability Zones
(`10.97.1.0/24`, `10.97.2.0/24` — an Application Load Balancer must span at least two), a security group, an
Application Load Balancer, a target group, and a listener forwarding to it, all named `infrena-live-<unix time>`
(the target group `infrena-live-tg-<unix time>`) and tagged `infrena-live-run: <unix time>`:

- **This is NOT free tier.** An ALB costs roughly $0.0225 an hour in `us-east-1` plus LCU-hours, billed per hour
  started, so one run costs well under a cent. Nothing else it creates costs anything, and no EC2 instance is
  registered as a target — an empty target group is legal and needs no instance to pay for or wait on.
- the load balancer with `Name`, `Type: application`, `Scheme: internal`, `IpAddressType: ipv4`, both subnets, and
  the security group; updated by adding a tag
- the target group with `Protocol: HTTP`, `Port: 80`, the same VPC, `TargetType: instance`, and every health check
  value spelled out (`HealthCheckPath: /`, `HealthCheckPort: traffic-port`, interval 30, timeout 5, thresholds 2
  and 3, `Matcher` `http_code: 200`) rather than left to AWS's defaults, which it reports on every read; updated by
  changing the health check path to `/healthz` and the interval to 10 seconds (the timeout must stay below it)
- the listener with `Protocol: HTTP`, `Port: 80` and one `forward` default action naming the target group. AWS
  answers a read of that action with more than was sent (a `ForwardConfig` and an `Order`); reconciliation drops
  keys AWS added that the configuration does not have, so it reads back as what was configured.
- `Scheme` is `internal`, not `internet-facing`, because AWS requires an internet gateway attached to the VPC
  before it will create an internet-facing load balancer. That would mean an `AWS::EC2::InternetGateway` and an
  `AWS::EC2::VPCGatewayAttachment` here, and the attachment would have to be deleted before the VPC while the
  ALB's network interfaces are still lingering — another resource in the middle of the one part of teardown most
  likely to need a retry. An internal ALB exercises the same three ELBv2 types, with nothing reachable from the
  internet.

It deletes everything it created in dependency order (listener, load balancer, target group, security group, the
subnets, then the VPC), waiting out a lingering dependency at each step, and is a separate test function so a run
can target it alone with `-run`.

`TestDNSAgainstRealAWS` creates a VPC (`10.96.0.0/16`), a **private** Route 53 hosted zone
(`infrena-live-<unix time>.internal.`, in the reserved `.internal` TLD, so it cannot collide with anything real)
associated with that VPC, and two record sets in it (an A record and a TXT record, both named
`<www|txt>.infrena-live-<unix time>.internal`), tagged `infrena-live-run: <unix time>` (record sets take no
tags at all — see below):

- **Cost:** a hosted zone is $0.50/month, prorated, and AWS does not charge for a zone deleted within 12 hours
  of creation — this test's zone lives for at most a few minutes. Record set queries are negligible (this test
  makes none; nothing here is reachable, being a private zone with no EC2 instance behind it).
- Route 53 is a **global** type in this plugin (`gen/overlay.yaml`'s `global:` list matches `AWS::Route53::*`):
  neither `aws.hostedzone` nor `aws.recordset` takes a region attribute, and both provider IDs start `global/`.
- both `Name` properties are configured already in the form Route 53's own `propertyTransform` normalises them
  to (a hosted zone's gets a trailing dot appended and is lowercased; a record set's has a trailing dot
  stripped and is lowercased) — see the comment in `live_test.go` for the exact transform text, taken from
  `schemas/CloudformationSchema.zip`'s `aws-route53-hostedzone.json` and `aws-route53-recordset.json`.
- it updates the A record's `TTL` (300 to 600) and the hosted zone's `HostedZoneConfig` comment — both cheap,
  in-place changes.
- `AWS::Route53::RecordSet`'s list handler needs a parent hosted zone, which is not expressed as a top-level
  `required` in its schema (the same shape of problem as the ELB listener noted above), so today this type is
  discovered as though it were plainly listable and 2026-09-16's live run found it cannot in fact be listed
  without one. A concurrent generator change fixes that; until it lands, this test only logs what discovery
  reports for its record sets rather than asserting on it, with a comment explaining why.
- a hosted zone and its record sets are quick: creates, updates and deletes each typically take AWS a few
  seconds. `-timeout 10m` is generous headroom alongside the VPC this test also creates and destroys.

It deletes everything it created in dependency order (both record sets, then the hosted zone, then the VPC —
`DeleteHostedZone` refuses a zone that still holds anything but its default NS/SOA records), and is a separate
test function so a run can target it alone with `-run`.

### The load balancer's longer timeouts

Creating an ALB typically takes AWS 2 to 4 minutes and deleting one 1 to 3, and the elastic network interfaces it
leaves in the subnets can refuse the security group and the subnets for a few minutes after Cloud Control reports
the load balancer gone (the test waits that out, for up to six minutes per resource). A run of
`TestLoadBalancersAgainstRealAWS` therefore needs `go test`'s own `-timeout` at `30m`.

### The DB instance's longer timeouts

`AWS::RDS::DBInstance`'s create, update and delete handlers are each registered with a 2160-minute (36-hour)
`timeoutInMinutes` in the catalog (`internal/ccprov/catalog_test.go`, sourced from
`schemas/CloudformationSchema.zip`'s `aws-rds-dbinstance.json`) — that is Cloud Control's own outer ceiling on
one request, not a prediction of how long a real create takes. A real create typically takes AWS 5 to 10
minutes and a delete several minutes more. That is far more than the few minutes the other tests take, so a run
that includes `TestDatabasesAgainstRealAWS` needs `go test`'s own `-timeout` raised well past the `30m` used
below — `45m` or more is reasonable headroom; raise it further if a run is timing out mid-create.

## Running it

A run that includes `TestDatabasesAgainstRealAWS` or `TestLoadBalancersAgainstRealAWS` (whether by itself or as
part of the whole package) needs a longer `-timeout` than the other tests alone do — see the two timeout sections
above:

```bash
# TestTheLifecycleAgainstRealAWS and TestStorageAndContainersAgainstRealAWS only, 30m is plenty:
INFRENA_AWS_LIVE_PROFILE=infrena-live INFRENA_AWS_LIVE_ACCOUNT=111111111111 \
  go test -tags live -count=1 -v -timeout 30m -run 'TestTheLifecycleAgainstRealAWS|TestStorageAndContainersAgainstRealAWS' ./live/

# TestLoadBalancersAgainstRealAWS alone: the ALB's create, delete and network interface cleanup need 30m:
INFRENA_AWS_LIVE_PROFILE=infrena-live INFRENA_AWS_LIVE_ACCOUNT=111111111111 \
  go test -tags live -count=1 -v -timeout 30m -run TestLoadBalancersAgainstRealAWS ./live/

# the whole package now runs the DB instance and the ALB back to back: give it 60m:
INFRENA_AWS_LIVE_PROFILE=infrena-live INFRENA_AWS_LIVE_ACCOUNT=111111111111 \
  go test -tags live -count=1 -v -timeout 60m ./live/

INFRENA_AWS_LIVE_PROFILE=infrena-live INFRENA_AWS_LIVE_ACCOUNT=111111111111 \
  go test -tags live -count=1 -v -timeout 45m -run TestDatabasesAgainstRealAWS ./live/

# TestDNSAgainstRealAWS alone: a hosted zone and its record sets are quick, 10m is generous:
INFRENA_AWS_LIVE_PROFILE=infrena-live INFRENA_AWS_LIVE_ACCOUNT=111111111111 \
  go test -tags live -count=1 -v -timeout 10m -run TestDNSAgainstRealAWS ./live/
```

## Cleaning up after a crashed run

If a run is interrupted before its cleanup runs, `TestSweepLeftovers` finds and deletes anything
this suite tagged more than an hour ago, across all fourteen taggable types (buckets and repositories included —
safe without checking for emptiness, since this suite never puts objects or images in them), listeners first,
then load balancers, then target groups, then the hosted zone right before the VPC it may be associated with, so
nothing is refused for still being in use. `AWS::Route53::RecordSet` is not among them: its schema declares
`tagging: {taggable: false}`, so a record set cannot be tagged and this sweep has no way to find one by the run
tag the way it finds everything else. A crash between `TestDNSAgainstRealAWS` creating a record set and its own
teardown running leaves that record behind untagged, which then makes the hosted zone sweep below fail (as a
non-fatal `t.Error`, not `t.Fatal`) until the record is removed by hand — narrow and cheap, since a record set
create is a few seconds of API calls and nothing about an orphaned record itself is billed:

```bash
INFRENA_AWS_LIVE_PROFILE=infrena-live INFRENA_AWS_LIVE_ACCOUNT=111111111111 \
  go test -tags live -count=1 -v -run TestSweepLeftovers ./live/
```

This suite never runs in CI: `.github/workflows/ci.yml` builds with `go vet ./...` and
`go test -count=1 ./...`, neither of which includes the `live` build tag.
