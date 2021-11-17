/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.1
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

// Fetch a specific Country.
func (c *ApiService) FetchTrunkingCountry(IsoCountry string) (*PricingV2TrunkingCountryInstance, error) {
	path := "/v2/Trunking/Countries/{IsoCountry}"
	path = strings.Replace(path, "{"+"IsoCountry"+"}", IsoCountry, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &PricingV2TrunkingCountryInstance{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListTrunkingCountry'
type ListTrunkingCountryParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListTrunkingCountryParams) SetPageSize(PageSize int) *ListTrunkingCountryParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListTrunkingCountryParams) SetLimit(Limit int) *ListTrunkingCountryParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of TrunkingCountry records from the API. Request is executed immediately.
func (c *ApiService) PageTrunkingCountry(params *ListTrunkingCountryParams, pageToken, pageNumber string) (*ListTrunkingCountryResponse, error) {
	path := "/v2/Trunking/Countries"

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

	ps := &ListTrunkingCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists TrunkingCountry records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListTrunkingCountry(params *ListTrunkingCountryParams) ([]PricingV2TrunkingCountry, error) {
	if params == nil {
		params = &ListTrunkingCountryParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageTrunkingCountry(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []PricingV2TrunkingCountry

	for response != nil {
		records = append(records, response.Countries...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListTrunkingCountryResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListTrunkingCountryResponse)
	}

	return records, err
}

// Streams TrunkingCountry records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamTrunkingCountry(params *ListTrunkingCountryParams) (chan PricingV2TrunkingCountry, error) {
	if params == nil {
		params = &ListTrunkingCountryParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageTrunkingCountry(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan PricingV2TrunkingCountry, 1)

	go func() {
		for response != nil {
			for item := range response.Countries {
				channel <- response.Countries[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListTrunkingCountryResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListTrunkingCountryResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListTrunkingCountryResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTrunkingCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
