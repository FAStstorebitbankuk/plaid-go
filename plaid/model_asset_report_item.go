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
	"time"
)

// AssetReportItem A representation of an Item within an Asset Report.
type AssetReportItem struct {
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	// The full financial institution name associated with the Item.
	InstitutionName string `json:"institution_name"`
	// The id of the financial institution associated with the Item.
	InstitutionId string `json:"institution_id"`
	// The date and time when this Item’s data was last retrieved from the financial institution, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.
	DateLastUpdated time.Time `json:"date_last_updated"`
	// Data about each of the accounts open on the Item.
	Accounts []AccountAssets `json:"accounts"`
	AdditionalProperties map[string]interface{}
}

type _AssetReportItem AssetReportItem

// NewAssetReportItem instantiates a new AssetReportItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportItem(itemId string, institutionName string, institutionId string, dateLastUpdated time.Time, accounts []AccountAssets) *AssetReportItem {
	this := AssetReportItem{}
	this.ItemId = itemId
	this.InstitutionName = institutionName
	this.InstitutionId = institutionId
	this.DateLastUpdated = dateLastUpdated
	this.Accounts = accounts
	return &this
}

// NewAssetReportItemWithDefaults instantiates a new AssetReportItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportItemWithDefaults() *AssetReportItem {
	this := AssetReportItem{}
	return &this
}

// GetItemId returns the ItemId field value
func (o *AssetReportItem) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *AssetReportItem) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *AssetReportItem) SetItemId(v string) {
	o.ItemId = v
}

// GetInstitutionName returns the InstitutionName field value
func (o *AssetReportItem) GetInstitutionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstitutionName
}

// GetInstitutionNameOk returns a tuple with the InstitutionName field value
// and a boolean to check if the value has been set.
func (o *AssetReportItem) GetInstitutionNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstitutionName, true
}

// SetInstitutionName sets field value
func (o *AssetReportItem) SetInstitutionName(v string) {
	o.InstitutionName = v
}

// GetInstitutionId returns the InstitutionId field value
func (o *AssetReportItem) GetInstitutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value
// and a boolean to check if the value has been set.
func (o *AssetReportItem) GetInstitutionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstitutionId, true
}

// SetInstitutionId sets field value
func (o *AssetReportItem) SetInstitutionId(v string) {
	o.InstitutionId = v
}

// GetDateLastUpdated returns the DateLastUpdated field value
func (o *AssetReportItem) GetDateLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DateLastUpdated
}

// GetDateLastUpdatedOk returns a tuple with the DateLastUpdated field value
// and a boolean to check if the value has been set.
func (o *AssetReportItem) GetDateLastUpdatedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DateLastUpdated, true
}

// SetDateLastUpdated sets field value
func (o *AssetReportItem) SetDateLastUpdated(v time.Time) {
	o.DateLastUpdated = v
}

// GetAccounts returns the Accounts field value
func (o *AssetReportItem) GetAccounts() []AccountAssets {
	if o == nil {
		var ret []AccountAssets
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *AssetReportItem) GetAccountsOk() (*[]AccountAssets, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Accounts, true
}

// SetAccounts sets field value
func (o *AssetReportItem) SetAccounts(v []AccountAssets) {
	o.Accounts = v
}

func (o AssetReportItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item_id"] = o.ItemId
	}
	if true {
		toSerialize["institution_name"] = o.InstitutionName
	}
	if true {
		toSerialize["institution_id"] = o.InstitutionId
	}
	if true {
		toSerialize["date_last_updated"] = o.DateLastUpdated
	}
	if true {
		toSerialize["accounts"] = o.Accounts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetReportItem) UnmarshalJSON(bytes []byte) (err error) {
	varAssetReportItem := _AssetReportItem{}

	if err = json.Unmarshal(bytes, &varAssetReportItem); err == nil {
		*o = AssetReportItem(varAssetReportItem)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "institution_name")
		delete(additionalProperties, "institution_id")
		delete(additionalProperties, "date_last_updated")
		delete(additionalProperties, "accounts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetReportItem struct {
	value *AssetReportItem
	isSet bool
}

func (v NullableAssetReportItem) Get() *AssetReportItem {
	return v.value
}

func (v *NullableAssetReportItem) Set(val *AssetReportItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportItem(val *AssetReportItem) *NullableAssetReportItem {
	return &NullableAssetReportItem{value: val, isSet: true}
}

func (v NullableAssetReportItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


