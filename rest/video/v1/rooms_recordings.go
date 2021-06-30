/*
 * Twilio - Video
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
	"time"
)

func (c *ApiService) DeleteRoomRecording(RoomSid string, Sid string) error {
	path := "/v1/Rooms/{RoomSid}/Recordings/{Sid}"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
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

func (c *ApiService) FetchRoomRecording(RoomSid string, Sid string) (*VideoV1RoomRoomRecording, error) {
	path := "/v1/Rooms/{RoomSid}/Recordings/{Sid}"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1RoomRoomRecording{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRoomRecording'
type ListRoomRecordingParams struct {
	// Read only the recordings with this status. Can be: `processing`, `completed`, or `deleted`.
	Status *string `json:"Status,omitempty"`
	// Read only the recordings that have this `source_sid`.
	SourceSid *string `json:"SourceSid,omitempty"`
	// Read only recordings that started on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) datetime with time zone.
	DateCreatedAfter *time.Time `json:"DateCreatedAfter,omitempty"`
	// Read only Recordings that started before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) datetime with time zone.
	DateCreatedBefore *time.Time `json:"DateCreatedBefore,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListRoomRecordingParams) SetStatus(Status string) *ListRoomRecordingParams {
	params.Status = &Status
	return params
}
func (params *ListRoomRecordingParams) SetSourceSid(SourceSid string) *ListRoomRecordingParams {
	params.SourceSid = &SourceSid
	return params
}
func (params *ListRoomRecordingParams) SetDateCreatedAfter(DateCreatedAfter time.Time) *ListRoomRecordingParams {
	params.DateCreatedAfter = &DateCreatedAfter
	return params
}
func (params *ListRoomRecordingParams) SetDateCreatedBefore(DateCreatedBefore time.Time) *ListRoomRecordingParams {
	params.DateCreatedBefore = &DateCreatedBefore
	return params
}
func (params *ListRoomRecordingParams) SetPageSize(PageSize int) *ListRoomRecordingParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListRoomRecording(RoomSid string, params *ListRoomRecordingParams) (*ListRoomRecordingResponse, error) {
	path := "/v1/Rooms/{RoomSid}/Recordings"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.SourceSid != nil {
		data.Set("SourceSid", *params.SourceSid)
	}
	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreatedAfter", fmt.Sprint((*params.DateCreatedAfter).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedBefore != nil {
		data.Set("DateCreatedBefore", fmt.Sprint((*params.DateCreatedBefore).Format(time.RFC3339)))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRoomRecordingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
