# aws.multiregionendpoint

**CloudFormation type:** `AWS::SES::MultiRegionEndpoint`

Resource Type definition for AWS::SES::MultiRegionEndpoint

Region attribute: `region`

**Import ID:** `<region>/EndpointName` (AWS::SES::MultiRegionEndpoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Details` |  | `map` | required, replaces on change |  | Contains details of a multi-region endpoint (global-endpoint) being created. |
| `EndpointName` | endpoint_name | `string` | required, replaces on change |  | The name of the multi-region endpoint (global-endpoint). |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An Array of objects that define the tags (keys and values) to associate with the multi-region endpoint (global-endpoint). |

Supports update: yes

Discovery: supported
