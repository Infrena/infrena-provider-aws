# aws.vpcblockpublicaccessexclusion

**CloudFormation type:** `AWS::EC2::VPCBlockPublicAccessExclusion`

Resource Type definition for AWS::EC2::VPCBlockPublicAccessExclusion.

Region attribute: `region`

**Import ID:** `<region>/ExclusionId` (AWS::EC2::VPCBlockPublicAccessExclusion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ExclusionId` | exclusion_id | `string` | computed |  | The ID of the exclusion |
| `InternetGatewayExclusionMode` | internet_gateway_exclusion_mode | `string` | required |  | The desired Block Public Access Exclusion Mode for a specific VPC/Subnet. |
| `SubnetId` | subnet_id | `string` | optional, computed, provider-chosen, replaces on change | aws.subnet.SubnetId | The ID of the subnet. Required only if you don't specify VpcId |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `VpcId` | vpc_id | `string` | optional, computed, provider-chosen, replaces on change | aws.vpc.VpcId | The ID of the vpc. Required only if you don't specify SubnetId. |

Supports update: yes

Discovery: supported
