/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateShortCode'
type CreateShortCodeParams struct {
	// The SID of a Twilio [ShortCode](https://www.twilio.com/docs/sms/api/short-code) resource that represents the short code you would like to assign to your Proxy Service.
	Sid *string `json:"Sid,omitempty"`
}

func (params *CreateShortCodeParams) SetSid(Sid string) *CreateShortCodeParams {
	params.Sid = &Sid
	return params
}

// Add a Short Code to the Proxy Number Pool for the Service.
func (c *ApiService) CreateShortCode(ServiceSid string, params *CreateShortCodeParams) (*ProxyV1ShortCode, error) {
	path := "/v1/Services/{ServiceSid}/ShortCodes"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Sid != nil {
		data.Set("Sid", *params.Sid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1ShortCode{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Short Code from a Service.
func (c *ApiService) DeleteShortCode(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/ShortCodes/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
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

// Fetch a specific Short Code.
func (c *ApiService) FetchShortCode(ServiceSid string, Sid string) (*ProxyV1ShortCode, error) {
	path := "/v1/Services/{ServiceSid}/ShortCodes/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1ShortCode{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListShortCode'
type ListShortCodeParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListShortCodeParams) SetPageSize(PageSize int) *ListShortCodeParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListShortCodeParams) SetLimit(Limit int) *ListShortCodeParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ShortCode records from the API. Request is executed immediately.
func (c *ApiService) PageShortCode(ServiceSid string, params *ListShortCodeParams, pageToken, pageNumber string) (*ListShortCodeResponse, error) {
	path := "/v1/Services/{ServiceSid}/ShortCodes"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListShortCodeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ShortCode records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListShortCode(ServiceSid string, params *ListShortCodeParams) ([]ProxyV1ShortCode, error) {
	if params == nil {
		params = &ListShortCodeParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageShortCode(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ProxyV1ShortCode

	for response != nil {
		records = append(records, response.ShortCodes...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListShortCodeResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListShortCodeResponse)
	}

	return records, err
}

// Streams ShortCode records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamShortCode(ServiceSid string, params *ListShortCodeParams) (chan ProxyV1ShortCode, error) {
	if params == nil {
		params = &ListShortCodeParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageShortCode(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ProxyV1ShortCode, 1)

	go func() {
		for response != nil {
			for item := range response.ShortCodes {
				channel <- response.ShortCodes[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListShortCodeResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListShortCodeResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListShortCodeResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListShortCodeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateShortCode'
type UpdateShortCodeParams struct {
	// Whether the short code should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information.
	IsReserved *bool `json:"IsReserved,omitempty"`
}

func (params *UpdateShortCodeParams) SetIsReserved(IsReserved bool) *UpdateShortCodeParams {
	params.IsReserved = &IsReserved
	return params
}

// Update a specific Short Code.
func (c *ApiService) UpdateShortCode(ServiceSid string, Sid string, params *UpdateShortCodeParams) (*ProxyV1ShortCode, error) {
	path := "/v1/Services/{ServiceSid}/ShortCodes/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.IsReserved != nil {
		data.Set("IsReserved", fmt.Sprint(*params.IsReserved))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1ShortCode{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
