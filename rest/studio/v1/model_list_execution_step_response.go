/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.18.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListExecutionStepResponse struct for ListExecutionStepResponse
type ListExecutionStepResponse struct {
	Meta  ListFlowResponseMeta                 `json:"meta,omitempty"`
	Steps []StudioV1FlowExecutionExecutionStep `json:"steps,omitempty"`
}
