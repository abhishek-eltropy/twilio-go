/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSubscribedEventResponse struct for ListSubscribedEventResponse
type ListSubscribedEventResponse struct {
	Meta  ListSchemaVersionResponseMeta `json:"meta,omitempty"`
	Types []EventsV1SubscribedEvent     `json:"types,omitempty"`
}
