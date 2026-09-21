# aws.rekognition.collection

**CloudFormation type:** `AWS::Rekognition::Collection`

The AWS::Rekognition::Collection type creates an Amazon Rekognition Collection. A collection is a logical grouping of information about detected faces which can later be referenced for searches on the group

Region attribute: `region`

**Import ID:** `<region>/CollectionId` (AWS::Rekognition::Collection)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CollectionId` | collection_id | `string` | required, replaces on change | aws.rekognition.collection.CollectionId | The name of the collection |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
