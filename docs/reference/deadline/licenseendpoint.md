# aws.licenseendpoint

**CloudFormation type:** `AWS::Deadline::LicenseEndpoint`

Resource Type definition for AWS::Deadline::LicenseEndpoint

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Deadline::LicenseEndpoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `DnsName` | dns_name | `string` | computed |  |  |
| `LicenseEndpointId` | license_endpoint_id | `string` | computed |  |  |
| `SecurityGroupIds` | security_group_ids | `list` | required, replaces on change | aws.securitygroup.Id |  |
| `Status` |  | `string` | computed |  |  |
| `StatusMessage` | status_message | `string` | computed |  |  |
| `SubnetIds` | subnet_ids | `list` | required, replaces on change | aws.subnet.SubnetId |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `VpcId` | vpc_id | `string` | required, replaces on change | aws.vpc.VpcId |  |

Supports update: yes

Discovery: supported
