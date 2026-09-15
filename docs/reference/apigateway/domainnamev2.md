# aws.domainnamev2

**CloudFormation type:** `AWS::ApiGateway::DomainNameV2`

Resource Type definition for AWS::ApiGateway::DomainNameV2.

Region attribute: `region`

**Import ID:** `<region>/DomainNameArn` (AWS::ApiGateway::DomainNameV2)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CertificateArn` | certificate_arn | `string` | optional, computed, provider-chosen |  |  |
| `DomainName` | domain_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DomainNameArn` | domain_name_arn | `string` | computed |  | The amazon resource name (ARN) of the domain name resource. |
| `DomainNameId` | domain_name_id | `string` | computed |  |  |
| `EndpointAccessMode` | endpoint_access_mode | `string` | optional, computed, provider-chosen |  |  |
| `EndpointConfiguration` | endpoint_configuration | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `Policy` |  | `string` | optional, computed, provider-chosen |  |  |
| `RoutingMode` | routing_mode | `string` | optional, computed, provider-chosen |  | The valid routing modes are [BASE_PATH_MAPPING_ONLY], [ROUTING_RULE_THEN_BASE_PATH_MAPPING] and [ROUTING_RULE_ONLY]. All other inputs are invalid. |
| `SecurityPolicy` | security_policy | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
