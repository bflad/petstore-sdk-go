// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/bflad/petstore-sdk/internal/utils"
	"time"
)

// OrderStatus - Order Status
type OrderStatus string

const (
	OrderStatusPlaced    OrderStatus = "placed"
	OrderStatusApproved  OrderStatus = "approved"
	OrderStatusDelivered OrderStatus = "delivered"
)

func (e OrderStatus) ToPointer() *OrderStatus {
	return &e
}
func (e *OrderStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "placed":
		fallthrough
	case "approved":
		fallthrough
	case "delivered":
		*e = OrderStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OrderStatus: %v", v)
	}
}

type Order struct {
	ID       *int64     `json:"id,omitempty"`
	PetID    *int64     `json:"petId,omitempty"`
	Quantity *int       `json:"quantity,omitempty"`
	ShipDate *time.Time `json:"shipDate,omitempty"`
	// Order Status
	Status   *OrderStatus `json:"status,omitempty"`
	Complete *bool        `json:"complete,omitempty"`
}

func (o Order) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *Order) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Order) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Order) GetPetID() *int64 {
	if o == nil {
		return nil
	}
	return o.PetID
}

func (o *Order) GetQuantity() *int {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *Order) GetShipDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.ShipDate
}

func (o *Order) GetStatus() *OrderStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Order) GetComplete() *bool {
	if o == nil {
		return nil
	}
	return o.Complete
}
