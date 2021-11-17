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

// Optional parameters for the method 'DeleteConnectApp'
type DeleteConnectAppParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteConnectAppParams) SetPathAccountSid(PathAccountSid string) *DeleteConnectAppParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete an instance of a connect-app
func (c *ApiService) DeleteConnectApp(Sid string, params *DeleteConnectAppParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/ConnectApps/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
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

// Optional parameters for the method 'FetchConnectApp'
type FetchConnectAppParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchConnectAppParams) SetPathAccountSid(PathAccountSid string) *FetchConnectAppParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an instance of a connect-app
func (c *ApiService) FetchConnectApp(Sid string, params *FetchConnectAppParams) (*ApiV2010ConnectApp, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/ConnectApps/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010ConnectApp{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListConnectApp'
type ListConnectAppParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListConnectAppParams) SetPathAccountSid(PathAccountSid string) *ListConnectAppParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListConnectAppParams) SetPageSize(PageSize int) *ListConnectAppParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListConnectAppParams) SetLimit(Limit int) *ListConnectAppParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ConnectApp records from the API. Request is executed immediately.
func (c *ApiService) PageConnectApp(params *ListConnectAppParams, pageToken, pageNumber string) (*ListConnectAppResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/ConnectApps.json"

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

	ps := &ListConnectAppResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ConnectApp records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListConnectApp(params *ListConnectAppParams) ([]ApiV2010ConnectApp, error) {
	if params == nil {
		params = &ListConnectAppParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageConnectApp(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010ConnectApp

	for response != nil {
		records = append(records, response.ConnectApps...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListConnectAppResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListConnectAppResponse)
	}

	return records, err
}

// Streams ConnectApp records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamConnectApp(params *ListConnectAppParams) (chan ApiV2010ConnectApp, error) {
	if params == nil {
		params = &ListConnectAppParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageConnectApp(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010ConnectApp, 1)

	go func() {
		for response != nil {
			for item := range response.ConnectApps {
				channel <- response.ConnectApps[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListConnectAppResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListConnectAppResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListConnectAppResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListConnectAppResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateConnectApp'
type UpdateConnectAppParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resources to update.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The URL to redirect the user to after we authenticate the user and obtain authorization to access the Connect App.
	AuthorizeRedirectUrl *string `json:"AuthorizeRedirectUrl,omitempty"`
	// The company name to set for the Connect App.
	CompanyName *string `json:"CompanyName,omitempty"`
	// The HTTP method to use when calling `deauthorize_callback_url`.
	DeauthorizeCallbackMethod *string `json:"DeauthorizeCallbackMethod,omitempty"`
	// The URL to call using the `deauthorize_callback_method` to de-authorize the Connect App.
	DeauthorizeCallbackUrl *string `json:"DeauthorizeCallbackUrl,omitempty"`
	// A description of the Connect App.
	Description *string `json:"Description,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// A public URL where users can obtain more information about this Connect App.
	HomepageUrl *string `json:"HomepageUrl,omitempty"`
	// A comma-separated list of the permissions you will request from the users of this ConnectApp.  Can include: `get-all` and `post-all`.
	Permissions *[]string `json:"Permissions,omitempty"`
}

func (params *UpdateConnectAppParams) SetPathAccountSid(PathAccountSid string) *UpdateConnectAppParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateConnectAppParams) SetAuthorizeRedirectUrl(AuthorizeRedirectUrl string) *UpdateConnectAppParams {
	params.AuthorizeRedirectUrl = &AuthorizeRedirectUrl
	return params
}
func (params *UpdateConnectAppParams) SetCompanyName(CompanyName string) *UpdateConnectAppParams {
	params.CompanyName = &CompanyName
	return params
}
func (params *UpdateConnectAppParams) SetDeauthorizeCallbackMethod(DeauthorizeCallbackMethod string) *UpdateConnectAppParams {
	params.DeauthorizeCallbackMethod = &DeauthorizeCallbackMethod
	return params
}
func (params *UpdateConnectAppParams) SetDeauthorizeCallbackUrl(DeauthorizeCallbackUrl string) *UpdateConnectAppParams {
	params.DeauthorizeCallbackUrl = &DeauthorizeCallbackUrl
	return params
}
func (params *UpdateConnectAppParams) SetDescription(Description string) *UpdateConnectAppParams {
	params.Description = &Description
	return params
}
func (params *UpdateConnectAppParams) SetFriendlyName(FriendlyName string) *UpdateConnectAppParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateConnectAppParams) SetHomepageUrl(HomepageUrl string) *UpdateConnectAppParams {
	params.HomepageUrl = &HomepageUrl
	return params
}
func (params *UpdateConnectAppParams) SetPermissions(Permissions []string) *UpdateConnectAppParams {
	params.Permissions = &Permissions
	return params
}

// Update a connect-app with the specified parameters
func (c *ApiService) UpdateConnectApp(Sid string, params *UpdateConnectAppParams) (*ApiV2010ConnectApp, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/ConnectApps/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AuthorizeRedirectUrl != nil {
		data.Set("AuthorizeRedirectUrl", *params.AuthorizeRedirectUrl)
	}
	if params != nil && params.CompanyName != nil {
		data.Set("CompanyName", *params.CompanyName)
	}
	if params != nil && params.DeauthorizeCallbackMethod != nil {
		data.Set("DeauthorizeCallbackMethod", *params.DeauthorizeCallbackMethod)
	}
	if params != nil && params.DeauthorizeCallbackUrl != nil {
		data.Set("DeauthorizeCallbackUrl", *params.DeauthorizeCallbackUrl)
	}
	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.HomepageUrl != nil {
		data.Set("HomepageUrl", *params.HomepageUrl)
	}
	if params != nil && params.Permissions != nil {
		for _, item := range *params.Permissions {
			data.Add("Permissions", item)
		}
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010ConnectApp{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
