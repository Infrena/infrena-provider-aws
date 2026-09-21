# aws.build

**CloudFormation type:** `AWS::GameLift::Build`

Resource Type definition for AWS::GameLift::Build

Region attribute: `region`

**Import ID:** `<region>/BuildId` (AWS::GameLift::Build)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BuildArn` | build_arn | `string` | computed |  | The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift build resource and uniquely identifies it. ARNs are unique across all Regions. In a GameLift build ARN, the resource ID matches the BuildId value. |
| `BuildId` | build_id | `string` | computed |  | A unique identifier for a build to be deployed on the new fleet. If you are deploying the fleet with a custom game build, you must specify this property. The build must have been successfully uploaded to Amazon GameLift and be in a READY status. This fleet setting cannot be changed once the fleet is created. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | A descriptive label that is associated with a build. Build names do not need to be unique. |
| `OperatingSystem` | operating_system | `string` | optional, computed, provider-chosen, replaces on change |  | The operating system that the game server binaries are built to run on. This value determines the type of fleet resources that you can use for this build. If your game build contains multiple executables, they all must run on the same operating system. If an operating system is not specified when creating a build, Amazon GameLift uses the default value (WINDOWS_2012). This value cannot be changed later. |
| `ServerSdkVersion` | server_sdk_version | `string` | optional, computed, provider-chosen, replaces on change |  | A server SDK version you used when integrating your game server build with Amazon GameLift. By default Amazon GameLift sets this value to 4.0.2. |
| `StorageLocation` | storage_location | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | Information indicating where your game build files are stored. Use this parameter only when creating a build with files stored in an Amazon S3 bucket that you own. The storage location must specify an Amazon S3 bucket name and key. The location must also specify a role ARN that you set up to allow Amazon GameLift to access your Amazon S3 bucket. The S3 bucket and your new build must be in the same Region. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `Version` |  | `string` | optional, computed, provider-chosen |  | Version information that is associated with this build. Version strings do not need to be unique. |

Supports update: yes

Discovery: supported
