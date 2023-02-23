package graphql

type FulfillmentOrderLineItems struct {
	ID       string `json:"id,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}
type LineItemsByFulfillmentOrder struct {
	FulfillmentOrderID        string                       `json:"fulfillmentOrderId,omitempty"`
	FulfillmentOrderLineItems []*FulfillmentOrderLineItems `json:"fulfillmentOrderLineItems,omitempty"`
}
type TrackingInfo struct {
	Company string   `json:"company,omitempty"`
	Number  string   `json:"number,omitempty"`
	Numbers []string `json:"numbers,omitempty"`
	URL     string   `json:"url,omitempty"`
	Urls    []string `json:"urls,omitempty"`
}
type FulfillmentCreateV2Fulfillment struct {
	LineItemsByFulfillmentOrder []*LineItemsByFulfillmentOrder `json:"lineItemsByFulfillmentOrder,omitempty"`
	NotifyCustomer              bool                           `json:"notifyCustomer,omitempty"`
	TrackingInfo                *TrackingInfo                  `json:"trackingInfo,omitempty"`
}

type FulfillmentCreateV2Request struct {
	Fulfillment *FulfillmentCreateV2Fulfillment `json:"fulfillment,omitempty"`
	Message     string                          `json:"message,omitempty"`
}

type FulfillmentResource struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}
