/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.18.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSyncMapPermissionResponse struct for ListSyncMapPermissionResponse
type ListSyncMapPermissionResponse struct {
	Meta        ListServiceResponseMeta                 `json:"meta,omitempty"`
	Permissions []SyncV1ServiceSyncMapSyncMapPermission `json:"permissions,omitempty"`
}
