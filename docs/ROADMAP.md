# Roadmap

The AWS provider for infrena. Docs live in this repository until the project goes live, then move to the wiki and a website.

## Shipped

| Version | infrena | What |
| --- | --- | --- |
| 0.1.0 | 0.4.0 | Every AWS type Cloud Control supports (1,584), generated from AWS's schemas. Friendly attribute names, nested values that converge, tags as a map, discovery and import. |
| 0.1.1 | 0.4.0 | IAM policy documents kept as written. |
| 0.1.2 | 0.4.0 | JSON documents compared by meaning at any depth. Live suite covers S3, ECR and ECS. |
| 0.1.3 | 0.5.0 | New variable syntax, `${var.x}`. |
| 0.2.0 | 0.6.2 | Whole-resource references (`vpc: ${vpc}`): 922 attributes declare what they refer to. |
| 0.3.0 | 0.7.0 | Plugin protocol 4. Discovered resources named from their `Name` tag; `import --generate` writes references. |
| 0.4.0 | 0.7.0 | System-owned resources flagged so `import` doesn't adopt them by accident: the default VPC, default subnets, the default security group, service-linked roles and CloudFormation-managed resources. The defaults need the EC2 API, because Cloud Control does not report them. |
| 0.5.0 | 0.7.1 | Update builds its patch from the state infrena now refreshes before planning, so the extra read before every update is gone: one fewer AWS call per changed resource. |
| 0.6.0 | 0.13.0 | **Fixes a plan that never converged.** A property whose every leaf is write-only — a Lambda function's `Code` among them — was generated as an ordinary attribute, and AWS returns it empty, so every plan wanted it back and every apply re-sent it. Fixed in the generator and again in the provider, which had been keeping a configured value only when AWS omitted a property rather than when it returned an empty one. 8 types gained an attribute. Also: 45 of the 46 pending reference edges settled (30 approved, 15 rejected), live coverage for ECS services, task definitions, Route 53 and Lambda, and a fix to discovery for types whose list handler needs a parent. |
| 0.6.1 | 0.13.0 | **Importing a resource no longer loses its tags.** The tag map was marked as a value the cloud chooses, which put it in the same bucket as a setting with a default — so `import --generate` wrote no tags at all for 905 of the 958 taggable types. Tags still reached state, so the plan straight after an import looked clean; the loss showed a step later, when adding one tag by hand diffed it against the three in state and proposed deleting the two you had never seen. Tags now carry no flag, which makes configuration authoritative for them. Note the consequence: a hand-written resource that omits `tags:` will now propose removing the tags AWS holds. |
| 0.7.0 | 0.13.0 | **`import` no longer adopts an instance a fleet launched.** An instance tagged `aws:autoscaling:groupName`, `aws:ec2:fleet-id` or `aws:ec2spot:fleet-request-id` belongs to a controller that replaces it on its own schedule, which on a spot fleet draining short jobs is minutes. Adopting one recorded an instance id that stopped existing almost at once, so every plan afterwards proposed creating it again. Reported as an engine bug and it was not one: the instances really were gone. `aws:ec2launchtemplate:id` is deliberately not evidence, because launching from a template is something people do by hand. Note the consequence: a discover or import that used to adopt an account's fleet instances now skips them unless you name one. |

Documentation ships with these: a generated reference page for every one of the 1,584 types, longer guides with examples for the common services, runnable `examples/`, and example modules for common setups.

Verified against real AWS — 19 types across 10 services, each created, updated, discovered and deleted in a real account, with the account checked for leftovers afterwards:

| Service | Types |
| --- | --- |
| EC2 | VPC, subnet, security group |
| IAM | role |
| S3 | bucket |
| ECR | repository |
| ECS | cluster, task definition, service |
| Lambda | function |
| CloudWatch Logs | log group |
| RDS | DB instance, DB subnet group, DB parameter group |
| Elastic Load Balancing v2 | load balancer, target group, listener |
| Route 53 | hosted zone, record set |

These runs are worth their cost: the Lambda one found a bug that would have left every inline-code function
planning a change forever, and it needed fixes in both the generator and the provider. A write-only test already
existed and stayed green throughout, because it only covered AWS omitting a property rather than returning an
empty one.

## Next

1. **Pending references.** 46 uncertain reference matches still await review (`docs/references-review.md`). These are judgement calls about whether an attribute really points at another resource type, and a wrong accept writes a bad reference into the catalog, so they wait for a human.
2. **More live coverage, chosen by behaviour rather than by service.** Every finding so far came from a *shape*, not a
   service: async handlers (RDS), rewritten values (IAM, ECR, RDS), silent cross-resource coupling (RDS parameter
   group family, ALB subnets), parent-only listing (ELB listener, Route 53 record set), deregister-not-delete (ECS
   task definition), and write-only values (Lambda). A type that is just another CRUD box teaches nothing. The
   shapes still untested, all free or nearly so: cross-resource wiring and a subscription flow (SQS queue, SNS
   topic and subscription); targets naming other resources (EventBridge rule, Scheduler schedule); a JSON document
   as a property (Step Functions state machine); deeply nested structures (API Gateway v2); composite IDs and
   untagged types (EC2 route table, route, internet gateway). Deliberately avoided: KMS keys, whose mandatory
   7–30 day deletion window means the suite cannot clean up after itself, and ACM certificates, which never reach
   `ISSUED` without DNS validation. Nothing is listed above on the strength of a passing fake.
3. **Types Cloud Control cannot manage at all**, which belongs in the docs before a user concludes the provider is
   broken. Six types in AWS's own bundle declare create/delete/update but no read handler, so they are correctly
   excluded from the catalog: `AWS::IAM::Policy`, `AWS::SNS::TopicPolicy`, `AWS::SQS::QueuePolicy`,
   `AWS::EC2::NetworkAclEntry`, `AWS::EC2::VPNGatewayRoutePropagation`, `AWS::AmazonMQ::ConfigurationAssociation`.
   **Documented 2026-09-20 in `docs/README.md`, and the open question is closed.** SNS and SQS both have a
   supported equivalent: `TopicInlinePolicy` and `QueueInlinePolicy`. IAM is covered by `RolePolicy`,
   `UserPolicy`, `GroupPolicy` and `ManagedPolicy`. The other three have no equivalent and cannot until AWS
   publishes a read handler.

## Later

- References on nested properties, once infrena supports them.
- Discovery for types that can only be listed under a parent resource.
- Schema differences between regions (the us-east-1 schemas are used for every region today).
- Moving docs to the wiki and a website at launch.
