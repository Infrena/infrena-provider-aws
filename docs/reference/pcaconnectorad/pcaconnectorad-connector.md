# aws.pcaconnectorad.connector

**CloudFormation type:** `AWS::PCAConnectorAD::Connector`

Represents a Connector that connects AWS PrivateCA and your directory

Region attribute: `region`

**Import ID:** `<region>/ConnectorArn` (AWS::PCAConnectorAD::Connector)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CertificateAuthorityArn` | certificate_authority_arn | `string` | required, replaces on change |  |  |
| `ConnectorArn` | connector_arn | `string` | computed |  |  |
| `DirectoryId` | directory_id | `string` | required, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |
| `VpcInformation` | vpc_information | `map` | required, replaces on change |  |  |

Supports update: yes

Discovery: supported
