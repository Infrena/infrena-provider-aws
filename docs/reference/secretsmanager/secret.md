# aws.secret

**CloudFormation type:** `AWS::SecretsManager::Secret`

Creates a new secret. A *secret* can be a password, a set of credentials such as a user name and password, an OAuth token, or other secret information that you store in an encrypted form in Secrets Manager.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::SecretsManager::Secret)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the secret. |
| `GenerateSecretString` | generate_secret_string | `map` | optional, computed, provider-chosen, write-only |  | Generates a random password. We recommend that you specify the maximum length and include every character type that the system you are generating a password for can support. |
| `Id` |  | `string` | computed |  |  |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen |  | The ARN, key ID, or alias of the KMS key that Secrets Manager uses to encrypt the secret value in the secret. An alias is always prefixed by ``alias/``, for example ``alias/aws/secretsmanager``. For more information, see [About aliases](https://docs.aws.amazon.com/kms/latest/developerguide/alias-about.html). |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the new secret. |
| `ReplicaRegions` | replica_regions | `list` | optional, computed, provider-chosen |  | A custom type that specifies a ``Region`` and the ``KmsKeyId`` for a replica secret. |
| `SecretString` | secret_string | `string` | optional, computed, provider-chosen, sensitive, write-only |  | The text to encrypt and store in the secret. We recommend you use a JSON structure of key/value pairs for your secret value. To generate a random password, use ``GenerateSecretString`` instead. If you omit both ``GenerateSecretString`` and ``SecretString``, you create an empty secret. When you make a change to this property, a new secret version is created. |
| `Tags` |  | `map` | tags map |  | A list of tags to attach to the secret. Each tag is a key and value pair of strings in a JSON text string, for example: |
| `Type` | type_value | `string` | optional, computed, provider-chosen |  | The exact string that identifies the third-party partner that holds the external secret. For more information, see [Managed external secret partners](https://docs.aws.amazon.com/secretsmanager/latest/userguide/mes-partners.html). |

Supports update: yes

Discovery: supported
