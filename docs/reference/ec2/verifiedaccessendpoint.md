# aws.verifiedaccessendpoint

**CloudFormation type:** `AWS::EC2::VerifiedAccessEndpoint`

The AWS::EC2::VerifiedAccessEndpoint resource creates an AWS EC2 Verified Access Endpoint.

Region attribute: `region`

**Import ID:** `<region>/VerifiedAccessEndpointId` (AWS::EC2::VerifiedAccessEndpoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationDomain` | application_domain | `string` | optional, computed, provider-chosen, replaces on change |  | The DNS name for users to reach your application. |
| `AttachmentType` | attachment_type | `string` | required, replaces on change |  | The type of attachment used to provide connectivity between the AWS Verified Access endpoint and the application. |
| `CidrOptions` | cidr_options | `map` | optional, computed, provider-chosen |  | The options for cidr type endpoint. |
| `CreationTime` | creation_time | `string` | computed |  | The creation time. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description for the AWS Verified Access endpoint. |
| `DeviceValidationDomain` | device_validation_domain | `string` | computed |  | Returned if endpoint has a device trust provider attached. |
| `DomainCertificateArn` | domain_certificate_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The ARN of a public TLS/SSL certificate imported into or created with ACM. |
| `EndpointDomain` | endpoint_domain | `string` | computed |  | A DNS name that is generated for the endpoint. |
| `EndpointDomainPrefix` | endpoint_domain_prefix | `string` | optional, computed, provider-chosen, replaces on change |  | A custom identifier that gets prepended to a DNS name that is generated for the endpoint. |
| `EndpointType` | endpoint_type | `string` | required, replaces on change |  | The type of AWS Verified Access endpoint. Incoming application requests will be sent to an IP address, load balancer or a network interface depending on the endpoint type specified.The type of AWS Verified Access endpoint. Incoming application requests will be sent to an IP address, load balancer or a network interface depending on the endpoint type specified. |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | The last updated time. |
| `LoadBalancerOptions` | load_balancer_options | `map` | optional, computed, provider-chosen |  | The load balancer details if creating the AWS Verified Access endpoint as load-balancertype. |
| `NetworkInterfaceOptions` | network_interface_options | `map` | optional, computed, provider-chosen |  | The options for network-interface type endpoint. |
| `PolicyDocument` | policy_document | `string` | optional, computed, provider-chosen |  | The AWS Verified Access policy document. |
| `PolicyEnabled` | policy_enabled | `boolean` | optional, computed, provider-chosen |  | The status of the Verified Access policy. |
| `RdsOptions` | rds_options | `map` | optional, computed, provider-chosen |  | The options for rds type endpoint. |
| `SecurityGroupIds` | security_group_ids | `list` | optional, computed, provider-chosen, replaces on change | aws.securitygroup.Id | The IDs of the security groups for the endpoint. |
| `SseSpecification` | sse_specification | `map` | optional, computed, provider-chosen |  | The configuration options for customer provided KMS encryption. |
| `Status` |  | `string` | computed |  | The endpoint status. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `VerifiedAccessEndpointId` | verified_access_endpoint_id | `string` | computed |  | The ID of the AWS Verified Access endpoint. |
| `VerifiedAccessGroupId` | verified_access_group_id | `string` | required | aws.verifiedaccessgroup.VerifiedAccessGroupId | The ID of the AWS Verified Access group. |
| `VerifiedAccessInstanceId` | verified_access_instance_id | `string` | computed |  | The ID of the AWS Verified Access instance. |

Supports update: yes

Discovery: supported
