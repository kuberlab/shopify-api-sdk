package goshopify

import (
	"fmt"
	"time"
)

const inventoryLevelsBasePath = "inventory_levels"

// InventoryLevelService is an interface for interacting with the
// inventory levels endpoints of the Shopify API
// See https://help.shopify.com/en/api/reference/inventory/inventorylevel
type InventoryLevelService interface {
	List(interface{}) ([]InventoryLevel, error)
	Set(InventoryLevel) (*InventoryLevel, error)
}

// InventoryLevelServiceOp is the default implementation of the InventoryLevelService interface
type InventoryLevelServiceOp struct {
	client *Client
}

// InventoryLevel represents a Shopify inventory level
type InventoryLevel struct {
	InventoryItemID   int64      `json:"inventory_item_id,omitempty"`
	LocationID        int64      `json:"location_id,omitempty"`
	Available         int        `json:"available,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	AdminGraphqlAPIID string     `json:"admin_graphql_api_id,omitempty"`
}

// InventoryLevelResource is used for handling single item requests and responses
type InventoryLevelResource struct {
	InventoryLevel *InventoryLevel `json:"inventory_level"`
}

// InventoryLevelsResource is used for handling multiple item response
type InventoryLevelsResource struct {
	InventoryLevels []InventoryLevel `json:"inventory_levels"`
}

type InventoryLevelOptions struct {
	InventoryItemIDS string    `url:"inventory_item_ids,omitempty"`
	LocationIDS      string    `url:"location_ids,omitempty"`
	Limit            int       `url:"limit,omitempty"`
	UpdatedAtMin     time.Time `url:"updated_at_min,omitempty"`
}

// List Retrieves a list of inventory levels
func (s *InventoryLevelServiceOp) List(options interface{}) ([]InventoryLevel, error) {
	path := fmt.Sprintf("%s.json", inventoryLevelsBasePath)
	resource := new(InventoryLevelsResource)
	err := s.client.Get(path, resource, options)
	return resource.InventoryLevels, err
}

// Set Sets the inventory level for an inventory item at a location
func (s *InventoryLevelServiceOp) Set(inventoryLevel InventoryLevel) (*InventoryLevel, error) {
	path := fmt.Sprintf("%s/set.json", inventoryLevelsBasePath)
	resource := new(InventoryLevelResource)
	err := s.client.Post(path, inventoryLevel, resource)
	return resource.InventoryLevel, err
}
