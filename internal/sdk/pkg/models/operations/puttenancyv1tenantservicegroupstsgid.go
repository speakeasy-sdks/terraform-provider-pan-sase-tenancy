// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"pan-sase-tenancy/internal/sdk/pkg/models/shared"
)

type PutTenancyV1TenantServiceGroupsTsgIDRequest struct {
	TenantServiceGroupUpdate shared.TenantServiceGroupUpdate `request:"mediaType=application/json"`
	// A unique identifier for the tenant service group.
	//
	TsgID string `pathParam:"style=simple,explode=false,name=tsg_id"`
}

type PutTenancyV1TenantServiceGroupsTsgIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response.
	TenantServiceGroup *shared.TenantServiceGroup
}
