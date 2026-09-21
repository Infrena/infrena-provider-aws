# aws.guardduty.detector

**CloudFormation type:** `AWS::GuardDuty::Detector`

Resource Type definition for AWS::GuardDuty::Detector

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::GuardDuty::Detector)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DataSources` | data_sources | `map` | optional, computed, provider-chosen |  |  |
| `Enable` |  | `boolean` | required |  |  |
| `Features` |  | `list` | optional, computed, provider-chosen |  |  |
| `FindingPublishingFrequency` | finding_publishing_frequency | `string` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
