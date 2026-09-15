# aws.virtualmfadevice

**CloudFormation type:** `AWS::IAM::VirtualMFADevice`

Resource Type definition for AWS::IAM::VirtualMFADevice

Global type (no region attribute)

**Import ID:** `global/SerialNumber` (AWS::IAM::VirtualMFADevice)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Path` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `SerialNumber` | serial_number | `string` | computed |  |  |
| `Tags` |  | `list` | optional, computed, provider-chosen |  |  |
| `Users` |  | `list` | required |  |  |
| `VirtualMfaDeviceName` | virtual_mfa_device_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |

Supports update: yes

Discovery: supported
