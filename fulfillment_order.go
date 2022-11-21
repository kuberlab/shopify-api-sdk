package goshopify

import (
	"fmt"
	"time"
)

const fulfillmentOrderBasePath = "fulfillment_orders"

// FulfillmentOrderService is an interface for interfacing with the fulfillment order
// of the Shopify API.
// https://shopify.dev/api/admin-rest/2022-10/resources/fulfillmentorder
type FulfillmentOrderService interface {
	List(int64) ([]FulfillmentOrder, error)
}

type FulfillmentOrder struct {
	ID                       int64                      `json:"id,omitempty"`
	ShopID                   int64                      `json:"shop_id,omitempty"`
	OrderID                  int64                      `json:"order_id,omitempty"`
	AssignedLocationID       int64                      `json:"assigned_location_id,omitempty"`
	FulfillmentServiceHandle string                     `json:"fulfillment_service_handle,omitempty"`
	RequestStatus            string                     `json:"request_status,omitempty"`
	Status                   string                     `json:"status,omitempty"`
	SupportedActions         []string                   `json:"supported_actions,omitempty"`
	Destination              *Address                   `json:"destination,omitempty"`
	LineItems                []FulfillmentOrderLineItem `json:"line_items,omitempty"`
	FulfillAt                *time.Time                 `json:"fulfill_at,omitempty"`
	FulfillBy                *time.Time                 `json:"fulfill_by,omitempty"`
	InternationalDuties      InternationalDutie         `json:"international_duties,omitempty"`
	FulfillmentHolds         []FulfillmentHold          `json:"fulfillment_holds,omitempty"`
	DeliveryMethod           *DeliveryMethod            `json:"delivery_method,omitempty"`
	AssignedLocation         *AssignedLocation          `json:"assigned_location,omitempty"`
	MerchantRequests         []MerchantRequest          `json:"merchant_requests,omitempty"`
}

type FulfillmentOrderLineItem struct {
	ID                  int64 `json:"id,omitempty"`
	ShopID              int64 `json:"shop_id,omitempty"`
	FulfillmentOrderID  int64 `json:"fulfillment_order_id,omitempty"`
	Quantity            int   `json:"quantity,omitempty"`
	LineItemID          int64 `json:"line_item_id,omitempty"`
	InventoryItemID     int64 `json:"inventory_item_id,omitempty"`
	FulfillableQuantity int   `json:"fulfillable_quantity,omitempty"`
	VariantID           int64 `json:"variant_id,omitempty"`
}

type InternationalDutie struct {
	Incoterm string `json:"incoterm,omitempty"`
}

type FulfillmentHold struct {
	Reason      string `json:"reason,omitempty"`
	ReasonNotes string `json:"reason_notes,omitempty"`
}

type DeliveryMethod struct {
	ID                  int64      `json:"id,omitempty"`
	MethodType          string     `json:"method_type,omitempty"`
	MinDeliveryDateTime *time.Time `json:"min_delivery_date_time,omitempty"`
	MaxDeliveryDateTime *time.Time `json:"max_delivery_date_time,omitempty"`
}

type AssignedLocation struct {
	Address1    string `json:"address1,omitempty"`
	Address2    string `json:"address2,omitempty"`
	City        string `json:"city,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	LocationID  int64  `json:"location_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Province    string `json:"province,omitempty"`
	Zip         string `json:"zip,omitempty"`
}

type MerchantRequest struct {
	Message        string          `json:"message,omitempty"`
	RequestOptions *RequestOptions `json:"request_options,omitempty"`
	Kind           string          `json:"kind,omitempty"`
}
type RequestOptions struct {
	ShippingMethod string     `json:"shipping_method,omitempty"`
	Note           string     `json:"note,omitempty"`
	Date           *time.Time `json:"date,omitempty"`
}

type FulfillmentOrderResource struct {
	FulfillmentOrder *FulfillmentOrder `json:"fulfillment_order,omitempty"`
}

type FulfillmentOrdersResource struct {
	FulfillmentOrders []FulfillmentOrder `json:"fulfillment_orders,omitempty"`
}

// FulfillmentOrderServiceOp handles communication with the FulfillmentOrder
// related methods of the Shopify API
type FulfillmentOrderServiceOp struct {
	client *Client
}

// List Receive a list of all FulfillmentServiceData
func (s *FulfillmentOrderServiceOp) List(orderID int64) ([]FulfillmentOrder, error) {
	path := fmt.Sprintf("%s/%d/%s.json", ordersResourceName, orderID, fulfillmentOrderBasePath)
	resource := new(FulfillmentOrdersResource)
	err := s.client.Get(path, resource, nil)
	return resource.FulfillmentOrders, err
}
