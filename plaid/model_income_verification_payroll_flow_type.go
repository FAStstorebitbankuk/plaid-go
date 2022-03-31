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
	"fmt"
)

// IncomeVerificationPayrollFlowType Flow types to retrieve payroll income data
type IncomeVerificationPayrollFlowType string

// List of IncomeVerificationPayrollFlowType
const (
	INCOMEVERIFICATIONPAYROLLFLOWTYPE_DIGITAL_INCOME IncomeVerificationPayrollFlowType = "payroll_digital_income"
	INCOMEVERIFICATIONPAYROLLFLOWTYPE_DOCUMENT_INCOME IncomeVerificationPayrollFlowType = "payroll_document_income"
)

var allowedIncomeVerificationPayrollFlowTypeEnumValues = []IncomeVerificationPayrollFlowType{
	"payroll_digital_income",
	"payroll_document_income",
}

func (v *IncomeVerificationPayrollFlowType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IncomeVerificationPayrollFlowType(value)
	for _, existing := range allowedIncomeVerificationPayrollFlowTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IncomeVerificationPayrollFlowType", value)
}

// NewIncomeVerificationPayrollFlowTypeFromValue returns a pointer to a valid IncomeVerificationPayrollFlowType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIncomeVerificationPayrollFlowTypeFromValue(v string) (*IncomeVerificationPayrollFlowType, error) {
	ev := IncomeVerificationPayrollFlowType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IncomeVerificationPayrollFlowType: valid values are %v", v, allowedIncomeVerificationPayrollFlowTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IncomeVerificationPayrollFlowType) IsValid() bool {
	for _, existing := range allowedIncomeVerificationPayrollFlowTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncomeVerificationPayrollFlowType value
func (v IncomeVerificationPayrollFlowType) Ptr() *IncomeVerificationPayrollFlowType {
	return &v
}

type NullableIncomeVerificationPayrollFlowType struct {
	value *IncomeVerificationPayrollFlowType
	isSet bool
}

func (v NullableIncomeVerificationPayrollFlowType) Get() *IncomeVerificationPayrollFlowType {
	return v.value
}

func (v *NullableIncomeVerificationPayrollFlowType) Set(val *IncomeVerificationPayrollFlowType) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationPayrollFlowType) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationPayrollFlowType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationPayrollFlowType(val *IncomeVerificationPayrollFlowType) *NullableIncomeVerificationPayrollFlowType {
	return &NullableIncomeVerificationPayrollFlowType{value: val, isSet: true}
}

func (v NullableIncomeVerificationPayrollFlowType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationPayrollFlowType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

