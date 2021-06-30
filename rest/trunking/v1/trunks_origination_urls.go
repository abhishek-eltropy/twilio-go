/*
 * Twilio - Trunking
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

// Optional parameters for the method 'CreateOriginationUrl'
type CreateOriginationUrlParams struct {
	// Whether the URL is enabled. The default is `true`.
	Enabled *bool `json:"Enabled,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The relative importance of the URI. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important URI.
	Priority *int `json:"Priority,omitempty"`
	// The SIP address you want Twilio to route your Origination calls to. This must be a `sip:` schema.
	SipUrl *string `json:"SipUrl,omitempty"`
	// The value that determines the relative share of the load the URI should receive compared to other URIs with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. URLs with higher values receive more load than those with lower ones with the same priority.
	Weight *int `json:"Weight,omitempty"`
}

func (params *CreateOriginationUrlParams) SetEnabled(Enabled bool) *CreateOriginationUrlParams {
	params.Enabled = &Enabled
	return params
}
func (params *CreateOriginationUrlParams) SetFriendlyName(FriendlyName string) *CreateOriginationUrlParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateOriginationUrlParams) SetPriority(Priority int) *CreateOriginationUrlParams {
	params.Priority = &Priority
	return params
}
func (params *CreateOriginationUrlParams) SetSipUrl(SipUrl string) *CreateOriginationUrlParams {
	params.SipUrl = &SipUrl
	return params
}
func (params *CreateOriginationUrlParams) SetWeight(Weight int) *CreateOriginationUrlParams {
	params.Weight = &Weight
	return params
}

func (c *ApiService) CreateOriginationUrl(TrunkSid string, params *CreateOriginationUrlParams) (*TrunkingV1TrunkOriginationUrl, error) {
	path := "/v1/Trunks/{TrunkSid}/OriginationUrls"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Enabled != nil {
		data.Set("Enabled", fmt.Sprint(*params.Enabled))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Priority != nil {
		data.Set("Priority", fmt.Sprint(*params.Priority))
	}
	if params != nil && params.SipUrl != nil {
		data.Set("SipUrl", *params.SipUrl)
	}
	if params != nil && params.Weight != nil {
		data.Set("Weight", fmt.Sprint(*params.Weight))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1TrunkOriginationUrl{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteOriginationUrl(TrunkSid string, Sid string) error {
	path := "/v1/Trunks/{TrunkSid}/OriginationUrls/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchOriginationUrl(TrunkSid string, Sid string) (*TrunkingV1TrunkOriginationUrl, error) {
	path := "/v1/Trunks/{TrunkSid}/OriginationUrls/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1TrunkOriginationUrl{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListOriginationUrl'
type ListOriginationUrlParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListOriginationUrlParams) SetPageSize(PageSize int) *ListOriginationUrlParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListOriginationUrl(TrunkSid string, params *ListOriginationUrlParams) (*ListOriginationUrlResponse, error) {
	path := "/v1/Trunks/{TrunkSid}/OriginationUrls"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

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

	ps := &ListOriginationUrlResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateOriginationUrl'
type UpdateOriginationUrlParams struct {
	// Whether the URL is enabled. The default is `true`.
	Enabled *bool `json:"Enabled,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The relative importance of the URI. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important URI.
	Priority *int `json:"Priority,omitempty"`
	// The SIP address you want Twilio to route your Origination calls to. This must be a `sip:` schema. `sips` is NOT supported.
	SipUrl *string `json:"SipUrl,omitempty"`
	// The value that determines the relative share of the load the URI should receive compared to other URIs with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. URLs with higher values receive more load than those with lower ones with the same priority.
	Weight *int `json:"Weight,omitempty"`
}

func (params *UpdateOriginationUrlParams) SetEnabled(Enabled bool) *UpdateOriginationUrlParams {
	params.Enabled = &Enabled
	return params
}
func (params *UpdateOriginationUrlParams) SetFriendlyName(FriendlyName string) *UpdateOriginationUrlParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateOriginationUrlParams) SetPriority(Priority int) *UpdateOriginationUrlParams {
	params.Priority = &Priority
	return params
}
func (params *UpdateOriginationUrlParams) SetSipUrl(SipUrl string) *UpdateOriginationUrlParams {
	params.SipUrl = &SipUrl
	return params
}
func (params *UpdateOriginationUrlParams) SetWeight(Weight int) *UpdateOriginationUrlParams {
	params.Weight = &Weight
	return params
}

func (c *ApiService) UpdateOriginationUrl(TrunkSid string, Sid string, params *UpdateOriginationUrlParams) (*TrunkingV1TrunkOriginationUrl, error) {
	path := "/v1/Trunks/{TrunkSid}/OriginationUrls/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Enabled != nil {
		data.Set("Enabled", fmt.Sprint(*params.Enabled))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Priority != nil {
		data.Set("Priority", fmt.Sprint(*params.Priority))
	}
	if params != nil && params.SipUrl != nil {
		data.Set("SipUrl", *params.SipUrl)
	}
	if params != nil && params.Weight != nil {
		data.Set("Weight", fmt.Sprint(*params.Weight))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1TrunkOriginationUrl{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
