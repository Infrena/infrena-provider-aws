# aws.opensearchservice.domain

**CloudFormation type:** `AWS::OpenSearchService::Domain`

An example resource schema demonstrating some basic constructs and validation rules.

Region attribute: `region`

**Import ID:** `<region>/DomainName` (AWS::OpenSearchService::Domain)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AIMLOptions` | aiml_options | `map` | optional, computed, provider-chosen |  |  |
| `AccessPolicies` | access_policies | `map` | optional, computed, provider-chosen |  |  |
| `AdvancedOptions` | advanced_options | `map` | optional, computed, provider-chosen |  |  |
| `AdvancedSecurityOptions` | advanced_security_options | `map` | optional, computed, provider-chosen |  |  |
| `Arn` |  | `string` | computed |  |  |
| `AutomatedSnapshotPauseOptions` | automated_snapshot_pause_options | `map` | optional, computed, provider-chosen |  |  |
| `ClusterConfig` | cluster_config | `map` | optional, computed, provider-chosen |  |  |
| `CognitoOptions` | cognito_options | `map` | optional, computed, provider-chosen |  |  |
| `DeploymentStrategyOptions` | deployment_strategy_options | `map` | optional, computed, provider-chosen |  |  |
| `DomainArn` | domain_arn | `string` | computed |  |  |
| `DomainEndpoint` | domain_endpoint | `string` | computed |  |  |
| `DomainEndpointOptions` | domain_endpoint_options | `map` | optional, computed, provider-chosen |  |  |
| `DomainEndpointV2` | domain_endpoint_v2 | `string` | computed |  |  |
| `DomainEndpoints` | domain_endpoints | `map` | computed |  |  |
| `DomainName` | domain_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `EBSOptions` | ebs_options | `map` | optional, computed, provider-chosen |  |  |
| `EncryptionAtRestOptions` | encryption_at_rest_options | `map` | optional, computed, provider-chosen |  |  |
| `EngineMode` | engine_mode | `string` | optional, computed, provider-chosen, replaces on change |  | The engine mode of the domain. Determines whether the domain runs the standard (GENERAL) engine or the optimized multi-engine (OPTIMIZED) engine. This value cannot be changed after the domain is created. |
| `EngineVersion` | engine_version | `string` | optional, computed, provider-chosen |  |  |
| `IPAddressType` | ip_address_type | `string` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | computed |  |  |
| `IdentityCenterOptions` | identity_center_options | `map` | optional, computed, provider-chosen |  | Options for configuring Identity Center |
| `LogPublishingOptions` | log_publishing_options | `map` | optional, computed, provider-chosen |  |  |
| `NodeToNodeEncryptionOptions` | node_to_node_encryption_options | `map` | optional, computed, provider-chosen |  |  |
| `OffPeakWindowOptions` | off_peak_window_options | `map` | optional, computed, provider-chosen |  |  |
| `ServiceSoftwareOptions` | service_software_options | `map` | computed |  |  |
| `SkipShardMigrationWait` | skip_shard_migration_wait | `boolean` | optional, computed, provider-chosen |  |  |
| `SnapshotOptions` | snapshot_options | `map` | optional, computed, provider-chosen |  |  |
| `SoftwareUpdateOptions` | software_update_options | `map` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An arbitrary set of tags (key-value pairs) for this Domain. |
| `UseCase` | use_case | `string` | optional, computed, provider-chosen |  | The primary use case of the domain. Determines the default configuration tuned for the workload. For GENERAL engine-mode domains, this value can be changed after creation. For OPTIMIZED engine-mode domains, this value cannot be changed after creation. |
| `VPCOptions` | vpc_options | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: not supported
