// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/bflad/petstore-sdk/internal/utils"
	"github.com/bflad/petstore-sdk/models/components"
)

// Status values that need to be considered for filter
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

type FindPetsByStatusRequest struct {
	// Status values that need to be considered for filter
	Status *Status `default:"available" queryParam:"style=form,explode=true,name=status"`
}

func (f FindPetsByStatusRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FindPetsByStatusRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *FindPetsByStatusRequest) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

type FindPetsByStatusResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// successful operation
	Pets []components.Pet
}

func (o *FindPetsByStatusResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *FindPetsByStatusResponse) GetPets() []components.Pet {
	if o == nil {
		return nil
	}
	return o.Pets
}