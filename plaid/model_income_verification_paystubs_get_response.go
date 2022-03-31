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

// IncomeVerificationPaystubsGetResponse IncomeVerificationPaystubsGetResponse defines the response schema for `/income/verification/paystubs/get`.
type IncomeVerificationPaystubsGetResponse struct {
	// Metadata for an income document.
	DocumentMetadata *[]DocumentMetadata `json:"document_metadata,omitempty"`
	Paystubs []Paystub `json:"paystubs"`
	Error *PlaidError `json:"error,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _IncomeVerificationPaystubsGetResponse IncomeVerificationPaystubsGetResponse

// NewIncomeVerificationPaystubsGetResponse instantiates a new IncomeVerificationPaystubsGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeVerificationPaystubsGetResponse(paystubs []Paystub, requestId string) *IncomeVerificationPaystubsGetResponse {
	this := IncomeVerificationPaystubsGetResponse{}
	this.Paystubs = paystubs
	this.RequestId = requestId
	return &this
}

// NewIncomeVerificationPaystubsGetResponseWithDefaults instantiates a new IncomeVerificationPaystubsGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeVerificationPaystubsGetResponseWithDefaults() *IncomeVerificationPaystubsGetResponse {
	this := IncomeVerificationPaystubsGetResponse{}
	return &this
}

// GetDocumentMetadata returns the DocumentMetadata field value if set, zero value otherwise.
func (o *IncomeVerificationPaystubsGetResponse) GetDocumentMetadata() []DocumentMetadata {
	if o == nil || o.DocumentMetadata == nil {
		var ret []DocumentMetadata
		return ret
	}
	return *o.DocumentMetadata
}

// GetDocumentMetadataOk returns a tuple with the DocumentMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPaystubsGetResponse) GetDocumentMetadataOk() (*[]DocumentMetadata, bool) {
	if o == nil || o.DocumentMetadata == nil {
		return nil, false
	}
	return o.DocumentMetadata, true
}

// HasDocumentMetadata returns a boolean if a field has been set.
func (o *IncomeVerificationPaystubsGetResponse) HasDocumentMetadata() bool {
	if o != nil && o.DocumentMetadata != nil {
		return true
	}

	return false
}

// SetDocumentMetadata gets a reference to the given []DocumentMetadata and assigns it to the DocumentMetadata field.
func (o *IncomeVerificationPaystubsGetResponse) SetDocumentMetadata(v []DocumentMetadata) {
	o.DocumentMetadata = &v
}

// GetPaystubs returns the Paystubs field value
func (o *IncomeVerificationPaystubsGetResponse) GetPaystubs() []Paystub {
	if o == nil {
		var ret []Paystub
		return ret
	}

	return o.Paystubs
}

// GetPaystubsOk returns a tuple with the Paystubs field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPaystubsGetResponse) GetPaystubsOk() (*[]Paystub, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Paystubs, true
}

// SetPaystubs sets field value
func (o *IncomeVerificationPaystubsGetResponse) SetPaystubs(v []Paystub) {
	o.Paystubs = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *IncomeVerificationPaystubsGetResponse) GetError() PlaidError {
	if o == nil || o.Error == nil {
		var ret PlaidError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPaystubsGetResponse) GetErrorOk() (*PlaidError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *IncomeVerificationPaystubsGetResponse) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given PlaidError and assigns it to the Error field.
func (o *IncomeVerificationPaystubsGetResponse) SetError(v PlaidError) {
	o.Error = &v
}

// GetRequestId returns the RequestId field value
func (o *IncomeVerificationPaystubsGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPaystubsGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *IncomeVerificationPaystubsGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o IncomeVerificationPaystubsGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentMetadata != nil {
		toSerialize["document_metadata"] = o.DocumentMetadata
	}
	if true {
		toSerialize["paystubs"] = o.Paystubs
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IncomeVerificationPaystubsGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varIncomeVerificationPaystubsGetResponse := _IncomeVerificationPaystubsGetResponse{}

	if err = json.Unmarshal(bytes, &varIncomeVerificationPaystubsGetResponse); err == nil {
		*o = IncomeVerificationPaystubsGetResponse(varIncomeVerificationPaystubsGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "document_metadata")
		delete(additionalProperties, "paystubs")
		delete(additionalProperties, "error")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIncomeVerificationPaystubsGetResponse struct {
	value *IncomeVerificationPaystubsGetResponse
	isSet bool
}

func (v NullableIncomeVerificationPaystubsGetResponse) Get() *IncomeVerificationPaystubsGetResponse {
	return v.value
}

func (v *NullableIncomeVerificationPaystubsGetResponse) Set(val *IncomeVerificationPaystubsGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationPaystubsGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationPaystubsGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationPaystubsGetResponse(val *IncomeVerificationPaystubsGetResponse) *NullableIncomeVerificationPaystubsGetResponse {
	return &NullableIncomeVerificationPaystubsGetResponse{value: val, isSet: true}
}

func (v NullableIncomeVerificationPaystubsGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationPaystubsGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


