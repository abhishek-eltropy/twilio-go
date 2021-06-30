/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.18.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListWorkerResponse struct for ListWorkerResponse
type ListWorkerResponse struct {
	Meta    ListWorkspaceResponseMeta     `json:"meta,omitempty"`
	Workers []TaskrouterV1WorkspaceWorker `json:"workers,omitempty"`
}
