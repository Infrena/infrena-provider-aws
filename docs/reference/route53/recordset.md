# aws.recordset

**CloudFormation type:** `AWS::Route53::RecordSet`

Resource Type definition for AWS::Route53::RecordSet.

Global type (no region attribute)

**Import ID:** `global/Name|HostedZoneId|Type|SetIdentifier` (AWS::Route53::RecordSet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AliasTarget` | alias_target | `map` | optional, computed, provider-chosen |  | Alias resource record sets only: Information about the AWS resource, such as a CloudFront distribution or an Amazon S3 bucket, that you want to route traffic to. |
| `CidrRoutingConfig` | cidr_routing_config | `map` | optional, computed, provider-chosen |  | The object that is specified in resource record set object when you are linking a resource record set to a CIDR location. |
| `Comment` |  | `string` | optional, computed, provider-chosen, write-only |  | Optional: Any comments you want to include about a change batch request. |
| `Failover` |  | `string` | optional, computed, provider-chosen |  | To configure failover, you add the Failover element to two resource record sets. For one resource record set, you specify PRIMARY as the value for Failover; for the other resource record set, you specify SECONDARY. In addition, you include the HealthCheckId element and specify the health check that you want Amazon Route 53 to perform for each resource record set. |
| `GeoLocation` | geo_location | `map` | optional, computed, provider-chosen |  | A complex type that lets you control how Amazon Route 53 responds to DNS queries based on the geographic origin of the query. |
| `HealthCheckId` | health_check_id | `string` | optional, computed, provider-chosen | aws.healthcheck.HealthCheckId | If you want Amazon Route 53 to return this resource record set in response to a DNS query only when the status of a health check is healthy, include the HealthCheckId element and specify the ID of the applicable health check. |
| `HostedZoneId` | hosted_zone_id | `string` | optional, computed, provider-chosen, replaces on change | aws.hostedzone.Id | The ID of the hosted zone that you want to create records in. |
| `HostedZoneName` | hosted_zone_name | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The name of the hosted zone that you want to create records in. You must include a trailing dot (for example, www.example.com.) as part of the HostedZoneName. |
| `MultiValueAnswer` | multi_value_answer | `boolean` | optional, computed, provider-chosen |  | To route traffic approximately randomly to multiple resources, such as web servers, create one multivalue answer record for each resource and specify true for MultiValueAnswer. |
| `Name` |  | `string` | required |  | The name of the record that you want to create, update, or delete. |
| `Region` |  | `string` | optional, computed, provider-chosen |  | The Amazon EC2 Region where you created the resource that this resource record set refers to. |
| `ResourceRecords` | resource_records | `list` | optional, computed, provider-chosen |  | One or more values that correspond with the value that you specified for the Type property. |
| `SetIdentifier` | set_identifier | `string` | optional, computed, provider-chosen |  | An identifier that differentiates among multiple resource record sets that have the same combination of name and type. |
| `TTL` |  | `string` | optional, computed, provider-chosen |  | The resource record cache time to live (TTL), in seconds. |
| `Type` | type_value | `string` | required |  | The DNS record type. |
| `Weight` |  | `integer` | optional, computed, provider-chosen |  | Among resource record sets that have the same combination of DNS name and type, a value that determines the proportion of DNS queries that Amazon Route 53 responds to using the current resource record set. Route 53 calculates the sum of the weights for the resource record sets that have the same combination of DNS name and type. Route 53 then responds to queries based on the ratio of a resource's weight to the total. |

Supports update: yes

Discovery: supported (parent resource required)
