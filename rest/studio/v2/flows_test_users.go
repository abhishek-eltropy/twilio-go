/*
 * Twilio - Studio
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

// Fetch flow test users
func (c *ApiService) FetchTestUser(Sid string) (*StudioV2FlowTestUser, error) {
	path := "/v2/Flows/{Sid}/TestUsers"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV2FlowTestUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateTestUser'
type UpdateTestUserParams struct {
	// List of test user identities that can test draft versions of the flow.
	TestUsers *[]string `json:"TestUsers,omitempty"`
}

func (params *UpdateTestUserParams) SetTestUsers(TestUsers []string) *UpdateTestUserParams {
	params.TestUsers = &TestUsers
	return params
}

// Update flow test users
func (c *ApiService) UpdateTestUser(Sid string, params *UpdateTestUserParams) (*StudioV2FlowTestUser, error) {
	path := "/v2/Flows/{Sid}/TestUsers"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.TestUsers != nil {
		for _, item := range *params.TestUsers {
			data.Add("TestUsers", item)
		}
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV2FlowTestUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
