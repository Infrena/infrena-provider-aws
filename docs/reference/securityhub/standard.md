# aws.standard

**CloudFormation type:** `AWS::SecurityHub::Standard`

The ``AWS::SecurityHub::Standard`` resource specifies the enablement of a security standard. The standard is identified by the ``StandardsArn`` property. To view a list of ASH standards and their Amazon Resource Names (ARNs), use the [DescribeStandards](https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_DescribeStandards.html) API operation.

Region attribute: `region`

**Import ID:** `<region>/StandardsSubscriptionArn` (AWS::SecurityHub::Standard)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DisabledStandardsControls` | disabled_standards_controls | `list` | optional, computed, provider-chosen |  | Specifies which controls are to be disabled in a standard. |
| `StandardsArn` | standards_arn | `string` | required, replaces on change |  | The ARN of the standard that you want to enable. To view a list of available ASH standards and their ARNs, use the [DescribeStandards](https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_DescribeStandards.html) API operation. |
| `StandardsSubscriptionArn` | standards_subscription_arn | `string` | computed |  |  |

Supports update: yes

Discovery: supported
