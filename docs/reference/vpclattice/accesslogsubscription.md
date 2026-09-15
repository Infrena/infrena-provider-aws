# aws.accesslogsubscription

**CloudFormation type:** `AWS::VpcLattice::AccessLogSubscription`

Enables access logs to be sent to Amazon CloudWatch, Amazon S3, and Amazon Kinesis Data Firehose. The service network owner can use the access logs to audit the services in the network. The service network owner will only see access logs from clients and services that are associated with their service network. Access log entries represent traffic originated from VPCs associated with that network.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::VpcLattice::AccessLogSubscription)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `DestinationArn` | destination_arn | `string` | required |  |  |
| `Id` |  | `string` | computed |  |  |
| `ResourceArn` | resource_arn | `string` | computed |  |  |
| `ResourceId` | resource_id | `string` | computed |  |  |
| `ResourceIdentifier` | resource_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `ServiceNetworkLogType` | service_network_log_type | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported (parent resource required)
