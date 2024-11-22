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

func TestUser_CreateUser(t *testing.T) {
	s := petstoresdk.New(
		petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.User.CreateUser(ctx, &components.User{
		ID:         petstoresdk.Int64(10),
		Username:   petstoresdk.String("theUser"),
		FirstName:  petstoresdk.String("John"),
		LastName:   petstoresdk.String("James"),
		Email:      petstoresdk.String("john@email.com"),
		Password:   petstoresdk.String("12345"),
		Phone:      petstoresdk.String("12345"),
		UserStatus: petstoresdk.Int(1),
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, &components.User{
		ID:         petstoresdk.Int64(10),
		Username:   petstoresdk.String("theUser"),
		FirstName:  petstoresdk.String("John"),
		LastName:   petstoresdk.String("James"),
		Email:      petstoresdk.String("john@email.com"),
		Password:   petstoresdk.String("12345"),
		Phone:      petstoresdk.String("12345"),
		UserStatus: petstoresdk.Int(1),
	}, res.User)
}

func TestUser_CreateUsersWithListInput(t *testing.T) {
	s := petstoresdk.New(
		petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.User.CreateUsersWithListInput(ctx, []components.User{
		components.User{
			ID:         petstoresdk.Int64(10),
			Username:   petstoresdk.String("theUser"),
			FirstName:  petstoresdk.String("John"),
			LastName:   petstoresdk.String("James"),
			Email:      petstoresdk.String("john@email.com"),
			Password:   petstoresdk.String("12345"),
			Phone:      petstoresdk.String("12345"),
			UserStatus: petstoresdk.Int(1),
		},
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, &components.User{
		ID:         petstoresdk.Int64(10),
		Username:   petstoresdk.String("theUser"),
		FirstName:  petstoresdk.String("John"),
		LastName:   petstoresdk.String("James"),
		Email:      petstoresdk.String("john@email.com"),
		Password:   petstoresdk.String("12345"),
		Phone:      petstoresdk.String("12345"),
		UserStatus: petstoresdk.Int(1),
	}, res.User)
}

func TestUser_LoginUser(t *testing.T) {
	s := petstoresdk.New(
		petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.User.LoginUser(ctx, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, petstoresdk.String("<value>"), res.String)
}

func TestUser_LogoutUser(t *testing.T) {
	s := petstoresdk.New(
		petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.User.LogoutUser(ctx)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestUser_GetUserByName(t *testing.T) {
	s := petstoresdk.New(
		petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.User.GetUserByName(ctx, "Zachery_Lubowitz15")
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, &components.User{
		ID:         petstoresdk.Int64(10),
		Username:   petstoresdk.String("theUser"),
		FirstName:  petstoresdk.String("John"),
		LastName:   petstoresdk.String("James"),
		Email:      petstoresdk.String("john@email.com"),
		Password:   petstoresdk.String("12345"),
		Phone:      petstoresdk.String("12345"),
		UserStatus: petstoresdk.Int(1),
	}, res.User)
}

func TestUser_UpdateUser(t *testing.T) {
	s := petstoresdk.New(
		petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.User.UpdateUser(ctx, "Dandre_Hand41", &components.User{
		ID:         petstoresdk.Int64(10),
		Username:   petstoresdk.String("theUser"),
		FirstName:  petstoresdk.String("John"),
		LastName:   petstoresdk.String("James"),
		Email:      petstoresdk.String("john@email.com"),
		Password:   petstoresdk.String("12345"),
		Phone:      petstoresdk.String("12345"),
		UserStatus: petstoresdk.Int(1),
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestUser_DeleteUser(t *testing.T) {
	s := petstoresdk.New(
		petstoresdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.User.DeleteUser(ctx, "Demetris_Schmitt")
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, &components.User{
		ID:         petstoresdk.Int64(10),
		Username:   petstoresdk.String("theUser"),
		FirstName:  petstoresdk.String("John"),
		LastName:   petstoresdk.String("James"),
		Email:      petstoresdk.String("john@email.com"),
		Password:   petstoresdk.String("12345"),
		Phone:      petstoresdk.String("12345"),
		UserStatus: petstoresdk.Int(1),
	}, res.User)
}
