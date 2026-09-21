# aws.document

**CloudFormation type:** `AWS::SSM::Document`

The AWS::SSM::Document resource is an SSM document in AWS Systems Manager that defines the actions that Systems Manager performs, which can be used to set up and run commands on your instances.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::SSM::Document)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Attachments` |  | `list` | optional, computed, provider-chosen, write-only |  | A list of key and value pairs that describe attachments to a version of a document. |
| `Content` |  | `string` | required |  | The content for the Systems Manager document in JSON, YAML or String format. |
| `DocumentFormat` | document_format | `string` | optional, computed, provider-chosen |  | Specify the document format for the request. The document format can be either JSON or YAML. JSON is the default format. |
| `DocumentType` | document_type | `string` | optional, computed, provider-chosen, replaces on change |  | The type of document to create. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A name for the Systems Manager document. |
| `Requires` |  | `list` | optional, computed, provider-chosen |  | A list of SSM documents required by a document. For example, an ApplicationConfiguration document requires an ApplicationConfigurationSchema document. |
| `Tags` |  | `map` | tags map |  | Optional metadata that you assign to a resource. Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. |
| `TargetType` | target_type | `string` | optional, computed, provider-chosen |  | Specify a target type to define the kinds of resources the document can run on. |
| `UpdateMethod` | update_method | `string` | optional, computed, provider-chosen, write-only |  | Update method - when set to 'Replace', the update will replace the existing document; when set to 'NewVersion', the update will create a new version. |
| `VersionName` | version_name | `string` | optional, computed, provider-chosen |  | An optional field specifying the version of the artifact you are creating with the document. This value is unique across all versions of a document, and cannot be changed. |

Supports update: yes

Discovery: supported
