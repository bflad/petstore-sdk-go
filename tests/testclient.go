// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package tests

import (
	petstoresdk "github.com/bflad/petstore-sdk"
	"net/http"
	"time"
)

func createTestHTTPClient(testName string) petstoresdk.HTTPClient {
	return &testClient{
		client:   &http.Client{Timeout: 60 * time.Second},
		testName: testName,
	}
}

type testClient struct {
	client   *http.Client
	testName string
}

var _ petstoresdk.HTTPClient = &testClient{}

func (c *testClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("x-speakeasy-test-name", c.testName)

	return c.client.Do(req)
}
