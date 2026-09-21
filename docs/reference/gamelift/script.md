# aws.script

**CloudFormation type:** `AWS::GameLift::Script`

The AWS::GameLift::Script resource creates a new script record for your Realtime Servers script. Realtime scripts are JavaScript that provide configuration settings and optional custom game logic for your game. The script is deployed when you create a Realtime Servers fleet to host your game sessions. Script logic is executed during an active game session.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::GameLift::Script)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift script resource and uniquely identifies it. ARNs are unique across all Regions. In a GameLift script ARN, the resource ID matches the Id value. |
| `CreationTime` | creation_time | `string` | computed |  | A time stamp indicating when this data object was created. Format is a number expressed in Unix time as milliseconds (for example "1469498468.057"). |
| `Id` |  | `string` | computed |  | A unique identifier for the Realtime script |
| `Name` |  | `string` | optional, computed, provider-chosen |  | A descriptive label that is associated with a script. Script names do not need to be unique. |
| `NodeJsVersion` | node_js_version | `string` | optional, computed, provider-chosen, replaces on change |  | The Node.js version used for execution of the Realtime script. |
| `SizeOnDisk` | size_on_disk | `integer` | computed |  | The file size of the uploaded Realtime script, expressed in bytes. When files are uploaded from an S3 location, this value remains at "0". |
| `StorageLocation` | storage_location | `map` | required |  | The location of the Amazon S3 bucket where a zipped file containing your Realtime scripts is stored. The storage location must specify the Amazon S3 bucket name, the zip file name (the "key"), and a role ARN that allows Amazon GameLift to access the Amazon S3 storage location. The S3 bucket must be in the same Region where you want to create a new script. By default, Amazon GameLift uploads the latest version of the zip file; if you have S3 object versioning turned on, you can use the ObjectVersion parameter to specify an earlier version. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `Version` |  | `string` | optional, computed, provider-chosen |  | The version that is associated with a script. Version strings do not need to be unique. |

Supports update: yes

Discovery: supported
