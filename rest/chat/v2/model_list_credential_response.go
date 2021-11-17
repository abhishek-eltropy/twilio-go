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

// ListCredentialResponse struct for ListCredentialResponse
type ListCredentialResponse struct {
	Credentials []ChatV2Credential         `json:"credentials,omitempty"`
	Meta        ListCredentialResponseMeta `json:"meta,omitempty"`
}
