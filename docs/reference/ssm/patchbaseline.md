# aws.patchbaseline

**CloudFormation type:** `AWS::SSM::PatchBaseline`

Resource Type definition for AWS::SSM::PatchBaseline

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::SSM::PatchBaseline)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApprovalRules` | approval_rules | `map` | optional, computed, provider-chosen |  | A set of rules defining the approval rules for a patch baseline. |
| `ApprovedPatches` | approved_patches | `list` | optional, computed, provider-chosen |  | A list of explicitly approved patches for the baseline. |
| `ApprovedPatchesComplianceLevel` | approved_patches_compliance_level | `string` | optional, computed, provider-chosen |  | Defines the compliance level for approved patches. This means that if an approved patch is reported as missing, this is the severity of the compliance violation. The default value is UNSPECIFIED. |
| `ApprovedPatchesEnableNonSecurity` | approved_patches_enable_non_security | `boolean` | optional, computed, provider-chosen |  | Indicates whether the list of approved patches includes non-security updates that should be applied to the instances. The default value is 'false'. Applies to Linux instances only. |
| `AvailableSecurityUpdatesComplianceStatus` | available_security_updates_compliance_status | `string` | optional, computed, provider-chosen |  | The compliance status for vendor recommended security updates that are not approved by this patch baseline. |
| `DefaultBaseline` | default_baseline | `boolean` | optional, computed, provider-chosen |  | Set the baseline as default baseline. Only registering to default patch baseline is allowed. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the patch baseline. |
| `GlobalFilters` | global_filters | `map` | optional, computed, provider-chosen |  | The patch filter group that defines the criteria for the rule. |
| `Id` |  | `string` | computed |  | The ID of the patch baseline. |
| `Name` |  | `string` | required |  | The name of the patch baseline. |
| `OperatingSystem` | operating_system | `string` | optional, computed, provider-chosen, replaces on change |  | Defines the operating system the patch baseline applies to. The Default value is WINDOWS. |
| `PatchGroups` | patch_groups | `list` | optional, computed, provider-chosen |  | PatchGroups is used to associate instances with a specific patch baseline |
| `RejectedPatches` | rejected_patches | `list` | optional, computed, provider-chosen |  | A list of explicitly rejected patches for the baseline. |
| `RejectedPatchesAction` | rejected_patches_action | `string` | optional, computed, provider-chosen |  | The action for Patch Manager to take on patches included in the RejectedPackages list. |
| `Sources` |  | `list` | optional, computed, provider-chosen |  | Information about the patches to use to update the instances, including target operating systems and source repository. Applies to Linux instances only. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Optional metadata that you assign to a resource. Tags enable you to categorize a resource in different ways. |

Supports update: yes

Discovery: supported
