# aws.hostedzone

**CloudFormation type:** `AWS::Route53::HostedZone`

Creates a new public or private hosted zone. You create records in a public hosted zone to define how you want to route traffic on the internet for a domain, such as example.com, and its subdomains (apex.example.com, acme.example.com). You create records in a private hosted zone to define how you want to route traffic for a domain and its subdomains within one or more Amazon Virtual Private Clouds (Amazon VPCs). 

Global type (no region attribute)

**Import ID:** `global/Id` (AWS::Route53::HostedZone)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `HostedZoneConfig` | hosted_zone_config | `map` | optional, computed, provider-chosen |  | A complex type that contains an optional comment about your hosted zone. If you don't want to specify a comment, omit both the ``HostedZoneConfig`` and ``Comment`` elements. |
| `HostedZoneFeatures` | hosted_zone_features | `map` | optional, computed, provider-chosen |  | Represents the features configuration for a hosted zone, including the status of various features and any associated failure reasons. |
| `HostedZoneTags` | hosted_zone_tags | `map` | tags map |  | Adds, edits, or deletes tags for a health check or a hosted zone. |
| `Id` |  | `string` | computed |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the domain. Specify a fully qualified domain name, for example, *www.example.com*. The trailing dot is optional; Amazon Route 53 assumes that the domain name is fully qualified. This means that Route 53 treats *www.example.com* (without a trailing dot) and *www.example.com.* (with a trailing dot) as identical. |
| `NameServers` | name_servers | `list` | computed |  |  |
| `QueryLoggingConfig` | query_logging_config | `map` | optional, computed, provider-chosen |  | A complex type that contains information about a configuration for DNS query logging. |
| `VPCs` | vp_cs | `list` | optional, computed, provider-chosen |  | *Private hosted zones:* A complex type that contains information about the VPCs that are associated with the specified hosted zone. |

Supports update: yes

Discovery: supported
