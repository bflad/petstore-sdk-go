// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// Status - pet status in the store
type Status string

const (
	StatusAvailable Status = "available"
	StatusPending   Status = "pending"
	StatusSold      Status = "sold"
)

func (e Status) ToPointer() *Status {
	return &e
}
func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "available":
		fallthrough
	case "pending":
		fallthrough
	case "sold":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

type Pet struct {
	ID        *int64    `json:"id,omitempty"`
	Name      string    `json:"name"`
	Category  *Category `json:"category,omitempty"`
	PhotoUrls []string  `json:"photoUrls"`
	Tags      []Tag     `json:"tags,omitempty"`
	// pet status in the store
	Status *Status `json:"status,omitempty"`
}

func (o *Pet) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Pet) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Pet) GetCategory() *Category {
	if o == nil {
		return nil
	}
	return o.Category
}

func (o *Pet) GetPhotoUrls() []string {
	if o == nil {
		return []string{}
	}
	return o.PhotoUrls
}

func (o *Pet) GetTags() []Tag {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Pet) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}
