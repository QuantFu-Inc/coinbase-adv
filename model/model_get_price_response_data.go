/*
Coinbase Advanced Trading API

OpenAPI 3.x specification for Coinbase Adavanced Trading

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// GetPriceResponseData struct for GetPriceResponseData
type GetPriceResponseData struct {
	Amount *float64 `json:"amount,omitempty,string"`
	Currency *string `json:"currency,omitempty"`
}

// NewGetPriceResponseData instantiates a new GetPriceResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPriceResponseData() *GetPriceResponseData {
	this := GetPriceResponseData{}
	return &this
}

// NewGetPriceResponseDataWithDefaults instantiates a new GetPriceResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPriceResponseDataWithDefaults() *GetPriceResponseData {
	this := GetPriceResponseData{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetPriceResponseData) GetAmount() float64 {
	if o == nil || isNil(o.Amount) {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPriceResponseData) GetAmountOk() (*float64, bool) {
	if o == nil || isNil(o.Amount) {
    return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetPriceResponseData) HasAmount() bool {
	if o != nil && !isNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *GetPriceResponseData) SetAmount(v float64) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *GetPriceResponseData) GetCurrency() string {
	if o == nil || isNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPriceResponseData) GetCurrencyOk() (*string, bool) {
	if o == nil || isNil(o.Currency) {
    return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *GetPriceResponseData) HasCurrency() bool {
	if o != nil && !isNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *GetPriceResponseData) SetCurrency(v string) {
	o.Currency = &v
}

func (o GetPriceResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !isNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

type NullableGetPriceResponseData struct {
	value *GetPriceResponseData
	isSet bool
}

func (v NullableGetPriceResponseData) Get() *GetPriceResponseData {
	return v.value
}

func (v *NullableGetPriceResponseData) Set(val *GetPriceResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPriceResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPriceResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPriceResponseData(val *GetPriceResponseData) *NullableGetPriceResponseData {
	return &NullableGetPriceResponseData{value: val, isSet: true}
}

func (v NullableGetPriceResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPriceResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


