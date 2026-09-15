# aws.directoryconfig

**CloudFormation type:** `AWS::AppStream::DirectoryConfig`

Resource Type definition for AWS::AppStream::DirectoryConfig

Region attribute: `region`

**Import ID:** `<region>/DirectoryName` (AWS::AppStream::DirectoryConfig)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CertificateBasedAuthProperties` | certificate_based_auth_properties | `map` | optional, computed, provider-chosen |  |  |
| `DirectoryName` | directory_name | `string` | required, replaces on change |  |  |
| `OrganizationalUnitDistinguishedNames` | organizational_unit_distinguished_names | `list` | required |  |  |
| `ServiceAccountCredentials` | service_account_credentials | `map` | required |  |  |

Supports update: yes

Discovery: supported
