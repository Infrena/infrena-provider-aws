# aws.dlm.lifecyclepolicy

**CloudFormation type:** `AWS::DLM::LifecyclePolicy`

Resource Type definition for AWS::DLM::LifecyclePolicy

Region attribute: `region`

**Import ID:** `<region>/PolicyId` (AWS::DLM::LifecyclePolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the lifecycle policy. |
| `CopyTags` | copy_tags | `boolean` | optional, computed, provider-chosen |  | **[Default policies only]** Indicates whether the policy should copy tags from the source resource to the snapshot or AMI. If you do not specify a value, the default is false. |
| `CreateInterval` | create_interval | `integer` | optional, computed, provider-chosen |  | **[Default policies only]** Specifies how often the policy should run and create snapshots or AMIs. The creation frequency can range from 1 to 7 days. |
| `CrossRegionCopyTargets` | cross_region_copy_targets | `list` | optional, computed, provider-chosen |  | **[Default policies only]** Specifies destination Regions for snapshot or AMI copies. You can specify up to 3 destination Regions. If you do not want to create cross-Region copies, omit this parameter. |
| `DefaultPolicy` | default_policy | `string` | optional, computed, provider-chosen |  | **[Default policies only]** Specify the type of default policy to create. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the lifecycle policy. The characters ^[0-9A-Za-z _-]+$ are supported. |
| `Exclusions` |  | `map` | optional, computed, provider-chosen |  | **[Default policies only]** Specifies exclusion parameters for volumes or instances for which you do not want to create snapshots or AMIs. The policy will not create snapshots or AMIs for target resources that match any of the specified exclusion parameters. |
| `ExecutionRoleArn` | execution_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The Amazon Resource Name (ARN) of the IAM role used to run the operations specified by the lifecycle policy. |
| `ExtendDeletion` | extend_deletion | `boolean` | optional, computed, provider-chosen |  | **[Default policies only]** Defines the snapshot or AMI retention behavior for the policy if the source volume or instance is deleted, or if the policy enters the error, disabled, or deleted state. |
| `PolicyDetails` | policy_details | `map` | optional, computed, provider-chosen |  | The configuration details of the lifecycle policy. |
| `PolicyId` | policy_id | `string` | computed |  | The identifier of the lifecycle policy. |
| `RetainInterval` | retain_interval | `integer` | optional, computed, provider-chosen |  | **[Default policies only]** Specifies how long the policy should retain snapshots or AMIs before deleting them. The retention period can range from 2 to 14 days, but it must be greater than the creation frequency to ensure that the policy retains at least 1 snapshot or AMI at any given time. |
| `State` |  | `string` | optional, computed, provider-chosen |  | The activation state of the lifecycle policy. |
| `Tags` |  | `map` | tags map |  | The tags to apply to the lifecycle policy during creation. |

Supports update: yes

Discovery: supported
