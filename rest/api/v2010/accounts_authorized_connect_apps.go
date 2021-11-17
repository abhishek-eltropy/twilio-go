/*
 * Twilio - Api
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

// Optional parameters for the method 'FetchAuthorizedConnectApp'
type FetchAuthorizedConnectAppParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the AuthorizedConnectApp resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchAuthorizedConnectAppParams) SetPathAccountSid(PathAccountSid string) *FetchAuthorizedConnectAppParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an instance of an authorized-connect-app
func (c *ApiService) FetchAuthorizedConnectApp(ConnectAppSid string, params *FetchAuthorizedConnectAppParams) (*ApiV2010AuthorizedConnectApp, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/AuthorizedConnectApps/{ConnectAppSid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ConnectAppSid"+"}", ConnectAppSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AuthorizedConnectApp{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListAuthorizedConnectApp'
type ListAuthorizedConnectAppParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the AuthorizedConnectApp resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListAuthorizedConnectAppParams) SetPathAccountSid(PathAccountSid string) *ListAuthorizedConnectAppParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListAuthorizedConnectAppParams) SetPageSize(PageSize int) *ListAuthorizedConnectAppParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListAuthorizedConnectAppParams) SetLimit(Limit int) *ListAuthorizedConnectAppParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of AuthorizedConnectApp records from the API. Request is executed immediately.
func (c *ApiService) PageAuthorizedConnectApp(params *ListAuthorizedConnectAppParams, pageToken, pageNumber string) (*ListAuthorizedConnectAppResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/AuthorizedConnectApps.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

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

	ps := &ListAuthorizedConnectAppResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists AuthorizedConnectApp records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListAuthorizedConnectApp(params *ListAuthorizedConnectAppParams) ([]ApiV2010AuthorizedConnectApp, error) {
	if params == nil {
		params = &ListAuthorizedConnectAppParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageAuthorizedConnectApp(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010AuthorizedConnectApp

	for response != nil {
		records = append(records, response.AuthorizedConnectApps...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListAuthorizedConnectAppResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListAuthorizedConnectAppResponse)
	}

	return records, err
}

// Streams AuthorizedConnectApp records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamAuthorizedConnectApp(params *ListAuthorizedConnectAppParams) (chan ApiV2010AuthorizedConnectApp, error) {
	if params == nil {
		params = &ListAuthorizedConnectAppParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageAuthorizedConnectApp(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010AuthorizedConnectApp, 1)

	go func() {
		for response != nil {
			for item := range response.AuthorizedConnectApps {
				channel <- response.AuthorizedConnectApps[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListAuthorizedConnectAppResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListAuthorizedConnectAppResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListAuthorizedConnectAppResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAuthorizedConnectAppResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
