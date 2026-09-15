# aws.serviceprincipalname

**CloudFormation type:** `AWS::PCAConnectorAD::ServicePrincipalName`

Definition of AWS::PCAConnectorAD::ServicePrincipalName Resource Type

Region attribute: `region`

**Import ID:** `<region>/ConnectorArn|DirectoryRegistrationArn` (AWS::PCAConnectorAD::ServicePrincipalName)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConnectorArn` | connector_arn | `string` | required, replaces on change | aws.pcaconnectorad.connector.ConnectorArn |  |
| `DirectoryRegistrationArn` | directory_registration_arn | `string` | required, replaces on change | aws.directoryregistration.DirectoryRegistrationArn |  |

Supports update: no

Discovery: supported (parent resource required)
