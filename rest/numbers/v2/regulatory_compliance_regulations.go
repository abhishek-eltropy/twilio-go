/*
 * Twilio - Numbers
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

// Fetch specific Regulation Instance.
func (c *ApiService) FetchRegulation(Sid string) (*NumbersV2RegulatoryComplianceRegulation, error) {
	path := "/v2/RegulatoryCompliance/Regulations/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2RegulatoryComplianceRegulation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRegulation'
type ListRegulationParams struct {
	// The type of End User the regulation requires - can be `individual` or `business`.
	EndUserType *string `json:"EndUserType,omitempty"`
	// The ISO country code of the phone number's country.
	IsoCountry *string `json:"IsoCountry,omitempty"`
	// The type of phone number that the regulatory requiremnt is restricting.
	NumberType *string `json:"NumberType,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListRegulationParams) SetEndUserType(EndUserType string) *ListRegulationParams {
	params.EndUserType = &EndUserType
	return params
}
func (params *ListRegulationParams) SetIsoCountry(IsoCountry string) *ListRegulationParams {
	params.IsoCountry = &IsoCountry
	return params
}
func (params *ListRegulationParams) SetNumberType(NumberType string) *ListRegulationParams {
	params.NumberType = &NumberType
	return params
}
func (params *ListRegulationParams) SetPageSize(PageSize int) *ListRegulationParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Regulations.
func (c *ApiService) ListRegulation(params *ListRegulationParams) (*ListRegulationResponse, error) {
	path := "/v2/RegulatoryCompliance/Regulations"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.EndUserType != nil {
		data.Set("EndUserType", *params.EndUserType)
	}
	if params != nil && params.IsoCountry != nil {
		data.Set("IsoCountry", *params.IsoCountry)
	}
	if params != nil && params.NumberType != nil {
		data.Set("NumberType", *params.NumberType)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRegulationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
