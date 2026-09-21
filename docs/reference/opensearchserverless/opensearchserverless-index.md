# aws.opensearchserverless.index

**CloudFormation type:** `AWS::OpenSearchServerless::Index`

Resource Type definition for AWS::OpenSearchServerless::Index

Region attribute: `region`

**Import ID:** `<region>/IndexName|CollectionEndpoint` (AWS::OpenSearchServerless::Index)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CollectionEndpoint` | collection_endpoint | `string` | required, replaces on change |  | The endpoint for the collection. |
| `IndexName` | index_name | `string` | required, replaces on change |  | The name of the OpenSearch Serverless index. |
| `Mappings` |  | `map` | optional, computed, provider-chosen |  | Index Mappings |
| `Settings` |  | `map` | optional, computed, provider-chosen |  | Index settings |
| `Uuid` |  | `string` | computed |  | The unique identifier for the index. |

Supports update: yes

Discovery: supported (parent resource required)
