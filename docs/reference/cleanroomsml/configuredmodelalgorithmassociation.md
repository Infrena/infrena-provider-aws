# aws.configuredmodelalgorithmassociation

**CloudFormation type:** `AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation`

Definition of AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation Resource Type

Region attribute: `region`

**Import ID:** `<region>/ConfiguredModelAlgorithmAssociationArn` (AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CollaborationIdentifier` | collaboration_identifier | `string` | computed |  |  |
| `ConfiguredModelAlgorithmArn` | configured_model_algorithm_arn | `string` | required, replaces on change | aws.configuredmodelalgorithm.ConfiguredModelAlgorithmArn |  |
| `ConfiguredModelAlgorithmAssociationArn` | configured_model_algorithm_association_arn | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `MembershipIdentifier` | membership_identifier | `string` | required, replaces on change |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `PrivacyConfiguration` | privacy_configuration | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An arbitrary set of tags (key-value pairs) for this cleanrooms-ml configured model algorithm association. |

Supports update: yes

Discovery: supported
