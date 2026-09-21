# aws.cidrcollection

**CloudFormation type:** `AWS::Route53::CidrCollection`

Resource Type definition for AWS::Route53::CidrCollection.

Global type (no region attribute)

**Import ID:** `global/Id` (AWS::Route53::CidrCollection)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon resource name (ARN) to uniquely identify the AWS resource. |
| `Id` |  | `string` | computed |  | UUID of the CIDR collection. |
| `Locations` |  | `list` | optional, computed, provider-chosen |  | A complex type that contains information about the list of CIDR locations. |
| `Name` |  | `string` | required, replaces on change |  | A unique name for the CIDR collection. |

Supports update: yes

Discovery: supported
