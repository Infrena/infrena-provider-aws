# aws.basepathmapping

**CloudFormation type:** `AWS::ApiGateway::BasePathMapping`

The ``AWS::ApiGateway::BasePathMapping`` resource creates a base path that clients who call your API must use in the invocation URL. Supported only for public custom domain names.

Region attribute: `region`

**Import ID:** `<region>/DomainName|BasePath` (AWS::ApiGateway::BasePathMapping)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BasePath` | base_path | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DomainName` | domain_name | `string` | required, replaces on change |  |  |
| `RestApiId` | rest_api_id | `string` | optional, computed, provider-chosen | aws.restapi.RestApiId |  |
| `Stage` |  | `string` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported (parent resource required)
