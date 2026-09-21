# aws.accelerator

**CloudFormation type:** `AWS::GlobalAccelerator::Accelerator`

Resource Type definition for AWS::GlobalAccelerator::Accelerator

Region attribute: `region`

**Import ID:** `<region>/AcceleratorArn` (AWS::GlobalAccelerator::Accelerator)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcceleratorArn` | accelerator_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the accelerator. |
| `DnsName` | dns_name | `string` | computed |  | The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IPv4 addresses. |
| `DualStackDnsName` | dual_stack_dns_name | `string` | computed |  | The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IPv4 and IPv6 addresses. |
| `Enabled` |  | `boolean` | optional, computed, provider-chosen |  | Indicates whether an accelerator is enabled. The value is true or false. |
| `FlowLogsEnabled` | flow_logs_enabled | `boolean` | optional, computed, provider-chosen |  | Indicates whether flow logs are enabled for the accelerator. |
| `FlowLogsS3Bucket` | flow_logs_s3_bucket | `string` | optional, computed, provider-chosen |  | The name of the Amazon S3 bucket for the flow logs. |
| `FlowLogsS3Prefix` | flow_logs_s3_prefix | `string` | optional, computed, provider-chosen |  | The prefix for the location in the Amazon S3 bucket for the flow logs. |
| `IpAddressType` | ip_address_type | `string` | optional, computed, provider-chosen |  | IP Address type. |
| `IpAddresses` | ip_addresses | `list` | optional, computed, provider-chosen |  | The IP addresses from BYOIP Prefix pool. |
| `Ipv4Addresses` | ipv4_addresses | `list` | computed |  | The IPv4 addresses assigned to the accelerator. |
| `Ipv6Addresses` | ipv6_addresses | `list` | computed |  | The IPv6 addresses assigned if the accelerator is dualstack |
| `Name` |  | `string` | required |  | Name of accelerator. |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
