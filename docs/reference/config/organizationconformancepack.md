# aws.organizationconformancepack

**CloudFormation type:** `AWS::Config::OrganizationConformancePack`

Resource Type definition for AWS::Config::OrganizationConformancePack.

Region attribute: `region`

**Import ID:** `<region>/OrganizationConformancePackName` (AWS::Config::OrganizationConformancePack)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConformancePackInputParameters` | conformance_pack_input_parameters | `list` | optional, computed, provider-chosen |  | A list of ConformancePackInputParameter objects. |
| `DeliveryS3Bucket` | delivery_s3_bucket | `string` | optional, computed, provider-chosen |  | AWS Config stores intermediate files while processing conformance pack template. |
| `DeliveryS3KeyPrefix` | delivery_s3_key_prefix | `string` | optional, computed, provider-chosen |  | The prefix for the delivery S3 bucket. |
| `ExcludedAccounts` | excluded_accounts | `list` | optional, computed, provider-chosen |  | A list of AWS accounts to be excluded from an organization conformance pack while deploying a conformance pack. |
| `OrganizationConformancePackArn` | organization_conformance_pack_arn | `string` | computed |  | Amazon Resource Name (ARN) of the organization conformance pack. |
| `OrganizationConformancePackName` | organization_conformance_pack_name | `string` | required, replaces on change |  | The name of the organization conformance pack. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags for the organization conformance pack. |
| `TemplateBody` | template_body | `string` | optional, computed, provider-chosen, write-only |  | A string containing full conformance pack template body. |
| `TemplateS3Uri` | template_s3_uri | `string` | optional, computed, provider-chosen, write-only |  | Location of file containing the template body. |

Supports update: yes

Discovery: supported
