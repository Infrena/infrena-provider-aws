# aws.lifecycleautomation

**CloudFormation type:** `AWS::SSMQuickSetup::LifecycleAutomation`

Resource Type definition for AWS::SSMQuickSetup::LifecycleAutomation that executes SSM Automation documents in response to CloudFormation lifecycle events.

Region attribute: `region`

**Import ID:** `<region>/AssociationId` (AWS::SSMQuickSetup::LifecycleAutomation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssociationId` | association_id | `string` | computed |  | The id from the association that is returned when creating the association |
| `AutomationDocument` | automation_document | `string` | required |  | The name of the Automation document to execute |
| `AutomationParameters` | automation_parameters | `map` | required |  | Parameters to be passed to the Automation Document |
| `ResourceKey` | resource_key | `string` | required, replaces on change |  | A unique identifier used for generating a unique logical ID for the custom resource |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
