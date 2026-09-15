# aws.classifier

**CloudFormation type:** `AWS::Glue::Classifier`

Resource Type definition for AWS::Glue::Classifier

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Glue::Classifier)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CsvClassifier` | csv_classifier | `map` | optional, computed, provider-chosen |  | A classifier for comma-separated values (CSV). |
| `GrokClassifier` | grok_classifier | `map` | optional, computed, provider-chosen |  | A classifier that uses grok. |
| `JsonClassifier` | json_classifier | `map` | optional, computed, provider-chosen |  | A classifier for JSON content. |
| `Name` |  | `string` | computed |  | One of XMLClassifier/Name, GrokClassifier/Name, JsonClassifier/Name or CsvClassifier/Name |
| `XMLClassifier` | xml_classifier | `map` | optional, computed, provider-chosen |  | A classifier for XML content. |

Supports update: yes

Discovery: supported
