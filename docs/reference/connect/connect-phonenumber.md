# aws.connect.phonenumber

**CloudFormation type:** `AWS::Connect::PhoneNumber`

Resource Type definition for AWS::Connect::PhoneNumber

Region attribute: `region`

**Import ID:** `<region>/PhoneNumberArn` (AWS::Connect::PhoneNumber)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Address` |  | `string` | computed |  | The phone number e164 address. |
| `CountryCode` | country_code | `string` | optional, computed, provider-chosen, replaces on change |  | The phone number country code. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the phone number. |
| `PhoneNumberArn` | phone_number_arn | `string` | computed |  | The phone number ARN |
| `Prefix` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The phone number prefix. |
| `SourcePhoneNumberArn` | source_phone_number_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.connect.phonenumber.PhoneNumberArn | The source phone number arn. |
| `Tags` |  | `map` | tags map |  | One or more tags. |
| `TargetArn` | target_arn | `string` | required |  | The ARN of the target the phone number is claimed to. |
| `Type` | type_value | `string` | optional, computed, provider-chosen, replaces on change |  | The phone number type |

Supports update: yes

Discovery: supported (parent resource required)
