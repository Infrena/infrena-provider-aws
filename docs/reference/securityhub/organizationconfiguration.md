# aws.organizationconfiguration

**CloudFormation type:** `AWS::SecurityHub::OrganizationConfiguration`

The AWS::SecurityHub::OrganizationConfiguration resource represents the configuration of your organization in Security Hub. Only the Security Hub administrator account can create Organization Configuration resource in each region and can opt-in to Central Configuration only in the aggregation region of FindingAggregator.

Region attribute: `region`

**Import ID:** `<region>/OrganizationConfigurationIdentifier` (AWS::SecurityHub::OrganizationConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AutoEnable` | auto_enable | `boolean` | required |  | Whether to automatically enable Security Hub in new member accounts when they join the organization. |
| `AutoEnableStandards` | auto_enable_standards | `string` | optional, computed, provider-chosen |  | Whether to automatically enable Security Hub default standards in new member accounts when they join the organization. |
| `ConfigurationType` | configuration_type | `string` | optional, computed, provider-chosen |  | Indicates whether the organization uses local or central configuration. |
| `MemberAccountLimitReached` | member_account_limit_reached | `boolean` | computed |  | Whether the maximum number of allowed member accounts are already associated with the Security Hub administrator account. |
| `OrganizationConfigurationIdentifier` | organization_configuration_identifier | `string` | computed |  | The identifier of the OrganizationConfiguration being created and assigned as the unique identifier. |
| `Status` |  | `string` | computed |  | Describes whether central configuration could be enabled as the ConfigurationType for the organization. |
| `StatusMessage` | status_message | `string` | computed |  | Provides an explanation if the value of Status is equal to FAILED when ConfigurationType is equal to CENTRAL. |

Supports update: yes

Discovery: supported
