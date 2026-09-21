# IAM

AWS Identity and Access Management (IAM) is a service that helps you manage access to AWS resources. With IAM, you can create users, roles, and policies to control who can do what in your AWS account.

IAM resources are global — they have no region attribute. They are managed in `us-east-1` by the plugin.

## Role

An IAM role is an entity that holds a set of permissions. Roles are assumed by services, users, or applications (via `STS AssumeRole`), and they define what those entities can do in your AWS account. Roles are commonly used by EC2 instances and Lambda functions to access other AWS services.

**Type:** `aws.role`

**Import ID:** `global/RoleName`

**Settable attributes:**
- `name` (alias for `RoleName`): a name for the role, up to 64 characters
- `assume_role_policy` (alias for `AssumeRolePolicyDocument`): the trust policy as a JSON document, typically written as a YAML block string
- `description`: a description of the role
- `max_session_duration`: the maximum session duration in seconds (default 3600)
- `path`: the path to the role (default `/`)
- `managed_policy_arns`: a list of ARNs of managed policies to attach
- `tags`: a map of tags

Example:

```yaml
  app_role:
    type: aws.role
    name: app-execution-role
    description: Execution role for the application
    assume_role_policy: |
      {
        "Version": "2012-10-17",
        "Statement": [
          {
            "Effect": "Allow",
            "Principal": {
              "Service": "ec2.amazonaws.com"
            },
            "Action": "sts:AssumeRole"
          }
        ]
      }
    managed_policy_arns:
      - arn:aws:iam::aws:policy/CloudWatchLogsFullAccess
    tags:
      Environment: production
      Team: platform
```

## Managed Policy

A managed policy is a standalone policy that has its own ARN. Managed policies can be attached to users, groups, or roles. Unlike inline policies, managed policies can be reused across multiple identities.

**Type:** `aws.managedpolicy`

**Import ID:** `global/PolicyArn`

**Settable attributes:**
- `name` (alias for `ManagedPolicyName`): a friendly name for the policy
- `policy` (alias for `PolicyDocument`): the policy document as a JSON string, typically written as a YAML block string
- `description`: a friendly description of the policy
- `path`: the path for the policy (default `/`)
- `roles`: list of role names to attach the policy to
- `users`: list of user names to attach the policy to
- `groups`: list of group names to attach the policy to

Example:

```yaml
  app_policy:
    type: aws.managedpolicy
    name: app-s3-access
    description: Policy for S3 bucket access
    policy: |
      {
        "Version": "2012-10-17",
        "Statement": [
          {
            "Effect": "Allow",
            "Action": ["s3:GetObject", "s3:PutObject"],
            "Resource": "arn:aws:s3:::my-bucket/*"
          }
        ]
      }
    roles:
      - app-execution-role
```

## Instance Profile

An instance profile is a container for an IAM role. It holds the role that an EC2 instance can assume. When you create an EC2 instance, you specify an instance profile, and the instance can then use the role's permissions to access AWS services.

**Type:** `aws.iam.instanceprofile`

**Import ID:** `global/InstanceProfileName`

**Settable attributes:**
- `instance_profile_name` (alias for `InstanceProfileName`): the name of the instance profile
- `path`: the path to the instance profile (default `/`)
- `roles`: list of role names to associate with the profile (required; only one role can be assigned)

Example:

```yaml
  app_instance_profile:
    type: aws.iam.instanceprofile
    instance_profile_name: app-instance-profile
    roles:
      - app-execution-role
```

## User

An IAM user is an entity within your AWS account that has specific credentials and permissions. Unlike roles, users have permanent credentials (access keys, passwords) and are typically used for human access or application-specific access.

**Type:** `aws.iam.user`

**Import ID:** `global/UserName`

**Settable attributes:**
- `user_name` (alias for `UserName`): the name of the user to create
- `path`: the path for the user (default `/`)
- `managed_policy_arns`: a list of ARNs of managed policies to attach to the user
- `permissions_boundary`: the ARN of the managed policy used to set the permissions boundary
- `groups`: list of group names to add the user to
- `tags`: a map of tags

Example:

```yaml
  app_user:
    type: aws.iam.user
    user_name: app-service-user
    managed_policy_arns:
      - arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess
    tags:
      Environment: production
      Purpose: application-access
```

## Common patterns and pitfalls

- **Trust policies are required for roles.** The `assume_role_policy` attribute defines which services or principals can assume the role. It is a JSON policy document, not a YAML object.
- **Policy documents are JSON strings.** Both `assume_role_policy` and managed policy documents are written as JSON text in YAML block strings (`|` or `|-`). This preserves the exact JSON and avoids parsing issues.
- **Managed policy ARNs are idempotent.** If you attach a managed policy that is already attached, the operation is safe and idempotent.
- **Instance profiles require a role.** An instance profile must have at least one role assigned. A role can only be assigned to one instance profile at a time.
- **IAM is global.** There is no region attribute for IAM resources. They are managed in `us-east-1` regardless of which region you specify in your configuration.

## Reference pages

See the generated reference pages for the complete attribute list and behavior:

- [aws.role](../reference/iam/role.md)
- [aws.managedpolicy](../reference/iam/managedpolicy.md)
- [aws.iam.instanceprofile](../reference/iam/iam-instanceprofile.md)
- [aws.iam.user](../reference/iam/iam-user.md)
