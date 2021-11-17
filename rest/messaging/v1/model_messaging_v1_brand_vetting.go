/*
 * Twilio - Messaging
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

// MessagingV1BrandVetting struct for MessagingV1BrandVetting
type MessagingV1BrandVetting struct {
	// The SID of the Account that created the vetting
	AccountSid *string `json:"account_sid,omitempty"`
	// A2P BrandRegistration Sid
	BrandSid *string `json:"brand_sid,omitempty"`
	// SID for third-party vetting record
	BrandVettingSid *string `json:"brand_vetting_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The absolute URL of the Brand Vetting
	Url *string `json:"url,omitempty"`
	// The type of vetting
	VettingClass *string `json:"vetting_class,omitempty"`
	// The unique ID of the vetting
	VettingId *string `json:"vetting_id,omitempty"`
	// Third-party provider that has conducted the vetting
	VettingProvider *string `json:"vetting_provider,omitempty"`
	// Status of vetting attempt
	VettingStatus *string `json:"vetting_status,omitempty"`
}
