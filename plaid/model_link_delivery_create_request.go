/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.197.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// LinkDeliveryCreateRequest LinkDeliveryCreateRequest defines the request schema for `/link_delivery/create`
type LinkDeliveryCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// A `link_token` from a previous invocation of `/link/token/create` with Link Delivery enabled
	LinkToken string `json:"link_token"`
	DeliveryMethod LinkDeliveryDeliveryMethod `json:"delivery_method"`
	// The email or phone number to be used to delivery the URL of the Link Delivery session
	DeliveryDestination string `json:"delivery_destination"`
}

// NewLinkDeliveryCreateRequest instantiates a new LinkDeliveryCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkDeliveryCreateRequest(linkToken string, deliveryMethod LinkDeliveryDeliveryMethod, deliveryDestination string) *LinkDeliveryCreateRequest {
	this := LinkDeliveryCreateRequest{}
	this.LinkToken = linkToken
	this.DeliveryMethod = deliveryMethod
	this.DeliveryDestination = deliveryDestination
	return &this
}

// NewLinkDeliveryCreateRequestWithDefaults instantiates a new LinkDeliveryCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkDeliveryCreateRequestWithDefaults() *LinkDeliveryCreateRequest {
	this := LinkDeliveryCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *LinkDeliveryCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDeliveryCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *LinkDeliveryCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *LinkDeliveryCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *LinkDeliveryCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDeliveryCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *LinkDeliveryCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *LinkDeliveryCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetLinkToken returns the LinkToken field value
func (o *LinkDeliveryCreateRequest) GetLinkToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkToken
}

// GetLinkTokenOk returns a tuple with the LinkToken field value
// and a boolean to check if the value has been set.
func (o *LinkDeliveryCreateRequest) GetLinkTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LinkToken, true
}

// SetLinkToken sets field value
func (o *LinkDeliveryCreateRequest) SetLinkToken(v string) {
	o.LinkToken = v
}

// GetDeliveryMethod returns the DeliveryMethod field value
func (o *LinkDeliveryCreateRequest) GetDeliveryMethod() LinkDeliveryDeliveryMethod {
	if o == nil {
		var ret LinkDeliveryDeliveryMethod
		return ret
	}

	return o.DeliveryMethod
}

// GetDeliveryMethodOk returns a tuple with the DeliveryMethod field value
// and a boolean to check if the value has been set.
func (o *LinkDeliveryCreateRequest) GetDeliveryMethodOk() (*LinkDeliveryDeliveryMethod, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DeliveryMethod, true
}

// SetDeliveryMethod sets field value
func (o *LinkDeliveryCreateRequest) SetDeliveryMethod(v LinkDeliveryDeliveryMethod) {
	o.DeliveryMethod = v
}

// GetDeliveryDestination returns the DeliveryDestination field value
func (o *LinkDeliveryCreateRequest) GetDeliveryDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeliveryDestination
}

// GetDeliveryDestinationOk returns a tuple with the DeliveryDestination field value
// and a boolean to check if the value has been set.
func (o *LinkDeliveryCreateRequest) GetDeliveryDestinationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DeliveryDestination, true
}

// SetDeliveryDestination sets field value
func (o *LinkDeliveryCreateRequest) SetDeliveryDestination(v string) {
	o.DeliveryDestination = v
}

func (o LinkDeliveryCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["link_token"] = o.LinkToken
	}
	if true {
		toSerialize["delivery_method"] = o.DeliveryMethod
	}
	if true {
		toSerialize["delivery_destination"] = o.DeliveryDestination
	}
	return json.Marshal(toSerialize)
}

type NullableLinkDeliveryCreateRequest struct {
	value *LinkDeliveryCreateRequest
	isSet bool
}

func (v NullableLinkDeliveryCreateRequest) Get() *LinkDeliveryCreateRequest {
	return v.value
}

func (v *NullableLinkDeliveryCreateRequest) Set(val *LinkDeliveryCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkDeliveryCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkDeliveryCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkDeliveryCreateRequest(val *LinkDeliveryCreateRequest) *NullableLinkDeliveryCreateRequest {
	return &NullableLinkDeliveryCreateRequest{value: val, isSet: true}
}

func (v NullableLinkDeliveryCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkDeliveryCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


