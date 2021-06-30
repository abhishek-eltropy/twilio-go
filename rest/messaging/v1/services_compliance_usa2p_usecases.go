/*
 * Twilio - Messaging
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
	"net/url"

	"strings"
)

func (c *ApiService) FetchUsAppToPersonUsecase(MessagingServiceSid string) (*MessagingV1ServiceUsAppToPersonUsecase, error) {
	path := "/v1/Services/{MessagingServiceSid}/Compliance/Usa2p/Usecases"
	path = strings.Replace(path, "{"+"MessagingServiceSid"+"}", MessagingServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1ServiceUsAppToPersonUsecase{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
