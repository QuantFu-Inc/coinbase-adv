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

// OutputOrderConfigurationStopLimitStopLimitGtd struct for OutputOrderConfigurationStopLimitStopLimitGtd
type OutputOrderConfigurationStopLimitStopLimitGtd struct {
	BaseSize *float64 `json:"base_size,omitempty,string"`
	LimitPrice *float64 `json:"limit_price,omitempty,string"`
	StopPrice *float64 `json:"stop_price,omitempty,string"`
	EndTime *string `json:"end_time,omitempty"`
	StopDirection *string `json:"stop_direction,omitempty"`
}

// NewOutputOrderConfigurationStopLimitStopLimitGtd instantiates a new OutputOrderConfigurationStopLimitStopLimitGtd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputOrderConfigurationStopLimitStopLimitGtd() *OutputOrderConfigurationStopLimitStopLimitGtd {
	this := OutputOrderConfigurationStopLimitStopLimitGtd{}
	return &this
}

// NewOutputOrderConfigurationStopLimitStopLimitGtdWithDefaults instantiates a new OutputOrderConfigurationStopLimitStopLimitGtd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputOrderConfigurationStopLimitStopLimitGtdWithDefaults() *OutputOrderConfigurationStopLimitStopLimitGtd {
	this := OutputOrderConfigurationStopLimitStopLimitGtd{}
	return &this
}

// GetBaseSize returns the BaseSize field value if set, zero value otherwise.
func (o *OutputOrderConfigurationStopLimitStopLimitGtd) GetBaseSize() float64 {
	if o == nil || isNil(o.BaseSize) {
		var ret float64
		return ret
	}
	return *o.BaseSize
}

// GetBaseSizeOk returns a tuple with the BaseSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputOrderConfigurationStopLimitStopLimitGtd) GetBaseSizeOk() (*float64, bool) {
	if o == nil || isNil(o.BaseSize) {
    return nil, false
	}
	return o.BaseSize, true
}

// HasBaseSize returns a boolean if a field has been set.
func (o *OutputOrderConfigurationStopLimitStopLimitGtd) HasBaseSize() bool {
	if o != nil && !isNil(o.BaseSize) {
		return true
	}

	return false
}

// SetBaseSize gets a reference to the given float64 and assigns it to the BaseSize field.
func (o *OutputOrderConfigurationStopLimitStopLimitGtd) SetBaseSize(v float64) {
	o.BaseSize = &v
}

// GetLimitPrice returns the LimitPrice field value if set, zero value otherwise.
func (o *OutputOrderConfigurationStopLimitStopLimitGtd) GetLimitPrice() float64 {
	if o == nil || isNil(o.LimitPrice) {
		var ret float64
		return ret
	}
	return *o.LimitPrice
}

// GetLimitPriceOk returns a tuple with the LimitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputOrderConfigurationStopLimitStopLimitGtd) GetLimitPriceOk() (*float64, bool) {
	if o == nil || isNil(o.LimitPrice) {
    return nil, false
	}
	return o.LimitPrice, true
}

// HasLimitPrice returns a boolean if a field has been set.
func (o *OutputOrderConfigurationStopLimitStopLimitGtd) HasLimitPrice() bool {
	if o != nil && !isNil(o.LimitPrice) {
		return true
	}

	return false
}

// SetLimitPrice gets a reference to the given float64 and assigns it to the LimitPrice field.
func (o *OutputOrderConfigurationStopLimitStopLimitGtd) SetLimitPrice(v float64) {
	o.LimitPrice = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *OutputOrderConfigurationStopLimitStopLimitGtd) GetStopPrice() float64 {
	if o == nil || isNil(o.StopPrice) {
		var ret float64
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputOrderConfigurationStopLimitStopLimitGtd) GetStopPriceOk() (*float64, bool) {
	if o == nil || isNil(o.StopPrice) {
    return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *OutputOrderConfigurationStopLimitStopLimitGtd) HasStopPrice() bool {
	if o != nil && !isNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given float64 and assigns it to the StopPrice field.
func (o *OutputOrderConfigurationStopLimitStopLimitGtd) SetStopPrice(v float64) {
	o.StopPrice = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *OutputOrderConfigurationStopLimitStopLimitGtd) GetEndTime() string {
	if o == nil || isNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputOrderConfigurationStopLimitStopLimitGtd) GetEndTimeOk() (*string, bool) {
	if o == nil || isNil(o.EndTime) {
    return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *OutputOrderConfigurationStopLimitStopLimitGtd) HasEndTime() bool {
	if o != nil && !isNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *OutputOrderConfigurationStopLimitStopLimitGtd) SetEndTime(v string) {
	o.EndTime = &v
}

// GetStopDirection returns the StopDirection field value if set, zero value otherwise.
func (o *OutputOrderConfigurationStopLimitStopLimitGtd) GetStopDirection() string {
	if o == nil || isNil(o.StopDirection) {
		var ret string
		return ret
	}
	return *o.StopDirection
}

// GetStopDirectionOk returns a tuple with the StopDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputOrderConfigurationStopLimitStopLimitGtd) GetStopDirectionOk() (*string, bool) {
	if o == nil || isNil(o.StopDirection) {
    return nil, false
	}
	return o.StopDirection, true
}

// HasStopDirection returns a boolean if a field has been set.
func (o *OutputOrderConfigurationStopLimitStopLimitGtd) HasStopDirection() bool {
	if o != nil && !isNil(o.StopDirection) {
		return true
	}

	return false
}

// SetStopDirection gets a reference to the given string and assigns it to the StopDirection field.
func (o *OutputOrderConfigurationStopLimitStopLimitGtd) SetStopDirection(v string) {
	o.StopDirection = &v
}

func (o OutputOrderConfigurationStopLimitStopLimitGtd) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BaseSize) {
		toSerialize["base_size"] = o.BaseSize
	}
	if !isNil(o.LimitPrice) {
		toSerialize["limit_price"] = o.LimitPrice
	}
	if !isNil(o.StopPrice) {
		toSerialize["stop_price"] = o.StopPrice
	}
	if !isNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !isNil(o.StopDirection) {
		toSerialize["stop_direction"] = o.StopDirection
	}
	return json.Marshal(toSerialize)
}

type NullableOutputOrderConfigurationStopLimitStopLimitGtd struct {
	value *OutputOrderConfigurationStopLimitStopLimitGtd
	isSet bool
}

func (v NullableOutputOrderConfigurationStopLimitStopLimitGtd) Get() *OutputOrderConfigurationStopLimitStopLimitGtd {
	return v.value
}

func (v *NullableOutputOrderConfigurationStopLimitStopLimitGtd) Set(val *OutputOrderConfigurationStopLimitStopLimitGtd) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputOrderConfigurationStopLimitStopLimitGtd) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputOrderConfigurationStopLimitStopLimitGtd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputOrderConfigurationStopLimitStopLimitGtd(val *OutputOrderConfigurationStopLimitStopLimitGtd) *NullableOutputOrderConfigurationStopLimitStopLimitGtd {
	return &NullableOutputOrderConfigurationStopLimitStopLimitGtd{value: val, isSet: true}
}

func (v NullableOutputOrderConfigurationStopLimitStopLimitGtd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputOrderConfigurationStopLimitStopLimitGtd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


