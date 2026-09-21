# aws.servicediscovery.instance

**CloudFormation type:** `AWS::ServiceDiscovery::Instance`

Resource Type definition for AWS::ServiceDiscovery::Instance

Region attribute: `region`

**Import ID:** `<region>/ServiceId|InstanceId` (AWS::ServiceDiscovery::Instance)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `InstanceAttributes` | instance_attributes | `map` | required |  | A string map that contains information for the service that is specified in ServiceId. |
| `InstanceId` | instance_id | `string` | optional, computed, provider-chosen, replaces on change | aws.servicediscovery.instance.InstanceId | An identifier that you want to associate with the instance. |
| `ServiceId` | service_id | `string` | required, replaces on change | aws.servicediscovery.service.Id | The ID or Amazon Resource Name (ARN) of the service that you want to use for settings for the instance. |

Supports update: yes

Discovery: supported (parent resource required)
