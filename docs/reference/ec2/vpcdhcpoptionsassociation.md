# aws.vpcdhcpoptionsassociation

**CloudFormation type:** `AWS::EC2::VPCDHCPOptionsAssociation`

Associates a set of DHCP options with a VPC, or associates no DHCP options with the VPC.

Region attribute: `region`

**Import ID:** `<region>/DhcpOptionsId|VpcId` (AWS::EC2::VPCDHCPOptionsAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DhcpOptionsId` | dhcp_options_id | `string` | required, replaces on change | aws.dhcpoptions.DhcpOptionsId | The ID of the DHCP options set, or default to associate no DHCP options with the VPC. |
| `VpcId` | vpc_id | `string` | required, replaces on change | aws.vpc.VpcId | The ID of the VPC. |

Supports update: yes

Discovery: supported
