/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// ChatV2Channel struct for ChatV2Channel
type ChatV2Channel struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The JSON string that stores application-specific data
	Attributes *string `json:"attributes,omitempty"`
	// The identity of the User that created the channel
	CreatedBy *string `json:"created_by,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// Absolute URLs to access the Members, Messages , Invites and, if it exists, the last Message for the Channel
	Links *map[string]interface{} `json:"links,omitempty"`
	// The number of Members in the Channel
	MembersCount *int `json:"members_count,omitempty"`
	// The number of Messages that have been passed in the Channel
	MessagesCount *int `json:"messages_count,omitempty"`
	// The SID of the Service that the resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The visibility of the channel. Can be: `public` or `private`
	Type *string `json:"type,omitempty"`
	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"unique_name,omitempty"`
	// The absolute URL of the Channel resource
	Url *string `json:"url,omitempty"`
}
