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

// Optional parameters for the method 'CreateUsageTrigger'
type CreateUsageTriggerParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is `POST`.
	CallbackMethod *string `json:"CallbackMethod,omitempty"`
	// The URL we should call using `callback_method` when the trigger fires.
	CallbackUrl *string `json:"CallbackUrl,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The frequency of a recurring UsageTrigger.  Can be: `daily`, `monthly`, or `yearly` for recurring triggers or empty for non-recurring triggers. A trigger will only fire once during each period. Recurring times are in GMT.
	Recurring *string `json:"Recurring,omitempty"`
	// The field in the [UsageRecord](https://www.twilio.com/docs/usage/api/usage-record) resource that should fire the trigger.  Can be: `count`, `usage`, or `price` as described in the [UsageRecords documentation](https://www.twilio.com/docs/usage/api/usage-record#usage-count-price).  The default is `usage`.
	TriggerBy *string `json:"TriggerBy,omitempty"`
	// The usage value at which the trigger should fire.  For convenience, you can use an offset value such as `+30` to specify a trigger_value that is 30 units more than the current usage value. Be sure to urlencode a `+` as `%2B`.
	TriggerValue *string `json:"TriggerValue,omitempty"`
	// The usage category that the trigger should watch.  Use one of the supported [usage categories](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) for this value.
	UsageCategory *string `json:"UsageCategory,omitempty"`
}

func (params *CreateUsageTriggerParams) SetPathAccountSid(PathAccountSid string) *CreateUsageTriggerParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateUsageTriggerParams) SetCallbackMethod(CallbackMethod string) *CreateUsageTriggerParams {
	params.CallbackMethod = &CallbackMethod
	return params
}
func (params *CreateUsageTriggerParams) SetCallbackUrl(CallbackUrl string) *CreateUsageTriggerParams {
	params.CallbackUrl = &CallbackUrl
	return params
}
func (params *CreateUsageTriggerParams) SetFriendlyName(FriendlyName string) *CreateUsageTriggerParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateUsageTriggerParams) SetRecurring(Recurring string) *CreateUsageTriggerParams {
	params.Recurring = &Recurring
	return params
}
func (params *CreateUsageTriggerParams) SetTriggerBy(TriggerBy string) *CreateUsageTriggerParams {
	params.TriggerBy = &TriggerBy
	return params
}
func (params *CreateUsageTriggerParams) SetTriggerValue(TriggerValue string) *CreateUsageTriggerParams {
	params.TriggerValue = &TriggerValue
	return params
}
func (params *CreateUsageTriggerParams) SetUsageCategory(UsageCategory string) *CreateUsageTriggerParams {
	params.UsageCategory = &UsageCategory
	return params
}

// Create a new UsageTrigger
func (c *ApiService) CreateUsageTrigger(params *CreateUsageTriggerParams) (*ApiV2010UsageTrigger, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Usage/Triggers.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CallbackMethod != nil {
		data.Set("CallbackMethod", *params.CallbackMethod)
	}
	if params != nil && params.CallbackUrl != nil {
		data.Set("CallbackUrl", *params.CallbackUrl)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Recurring != nil {
		data.Set("Recurring", *params.Recurring)
	}
	if params != nil && params.TriggerBy != nil {
		data.Set("TriggerBy", *params.TriggerBy)
	}
	if params != nil && params.TriggerValue != nil {
		data.Set("TriggerValue", *params.TriggerValue)
	}
	if params != nil && params.UsageCategory != nil {
		data.Set("UsageCategory", *params.UsageCategory)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010UsageTrigger{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteUsageTrigger'
type DeleteUsageTriggerParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteUsageTriggerParams) SetPathAccountSid(PathAccountSid string) *DeleteUsageTriggerParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

func (c *ApiService) DeleteUsageTrigger(Sid string, params *DeleteUsageTriggerParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Usage/Triggers/{Sid}.json"
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

// Optional parameters for the method 'FetchUsageTrigger'
type FetchUsageTriggerParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchUsageTriggerParams) SetPathAccountSid(PathAccountSid string) *FetchUsageTriggerParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch and instance of a usage-trigger
func (c *ApiService) FetchUsageTrigger(Sid string, params *FetchUsageTriggerParams) (*ApiV2010UsageTrigger, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Usage/Triggers/{Sid}.json"
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

	ps := &ApiV2010UsageTrigger{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListUsageTrigger'
type ListUsageTriggerParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The frequency of recurring UsageTriggers to read. Can be: `daily`, `monthly`, or `yearly` to read recurring UsageTriggers. An empty value or a value of `alltime` reads non-recurring UsageTriggers.
	Recurring *string `json:"Recurring,omitempty"`
	// The trigger field of the UsageTriggers to read.  Can be: `count`, `usage`, or `price` as described in the [UsageRecords documentation](https://www.twilio.com/docs/usage/api/usage-record#usage-count-price).
	TriggerBy *string `json:"TriggerBy,omitempty"`
	// The usage category of the UsageTriggers to read. Must be a supported [usage categories](https://www.twilio.com/docs/usage/api/usage-record#usage-categories).
	UsageCategory *string `json:"UsageCategory,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListUsageTriggerParams) SetPathAccountSid(PathAccountSid string) *ListUsageTriggerParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListUsageTriggerParams) SetRecurring(Recurring string) *ListUsageTriggerParams {
	params.Recurring = &Recurring
	return params
}
func (params *ListUsageTriggerParams) SetTriggerBy(TriggerBy string) *ListUsageTriggerParams {
	params.TriggerBy = &TriggerBy
	return params
}
func (params *ListUsageTriggerParams) SetUsageCategory(UsageCategory string) *ListUsageTriggerParams {
	params.UsageCategory = &UsageCategory
	return params
}
func (params *ListUsageTriggerParams) SetPageSize(PageSize int) *ListUsageTriggerParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListUsageTriggerParams) SetLimit(Limit int) *ListUsageTriggerParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of UsageTrigger records from the API. Request is executed immediately.
func (c *ApiService) PageUsageTrigger(params *ListUsageTriggerParams, pageToken, pageNumber string) (*ListUsageTriggerResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Usage/Triggers.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Recurring != nil {
		data.Set("Recurring", *params.Recurring)
	}
	if params != nil && params.TriggerBy != nil {
		data.Set("TriggerBy", *params.TriggerBy)
	}
	if params != nil && params.UsageCategory != nil {
		data.Set("UsageCategory", *params.UsageCategory)
	}
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

	ps := &ListUsageTriggerResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists UsageTrigger records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListUsageTrigger(params *ListUsageTriggerParams) ([]ApiV2010UsageTrigger, error) {
	if params == nil {
		params = &ListUsageTriggerParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageUsageTrigger(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010UsageTrigger

	for response != nil {
		records = append(records, response.UsageTriggers...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListUsageTriggerResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListUsageTriggerResponse)
	}

	return records, err
}

// Streams UsageTrigger records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamUsageTrigger(params *ListUsageTriggerParams) (chan ApiV2010UsageTrigger, error) {
	if params == nil {
		params = &ListUsageTriggerParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageUsageTrigger(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010UsageTrigger, 1)

	go func() {
		for response != nil {
			for item := range response.UsageTriggers {
				channel <- response.UsageTriggers[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListUsageTriggerResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListUsageTriggerResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListUsageTriggerResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUsageTriggerResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateUsageTrigger'
type UpdateUsageTriggerParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resources to update.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is `POST`.
	CallbackMethod *string `json:"CallbackMethod,omitempty"`
	// The URL we should call using `callback_method` when the trigger fires.
	CallbackUrl *string `json:"CallbackUrl,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateUsageTriggerParams) SetPathAccountSid(PathAccountSid string) *UpdateUsageTriggerParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateUsageTriggerParams) SetCallbackMethod(CallbackMethod string) *UpdateUsageTriggerParams {
	params.CallbackMethod = &CallbackMethod
	return params
}
func (params *UpdateUsageTriggerParams) SetCallbackUrl(CallbackUrl string) *UpdateUsageTriggerParams {
	params.CallbackUrl = &CallbackUrl
	return params
}
func (params *UpdateUsageTriggerParams) SetFriendlyName(FriendlyName string) *UpdateUsageTriggerParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Update an instance of a usage trigger
func (c *ApiService) UpdateUsageTrigger(Sid string, params *UpdateUsageTriggerParams) (*ApiV2010UsageTrigger, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Usage/Triggers/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CallbackMethod != nil {
		data.Set("CallbackMethod", *params.CallbackMethod)
	}
	if params != nil && params.CallbackUrl != nil {
		data.Set("CallbackUrl", *params.CallbackUrl)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010UsageTrigger{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
