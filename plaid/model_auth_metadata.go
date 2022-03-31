/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.94.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// AuthMetadata Metadata that captures information about the Auth features of an institution.
type AuthMetadata struct {
	SupportedMethods NullableAuthSupportedMethods `json:"supported_methods"`
	AdditionalProperties map[string]interface{}
}

type _AuthMetadata AuthMetadata

// NewAuthMetadata instantiates a new AuthMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthMetadata(supportedMethods NullableAuthSupportedMethods) *AuthMetadata {
	this := AuthMetadata{}
	this.SupportedMethods = supportedMethods
	return &this
}

// NewAuthMetadataWithDefaults instantiates a new AuthMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthMetadataWithDefaults() *AuthMetadata {
	this := AuthMetadata{}
	return &this
}

// GetSupportedMethods returns the SupportedMethods field value
// If the value is explicit nil, the zero value for AuthSupportedMethods will be returned
func (o *AuthMetadata) GetSupportedMethods() AuthSupportedMethods {
	if o == nil || o.SupportedMethods.Get() == nil {
		var ret AuthSupportedMethods
		return ret
	}

	return *o.SupportedMethods.Get()
}

// GetSupportedMethodsOk returns a tuple with the SupportedMethods field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthMetadata) GetSupportedMethodsOk() (*AuthSupportedMethods, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SupportedMethods.Get(), o.SupportedMethods.IsSet()
}

// SetSupportedMethods sets field value
func (o *AuthMetadata) SetSupportedMethods(v AuthSupportedMethods) {
	o.SupportedMethods.Set(&v)
}

func (o AuthMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["supported_methods"] = o.SupportedMethods.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthMetadata) UnmarshalJSON(bytes []byte) (err error) {
	varAuthMetadata := _AuthMetadata{}

	if err = json.Unmarshal(bytes, &varAuthMetadata); err == nil {
		*o = AuthMetadata(varAuthMetadata)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "supported_methods")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthMetadata struct {
	value *AuthMetadata
	isSet bool
}

func (v NullableAuthMetadata) Get() *AuthMetadata {
	return v.value
}

func (v *NullableAuthMetadata) Set(val *AuthMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthMetadata(val *AuthMetadata) *NullableAuthMetadata {
	return &NullableAuthMetadata{value: val, isSet: true}
}

func (v NullableAuthMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


