# aws.amplify.domain

**CloudFormation type:** `AWS::Amplify::Domain`

The AWS::Amplify::Domain resource allows you to connect a custom domain to your app.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Amplify::Domain)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AppId` | app_id | `string` | required, replaces on change | aws.amplify.app.AppId |  |
| `Arn` |  | `string` | computed |  |  |
| `AutoSubDomainCreationPatterns` | auto_sub_domain_creation_patterns | `list` | optional, computed, provider-chosen |  |  |
| `AutoSubDomainIAMRole` | auto_sub_domain_iam_role | `string` | optional, computed, provider-chosen |  |  |
| `Certificate` |  | `map` | computed |  |  |
| `CertificateRecord` | certificate_record | `string` | computed |  |  |
| `CertificateSettings` | certificate_settings | `map` | optional, computed, provider-chosen, write-only |  |  |
| `DomainName` | domain_name | `string` | required, replaces on change |  |  |
| `DomainStatus` | domain_status | `string` | computed |  |  |
| `EnableAutoSubDomain` | enable_auto_sub_domain | `boolean` | optional, computed, provider-chosen |  |  |
| `StatusReason` | status_reason | `string` | computed |  |  |
| `SubDomainSettings` | sub_domain_settings | `list` | required |  |  |
| `UpdateStatus` | update_status | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
