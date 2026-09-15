# aws.principalpermissions

**CloudFormation type:** `AWS::LakeFormation::PrincipalPermissions`

The ``AWS::LakeFormation::PrincipalPermissions`` resource represents the permissions that a principal has on a GLUDC resource (such as GLUlong databases or GLUlong tables). When you create a ``PrincipalPermissions`` resource, the permissions are granted via the LFlong``GrantPermissions`` API operation. When you delete a ``PrincipalPermissions`` resource, the permissions on principal-resource pair are revoked via the LFlong``RevokePermissions`` API operation.

Region attribute: `region`

**Import ID:** `<region>/PrincipalIdentifier|ResourceIdentifier` (AWS::LakeFormation::PrincipalPermissions)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Catalog` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The identifier for the GLUDC. By default, the account ID. The GLUDC is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment. |
| `Permissions` |  | `list` | required, replaces on change |  | The permissions granted or revoked. |
| `PermissionsWithGrantOption` | permissions_with_grant_option | `list` | required, replaces on change |  | Indicates the ability to grant permissions (as a subset of permissions granted). |
| `Principal` |  | `map` | required, replaces on change |  | The LFlong principal. |
| `PrincipalIdentifier` | principal_identifier | `string` | computed |  |  |
| `Resource` |  | `map` | required, replaces on change |  | A structure for the resource. |
| `ResourceIdentifier` | resource_identifier | `string` | computed |  |  |

Supports update: no

Discovery: supported (parent resource required)
