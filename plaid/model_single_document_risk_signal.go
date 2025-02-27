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

// SingleDocumentRiskSignal Object containing all risk signals and relevant metadata for a single document
type SingleDocumentRiskSignal struct {
	DocumentReference RiskSignalDocumentReference `json:"document_reference"`
	// Array of attributes that indicate whether or not there is fraud risk with a document
	RiskSignals []DocumentRiskSignal `json:"risk_signals"`
	AdditionalProperties map[string]interface{}
}

type _SingleDocumentRiskSignal SingleDocumentRiskSignal

// NewSingleDocumentRiskSignal instantiates a new SingleDocumentRiskSignal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleDocumentRiskSignal(documentReference RiskSignalDocumentReference, riskSignals []DocumentRiskSignal) *SingleDocumentRiskSignal {
	this := SingleDocumentRiskSignal{}
	this.DocumentReference = documentReference
	this.RiskSignals = riskSignals
	return &this
}

// NewSingleDocumentRiskSignalWithDefaults instantiates a new SingleDocumentRiskSignal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleDocumentRiskSignalWithDefaults() *SingleDocumentRiskSignal {
	this := SingleDocumentRiskSignal{}
	return &this
}

// GetDocumentReference returns the DocumentReference field value
func (o *SingleDocumentRiskSignal) GetDocumentReference() RiskSignalDocumentReference {
	if o == nil {
		var ret RiskSignalDocumentReference
		return ret
	}

	return o.DocumentReference
}

// GetDocumentReferenceOk returns a tuple with the DocumentReference field value
// and a boolean to check if the value has been set.
func (o *SingleDocumentRiskSignal) GetDocumentReferenceOk() (*RiskSignalDocumentReference, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentReference, true
}

// SetDocumentReference sets field value
func (o *SingleDocumentRiskSignal) SetDocumentReference(v RiskSignalDocumentReference) {
	o.DocumentReference = v
}

// GetRiskSignals returns the RiskSignals field value
func (o *SingleDocumentRiskSignal) GetRiskSignals() []DocumentRiskSignal {
	if o == nil {
		var ret []DocumentRiskSignal
		return ret
	}

	return o.RiskSignals
}

// GetRiskSignalsOk returns a tuple with the RiskSignals field value
// and a boolean to check if the value has been set.
func (o *SingleDocumentRiskSignal) GetRiskSignalsOk() (*[]DocumentRiskSignal, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RiskSignals, true
}

// SetRiskSignals sets field value
func (o *SingleDocumentRiskSignal) SetRiskSignals(v []DocumentRiskSignal) {
	o.RiskSignals = v
}

func (o SingleDocumentRiskSignal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["document_reference"] = o.DocumentReference
	}
	if true {
		toSerialize["risk_signals"] = o.RiskSignals
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SingleDocumentRiskSignal) UnmarshalJSON(bytes []byte) (err error) {
	varSingleDocumentRiskSignal := _SingleDocumentRiskSignal{}

	if err = json.Unmarshal(bytes, &varSingleDocumentRiskSignal); err == nil {
		*o = SingleDocumentRiskSignal(varSingleDocumentRiskSignal)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "document_reference")
		delete(additionalProperties, "risk_signals")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSingleDocumentRiskSignal struct {
	value *SingleDocumentRiskSignal
	isSet bool
}

func (v NullableSingleDocumentRiskSignal) Get() *SingleDocumentRiskSignal {
	return v.value
}

func (v *NullableSingleDocumentRiskSignal) Set(val *SingleDocumentRiskSignal) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleDocumentRiskSignal) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleDocumentRiskSignal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleDocumentRiskSignal(val *SingleDocumentRiskSignal) *NullableSingleDocumentRiskSignal {
	return &NullableSingleDocumentRiskSignal{value: val, isSet: true}
}

func (v NullableSingleDocumentRiskSignal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleDocumentRiskSignal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


