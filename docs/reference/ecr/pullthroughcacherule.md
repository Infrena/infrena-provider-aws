# aws.pullthroughcacherule

**CloudFormation type:** `AWS::ECR::PullThroughCacheRule`

The ``AWS::ECR::PullThroughCacheRule`` resource creates or updates a pull through cache rule. A pull through cache rule provides a way to cache images from an upstream registry in your Amazon ECR private registry.

Region attribute: `region`

**Import ID:** `<region>/EcrRepositoryPrefix` (AWS::ECR::PullThroughCacheRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CredentialArn` | credential_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that identifies the credentials to authenticate to the upstream registry. |
| `CustomRoleArn` | custom_role_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.role.Arn | The ARN of the IAM role to be assumed by Amazon ECR to authenticate to ECR upstream registry. This role must be in the same account as the registry that you are configuring. |
| `EcrRepositoryPrefix` | ecr_repository_prefix | `string` | optional, computed, provider-chosen, replaces on change |  | The ECRRepositoryPrefix is a custom alias for upstream registry url. |
| `UpstreamRegistry` | upstream_registry | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The name of the upstream registry. |
| `UpstreamRegistryUrl` | upstream_registry_url | `string` | optional, computed, provider-chosen, replaces on change |  | The upstreamRegistryUrl is the endpoint of upstream registry url of the public repository to be cached |
| `UpstreamRepositoryPrefix` | upstream_repository_prefix | `string` | optional, computed, provider-chosen, replaces on change |  | The repository name prefix of upstream registry to match with the upstream repository name. When this field isn't specified, Amazon ECR will use the `ROOT`. |

Supports update: yes

Discovery: supported
