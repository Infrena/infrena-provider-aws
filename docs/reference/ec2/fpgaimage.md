# aws.fpgaimage

**CloudFormation type:** `AWS::EC2::FpgaImage`

Creates and manages an Amazon FPGA Image (AFI) from a design checkpoint (DCP).

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::EC2::FpgaImage)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the FPGA image. |
| `CreateTime` | create_time | `string` | computed |  | The date and time the AFI was created. |
| `DataRetentionSupport` | data_retention_support | `boolean` | computed |  | Indicates whether data retention support is enabled for the AFI. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description for the AFI. |
| `FpgaImageGlobalId` | fpga_image_global_id | `string` | computed |  | The global FPGA image identifier (AGFI ID). |
| `FpgaImageId` | fpga_image_id | `string` | computed |  | The FPGA image identifier (AFI ID). |
| `InputStorageLocation` | input_storage_location | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | Describes a storage location in Amazon S3. |
| `LogsStorageLocation` | logs_storage_location | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | Describes a storage location in Amazon S3. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | A name for the AFI. |
| `OwnerId` | owner_id | `string` | computed |  | The ID of the AWS account that owns the AFI. |
| `Public` |  | `boolean` | computed |  | Indicates whether the AFI is public. |
| `State` |  | `string` | computed |  | The state of the AFI (pending \| available \| failed \| unavailable). |
| `Tags` |  | `map` | tags map |  | The tags assigned to the FPGA image. |
| `UpdateTime` | update_time | `string` | computed |  | The time of the most recent update to the AFI. |

Supports update: yes

Discovery: supported
