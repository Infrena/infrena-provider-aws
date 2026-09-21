# aws.aggregationauthorization

**CloudFormation type:** `AWS::Config::AggregationAuthorization`

Resource Type definition for AWS::Config::AggregationAuthorization

Region attribute: `region`

**Import ID:** `<region>/AuthorizedAccountId|AuthorizedAwsRegion` (AWS::Config::AggregationAuthorization)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AggregationAuthorizationArn` | aggregation_authorization_arn | `string` | computed |  | The ARN of the AggregationAuthorization. |
| `AuthorizedAccountId` | authorized_account_id | `string` | required, replaces on change |  | The 12-digit account ID of the account authorized to aggregate data. |
| `AuthorizedAwsRegion` | authorized_aws_region | `string` | required, replaces on change |  | The region authorized to collect aggregated data. |
| `Tags` |  | `map` | tags map |  | The tags for the AggregationAuthorization. |

Supports update: yes

Discovery: supported
