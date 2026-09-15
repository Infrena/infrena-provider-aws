# aws.b2bi.transformer

**CloudFormation type:** `AWS::B2BI::Transformer`

Definition of AWS::B2BI::Transformer Resource Type

Region attribute: `region`

**Import ID:** `<region>/TransformerId` (AWS::B2BI::Transformer)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  |  |
| `EdiType` | edi_type | `string` | optional, computed, provider-chosen |  |  |
| `FileFormat` | file_format | `string` | optional, computed, provider-chosen |  |  |
| `InputConversion` | input_conversion | `map` | optional, computed, provider-chosen |  |  |
| `Mapping` |  | `map` | optional, computed, provider-chosen |  |  |
| `MappingTemplate` | mapping_template | `string` | optional, computed, provider-chosen |  | This shape is deprecated: This is a legacy trait. Please use input-conversion or output-conversion. |
| `ModifiedAt` | modified_at | `string` | computed |  |  |
| `Name` |  | `string` | required |  |  |
| `OutputConversion` | output_conversion | `map` | optional, computed, provider-chosen |  |  |
| `SampleDocument` | sample_document | `string` | optional, computed, provider-chosen |  | This shape is deprecated: This is a legacy trait. Please use input-conversion or output-conversion. |
| `SampleDocuments` | sample_documents | `map` | optional, computed, provider-chosen |  |  |
| `Status` |  | `string` | required |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TransformerArn` | transformer_arn | `string` | computed |  |  |
| `TransformerId` | transformer_id | `string` | computed |  |  |

Supports update: yes

Discovery: supported
