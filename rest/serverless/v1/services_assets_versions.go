/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.18.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
)

// Retrieve a specific Asset Version.
func (c *ApiService) FetchAssetVersion(ServiceSid string, AssetSid string, Sid string) (*ServerlessV1ServiceAssetAssetVersion, error) {
	path := "/v1/Services/{ServiceSid}/Assets/{AssetSid}/Versions/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"AssetSid"+"}", AssetSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ServerlessV1ServiceAssetAssetVersion{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListAssetVersion'
type ListAssetVersionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListAssetVersionParams) SetPageSize(PageSize int) *ListAssetVersionParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Asset Versions.
func (c *ApiService) ListAssetVersion(ServiceSid string, AssetSid string, params *ListAssetVersionParams) (*ListAssetVersionResponse, error) {
	path := "/v1/Services/{ServiceSid}/Assets/{AssetSid}/Versions"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"AssetSid"+"}", AssetSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAssetVersionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
