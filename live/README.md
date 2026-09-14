# live: the opt-in suite against real AWS

This suite is never run in CI. It is run by hand, with James's approval each time, against a real
AWS account. It creates and destroys real resources through Cloud Control and IAM, and every value
configured is checked against what AWS reports back.

## What it needs

A dedicated AWS account, and a profile in `~/.aws/config` / `~/.aws/credentials` for an IAM user or
role in that account. **Never root keys** — `guard` in `live_test.go` refuses a caller identity ARN
ending `:root` before making any Cloud Control call.

The identity needs a policy allowing:

- `cloudcontrol:*`
- `sts:GetCallerIdentity`
- the handler permissions for the four types this suite exercises:
  `ec2:CreateVpc`, `ec2:DeleteVpc`, `ec2:ModifyVpcAttribute`, `ec2:CreateSubnet`, `ec2:DeleteSubnet`,
  `ec2:CreateSecurityGroup`, `ec2:DeleteSecurityGroup`, `ec2:AuthorizeSecurityGroupIngress`,
  `ec2:RevokeSecurityGroupIngress`, `ec2:AuthorizeSecurityGroupEgress`, `ec2:RevokeSecurityGroupEgress`,
  `ec2:Describe*`, `ec2:CreateTags`, `ec2:DeleteTags`, `iam:CreateRole`, `iam:DeleteRole`,
  `iam:GetRole`, `iam:UpdateRole`, `iam:ListRoles`, `iam:TagRole`, `iam:UntagRole`,
  `iam:ListRolePolicies`, `iam:ListAttachedRolePolicies`

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

## Running it

```bash
INFRENA_AWS_LIVE_PROFILE=infrena-live INFRENA_AWS_LIVE_ACCOUNT=111111111111 \
  go test -tags live -count=1 -v -timeout 30m ./live/
```

## Cleaning up after a crashed run

If a run is interrupted before its cleanup runs, `TestSweepLeftovers` finds and deletes anything
this suite tagged more than an hour ago:

```bash
INFRENA_AWS_LIVE_PROFILE=infrena-live INFRENA_AWS_LIVE_ACCOUNT=111111111111 \
  go test -tags live -count=1 -v -run TestSweepLeftovers ./live/
```

This suite never runs in CI: `.github/workflows/ci.yml` builds with `go vet ./...` and
`go test -count=1 ./...`, neither of which includes the `live` build tag.
