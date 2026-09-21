# aws.connect.testcase

**CloudFormation type:** `AWS::Connect::TestCase`

Resource Type definition for AWS::Connect::TestCase

Region attribute: `region`

**Import ID:** `<region>/TestCaseArn` (AWS::Connect::TestCase)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Content` |  | `string` | required |  | The content of the test case. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the test case. |
| `EntryPoint` | entry_point | `map` | optional, computed, provider-chosen |  | The Entry Point associated with the test case |
| `InitializationData` | initialization_data | `string` | optional, computed, provider-chosen |  | The initialization data of the test case. |
| `InstanceArn` | instance_arn | `string` | required | aws.connect.instance.Arn | The identifier of the Amazon Connect instance. |
| `LastModifiedRegion` | last_modified_region | `string` | computed |  | Last modified region. |
| `LastModifiedTime` | last_modified_time | `float` | computed |  | A epoch time stamp field used for test case operations. |
| `Name` |  | `string` | required |  | The name of the test case. |
| `Status` |  | `string` | optional, computed, provider-chosen |  | The status of the test case. |
| `Tags` |  | `map` | tags map |  | One or more tags. |
| `TestCaseArn` | test_case_arn | `string` | computed |  | The identifier of the test case. |

Supports update: yes

Discovery: supported (parent resource required)
