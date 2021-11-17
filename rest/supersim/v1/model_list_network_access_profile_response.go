/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListNetworkAccessProfileResponse struct for ListNetworkAccessProfileResponse
type ListNetworkAccessProfileResponse struct {
	Meta                  ListCommandResponseMeta          `json:"meta,omitempty"`
	NetworkAccessProfiles []SupersimV1NetworkAccessProfile `json:"network_access_profiles,omitempty"`
}
