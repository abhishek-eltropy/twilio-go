/*
 * Twilio - Events
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
	"net/url"

	"strings"
)

// Create a new Sink Test Event for the given Sink.
func (c *ApiService) CreateSinkTest(Sid string) (*EventsV1SinkTest, error) {
	path := "/v1/Sinks/{Sid}/Test"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1SinkTest{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
