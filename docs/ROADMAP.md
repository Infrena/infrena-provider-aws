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

Documentation ships with these: a generated reference page for every one of the 1,584 types, longer guides with examples for the common services, runnable `examples/`, and example modules for common setups.

Verified against real AWS — 18 types across 9 services, each created, updated, discovered and deleted in a real account, with the account checked for leftovers afterwards:

| Service | Types |
| --- | --- |
| EC2 | VPC, subnet, security group |
| IAM | role |
| S3 | bucket |
| ECR | repository |
| ECS | cluster, task definition, service |
| CloudWatch Logs | log group |
| RDS | DB instance, DB subnet group, DB parameter group |
| Elastic Load Balancing v2 | load balancer, target group, listener |
| Route 53 | hosted zone, record set |

## Next

1. **More live coverage.** Lambda is the last one outstanding. RDS, load balancers, Route 53, and ECS services and task definitions are done. Each needs a real run before it is called verified — nothing is listed above on the strength of a passing fake.
2. **Pending references.** 46 uncertain reference matches still await review (`docs/references-review.md`). These are judgement calls about whether an attribute really points at another resource type, and a wrong accept writes a bad reference into the catalog, so they wait for a human.

## Later

- References on nested properties, once infrena supports them.
- Discovery for types that can only be listed under a parent resource.
- Schema differences between regions (the us-east-1 schemas are used for every region today).
- Moving docs to the wiki and a website at launch.
