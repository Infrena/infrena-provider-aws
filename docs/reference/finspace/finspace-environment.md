# aws.finspace.environment

**CloudFormation type:** `AWS::FinSpace::Environment`

An example resource schema demonstrating some basic constructs and validation rules.

Region attribute: `region`

**Import ID:** `<region>/EnvironmentId` (AWS::FinSpace::Environment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AwsAccountId` | aws_account_id | `string` | computed |  | AWS account ID associated with the Environment |
| `DataBundles` | data_bundles | `list` | optional, computed, provider-chosen, replaces on change |  | ARNs of FinSpace Data Bundles to install |
| `DedicatedServiceAccountId` | dedicated_service_account_id | `string` | computed |  | ID for FinSpace created account used to store Environment artifacts |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the Environment |
| `EnvironmentArn` | environment_arn | `string` | computed |  | ARN of the Environment |
| `EnvironmentId` | environment_id | `string` | computed |  | Unique identifier for representing FinSpace Environment |
| `EnvironmentUrl` | environment_url | `string` | computed |  | URL used to login to the Environment |
| `FederationMode` | federation_mode | `string` | optional, computed, provider-chosen |  | Federation mode used with the Environment |
| `FederationParameters` | federation_parameters | `map` | optional, computed, provider-chosen, replaces on change |  | Additional parameters to identify Federation mode |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | KMS key used to encrypt customer data within FinSpace Environment infrastructure |
| `Name` |  | `string` | required |  | Name of the Environment |
| `SageMakerStudioDomainUrl` | sage_maker_studio_domain_url | `string` | computed |  | SageMaker Studio Domain URL associated with the Environment |
| `Status` |  | `string` | computed |  | State of the Environment |
| `SuperuserParameters` | superuser_parameters | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | Parameters of the first Superuser for the FinSpace Environment |
| `Tags` |  | `map` | replaces on change, write-only, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
