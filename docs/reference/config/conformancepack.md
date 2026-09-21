# aws.conformancepack

**CloudFormation type:** `AWS::Config::ConformancePack`

A conformance pack is a collection of AWS Config rules and remediation actions that can be easily deployed as a single entity in an account and a region or across an entire AWS Organization.

Region attribute: `region`

**Import ID:** `<region>/ConformancePackName` (AWS::Config::ConformancePack)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConformancePackArn` | conformance_pack_arn | `string` | computed |  | Amazon Resource Name (ARN) of the conformance pack. |
| `ConformancePackInputParameters` | conformance_pack_input_parameters | `list` | optional, computed, provider-chosen |  | A list of ConformancePackInputParameter objects. |
| `ConformancePackName` | conformance_pack_name | `string` | required, replaces on change |  | Name of the conformance pack which will be assigned as the unique identifier. |
| `DeliveryS3Bucket` | delivery_s3_bucket | `string` | optional, computed, provider-chosen |  | AWS Config stores intermediate files while processing conformance pack template. |
| `DeliveryS3KeyPrefix` | delivery_s3_key_prefix | `string` | optional, computed, provider-chosen |  | The prefix for delivery S3 bucket. |
| `Tags` |  | `map` | tags map |  | The tags for the conformance pack. |
| `TemplateBody` | template_body | `string` | optional, computed, provider-chosen, write-only |  | A string containing full conformance pack template body. You can only specify one of the template body or template S3Uri fields. |
| `TemplateS3Uri` | template_s3_uri | `string` | optional, computed, provider-chosen, write-only |  | Location of file containing the template body which points to the conformance pack template that is located in an Amazon S3 bucket. You can only specify one of the template body or template S3Uri fields. |
| `TemplateSSMDocumentDetails` | template_ssm_document_details | `map` | optional, computed, provider-chosen, write-only |  | The TemplateSSMDocumentDetails object contains the name of the SSM document and the version of the SSM document. |

Supports update: yes

Discovery: supported
