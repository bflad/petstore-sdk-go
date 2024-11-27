// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package petstoresdk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bflad/petstore-sdk/internal/hooks"
	"github.com/bflad/petstore-sdk/internal/utils"
	"github.com/bflad/petstore-sdk/models/components"
	"github.com/bflad/petstore-sdk/retry"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// Mock API server.
	"http://localhost:18080",
	// A per-environment API.
	"https://{environment}.petstore.io",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	ServerDefaults    []map[string]string
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *retry.Config
	Hooks             *hooks.Hooks
	Timeout           *time.Duration
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], c.ServerDefaults[c.ServerIndex]
}

// PetstoreSDK - Petstore - OpenAPI 3.1: This is a sample Pet Store Server based on the OpenAPI 3.1 specification.
//
// Some useful links:
// - [OpenAPI Reference](https://www.speakeasyapi.dev/openapi)
// - [The Pet Store repository](https://github.com/swagger-api/swagger-petstore)
// - [The source API definition for the Pet Store](https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml)
//
// http://swagger.io - Find out more about Swagger
type PetstoreSDK struct {
	// Everything about your Pets
	//
	// http://swagger.io - Find out more
	Pet *Pet
	// Access to Petstore orders
	//
	// http://swagger.io - Find out more about our store
	Store *Store
	// Operations about user
	User *User

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*PetstoreSDK)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *PetstoreSDK) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *PetstoreSDK) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *PetstoreSDK) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// ServerEnvironment - The environment name. Defaults to the production environment.
type ServerEnvironment string

const (
	ServerEnvironmentProd    ServerEnvironment = "prod"
	ServerEnvironmentStaging ServerEnvironment = "staging"
	ServerEnvironmentDev     ServerEnvironment = "dev"
)

func (e ServerEnvironment) ToPointer() *ServerEnvironment {
	return &e
}
func (e *ServerEnvironment) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "prod":
		fallthrough
	case "staging":
		fallthrough
	case "dev":
		*e = ServerEnvironment(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ServerEnvironment: %v", v)
	}
}

// WithEnvironment allows setting the environment variable for url substitution
func WithEnvironment(environment ServerEnvironment) SDKOption {
	return func(sdk *PetstoreSDK) {
		for idx := range sdk.sdkConfiguration.ServerDefaults {
			if _, ok := sdk.sdkConfiguration.ServerDefaults[idx]["environment"]; !ok {
				continue
			}

			sdk.sdkConfiguration.ServerDefaults[idx]["environment"] = fmt.Sprintf("%v", environment)
		}
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *PetstoreSDK) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(apiKey string) SDKOption {
	return func(sdk *PetstoreSDK) {
		security := components.Security{APIKey: apiKey}
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(&security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (components.Security, error)) SDKOption {
	return func(sdk *PetstoreSDK) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *PetstoreSDK) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *PetstoreSDK) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *PetstoreSDK {
	sdk := &PetstoreSDK{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.0.0",
			SDKVersion:        "0.3.4",
			GenVersion:        "2.467.4",
			UserAgent:         "speakeasy-sdk/go 0.3.4 2.467.4 1.0.0 github.com/bflad/petstore-sdk",
			ServerDefaults: []map[string]string{
				{},
				{
					"environment": "prod",
				},
			},
			Hooks: hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.Pet = newPet(sdk.sdkConfiguration)

	sdk.Store = newStore(sdk.sdkConfiguration)

	sdk.User = newUser(sdk.sdkConfiguration)

	return sdk
}
