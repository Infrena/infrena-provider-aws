# aws.lightsail.instance

**CloudFormation type:** `AWS::Lightsail::Instance`

Resource Type definition for AWS::Lightsail::Instance

Region attribute: `region`

**Import ID:** `<region>/InstanceName` (AWS::Lightsail::Instance)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AddOns` | add_ons | `list` | optional, computed, provider-chosen |  | An array of objects representing the add-ons to enable for the new instance. |
| `AvailabilityZone` | availability_zone | `string` | optional, computed, provider-chosen, replaces on change |  | The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request. |
| `BlueprintId` | blueprint_id | `string` | required, replaces on change |  | The ID for a virtual private server image (e.g., app_wordpress_4_4 or app_lamp_7_0 ). Use the get blueprints operation to return a list of available images (or blueprints ). |
| `BundleId` | bundle_id | `string` | required, replaces on change |  | The bundle of specification information for your virtual private server (or instance ), including the pricing plan (e.g., micro_1_0 ). |
| `Hardware` |  | `map` | optional, computed, provider-chosen |  | Hardware of the Instance. |
| `InstanceArn` | instance_arn | `string` | computed |  |  |
| `InstanceName` | instance_name | `string` | required, replaces on change |  | The names to use for your new Lightsail instance. |
| `Ipv6Addresses` | ipv6_addresses | `list` | computed |  | IPv6 addresses of the instance |
| `IsStaticIp` | is_static_ip | `boolean` | computed |  | Is the IP Address of the Instance is the static IP |
| `KeyPairName` | key_pair_name | `string` | optional, computed, provider-chosen |  | The name of your key pair. |
| `Location` |  | `map` | optional, computed, provider-chosen |  | Location of a resource. |
| `Networking` |  | `map` | optional, computed, provider-chosen |  | Networking of the Instance. |
| `PrivateIpAddress` | private_ip_address | `string` | computed |  | Private IP Address of the Instance |
| `PublicIpAddress` | public_ip_address | `string` | computed |  | Public IP Address of the Instance |
| `ResourceType` | resource_type | `string` | computed |  | Resource type of Lightsail instance. |
| `SshKeyName` | ssh_key_name | `string` | computed |  | SSH Key Name of the  Lightsail instance. |
| `State` |  | `map` | optional, computed, provider-chosen |  | Current State of the Instance. |
| `SupportCode` | support_code | `string` | computed |  | Support code to help identify any issues |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `UserData` | user_data | `string` | optional, computed, provider-chosen, write-only |  | A launch script you can create that configures a server with additional user data. For example, you might want to run apt-get -y update. |
| `UserName` | user_name | `string` | computed |  | Username of the  Lightsail instance. |

Supports update: yes

Discovery: supported
