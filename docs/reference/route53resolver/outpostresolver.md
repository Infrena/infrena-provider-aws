# aws.outpostresolver

**CloudFormation type:** `AWS::Route53Resolver::OutpostResolver`

Resource schema for AWS::Route53Resolver::OutpostResolver.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Route53Resolver::OutpostResolver)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The OutpostResolver ARN. |
| `CreationTime` | creation_time | `string` | computed |  | The OutpostResolver creation time |
| `CreatorRequestId` | creator_request_id | `string` | computed |  | The id of the creator request. |
| `Id` |  | `string` | computed |  | Id |
| `InstanceCount` | instance_count | `integer` | optional, computed, provider-chosen |  | The number of OutpostResolvers. |
| `ModificationTime` | modification_time | `string` | computed |  | The OutpostResolver last modified time |
| `Name` |  | `string` | required |  | The OutpostResolver name. |
| `OutpostArn` | outpost_arn | `string` | required, replaces on change |  | The Outpost ARN. |
| `PreferredInstanceType` | preferred_instance_type | `string` | required |  | The OutpostResolver instance type. |
| `Status` |  | `string` | computed |  | The OutpostResolver status, possible values are CREATING, OPERATIONAL, UPDATING, DELETING, ACTION_NEEDED, FAILED_CREATION and FAILED_DELETION. |
| `StatusMessage` | status_message | `string` | computed |  | The OutpostResolver status message. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
