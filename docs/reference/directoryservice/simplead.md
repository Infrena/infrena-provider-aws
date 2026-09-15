# aws.simplead

**CloudFormation type:** `AWS::DirectoryService::SimpleAD`

Resource Type definition for AWS::DirectoryService::SimpleAD

Region attribute: `region`

**Import ID:** `<region>/DirectoryId` (AWS::DirectoryService::SimpleAD)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Alias` |  | `string` | computed |  | The alias for a directory. |
| `CreateAlias` | create_alias | `boolean` | optional, computed, provider-chosen, replaces on change |  | The name of the configuration set. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Description for the directory. |
| `DirectoryId` | directory_id | `string` | computed |  | The unique identifier for a directory. |
| `DnsIpAddresses` | dns_ip_addresses | `list` | computed |  | The IP addresses of the DNS servers for the directory, such as [ "172.31.3.154", "172.31.63.203" ]. |
| `EnableSso` | enable_sso | `boolean` | optional, computed, provider-chosen |  | Whether to enable single sign-on for a Simple Active Directory in AWS. |
| `Name` |  | `string` | required, replaces on change |  | The fully qualified domain name for the AWS Managed Simple AD directory. |
| `Password` |  | `string` | optional, computed, provider-chosen, replaces on change, sensitive, write-only |  | The password for the default administrative user named Admin. |
| `ShortName` | short_name | `string` | optional, computed, provider-chosen, replaces on change |  | The NetBIOS name for your domain. |
| `Size` |  | `string` | required, replaces on change |  | The size of the directory. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `VpcSettings` | vpc_settings | `map` | required, replaces on change |  | VPC settings of the Simple AD directory server in AWS. |

Supports update: yes

Discovery: supported
