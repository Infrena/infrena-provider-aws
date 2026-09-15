# aws.procurementportalpreference

**CloudFormation type:** `AWS::Invoicing::ProcurementPortalPreference`

Creates and manages a procurement portal preference configuration for e-invoice delivery and purchase order retrieval.

Region attribute: `region`

**Import ID:** `<region>/ProcurementPortalPreferenceArn` (AWS::Invoicing::ProcurementPortalPreference)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AwsAccountId` | aws_account_id | `string` | computed |  | The AWS account ID associated with this procurement portal preference. |
| `BuyerDomain` | buyer_domain | `string` | required, replaces on change |  | The domain identifier for the buyer in the procurement portal. |
| `BuyerIdentifier` | buyer_identifier | `string` | required, replaces on change |  | The unique identifier for the buyer in the procurement portal. |
| `Contacts` |  | `list` | required |  | List of contact information for portal administrators and technical contacts. |
| `CreateDate` | create_date | `string` | computed |  | The date and time when the procurement portal preference was created. |
| `EinvoiceDeliveryEnabled` | einvoice_delivery_enabled | `boolean` | required |  | Indicates whether e-invoice delivery is enabled for this procurement portal preference. |
| `EinvoiceDeliveryPreference` | einvoice_delivery_preference | `map` | optional, computed, provider-chosen |  | Specifies the preferences for e-invoice delivery. |
| `EinvoiceDeliveryPreferenceStatus` | einvoice_delivery_preference_status | `string` | computed |  | The current status of the e-invoice delivery preference. |
| `LastUpdateDate` | last_update_date | `string` | computed |  | The date and time when the procurement portal preference was last updated. |
| `ProcurementPortalInstanceEndpoint` | procurement_portal_instance_endpoint | `string` | optional, computed, provider-chosen |  | The endpoint URL where e-invoices are delivered to the procurement portal. |
| `ProcurementPortalName` | procurement_portal_name | `string` | required, replaces on change |  | The name of the procurement portal. |
| `ProcurementPortalPreferenceArn` | procurement_portal_preference_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the procurement portal preference. |
| `ProcurementPortalSharedSecret` | procurement_portal_shared_secret | `string` | optional, computed, provider-chosen, sensitive, write-only |  | The shared secret or authentication credential used for secure communication with the procurement portal. |
| `PurchaseOrderRetrievalEnabled` | purchase_order_retrieval_enabled | `boolean` | required |  | Indicates whether purchase order retrieval is enabled for this procurement portal preference. |
| `PurchaseOrderRetrievalEndpoint` | purchase_order_retrieval_endpoint | `string` | computed |  | The endpoint URL used for retrieving purchase orders from the procurement portal. |
| `PurchaseOrderRetrievalPreferenceStatus` | purchase_order_retrieval_preference_status | `string` | computed |  | The current status of the purchase order retrieval preference. |
| `Selector` |  | `map` | optional, computed, provider-chosen |  | Specifies criteria for selecting which invoices should be processed. |
| `SupplierDomain` | supplier_domain | `string` | required, replaces on change |  | The domain identifier for the supplier in the procurement portal. |
| `SupplierIdentifier` | supplier_identifier | `string` | required, replaces on change |  | The unique identifier for the supplier in the procurement portal. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags associated with this procurement portal preference. |
| `TestEnvPreference` | test_env_preference | `map` | optional, computed, provider-chosen |  | Configuration settings for the test environment of the procurement portal. |
| `Version` |  | `integer` | computed |  | The version number of the procurement portal preference configuration. |

Supports update: yes

Discovery: supported
