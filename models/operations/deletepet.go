// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/bflad/petstore-sdk/models/components"
)

type DeletePetRequest struct {
	APIKey *string `header:"style=simple,explode=false,name=api_key"`
	// Pet id to delete
	PetID int64 `pathParam:"style=simple,explode=false,name=petId"`
}

func (o *DeletePetRequest) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}

func (o *DeletePetRequest) GetPetID() int64 {
	if o == nil {
		return 0
	}
	return o.PetID
}

type DeletePetResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Pet deleted
	Pet *components.Pet
}

func (o *DeletePetResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeletePetResponse) GetPet() *components.Pet {
	if o == nil {
		return nil
	}
	return o.Pet
}
