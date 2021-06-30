/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.18.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListItemAssignmentResponse struct for ListItemAssignmentResponse
type ListItemAssignmentResponse struct {
	Meta    ListBundleResponseMeta                              `json:"meta,omitempty"`
	Results []NumbersV2RegulatoryComplianceBundleItemAssignment `json:"results,omitempty"`
}
