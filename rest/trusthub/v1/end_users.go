/*
 * Twilio - Trusthub
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

// Optional parameters for the method 'CreateEndUser'
type CreateEndUserParams struct {
	// The set of parameters that are the attributes of the End User resource which are derived End User Types.
	Attributes *map[string]interface{} `json:"Attributes,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The type of end user of the Bundle resource - can be `individual` or `business`.
	Type *string `json:"Type,omitempty"`
}

func (params *CreateEndUserParams) SetAttributes(Attributes map[string]interface{}) *CreateEndUserParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateEndUserParams) SetFriendlyName(FriendlyName string) *CreateEndUserParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateEndUserParams) SetType(Type string) *CreateEndUserParams {
	params.Type = &Type
	return params
}

// Create a new End User.
func (c *ApiService) CreateEndUser(params *CreateEndUserParams) (*TrusthubV1EndUser, error) {
	path := "/v1/EndUsers"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		v, err := json.Marshal(params.Attributes)

		if err != nil {
			return nil, err
		}

		data.Set("Attributes", string(v))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Type != nil {
		data.Set("Type", *params.Type)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1EndUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific End User.
func (c *ApiService) DeleteEndUser(Sid string) error {
	path := "/v1/EndUsers/{Sid}"
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

// Fetch specific End User Instance.
func (c *ApiService) FetchEndUser(Sid string) (*TrusthubV1EndUser, error) {
	path := "/v1/EndUsers/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1EndUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListEndUser'
type ListEndUserParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListEndUserParams) SetPageSize(PageSize int) *ListEndUserParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all End User for an account.
func (c *ApiService) ListEndUser(params *ListEndUserParams) (*ListEndUserResponse, error) {
	path := "/v1/EndUsers"

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

	ps := &ListEndUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateEndUser'
type UpdateEndUserParams struct {
	// The set of parameters that are the attributes of the End User resource which are derived End User Types.
	Attributes *map[string]interface{} `json:"Attributes,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateEndUserParams) SetAttributes(Attributes map[string]interface{}) *UpdateEndUserParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateEndUserParams) SetFriendlyName(FriendlyName string) *UpdateEndUserParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Update an existing End User.
func (c *ApiService) UpdateEndUser(Sid string, params *UpdateEndUserParams) (*TrusthubV1EndUser, error) {
	path := "/v1/EndUsers/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		v, err := json.Marshal(params.Attributes)

		if err != nil {
			return nil, err
		}

		data.Set("Attributes", string(v))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1EndUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
