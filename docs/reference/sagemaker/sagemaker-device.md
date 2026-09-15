# aws.sagemaker.device

**CloudFormation type:** `AWS::SageMaker::Device`

Resource schema for AWS::SageMaker::Device

Region attribute: `region`

**Import ID:** `<region>/Device/DeviceName` (AWS::SageMaker::Device)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Device` |  | `map` | optional, computed, provider-chosen |  | Edge device you want to create |
| `DeviceFleetName` | device_fleet_name | `string` | required |  | The name of the edge device fleet |
| `Tags` |  | `list` | optional, computed, provider-chosen |  | Associate tags with the resource |

Supports update: yes

Discovery: not supported
