# aws.pcaconnectorscep.connector

**CloudFormation type:** `AWS::PCAConnectorSCEP::Connector`

Represents a Connector that allows certificate issuance through Simple Certificate Enrollment Protocol (SCEP)

Region attribute: `region`

**Import ID:** `<region>/ConnectorArn` (AWS::PCAConnectorSCEP::Connector)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CertificateAuthorityArn` | certificate_authority_arn | `string` | required, replaces on change |  |  |
| `ConnectorArn` | connector_arn | `string` | computed |  |  |
| `Endpoint` |  | `string` | computed |  |  |
| `MobileDeviceManagement` | mobile_device_management | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `OpenIdConfiguration` | open_id_configuration | `map` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |
| `Type` | type_value | `string` | computed |  |  |
| `VpcEndpointId` | vpc_endpoint_id | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |

Supports update: yes

Discovery: supported
