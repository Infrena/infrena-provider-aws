# aws.ecr.registrypolicy

**CloudFormation type:** `AWS::ECR::RegistryPolicy`

The ``AWS::ECR::RegistryPolicy`` resource creates or updates the permissions policy for a private registry.

Region attribute: `region`

**Import ID:** `<region>/RegistryId` (AWS::ECR::RegistryPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `PolicyText` | policy_text | `map` | required |  | The JSON policy text for your registry. |
| `RegistryId` | registry_id | `string` | computed |  | The registry id. |

Supports update: yes

Discovery: supported
