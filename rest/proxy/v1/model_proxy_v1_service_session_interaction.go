/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.18.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// ProxyV1ServiceSessionInteraction struct for ProxyV1ServiceSessionInteraction
type ProxyV1ServiceSessionInteraction struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// A JSON string that includes the message body of message interactions
	Data *string `json:"data,omitempty"`
	// The ISO 8601 date and time in GMT when the Interaction was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The SID of the inbound Participant resource
	InboundParticipantSid *string `json:"inbound_participant_sid,omitempty"`
	// The SID of the inbound resource
	InboundResourceSid *string `json:"inbound_resource_sid,omitempty"`
	// The inbound resource status of the Interaction
	InboundResourceStatus *string `json:"inbound_resource_status,omitempty"`
	// The inbound resource type
	InboundResourceType *string `json:"inbound_resource_type,omitempty"`
	// The URL of the Twilio inbound resource
	InboundResourceUrl *string `json:"inbound_resource_url,omitempty"`
	// The SID of the outbound Participant
	OutboundParticipantSid *string `json:"outbound_participant_sid,omitempty"`
	// The SID of the outbound resource
	OutboundResourceSid *string `json:"outbound_resource_sid,omitempty"`
	// The outbound resource status of the Interaction
	OutboundResourceStatus *string `json:"outbound_resource_status,omitempty"`
	// The outbound resource type
	OutboundResourceType *string `json:"outbound_resource_type,omitempty"`
	// The URL of the Twilio outbound resource
	OutboundResourceUrl *string `json:"outbound_resource_url,omitempty"`
	// The SID of the resource's parent Service
	ServiceSid *string `json:"service_sid,omitempty"`
	// The SID of the resource's parent Session
	SessionSid *string `json:"session_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The Type of the Interaction
	Type *string `json:"type,omitempty"`
	// The absolute URL of the Interaction resource
	Url *string `json:"url,omitempty"`
}
