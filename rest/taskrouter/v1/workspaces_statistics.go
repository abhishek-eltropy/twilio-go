/*
 * Twilio - Taskrouter
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
)

// Optional parameters for the method 'FetchWorkspaceStatistics'
type FetchWorkspaceStatisticsParams struct {
	// Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends.
	Minutes *int `json:"Minutes,omitempty"`
	// Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	StartDate *time.Time `json:"StartDate,omitempty"`
	// Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
	EndDate *time.Time `json:"EndDate,omitempty"`
	// Only calculate statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`.
	TaskChannel *string `json:"TaskChannel,omitempty"`
	// A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed. For example, `5,30` would show splits of Tasks that were canceled or accepted before and after 5 seconds and before and after 30 seconds. This can be used to show short abandoned Tasks or Tasks that failed to meet an SLA.
	SplitByWaitTime *string `json:"SplitByWaitTime,omitempty"`
}

func (params *FetchWorkspaceStatisticsParams) SetMinutes(Minutes int) *FetchWorkspaceStatisticsParams {
	params.Minutes = &Minutes
	return params
}
func (params *FetchWorkspaceStatisticsParams) SetStartDate(StartDate time.Time) *FetchWorkspaceStatisticsParams {
	params.StartDate = &StartDate
	return params
}
func (params *FetchWorkspaceStatisticsParams) SetEndDate(EndDate time.Time) *FetchWorkspaceStatisticsParams {
	params.EndDate = &EndDate
	return params
}
func (params *FetchWorkspaceStatisticsParams) SetTaskChannel(TaskChannel string) *FetchWorkspaceStatisticsParams {
	params.TaskChannel = &TaskChannel
	return params
}
func (params *FetchWorkspaceStatisticsParams) SetSplitByWaitTime(SplitByWaitTime string) *FetchWorkspaceStatisticsParams {
	params.SplitByWaitTime = &SplitByWaitTime
	return params
}

func (c *ApiService) FetchWorkspaceStatistics(WorkspaceSid string, params *FetchWorkspaceStatisticsParams) (*TaskrouterV1WorkspaceStatistics, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Statistics"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Minutes != nil {
		data.Set("Minutes", fmt.Sprint(*params.Minutes))
	}
	if params != nil && params.StartDate != nil {
		data.Set("StartDate", fmt.Sprint((*params.StartDate).Format(time.RFC3339)))
	}
	if params != nil && params.EndDate != nil {
		data.Set("EndDate", fmt.Sprint((*params.EndDate).Format(time.RFC3339)))
	}
	if params != nil && params.TaskChannel != nil {
		data.Set("TaskChannel", *params.TaskChannel)
	}
	if params != nil && params.SplitByWaitTime != nil {
		data.Set("SplitByWaitTime", *params.SplitByWaitTime)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1WorkspaceStatistics{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
