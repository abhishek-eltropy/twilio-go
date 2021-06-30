/*
 * Twilio - Autopilot
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

// Optional parameters for the method 'CreateSample'
type CreateSampleParams struct {
	// The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the new sample. For example: `en-US`.
	Language *string `json:"Language,omitempty"`
	// The communication channel from which the new sample was captured. Can be: `voice`, `sms`, `chat`, `alexa`, `google-assistant`, `slack`, or null if not included.
	SourceChannel *string `json:"SourceChannel,omitempty"`
	// The text example of how end users might express the task. The sample can contain [Field tag blocks](https://www.twilio.com/docs/autopilot/api/task-sample#field-tagging).
	TaggedText *string `json:"TaggedText,omitempty"`
}

func (params *CreateSampleParams) SetLanguage(Language string) *CreateSampleParams {
	params.Language = &Language
	return params
}
func (params *CreateSampleParams) SetSourceChannel(SourceChannel string) *CreateSampleParams {
	params.SourceChannel = &SourceChannel
	return params
}
func (params *CreateSampleParams) SetTaggedText(TaggedText string) *CreateSampleParams {
	params.TaggedText = &TaggedText
	return params
}

func (c *ApiService) CreateSample(AssistantSid string, TaskSid string, params *CreateSampleParams) (*AutopilotV1AssistantTaskSample, error) {
	path := "/v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"TaskSid"+"}", TaskSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Language != nil {
		data.Set("Language", *params.Language)
	}
	if params != nil && params.SourceChannel != nil {
		data.Set("SourceChannel", *params.SourceChannel)
	}
	if params != nil && params.TaggedText != nil {
		data.Set("TaggedText", *params.TaggedText)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1AssistantTaskSample{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteSample(AssistantSid string, TaskSid string, Sid string) error {
	path := "/v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"TaskSid"+"}", TaskSid, -1)
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

func (c *ApiService) FetchSample(AssistantSid string, TaskSid string, Sid string) (*AutopilotV1AssistantTaskSample, error) {
	path := "/v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"TaskSid"+"}", TaskSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1AssistantTaskSample{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSample'
type ListSampleParams struct {
	// The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the sample. For example: `en-US`.
	Language *string `json:"Language,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListSampleParams) SetLanguage(Language string) *ListSampleParams {
	params.Language = &Language
	return params
}
func (params *ListSampleParams) SetPageSize(PageSize int) *ListSampleParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListSample(AssistantSid string, TaskSid string, params *ListSampleParams) (*ListSampleResponse, error) {
	path := "/v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"TaskSid"+"}", TaskSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Language != nil {
		data.Set("Language", *params.Language)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSampleResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateSample'
type UpdateSampleParams struct {
	// The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the sample. For example: `en-US`.
	Language *string `json:"Language,omitempty"`
	// The communication channel from which the sample was captured. Can be: `voice`, `sms`, `chat`, `alexa`, `google-assistant`, `slack`, or null if not included.
	SourceChannel *string `json:"SourceChannel,omitempty"`
	// The text example of how end users might express the task. The sample can contain [Field tag blocks](https://www.twilio.com/docs/autopilot/api/task-sample#field-tagging).
	TaggedText *string `json:"TaggedText,omitempty"`
}

func (params *UpdateSampleParams) SetLanguage(Language string) *UpdateSampleParams {
	params.Language = &Language
	return params
}
func (params *UpdateSampleParams) SetSourceChannel(SourceChannel string) *UpdateSampleParams {
	params.SourceChannel = &SourceChannel
	return params
}
func (params *UpdateSampleParams) SetTaggedText(TaggedText string) *UpdateSampleParams {
	params.TaggedText = &TaggedText
	return params
}

func (c *ApiService) UpdateSample(AssistantSid string, TaskSid string, Sid string, params *UpdateSampleParams) (*AutopilotV1AssistantTaskSample, error) {
	path := "/v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"TaskSid"+"}", TaskSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Language != nil {
		data.Set("Language", *params.Language)
	}
	if params != nil && params.SourceChannel != nil {
		data.Set("SourceChannel", *params.SourceChannel)
	}
	if params != nil && params.TaggedText != nil {
		data.Set("TaggedText", *params.TaggedText)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1AssistantTaskSample{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
