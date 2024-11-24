// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package tests

import (
	"context"
	petstoresdk "github.com/bflad/petstore-sdk"
	"github.com/bflad/petstore-sdk/models/components"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStore_GetInventory(t *testing.T) {
	s := petstoresdk.New(
		petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Store.GetInventory(ctx)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, map[string]int{
		"key": 373538,
	}, res.Object)
}

func TestStore_PlaceOrder(t *testing.T) {
	s := petstoresdk.New(
		petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Store.PlaceOrder(ctx, &components.Order{
		ID:       petstoresdk.Int64(10),
		PetID:    petstoresdk.Int64(198772),
		Quantity: petstoresdk.Int(7),
		Status:   components.OrderStatusApproved.ToPointer(),
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, &components.Order{
		ID:       petstoresdk.Int64(10),
		PetID:    petstoresdk.Int64(198772),
		Quantity: petstoresdk.Int(7),
		Status:   components.OrderStatusApproved.ToPointer(),
	}, res.Order)
}

func TestStore_GetOrderByID(t *testing.T) {
	s := petstoresdk.New(
		petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Store.GetOrderByID(ctx, 614993)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, &components.Order{
		ID:       petstoresdk.Int64(10),
		PetID:    petstoresdk.Int64(198772),
		Quantity: petstoresdk.Int(7),
		Status:   components.OrderStatusApproved.ToPointer(),
	}, res.Order)
}

func TestStore_DeleteOrder(t *testing.T) {
	s := petstoresdk.New(
		petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Store.DeleteOrder(ctx, 127902)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, &components.Order{
		ID:       petstoresdk.Int64(10),
		PetID:    petstoresdk.Int64(198772),
		Quantity: petstoresdk.Int(7),
		Status:   components.OrderStatusApproved.ToPointer(),
	}, res.Order)
}
