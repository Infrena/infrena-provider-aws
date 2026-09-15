# aws.serverlessrepo.application

**CloudFormation type:** `AWS::ServerlessRepo::Application`

Resource type definition for an AWS Serverless Application Repository application.

Region attribute: `region`

**Import ID:** `<region>/ApplicationId` (AWS::ServerlessRepo::Application)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationId` | application_id | `string` | computed |  | The Amazon Resource Name (ARN) of the application. |
| `Author` |  | `string` | required |  | The name of the author publishing the app. |
| `CreationTime` | creation_time | `string` | computed |  | The date and time this resource was created. |
| `Description` |  | `string` | required |  | The description of the application. |
| `HomePageUrl` | home_page_url | `string` | optional, computed, provider-chosen |  | A URL with more information about the application. |
| `IsVerifiedAuthor` | is_verified_author | `boolean` | computed |  | Whether the author of this application has been verified. |
| `Labels` |  | `list` | optional, computed, provider-chosen |  | Labels to improve discovery of apps in search results. |
| `LicenseBody` | license_body | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | A local text file that contains the license of the app. |
| `Name` |  | `string` | required, replaces on change |  | The name of the application. |
| `ReadmeBody` | readme_body | `string` | optional, computed, provider-chosen, write-only |  | A text readme file in Markdown language that contains a more detailed description of the application. |
| `SemanticVersion` | semantic_version | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The semantic version of the application. |
| `SourceCodeUrl` | source_code_url | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | A link to a public repository for the source code of your application. |
| `SpdxLicenseId` | spdx_license_id | `string` | optional, computed, provider-chosen, replaces on change |  | A valid identifier from https://spdx.org/licenses/. |
| `TemplateBody` | template_body | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The local raw packaged AWS SAM template file of your application. |

Supports update: yes

Discovery: supported
