# aws.verifiedaccessinstance

**CloudFormation type:** `AWS::EC2::VerifiedAccessInstance`

The AWS::EC2::VerifiedAccessInstance resource creates an AWS EC2 Verified Access Instance.

Region attribute: `region`

**Import ID:** `<region>/VerifiedAccessInstanceId` (AWS::EC2::VerifiedAccessInstance)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CidrEndpointsCustomSubDomain` | cidr_endpoints_custom_sub_domain | `string` | optional, computed, provider-chosen |  | Introduce CidrEndpointsCustomSubDomain property to represent the domain (say, ava.my-company.com) |
| `CidrEndpointsCustomSubDomainNameServers` | cidr_endpoints_custom_sub_domain_name_servers | `list` | computed |  | Property to represent the name servers assoicated with the domain that AVA manages (say, ['ns1.amazonaws.com', 'ns2.amazonaws.com', 'ns3.amazonaws.com', 'ns4.amazonaws.com']). |
| `CreationTime` | creation_time | `string` | computed |  | Time this Verified Access Instance was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description for the AWS Verified Access instance. |
| `FipsEnabled` | fips_enabled | `boolean` | optional, computed, provider-chosen |  | Indicates whether FIPS is enabled |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | Time this Verified Access Instance was last updated. |
| `LoggingConfigurations` | logging_configurations | `map` | optional, computed, provider-chosen |  | The configuration options for AWS Verified Access instances. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `VerifiedAccessInstanceId` | verified_access_instance_id | `string` | computed |  | The ID of the AWS Verified Access instance. |
| `VerifiedAccessTrustProviderIds` | verified_access_trust_provider_ids | `list` | optional, computed, provider-chosen | aws.verifiedaccesstrustprovider.VerifiedAccessTrustProviderId | The IDs of the AWS Verified Access trust providers. |
| `VerifiedAccessTrustProviders` | verified_access_trust_providers | `list` | optional, computed, provider-chosen |  | AWS Verified Access trust providers. |

Supports update: yes

Discovery: supported
