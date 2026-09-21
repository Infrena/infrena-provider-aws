# aws.launchwizard.deployment

**CloudFormation type:** `AWS::LaunchWizard::Deployment`

Definition of AWS::LaunchWizard::Deployment Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::LaunchWizard::Deployment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | ARN of the LaunchWizard deployment |
| `CreatedAt` | created_at | `string` | computed |  | Timestamp of LaunchWizard deployment creation |
| `DeletedAt` | deleted_at | `string` | computed |  | Timestamp of LaunchWizard deployment deletion |
| `DeploymentId` | deployment_id | `string` | computed |  | Deployment ID of the LaunchWizard deployment |
| `DeploymentPatternName` | deployment_pattern_name | `string` | required, replaces on change |  | Workload deployment pattern name |
| `Name` |  | `string` | required, replaces on change |  | Name of LaunchWizard deployment |
| `ResourceGroup` | resource_group | `string` | computed |  | Resource Group Name created for LaunchWizard deployment |
| `Specifications` |  | `map` | optional, computed, provider-chosen, write-only |  | LaunchWizard deployment specifications |
| `Status` |  | `string` | computed |  | Status of LaunchWizard deployment |
| `Tags` |  | `map` | tags map |  | Tags for LaunchWizard deployment |
| `WorkloadName` | workload_name | `string` | required, replaces on change |  | Workload Name for LaunchWizard deployment |

Supports update: yes

Discovery: supported
