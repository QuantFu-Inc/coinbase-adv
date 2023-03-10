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

// OrderOrderConfigurationMarketMarketIoc struct for OrderOrderConfigurationMarketMarketIoc
type OrderOrderConfigurationMarketMarketIoc struct {
	QuoteSize *float64 `json:"quote_size,omitempty,string"`
	BaseSize *float64 `json:"base_size,omitempty,string"`
}

// NewOrderOrderConfigurationMarketMarketIoc instantiates a new OrderOrderConfigurationMarketMarketIoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderOrderConfigurationMarketMarketIoc() *OrderOrderConfigurationMarketMarketIoc {
	this := OrderOrderConfigurationMarketMarketIoc{}
	return &this
}

// NewOrderOrderConfigurationMarketMarketIocWithDefaults instantiates a new OrderOrderConfigurationMarketMarketIoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderOrderConfigurationMarketMarketIocWithDefaults() *OrderOrderConfigurationMarketMarketIoc {
	this := OrderOrderConfigurationMarketMarketIoc{}
	return &this
}

// GetQuoteSize returns the QuoteSize field value if set, zero value otherwise.
func (o *OrderOrderConfigurationMarketMarketIoc) GetQuoteSize() float64 {
	if o == nil || isNil(o.QuoteSize) {
		var ret float64
		return ret
	}
	return *o.QuoteSize
}

// GetQuoteSizeOk returns a tuple with the QuoteSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderConfigurationMarketMarketIoc) GetQuoteSizeOk() (*float64, bool) {
	if o == nil || isNil(o.QuoteSize) {
    return nil, false
	}
	return o.QuoteSize, true
}

// HasQuoteSize returns a boolean if a field has been set.
func (o *OrderOrderConfigurationMarketMarketIoc) HasQuoteSize() bool {
	if o != nil && !isNil(o.QuoteSize) {
		return true
	}

	return false
}

// SetQuoteSize gets a reference to the given float64 and assigns it to the QuoteSize field.
func (o *OrderOrderConfigurationMarketMarketIoc) SetQuoteSize(v float64) {
	o.QuoteSize = &v
}

// GetBaseSize returns the BaseSize field value if set, zero value otherwise.
func (o *OrderOrderConfigurationMarketMarketIoc) GetBaseSize() float64 {
	if o == nil || isNil(o.BaseSize) {
		var ret float64
		return ret
	}
	return *o.BaseSize
}

// GetBaseSizeOk returns a tuple with the BaseSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderConfigurationMarketMarketIoc) GetBaseSizeOk() (*float64, bool) {
	if o == nil || isNil(o.BaseSize) {
    return nil, false
	}
	return o.BaseSize, true
}

// HasBaseSize returns a boolean if a field has been set.
func (o *OrderOrderConfigurationMarketMarketIoc) HasBaseSize() bool {
	if o != nil && !isNil(o.BaseSize) {
		return true
	}

	return false
}

// SetBaseSize gets a reference to the given float64 and assigns it to the BaseSize field.
func (o *OrderOrderConfigurationMarketMarketIoc) SetBaseSize(v float64) {
	o.BaseSize = &v
}

func (o OrderOrderConfigurationMarketMarketIoc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.QuoteSize) {
		toSerialize["quote_size"] = o.QuoteSize
	}
	if !isNil(o.BaseSize) {
		toSerialize["base_size"] = o.BaseSize
	}
	return json.Marshal(toSerialize)
}

type NullableOrderOrderConfigurationMarketMarketIoc struct {
	value *OrderOrderConfigurationMarketMarketIoc
	isSet bool
}

func (v NullableOrderOrderConfigurationMarketMarketIoc) Get() *OrderOrderConfigurationMarketMarketIoc {
	return v.value
}

func (v *NullableOrderOrderConfigurationMarketMarketIoc) Set(val *OrderOrderConfigurationMarketMarketIoc) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderOrderConfigurationMarketMarketIoc) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderOrderConfigurationMarketMarketIoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderOrderConfigurationMarketMarketIoc(val *OrderOrderConfigurationMarketMarketIoc) *NullableOrderOrderConfigurationMarketMarketIoc {
	return &NullableOrderOrderConfigurationMarketMarketIoc{value: val, isSet: true}
}

func (v NullableOrderOrderConfigurationMarketMarketIoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderOrderConfigurationMarketMarketIoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


