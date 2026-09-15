# S3 Private Bucket Module

Creates a secure S3 bucket with encryption, versioning, and public access blocking.

## Resources

- S3 bucket with:
  - Versioning enabled
  - Server-side encryption (AES256)
  - Public access block (all options enabled)

## Inputs

| Input | Type | Default | Description |
|-------|------|---------|-------------|
| `name` | string | (required) | S3 bucket name (must be globally unique) |
| `tags` | map | `{}` | Key-value tags to apply to the bucket |

## Outputs

| Output | Description |
|--------|-------------|
| `bucket_name` | The S3 bucket name |
| `bucket_arn` | The S3 bucket ARN |

## Usage

```yaml
modules:
  - ./modules/s3-private-bucket

resources:
  data_bucket:
    type: module.s3_private_bucket
    name: my-app-data-bucket-12345
    tags:
      Environment: production
      Application: data-pipeline
```

Note: S3 bucket names must be globally unique across AWS. Bucket names must be lowercase letters, numbers, periods, and hyphens.
