# aws.dataexchange.dataset

**CloudFormation type:** `AWS::DataExchange::DataSet`

Definition of AWS::DataExchange::DataSet Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::DataExchange::DataSet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN for the data set. |
| `AssetType` | asset_type | `string` | required, replaces on change |  | The type of asset that is added to a data set. |
| `CreatedAt` | created_at | `string` | computed |  | The date and time that the data set was created, in ISO 8601 format. |
| `Description` |  | `string` | required |  | A description for the data set. |
| `Id` |  | `string` | computed |  | The unique identifier for the data set. |
| `Name` |  | `string` | required |  | The name of the data set. |
| `Origin` |  | `string` | computed |  | A property that defines the data set as OWNED by the account (for providers) or ENTITLED to the account (for subscribers). |
| `Tags` |  | `map` | tags map |  | Tags for the data set. |
| `UpdatedAt` | updated_at | `string` | computed |  | The date and time that the data set was last updated, in ISO 8601 format. |

Supports update: yes

Discovery: supported
