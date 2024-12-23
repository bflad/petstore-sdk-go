// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GetPetByIDRequest struct {
	// ID of pet to return
	PetID int64 `pathParam:"style=simple,explode=false,name=petId"`
}

func (o *GetPetByIDRequest) GetPetID() int64 {
	if o == nil {
		return 0
	}
	return o.PetID
}

type GetPetByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// successful operation
	Pet *components.Pet
}

func (o *GetPetByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetPetByIDResponse) GetPet() *components.Pet {
	if o == nil {
		return nil
	}
	return o.Pet
}
