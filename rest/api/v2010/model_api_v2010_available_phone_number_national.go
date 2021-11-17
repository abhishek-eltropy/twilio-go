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

// ApiV2010AvailablePhoneNumberNational struct for ApiV2010AvailablePhoneNumberNational
type ApiV2010AvailablePhoneNumberNational struct {
	// The type of Address resource the phone number requires
	AddressRequirements *string `json:"address_requirements,omitempty"`
	// Whether the phone number is new to the Twilio platform
	Beta         *bool                                                                            `json:"beta,omitempty"`
	Capabilities *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalCapabilities `json:"capabilities,omitempty"`
	// A formatted version of the phone number
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The ISO country code of this phone number
	IsoCountry *string `json:"iso_country,omitempty"`
	// The LATA of this phone number
	Lata *string `json:"lata,omitempty"`
	// The latitude of this phone number's location
	Latitude *float32 `json:"latitude,omitempty"`
	// The locality or city of this phone number's location
	Locality *string `json:"locality,omitempty"`
	// The longitude of this phone number's location
	Longitude *float32 `json:"longitude,omitempty"`
	// The phone number in E.164 format
	PhoneNumber *string `json:"phone_number,omitempty"`
	// The postal or ZIP code of this phone number's location
	PostalCode *string `json:"postal_code,omitempty"`
	// The rate center of this phone number
	RateCenter *string `json:"rate_center,omitempty"`
	// The two-letter state or province abbreviation of this phone number's location
	Region *string `json:"region,omitempty"`
}
