# aws.webapp

**CloudFormation type:** `AWS::Transfer::WebApp`

Resource Type definition for AWS::Transfer::WebApp

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Transfer::WebApp)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessEndpoint` | access_endpoint | `string` | optional, computed, provider-chosen |  | The AccessEndpoint is the URL that you provide to your users for them to interact with the Transfer Family web app. You can specify a custom URL or use the default value. |
| `Arn` |  | `string` | computed |  | Specifies the unique Amazon Resource Name (ARN) for the web app. |
| `EndpointDetails` | endpoint_details | `map` | optional, computed, provider-chosen, write-only |  |  |
| `IdentityProviderDetails` | identity_provider_details | `map` | required |  | You can provide a structure that contains the details for the identity provider to use with your web app. |
| `Tags` |  | `map` | tags map |  | Key-value pairs that can be used to group and search for web apps. |
| `VpcEndpointId` | vpc_endpoint_id | `string` | computed |  |  |
| `WebAppCustomization` | web_app_customization | `map` | optional, computed, provider-chosen |  |  |
| `WebAppEndpointPolicy` | web_app_endpoint_policy | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `WebAppId` | web_app_id | `string` | computed |  | A unique identifier for the web app. |
| `WebAppUnits` | web_app_units | `string` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
