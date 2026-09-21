# aws.browserprofile

**CloudFormation type:** `AWS::BedrockAgentCore::BrowserProfile`

Resource definition for AWS::BedrockAgentCore::BrowserProfile

Region attribute: `region`

**Import ID:** `<region>/ProfileId` (AWS::BedrockAgentCore::BrowserProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | Timestamp when the browser profile was created. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The description of the browser profile. |
| `LastSavedAt` | last_saved_at | `string` | computed |  | Timestamp when the browser profile was last saved. |
| `LastSavedBrowserId` | last_saved_browser_id | `string` | computed |  | ID of the last saved browser. |
| `LastSavedBrowserSessionId` | last_saved_browser_session_id | `string` | computed |  | ID of the last saved browser session. |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | Timestamp when the browser profile was last updated. |
| `Name` |  | `string` | required, replaces on change |  | The name of the browser profile. |
| `ProfileArn` | profile_arn | `string` | computed |  | The ARN of a BrowserProfile resource. |
| `ProfileId` | profile_id | `string` | computed |  | The id of the browser profile. |
| `Status` |  | `string` | computed |  | Status of browser profile |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A map of tag keys and values |

Supports update: yes

Discovery: supported
