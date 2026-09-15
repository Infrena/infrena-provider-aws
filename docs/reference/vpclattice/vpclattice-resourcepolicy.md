# aws.vpclattice.resourcepolicy

**CloudFormation type:** `AWS::VpcLattice::ResourcePolicy`

Retrieves information about the resource policy. The resource policy is an IAM policy created by AWS RAM on behalf of the resource owner when they share a resource.

Region attribute: `region`

**Import ID:** `<region>/ResourceArn` (AWS::VpcLattice::ResourcePolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Policy` |  | `map` | required |  |  |
| `ResourceArn` | resource_arn | `string` | required, replaces on change |  |  |

Supports update: yes

Discovery: not supported
