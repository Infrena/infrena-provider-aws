# aws.accessgrant

**CloudFormation type:** `AWS::S3::AccessGrant`

The AWS::S3::AccessGrant resource is an Amazon S3 resource type representing permissions to a specific S3 bucket or prefix hosted in an S3 Access Grants instance.

Region attribute: `region`

**Import ID:** `<region>/AccessGrantId` (AWS::S3::AccessGrant)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessGrantArn` | access_grant_arn | `string` | computed |  | the Amazon Resource Name (ARN) of the specified access grant. |
| `AccessGrantId` | access_grant_id | `string` | computed |  | The ID assigned to this access grant. |
| `AccessGrantsLocationConfiguration` | access_grants_location_configuration | `map` | optional, computed, provider-chosen |  | The configuration options of the grant location, which is the S3 path to the data to which you are granting access. |
| `AccessGrantsLocationId` | access_grants_location_id | `string` | required | aws.accessgrantslocation.AccessGrantsLocationId | The custom S3 location to be accessed by the grantee |
| `ApplicationArn` | application_arn | `string` | optional, computed, provider-chosen |  | The ARN of the application grantees will use to access the location |
| `GrantScope` | grant_scope | `string` | computed |  | The S3 path of the data to which you are granting access. It is a combination of the S3 path of the registered location and the subprefix. |
| `Grantee` |  | `map` | required |  | The principal who will be granted permission to access S3. |
| `Permission` |  | `string` | required |  | The level of access to be afforded to the grantee |
| `S3PrefixType` | s3_prefix_type | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The type of S3SubPrefix. |
| `Tags` |  | `map` | replaces on change, tags map |  |  |

Supports update: yes

Discovery: supported
