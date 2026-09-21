# aws.hubv2

**CloudFormation type:** `AWS::SecurityHub::HubV2`

The AWS::SecurityHub::HubV2 resource represents the implementation of the AWS Security Hub V2 service in your account. Only one hubv2 resource can created in each region in which you enable Security Hub V2.

Region attribute: `region`

**Import ID:** `<region>/HubV2Arn` (AWS::SecurityHub::HubV2)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `HubV2Arn` | hub_v2_arn | `string` | computed |  | The Amazon Resource Name of the Security Hub V2 resource. |
| `NetworkScanning` | network_scanning | `map` | optional, computed, provider-chosen |  | Configuration for the Network Scanning opt-in feature of Security Hub V2. Network Scanning is available in the AWS commercial partition only; specifying this property in another partition, such as AWS GovCloud (US) or China, fails. This property is desired state: if you remove it from a stack that previously set it, the feature is disabled. If a stack has never set it, the feature is left as-is, so a stack that does not manage Network Scanning will not disable it. Network Scanning requires Security Hub V2 to be enabled in the same account and Region. |
| `SubscribedAt` | subscribed_at | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A key-value pair to associate with the Security Hub V2 resource. You can specify a key that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. |

Supports update: yes

Discovery: supported
