// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type DeleteTenancyV1TenantServiceGroupsTsgIDSecurity struct {
	Bearer string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type DeleteTenancyV1TenantServiceGroupsTsgIDRequest struct {
	// A unique identifier for the tenant service group.
	//
	TsgID string `pathParam:"style=simple,explode=false,name=tsg_id"`
}

type DeleteTenancyV1TenantServiceGroupsTsgIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response.
	TenantServiceGroup *shared.TenantServiceGroup
}