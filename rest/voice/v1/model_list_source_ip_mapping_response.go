/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSourceIpMappingResponse struct for ListSourceIpMappingResponse
type ListSourceIpMappingResponse struct {
	Meta             ListByocTrunkResponseMeta `json:"meta,omitempty"`
	SourceIpMappings []VoiceV1SourceIpMapping  `json:"source_ip_mappings,omitempty"`
}
