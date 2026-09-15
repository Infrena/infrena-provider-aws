# aws.dataautomationlibrary

**CloudFormation type:** `AWS::Bedrock::DataAutomationLibrary`

Resource Type definition for AWS::Bedrock::DataAutomationLibrary

Region attribute: `region`

**Import ID:** `<region>/LibraryArn` (AWS::Bedrock::DataAutomationLibrary)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | Time Stamp |
| `EncryptionConfiguration` | encryption_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | KMS Encryption Configuration |
| `EntityTypes` | entity_types | `list` | computed |  | List of info for each entity type in the DataAutomationLibrary |
| `LibraryArn` | library_arn | `string` | computed |  | ARN generated at the server side when a DataAutomationLibrary is created |
| `LibraryDescription` | library_description | `string` | optional, computed, provider-chosen |  | Description of the DataAutomationLibrary |
| `LibraryName` | library_name | `string` | required, replaces on change |  | Name of the DataAutomationLibrary |
| `Status` |  | `string` | computed |  | Status of DataAutomationLibrary |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | List of tags |

Supports update: yes

Discovery: supported
