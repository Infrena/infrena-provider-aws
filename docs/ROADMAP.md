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

Verified against real AWS: VPC, subnet, security group, IAM role, S3 bucket, ECR repository, ECS cluster.

## Next

1. **Documentation.** A reference page for every type, generated from the catalog; detailed pages with examples for the common services; runnable `examples/`; example modules for common setups.
2. **System-owned resources.** Flag what AWS or CloudFormation owns so `import` doesn't adopt it by accident: the default VPC, default subnets, the default security group, service-linked roles, and CloudFormation-managed resources. The default VPC and subnets need the EC2 API, because Cloud Control does not report them.
3. **More live coverage.** RDS, load balancers, ECS services and task definitions, Route 53, Lambda. Each needs a real run before it is called verified.
4. **Pending references.** 46 uncertain reference matches still await review (`docs/references-review.md`).

## Later

- A least-privilege IAM policy for the live suite, generated from the types it uses.
- References on nested properties, once infrena supports them.
- Discovery for types that can only be listed under a parent resource.
- Schema differences between regions (the us-east-1 schemas are used for every region today).
- Moving docs to the wiki and a website at launch.
