# aws.policygrant

**CloudFormation type:** `AWS::DataZone::PolicyGrant`

Policy Grant in AWS DataZone is an explicit authorization assignment that allows a specific principal (user, group, or project) to perform particular actions (such as creating glossary terms, managing projects, or accessing resources) on governed resources within a certain scope (like a Domain Unit or Project). Policy Grants are essentially the mechanism by which DataZone enforces fine-grained, role-based access control beyond what is possible through AWS IAM alone.

Region attribute: `region`

**Import ID:** `<region>/DomainIdentifier|GrantId|EntityIdentifier|EntityType|PolicyType` (AWS::DataZone::PolicyGrant)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | Specifies the timestamp at which policy grant member was created. |
| `CreatedBy` | created_by | `string` | computed |  | Specifies the user who created the policy grant member. |
| `Detail` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DomainIdentifier` | domain_identifier | `string` | required, replaces on change |  |  |
| `EntityIdentifier` | entity_identifier | `string` | required, replaces on change |  |  |
| `EntityType` | entity_type | `string` | required, replaces on change |  |  |
| `GrantId` | grant_id | `string` | computed |  | The unique identifier of the policy grant returned by the AddPolicyGrant API |
| `PolicyType` | policy_type | `string` | required, replaces on change |  |  |
| `Principal` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |

Supports update: no

Discovery: supported
