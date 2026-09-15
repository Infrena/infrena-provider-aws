# Reference edges: review

> **NOTHING IN TIER 2 IS APPROVED YET.** `references.approve_targets` in `gen/overlay.yaml` is empty, so all 175 tier-2 edges are `pending` in `gen/references.lock.json` and none of them reach the catalog. Approving is James's decision.

Generated on 2026-09-14 from `gen/references.lock.json` and the derivation over bundle `543f4b1846d7`, for the resource references design (infrena `docs/superpowers/specs/2026-09-14-resource-references-design.md` §4). The verdicts below are a first reading of property names and descriptions, not verified against AWS.

## How to act on this

- **Approve a target:** add its CloudFormation type to `references.approve_targets`. Every tier-2 edge pointing at it becomes `approved`.
- **Refuse an edge:** add `Type.Property` to `references.reject`. That works on tier 1 too, and a reject beats an approval, so a MIXED target can be approved with its bad edges rejected.
- Then run `go run ./cmd/gen-cloudcontrol` and commit the overlay, lock and catalog diff together.
- A schema refresh that derives an edge the lock does not hold fails generation. Review the new edges, then run it with `-accept-new-references`.

## Counts

| | this run | infrena session |
| --- | ---: | ---: |
| candidate properties | 1503 | 1614 |
| tier 1, exact match (accepted) | 816 | 832 |
| tier 2, suffix match (pending) | 175 | 161 |
| ambiguous | 285 | 399 |
| no such type | 143 | 159 |
| target lacks the attribute | 84 | 63 |
| resolved | 991 (65.9%) | 993 (61.5%) |

Why the numbers differ: the infrena session took candidate sources from every AWS type in the bundle. Over all 1729 AWS types there are 1617 candidates, which matches its 1614. This run takes sources and targets from the 1584 provisionable types only, the ones in the catalog. Which attribute of the target is picked, and in what order ambiguity and a missing attribute are checked, can also move edges between tiers and buckets. This run picks a property named exactly `Id` or `Arn`, then `<Segment>Id`/`<Segment>Arn`, then a single primary identifier ending in the kind. For example, `AWS::EC2::SecurityGroup.Id` takes 14 tier-2 edges here but is not in the session's top ten.

Only top-level properties are considered: infrena refuses references on nested attributes. There are 4377 writable nested properties named like references, in 534 provisionable types, and all of them are skipped.

All four `requirements:` in the overlay are backed by an accepted edge (Subnet.VpcId, RouteTable.VpcId, Route.RouteTableId, DBSubnetGroup.SubnetIds), so the cross-check wrote no warnings.

## Tier 2 by target type

| # | target | infrena type | edges | cumulative | verdict |
| ---: | --- | --- | ---: | ---: | --- |
| 1 | `AWS::IAM::Role` | `aws.role` | 83 | 47.4% | looks right: every source is a `*RoleArn` naming the role a service assumes |
| 2 | `AWS::EC2::SecurityGroup` | `aws.securitygroup` | 14 | 55.4% | looks right: `VpcSecurityGroupIds` and EMR Studio's engine and workspace groups are EC2 security group ids |
| 3 | `AWS::EC2::IPAMPool` | `aws.ipampool` | 8 | 60.0% | looks right |
| 4 | `AWS::EC2::Subnet` | `aws.subnet` | 7 | 64.0% | looks right |
| 5 | `AWS::DMS::Endpoint` | `aws.dms.endpoint` | 4 | 66.3% | looks right: source and target endpoints of DMS tasks |
| 6 | `AWS::EC2::RouteTable` | `aws.routetable` | 4 | 68.6% | MIXED: TransitGateway's AssociationDefaultRouteTableId and PropagationDefaultRouteTableId, and TransitGatewayPolicyTableEntry.TargetRouteTableId, hold transit gateway route table ids, not VPC route tables. Approve this target only together with rejects for those three |
| 7 | `AWS::EC2::TransitGatewayAttachment` | `aws.transitgatewayattachment` | 4 | 70.9% | looks right |
| 8 | `AWS::EC2::PrefixList` | `aws.prefixlist` | 3 | 72.6% | looks right (the id may also be an AWS-managed prefix list, which is still the same kind of id) |
| 9 | `AWS::EC2::VPC` | `aws.vpc` | 3 | 74.3% | looks right (a peer VPC can live in another account, and ClassicLinkVPCId is legacy) |
| 10 | `AWS::GroundStation::Config` | `aws.config` | 3 | 76.0% | MIXED: the two MissionProfile config arns are right; Route53Resolver's ResolverQueryLogConfigId -> Config.Id is FABRICATED |
| 11 | `AWS::ApiGateway::Resource` | `aws.resource` | 2 | 77.1% | FABRICATED: RDS resource ids (dbi-/cluster- ids) |
| 12 | `AWS::Connect::Notification` | `aws.notification` | 2 | 78.3% | FABRICATED: SNS topic arns and VPC endpoint connection notifications, nothing to do with Connect |
| 13 | `AWS::KMS::Key` | `aws.kms.key` | 2 | 79.4% | looks right: ReplicaKey.PrimaryKeyArn and Alias.TargetKeyId |
| 14 | `AWS::Lambda::Version` | `aws.version` | 2 | 80.6% | FABRICATED: CloudFormation extension type version arns |
| 15 | `AWS::Logs::LogGroup` | `aws.loggroup` | 2 | 81.7% | looks right |
| 16 | `AWS::Notifications::NotificationConfiguration` | `aws.notificationconfiguration` | 2 | 82.9% | probably FABRICATED: a managed notification configuration is AWS-owned, not this resource type |
| 17 | `AWS::ResilienceHubV2::System` | `aws.system` | 2 | 84.0% | FABRICATED: DBSystemId is an Oracle CDB name |
| 18 | `AWS::Transfer::Profile` | `aws.transfer.profile` | 2 | 85.1% | looks right |
| 19 | `AWS::AmazonMQ::Broker` | `aws.broker` | 1 | 85.7% | looks right (self reference) |
| 20 | `AWS::AmplifyUIBuilder::Form` | `aws.form` | 1 | 86.3% | FABRICATED: Signer's PlatformId is a signing platform |
| 21 | `AWS::BedrockAgentCore::Runtime` | `aws.runtime` | 1 | 86.9% | looks right |
| 22 | `AWS::Connect::ContactFlow` | `aws.contactflow` | 1 | 87.4% | looks right |
| 23 | `AWS::Connect::PhoneNumber` | `aws.connect.phonenumber` | 1 | 88.0% | looks right (self reference) |
| 24 | `AWS::Connect::Queue` | `aws.connect.queue` | 1 | 88.6% | looks right |
| 25 | `AWS::Deadline::Limit` | `aws.limit` | 1 | 89.1% | FABRICATED: a BedrockAgentCore gateway rate limit is not a Deadline limit |
| 26 | `AWS::Deadline::StorageProfile` | `aws.storageprofile` | 1 | 89.7% | looks right |
| 27 | `AWS::DevOpsAgent::Association` | `aws.devopsagent.association` | 1 | 90.3% | looks right (self reference) |
| 28 | `AWS::DirectConnect::Lag` | `aws.lag` | 1 | 90.9% | FABRICATED: an Outpost LAG, not a Direct Connect LAG |
| 29 | `AWS::EC2::TransitGateway` | `aws.transitgateway` | 1 | 91.4% | looks right |
| 30 | `AWS::EC2::Volume` | `aws.ec2.volume` | 1 | 92.0% | looks right (self reference) |
| 31 | `AWS::EKS::CertificateAuthority` | `aws.eks.certificateauthority` | 1 | 92.6% | unsure: check what Cluster.ActiveCertificateAuthorityId holds |
| 32 | `AWS::GameLift::Location` | `aws.location` | 1 | 93.1% | FABRICATED: a DataSync location, not a GameLift location |
| 33 | `AWS::Glue::Catalog` | `aws.catalog` | 1 | 93.7% | unsure: a Lake Formation catalog id is usually an account id |
| 34 | `AWS::Inspector::ResourceGroup` | `aws.resourcegroup` | 1 | 94.3% | FABRICATED: EC2 host resource groups belong to License Manager |
| 35 | `AWS::IoT::Certificate` | `aws.iot.certificate` | 1 | 94.9% | looks right |
| 36 | `AWS::IoTTwinMaker::Entity` | `aws.entity` | 1 | 95.4% | looks right (self reference) |
| 37 | `AWS::Lambda::Function` | `aws.lambda.function` | 1 | 96.0% | looks right |
| 38 | `AWS::MPA::ApprovalTeam` | `aws.approvalteam` | 1 | 96.6% | looks right |
| 39 | `AWS::MSK::Cluster` | `aws.msk.cluster` | 1 | 97.1% | looks right |
| 40 | `AWS::Panorama::Package` | `aws.package` | 1 | 97.7% | FABRICATED: Inspector rules packages |
| 41 | `AWS::QuickSight::Folder` | `aws.folder` | 1 | 98.3% | looks right (self reference) |
| 42 | `AWS::QuickSight::Theme` | `aws.quicksight.theme` | 1 | 98.9% | looks right (self reference) |
| 43 | `AWS::ServiceCatalog::Portfolio` | `aws.portfolio` | 1 | 99.4% | looks right |
| 44 | `AWS::WorkSpacesWeb::IdentityProvider` | `aws.identityprovider` | 1 | 100.0% | FABRICATED: QBusiness wants an IAM identity provider |

### 1. AWS::IAM::Role (83), at `Arn`

looks right: every source is a `*RoleArn` naming the role a service assumes

- `AWS::Amplify::App.ComputeRoleArn`
- `AWS::Amplify::Branch.ComputeRoleArn`
- `AWS::ApiGateway::Account.CloudWatchRoleArn`
- `AWS::AppConfig::ConfigurationProfile.RetrievalRoleArn`
- `AWS::AppStream::AppBlockBuilder.IamRoleArn`
- `AWS::AppStream::ImageBuilder.IamRoleArn`
- `AWS::AppSync::DataSource.ServiceRoleArn`
- `AWS::AppSync::GraphQLApi.MergedApiExecutionRoleArn`
- `AWS::Backup::RestoreTestingSelection.IamRoleArn`
- `AWS::Bedrock::Agent.AgentResourceRoleArn`
- `AWS::Bedrock::Flow.ExecutionRoleArn`
- `AWS::BedrockAgentCore::BrowserCustom.ExecutionRoleArn`
- `AWS::BedrockAgentCore::CodeInterpreterCustom.ExecutionRoleArn`
- `AWS::BedrockAgentCore::Harness.ExecutionRoleArn`
- `AWS::BedrockAgentCore::Memory.MemoryExecutionRoleArn`
- `AWS::BedrockAgentCore::OnlineEvaluationConfig.EvaluationExecutionRoleArn`
- `AWS::Budgets::BudgetsAction.ExecutionRoleArn`
- `AWS::Chatbot::MicrosoftTeamsChannelConfiguration.IamRoleArn`
- `AWS::Chatbot::SlackChannelConfiguration.IamRoleArn`
- `AWS::CloudFormation::HookVersion.ExecutionRoleArn`
- `AWS::CloudFormation::ResourceVersion.ExecutionRoleArn`
- `AWS::CloudFormation::TypeActivation.ExecutionRoleArn`
- `AWS::CloudTrail::EventDataStore.FederationRoleArn`
- `AWS::CloudTrail::Trail.CloudWatchLogsRoleArn`
- `AWS::CloudWatch::LogAlarm.ActionLogLineRoleArn`
- `AWS::CodeDeploy::DeploymentGroup.ServiceRoleArn`
- `AWS::Comprehend::DocumentClassifier.DataAccessRoleArn`
- `AWS::Comprehend::Flywheel.DataAccessRoleArn`
- `AWS::DLM::LifecyclePolicy.ExecutionRoleArn`
- `AWS::DMS::DataMigration.ServiceAccessRoleArn`
- `AWS::DataSync::LocationEFS.FileSystemAccessRoleArn`
- `AWS::DataZone::Environment.EnvironmentRoleArn`
- `AWS::DataZone::EnvironmentBlueprintConfiguration.ManageAccessRoleArn`
- `AWS::DataZone::EnvironmentBlueprintConfiguration.ProvisioningRoleArn`
- `AWS::DirectConnect::DirectConnectGatewayAssociation.AcceptDirectConnectGatewayAssociationProposalRoleArn`
- `AWS::DirectConnect::PrivateVirtualInterface.AllocatePrivateVirtualInterfaceRoleArn`
- `AWS::DirectConnect::PublicVirtualInterface.AllocatePublicVirtualInterfaceRoleArn`
- `AWS::DirectConnect::TransitVirtualInterface.AllocateTransitVirtualInterfaceRoleArn`
- `AWS::EC2::VPCPeeringConnection.PeerRoleArn`
- `AWS::ECR::PullThroughCacheRule.CustomRoleArn`
- `AWS::ECR::RepositoryCreationTemplate.CustomRoleArn`
- `AWS::ECS::DaemonTaskDefinition.ExecutionRoleArn`
- `AWS::ECS::DaemonTaskDefinition.TaskRoleArn`
- `AWS::ECS::ExpressGatewayService.ExecutionRoleArn`
- `AWS::ECS::ExpressGatewayService.InfrastructureRoleArn`
- `AWS::ECS::ExpressGatewayService.TaskRoleArn`
- `AWS::ECS::TaskDefinition.ExecutionRoleArn`
- `AWS::ECS::TaskDefinition.TaskRoleArn`
- `AWS::EKS::Addon.ServiceAccountRoleArn`
- `AWS::EKS::FargateProfile.PodExecutionRoleArn`
- `AWS::EKS::PodIdentityAssociation.TargetRoleArn`
- `AWS::EMRContainers::Endpoint.ExecutionRoleArn`
- `AWS::ElementalInference::Feed.AccessRoleArn`
- `AWS::GameLift::ContainerFleet.FleetRoleArn`
- `AWS::IoT::EncryptionConfiguration.KmsAccessRoleArn`
- `AWS::IoT::ProvisioningTemplate.ProvisioningRoleArn`
- `AWS::KafkaConnect::Connector.ServiceExecutionRoleArn`
- `AWS::Lambda::MicrovmImage.BuildRoleArn`
- `AWS::Logs::ScheduledQuery.ExecutionRoleArn`
- `AWS::MSK::Replicator.ServiceExecutionRoleArn`
- `AWS::MWAA::Environment.ExecutionRoleArn`
- `AWS::MediaLive::Cluster.InstanceRoleArn`
- `AWS::MediaPackage::Asset.SourceRoleArn`
- `AWS::OSIS::Pipeline.PipelineRoleArn`
- `AWS::Panorama::ApplicationInstance.RuntimeRoleArn`
- `AWS::Proton::EnvironmentAccountConnection.CodebuildRoleArn`
- `AWS::Proton::EnvironmentAccountConnection.ComponentRoleArn`
- `AWS::RDS::DBCluster.MonitoringRoleArn`
- `AWS::RDS::DBInstance.MonitoringRoleArn`
- `AWS::RedshiftServerless::Namespace.DefaultIamRoleArn`
- `AWS::S3::AccessGrantsLocation.IamRoleArn`
- `AWS::SNS::Subscription.SubscriptionRoleArn`
- `AWS::SSM::MaintenanceWindowTask.ServiceRoleArn`
- `AWS::SageMaker::EndpointConfig.ExecutionRoleArn`
- `AWS::SageMaker::Image.ImageRoleArn`
- `AWS::SageMaker::Model.ExecutionRoleArn`
- `AWS::SageMaker::PartnerApp.ExecutionRoleArn`
- `AWS::SecretsManager::RotationSchedule.ExternalSecretRotationRoleArn`
- `AWS::SecurityLake::DataLake.MetaStoreManagerRoleArn`
- `AWS::SupportApp::SlackChannelConfiguration.ChannelRoleArn`
- `AWS::Synthetics::Canary.ExecutionRoleArn`
- `AWS::Timestream::ScheduledQuery.ScheduledQueryExecutionRoleArn`
- `AWS::Transcribe::VocabularyFilter.DataAccessRoleArn`

### 2. AWS::EC2::SecurityGroup (14), at `Id`

looks right: `VpcSecurityGroupIds` and EMR Studio's engine and workspace groups are EC2 security group ids

- `AWS::DocDBElastic::Cluster.VpcSecurityGroupIds`
- `AWS::EC2::SecurityGroupEgress.DestinationSecurityGroupId`
- `AWS::EC2::SecurityGroupIngress.SourceSecurityGroupId`
- `AWS::EMR::Studio.EngineSecurityGroupId`
- `AWS::EMR::Studio.WorkspaceSecurityGroupId`
- `AWS::ElastiCache::CacheCluster.VpcSecurityGroupIds`
- `AWS::Neptune::DBCluster.VpcSecurityGroupIds`
- `AWS::RDS::DBCluster.VpcSecurityGroupIds`
- `AWS::RDS::DBProxy.VpcSecurityGroupIds`
- `AWS::RDS::DBProxyEndpoint.VpcSecurityGroupIds`
- `AWS::Redshift::Cluster.VpcSecurityGroupIds`
- `AWS::Redshift::EndpointAccess.VpcSecurityGroupIds`
- `AWS::Timestream::InfluxDBCluster.VpcSecurityGroupIds`
- `AWS::Timestream::InfluxDBInstance.VpcSecurityGroupIds`

### 3. AWS::EC2::IPAMPool (8), at `IpamPoolId`

looks right

- `AWS::EC2::IPAMPool.SourceIpamPoolId`
- `AWS::EC2::Subnet.Ipv4IpamPoolId`
- `AWS::EC2::Subnet.Ipv6IpamPoolId`
- `AWS::EC2::SubnetCidrBlock.Ipv6IpamPoolId`
- `AWS::EC2::VPC.Ipv4IpamPoolId`
- `AWS::EC2::VPCCidrBlock.Ipv4IpamPoolId`
- `AWS::EC2::VPCCidrBlock.Ipv6IpamPoolId`
- `AWS::ElasticLoadBalancingV2::LoadBalancer.Ipv4IpamPoolId`

### 4. AWS::EC2::Subnet (7), at `SubnetId`

looks right

- `AWS::EC2::TransitGatewayVpcAttachment.AddSubnetIds`
- `AWS::EC2::TransitGatewayVpcAttachment.RemoveSubnetIds`
- `AWS::EVS::Environment.ServiceAccessSubnetId`
- `AWS::RDS::DBProxy.VpcSubnetIds`
- `AWS::RDS::DBProxyEndpoint.VpcSubnetIds`
- `AWS::Timestream::InfluxDBCluster.VpcSubnetIds`
- `AWS::Timestream::InfluxDBInstance.VpcSubnetIds`

### 5. AWS::DMS::Endpoint (4), at `EndpointArn`

looks right: source and target endpoints of DMS tasks

- `AWS::DMS::ReplicationConfig.SourceEndpointArn`
- `AWS::DMS::ReplicationConfig.TargetEndpointArn`
- `AWS::DMS::ReplicationTask.SourceEndpointArn`
- `AWS::DMS::ReplicationTask.TargetEndpointArn`

### 6. AWS::EC2::RouteTable (4), at `RouteTableId`

MIXED: TransitGateway's AssociationDefaultRouteTableId and PropagationDefaultRouteTableId, and TransitGatewayPolicyTableEntry.TargetRouteTableId, hold transit gateway route table ids, not VPC route tables. Approve this target only together with rejects for those three

- `AWS::EC2::TransitGateway.AssociationDefaultRouteTableId`
- `AWS::EC2::TransitGateway.PropagationDefaultRouteTableId`
- `AWS::EC2::TransitGatewayPolicyTableEntry.TargetRouteTableId`
- `AWS::ODB::OdbPeeringConnection.PeerNetworkRouteTableIds`

### 7. AWS::EC2::TransitGatewayAttachment (4), at `Id`

looks right

- `AWS::EC2::TransitGatewayConnect.TransportTransitGatewayAttachmentId`
- `AWS::EC2::TransitGatewayMeteringPolicyEntry.DestinationTransitGatewayAttachmentId`
- `AWS::EC2::TransitGatewayMeteringPolicyEntry.SourceTransitGatewayAttachmentId`
- `AWS::EC2::VPNConnection.TransportTransitGatewayAttachmentId`

### 8. AWS::EC2::PrefixList (3), at `PrefixListId`

looks right (the id may also be an AWS-managed prefix list, which is still the same kind of id)

- `AWS::EC2::Route.DestinationPrefixListId`
- `AWS::EC2::SecurityGroupEgress.DestinationPrefixListId`
- `AWS::EC2::SecurityGroupIngress.SourcePrefixListId`

### 9. AWS::EC2::VPC (3), at `VpcId`

looks right (a peer VPC can live in another account, and ClassicLinkVPCId is legacy)

- `AWS::AutoScaling::LaunchConfiguration.ClassicLinkVPCId`
- `AWS::EC2::VPCPeeringConnection.PeerVpcId`
- `AWS::GameLift::Fleet.PeerVpcId`

### 10. AWS::GroundStation::Config (3), at `Arn`, `Id`

MIXED: the two MissionProfile config arns are right; Route53Resolver's ResolverQueryLogConfigId -> Config.Id is FABRICATED

- `AWS::GroundStation::MissionProfile.TelemetrySinkConfigArn` at `Arn`
- `AWS::GroundStation::MissionProfile.TrackingConfigArn` at `Arn`
- `AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation.ResolverQueryLogConfigId` at `Id`

### 11. AWS::ApiGateway::Resource (2), at `ResourceId`

FABRICATED: RDS resource ids (dbi-/cluster- ids)

- `AWS::RDS::DBCluster.SourceDbClusterResourceId`
- `AWS::RDS::DBInstance.SourceDbiResourceId`

### 12. AWS::Connect::Notification (2), at `Arn`

FABRICATED: SNS topic arns and VPC endpoint connection notifications, nothing to do with Connect

- `AWS::ApplicationInsights::Application.SNSNotificationArn`
- `AWS::EC2::VPCEndpointConnectionNotification.ConnectionNotificationArn`

### 13. AWS::KMS::Key (2), at `Arn`, `KeyId`

looks right: ReplicaKey.PrimaryKeyArn and Alias.TargetKeyId

- `AWS::KMS::Alias.TargetKeyId` at `KeyId`
- `AWS::KMS::ReplicaKey.PrimaryKeyArn` at `Arn`

### 14. AWS::Lambda::Version (2), at `FunctionArn`

FABRICATED: CloudFormation extension type version arns

- `AWS::CloudFormation::HookDefaultVersion.TypeVersionArn`
- `AWS::CloudFormation::ResourceDefaultVersion.TypeVersionArn`

### 15. AWS::Logs::LogGroup (2), at `Arn`

looks right

- `AWS::CloudTrail::Trail.CloudWatchLogsLogGroupArn`
- `AWS::DataSync::Task.CloudWatchLogGroupArn`

### 16. AWS::Notifications::NotificationConfiguration (2), at `Arn`

probably FABRICATED: a managed notification configuration is AWS-owned, not this resource type

- `AWS::Notifications::ManagedNotificationAccountContactAssociation.ManagedNotificationConfigurationArn`
- `AWS::Notifications::ManagedNotificationAdditionalChannelAssociation.ManagedNotificationConfigurationArn`

### 17. AWS::ResilienceHubV2::System (2), at `SystemId`

FABRICATED: DBSystemId is an Oracle CDB name

- `AWS::RDS::DBCluster.DBSystemId`
- `AWS::RDS::DBInstance.DBSystemId`

### 18. AWS::Transfer::Profile (2), at `ProfileId`

looks right

- `AWS::Transfer::Agreement.LocalProfileId`
- `AWS::Transfer::Agreement.PartnerProfileId`

### 19. AWS::AmazonMQ::Broker (1), at `Arn`

looks right (self reference)

- `AWS::AmazonMQ::Broker.DataReplicationPrimaryBrokerArn`

### 20. AWS::AmplifyUIBuilder::Form (1), at `Id`

FABRICATED: Signer's PlatformId is a signing platform

- `AWS::Signer::SigningProfile.PlatformId`

### 21. AWS::BedrockAgentCore::Runtime (1), at `AgentRuntimeId`

looks right

- `AWS::BedrockAgentCore::RuntimeEndpoint.AgentRuntimeId`

### 22. AWS::Connect::ContactFlow (1), at `ContactFlowArn`

looks right

- `AWS::Connect::TaskTemplate.SelfAssignContactFlowArn`

### 23. AWS::Connect::PhoneNumber (1), at `PhoneNumberArn`

looks right (self reference)

- `AWS::Connect::PhoneNumber.SourcePhoneNumberArn`

### 24. AWS::Connect::Queue (1), at `QueueArn`

looks right

- `AWS::Connect::RoutingProfile.DefaultOutboundQueueArn`

### 25. AWS::Deadline::Limit (1), at `LimitId`

FABRICATED: a BedrockAgentCore gateway rate limit is not a Deadline limit

- `AWS::BedrockAgentCore::GatewayRateLimit.RateLimitId`

### 26. AWS::Deadline::StorageProfile (1), at `StorageProfileId`

looks right

- `AWS::Deadline::Queue.AllowedStorageProfileIds`

### 27. AWS::DevOpsAgent::Association (1), at `AssociationId`

looks right (self reference)

- `AWS::DevOpsAgent::Association.LinkedAssociationIds`

### 28. AWS::DirectConnect::Lag (1), at `LagId`

FABRICATED: an Outpost LAG, not a Direct Connect LAG

- `AWS::EC2::LocalGatewayVirtualInterface.OutpostLagId`

### 29. AWS::EC2::TransitGateway (1), at `Id`

looks right

- `AWS::EC2::TransitGatewayPeeringAttachment.PeerTransitGatewayId`

### 30. AWS::EC2::Volume (1), at `VolumeId`

looks right (self reference)

- `AWS::EC2::Volume.SourceVolumeId`

### 31. AWS::EKS::CertificateAuthority (1), at `Id`

unsure: check what Cluster.ActiveCertificateAuthorityId holds

- `AWS::EKS::Cluster.ActiveCertificateAuthorityId`

### 32. AWS::GameLift::Location (1), at `LocationArn`

FABRICATED: a DataSync location, not a GameLift location

- `AWS::DataSync::Task.DestinationLocationArn`

### 33. AWS::Glue::Catalog (1), at `CatalogId`

unsure: a Lake Formation catalog id is usually an account id

- `AWS::LakeFormation::DataCellsFilter.TableCatalogId`

### 34. AWS::Inspector::ResourceGroup (1), at `Arn`

FABRICATED: EC2 host resource groups belong to License Manager

- `AWS::EC2::Instance.HostResourceGroupArn`

### 35. AWS::IoT::Certificate (1), at `Arn`

looks right

- `AWS::IoT::DomainConfiguration.ValidationCertificateArn`

### 36. AWS::IoTTwinMaker::Entity (1), at `EntityId`

looks right (self reference)

- `AWS::IoTTwinMaker::Entity.ParentEntityId`

### 37. AWS::Lambda::Function (1), at `Arn`

looks right

- `AWS::Lambda::Url.TargetFunctionArn`

### 38. AWS::MPA::ApprovalTeam (1), at `Arn`

looks right

- `AWS::Backup::LogicallyAirGappedBackupVault.MpaApprovalTeamArn`

### 39. AWS::MSK::Cluster (1), at `Arn`

looks right

- `AWS::MSK::VpcConnection.TargetClusterArn`

### 40. AWS::Panorama::Package (1), at `Arn`

FABRICATED: Inspector rules packages

- `AWS::Inspector::AssessmentTemplate.RulesPackageArns`

### 41. AWS::QuickSight::Folder (1), at `Arn`

looks right (self reference)

- `AWS::QuickSight::Folder.ParentFolderArn`

### 42. AWS::QuickSight::Theme (1), at `ThemeId`

looks right (self reference)

- `AWS::QuickSight::Theme.BaseThemeId`

### 43. AWS::ServiceCatalog::Portfolio (1), at `Id`

looks right

- `AWS::ServiceCatalog::PortfolioProductAssociation.SourcePortfolioId`

### 44. AWS::WorkSpacesWeb::IdentityProvider (1), at `IdentityProviderArn`

FABRICATED: QBusiness wants an IAM identity provider

- `AWS::QBusiness::Application.IamIdentityProviderArn`

## Fabrications

### In tier 2 (pending, so not in the catalog)

- `AWS::EC2::RouteTable`: `AWS::EC2::TransitGateway.AssociationDefaultRouteTableId`, `AWS::EC2::TransitGateway.PropagationDefaultRouteTableId`, `AWS::EC2::TransitGatewayPolicyTableEntry.TargetRouteTableId`, `AWS::ODB::OdbPeeringConnection.PeerNetworkRouteTableIds`. MIXED: TransitGateway's AssociationDefaultRouteTableId and PropagationDefaultRouteTableId, and TransitGatewayPolicyTableEntry.TargetRouteTableId, hold transit gateway route table ids, not VPC route tables. Approve this target only together with rejects for those three
- `AWS::GroundStation::Config`: `AWS::GroundStation::MissionProfile.TelemetrySinkConfigArn`, `AWS::GroundStation::MissionProfile.TrackingConfigArn`, `AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation.ResolverQueryLogConfigId`. MIXED: the two MissionProfile config arns are right; Route53Resolver's ResolverQueryLogConfigId -> Config.Id is FABRICATED
- `AWS::ApiGateway::Resource`: `AWS::RDS::DBCluster.SourceDbClusterResourceId`, `AWS::RDS::DBInstance.SourceDbiResourceId`. FABRICATED: RDS resource ids (dbi-/cluster- ids)
- `AWS::Connect::Notification`: `AWS::ApplicationInsights::Application.SNSNotificationArn`, `AWS::EC2::VPCEndpointConnectionNotification.ConnectionNotificationArn`. FABRICATED: SNS topic arns and VPC endpoint connection notifications, nothing to do with Connect
- `AWS::Lambda::Version`: `AWS::CloudFormation::HookDefaultVersion.TypeVersionArn`, `AWS::CloudFormation::ResourceDefaultVersion.TypeVersionArn`. FABRICATED: CloudFormation extension type version arns
- `AWS::Notifications::NotificationConfiguration`: `AWS::Notifications::ManagedNotificationAccountContactAssociation.ManagedNotificationConfigurationArn`, `AWS::Notifications::ManagedNotificationAdditionalChannelAssociation.ManagedNotificationConfigurationArn`. probably FABRICATED: a managed notification configuration is AWS-owned, not this resource type
- `AWS::ResilienceHubV2::System`: `AWS::RDS::DBCluster.DBSystemId`, `AWS::RDS::DBInstance.DBSystemId`. FABRICATED: DBSystemId is an Oracle CDB name
- `AWS::AmplifyUIBuilder::Form`: `AWS::Signer::SigningProfile.PlatformId`. FABRICATED: Signer's PlatformId is a signing platform
- `AWS::Deadline::Limit`: `AWS::BedrockAgentCore::GatewayRateLimit.RateLimitId`. FABRICATED: a BedrockAgentCore gateway rate limit is not a Deadline limit
- `AWS::DirectConnect::Lag`: `AWS::EC2::LocalGatewayVirtualInterface.OutpostLagId`. FABRICATED: an Outpost LAG, not a Direct Connect LAG
- `AWS::GameLift::Location`: `AWS::DataSync::Task.DestinationLocationArn`. FABRICATED: a DataSync location, not a GameLift location
- `AWS::Inspector::ResourceGroup`: `AWS::EC2::Instance.HostResourceGroupArn`. FABRICATED: EC2 host resource groups belong to License Manager
- `AWS::Panorama::Package`: `AWS::Inspector::AssessmentTemplate.RulesPackageArns`. FABRICATED: Inspector rules packages
- `AWS::WorkSpacesWeb::IdentityProvider`: `AWS::QBusiness::Application.IamIdentityProviderArn`. FABRICATED: QBusiness wants an IAM identity provider

### In tier 1 (ALREADY ACCEPTED, and in the catalog)

An exact name match is still wrong when the name is generic. These are only the ones found by scanning tier-1 edges that cross services. Same-service tier-1 edges were not reviewed. None are rejected yet; each is a candidate for `references.reject`.

- `AWS::ApiGateway::Resource` (9): FABRICATED: a generic `ResourceId` (an autoscaling target, a flow log's VPC or subnet, a Route 53 resolver's VPC, and so on), not an API Gateway resource. Sources: `AWS::ApplicationAutoScaling::ScalableTarget.ResourceId`, `AWS::ApplicationAutoScaling::ScalingPolicy.ResourceId`, `AWS::EC2::FlowLog.ResourceId`, `AWS::MediaPackage::Asset.ResourceId`, `AWS::Route53Profiles::ProfileAssociation.ResourceId`, `AWS::Route53Resolver::ResolverConfig.ResourceId`, `AWS::Route53Resolver::ResolverDNSSECConfig.ResourceId`, `AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation.ResourceId`, `AWS::ServiceCatalog::TagOptionAssociation.ResourceId`
- `AWS::Connect::Notification` (2): FABRICATED: Service Catalog's `NotificationArns` are SNS topic arns. Sources: `AWS::ServiceCatalog::CloudFormationProvisionedProduct.NotificationArns`, `AWS::ServiceCatalog::LaunchNotificationConstraint.NotificationArns`
- `AWS::GuardDuty::Member` (2): FABRICATED: Detective's `MemberId` is an account id; IdentityStore's is a user or group. Sources: `AWS::Detective::MemberInvitation.MemberId`, `AWS::IdentityStore::GroupMembership.MemberId`
- `AWS::MediaTailor::SourceLocation` (1): FABRICATED: a DataSync source location. Sources: `AWS::DataSync::Task.SourceLocationArn`
- `AWS::Glue::Catalog` (1): unsure: Lake Formation's `CatalogId` is usually an account id. Sources: `AWS::LakeFormation::Tag.CatalogId`

## Unresolved: ambiguous (285 properties, 51 candidate sets)

More than one type still matched after the same-service tiebreak, so no edge was written. Nothing is guessed. A future overlay entry could pin these by hand; the biggest sets are the most valuable.

- **104** among `AWS::KMS::Key`, `AWS::PaymentCryptography::Key`: `AWS::APS::Workspace.KmsKeyArn`, `AWS::ApiGateway::UsagePlanKey.KeyId`, `AWS::Backup::BackupVault.EncryptionKeyArn`, `AWS::Backup::LogicallyAirGappedBackupVault.EncryptionKeyArn`, `AWS::BackupGateway::Hypervisor.KmsKeyArn`, `AWS::Bedrock::Agent.CustomerEncryptionKeyArn`, `AWS::Bedrock::AutomatedReasoningPolicy.KmsKeyId`, `AWS::Bedrock::Blueprint.KmsKeyId`, `AWS::Bedrock::DataAutomationProject.KmsKeyId`, `AWS::Bedrock::Flow.CustomerEncryptionKeyArn`, `AWS::Bedrock::Guardrail.KmsKeyArn`, `AWS::Bedrock::Prompt.CustomerEncryptionKeyArn`, `AWS::Bedrock::Session.EncryptionKeyArn`, `AWS::BedrockAgentCore::ConfigurationBundle.KmsKeyArn`, `AWS::BedrockAgentCore::Dataset.KmsKeyArn`, `AWS::BedrockAgentCore::Evaluator.KmsKeyArn`, `AWS::BedrockAgentCore::Gateway.KmsKeyArn`, `AWS::BedrockAgentCore::Memory.EncryptionKeyArn`, `AWS::BedrockAgentCore::PolicyEngine.EncryptionKeyArn`, `AWS::CleanRooms::IdMappingTable.KmsKeyArn`, `AWS::CleanRooms::IntermediateTable.KmsKeyArn`, `AWS::CleanRoomsML::ConfiguredModelAlgorithm.KmsKeyArn`, `AWS::CloudTrail::EventDataStore.KmsKeyId`, `AWS::CloudTrail::Trail.KMSKeyId`, `AWS::CodeCommit::Repository.KmsKeyId`, `AWS::CodeStarConnections::RepositoryLink.EncryptionKeyArn`, `AWS::Comprehend::DocumentClassifier.ModelKmsKeyId`, `AWS::Comprehend::DocumentClassifier.VolumeKmsKeyId`, `AWS::Config::DeliveryChannel.S3KmsKeyArn`, `AWS::DMS::Endpoint.KmsKeyId`, `AWS::DMS::InstanceProfile.KmsKeyArn`, `AWS::DataBrew::Job.EncryptionKeyArn`, `AWS::Deadline::Farm.KmsKeyArn`, `AWS::DevOpsAgent::AgentSpace.KmsKeyArn`, `AWS::DevOpsAgent::Service.KmsKeyArn`, `AWS::DocDBElastic::Cluster.KmsKeyId`, `AWS::EC2::Volume.KmsKeyId`, `AWS::EFS::FileSystem.KmsKeyId`, `AWS::EMR::Step.EncryptionKeyArn`, `AWS::EMR::Studio.EncryptionKeyArn`, `AWS::EVS::Environment.KmsKeyId`, `AWS::ElastiCache::ReplicationGroup.KmsKeyId`, `AWS::ElastiCache::ServerlessCache.KmsKeyId`, `AWS::ElastiCache::ServerlessCacheSnapshot.KmsKeyId`, `AWS::FinSpace::Environment.KmsKeyId`, `AWS::Glue::Integration.KmsKeyId`, `AWS::HealthImaging::Datastore.KmsKeyArn`, `AWS::HealthLake::DataTransformationProfile.KmsKeyId`, `AWS::ImageBuilder::Component.KmsKeyId`, `AWS::ImageBuilder::ContainerRecipe.KmsKeyId`, `AWS::ImageBuilder::Workflow.KmsKeyId`, `AWS::IoT::EncryptionConfiguration.KmsKeyArn`, `AWS::IoTSiteWise::Workspace.KmsKeyId`, `AWS::KinesisVideo::Stream.KmsKeyId`, `AWS::Lambda::CapacityProvider.KmsKeyArn`, `AWS::Lambda::EventSourceMapping.KmsKeyArn`, `AWS::Lambda::Function.KmsKeyArn`, `AWS::Location::GeofenceCollection.KmsKeyId`, `AWS::Location::Tracker.KmsKeyId`, `AWS::Logs::LogAnomalyDetector.KmsKeyId`, `AWS::Logs::LogGroup.KmsKeyId`, `AWS::LookoutEquipment::InferenceScheduler.ServerSideKmsKeyId`, `AWS::M2::Application.KmsKeyId`, `AWS::M2::Environment.KmsKeyId`, `AWS::MemoryDB::Cluster.KmsKeyId`, `AWS::Neptune::DBCluster.KmsKeyId`, `AWS::OpenSearchService::Application.KmsKeyArn`, `AWS::Personalize::DatasetGroup.KmsKeyArn`, `AWS::RDS::CustomDBEngineVersion.KMSKeyId`, `AWS::RDS::DBCluster.KmsKeyId`, `AWS::RDS::DBCluster.PerformanceInsightsKmsKeyId`, `AWS::RDS::DBInstance.AutomaticBackupReplicationKmsKeyId`, `AWS::RDS::DBInstance.KmsKeyId`, `AWS::RDS::DBInstance.PerformanceInsightsKMSKeyId`, `AWS::RDS::Integration.KMSKeyId`, `AWS::Redshift::Cluster.KmsKeyId`, `AWS::Redshift::Cluster.MasterPasswordSecretKmsKeyId`, `AWS::Redshift::Integration.KMSKeyId`, `AWS::RedshiftServerless::Namespace.AdminPasswordSecretKmsKeyId`, `AWS::RedshiftServerless::Namespace.KmsKeyId`, `AWS::Rekognition::StreamProcessor.KmsKeyId`, `AWS::ResilienceHubV2::Policy.KmsKeyId`, `AWS::ResilienceHubV2::Service.KmsKeyId`, `AWS::ResilienceHubV2::System.KmsKeyId`, `AWS::S3Files::FileSystem.KmsKeyId`, `AWS::SES::MailManagerArchive.KmsKeyArn`, `AWS::SNS::Topic.KmsMasterKeyId`, `AWS::SQS::Queue.KmsMasterKeyId`, `AWS::SSM::ResourceDataSync.KMSKeyArn`, `AWS::SageMaker::Domain.KmsKeyId`, `AWS::SageMaker::EndpointConfig.KmsKeyId`, `AWS::SageMaker::NotebookInstance.KmsKeyId`, `AWS::SageMaker::PartnerApp.KmsKeyId`, `AWS::Scheduler::Schedule.KmsKeyArn`, `AWS::SecretsManager::Secret.KmsKeyId`, `AWS::SecurityAgent::AgentSpace.KmsKeyId`, `AWS::SecurityAgent::Application.DefaultKmsKeyId`, `AWS::SecurityAgent::SecurityRequirementPack.KmsKeyId`, `AWS::SecurityHub::ConnectorV2.KmsKeyArn`, `AWS::Synthetics::Canary.KmsKeyArn`, `AWS::Timestream::Database.KmsKeyId`, `AWS::Timestream::ScheduledQuery.KmsKeyId`, `AWS::WorkSpacesThinClient::Environment.KmsKeyArn`, `AWS::WorkspacesInstances::Volume.KmsKeyId`
- **39** among `AWS::ApiGateway::Account`, `AWS::CertificateManager::Account`, `AWS::Organizations::Account`: `AWS::BillingConductor::BillingGroup.PrimaryAccountId`, `AWS::BillingConductor::CustomLineItem.AccountId`, `AWS::Config::AggregationAuthorization.AuthorizedAccountId`, `AWS::Connect::DataLakeAssociation.TargetAccountId`, `AWS::DataZone::EnvironmentProfile.AwsAccountId`, `AWS::Detective::OrganizationAdmin.AccountId`, `AWS::EC2::TransitGatewayPeeringAttachment.PeerAccountId`, `AWS::FIS::TargetAccountConfiguration.AccountId`, `AWS::GameLift::Fleet.PeerVpcAwsAccountId`, `AWS::InternetMonitor::Monitor.LinkedAccountId`, `AWS::IoT::AccountAuditConfiguration.AccountId`, `AWS::IoT::Logging.AccountId`, `AWS::Logs::LogAnomalyDetector.AccountId`, `AWS::Proton::EnvironmentAccountConnection.EnvironmentAccountId`, `AWS::Proton::EnvironmentAccountConnection.ManagementAccountId`, `AWS::QuickSight::ActionConnector.AwsAccountId`, `AWS::QuickSight::Agent.AwsAccountId`, `AWS::QuickSight::Analysis.AwsAccountId`, `AWS::QuickSight::CustomPermissions.AwsAccountId`, `AWS::QuickSight::DLPSetting.AwsAccountId`, `AWS::QuickSight::Dashboard.AwsAccountId`, `AWS::QuickSight::DataSet.AwsAccountId`, `AWS::QuickSight::DataSource.AwsAccountId`, `AWS::QuickSight::Flow.AwsAccountId`, `AWS::QuickSight::Folder.AwsAccountId`, `AWS::QuickSight::KnowledgeBase.AwsAccountId`, `AWS::QuickSight::LimitsProfile.AccountId`, `AWS::QuickSight::RefreshSchedule.AwsAccountId`, `AWS::QuickSight::Space.AwsAccountId`, `AWS::QuickSight::Template.AwsAccountId`, `AWS::QuickSight::Theme.AwsAccountId`, `AWS::QuickSight::Topic.AwsAccountId`, `AWS::QuickSight::TopicV2.AwsAccountId`, `AWS::QuickSight::VPCConnection.AwsAccountId`, `AWS::S3::AccessPoint.BucketAccountId`, `AWS::S3Express::AccessPoint.BucketAccountId`, `AWS::SecurityHub::DelegatedAdmin.AdminAccountId`, `AWS::ServiceCatalog::PortfolioShare.AccountId`, `AWS::WellArchitected::Workload.AccountIds`
- **13** among `AWS::MSK::Topic`, `AWS::QuickSight::Topic`, `AWS::SNS::Topic`: `AWS::ApplicationInsights::Application.OpsItemSNSTopicArn`, `AWS::Chatbot::MicrosoftTeamsChannelConfiguration.SnsTopicArns`, `AWS::Chatbot::SlackChannelConfiguration.SnsTopicArns`, `AWS::DMS::EventSubscription.SnsTopicArn`, `AWS::DocDB::EventSubscription.SnsTopicArn`, `AWS::ElastiCache::CacheCluster.NotificationTopicArn`, `AWS::ElastiCache::ReplicationGroup.NotificationTopicArn`, `AWS::FMS::NotificationChannel.SnsTopicArn`, `AWS::ImageBuilder::InfrastructureConfiguration.SnsTopicArn`, `AWS::MemoryDB::Cluster.SnsTopicArn`, `AWS::Neptune::EventSubscription.SnsTopicArn`, `AWS::RDS::EventSubscription.SnsTopicArn`, `AWS::Redshift::EventSubscription.SnsTopicArn`
- **12** among `AWS::BedrockAgentCore::Gateway`, `AWS::IoTSiteWise::Gateway`, `AWS::MediaConnect::Gateway`: `AWS::DirectConnect::DirectConnectGatewayAssociation.AssociatedGatewayId`, `AWS::DirectConnect::PrivateVirtualInterface.VirtualGatewayId`, `AWS::EC2::GatewayRouteTableAssociation.GatewayId`, `AWS::EC2::LocalGatewayRouteTable.LocalGatewayId`, `AWS::EC2::LocalGatewayVirtualInterfaceGroup.LocalGatewayId`, `AWS::EC2::Route.GatewayId`, `AWS::EC2::Route.LocalGatewayId`, `AWS::RTBFabric::InboundExternalLink.GatewayId`, `AWS::RTBFabric::Link.GatewayId`, `AWS::RTBFabric::Link.PeerGatewayId`, `AWS::RTBFabric::LinkRoutingRule.GatewayId`, `AWS::RTBFabric::OutboundExternalLink.GatewayId`
- **12** among `AWS::Connect::Instance`, `AWS::EC2::Instance`, `AWS::Lightsail::Instance`, `AWS::SSO::Instance`, `AWS::ServiceDiscovery::Instance`: `AWS::AutoScaling::AutoScalingGroup.InstanceId`, `AWS::AutoScaling::LaunchConfiguration.InstanceId`, `AWS::ConnectCampaigns::Campaign.ConnectInstanceArn`, `AWS::ConnectCampaignsV2::Campaign.ConnectInstanceId`, `AWS::DMS::ReplicationTask.ReplicationInstanceArn`, `AWS::Deadline::Monitor.IdentityCenterInstanceArn`, `AWS::EMR::Studio.IdcInstanceArn`, `AWS::Glue::IdentityCenterConfiguration.InstanceArn`, `AWS::QBusiness::Application.IdentityCenterInstanceArn`, `AWS::SCN::Dataset.InstanceId`, `AWS::SCN::Namespace.InstanceId`, `AWS::SSM::Association.InstanceId`
- **10** among `AWS::ACMPCA::Certificate`, `AWS::CertificateManager::Certificate`, `AWS::DMS::Certificate`, `AWS::IoT::Certificate`, `AWS::Lightsail::Certificate`, `AWS::Transfer::Certificate`: `AWS::ApiGateway::DomainName.CertificateArn`, `AWS::ApiGateway::DomainName.OwnershipVerificationCertificateArn`, `AWS::ApiGateway::DomainName.RegionalCertificateArn`, `AWS::ApiGateway::DomainNameV2.CertificateArn`, `AWS::AppSync::DomainName.CertificateArn`, `AWS::EC2::CustomerGateway.CertificateArn`, `AWS::EC2::EnclaveCertificateIamRoleAssociation.CertificateArn`, `AWS::EC2::VerifiedAccessEndpoint.DomainCertificateArn`, `AWS::RTBFabric::ResponderGateway.AcmCertificateArn`, `AWS::VpcLattice::Service.CertificateArn`
- **9** among `AWS::IAM::Group`, `AWS::IdentityStore::Group`, `AWS::ResourceGroups::Group`, `AWS::Synthetics::Group`, `AWS::XRay::Group`: `AWS::Connect::SecurityProfile.AllowedAccessControlHierarchyGroupId`, `AWS::Connect::User.HierarchyGroupArn`, `AWS::Connect::UserHierarchyGroup.ParentGroupArn`, `AWS::EC2::PlacementGroup.ParentGroupId`, `AWS::EC2::SecurityGroupEgress.GroupId`, `AWS::EC2::SecurityGroupIngress.GroupId`, `AWS::EC2::SecurityGroupVpcAssociation.GroupId`, `AWS::RoboMaker::Robot.GreengrassGroupId`, `AWS::VpcLattice::ResourceConfiguration.ResourceConfigurationGroupId`
- **6** among `AWS::ImageBuilder::Image`, `AWS::SageMaker::Image`: `AWS::AppStream::ImageBuilder.ImageArn`, `AWS::AutoScaling::LaunchConfiguration.ImageId`, `AWS::CodeBuild::Fleet.ImageId`, `AWS::EC2::Instance.ImageId`, `AWS::Lambda::MicrovmImage.BaseImageArn`, `AWS::RDS::CustomDBEngineVersion.ImageId`
- **5** among `AWS::EFS::FileSystem`, `AWS::S3Files::FileSystem`: `AWS::DataSync::LocationEFS.EfsFilesystemArn`, `AWS::DataSync::LocationFSxLustre.FsxFilesystemArn`, `AWS::DataSync::LocationFSxOpenZFS.FsxFilesystemArn`, `AWS::DataSync::LocationFSxWindows.FsxFilesystemArn`, `AWS::FSx::DataRepositoryAssociation.FileSystemId`
- **4** among `AWS::BedrockAgentCore::Policy`, `AWS::FMS::Policy`, `AWS::IoT::Policy`, `AWS::Organizations::Policy`, `AWS::ResilienceHubV2::Policy`, `AWS::VerifiedPermissions::Policy`: `AWS::Bedrock::AutomatedReasoningPolicyVersion.PolicyArn`, `AWS::EMR::StudioSessionMapping.SessionPolicyArn`, `AWS::QuickSight::ApprovalPolicy.PolicyId`, `AWS::SES::MailManagerIngressPoint.TrafficPolicyId`
- **4** among `AWS::CloudTrail::Channel`, `AWS::IVS::Channel`, `AWS::IoTAnalytics::Channel`, `AWS::MSK::Channel`, `AWS::MediaPackage::Channel`, `AWS::MediaPackageV2::Channel`, `AWS::MediaTailor::Channel`: `AWS::Chatbot::MicrosoftTeamsChannelConfiguration.TeamsChannelId`, `AWS::Chatbot::SlackChannelConfiguration.SlackChannelId`, `AWS::Notifications::ManagedNotificationAdditionalChannelAssociation.ChannelArn`, `AWS::SupportApp::SlackChannelConfiguration.ChannelId`
- **4** among `AWS::IoT::Stream`, `AWS::Kinesis::Stream`, `AWS::KinesisVideo::Stream`, `AWS::QLDB::Stream`: `AWS::MediaConnect::FlowOutput.StreamId`, `AWS::MediaConnect::FlowSource.StreamId`, `AWS::MediaTailor::PrefetchSchedule.StreamId`, `AWS::WorkSpacesWeb::UserAccessLoggingSettings.KinesisStreamArn`
- **3** among `AWS::ACMPCA::CertificateAuthority`, `AWS::EKS::CertificateAuthority`: `AWS::CertificateManager::Certificate.CertificateAuthorityArn`, `AWS::PCAConnectorAD::Connector.CertificateAuthorityArn`, `AWS::PCAConnectorSCEP::Connector.CertificateAuthorityArn`
- **3** among `AWS::Amplify::App`, `AWS::ResilienceHub::App`, `AWS::SageMaker::App`: `AWS::AmplifyUIBuilder::Component.AppId`, `AWS::AmplifyUIBuilder::Form.AppId`, `AWS::AmplifyUIBuilder::Theme.AppId`
- **3** among `AWS::AppRunner::Service`, `AWS::DevOpsAgent::Service`, `AWS::ECS::Service`, `AWS::RefactorSpaces::Service`, `AWS::ResilienceHubV2::Service`, `AWS::ServiceDiscovery::Service`, `AWS::VpcLattice::Service`: `AWS::EC2::VPCEndpointConnectionNotification.ServiceId`, `AWS::EC2::VPCEndpointServicePermissions.ServiceId`, `AWS::Route53::KeySigningKey.KeyManagementServiceArn`
- **3** among `AWS::CodeConnections::Connection`, `AWS::CodeStarConnections::Connection`, `AWS::DataZone::Connection`, `AWS::DirectConnect::Connection`, `AWS::Events::Connection`, `AWS::Glue::Connection`, `AWS::Interconnect::Connection`: `AWS::ApiGatewayV2::Integration.ConnectionId`, `AWS::CloudFormation::Publisher.ConnectionArn`, `AWS::CodeGuruReviewer::RepositoryAssociation.ConnectionArn`
- **3** among `AWS::ElasticLoadBalancing::LoadBalancer`, `AWS::ElasticLoadBalancingV2::LoadBalancer`, `AWS::Lightsail::LoadBalancer`: `AWS::EC2::TrafficMirrorTarget.NetworkLoadBalancerArn`, `AWS::EC2::VPCEndpointService.GatewayLoadBalancerArns`, `AWS::EC2::VPCEndpointService.NetworkLoadBalancerArns`
- **3** among `AWS::Lightsail::Bucket`, `AWS::S3::Bucket`, `AWS::S3Outposts::Bucket`: `AWS::DRS::LaunchConfigurationTemplate.ExportBucketArn`, `AWS::DataSync::LocationS3.S3BucketArn`, `AWS::MWAA::Environment.SourceBucketArn`
- **2** among `AWS::AccountAccess::Application`, `AWS::AppConfig::Application`, `AWS::AppIntegrations::Application`, `AWS::AppStream::Application`, `AWS::ApplicationInsights::Application`, `AWS::CodeDeploy::Application`, `AWS::EMRServerless::Application`, `AWS::ElasticBeanstalk::Application`, `AWS::KinesisAnalyticsV2::Application`, `AWS::M2::Application`, `AWS::OpenSearchService::Application`, `AWS::QBusiness::Application`, `AWS::RefactorSpaces::Application`, `AWS::SSO::Application`, `AWS::SecurityAgent::Application`, `AWS::ServerlessRepo::Application`, `AWS::ServiceCatalogAppRegistry::Application`, `AWS::SystemsManagerSAP::Application`: `AWS::RedshiftServerless::Namespace.RedshiftIdcApplicationArn`, `AWS::S3::AccessGrant.ApplicationArn`
- **2** among `AWS::ApiGateway::Model`, `AWS::ApiGatewayV2::Model`, `AWS::SageMaker::Model`: `AWS::Comprehend::Flywheel.ActiveModelArn`, `AWS::Wisdom::AIPrompt.ModelId`
- **2** among `AWS::AppFlow::Flow`, `AWS::Bedrock::Flow`, `AWS::MediaConnect::Flow`, `AWS::QuickSight::Flow`: `AWS::ConnectCampaignsV2::Campaign.ConnectCampaignFlowArn`, `AWS::EMR::Step.JobFlowId`
- **2** among `AWS::AppSync::Type`, `AWS::Cassandra::Type`: `AWS::CloudFormation::HookTypeConfig.TypeArn`, `AWS::CloudFormation::TypeActivation.PublicTypeArn`
- **2** among `AWS::Bedrock::Blueprint`, `AWS::Glue::Blueprint`: `AWS::Lightsail::Database.RelationalDatabaseBlueprintId`, `AWS::Lightsail::Instance.BlueprintId`
- **2** among `AWS::Cases::Template`, `AWS::PCAConnectorAD::Template`, `AWS::QuickSight::Template`, `AWS::SES::Template`: `AWS::ACMPCA::Certificate.TemplateArn`, `AWS::CustomerProfiles::ObjectType.TemplateId`
- **2** among `AWS::CloudFront::Function`, `AWS::Lambda::Function`, `AWS::MediaTailor::Function`: `AWS::IoT::Authorizer.AuthorizerFunctionArn`, `AWS::IoT::CertificateProvider.LambdaFunctionArn`
- **2** among `AWS::CloudHSM::Cluster`, `AWS::DSQL::Cluster`, `AWS::DocDBElastic::Cluster`, `AWS::ECS::Cluster`, `AWS::EKS::Cluster`, `AWS::MSK::Cluster`, `AWS::MediaLive::Cluster`, `AWS::MemoryDB::Cluster`, `AWS::PCS::Cluster`, `AWS::Redshift::Cluster`, `AWS::Route53RecoveryControl::Cluster`, `AWS::SageMaker::Cluster`: `AWS::ElastiCache::ReplicationGroup.PrimaryClusterId`, `AWS::ElastiCache::ReplicationGroup.SnapshottingClusterId`
- **2** among `AWS::EC2::VPCEndpoint`, `AWS::OpenSearchServerless::VpcEndpoint`: `AWS::DataSync::Agent.VpcEndpointId`, `AWS::PCAConnectorSCEP::Connector.VpcEndpointId`
- **2** among `AWS::IoTWireless::Destination`, `AWS::Logs::Destination`: `AWS::Route53Resolver::ResolverQueryLoggingConfig.DestinationArn`, `AWS::VpcLattice::AccessLogSubscription.DestinationArn`
- **2** among `AWS::MediaLive::Network`, `AWS::Wickr::Network`: `AWS::EC2::TrafficMirrorSession.VirtualNetworkId`, `AWS::ODB::OdbPeeringConnection.PeerNetworkId`
- **2** among `AWS::SageMaker::Artifact`, `AWS::SecurityAgent::Artifact`: `AWS::ServiceCatalog::CloudFormationProvisionedProduct.ProvisioningArtifactId`, `AWS::ServiceCatalog::ServiceActionAssociation.ProvisioningArtifactId`
- **1** among `AWS::ACMPCA::Permission`, `AWS::Lambda::Permission`, `AWS::QBusiness::Permission`, `AWS::RAM::Permission`: `AWS::EC2::FlowLog.DeliverLogsPermissionArn`
- **1** among `AWS::APS::Workspace`, `AWS::AWSExternalAnthropic::Workspace`, `AWS::Connect::Workspace`, `AWS::Grafana::Workspace`, `AWS::IoTSiteWise::Workspace`, `AWS::IoTTwinMaker::Workspace`, `AWS::WorkSpaces::Workspace`: `AWS::Chatbot::SlackChannelConfiguration.SlackWorkspaceId`
- **1** among `AWS::AccountAccess::Entitlement`, `AWS::AppStream::Entitlement`: `AWS::MediaConnect::FlowSource.EntitlementArn`
- **1** among `AWS::AmazonMQ::Configuration`, `AWS::MSK::Configuration`, `AWS::Omics::Configuration`: `AWS::DataZone::Environment.EnvironmentConfigurationId`
- **1** among `AWS::ApiGatewayV2::Integration`, `AWS::CustomerProfiles::Integration`, `AWS::Glue::Integration`, `AWS::Logs::Integration`, `AWS::RDS::Integration`, `AWS::Redshift::Integration`: `AWS::Connect::IntegrationAssociation.IntegrationArn`
- **1** among `AWS::AppConfig::Environment`, `AWS::DataZone::Environment`, `AWS::EVS::Environment`, `AWS::ElasticBeanstalk::Environment`, `AWS::FinSpace::Environment`, `AWS::M2::Environment`, `AWS::MWAA::Environment`, `AWS::RefactorSpaces::Environment`, `AWS::WorkSpacesThinClient::Environment`: `AWS::Interconnect::Connection.EnvironmentId`
- **1** among `AWS::AppFlow::Connector`, `AWS::Config::Connector`, `AWS::InspectorV2::Connector`, `AWS::KafkaConnect::Connector`, `AWS::PCAConnectorAD::Connector`, `AWS::PCAConnectorSCEP::Connector`, `AWS::SecurityHub::Connector`, `AWS::Transfer::Connector`: `AWS::SSM::CloudConnector.ConfigConnectorArn`
- **1** among `AWS::AppStream::User`, `AWS::Connect::User`, `AWS::ElastiCache::User`, `AWS::IAM::User`, `AWS::IdentityStore::User`, `AWS::MemoryDB::User`, `AWS::Transfer::User`: `AWS::IVS::IngestConfiguration.UserId`
- **1** among `AWS::BedrockAgentCore::Dataset`, `AWS::DataBrew::Dataset`, `AWS::DataExchange::DataSet`, `AWS::Forecast::Dataset`, `AWS::IoTAnalytics::Dataset`, `AWS::IoTSiteWise::Dataset`, `AWS::Personalize::Dataset`, `AWS::QuickSight::DataSet`, `AWS::Rekognition::Dataset`, `AWS::SCN::Dataset`: `AWS::Connect::DataLakeAssociation.DataSetId`
- **1** among `AWS::CodeConnections::Host`, `AWS::EC2::Host`: `AWS::CodeStarConnections::Connection.HostArn`
- **1** among `AWS::Connect::Rule`, `AWS::Events::Rule`, `AWS::Rbin::Rule`, `AWS::VpcLattice::Rule`: `AWS::GuardDuty::CustomDetectionRuleAssociation.RuleId`
- **1** among `AWS::DMS::Endpoint`, `AWS::EMRContainers::Endpoint`, `AWS::Events::Endpoint`, `AWS::S3Outposts::Endpoint`, `AWS::SageMaker::Endpoint`: `AWS::EC2::TrafficMirrorTarget.GatewayLoadBalancerEndpointId`
- **1** among `AWS::DMS::InstanceProfile`, `AWS::IAM::InstanceProfile`: `AWS::PCS::ComputeNodeGroup.IamInstanceProfileArn`
- **1** among `AWS::DataSync::Task`, `AWS::IoTSiteWise::Task`: `AWS::SSM::MaintenanceWindowTask.TaskArn`
- **1** among `AWS::DevOpsAgent::Asset`, `AWS::IoTSiteWise::Asset`, `AWS::MediaPackage::Asset`: `AWS::EC2::Host.AssetId`
- **1** among `AWS::EFS::AccessPoint`, `AWS::S3::AccessPoint`, `AWS::S3Express::AccessPoint`, `AWS::S3Files::AccessPoint`, `AWS::S3ObjectLambda::AccessPoint`, `AWS::S3Outposts::AccessPoint`: `AWS::DataSync::LocationEFS.AccessPointArn`
- **1** among `AWS::Glue::Database`, `AWS::Lightsail::Database`, `AWS::Timestream::Database`: `AWS::SystemsManagerSAP::Application.DatabaseArn`
- **1** among `AWS::NetworkManager::Device`, `AWS::SageMaker::Device`: `AWS::Braket::SpendingLimit.DeviceArn`
- **1** among `AWS::NetworkManager::Site`, `AWS::Outposts::Site`: `AWS::EVS::Environment.SiteId`
- **1** among `AWS::PricingPlanManager::Subscription`, `AWS::SNS::Subscription`: `AWS::SES::MailManagerAddonInstance.AddonSubscriptionId`
- **1** among `AWS::RedshiftServerless::Namespace`, `AWS::S3Tables::Namespace`, `AWS::SCN::Namespace`: `AWS::ServiceDiscovery::Service.NamespaceId`

## Unresolved: target lacks the attribute (84)

One type matched, but it has no `Id`/`Arn`, `<Segment>Id`/`<Segment>Arn`, or single matching primary identifier. Many are correct negatives: generic `ResourceArn`s, owner account ids, snapshots.

`AWS::AuditManager::Assessment.FrameworkId`, `AWS::AutoScaling::LaunchConfiguration.RamDiskId`, `AWS::Backup::RestoreTestingSelection.ProtectedResourceArns`, `AWS::Bedrock::ResourcePolicy.ResourceArn`, `AWS::BedrockAgentCore::ResourcePolicy.ResourceArn`, `AWS::Chatbot::MicrosoftTeamsChannelConfiguration.CustomizationResourceArns`, `AWS::Chatbot::MicrosoftTeamsChannelConfiguration.TeamsTenantId`, `AWS::Chatbot::SlackChannelConfiguration.CustomizationResourceArns`, `AWS::CloudFormation::HookDefaultVersion.VersionId`, `AWS::CloudFormation::ModuleDefaultVersion.VersionId`, `AWS::CloudFormation::ResourceDefaultVersion.VersionId`, `AWS::CloudTrail::ResourcePolicy.ResourceArn`, `AWS::CloudWatch::Alarm.ThresholdMetricId`, `AWS::CodeStarConnections::RepositoryLink.OwnerId`, `AWS::CodeStarNotifications::NotificationRule.EventTypeId`, `AWS::CodeStarNotifications::NotificationRule.EventTypeIds`, `AWS::Connect::ContactFlowModuleAlias.ContactFlowModuleId`, `AWS::Connect::ContactFlowModuleVersion.ContactFlowModuleId`, `AWS::Connect::ContactFlowVersion.ContactFlowId`, `AWS::Connect::User.DirectoryUserId`, `AWS::DataSync::Agent.SecurityGroupArns`, `AWS::DataSync::Agent.SubnetArns`, `AWS::DataSync::LocationFSxLustre.SecurityGroupArns`, `AWS::DataSync::LocationFSxONTAP.SecurityGroupArns`, `AWS::DataSync::LocationFSxOpenZFS.SecurityGroupArns`, `AWS::DataSync::LocationFSxWindows.SecurityGroupArns`, `AWS::EC2::CapacityReservation.PlacementGroupArn`, `AWS::EC2::CapacityReservation.UnusedReservationBillingOwnerId`, `AWS::EC2::EIPAssociation.AllocationId`, `AWS::EC2::Instance.RamdiskId`, `AWS::EC2::NatGateway.AllocationId`, `AWS::EC2::NatGateway.SecondaryAllocationIds`, `AWS::EC2::SecurityGroupIngress.SourceSecurityGroupOwnerId`, `AWS::EC2::TrafficMirrorSession.OwnerId`, `AWS::EC2::VPCPeeringConnection.PeerOwnerId`, `AWS::EC2::Volume.SnapshotId`, `AWS::ECS::Daemon.CapacityProviderArns`, `AWS::ElastiCache::CacheCluster.SnapshotArns`, `AWS::ElastiCache::ReplicationGroup.SnapshotArns`, `AWS::ElasticBeanstalk::ConfigurationTemplate.EnvironmentId`, `AWS::ElasticBeanstalk::ConfigurationTemplate.PlatformArn`, `AWS::ElasticBeanstalk::Environment.PlatformArn`, `AWS::GameLift::MatchmakingConfiguration.RuleSetArn`, `AWS::Glue::IntegrationResourceProperty.ResourceArn`, `AWS::Kinesis::ResourcePolicy.ResourceArn`, `AWS::Lambda::ResourcePolicy.ResourceArn`, `AWS::Lex::ResourcePolicy.ResourceArn`, `AWS::Logs::DeliveryDestination.DestinationResourceArn`, `AWS::Logs::DeliverySource.ResourceArn`, `AWS::MemoryDB::Cluster.SnapshotArns`, `AWS::NetworkManager::ConnectPeer.SubnetArn`, `AWS::NetworkManager::CustomerGatewayAssociation.CustomerGatewayArn`, `AWS::NetworkManager::SiteToSiteVpnAttachment.VpnConnectionArn`, `AWS::NetworkManager::TransitGatewayRouteTableAttachment.TransitGatewayRouteTableArn`, `AWS::NetworkManager::VpcAttachment.SubnetArns`, `AWS::NetworkManager::VpcAttachment.VpcArn`, `AWS::Omics::RunCache.CacheBucketOwnerId`, `AWS::Omics::Workflow.WorkflowBucketOwnerId`, `AWS::Omics::WorkflowVersion.WorkflowBucketOwnerId`, `AWS::PaymentCryptography::Alias.KeyArn`, `AWS::Personalize::Solution.RecipeArn`, `AWS::PricingPlanManager::Subscription.ResourceArns`, `AWS::QuickSight::KnowledgeBase.PrimaryOwnerArn`, `AWS::RAM::ResourceShare.ResourceArns`, `AWS::RDS::DBInstance.DomainAuthSecretArn`, `AWS::RedshiftServerless::Workgroup.SnapshotArn`, `AWS::Route53GlobalResolver::HostedZoneAssociation.ResourceArn`, `AWS::Route53Profiles::ProfileResourceAssociation.ResourceArn`, `AWS::SES::MailManagerIngressPoint.RuleSetId`, `AWS::SMSVOICE::ResourcePolicy.ResourceArn`, `AWS::SSM::ResourcePolicy.ResourceArn`, `AWS::SSMContacts::ContactChannel.ContactId`, `AWS::SSMContacts::Plan.ContactId`, `AWS::SSMContacts::Plan.RotationIds`, `AWS::SSMContacts::Rotation.ContactIds`, `AWS::ServerlessRepo::Application.SpdxLicenseId`, `AWS::Shield::Protection.HealthCheckArns`, `AWS::Shield::Protection.ResourceArn`, `AWS::SupportApp::SlackWorkspaceConfiguration.VersionId`, `AWS::Synthetics::Group.ResourceArns`, `AWS::VpcLattice::ResourcePolicy.ResourceArn`, `AWS::WAFv2::LoggingConfiguration.ResourceArn`, `AWS::WAFv2::WebACLAssociation.ResourceArn`, `AWS::WorkspacesInstances::Volume.SnapshotId`

## Unresolved: no such type (143)

No provisionable type's last segment matches, even as a suffix. These are mostly correct negatives, such as AMIs, availability zones and kernels, which CloudFormation does not manage.

`AWS::AmplifyUIBuilder::Component.SourceId`, `AWS::ApiGateway::ApiKey.CustomerId`, `AWS::ApiGateway::ApiKey.GenerateDistinctId`, `AWS::ApiGateway::Resource.ParentId`, `AWS::ApiGateway::VpcLink.TargetArns`, `AWS::ApiGatewayV2::Api.CredentialsArn`, `AWS::ApiGatewayV2::Authorizer.AuthorizerCredentialsArn`, `AWS::ApiGatewayV2::Integration.CredentialsArn`, `AWS::AppFlow::ConnectorProfile.KMSArn`, `AWS::AppFlow::Flow.KMSArn`, `AWS::ApplicationAutoScaling::ScalingPolicy.ScalingTargetId`, `AWS::AutoScaling::AutoScalingGroup.AvailabilityZoneIds`, `AWS::AutoScaling::LaunchConfiguration.KernelId`, `AWS::BcmPricingCalculator::BillScenario.CostCategoryGroupSharingPreferenceArn`, `AWS::Cases::Case.CustomerId`, `AWS::Chatbot::MicrosoftTeamsChannelConfiguration.TeamId`, `AWS::CloudWatch::MetricStream.FirehoseArn`, `AWS::Cognito::ManagedLoginBranding.ClientId`, `AWS::Cognito::Terms.ClientId`, `AWS::Cognito::UserPoolRiskConfigurationAttachment.ClientId`, `AWS::Cognito::UserPoolUICustomizationAttachment.ClientId`, `AWS::Config::RemediationConfiguration.TargetId`, `AWS::Connect::Instance.DirectoryId`, `AWS::Connect::PhoneNumber.TargetArn`, `AWS::DMS::EventSubscription.SourceIds`, `AWS::DataBrew::Ruleset.TargetArn`, `AWS::DataSync::LocationFSxONTAP.StorageVirtualMachineArn`, `AWS::DataZone::GroupProfile.RolePrincipalArn`, `AWS::Deadline::MeteredProduct.ProductId`, `AWS::DocDB::EventSubscription.SourceIds`, `AWS::DynamoDB::GlobalTable.GlobalTableSourceArn`, `AWS::EC2::CapacityReservation.AvailabilityZoneId`, `AWS::EC2::CapacityReservation.OutPostArn`, `AWS::EC2::Host.OutpostArn`, `AWS::EC2::Instance.KernelId`, `AWS::EC2::NetworkInsightsAnalysis.FilterInArns`, `AWS::EC2::NetworkInsightsAnalysis.FilterOutArns`, `AWS::EC2::Subnet.AvailabilityZoneId`, `AWS::EC2::Subnet.OutpostArn`, `AWS::EC2::TransitGatewayMeteringPolicy.MiddleboxAttachmentIds`, `AWS::EC2::Volume.AvailabilityZoneId`, `AWS::EC2::Volume.OutpostArn`, `AWS::ECR::PullThroughCacheRule.CredentialArn`, `AWS::ECR::PullTimeUpdateExclusion.PrincipalArn`, `AWS::ECS::TaskSet.ExternalId`, `AWS::EKS::AccessEntry.PrincipalArn`, `AWS::EntityResolution::PolicyStatement.StatementId`, `AWS::EventSchemas::Discoverer.SourceArn`, `AWS::EventSchemas::RegistryPolicy.RevisionId`, `AWS::Events::Archive.SourceArn`, `AWS::Events::EventBusPolicy.StatementId`, `AWS::Glue::Integration.SourceArn`, `AWS::Glue::Integration.TargetArn`, `AWS::GreengrassV2::Deployment.ParentTargetArn`, `AWS::GreengrassV2::Deployment.TargetArn`, `AWS::GuardDuty::Master.InvitationId`, `AWS::IdentityStore::Group.IdentityStoreId`, `AWS::IdentityStore::GroupMembership.IdentityStoreId`, `AWS::IdentityStore::User.IdentityStoreId`, `AWS::IoT::SecurityProfile.TargetArns`, `AWS::IoTFleetWise::Campaign.TargetArn`, `AWS::IoTSiteWise::Asset.AssetExternalId`, `AWS::IoTSiteWise::AssetModel.AssetModelExternalId`, `AWS::Lambda::EventSourceMapping.EventSourceArn`, `AWS::Lambda::Permission.SourceArn`, `AWS::Lightsail::Bucket.BundleId`, `AWS::Lightsail::Database.RelationalDatabaseBundleId`, `AWS::Lightsail::Distribution.BundleId`, `AWS::Lightsail::Instance.BundleId`, `AWS::Location::TrackerConsumer.ConsumerArn`, `AWS::Logs::Destination.TargetArn`, `AWS::MediaConnect::Bridge.PlacementArn`, `AWS::MediaConnect::FlowOutput.RemoteId`, `AWS::MediaLive::SignalMap.DiscoveryEntryPointArn`, `AWS::MediaPackage::Asset.SourceArn`, `AWS::Neptune::EventSubscription.SourceIds`, `AWS::NetworkFlowMonitor::Monitor.ScopeArn`, `AWS::NetworkManager::ConnectAttachment.TransportAttachmentId`, `AWS::NetworkManager::TransitGatewayRouteTableAttachment.PeeringId`, `AWS::ODB::CloudExadataInfrastructure.AvailabilityZoneId`, `AWS::ODB::OdbNetwork.AvailabilityZoneId`, `AWS::Organizations::Account.ParentIds`, `AWS::Organizations::OrganizationalUnit.ParentId`, `AWS::Organizations::Policy.TargetIds`, `AWS::PCAConnectorAD::Connector.DirectoryId`, `AWS::PCAConnectorAD::DirectoryRegistration.DirectoryId`, `AWS::PCS::ComputeNodeGroup.AmiId`, `AWS::QBusiness::Permission.StatementId`, `AWS::QuickSight::Agent.IconId`, `AWS::QuickSight::OAuthClientApplication.ClientId`, `AWS::RDS::DBInstance.SourceDBInstanceAutomatedBackupsArn`, `AWS::RDS::DBInstance.TdeCredentialArn`, `AWS::RDS::EventSubscription.SourceIds`, `AWS::RDS::Integration.SourceArn`, `AWS::RDS::Integration.TargetArn`, `AWS::Redshift::EventSubscription.SourceIds`, `AWS::Redshift::Integration.SourceArn`, `AWS::Redshift::Integration.TargetArn`, `AWS::RedshiftServerless::Workgroup.RecoveryPointId`, `AWS::RoboMaker::RobotApplication.CurrentRevisionId`, `AWS::RoboMaker::RobotApplicationVersion.CurrentRevisionId`, `AWS::RoboMaker::SimulationApplication.CurrentRevisionId`, `AWS::RoboMaker::SimulationApplicationVersion.CurrentRevisionId`, `AWS::Route53Resolver::OutpostResolver.OutpostArn`, `AWS::Route53Resolver::ResolverEndpoint.OutpostArn`, `AWS::S3::AccessGrantsInstance.IdentityCenterArn`, `AWS::S3Outposts::Bucket.OutpostId`, `AWS::S3Outposts::Endpoint.OutpostId`, `AWS::SMSVOICE::ConfigurationSet.DefaultSenderId`, `AWS::SMSVOICE::SenderId.SenderId`, `AWS::SSM::MaintenanceWindowTarget.WindowId`, `AWS::SSM::MaintenanceWindowTask.WindowId`, `AWS::SSM::ServiceSetting.SettingId`, `AWS::SSMContacts::Rotation.TimeZoneId`, `AWS::SSO::Application.ApplicationProviderArn`, `AWS::SSO::ApplicationAssignment.PrincipalId`, `AWS::SSO::Assignment.PrincipalId`, `AWS::SSO::Assignment.TargetId`, `AWS::SecretsManager::SecretTargetAttachment.TargetId`, `AWS::SecurityHub::PolicyAssociation.TargetId`, `AWS::SecurityHub::ProductSubscription.ProductArn`, `AWS::SecurityHub::Standard.StandardsArn`, `AWS::ServiceCatalog::CloudFormationProvisionedProduct.PathId`, `AWS::ServiceCatalog::CloudFormationProvisionedProduct.ProductId`, `AWS::ServiceCatalog::LaunchNotificationConstraint.ProductId`, `AWS::ServiceCatalog::LaunchRoleConstraint.ProductId`, `AWS::ServiceCatalog::LaunchTemplateConstraint.ProductId`, `AWS::ServiceCatalog::PortfolioProductAssociation.ProductId`, `AWS::ServiceCatalog::ResourceUpdateConstraint.ProductId`, `AWS::ServiceCatalog::ServiceActionAssociation.ProductId`, `AWS::ServiceCatalog::StackSetConstraint.ProductId`, `AWS::Signer::ProfilePermission.StatementId`, `AWS::StepFunctions::StateMachineVersion.StateMachineRevisionId`, `AWS::SupportApp::SlackChannelConfiguration.TeamId`, `AWS::SupportApp::SlackWorkspaceConfiguration.TeamId`, `AWS::SupportAuthZ::SupportPermit.SupportCaseDisplayId`, `AWS::Transfer::Profile.As2Id`, `AWS::WorkSpaces::Workspace.BundleId`, `AWS::WorkSpaces::Workspace.DirectoryId`, `AWS::WorkSpaces::WorkspacesPool.BundleId`, `AWS::WorkSpaces::WorkspacesPool.DirectoryId`, `AWS::WorkSpacesThinClient::Environment.DesiredSoftwareSetId`, `AWS::WorkSpacesThinClient::Environment.DesktopArn`
