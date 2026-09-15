# aws.datarepositoryassociation

**CloudFormation type:** `AWS::FSx::DataRepositoryAssociation`

Resource Type definition for AWS::FSx::DataRepositoryAssociation

Region attribute: `region`

**Import ID:** `<region>/AssociationId` (AWS::FSx::DataRepositoryAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssociationId` | association_id | `string` | computed |  | The system-generated, unique ID of the data repository association. |
| `BatchImportMetaDataOnCreate` | batch_import_meta_data_on_create | `boolean` | optional, computed, provider-chosen, replaces on change |  | A boolean flag indicating whether an import data repository task to import metadata should run after the data repository association is created. The task runs if this flag is set to true. |
| `DataRepositoryPath` | data_repository_path | `string` | required, replaces on change |  | The path to the Amazon S3 data repository that will be linked to the file system. The path can be an S3 bucket or prefix in the format s3://myBucket/myPrefix/ . This path specifies where in the S3 data repository files will be imported from or exported to. |
| `FileSystemId` | file_system_id | `string` | required, replaces on change |  | The globally unique ID of the file system, assigned by Amazon FSx. |
| `FileSystemPath` | file_system_path | `string` | required, replaces on change |  | This path specifies where in your file system files will be exported from or imported to. This file system directory can be linked to only one Amazon S3 bucket, and no other S3 bucket can be linked to the directory. |
| `ImportedFileChunkSize` | imported_file_chunk_size | `integer` | optional, computed, provider-chosen |  | For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system. |
| `ResourceARN` | resource_arn | `string` | computed |  | The Amazon Resource Name (ARN) for a given resource. ARNs uniquely identify Amazon Web Services resources. We require an ARN when you need to specify a resource unambiguously across all of Amazon Web Services. For more information, see Amazon Resource Names (ARNs) in the Amazon Web Services General Reference. |
| `S3` |  | `map` | optional, computed, provider-chosen |  | The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association. The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of Tag values, with a maximum of 50 elements. |

Supports update: yes

Discovery: supported
