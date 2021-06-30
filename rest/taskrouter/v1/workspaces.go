/*
 * Twilio - Taskrouter
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

// Optional parameters for the method 'CreateWorkspace'
type CreateWorkspaceParams struct {
	// The URL we should call when an event occurs. If provided, the Workspace will publish events to this URL, for example, to collect data for reporting. See [Workspace Events](https://www.twilio.com/docs/taskrouter/api/event) for more information.
	EventCallbackUrl *string `json:"EventCallbackUrl,omitempty"`
	// The list of Workspace events for which to call event_callback_url. For example, if `EventsFilter=task.created, task.canceled, worker.activity.update`, then TaskRouter will call event_callback_url only when a task is created, canceled, or a Worker activity is updated.
	EventsFilter *string `json:"EventsFilter,omitempty"`
	// A descriptive string that you create to describe the Workspace resource. It can be up to 64 characters long. For example: `Customer Support` or `2014 Election Campaign`.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether to enable multi-tasking. Can be: `true` to enable multi-tasking, or `false` to disable it. The default is `false`. Multi-tasking allows Workers to handle multiple Tasks simultaneously. When enabled (`true`), each Worker can receive parallel reservations up to the per-channel maximums defined in the Workers section. Otherwise, each Worker will only receive a new reservation when the previous task is completed. Learn more at [Multitasking](https://www.twilio.com/docs/taskrouter/multitasking).
	MultiTaskEnabled *bool `json:"MultiTaskEnabled,omitempty"`
	// The type of TaskQueue to prioritize when Workers are receiving Tasks from both types of TaskQueues. Can be: `LIFO` or `FIFO` and the default is `FIFO`. For more information, see [Queue Ordering](https://www.twilio.com/docs/taskrouter/queue-ordering-last-first-out-lifo).
	PrioritizeQueueOrder *string `json:"PrioritizeQueueOrder,omitempty"`
	// An available template name. Can be: `NONE` or `FIFO` and the default is `NONE`. Pre-configures the Workspace with the Workflow and Activities specified in the template. `NONE` will create a Workspace with only a set of default activities. `FIFO` will configure TaskRouter with a set of default activities and a single TaskQueue for first-in, first-out distribution, which can be useful when you are getting started with TaskRouter.
	Template *string `json:"Template,omitempty"`
}

func (params *CreateWorkspaceParams) SetEventCallbackUrl(EventCallbackUrl string) *CreateWorkspaceParams {
	params.EventCallbackUrl = &EventCallbackUrl
	return params
}
func (params *CreateWorkspaceParams) SetEventsFilter(EventsFilter string) *CreateWorkspaceParams {
	params.EventsFilter = &EventsFilter
	return params
}
func (params *CreateWorkspaceParams) SetFriendlyName(FriendlyName string) *CreateWorkspaceParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateWorkspaceParams) SetMultiTaskEnabled(MultiTaskEnabled bool) *CreateWorkspaceParams {
	params.MultiTaskEnabled = &MultiTaskEnabled
	return params
}
func (params *CreateWorkspaceParams) SetPrioritizeQueueOrder(PrioritizeQueueOrder string) *CreateWorkspaceParams {
	params.PrioritizeQueueOrder = &PrioritizeQueueOrder
	return params
}
func (params *CreateWorkspaceParams) SetTemplate(Template string) *CreateWorkspaceParams {
	params.Template = &Template
	return params
}

func (c *ApiService) CreateWorkspace(params *CreateWorkspaceParams) (*TaskrouterV1Workspace, error) {
	path := "/v1/Workspaces"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.EventCallbackUrl != nil {
		data.Set("EventCallbackUrl", *params.EventCallbackUrl)
	}
	if params != nil && params.EventsFilter != nil {
		data.Set("EventsFilter", *params.EventsFilter)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.MultiTaskEnabled != nil {
		data.Set("MultiTaskEnabled", fmt.Sprint(*params.MultiTaskEnabled))
	}
	if params != nil && params.PrioritizeQueueOrder != nil {
		data.Set("PrioritizeQueueOrder", *params.PrioritizeQueueOrder)
	}
	if params != nil && params.Template != nil {
		data.Set("Template", *params.Template)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1Workspace{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteWorkspace(Sid string) error {
	path := "/v1/Workspaces/{Sid}"
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

func (c *ApiService) FetchWorkspace(Sid string) (*TaskrouterV1Workspace, error) {
	path := "/v1/Workspaces/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1Workspace{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListWorkspace'
type ListWorkspaceParams struct {
	// The `friendly_name` of the Workspace resources to read. For example `Customer Support` or `2014 Election Campaign`.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListWorkspaceParams) SetFriendlyName(FriendlyName string) *ListWorkspaceParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListWorkspaceParams) SetPageSize(PageSize int) *ListWorkspaceParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListWorkspace(params *ListWorkspaceParams) (*ListWorkspaceResponse, error) {
	path := "/v1/Workspaces"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListWorkspaceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateWorkspace'
type UpdateWorkspaceParams struct {
	// The SID of the Activity that will be used when new Workers are created in the Workspace.
	DefaultActivitySid *string `json:"DefaultActivitySid,omitempty"`
	// The URL we should call when an event occurs. See [Workspace Events](https://www.twilio.com/docs/taskrouter/api/event) for more information.
	EventCallbackUrl *string `json:"EventCallbackUrl,omitempty"`
	// The list of Workspace events for which to call event_callback_url. For example if `EventsFilter=task.created,task.canceled,worker.activity.update`, then TaskRouter will call event_callback_url only when a task is created, canceled, or a Worker activity is updated.
	EventsFilter *string `json:"EventsFilter,omitempty"`
	// A descriptive string that you create to describe the Workspace resource. For example: `Sales Call Center` or `Customer Support Team`.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether to enable multi-tasking. Can be: `true` to enable multi-tasking, or `false` to disable it. The default is `false`. Multi-tasking allows Workers to handle multiple Tasks simultaneously. When enabled (`true`), each Worker can receive parallel reservations up to the per-channel maximums defined in the Workers section. Otherwise, each Worker will only receive a new reservation when the previous task is completed. Learn more at [Multitasking](https://www.twilio.com/docs/taskrouter/multitasking).
	MultiTaskEnabled *bool `json:"MultiTaskEnabled,omitempty"`
	// The type of TaskQueue to prioritize when Workers are receiving Tasks from both types of TaskQueues. Can be: `LIFO` or `FIFO` and the default is `FIFO`. For more information, see [Queue Ordering](https://www.twilio.com/docs/taskrouter/queue-ordering-last-first-out-lifo).
	PrioritizeQueueOrder *string `json:"PrioritizeQueueOrder,omitempty"`
	// The SID of the Activity that will be assigned to a Worker when a Task reservation times out without a response.
	TimeoutActivitySid *string `json:"TimeoutActivitySid,omitempty"`
}

func (params *UpdateWorkspaceParams) SetDefaultActivitySid(DefaultActivitySid string) *UpdateWorkspaceParams {
	params.DefaultActivitySid = &DefaultActivitySid
	return params
}
func (params *UpdateWorkspaceParams) SetEventCallbackUrl(EventCallbackUrl string) *UpdateWorkspaceParams {
	params.EventCallbackUrl = &EventCallbackUrl
	return params
}
func (params *UpdateWorkspaceParams) SetEventsFilter(EventsFilter string) *UpdateWorkspaceParams {
	params.EventsFilter = &EventsFilter
	return params
}
func (params *UpdateWorkspaceParams) SetFriendlyName(FriendlyName string) *UpdateWorkspaceParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateWorkspaceParams) SetMultiTaskEnabled(MultiTaskEnabled bool) *UpdateWorkspaceParams {
	params.MultiTaskEnabled = &MultiTaskEnabled
	return params
}
func (params *UpdateWorkspaceParams) SetPrioritizeQueueOrder(PrioritizeQueueOrder string) *UpdateWorkspaceParams {
	params.PrioritizeQueueOrder = &PrioritizeQueueOrder
	return params
}
func (params *UpdateWorkspaceParams) SetTimeoutActivitySid(TimeoutActivitySid string) *UpdateWorkspaceParams {
	params.TimeoutActivitySid = &TimeoutActivitySid
	return params
}

func (c *ApiService) UpdateWorkspace(Sid string, params *UpdateWorkspaceParams) (*TaskrouterV1Workspace, error) {
	path := "/v1/Workspaces/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.DefaultActivitySid != nil {
		data.Set("DefaultActivitySid", *params.DefaultActivitySid)
	}
	if params != nil && params.EventCallbackUrl != nil {
		data.Set("EventCallbackUrl", *params.EventCallbackUrl)
	}
	if params != nil && params.EventsFilter != nil {
		data.Set("EventsFilter", *params.EventsFilter)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.MultiTaskEnabled != nil {
		data.Set("MultiTaskEnabled", fmt.Sprint(*params.MultiTaskEnabled))
	}
	if params != nil && params.PrioritizeQueueOrder != nil {
		data.Set("PrioritizeQueueOrder", *params.PrioritizeQueueOrder)
	}
	if params != nil && params.TimeoutActivitySid != nil {
		data.Set("TimeoutActivitySid", *params.TimeoutActivitySid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1Workspace{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
