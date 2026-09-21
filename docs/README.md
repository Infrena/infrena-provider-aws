# infrena AWS provider documentation

The AWS provider plugin lets infrena manage every resource type AWS Cloud Control API supports, 1,584 of them, generated from AWS's own schemas.

## Start here

- [README](../README.md): installing the plugin, credentials, regions, the attribute naming rules and a worked example.
- [Roadmap](ROADMAP.md): what's shipped and what's next.

## Service guides

Detailed guides with examples for the most used services. Every example is a real project under [`examples/`](../examples) that `scripts/check-examples` compiles with infrena.

| Area | Guide |
| --- | --- |
| Networking | [VPC, subnets, gateways, routes, security groups](services/vpc.md) |
| Compute | [EC2 instances, launch templates, key pairs, volumes](services/ec2.md) |
| Load balancing | [Application and network load balancers](services/elb.md) |
| Identity | [IAM roles, policies, instance profiles, users](services/iam.md) |
| Storage | [S3 buckets and bucket policies](services/s3.md) |
| Encryption | [KMS keys and aliases](services/kms.md) |
| Containers | [ECS clusters, task definitions, services](services/ecs.md) |
| Container images | [ECR repositories](services/ecr.md) |
| Functions | [Lambda functions, permissions, URLs](services/lambda.md) |
| Databases | [RDS instances, subnet and parameter groups](services/rds.md) |
| DNS | [Route 53 hosted zones and records](services/route53.md) |
| Logs | [CloudWatch Logs groups and metric filters](services/logs.md) |

## Reference for every type

[`reference/`](reference/README.md) has one page per resource type: every attribute, the spellings it accepts, whether AWS chooses a value when you leave it unset, what forces a replacement, which attributes refer to other resources, and the import ID. The pages are generated from the plugin's catalog with `go run ./cmd/gen-docs`, so they always match the released plugin.

## Modules

[`modules/`](../modules) holds example modules for common setups: a VPC with public and private subnets, an ECS Fargate service, a private S3 bucket and a Lambda function.

## Types this plugin cannot manage

Six CloudFormation types declare create, update and delete but **no read handler**. Without a
read there is no way to detect drift or confirm what was applied, so they are excluded from
the catalog rather than half-supported. If you go looking for one of these and cannot find
it, the provider is not broken.

Three have a direct equivalent that is supported, and you should use it instead:

| Not available | Use instead |
| --- | --- |
| `AWS::IAM::Policy` | `aws.iam.rolepolicy`, `aws.iam.userpolicy`, `aws.iam.grouppolicy`, or `aws.iam.managedpolicy` |
| `AWS::SNS::TopicPolicy` | `aws.sns.topicinlinepolicy` |
| `AWS::SQS::QueuePolicy` | `aws.sqs.queueinlinepolicy` |

Three have no equivalent, and there is nothing this plugin can do about it until AWS adds a
read handler:

- `AWS::EC2::NetworkAclEntry` — the network ACL itself is supported, its entries are not
- `AWS::EC2::VPNGatewayRoutePropagation`
- `AWS::AmazonMQ::ConfigurationAssociation` — brokers and configurations are supported, the
  association between them is not

## How the plugin decides things

- [Reference relationships](references-review.md): which attributes refer to other resources, and how uncertain matches were reviewed.
- [Design](specs/2026-09-14-generic-cloudcontrol-provider.md) and [implementation plan](plans/2026-09-14-generic-cloudcontrol-provider.md): the records of how the provider was built.
