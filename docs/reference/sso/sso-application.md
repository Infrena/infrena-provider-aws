# aws.sso.application

**CloudFormation type:** `AWS::SSO::Application`

Resource Type definition for Identity Center (SSO) Application

Region attribute: `region`

**Import ID:** `<region>/ApplicationArn` (AWS::SSO::Application)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationArn` | application_arn | `string` | computed |  | The Application ARN that is returned upon creation of the Identity Center (SSO) Application |
| `ApplicationProviderArn` | application_provider_arn | `string` | required, replaces on change |  | The ARN of the application provider under which the operation will run |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description information for the Identity Center (SSO) Application |
| `IdentityStoreArn` | identity_store_arn | `string` | computed |  | The ARN of the identity store associated with the Identity Center instance |
| `InstanceArn` | instance_arn | `string` | required, replaces on change | aws.sso.instance.InstanceArn | The ARN of the instance of IAM Identity Center under which the operation will run |
| `Name` |  | `string` | required |  | The name you want to assign to this Identity Center (SSO) Application |
| `PortalOptions` | portal_options | `map` | optional, computed, provider-chosen |  | A structure that describes the options for the access portal associated with an application |
| `Status` |  | `string` | optional, computed, provider-chosen |  | Specifies whether the application is enabled or disabled |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported (parent resource required)
