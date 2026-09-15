# aws.respondergateway

**CloudFormation type:** `AWS::RTBFabric::ResponderGateway`

Resource Type definition for AWS::RTBFabric::ResponderGateway Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::RTBFabric::ResponderGateway)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcmCertificateArn` | acm_certificate_arn | `string` | optional, computed, provider-chosen |  |  |
| `Arn` |  | `string` | computed |  |  |
| `CertificateAssociationStatus` | certificate_association_status | `string` | computed |  |  |
| `CreatedTimestamp` | created_timestamp | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DomainName` | domain_name | `string` | optional, computed, provider-chosen |  |  |
| `ExternalInboundEndpoint` | external_inbound_endpoint | `string` | computed |  |  |
| `GatewayId` | gateway_id | `string` | computed |  |  |
| `GatewayType` | gateway_type | `string` | optional, computed, provider-chosen |  |  |
| `ListenerConfig` | listener_config | `map` | optional, computed, provider-chosen |  |  |
| `ManagedEndpointConfiguration` | managed_endpoint_configuration | `map` | optional, computed, provider-chosen |  |  |
| `Port` |  | `integer` | required |  |  |
| `Protocol` |  | `string` | required |  |  |
| `ResponderGatewayStatus` | responder_gateway_status | `string` | computed |  |  |
| `SecurityGroupIds` | security_group_ids | `list` | required | aws.securitygroup.Id | The ID of one or more security groups in order to create a gateway. |
| `SubnetIds` | subnet_ids | `list` | required | aws.subnet.SubnetId | The ID of one or more subnets in order to create a gateway. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags to assign to the Responder Gateway. |
| `TrustStoreConfiguration` | trust_store_configuration | `map` | optional, computed, provider-chosen |  |  |
| `UpdatedTimestamp` | updated_timestamp | `string` | computed |  |  |
| `VpcId` | vpc_id | `string` | required | aws.vpc.VpcId |  |

Supports update: yes

Discovery: supported
