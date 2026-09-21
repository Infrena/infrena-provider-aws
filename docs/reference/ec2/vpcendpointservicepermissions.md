# aws.vpcendpointservicepermissions

**CloudFormation type:** `AWS::EC2::VPCEndpointServicePermissions`

Resource Type definition for AWS::EC2::VPCEndpointServicePermissions

Region attribute: `region`

**Import ID:** `<region>/ServiceId` (AWS::EC2::VPCEndpointServicePermissions)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllowedPrincipals` | allowed_principals | `list` | optional, computed, provider-chosen |  |  |
| `ServiceId` | service_id | `string` | required, replaces on change |  |  |

Supports update: yes

Discovery: supported
