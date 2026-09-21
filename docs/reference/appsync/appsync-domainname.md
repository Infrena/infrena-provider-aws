# aws.appsync.domainname

**CloudFormation type:** `AWS::AppSync::DomainName`

Resource Type definition for AWS::AppSync::DomainName

Region attribute: `region`

**Import ID:** `<region>/DomainName` (AWS::AppSync::DomainName)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AppSyncDomainName` | app_sync_domain_name | `string` | computed |  |  |
| `CertificateArn` | certificate_arn | `string` | required, replaces on change |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DomainName` | domain_name | `string` | required, replaces on change |  |  |
| `DomainNameArn` | domain_name_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the Domain Name. |
| `HostedZoneId` | hosted_zone_id | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  | An arbitrary set of tags (key-value pairs) for this Domain Name. |

Supports update: yes

Discovery: supported
