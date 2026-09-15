# aws.publisher

**CloudFormation type:** `AWS::CloudFormation::Publisher`

Register as a publisher in the CloudFormation Registry.

Region attribute: `region`

**Import ID:** `<region>/PublisherId` (AWS::CloudFormation::Publisher)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcceptTermsAndConditions` | accept_terms_and_conditions | `boolean` | required, replaces on change |  | Whether you accept the terms and conditions for publishing extensions in the CloudFormation registry. You must accept the terms and conditions in order to publish public extensions to the CloudFormation registry. The terms and conditions can be found at https://cloudformation-registry-documents.s3.amazonaws.com/Terms_and_Conditions_for_AWS_CloudFormation_Registry_Publishers.pdf |
| `ConnectionArn` | connection_arn | `string` | optional, computed, provider-chosen, replaces on change |  | If you are using a Bitbucket or GitHub account for identity verification, the Amazon Resource Name (ARN) for your connection to that account. |
| `IdentityProvider` | identity_provider | `string` | computed |  | The type of account used as the identity provider when registering this publisher with CloudFormation. |
| `PublisherId` | publisher_id | `string` | computed |  | The reserved publisher id for this type, or the publisher id assigned by CloudFormation for publishing in this region. |
| `PublisherProfile` | publisher_profile | `string` | computed |  | The URL to the publisher's profile with the identity provider. |
| `PublisherStatus` | publisher_status | `string` | computed |  | Whether the publisher is verified. |

Supports update: no

Discovery: supported
