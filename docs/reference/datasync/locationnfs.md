# aws.locationnfs

**CloudFormation type:** `AWS::DataSync::LocationNFS`

Resource schema for AWS::DataSync::LocationNFS

Region attribute: `region`

**Import ID:** `<region>/LocationArn` (AWS::DataSync::LocationNFS)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `LocationArn` | location_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the NFS location. |
| `LocationUri` | location_uri | `string` | computed |  | The URL of the NFS location that was described. |
| `MountOptions` | mount_options | `map` | optional, computed, provider-chosen |  | The NFS mount options that DataSync can use to mount your NFS share. |
| `OnPremConfig` | on_prem_config | `map` | required |  | Contains a list of Amazon Resource Names (ARNs) of agents that are used to connect an NFS server. |
| `ServerHostname` | server_hostname | `string` | optional, computed, provider-chosen, write-only |  | The name of the NFS server. This value is the IP address or DNS name of the NFS server. |
| `Subdirectory` |  | `string` | optional, computed, provider-chosen, write-only |  | The subdirectory in the NFS file system that is used to read data from the NFS source location or write data to the NFS destination. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
