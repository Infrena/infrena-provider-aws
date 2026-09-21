# aws.pricingplanmanager.subscription

**CloudFormation type:** `AWS::PricingPlanManager::Subscription`

Resource type definition for AWS::PricingPlanManager::Subscription. Deleting an activated subscription does not terminate it immediately; it schedules a cancellation that takes effect at the end of the current billing period. Until that date the subscription remains active and billing continues. Deleting a subscription that has not yet been activated (PENDING_APPROVAL status) removes it immediately with no further charges.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::PricingPlanManager::Subscription)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the subscription. |
| `CreatedAt` | created_at | `string` | computed |  | The time the subscription was created, in ISO 8601 format. |
| `CurrentPlanTier` | current_plan_tier | `string` | computed |  | The plan tier currently active on the subscription as reported by the API. Populated by the Read handler. Diverges from PlanTier after a non-reversible CloudFormation rollback; surface via drift detection. |
| `PlanFamily` | plan_family | `string` | required, replaces on change |  | The name of the pricing plan family. |
| `PlanTier` | plan_tier | `string` | required |  | The tier of the pricing plan. CloudFormation does not change the tier of an existing subscription; a stack update that changes the tier, upgrading or downgrading it, is rejected. |
| `ResourceArns` | resource_arns | `list` | required |  | The ARNs of resources associated with the subscription. |
| `Status` |  | `string` | computed |  | The status of the subscription. PENDING_APPROVAL means a paid-tier subscription has been created but is not yet active and incurs no charges until it is approved out of band via a separate ApprovePaidSubscription call; CloudFormation never approves it. Free-tier subscriptions are activated immediately and do not use this status. ACTIVE means the subscription is in effect and, for paid tiers, billing has started. SYNC_IN_PROGRESS means a change is being applied. FAILED means provisioning did not complete; see StatusReason. |
| `StatusReason` | status_reason | `string` | computed |  | A human-readable explanation of why the subscription is in its current status. Populated only when Status is FAILED, where it carries the reason the subscription could not be provisioned. Empty for all other statuses. |
| `UpdatedAt` | updated_at | `string` | computed |  | The time the subscription was last modified, in ISO 8601 format. |
| `UsageLevel` | usage_level | `string` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
