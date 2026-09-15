# aws.dnssec

**CloudFormation type:** `AWS::Route53::DNSSEC`

Resource used to control (enable/disable) DNSSEC in a specific hosted zone.

Global type (no region attribute)

**Import ID:** `global/HostedZoneId` (AWS::Route53::DNSSEC)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `HostedZoneId` | hosted_zone_id | `string` | required, replaces on change | aws.hostedzone.Id | The unique string (ID) used to identify a hosted zone. |

Supports update: no

Discovery: supported
