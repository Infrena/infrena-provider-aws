# aws.primarytaskset

**CloudFormation type:** `AWS::ECS::PrimaryTaskSet`

A pseudo-resource that manages which of your ECS task sets is primary.

Region attribute: `region`

**Import ID:** `<region>/Cluster|Service` (AWS::ECS::PrimaryTaskSet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Cluster` |  | `string` | required, replaces on change |  | The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in. |
| `Service` |  | `string` | required, replaces on change |  | The short name or full Amazon Resource Name (ARN) of the service to create the task set in. |
| `TaskSetId` | task_set_id | `string` | required | aws.taskset.Id | The ID or full Amazon Resource Name (ARN) of the task set. |

Supports update: yes

Discovery: not supported
