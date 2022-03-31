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

// TransferRepaymentReturnListResponse Defines the response schema for `/transfer/repayments/return/list`
type TransferRepaymentReturnListResponse struct {
	RepaymentReturns []TransferRepaymentReturn `json:"repayment_returns"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferRepaymentReturnListResponse TransferRepaymentReturnListResponse

// NewTransferRepaymentReturnListResponse instantiates a new TransferRepaymentReturnListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRepaymentReturnListResponse(repaymentReturns []TransferRepaymentReturn, requestId string) *TransferRepaymentReturnListResponse {
	this := TransferRepaymentReturnListResponse{}
	this.RepaymentReturns = repaymentReturns
	this.RequestId = requestId
	return &this
}

// NewTransferRepaymentReturnListResponseWithDefaults instantiates a new TransferRepaymentReturnListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRepaymentReturnListResponseWithDefaults() *TransferRepaymentReturnListResponse {
	this := TransferRepaymentReturnListResponse{}
	return &this
}

// GetRepaymentReturns returns the RepaymentReturns field value
func (o *TransferRepaymentReturnListResponse) GetRepaymentReturns() []TransferRepaymentReturn {
	if o == nil {
		var ret []TransferRepaymentReturn
		return ret
	}

	return o.RepaymentReturns
}

// GetRepaymentReturnsOk returns a tuple with the RepaymentReturns field value
// and a boolean to check if the value has been set.
func (o *TransferRepaymentReturnListResponse) GetRepaymentReturnsOk() (*[]TransferRepaymentReturn, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RepaymentReturns, true
}

// SetRepaymentReturns sets field value
func (o *TransferRepaymentReturnListResponse) SetRepaymentReturns(v []TransferRepaymentReturn) {
	o.RepaymentReturns = v
}

// GetRequestId returns the RequestId field value
func (o *TransferRepaymentReturnListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferRepaymentReturnListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferRepaymentReturnListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferRepaymentReturnListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["repayment_returns"] = o.RepaymentReturns
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferRepaymentReturnListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferRepaymentReturnListResponse := _TransferRepaymentReturnListResponse{}

	if err = json.Unmarshal(bytes, &varTransferRepaymentReturnListResponse); err == nil {
		*o = TransferRepaymentReturnListResponse(varTransferRepaymentReturnListResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "repayment_returns")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferRepaymentReturnListResponse struct {
	value *TransferRepaymentReturnListResponse
	isSet bool
}

func (v NullableTransferRepaymentReturnListResponse) Get() *TransferRepaymentReturnListResponse {
	return v.value
}

func (v *NullableTransferRepaymentReturnListResponse) Set(val *TransferRepaymentReturnListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRepaymentReturnListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRepaymentReturnListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRepaymentReturnListResponse(val *TransferRepaymentReturnListResponse) *NullableTransferRepaymentReturnListResponse {
	return &NullableTransferRepaymentReturnListResponse{value: val, isSet: true}
}

func (v NullableTransferRepaymentReturnListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRepaymentReturnListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


