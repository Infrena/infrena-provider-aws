# aws.networksettings

**CloudFormation type:** `AWS::WorkSpacesWeb::NetworkSettings`

Definition of AWS::WorkSpacesWeb::NetworkSettings Resource Type

Region attribute: `region`

**Import ID:** `<region>/NetworkSettingsArn` (AWS::WorkSpacesWeb::NetworkSettings)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssociatedPortalArns` | associated_portal_arns | `list` | computed |  |  |
| `NetworkSettingsArn` | network_settings_arn | `string` | computed |  |  |
| `SecurityGroupIds` | security_group_ids | `list` | required | aws.securitygroup.Id |  |
| `SubnetIds` | subnet_ids | `list` | required | aws.subnet.SubnetId |  |
| `Tags` |  | `map` | tags map |  |  |
| `VpcId` | vpc_id | `string` | required | aws.vpc.VpcId |  |

Supports update: yes

Discovery: supported
