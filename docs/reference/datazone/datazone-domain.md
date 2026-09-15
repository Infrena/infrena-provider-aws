# aws.datazone.domain

**CloudFormation type:** `AWS::DataZone::Domain`

A domain is an organizing entity for connecting together assets, users, and their projects

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::DataZone::Domain)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the Amazon DataZone domain. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp of when the Amazon DataZone domain was last updated. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the Amazon DataZone domain. |
| `DomainExecutionRole` | domain_execution_role | `string` | optional, computed, provider-chosen |  | The domain execution role that is created when an Amazon DataZone domain is created. The domain execution role is created in the AWS account that houses the Amazon DataZone domain. |
| `DomainVersion` | domain_version | `string` | optional, computed, provider-chosen, replaces on change |  | The version of the domain. |
| `Id` |  | `string` | computed |  | The id of the Amazon DataZone domain. |
| `KmsKeyIdentifier` | kms_key_identifier | `string` | optional, computed, provider-chosen, replaces on change |  | The identifier of the AWS Key Management Service (KMS) key that is used to encrypt the Amazon DataZone domain, metadata, and reporting data. |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | The timestamp of when the Amazon DataZone domain was last updated. |
| `ManagedAccountId` | managed_account_id | `string` | computed |  | The identifier of the AWS account that manages the domain. |
| `Name` |  | `string` | required |  | The name of the Amazon DataZone domain. |
| `PortalUrl` | portal_url | `string` | computed |  | The URL of the data portal for this Amazon DataZone domain. |
| `RootDomainUnitId` | root_domain_unit_id | `string` | computed |  | The ID of the root domain in Amazon Datazone. |
| `ServiceRole` | service_role | `string` | optional, computed, provider-chosen, write-only |  | The service role of the domain that is created. |
| `SingleSignOn` | single_sign_on | `map` | optional, computed, provider-chosen |  | The single-sign on configuration of the Amazon DataZone domain. |
| `Status` |  | `string` | computed |  | The status of the Amazon DataZone domain. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags specified for the Amazon DataZone domain. |

Supports update: yes

Discovery: supported
