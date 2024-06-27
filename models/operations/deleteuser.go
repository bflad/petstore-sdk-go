// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/bflad/petstore-sdk/models/components"
)

type DeleteUserRequest struct {
	// The name that needs to be deleted
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

func (o *DeleteUserRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type DeleteUserResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// User deleted
	User *components.User
}

func (o *DeleteUserResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteUserResponse) GetUser() *components.User {
	if o == nil {
		return nil
	}
	return o.User
}
