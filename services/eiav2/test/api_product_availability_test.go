/*
Equinix Internet Access API

Testing ProductAvailabilityApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package eiav2

import (
	"context"
	"testing"

	openapiclient "github.com/equinix/equinix-sdk-go/services/eiav2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_eiav2_ProductAvailabilityApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProductAvailabilityApiService GetIbxs", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProductAvailabilityApi.GetIbxs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}