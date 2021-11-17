/*
 * Twilio - Conversations
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
	"time"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateConversationParticipant'
type CreateConversationParticipantParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	// An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
	Attributes *string `json:"Attributes,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversations SDK to communicate. Limited to 256 characters.
	Identity *string `json:"Identity,omitempty"`
	// The address of the participant's device, e.g. a phone or WhatsApp number. Together with the Proxy address, this determines a participant uniquely. This field (with proxy_address) is only null when the participant is interacting from an SDK endpoint (see the 'identity' field).
	MessagingBindingAddress *string `json:"MessagingBinding.Address,omitempty"`
	// The address of the Twilio phone number that is used in Group MMS. Communication mask for the Conversation participant with Identity.
	MessagingBindingProjectedAddress *string `json:"MessagingBinding.ProjectedAddress,omitempty"`
	// The address of the Twilio phone number (or WhatsApp number) that the participant is in contact with. This field, together with participant address, is only null when the participant is interacting from an SDK endpoint (see the 'identity' field).
	MessagingBindingProxyAddress *string `json:"MessagingBinding.ProxyAddress,omitempty"`
	// The SID of a conversation-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the participant.
	RoleSid *string `json:"RoleSid,omitempty"`
}

func (params *CreateConversationParticipantParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *CreateConversationParticipantParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *CreateConversationParticipantParams) SetAttributes(Attributes string) *CreateConversationParticipantParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateConversationParticipantParams) SetDateCreated(DateCreated time.Time) *CreateConversationParticipantParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *CreateConversationParticipantParams) SetDateUpdated(DateUpdated time.Time) *CreateConversationParticipantParams {
	params.DateUpdated = &DateUpdated
	return params
}
func (params *CreateConversationParticipantParams) SetIdentity(Identity string) *CreateConversationParticipantParams {
	params.Identity = &Identity
	return params
}
func (params *CreateConversationParticipantParams) SetMessagingBindingAddress(MessagingBindingAddress string) *CreateConversationParticipantParams {
	params.MessagingBindingAddress = &MessagingBindingAddress
	return params
}
func (params *CreateConversationParticipantParams) SetMessagingBindingProjectedAddress(MessagingBindingProjectedAddress string) *CreateConversationParticipantParams {
	params.MessagingBindingProjectedAddress = &MessagingBindingProjectedAddress
	return params
}
func (params *CreateConversationParticipantParams) SetMessagingBindingProxyAddress(MessagingBindingProxyAddress string) *CreateConversationParticipantParams {
	params.MessagingBindingProxyAddress = &MessagingBindingProxyAddress
	return params
}
func (params *CreateConversationParticipantParams) SetRoleSid(RoleSid string) *CreateConversationParticipantParams {
	params.RoleSid = &RoleSid
	return params
}

// Add a new participant to the conversation
func (c *ApiService) CreateConversationParticipant(ConversationSid string, params *CreateConversationParticipantParams) (*ConversationsV1ConversationParticipant, error) {
	path := "/v1/Conversations/{ConversationSid}/Participants"
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
	}
	if params != nil && params.DateUpdated != nil {
		data.Set("DateUpdated", fmt.Sprint((*params.DateUpdated).Format(time.RFC3339)))
	}
	if params != nil && params.Identity != nil {
		data.Set("Identity", *params.Identity)
	}
	if params != nil && params.MessagingBindingAddress != nil {
		data.Set("MessagingBinding.Address", *params.MessagingBindingAddress)
	}
	if params != nil && params.MessagingBindingProjectedAddress != nil {
		data.Set("MessagingBinding.ProjectedAddress", *params.MessagingBindingProjectedAddress)
	}
	if params != nil && params.MessagingBindingProxyAddress != nil {
		data.Set("MessagingBinding.ProxyAddress", *params.MessagingBindingProxyAddress)
	}
	if params != nil && params.RoleSid != nil {
		data.Set("RoleSid", *params.RoleSid)
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConversationParticipant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteConversationParticipant'
type DeleteConversationParticipantParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
}

func (params *DeleteConversationParticipantParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *DeleteConversationParticipantParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}

// Remove a participant from the conversation
func (c *ApiService) DeleteConversationParticipant(ConversationSid string, Sid string, params *DeleteConversationParticipantParams) error {
	path := "/v1/Conversations/{ConversationSid}/Participants/{Sid}"
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a participant of the conversation
func (c *ApiService) FetchConversationParticipant(ConversationSid string, Sid string) (*ConversationsV1ConversationParticipant, error) {
	path := "/v1/Conversations/{ConversationSid}/Participants/{Sid}"
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConversationParticipant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListConversationParticipant'
type ListConversationParticipantParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListConversationParticipantParams) SetPageSize(PageSize int) *ListConversationParticipantParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListConversationParticipantParams) SetLimit(Limit int) *ListConversationParticipantParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ConversationParticipant records from the API. Request is executed immediately.
func (c *ApiService) PageConversationParticipant(ConversationSid string, params *ListConversationParticipantParams, pageToken, pageNumber string) (*ListConversationParticipantResponse, error) {
	path := "/v1/Conversations/{ConversationSid}/Participants"

	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

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

	ps := &ListConversationParticipantResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ConversationParticipant records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListConversationParticipant(ConversationSid string, params *ListConversationParticipantParams) ([]ConversationsV1ConversationParticipant, error) {
	if params == nil {
		params = &ListConversationParticipantParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageConversationParticipant(ConversationSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ConversationsV1ConversationParticipant

	for response != nil {
		records = append(records, response.Participants...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListConversationParticipantResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListConversationParticipantResponse)
	}

	return records, err
}

// Streams ConversationParticipant records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamConversationParticipant(ConversationSid string, params *ListConversationParticipantParams) (chan ConversationsV1ConversationParticipant, error) {
	if params == nil {
		params = &ListConversationParticipantParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageConversationParticipant(ConversationSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ConversationsV1ConversationParticipant, 1)

	go func() {
		for response != nil {
			for item := range response.Participants {
				channel <- response.Participants[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListConversationParticipantResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListConversationParticipantResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListConversationParticipantResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListConversationParticipantResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateConversationParticipant'
type UpdateConversationParticipantParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	// An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
	Attributes *string `json:"Attributes,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversations SDK to communicate. Limited to 256 characters.
	Identity *string `json:"Identity,omitempty"`
	// Index of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant.
	LastReadMessageIndex *int `json:"LastReadMessageIndex,omitempty"`
	// Timestamp of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant.
	LastReadTimestamp *string `json:"LastReadTimestamp,omitempty"`
	// The address of the Twilio phone number that is used in Group MMS. 'null' value will remove it.
	MessagingBindingProjectedAddress *string `json:"MessagingBinding.ProjectedAddress,omitempty"`
	// The address of the Twilio phone number that the participant is in contact with. 'null' value will remove it.
	MessagingBindingProxyAddress *string `json:"MessagingBinding.ProxyAddress,omitempty"`
	// The SID of a conversation-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the participant.
	RoleSid *string `json:"RoleSid,omitempty"`
}

func (params *UpdateConversationParticipantParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *UpdateConversationParticipantParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *UpdateConversationParticipantParams) SetAttributes(Attributes string) *UpdateConversationParticipantParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateConversationParticipantParams) SetDateCreated(DateCreated time.Time) *UpdateConversationParticipantParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *UpdateConversationParticipantParams) SetDateUpdated(DateUpdated time.Time) *UpdateConversationParticipantParams {
	params.DateUpdated = &DateUpdated
	return params
}
func (params *UpdateConversationParticipantParams) SetIdentity(Identity string) *UpdateConversationParticipantParams {
	params.Identity = &Identity
	return params
}
func (params *UpdateConversationParticipantParams) SetLastReadMessageIndex(LastReadMessageIndex int) *UpdateConversationParticipantParams {
	params.LastReadMessageIndex = &LastReadMessageIndex
	return params
}
func (params *UpdateConversationParticipantParams) SetLastReadTimestamp(LastReadTimestamp string) *UpdateConversationParticipantParams {
	params.LastReadTimestamp = &LastReadTimestamp
	return params
}
func (params *UpdateConversationParticipantParams) SetMessagingBindingProjectedAddress(MessagingBindingProjectedAddress string) *UpdateConversationParticipantParams {
	params.MessagingBindingProjectedAddress = &MessagingBindingProjectedAddress
	return params
}
func (params *UpdateConversationParticipantParams) SetMessagingBindingProxyAddress(MessagingBindingProxyAddress string) *UpdateConversationParticipantParams {
	params.MessagingBindingProxyAddress = &MessagingBindingProxyAddress
	return params
}
func (params *UpdateConversationParticipantParams) SetRoleSid(RoleSid string) *UpdateConversationParticipantParams {
	params.RoleSid = &RoleSid
	return params
}

// Update an existing participant in the conversation
func (c *ApiService) UpdateConversationParticipant(ConversationSid string, Sid string, params *UpdateConversationParticipantParams) (*ConversationsV1ConversationParticipant, error) {
	path := "/v1/Conversations/{ConversationSid}/Participants/{Sid}"
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
	}
	if params != nil && params.DateUpdated != nil {
		data.Set("DateUpdated", fmt.Sprint((*params.DateUpdated).Format(time.RFC3339)))
	}
	if params != nil && params.Identity != nil {
		data.Set("Identity", *params.Identity)
	}
	if params != nil && params.LastReadMessageIndex != nil {
		data.Set("LastReadMessageIndex", fmt.Sprint(*params.LastReadMessageIndex))
	}
	if params != nil && params.LastReadTimestamp != nil {
		data.Set("LastReadTimestamp", *params.LastReadTimestamp)
	}
	if params != nil && params.MessagingBindingProjectedAddress != nil {
		data.Set("MessagingBinding.ProjectedAddress", *params.MessagingBindingProjectedAddress)
	}
	if params != nil && params.MessagingBindingProxyAddress != nil {
		data.Set("MessagingBinding.ProxyAddress", *params.MessagingBindingProxyAddress)
	}
	if params != nil && params.RoleSid != nil {
		data.Set("RoleSid", *params.RoleSid)
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConversationParticipant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
