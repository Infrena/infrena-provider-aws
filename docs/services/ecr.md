# Amazon ECR (Elastic Container Registry)

Amazon ECR is a fully managed container image registry that makes it easy to push, pull, and manage Docker images. Use ECR to store container images for applications running on ECS, EKS, Lambda, and other services.

## ECR Repository

An ECR repository is a private Docker image registry where you push and pull container images. Each repository can be configured with image tag mutability, scanning, and lifecycle policies.

### Key attributes

- **repository_name**: Name for the repository (can include namespace like `project-a/nginx`). Defaults to an auto-generated name if omitted.
- **image_tag_mutability**: Tag mutability setting (`MUTABLE` or `IMMUTABLE`). Defaults to `MUTABLE`, allowing tags to be overwritten. Set to `IMMUTABLE` to prevent tag reuse.
- **image_scanning_configuration**: Configuration for scanning images on push. Specify `ScanOnPush: true` to automatically scan pushed images for vulnerabilities.
- **encryption_configuration**: Encryption configuration for stored images. Defaults to AWS-managed keys if omitted. Specify a customer managed KMS key ARN for encrypted storage.
- **lifecycle_policy**: Lifecycle policy as a JSON block string to automatically expire images. Use `imageSizeGreaterThanBytes` and `imageCountMoreThan` to define retention rules.
- **repository_policy_text**: JSON repository policy text (as a string) controlling who can push and pull images. Defaults to private (no policy).
- **empty_on_delete**: Force delete repository contents when destroying (optional, defaults to `false`).
- **tags**: Metadata tags for the repository

### Example

```yaml
app_repo:
  type: aws.ecr.repository
  repository_name: myapp/api
  image_tag_mutability: IMMUTABLE
  image_scanning_configuration:
    ScanOnPush: true
  lifecycle_policy:
    Rules:
      - RulePriority: 1
        Description: Expire untagged images after 7 days
        Selection:
          TagStatus: untagged
          CountUnit: imageCountMoreThan
          CountNumber: 0
          DaysSinceImagePushed: 7
        Action:
          Type: expire
  tags:
    Environment: production
    Application: myapp
```

Note: The lifecycle policy is a JSON object, not a string. The plugin accepts it as a map and serializes it correctly.

## Common pitfalls

- **Empty repository cannot be deleted**: By default, ECR refuses to delete repositories with images. Set `empty_on_delete: true` to force deletion.
- **Tag immutability is permanent once set**: Once you set `image_tag_mutability: IMMUTABLE`, you must create new versions of images with different tags. You cannot overwrite existing tags.
- **Lifecycle policies expire in UTC**: Date-based expiration rules use UTC time. A rule set to expire images `7` days old uses UTC now, not local time.
- **Repository policy is optional**: Without a policy, the repository remains private to the AWS account that created it. Only set `repository_policy_text` if you need cross-account access.
- **Scanning requires explicit configuration**: Images are not scanned for vulnerabilities by default. Set `image_scanning_configuration: {ScanOnPush: true}` to enable automated scanning.

## Reference pages

- [aws.ecr.repository](../reference/ecr/ecr-repository.md) — Repository configuration and attributes
