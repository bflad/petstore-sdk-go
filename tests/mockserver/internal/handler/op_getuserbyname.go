// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package handler

import (
	"mockserver/internal/handler/assert"
	"mockserver/internal/sdk/models/components"
	"mockserver/internal/sdk/types"
	"mockserver/internal/sdk/utils"
	"net/http"
)

func opGetUserByName() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if err := assert.SecurityHeader(req, "api_key", false); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := assert.ParameterPathValue(req, "username", "Zachery_Lubowitz15"); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := assert.HeaderValues(req, "Accept", []string{"application/json"}); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := assert.HeaderExists(req, "User-Agent"); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		respBody := &components.User{
			ID:         types.Int64(10),
			Username:   types.String("theUser"),
			FirstName:  types.String("John"),
			LastName:   types.String("James"),
			Email:      types.String("john@email.com"),
			Password:   types.String("12345"),
			Phone:      types.String("12345"),
			UserStatus: types.Int(1),
		}
		respBodyBytes, err := utils.MarshalJSON(respBody, "", true)

		if err != nil {
			http.Error(
				w,
				"Unable to encode response body as JSON: "+err.Error(),
				http.StatusInternalServerError,
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(respBodyBytes)
	}
}
