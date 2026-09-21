# aws.distributiontenant

**CloudFormation type:** `AWS::CloudFront::DistributionTenant`

The distribution tenant.

Global type (no region attribute)

**Import ID:** `global/Id` (AWS::CloudFront::DistributionTenant)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `ConnectionGroupId` | connection_group_id | `string` | optional, computed, provider-chosen | aws.connectiongroup.Id | The ID of the connection group for the distribution tenant. If you don't specify a connection group, CloudFront uses the default connection group. |
| `CreatedTime` | created_time | `string` | computed |  |  |
| `Customizations` |  | `map` | optional, computed, provider-chosen |  | Customizations for the distribution tenant. For each distribution tenant, you can specify the geographic restrictions, and the Amazon Resource Names (ARNs) for the ACM certificate and WAF web ACL. These are specific values that you can override or disable from the multi-tenant distribution that was used to create the distribution tenant. |
| `DistributionId` | distribution_id | `string` | required | aws.cloudfront.distribution.Id | The ID of the multi-tenant distribution. |
| `DomainResults` | domain_results | `list` | computed |  |  |
| `Domains` |  | `list` | required |  | The domains associated with the distribution tenant. |
| `ETag` | e_tag | `string` | computed |  |  |
| `Enabled` |  | `boolean` | optional, computed, provider-chosen |  | Indicates whether the distribution tenant is in an enabled state. If disabled, the distribution tenant won't serve traffic. |
| `Id` |  | `string` | computed |  |  |
| `LastModifiedTime` | last_modified_time | `string` | computed |  |  |
| `ManagedCertificateRequest` | managed_certificate_request | `map` | optional, computed, provider-chosen, write-only |  | An object that represents the request for the Amazon CloudFront managed ACM certificate. |
| `Name` |  | `string` | required, replaces on change |  | The name of the distribution tenant. |
| `Parameters` |  | `list` | optional, computed, provider-chosen |  | A list of parameter values to add to the resource. A parameter is specified as a key-value pair. A valid parameter value must exist for any parameter that is marked as required in the multi-tenant distribution. |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  | A complex type that contains zero or more ``Tag`` elements. |

Supports update: yes

Discovery: supported
