# aws.domainverification

**CloudFormation type:** `AWS::VpcLattice::DomainVerification`

Creates a Lattice Domain Verification

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::VpcLattice::DomainVerification)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `DomainName` | domain_name | `string` | required, replaces on change |  |  |
| `Id` |  | `string` | computed |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `TxtMethodConfig` | txt_method_config | `map` | computed |  |  |

Supports update: yes

Discovery: supported
