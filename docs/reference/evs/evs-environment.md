# aws.evs.environment

**CloudFormation type:** `AWS::EVS::Environment`

An environment created within the EVS service

Region attribute: `region`

**Import ID:** `<region>/EnvironmentId` (AWS::EVS::Environment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Checks` |  | `list` | computed |  |  |
| `ConnectivityInfo` | connectivity_info | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `Credentials` |  | `list` | computed |  |  |
| `EnvironmentArn` | environment_arn | `string` | computed |  |  |
| `EnvironmentId` | environment_id | `string` | computed |  |  |
| `EnvironmentName` | environment_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of an EVS environment |
| `EnvironmentState` | environment_state | `string` | computed |  |  |
| `Hosts` |  | `list` | optional, computed, provider-chosen, write-only |  | The initial hosts for environment only required upon creation. Modification after creation will have no effect |
| `InitialVlans` | initial_vlans | `map` | optional, computed, provider-chosen, write-only |  | The initial Vlan configuration only required upon creation. Modification after creation will have no effect |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `LicenseInfo` | license_info | `map` | optional, computed, provider-chosen, replaces on change |  | The license information for an EVS environment |
| `ModifiedAt` | modified_at | `string` | computed |  |  |
| `ServiceAccessSecurityGroups` | service_access_security_groups | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `ServiceAccessSubnetId` | service_access_subnet_id | `string` | required, replaces on change | aws.subnet.SubnetId |  |
| `SiteId` | site_id | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `StateDetails` | state_details | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `TermsAccepted` | terms_accepted | `boolean` | required, replaces on change |  |  |
| `VcfHostnames` | vcf_hostnames | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `VcfVersion` | vcf_version | `string` | required, replaces on change |  |  |
| `VpcId` | vpc_id | `string` | required, replaces on change | aws.vpc.VpcId |  |

Supports update: yes

Discovery: supported
