# aws.devicefleet

**CloudFormation type:** `AWS::SageMaker::DeviceFleet`

Resource schema for AWS::SageMaker::DeviceFleet

Region attribute: `region`

**Import ID:** `<region>/DeviceFleetName` (AWS::SageMaker::DeviceFleet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description for the edge device fleet |
| `DeviceFleetName` | device_fleet_name | `string` | required, replaces on change |  | The name of the edge device fleet |
| `OutputConfig` | output_config | `map` | required |  | S3 bucket and an ecryption key id (if available) to store outputs for the fleet |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | Role associated with the device fleet |
| `Tags` |  | `list` | optional, computed, provider-chosen |  | Associate tags with the resource |

Supports update: yes

Discovery: not supported
