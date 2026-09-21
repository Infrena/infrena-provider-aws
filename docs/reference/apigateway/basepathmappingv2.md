# aws.basepathmappingv2

**CloudFormation type:** `AWS::ApiGateway::BasePathMappingV2`

Resource Type definition for AWS::ApiGateway::BasePathMappingV2

Region attribute: `region`

**Import ID:** `<region>/BasePathMappingArn` (AWS::ApiGateway::BasePathMappingV2)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BasePath` | base_path | `string` | optional, computed, provider-chosen, replaces on change |  | The base path name that callers of the API must provide in the URL after the domain name. |
| `BasePathMappingArn` | base_path_mapping_arn | `string` | computed |  | Amazon Resource Name (ARN) of the resource. |
| `DomainNameArn` | domain_name_arn | `string` | required, replaces on change | aws.apigateway.domainname.DomainNameArn | The Arn of an AWS::ApiGateway::DomainNameV2 resource. |
| `RestApiId` | rest_api_id | `string` | required | aws.restapi.RestApiId | The ID of the API. |
| `Stage` |  | `string` | optional, computed, provider-chosen |  | The name of the API's stage. |

Supports update: yes

Discovery: supported (parent resource required)
