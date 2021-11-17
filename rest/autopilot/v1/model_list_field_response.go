/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListFieldResponse struct for ListFieldResponse
type ListFieldResponse struct {
	Fields []AutopilotV1Field        `json:"fields,omitempty"`
	Meta   ListAssistantResponseMeta `json:"meta,omitempty"`
}
