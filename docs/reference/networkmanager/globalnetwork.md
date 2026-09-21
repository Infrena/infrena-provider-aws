# aws.globalnetwork

**CloudFormation type:** `AWS::NetworkManager::GlobalNetwork`

The AWS::NetworkManager::GlobalNetwork type specifies a global network of the user's account

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::NetworkManager::GlobalNetwork)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the global network. |
| `CreatedAt` | created_at | `string` | optional, computed, provider-chosen |  | The date and time that the global network was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the global network. |
| `Id` |  | `string` | computed |  | The ID of the global network. |
| `State` |  | `string` | optional, computed, provider-chosen |  | The state of the global network. |
| `Tags` |  | `map` | tags map |  | The tags for the global network. |

Supports update: yes

Discovery: supported
