# aws.applicationinstance

**CloudFormation type:** `AWS::Panorama::ApplicationInstance`

Creates an application instance and deploys it to a device.

Region attribute: `region`

**Import ID:** `<region>/ApplicationInstanceId` (AWS::Panorama::ApplicationInstance)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationInstanceId` | application_instance_id | `string` | computed |  |  |
| `ApplicationInstanceIdToReplace` | application_instance_id_to_replace | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The ID of an application instance to replace with the new instance. |
| `Arn` |  | `string` | computed |  |  |
| `CreatedTime` | created_time | `integer` | computed |  |  |
| `DefaultRuntimeContextDevice` | default_runtime_context_device | `string` | required, replaces on change |  | The device's ID. |
| `DefaultRuntimeContextDeviceName` | default_runtime_context_device_name | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A description for the application instance. |
| `HealthStatus` | health_status | `string` | computed |  |  |
| `LastUpdatedTime` | last_updated_time | `integer` | computed |  |  |
| `ManifestOverridesPayload` | manifest_overrides_payload | `map` | optional, computed, provider-chosen, replaces on change |  | Parameter overrides for an application instance. This is a JSON document that has a single key (``PayloadData``) where the value is an escaped string representation of the overrides document. |
| `ManifestPayload` | manifest_payload | `map` | required, replaces on change |  | A application verion's manifest file. This is a JSON document that has a single key (``PayloadData``) where the value is an escaped string representation of the application manifest (``graph.json``). This file is located in the ``graphs`` folder in your application source. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A name for the application instance. |
| `RuntimeRoleArn` | runtime_role_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.role.Arn | The ARN of a runtime role for the application instance. |
| `Status` |  | `string` | computed |  |  |
| `StatusDescription` | status_description | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | List of tags |

Supports update: yes

Discovery: supported
