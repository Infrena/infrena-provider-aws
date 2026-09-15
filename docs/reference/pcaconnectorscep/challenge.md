# aws.challenge

**CloudFormation type:** `AWS::PCAConnectorSCEP::Challenge`

Represents a SCEP Challenge that is used for certificate enrollment

Region attribute: `region`

**Import ID:** `<region>/ChallengeArn` (AWS::PCAConnectorSCEP::Challenge)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ChallengeArn` | challenge_arn | `string` | computed |  |  |
| `ConnectorArn` | connector_arn | `string` | required, replaces on change | aws.pcaconnectorscep.connector.ConnectorArn |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported (parent resource required)
