/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListTrustProductEntityAssignmentResponse struct for ListTrustProductEntityAssignmentResponse
type ListTrustProductEntityAssignmentResponse struct {
	Meta    ListCustomerProfileResponseMeta          `json:"meta,omitempty"`
	Results []TrusthubV1TrustProductEntityAssignment `json:"results,omitempty"`
}
