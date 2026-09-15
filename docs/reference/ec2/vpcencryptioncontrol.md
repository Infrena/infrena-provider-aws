# aws.vpcencryptioncontrol

**CloudFormation type:** `AWS::EC2::VPCEncryptionControl`

Resource Type definition for AWS::EC2::VPCEncryptionControl

Region attribute: `region`

**Import ID:** `<region>/VpcEncryptionControlId` (AWS::EC2::VPCEncryptionControl)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `EgressOnlyInternetGatewayExclusionInput` | egress_only_internet_gateway_exclusion_input | `string` | optional, computed, provider-chosen, write-only |  | Used to enable or disable EIGW exclusion |
| `ElasticFileSystemExclusionInput` | elastic_file_system_exclusion_input | `string` | optional, computed, provider-chosen, write-only |  | Used to enable or disable EFS exclusion |
| `InternetGatewayExclusionInput` | internet_gateway_exclusion_input | `string` | optional, computed, provider-chosen, write-only |  | Used to enable or disable IGW exclusion |
| `LambdaExclusionInput` | lambda_exclusion_input | `string` | optional, computed, provider-chosen, write-only |  | Used to enable or disable Lambda exclusion |
| `Mode` |  | `string` | optional, computed, provider-chosen |  | The VPC encryption control mode, either monitor or enforce. |
| `NatGatewayExclusionInput` | nat_gateway_exclusion_input | `string` | optional, computed, provider-chosen, write-only |  | Used to enable or disable Nat gateway exclusion |
| `ResourceExclusions` | resource_exclusions | `map` | computed |  | Enumerates the states of all the VPC encryption control resource exclusions |
| `State` |  | `string` | computed |  | The current state of the VPC encryption control. |
| `StateMessage` | state_message | `string` | computed |  | Provides additional context on the state of the VPC encryption control. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags to assign to the VPC encryption control. |
| `VirtualPrivateGatewayExclusionInput` | virtual_private_gateway_exclusion_input | `string` | optional, computed, provider-chosen, write-only |  | Used to enable or disable VGW exclusion |
| `VpcEncryptionControlId` | vpc_encryption_control_id | `string` | computed |  | The VPC encryption control resource id. |
| `VpcId` | vpc_id | `string` | optional, computed, provider-chosen, replaces on change | aws.vpc.VpcId | The VPC on which this VPC encryption control is applied. |
| `VpcLatticeExclusionInput` | vpc_lattice_exclusion_input | `string` | optional, computed, provider-chosen, write-only |  | Used to enable or disable Vpc Lattice exclusion |
| `VpcPeeringExclusionInput` | vpc_peering_exclusion_input | `string` | optional, computed, provider-chosen, write-only |  | Used to enable or disable VPC peering exclusion |

Supports update: yes

Discovery: supported
