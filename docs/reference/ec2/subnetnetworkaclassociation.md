# aws.subnetnetworkaclassociation

**CloudFormation type:** `AWS::EC2::SubnetNetworkAclAssociation`

Resource Type definition for AWS::EC2::SubnetNetworkAclAssociation

Region attribute: `region`

**Import ID:** `<region>/AssociationId` (AWS::EC2::SubnetNetworkAclAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssociationId` | association_id | `string` | computed |  |  |
| `NetworkAclId` | network_acl_id | `string` | required, replaces on change | aws.networkacl.Id | The ID of the network ACL |
| `SubnetId` | subnet_id | `string` | required, replaces on change | aws.subnet.SubnetId | The ID of the subnet |

Supports update: no

Discovery: supported
