// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/bflad/petstore-sdk/models/components"
)

type CreateUsersWithListInputResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful operation
	User *components.User
}

func (o *CreateUsersWithListInputResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateUsersWithListInputResponse) GetUser() *components.User {
	if o == nil {
		return nil
	}
	return o.User
}
