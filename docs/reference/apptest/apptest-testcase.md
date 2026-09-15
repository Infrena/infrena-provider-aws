# aws.apptest.testcase

**CloudFormation type:** `AWS::AppTest::TestCase`

Represents a Test Case that can be captured and executed

Region attribute: `region`

**Import ID:** `<region>/TestCaseId` (AWS::AppTest::TestCase)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `LastUpdateTime` | last_update_time | `string` | computed |  |  |
| `LatestVersion` | latest_version | `map` | computed |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `Status` |  | `string` | computed |  |  |
| `Steps` |  | `list` | required |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |
| `TestCaseArn` | test_case_arn | `string` | computed |  |  |
| `TestCaseId` | test_case_id | `string` | computed |  |  |
| `TestCaseVersion` | test_case_version | `float` | computed |  |  |

Supports update: yes

Discovery: supported
