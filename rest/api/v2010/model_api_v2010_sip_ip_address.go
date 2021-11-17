/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010SipIpAddress struct for ApiV2010SipIpAddress
type ApiV2010SipIpAddress struct {
	// The unique id of the Account that is responsible for this resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// An integer representing the length of the CIDR prefix to use with this IP address when accepting traffic. By default the entire IP address is used.
	CidrPrefixLength *int `json:"cidr_prefix_length,omitempty"`
	// The date that this resource was created, given as GMT in RFC 2822 format.
	DateCreated *string `json:"date_created,omitempty"`
	// The date that this resource was last updated, given as GMT in RFC 2822 format.
	DateUpdated *string `json:"date_updated,omitempty"`
	// A human readable descriptive text for this resource, up to 64 characters long.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The unique id of the IpAccessControlList resource that includes this resource.
	IpAccessControlListSid *string `json:"ip_access_control_list_sid,omitempty"`
	// An IP address in dotted decimal notation from which you want to accept traffic. Any SIP requests from this IP address will be allowed by Twilio. IPv4 only supported today.
	IpAddress *string `json:"ip_address,omitempty"`
	// A 34 character string that uniquely identifies this resource.
	Sid *string `json:"sid,omitempty"`
	// The URI for this resource, relative to https://api.twilio.com
	Uri *string `json:"uri,omitempty"`
}
