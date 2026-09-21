# aws.datacatalog

**CloudFormation type:** `AWS::Athena::DataCatalog`

Resource schema for AWS::Athena::DataCatalog

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Athena::DataCatalog)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConnectionType` | connection_type | `string` | optional, computed, provider-chosen |  | The type of connection for a FEDERATED data catalog |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the data catalog to be created. |
| `Error` |  | `string` | optional, computed, provider-chosen |  | Text of the error that occurred during data catalog creation or deletion. |
| `Name` |  | `string` | required, replaces on change |  | The name of the data catalog to create. The catalog name must be unique for the AWS account and can use a maximum of 128 alphanumeric, underscore, at sign, or hyphen characters. |
| `Parameters` |  | `map` | optional, computed, provider-chosen |  | Specifies the Lambda function or functions to use for creating the data catalog. This is a mapping whose values depend on the catalog type. |
| `Status` |  | `string` | optional, computed, provider-chosen |  | The status of the creation or deletion of the data catalog. LAMBDA, GLUE, and HIVE data catalog types are created synchronously. Their status is either CREATE_COMPLETE or CREATE_FAILED. The FEDERATED data catalog type is created asynchronously. |
| `Tags` |  | `map` | tags map |  | A list of comma separated tags to add to the data catalog that is created. |
| `Type` | type_value | `string` | required |  | The type of data catalog to create: LAMBDA for a federated catalog, GLUE for AWS Glue Catalog, or HIVE for an external hive metastore. FEDERATED is a federated catalog for which Athena creates the connection and the Lambda function for you based on the parameters that you pass. |

Supports update: yes

Discovery: supported
