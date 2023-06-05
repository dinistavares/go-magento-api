package magento

import (
  "net/http"
)

// Customer service
type OrdersService service

type OrderItems struct {
  Orders []Order `json:"items,omitempty"`
}

type Order struct {
  AdjustmentNegative                      float64                 `json:"adjustment_negative,omitempty"`
  AdjustmentPositive                      float64                 `json:"adjustment_positive,omitempty"`
  AppliedRuleIds                          string                  `json:"applied_rule_ids,omitempty"`
  BaseAdjustmentNegative                  float64                 `json:"base_adjustment_negative,omitempty"`
  BaseAdjustmentPositive                  float64                 `json:"base_adjustment_positive,omitempty"`
  BaseCurrencyCode                        string                  `json:"base_currency_code,omitempty"`
  BaseDiscountAmount                      float64                 `json:"base_discount_amount,omitempty"`
  BaseDiscountCanceled                    float64                 `json:"base_discount_canceled,omitempty"`
  BaseDiscountInvoiced                    float64                 `json:"base_discount_invoiced,omitempty"`
  BaseDiscountRefunded                    float64                 `json:"base_discount_refunded,omitempty"`
  BaseGrandTotal                          float64                 `json:"base_grand_total,omitempty"`
  BaseDiscountTaxCompensationAmount       float64                 `json:"base_discount_tax_compensation_amount,omitempty"`
  BaseDiscountTaxCompensationInvoiced     float64                 `json:"base_discount_tax_compensation_invoiced,omitempty"`
  BaseDiscountTaxCompensationRefunded     float64                 `json:"base_discount_tax_compensation_refunded,omitempty"`
  BaseShippingAmount                      float64                 `json:"base_shipping_amount,omitempty"`
  BaseShippingCanceled                    float64                 `json:"base_shipping_canceled,omitempty"`
  BaseShippingDiscountAmount              float64                 `json:"base_shipping_discount_amount,omitempty"`
  BaseShippingDiscountTaxCompensationAmnt float64                 `json:"base_shipping_discount_tax_compensation_amnt,omitempty"`
  BaseShippingInclTax                     float64                 `json:"base_shipping_incl_tax,omitempty"`
  BaseShippingInvoiced                    float64                 `json:"base_shipping_invoiced,omitempty"`
  BaseShippingRefunded                    float64                 `json:"base_shipping_refunded,omitempty"`
  BaseShippingTaxAmount                   float64                 `json:"base_shipping_tax_amount,omitempty"`
  BaseShippingTaxRefunded                 float64                 `json:"base_shipping_tax_refunded,omitempty"`
  BaseSubtotal                            float64                 `json:"base_subtotal,omitempty"`
  BaseSubtotalCanceled                    float64                 `json:"base_subtotal_canceled,omitempty"`
  BaseSubtotalInclTax                     float64                 `json:"base_subtotal_incl_tax,omitempty"`
  BaseSubtotalInvoiced                    float64                 `json:"base_subtotal_invoiced,omitempty"`
  BaseSubtotalRefunded                    float64                 `json:"base_subtotal_refunded,omitempty"`
  BaseTaxAmount                           float64                 `json:"base_tax_amount,omitempty"`
  BaseTaxCanceled                         float64                 `json:"base_tax_canceled,omitempty"`
  BaseTaxInvoiced                         float64                 `json:"base_tax_invoiced,omitempty"`
  BaseTaxRefunded                         float64                 `json:"base_tax_refunded,omitempty"`
  BaseTotalCanceled                       float64                 `json:"base_total_canceled,omitempty"`
  BaseTotalDue                            float64                 `json:"base_total_due,omitempty"`
  BaseTotalInvoiced                       float64                 `json:"base_total_invoiced,omitempty"`
  BaseTotalInvoicedCost                   float64                 `json:"base_total_invoiced_cost,omitempty"`
  BaseTotalOfflineRefunded                float64                 `json:"base_total_offline_refunded,omitempty"`
  BaseTotalOnlineRefunded                 float64                 `json:"base_total_online_refunded,omitempty"`
  BaseTotalPaid                           float64                 `json:"base_total_paid,omitempty"`
  BaseTotalQtyOrdered                     float64                 `json:"base_total_qty_ordered,omitempty"`
  BaseTotalRefunded                       float64                 `json:"base_total_refunded,omitempty"`
  BaseToGlobalRate                        float64                 `json:"base_to_global_rate,omitempty"`
  BaseToOrderRate                         float64                 `json:"base_to_order_rate,omitempty"`
  BillingAddressID                        float64                 `json:"billing_address_id,omitempty"`
  CanShipPartially                        float64                 `json:"can_ship_partially,omitempty"`
  CanShipPartiallyItem                    float64                 `json:"can_ship_partially_item,omitempty"`
  CouponCode                              string                  `json:"coupon_code,omitempty"`
  CreatedAt                               string                  `json:"created_at,omitempty"`
  CustomerDob                             string                  `json:"customer_dob,omitempty"`
  CustomerEmail                           string                  `json:"customer_email,omitempty"`
  CustomerFirstname                       string                  `json:"customer_firstname,omitempty"`
  CustomerGender                          float64                 `json:"customer_gender,omitempty"`
  CustomerGroupID                         float64                 `json:"customer_group_id,omitempty"`
  CustomerID                              float64                 `json:"customer_id,omitempty"`
  CustomerIsGuest                         float64                 `json:"customer_is_guest,omitempty"`
  CustomerLastname                        string                  `json:"customer_lastname,omitempty"`
  CustomerMiddlename                      string                  `json:"customer_middlename,omitempty"`
  CustomerNote                            string                  `json:"customer_note,omitempty"`
  CustomerNoteNotify                      float64                 `json:"customer_note_notify,omitempty"`
  CustomerPrefix                          string                  `json:"customer_prefix,omitempty"`
  CustomerSuffix                          string                  `json:"customer_suffix,omitempty"`
  CustomerTaxvat                          string                  `json:"customer_taxvat,omitempty"`
  DiscountAmount                          float64                 `json:"discount_amount,omitempty"`
  DiscountCanceled                        float64                 `json:"discount_canceled,omitempty"`
  DiscountDescription                     string                  `json:"discount_description,omitempty"`
  DiscountInvoiced                        float64                 `json:"discount_invoiced,omitempty"`
  DiscountRefunded                        float64                 `json:"discount_refunded,omitempty"`
  EditIncrement                           float64                 `json:"edit_increment,omitempty"`
  EmailSent                               float64                 `json:"email_sent,omitempty"`
  EntityID                                float64                 `json:"entity_id,omitempty"`
  ExtCustomerID                           string                  `json:"ext_customer_id,omitempty"`
  ExtOrderID                              string                  `json:"ext_order_id,omitempty"`
  ForcedShipmentWithInvoice               float64                 `json:"forced_shipment_with_invoice,omitempty"`
  GlobalCurrencyCode                      string                  `json:"global_currency_code,omitempty"`
  GrandTotal                              float64                 `json:"grand_total,omitempty"`
  DiscountTaxCompensationAmount           float64                 `json:"discount_tax_compensation_amount,omitempty"`
  DiscountTaxCompensationInvoiced         float64                 `json:"discount_tax_compensation_invoiced,omitempty"`
  DiscountTaxCompensationRefunded         float64                 `json:"discount_tax_compensation_refunded,omitempty"`
  HoldBeforeState                         string                  `json:"hold_before_state,omitempty"`
  HoldBeforeStatus                        string                  `json:"hold_before_status,omitempty"`
  IncrementID                             string                  `json:"increment_id,omitempty"`
  IsVirtual                               float64                 `json:"is_virtual,omitempty"`
  OrderCurrencyCode                       string                  `json:"order_currency_code,omitempty"`
  OriginalIncrementID                     string                  `json:"original_increment_id,omitempty"`
  PaymentAuthorizationAmount              float64                 `json:"payment_authorization_amount,omitempty"`
  PaymentAuthExpiration                   float64                 `json:"payment_auth_expiration,omitempty"`
  ProtectCode                             string                  `json:"protect_code,omitempty"`
  QuoteAddressID                          float64                 `json:"quote_address_id,omitempty"`
  QuoteID                                 float64                 `json:"quote_id,omitempty"`
  RelationChildID                         string                  `json:"relation_child_id,omitempty"`
  RelationChildRealID                     string                  `json:"relation_child_real_id,omitempty"`
  RelationParentID                        string                  `json:"relation_parent_id,omitempty"`
  RelationParentRealID                    string                  `json:"relation_parent_real_id,omitempty"`
  RemoteIP                                string                  `json:"remote_ip,omitempty"`
  ShippingAmount                          float64                 `json:"shipping_amount,omitempty"`
  ShippingCanceled                        float64                 `json:"shipping_canceled,omitempty"`
  ShippingDescription                     string                  `json:"shipping_description,omitempty"`
  ShippingDiscountAmount                  float64                 `json:"shipping_discount_amount,omitempty"`
  ShippingDiscountTaxCompensationAmount   float64                 `json:"shipping_discount_tax_compensation_amount,omitempty"`
  ShippingInclTax                         float64                 `json:"shipping_incl_tax,omitempty"`
  ShippingInvoiced                        float64                 `json:"shipping_invoiced,omitempty"`
  ShippingRefunded                        float64                 `json:"shipping_refunded,omitempty"`
  ShippingTaxAmount                       float64                 `json:"shipping_tax_amount,omitempty"`
  ShippingTaxRefunded                     float64                 `json:"shipping_tax_refunded,omitempty"`
  State                                   string                  `json:"state,omitempty"`
  Status                                  string                  `json:"status,omitempty"`
  StoreCurrencyCode                       string                  `json:"store_currency_code,omitempty"`
  StoreID                                 float64                 `json:"store_id,omitempty"`
  StoreName                               string                  `json:"store_name,omitempty"`
  StoreToBaseRate                         float64                 `json:"store_to_base_rate,omitempty"`
  StoreToOrderRate                        float64                 `json:"store_to_order_rate,omitempty"`
  Subtotal                                float64                 `json:"subtotal,omitempty"`
  SubtotalCanceled                        float64                 `json:"subtotal_canceled,omitempty"`
  SubtotalInclTax                         float64                 `json:"subtotal_incl_tax,omitempty"`
  SubtotalInvoiced                        float64                 `json:"subtotal_invoiced,omitempty"`
  SubtotalRefunded                        float64                 `json:"subtotal_refunded,omitempty"`
  TaxAmount                               float64                 `json:"tax_amount,omitempty"`
  TaxCanceled                             float64                 `json:"tax_canceled,omitempty"`
  TaxInvoiced                             float64                 `json:"tax_invoiced,omitempty"`
  TaxRefunded                             float64                 `json:"tax_refunded,omitempty"`
  TotalCanceled                           float64                 `json:"total_canceled,omitempty"`
  TotalDue                                float64                 `json:"total_due,omitempty"`
  TotalInvoiced                           float64                 `json:"total_invoiced,omitempty"`
  TotalItemCount                          float64                 `json:"total_item_count,omitempty"`
  TotalOfflineRefunded                    float64                 `json:"total_offline_refunded,omitempty"`
  TotalOnlineRefunded                     float64                 `json:"total_online_refunded,omitempty"`
  TotalPaid                               float64                 `json:"total_paid,omitempty"`
  TotalQtyOrdered                         float64                 `json:"total_qty_ordered,omitempty"`
  TotalRefunded                           float64                 `json:"total_refunded,omitempty"`
  UpdatedAt                               string                  `json:"updated_at,omitempty"`
  Weight                                  float64                 `json:"weight,omitempty"`
  XForwardedFor                           string                  `json:"x_forwarded_for,omitempty"`
  Items                                   *[]OrderItem         `json:"items,omitempty"`
  BillingAddress                          *BillingAddress      `json:"billing_address,omitempty"`
  Payment                                 *Payment             `json:"payment,omitempty"`
  StatusHistories                         *[]StatusHistories   `json:"status_histories,omitempty"`
}

type CustomOptions struct {
  OptionID            string                  `json:"option_id,omitempty"`
  OptionValue         string                  `json:"option_value,omitempty"`
}

type BundleOptions struct {
  OptionID            float64                 `json:"option_id,omitempty"`
  OptionQty           float64                 `json:"option_qty,omitempty"`
  OptionSelections    *[]interface{}          `json:"option_selections,omitempty"`
}

type DownloadableOption struct {
  DownloadableLinks *[]float64 `json:"downloadable_links,omitempty"`
}

type GiftcardItemOption struct {
  GiftcardAmount         string                 `json:"giftcard_amount,omitempty"`
  CustomGiftcardAmount   float64                 `json:"custom_giftcard_amount,omitempty"`
  GiftcardSenderName     string                 `json:"giftcard_sender_name,omitempty"`
  GiftcardRecipientName  string                 `json:"giftcard_recipient_name,omitempty"`
  GiftcardSenderEmail    string                 `json:"giftcard_sender_email,omitempty"`
  GiftcardRecipientEmail string                 `json:"giftcard_recipient_email,omitempty"`
  GiftcardMessage        string                 `json:"giftcard_message,omitempty"`
}

type ConfigurableItemOptions struct {
  OptionID            string                  `json:"option_id,omitempty"`
  OptionValue         float64                 `json:"option_value,omitempty"`
}

type GiftMessage struct {
  GiftMessageID       float64                 `json:"gift_message_id,omitempty"`
  CustomerID          float64                 `json:"customer_id,omitempty"`
  Sender              string                  `json:"sender,omitempty"`
  Recipient           string                  `json:"recipient,omitempty"`
  Message             string                  `json:"message,omitempty"`
}

type OrderItem struct {
  AdditionalData                      string                  `json:"additional_data,omitempty"`
  AmountRefunded                      float64                 `json:"amount_refunded,omitempty"`
  AppliedRuleIds                      string                  `json:"applied_rule_ids,omitempty"`
  BaseAmountRefunded                  float64                 `json:"base_amount_refunded,omitempty"`
  BaseCost                            float64                 `json:"base_cost,omitempty"`
  BaseDiscountAmount                  float64                 `json:"base_discount_amount,omitempty"`
  BaseDiscountInvoiced                float64                 `json:"base_discount_invoiced,omitempty"`
  BaseDiscountRefunded                float64                 `json:"base_discount_refunded,omitempty"`
  BaseDiscountTaxCompensationAmount   float64                 `json:"base_discount_tax_compensation_amount,omitempty"`
  BaseDiscountTaxCompensationInvoiced float64                 `json:"base_discount_tax_compensation_invoiced,omitempty"`
  BaseDiscountTaxCompensationRefunded float64                 `json:"base_discount_tax_compensation_refunded,omitempty"`
  BaseOriginalPrice                   float64                 `json:"base_original_price,omitempty"`
  BasePrice                           float64                 `json:"base_price,omitempty"`
  BasePriceInclTax                    float64                 `json:"base_price_incl_tax,omitempty"`
  BaseRowInvoiced                     float64                 `json:"base_row_invoiced,omitempty"`
  BaseRowTotal                        float64                 `json:"base_row_total,omitempty"`
  BaseRowTotalInclTax                 float64                 `json:"base_row_total_incl_tax,omitempty"`
  BaseTaxAmount                       float64                 `json:"base_tax_amount,omitempty"`
  BaseTaxBeforeDiscount               float64                 `json:"base_tax_before_discount,omitempty"`
  BaseTaxInvoiced                     float64                 `json:"base_tax_invoiced,omitempty"`
  BaseTaxRefunded                     float64                 `json:"base_tax_refunded,omitempty"`
  BaseWeeeTaxAppliedAmount            float64                 `json:"base_weee_tax_applied_amount,omitempty"`
  BaseWeeeTaxAppliedRowAmnt           float64                 `json:"base_weee_tax_applied_row_amnt,omitempty"`
  BaseWeeeTaxDisposition              float64                 `json:"base_weee_tax_disposition,omitempty"`
  BaseWeeeTaxRowDisposition           float64                 `json:"base_weee_tax_row_disposition,omitempty"`
  CreatedAt                           string                  `json:"created_at,omitempty"`
  Description                         string                  `json:"description,omitempty"`
  DiscountAmount                      float64                 `json:"discount_amount,omitempty"`
  DiscountInvoiced                    float64                 `json:"discount_invoiced,omitempty"`
  DiscountPercent                     float64                 `json:"discount_percent,omitempty"`
  DiscountRefunded                    float64                 `json:"discount_refunded,omitempty"`
  EventID                             float64                 `json:"event_id,omitempty"`
  ExtOrderItemID                      string                  `json:"ext_order_item_id,omitempty"`
  FreeShipping                        float64                 `json:"free_shipping,omitempty"`
  GwBasePrice                         float64                 `json:"gw_base_price,omitempty"`
  GwBasePriceInvoiced                 float64                 `json:"gw_base_price_invoiced,omitempty"`
  GwBasePriceRefunded                 float64                 `json:"gw_base_price_refunded,omitempty"`
  GwBaseTaxAmount                     float64                 `json:"gw_base_tax_amount,omitempty"`
  GwBaseTaxAmountInvoiced             float64                 `json:"gw_base_tax_amount_invoiced,omitempty"`
  GwBaseTaxAmountRefunded             float64                 `json:"gw_base_tax_amount_refunded,omitempty"`
  GwID                                float64                 `json:"gw_id,omitempty"`
  GwPrice                             float64                 `json:"gw_price,omitempty"`
  GwPriceInvoiced                     float64                 `json:"gw_price_invoiced,omitempty"`
  GwPriceRefunded                     float64                 `json:"gw_price_refunded,omitempty"`
  GwTaxAmount                         float64                 `json:"gw_tax_amount,omitempty"`
  GwTaxAmountInvoiced                 float64                 `json:"gw_tax_amount_invoiced,omitempty"`
  GwTaxAmountRefunded                 float64                 `json:"gw_tax_amount_refunded,omitempty"`
  DiscountTaxCompensationAmount       float64                 `json:"discount_tax_compensation_amount,omitempty"`
  DiscountTaxCompensationCanceled     float64                 `json:"discount_tax_compensation_canceled,omitempty"`
  DiscountTaxCompensationInvoiced     float64                 `json:"discount_tax_compensation_invoiced,omitempty"`
  DiscountTaxCompensationRefunded     float64                 `json:"discount_tax_compensation_refunded,omitempty"`
  IsQtyDecimal                        float64                 `json:"is_qty_decimal,omitempty"`
  IsVirtual                           float64                 `json:"is_virtual,omitempty"`
  ItemID                              float64                 `json:"item_id,omitempty"`
  LockedDoInvoice                     float64                 `json:"locked_do_invoice,omitempty"`
  LockedDoShip                        float64                 `json:"locked_do_ship,omitempty"`
  Name                                string                  `json:"name,omitempty"`
  NoDiscount                          float64                 `json:"no_discount,omitempty"`
  OrderID                             float64                 `json:"order_id,omitempty"`
  OriginalPrice                       float64                 `json:"original_price,omitempty"`
  ParentItemID                        float64                 `json:"parent_item_id,omitempty"`
  Price                               float64                 `json:"price,omitempty"`
  PriceInclTax                        float64                 `json:"price_incl_tax,omitempty"`
  ProductID                           float64                 `json:"product_id,omitempty"`
  ProductType                         string                  `json:"product_type,omitempty"`
  QtyBackordered                      float64                 `json:"qty_backordered,omitempty"`
  QtyCanceled                         float64                 `json:"qty_canceled,omitempty"`
  QtyInvoiced                         float64                 `json:"qty_invoiced,omitempty"`
  QtyOrdered                          float64                 `json:"qty_ordered,omitempty"`
  QtyRefunded                         float64                 `json:"qty_refunded,omitempty"`
  QtyReturned                         float64                 `json:"qty_returned,omitempty"`
  QtyShipped                          float64                 `json:"qty_shipped,omitempty"`
  QuoteItemID                         float64                 `json:"quote_item_id,omitempty"`
  RowInvoiced                         float64                 `json:"row_invoiced,omitempty"`
  RowTotal                            float64                 `json:"row_total,omitempty"`
  RowTotalInclTax                     float64                 `json:"row_total_incl_tax,omitempty"`
  RowWeight                           float64                 `json:"row_weight,omitempty"`
  Sku                                 string                  `json:"sku,omitempty"`
  StoreID                             float64                 `json:"store_id,omitempty"`
  TaxAmount                           float64                 `json:"tax_amount,omitempty"`
  TaxBeforeDiscount                   float64                 `json:"tax_before_discount,omitempty"`
  TaxCanceled                         float64                 `json:"tax_canceled,omitempty"`
  TaxInvoiced                         float64                 `json:"tax_invoiced,omitempty"`
  TaxPercent                          float64                 `json:"tax_percent,omitempty"`
  TaxRefunded                         float64                 `json:"tax_refunded,omitempty"`
  UpdatedAt                           string                  `json:"updated_at,omitempty"`
  WeeeTaxApplied                      string                  `json:"weee_tax_applied,omitempty"`
  WeeeTaxAppliedAmount                float64                 `json:"weee_tax_applied_amount,omitempty"`
  WeeeTaxAppliedRowAmount             float64                 `json:"weee_tax_applied_row_amount,omitempty"`
  WeeeTaxDisposition                  float64                 `json:"weee_tax_disposition,omitempty"`
  WeeeTaxRowDisposition               float64                 `json:"weee_tax_row_disposition,omitempty"`
  Weight                              float64                 `json:"weight,omitempty"`
  ParentItem                          *interface{}            `json:"parent_item,omitempty"`
}

type BillingAddress struct {
  AddressType         string                  `json:"address_type,omitempty"`
  City                string                  `json:"city,omitempty"`
  Company             string                  `json:"company,omitempty"`
  CountryID           string                  `json:"country_id,omitempty"`
  CustomerAddressID   float64                 `json:"customer_address_id,omitempty"`
  CustomerID          float64                 `json:"customer_id,omitempty"`
  Email               string                  `json:"email,omitempty"`
  EntityID            float64                 `json:"entity_id,omitempty"`
  Fax                 string                  `json:"fax,omitempty"`
  Firstname           string                  `json:"firstname,omitempty"`
  Lastname            string                  `json:"lastname,omitempty"`
  Middlename          string                  `json:"middlename,omitempty"`
  ParentID            float64                 `json:"parent_id,omitempty"`
  Postcode            string                  `json:"postcode,omitempty"`
  Prefix              string                  `json:"prefix,omitempty"`
  Region              string                  `json:"region,omitempty"`
  RegionCode          string                  `json:"region_code,omitempty"`
  RegionID            float64                 `json:"region_id,omitempty"`
  Street              []string                `json:"street,omitempty"`
  Suffix              string                  `json:"suffix,omitempty"`
  Telephone           string                  `json:"telephone,omitempty"`
  VatID               string                  `json:"vat_id,omitempty"`
  VatIsValid          float64                 `json:"vat_is_valid,omitempty"`
  VatRequestDate      string                  `json:"vat_request_date,omitempty"`
  VatRequestID        string                  `json:"vat_request_id,omitempty"`
  VatRequestSuccess   float64                 `json:"vat_request_success,omitempty"`
}

type VaultPaymentToken struct {
  EntityID          float64    `json:"entity_id,omitempty"`
  CustomerID        float64    `json:"customer_id,omitempty"`
  PublicHash        string     `json:"public_hash,omitempty"`
  PaymentMethodCode string     `json:"payment_method_code,omitempty"`
  Type              string     `json:"type,omitempty"`
  CreatedAt         string     `json:"created_at,omitempty"`
  ExpiresAt         string     `json:"expires_at,omitempty"`
  GatewayToken      string     `json:"gateway_token,omitempty"`
  TokenDetails      string     `json:"token_details,omitempty"`
  IsActive          bool       `json:"is_active,omitempty"`
  IsVisible         bool       `json:"is_visible,omitempty"`
}

type Payment struct {
  AccountStatus             string                  `json:"account_status,omitempty"`
  AdditionalData            string                  `json:"additional_data,omitempty"`
  AdditionalInformation     []string                `json:"additional_information,omitempty"`
  AddressStatus             string                  `json:"address_status,omitempty"`
  AmountAuthorized          float64                 `json:"amount_authorized,omitempty"`
  AmountCanceled            float64                 `json:"amount_canceled,omitempty"`
  AmountOrdered             float64                 `json:"amount_ordered,omitempty"`
  AmountPaid                float64                 `json:"amount_paid,omitempty"`
  AmountRefunded            float64                 `json:"amount_refunded,omitempty"`
  AnetTransMethod           string                  `json:"anet_trans_method,omitempty"`
  BaseAmountAuthorized      float64                 `json:"base_amount_authorized,omitempty"`
  BaseAmountCanceled        float64                 `json:"base_amount_canceled,omitempty"`
  BaseAmountOrdered         float64                 `json:"base_amount_ordered,omitempty"`
  BaseAmountPaid            float64                 `json:"base_amount_paid,omitempty"`
  BaseAmountPaidOnline      float64                 `json:"base_amount_paid_online,omitempty"`
  BaseAmountRefunded        float64                 `json:"base_amount_refunded,omitempty"`
  BaseAmountRefundedOnline  float64                 `json:"base_amount_refunded_online,omitempty"`
  BaseShippingAmount        float64                 `json:"base_shipping_amount,omitempty"`
  BaseShippingCaptured      float64                 `json:"base_shipping_captured,omitempty"`
  BaseShippingRefunded      float64                 `json:"base_shipping_refunded,omitempty"`
  CcApproval                string                  `json:"cc_approval,omitempty"`
  CcAvsStatus               string                  `json:"cc_avs_status,omitempty"`
  CcCidStatus               string                  `json:"cc_cid_status,omitempty"`
  CcDebugRequestBody        string                  `json:"cc_debug_request_body,omitempty"`
  CcDebugResponseBody       string                  `json:"cc_debug_response_body,omitempty"`
  CcDebugResponseSerialized string                  `json:"cc_debug_response_serialized,omitempty"`
  CcExpMonth                string                  `json:"cc_exp_month,omitempty"`
  CcExpYear                 string                  `json:"cc_exp_year,omitempty"`
  CcLast4                   string                  `json:"cc_last4,omitempty"`
  CcNumberEnc               string                  `json:"cc_number_enc,omitempty"`
  CcOwner                   string                  `json:"cc_owner,omitempty"`
  CcSecureVerify            string                  `json:"cc_secure_verify,omitempty"`
  CcSsIssue                 string                  `json:"cc_ss_issue,omitempty"`
  CcSsStartMonth            string                  `json:"cc_ss_start_month,omitempty"`
  CcSsStartYear             string                  `json:"cc_ss_start_year,omitempty"`
  CcStatus                  string                  `json:"cc_status,omitempty"`
  CcStatusDescription       string                  `json:"cc_status_description,omitempty"`
  CcTransID                 string                  `json:"cc_trans_id,omitempty"`
  CcType                    string                  `json:"cc_type,omitempty"`
  EcheckAccountName         string                  `json:"echeck_account_name,omitempty"`
  EcheckAccountType         string                  `json:"echeck_account_type,omitempty"`
  EcheckBankName            string                  `json:"echeck_bank_name,omitempty"`
  EcheckRoutingNumber       string                  `json:"echeck_routing_number,omitempty"`
  EcheckType                string                  `json:"echeck_type,omitempty"`
  EntityID                  float64                 `json:"entity_id,omitempty"`
  LastTransID               string                  `json:"last_trans_id,omitempty"`
  Method                    string                  `json:"method,omitempty"`
  ParentID                  float64                 `json:"parent_id,omitempty"`
  PoNumber                  string                  `json:"po_number,omitempty"`
  ProtectionEligibility     string                  `json:"protection_eligibility,omitempty"`
  QuotePaymentID            float64                 `json:"quote_payment_id,omitempty"`
  ShippingAmount            float64                 `json:"shipping_amount,omitempty"`
  ShippingCaptured          float64                 `json:"shipping_captured,omitempty"`
  ShippingRefunded          float64                 `json:"shipping_refunded,omitempty"`
}

type StatusHistories struct {
  Comment             string                  `json:"comment,omitempty"`
  CreatedAt           string                  `json:"created_at,omitempty"`
  EntityID            float64                 `json:"entity_id,omitempty"`
  EntityName          string                  `json:"entity_name,omitempty"`
  IsCustomerNotified  float64                 `json:"is_customer_notified,omitempty"`
  IsVisibleOnFront    float64                 `json:"is_visible_on_front,omitempty"`
  ParentID            float64                 `json:"parent_id,omitempty"`
  Status              string                  `json:"status,omitempty"`
}

type Address struct {
  AddressType         string                  `json:"address_type,omitempty"`
  City                string                  `json:"city,omitempty"`
  Company             string                  `json:"company,omitempty"`
  CountryID           string                  `json:"country_id,omitempty"`
  CustomerAddressID   float64                 `json:"customer_address_id,omitempty"`
  CustomerID          float64                 `json:"customer_id,omitempty"`
  Email               string                  `json:"email,omitempty"`
  EntityID            float64                 `json:"entity_id,omitempty"`
  Fax                 string                  `json:"fax,omitempty"`
  Firstname           string                  `json:"firstname,omitempty"`
  Lastname            string                  `json:"lastname,omitempty"`
  Middlename          string                  `json:"middlename,omitempty"`
  ParentID            float64                 `json:"parent_id,omitempty"`
  Postcode            string                  `json:"postcode,omitempty"`
  Prefix              string                  `json:"prefix,omitempty"`
  Region              string                  `json:"region,omitempty"`
  RegionCode          string                  `json:"region_code,omitempty"`
  RegionID            float64                 `json:"region_id,omitempty"`
  Street              *[]string               `json:"street,omitempty"`
  Suffix              string                  `json:"suffix,omitempty"`
  Telephone           string                  `json:"telephone,omitempty"`
  VatID               string                  `json:"vat_id,omitempty"`
  VatIsValid          float64                 `json:"vat_is_valid,omitempty"`
  VatRequestDate      string                  `json:"vat_request_date,omitempty"`
  VatRequestID        string                  `json:"vat_request_id,omitempty"`
  VatRequestSuccess   float64                 `json:"vat_request_success,omitempty"`
}

type Total struct {
  BaseShippingAmount                      float64                 `json:"base_shipping_amount,omitempty"`
  BaseShippingCanceled                    float64                 `json:"base_shipping_canceled,omitempty"`
  BaseShippingDiscountAmount              float64                 `json:"base_shipping_discount_amount,omitempty"`
  BaseShippingDiscountTaxCompensationAmnt float64                 `json:"base_shipping_discount_tax_compensation_amnt,omitempty"`
  BaseShippingInclTax                     float64                 `json:"base_shipping_incl_tax,omitempty"`
  BaseShippingInvoiced                    float64                 `json:"base_shipping_invoiced,omitempty"`
  BaseShippingRefunded                    float64                 `json:"base_shipping_refunded,omitempty"`
  BaseShippingTaxAmount                   float64                 `json:"base_shipping_tax_amount,omitempty"`
  BaseShippingTaxRefunded                 float64                 `json:"base_shipping_tax_refunded,omitempty"`
  ShippingAmount                          float64                 `json:"shipping_amount,omitempty"`
  ShippingCanceled                        float64                 `json:"shipping_canceled,omitempty"`
  ShippingDiscountAmount                  float64                 `json:"shipping_discount_amount,omitempty"`
  ShippingDiscountTaxCompensationAmount   float64                 `json:"shipping_discount_tax_compensation_amount,omitempty"`
  ShippingInclTax                         float64                 `json:"shipping_incl_tax,omitempty"`
  ShippingInvoiced                        float64                 `json:"shipping_invoiced,omitempty"`
  ShippingRefunded                        float64                 `json:"shipping_refunded,omitempty"`
  ShippingTaxAmount                       float64                 `json:"shipping_tax_amount,omitempty"`
  ShippingTaxRefunded                     float64                 `json:"shipping_tax_refunded,omitempty"`
}

type Shipping struct {
  Address             *Address             `json:"address,omitempty"`
  Method              string               `json:"method,omitempty"`
  Total               *Total               `json:"total,omitempty"`
}

type ShippingAssignments struct {
  Shipping            *Shipping            `json:"shipping,omitempty"`
  Items               *[]OrderItem         `json:"items,omitempty"`
  StockID             float64              `json:"stock_id,omitempty"`
}

type PaymentAdditionalInfo struct {
  Key   string     `json:"key,omitempty"`
  Value string     `json:"value,omitempty"`
}

type CompanyOrderAttributes struct {
  OrderID             float64                 `json:"order_id,omitempty"`
  CompanyID           float64                 `json:"company_id,omitempty"`
  CompanyName         string                  `json:"company_name,omitempty"`
}

type Rates struct {
  Code                string                  `json:"code,omitempty"`
  Title               string                  `json:"title,omitempty"`
  Percent             float64                 `json:"percent,omitempty"`
}

type AppliedTaxes struct {
  Code                string                  `json:"code,omitempty"`
  Title               string                  `json:"title,omitempty"`
  Percent             float64                 `json:"percent,omitempty"`
  Amount              float64                 `json:"amount,omitempty"`
  BaseAmount          float64                 `json:"base_amount,omitempty"`
}

type ItemAppliedTaxes struct {
  Type                string                  `json:"type,omitempty"`
  ItemID              float64                 `json:"item_id,omitempty"`
  AssociatedItemID    float64                 `json:"associated_item_id,omitempty"`
  AppliedTaxes        *[]AppliedTaxes         `json:"applied_taxes,omitempty"`
}

type GiftCards struct {
  ID         float64    `json:"id,omitempty"`
  Code       string    `json:"code,omitempty"`
  Amount     float64    `json:"amount,omitempty"`
  BaseAmount float64    `json:"base_amount,omitempty"`
}

// Get a order by ID. Reference: https://adobe-commerce.redoc.ly/2.4.6-admin/tag/ordersid#operation/GetV1OrdersId
func (service *OrdersService) Get(customerID string   ) (*Customer, *http.Response, error) {
  _url := "/orders/" + customerID
  req, _ := service.client.NewRequest("GET", _url, nil, nil)

  customer := new(Customer)
  response, err := service.client.Do(req, customer)

  if err != nil {
    return nil, response, err
  }

  return customer, response, nil
}

// Search for orders. Reference: https://adobe-commerce.redoc.ly/2.4.6-admin/tag/orders#operation/GetV1Orders
func (service *OrdersService) Search(opts *SearchParams) (*[]Order, *http.Response, error) {
  _url := "/orders"
  req, _ := service.client.NewRequest("GET", _url, opts, nil)

  orderItems := new(OrderItems)
  response, err := service.client.Do(req, orderItems)

  if err != nil {
    return nil, response, err
  }

  return &orderItems.Orders, response, nil
}