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

// GetProductResponse struct for GetProductResponse
type GetProductResponse struct {
	ProductId *string `json:"product_id,omitempty"`
	Price *float64 `json:"price,omitempty,string"`
	PricePercentageChange24h *float64 `json:"price_percentage_change_24h,omitempty,string"`
	Volume24h *float64 `json:"volume_24h,omitempty,string"`
	VolumePercentageChange24h *float64 `json:"volume_percentage_change_24h,omitempty,string"`
	BaseIncrement *float64 `json:"base_increment,omitempty,string"`
	QuoteIncrement *float64 `json:"quote_increment,omitempty,string"`
	QuoteMinSize *float64 `json:"quote_min_size,omitempty,string"`
	QuoteMaxSize *float64 `json:"quote_max_size,omitempty,string"`
	BaseMinSize *float64 `json:"base_min_size,omitempty,string"`
	BaseMaxSize *float64 `json:"base_max_size,omitempty,string"`
	BaseName *string `json:"base_name,omitempty"`
	QuoteName *string `json:"quote_name,omitempty"`
	Watched *bool `json:"watched,omitempty"`
	IsDisabled *bool `json:"is_disabled,omitempty"`
	New *bool `json:"new,omitempty"`
	Status *string `json:"status,omitempty"`
	CancelOnly *bool `json:"cancel_only,omitempty"`
	LimitOnly *bool `json:"limit_only,omitempty"`
	PostOnly *bool `json:"post_only,omitempty"`
	TradingDisabled *bool `json:"trading_disabled,omitempty"`
	AuctionMode *bool `json:"auction_mode,omitempty"`
	ProductType *string `json:"product_type,omitempty"`
	QuoteCurrencyId *string `json:"quote_currency_id,omitempty"`
	BaseCurrencyId *string `json:"base_currency_id,omitempty"`
	FcmTradingSessionDetails *string `json:"fcm_trading_session_details,omitempty"`
	MidMarketPrice *string `json:"mid_market_price,omitempty"`
	BaseDisplaySymbol *string `json:"base_display_symbol,omitempty"`
	QuoteDisplaySymbol *string `json:"quote_display_symbol,omitempty"`
}

// NewGetProductResponse instantiates a new GetProductResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductResponse() *GetProductResponse {
	this := GetProductResponse{}
	return &this
}

// NewGetProductResponseWithDefaults instantiates a new GetProductResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductResponseWithDefaults() *GetProductResponse {
	this := GetProductResponse{}
	return &this
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *GetProductResponse) GetProductId() string {
	if o == nil || isNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetProductIdOk() (*string, bool) {
	if o == nil || isNil(o.ProductId) {
    return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *GetProductResponse) HasProductId() bool {
	if o != nil && !isNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *GetProductResponse) SetProductId(v string) {
	o.ProductId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *GetProductResponse) GetPrice() float64 {
	if o == nil || isNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetPriceOk() (*float64, bool) {
	if o == nil || isNil(o.Price) {
    return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *GetProductResponse) HasPrice() bool {
	if o != nil && !isNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *GetProductResponse) SetPrice(v float64) {
	o.Price = &v
}

// GetPricePercentageChange24h returns the PricePercentageChange24h field value if set, zero value otherwise.
func (o *GetProductResponse) GetPricePercentageChange24h() float64 {
	if o == nil || isNil(o.PricePercentageChange24h) {
		var ret float64
		return ret
	}
	return *o.PricePercentageChange24h
}

// GetPricePercentageChange24hOk returns a tuple with the PricePercentageChange24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetPricePercentageChange24hOk() (*float64, bool) {
	if o == nil || isNil(o.PricePercentageChange24h) {
    return nil, false
	}
	return o.PricePercentageChange24h, true
}

// HasPricePercentageChange24h returns a boolean if a field has been set.
func (o *GetProductResponse) HasPricePercentageChange24h() bool {
	if o != nil && !isNil(o.PricePercentageChange24h) {
		return true
	}

	return false
}

// SetPricePercentageChange24h gets a reference to the given float64 and assigns it to the PricePercentageChange24h field.
func (o *GetProductResponse) SetPricePercentageChange24h(v float64) {
	o.PricePercentageChange24h = &v
}

// GetVolume24h returns the Volume24h field value if set, zero value otherwise.
func (o *GetProductResponse) GetVolume24h() float64 {
	if o == nil || isNil(o.Volume24h) {
		var ret float64
		return ret
	}
	return *o.Volume24h
}

// GetVolume24hOk returns a tuple with the Volume24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetVolume24hOk() (*float64, bool) {
	if o == nil || isNil(o.Volume24h) {
    return nil, false
	}
	return o.Volume24h, true
}

// HasVolume24h returns a boolean if a field has been set.
func (o *GetProductResponse) HasVolume24h() bool {
	if o != nil && !isNil(o.Volume24h) {
		return true
	}

	return false
}

// SetVolume24h gets a reference to the given float64 and assigns it to the Volume24h field.
func (o *GetProductResponse) SetVolume24h(v float64) {
	o.Volume24h = &v
}

// GetVolumePercentageChange24h returns the VolumePercentageChange24h field value if set, zero value otherwise.
func (o *GetProductResponse) GetVolumePercentageChange24h() float64 {
	if o == nil || isNil(o.VolumePercentageChange24h) {
		var ret float64
		return ret
	}
	return *o.VolumePercentageChange24h
}

// GetVolumePercentageChange24hOk returns a tuple with the VolumePercentageChange24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetVolumePercentageChange24hOk() (*float64, bool) {
	if o == nil || isNil(o.VolumePercentageChange24h) {
    return nil, false
	}
	return o.VolumePercentageChange24h, true
}

// HasVolumePercentageChange24h returns a boolean if a field has been set.
func (o *GetProductResponse) HasVolumePercentageChange24h() bool {
	if o != nil && !isNil(o.VolumePercentageChange24h) {
		return true
	}

	return false
}

// SetVolumePercentageChange24h gets a reference to the given float64 and assigns it to the VolumePercentageChange24h field.
func (o *GetProductResponse) SetVolumePercentageChange24h(v float64) {
	o.VolumePercentageChange24h = &v
}

// GetBaseIncrement returns the BaseIncrement field value if set, zero value otherwise.
func (o *GetProductResponse) GetBaseIncrement() float64 {
	if o == nil || isNil(o.BaseIncrement) {
		var ret float64
		return ret
	}
	return *o.BaseIncrement
}

// GetBaseIncrementOk returns a tuple with the BaseIncrement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetBaseIncrementOk() (*float64, bool) {
	if o == nil || isNil(o.BaseIncrement) {
    return nil, false
	}
	return o.BaseIncrement, true
}

// HasBaseIncrement returns a boolean if a field has been set.
func (o *GetProductResponse) HasBaseIncrement() bool {
	if o != nil && !isNil(o.BaseIncrement) {
		return true
	}

	return false
}

// SetBaseIncrement gets a reference to the given float64 and assigns it to the BaseIncrement field.
func (o *GetProductResponse) SetBaseIncrement(v float64) {
	o.BaseIncrement = &v
}

// GetQuoteIncrement returns the QuoteIncrement field value if set, zero value otherwise.
func (o *GetProductResponse) GetQuoteIncrement() float64 {
	if o == nil || isNil(o.QuoteIncrement) {
		var ret float64
		return ret
	}
	return *o.QuoteIncrement
}

// GetQuoteIncrementOk returns a tuple with the QuoteIncrement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetQuoteIncrementOk() (*float64, bool) {
	if o == nil || isNil(o.QuoteIncrement) {
    return nil, false
	}
	return o.QuoteIncrement, true
}

// HasQuoteIncrement returns a boolean if a field has been set.
func (o *GetProductResponse) HasQuoteIncrement() bool {
	if o != nil && !isNil(o.QuoteIncrement) {
		return true
	}

	return false
}

// SetQuoteIncrement gets a reference to the given float64 and assigns it to the QuoteIncrement field.
func (o *GetProductResponse) SetQuoteIncrement(v float64) {
	o.QuoteIncrement = &v
}

// GetQuoteMinSize returns the QuoteMinSize field value if set, zero value otherwise.
func (o *GetProductResponse) GetQuoteMinSize() float64 {
	if o == nil || isNil(o.QuoteMinSize) {
		var ret float64
		return ret
	}
	return *o.QuoteMinSize
}

// GetQuoteMinSizeOk returns a tuple with the QuoteMinSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetQuoteMinSizeOk() (*float64, bool) {
	if o == nil || isNil(o.QuoteMinSize) {
    return nil, false
	}
	return o.QuoteMinSize, true
}

// HasQuoteMinSize returns a boolean if a field has been set.
func (o *GetProductResponse) HasQuoteMinSize() bool {
	if o != nil && !isNil(o.QuoteMinSize) {
		return true
	}

	return false
}

// SetQuoteMinSize gets a reference to the given float64 and assigns it to the QuoteMinSize field.
func (o *GetProductResponse) SetQuoteMinSize(v float64) {
	o.QuoteMinSize = &v
}

// GetQuoteMaxSize returns the QuoteMaxSize field value if set, zero value otherwise.
func (o *GetProductResponse) GetQuoteMaxSize() float64 {
	if o == nil || isNil(o.QuoteMaxSize) {
		var ret float64
		return ret
	}
	return *o.QuoteMaxSize
}

// GetQuoteMaxSizeOk returns a tuple with the QuoteMaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetQuoteMaxSizeOk() (*float64, bool) {
	if o == nil || isNil(o.QuoteMaxSize) {
    return nil, false
	}
	return o.QuoteMaxSize, true
}

// HasQuoteMaxSize returns a boolean if a field has been set.
func (o *GetProductResponse) HasQuoteMaxSize() bool {
	if o != nil && !isNil(o.QuoteMaxSize) {
		return true
	}

	return false
}

// SetQuoteMaxSize gets a reference to the given float64 and assigns it to the QuoteMaxSize field.
func (o *GetProductResponse) SetQuoteMaxSize(v float64) {
	o.QuoteMaxSize = &v
}

// GetBaseMinSize returns the BaseMinSize field value if set, zero value otherwise.
func (o *GetProductResponse) GetBaseMinSize() float64 {
	if o == nil || isNil(o.BaseMinSize) {
		var ret float64
		return ret
	}
	return *o.BaseMinSize
}

// GetBaseMinSizeOk returns a tuple with the BaseMinSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetBaseMinSizeOk() (*float64, bool) {
	if o == nil || isNil(o.BaseMinSize) {
    return nil, false
	}
	return o.BaseMinSize, true
}

// HasBaseMinSize returns a boolean if a field has been set.
func (o *GetProductResponse) HasBaseMinSize() bool {
	if o != nil && !isNil(o.BaseMinSize) {
		return true
	}

	return false
}

// SetBaseMinSize gets a reference to the given float64 and assigns it to the BaseMinSize field.
func (o *GetProductResponse) SetBaseMinSize(v float64) {
	o.BaseMinSize = &v
}

// GetBaseMaxSize returns the BaseMaxSize field value if set, zero value otherwise.
func (o *GetProductResponse) GetBaseMaxSize() float64 {
	if o == nil || isNil(o.BaseMaxSize) {
		var ret float64
		return ret
	}
	return *o.BaseMaxSize
}

// GetBaseMaxSizeOk returns a tuple with the BaseMaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetBaseMaxSizeOk() (*float64, bool) {
	if o == nil || isNil(o.BaseMaxSize) {
    return nil, false
	}
	return o.BaseMaxSize, true
}

// HasBaseMaxSize returns a boolean if a field has been set.
func (o *GetProductResponse) HasBaseMaxSize() bool {
	if o != nil && !isNil(o.BaseMaxSize) {
		return true
	}

	return false
}

// SetBaseMaxSize gets a reference to the given float64 and assigns it to the BaseMaxSize field.
func (o *GetProductResponse) SetBaseMaxSize(v float64) {
	o.BaseMaxSize = &v
}

// GetBaseName returns the BaseName field value if set, zero value otherwise.
func (o *GetProductResponse) GetBaseName() string {
	if o == nil || isNil(o.BaseName) {
		var ret string
		return ret
	}
	return *o.BaseName
}

// GetBaseNameOk returns a tuple with the BaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetBaseNameOk() (*string, bool) {
	if o == nil || isNil(o.BaseName) {
    return nil, false
	}
	return o.BaseName, true
}

// HasBaseName returns a boolean if a field has been set.
func (o *GetProductResponse) HasBaseName() bool {
	if o != nil && !isNil(o.BaseName) {
		return true
	}

	return false
}

// SetBaseName gets a reference to the given string and assigns it to the BaseName field.
func (o *GetProductResponse) SetBaseName(v string) {
	o.BaseName = &v
}

// GetQuoteName returns the QuoteName field value if set, zero value otherwise.
func (o *GetProductResponse) GetQuoteName() string {
	if o == nil || isNil(o.QuoteName) {
		var ret string
		return ret
	}
	return *o.QuoteName
}

// GetQuoteNameOk returns a tuple with the QuoteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetQuoteNameOk() (*string, bool) {
	if o == nil || isNil(o.QuoteName) {
    return nil, false
	}
	return o.QuoteName, true
}

// HasQuoteName returns a boolean if a field has been set.
func (o *GetProductResponse) HasQuoteName() bool {
	if o != nil && !isNil(o.QuoteName) {
		return true
	}

	return false
}

// SetQuoteName gets a reference to the given string and assigns it to the QuoteName field.
func (o *GetProductResponse) SetQuoteName(v string) {
	o.QuoteName = &v
}

// GetWatched returns the Watched field value if set, zero value otherwise.
func (o *GetProductResponse) GetWatched() bool {
	if o == nil || isNil(o.Watched) {
		var ret bool
		return ret
	}
	return *o.Watched
}

// GetWatchedOk returns a tuple with the Watched field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetWatchedOk() (*bool, bool) {
	if o == nil || isNil(o.Watched) {
    return nil, false
	}
	return o.Watched, true
}

// HasWatched returns a boolean if a field has been set.
func (o *GetProductResponse) HasWatched() bool {
	if o != nil && !isNil(o.Watched) {
		return true
	}

	return false
}

// SetWatched gets a reference to the given bool and assigns it to the Watched field.
func (o *GetProductResponse) SetWatched(v bool) {
	o.Watched = &v
}

// GetIsDisabled returns the IsDisabled field value if set, zero value otherwise.
func (o *GetProductResponse) GetIsDisabled() bool {
	if o == nil || isNil(o.IsDisabled) {
		var ret bool
		return ret
	}
	return *o.IsDisabled
}

// GetIsDisabledOk returns a tuple with the IsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetIsDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsDisabled) {
    return nil, false
	}
	return o.IsDisabled, true
}

// HasIsDisabled returns a boolean if a field has been set.
func (o *GetProductResponse) HasIsDisabled() bool {
	if o != nil && !isNil(o.IsDisabled) {
		return true
	}

	return false
}

// SetIsDisabled gets a reference to the given bool and assigns it to the IsDisabled field.
func (o *GetProductResponse) SetIsDisabled(v bool) {
	o.IsDisabled = &v
}

// GetNew returns the New field value if set, zero value otherwise.
func (o *GetProductResponse) GetNew() bool {
	if o == nil || isNil(o.New) {
		var ret bool
		return ret
	}
	return *o.New
}

// GetNewOk returns a tuple with the New field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetNewOk() (*bool, bool) {
	if o == nil || isNil(o.New) {
    return nil, false
	}
	return o.New, true
}

// HasNew returns a boolean if a field has been set.
func (o *GetProductResponse) HasNew() bool {
	if o != nil && !isNil(o.New) {
		return true
	}

	return false
}

// SetNew gets a reference to the given bool and assigns it to the New field.
func (o *GetProductResponse) SetNew(v bool) {
	o.New = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetProductResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetProductResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetProductResponse) SetStatus(v string) {
	o.Status = &v
}

// GetCancelOnly returns the CancelOnly field value if set, zero value otherwise.
func (o *GetProductResponse) GetCancelOnly() bool {
	if o == nil || isNil(o.CancelOnly) {
		var ret bool
		return ret
	}
	return *o.CancelOnly
}

// GetCancelOnlyOk returns a tuple with the CancelOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetCancelOnlyOk() (*bool, bool) {
	if o == nil || isNil(o.CancelOnly) {
    return nil, false
	}
	return o.CancelOnly, true
}

// HasCancelOnly returns a boolean if a field has been set.
func (o *GetProductResponse) HasCancelOnly() bool {
	if o != nil && !isNil(o.CancelOnly) {
		return true
	}

	return false
}

// SetCancelOnly gets a reference to the given bool and assigns it to the CancelOnly field.
func (o *GetProductResponse) SetCancelOnly(v bool) {
	o.CancelOnly = &v
}

// GetLimitOnly returns the LimitOnly field value if set, zero value otherwise.
func (o *GetProductResponse) GetLimitOnly() bool {
	if o == nil || isNil(o.LimitOnly) {
		var ret bool
		return ret
	}
	return *o.LimitOnly
}

// GetLimitOnlyOk returns a tuple with the LimitOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetLimitOnlyOk() (*bool, bool) {
	if o == nil || isNil(o.LimitOnly) {
    return nil, false
	}
	return o.LimitOnly, true
}

// HasLimitOnly returns a boolean if a field has been set.
func (o *GetProductResponse) HasLimitOnly() bool {
	if o != nil && !isNil(o.LimitOnly) {
		return true
	}

	return false
}

// SetLimitOnly gets a reference to the given bool and assigns it to the LimitOnly field.
func (o *GetProductResponse) SetLimitOnly(v bool) {
	o.LimitOnly = &v
}

// GetPostOnly returns the PostOnly field value if set, zero value otherwise.
func (o *GetProductResponse) GetPostOnly() bool {
	if o == nil || isNil(o.PostOnly) {
		var ret bool
		return ret
	}
	return *o.PostOnly
}

// GetPostOnlyOk returns a tuple with the PostOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetPostOnlyOk() (*bool, bool) {
	if o == nil || isNil(o.PostOnly) {
    return nil, false
	}
	return o.PostOnly, true
}

// HasPostOnly returns a boolean if a field has been set.
func (o *GetProductResponse) HasPostOnly() bool {
	if o != nil && !isNil(o.PostOnly) {
		return true
	}

	return false
}

// SetPostOnly gets a reference to the given bool and assigns it to the PostOnly field.
func (o *GetProductResponse) SetPostOnly(v bool) {
	o.PostOnly = &v
}

// GetTradingDisabled returns the TradingDisabled field value if set, zero value otherwise.
func (o *GetProductResponse) GetTradingDisabled() bool {
	if o == nil || isNil(o.TradingDisabled) {
		var ret bool
		return ret
	}
	return *o.TradingDisabled
}

// GetTradingDisabledOk returns a tuple with the TradingDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetTradingDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.TradingDisabled) {
    return nil, false
	}
	return o.TradingDisabled, true
}

// HasTradingDisabled returns a boolean if a field has been set.
func (o *GetProductResponse) HasTradingDisabled() bool {
	if o != nil && !isNil(o.TradingDisabled) {
		return true
	}

	return false
}

// SetTradingDisabled gets a reference to the given bool and assigns it to the TradingDisabled field.
func (o *GetProductResponse) SetTradingDisabled(v bool) {
	o.TradingDisabled = &v
}

// GetAuctionMode returns the AuctionMode field value if set, zero value otherwise.
func (o *GetProductResponse) GetAuctionMode() bool {
	if o == nil || isNil(o.AuctionMode) {
		var ret bool
		return ret
	}
	return *o.AuctionMode
}

// GetAuctionModeOk returns a tuple with the AuctionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetAuctionModeOk() (*bool, bool) {
	if o == nil || isNil(o.AuctionMode) {
    return nil, false
	}
	return o.AuctionMode, true
}

// HasAuctionMode returns a boolean if a field has been set.
func (o *GetProductResponse) HasAuctionMode() bool {
	if o != nil && !isNil(o.AuctionMode) {
		return true
	}

	return false
}

// SetAuctionMode gets a reference to the given bool and assigns it to the AuctionMode field.
func (o *GetProductResponse) SetAuctionMode(v bool) {
	o.AuctionMode = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *GetProductResponse) GetProductType() string {
	if o == nil || isNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetProductTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProductType) {
    return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *GetProductResponse) HasProductType() bool {
	if o != nil && !isNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *GetProductResponse) SetProductType(v string) {
	o.ProductType = &v
}

// GetQuoteCurrencyId returns the QuoteCurrencyId field value if set, zero value otherwise.
func (o *GetProductResponse) GetQuoteCurrencyId() string {
	if o == nil || isNil(o.QuoteCurrencyId) {
		var ret string
		return ret
	}
	return *o.QuoteCurrencyId
}

// GetQuoteCurrencyIdOk returns a tuple with the QuoteCurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetQuoteCurrencyIdOk() (*string, bool) {
	if o == nil || isNil(o.QuoteCurrencyId) {
    return nil, false
	}
	return o.QuoteCurrencyId, true
}

// HasQuoteCurrencyId returns a boolean if a field has been set.
func (o *GetProductResponse) HasQuoteCurrencyId() bool {
	if o != nil && !isNil(o.QuoteCurrencyId) {
		return true
	}

	return false
}

// SetQuoteCurrencyId gets a reference to the given string and assigns it to the QuoteCurrencyId field.
func (o *GetProductResponse) SetQuoteCurrencyId(v string) {
	o.QuoteCurrencyId = &v
}

// GetBaseCurrencyId returns the BaseCurrencyId field value if set, zero value otherwise.
func (o *GetProductResponse) GetBaseCurrencyId() string {
	if o == nil || isNil(o.BaseCurrencyId) {
		var ret string
		return ret
	}
	return *o.BaseCurrencyId
}

// GetBaseCurrencyIdOk returns a tuple with the BaseCurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetBaseCurrencyIdOk() (*string, bool) {
	if o == nil || isNil(o.BaseCurrencyId) {
    return nil, false
	}
	return o.BaseCurrencyId, true
}

// HasBaseCurrencyId returns a boolean if a field has been set.
func (o *GetProductResponse) HasBaseCurrencyId() bool {
	if o != nil && !isNil(o.BaseCurrencyId) {
		return true
	}

	return false
}

// SetBaseCurrencyId gets a reference to the given string and assigns it to the BaseCurrencyId field.
func (o *GetProductResponse) SetBaseCurrencyId(v string) {
	o.BaseCurrencyId = &v
}

// GetFcmTradingSessionDetails returns the FcmTradingSessionDetails field value if set, zero value otherwise.
func (o *GetProductResponse) GetFcmTradingSessionDetails() string {
	if o == nil || isNil(o.FcmTradingSessionDetails) {
		var ret string
		return ret
	}
	return *o.FcmTradingSessionDetails
}

// GetFcmTradingSessionDetailsOk returns a tuple with the FcmTradingSessionDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetFcmTradingSessionDetailsOk() (*string, bool) {
	if o == nil || isNil(o.FcmTradingSessionDetails) {
    return nil, false
	}
	return o.FcmTradingSessionDetails, true
}

// HasFcmTradingSessionDetails returns a boolean if a field has been set.
func (o *GetProductResponse) HasFcmTradingSessionDetails() bool {
	if o != nil && !isNil(o.FcmTradingSessionDetails) {
		return true
	}

	return false
}

// SetFcmTradingSessionDetails gets a reference to the given string and assigns it to the FcmTradingSessionDetails field.
func (o *GetProductResponse) SetFcmTradingSessionDetails(v string) {
	o.FcmTradingSessionDetails = &v
}

// GetMidMarketPrice returns the MidMarketPrice field value if set, zero value otherwise.
func (o *GetProductResponse) GetMidMarketPrice() string {
	if o == nil || isNil(o.MidMarketPrice) {
		var ret string
		return ret
	}
	return *o.MidMarketPrice
}

// GetMidMarketPriceOk returns a tuple with the MidMarketPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetMidMarketPriceOk() (*string, bool) {
	if o == nil || isNil(o.MidMarketPrice) {
    return nil, false
	}
	return o.MidMarketPrice, true
}

// HasMidMarketPrice returns a boolean if a field has been set.
func (o *GetProductResponse) HasMidMarketPrice() bool {
	if o != nil && !isNil(o.MidMarketPrice) {
		return true
	}

	return false
}

// SetMidMarketPrice gets a reference to the given string and assigns it to the MidMarketPrice field.
func (o *GetProductResponse) SetMidMarketPrice(v string) {
	o.MidMarketPrice = &v
}

// GetBaseDisplaySymbol returns the BaseDisplaySymbol field value if set, zero value otherwise.
func (o *GetProductResponse) GetBaseDisplaySymbol() string {
	if o == nil || isNil(o.BaseDisplaySymbol) {
		var ret string
		return ret
	}
	return *o.BaseDisplaySymbol
}

// GetBaseDisplaySymbolOk returns a tuple with the BaseDisplaySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetBaseDisplaySymbolOk() (*string, bool) {
	if o == nil || isNil(o.BaseDisplaySymbol) {
    return nil, false
	}
	return o.BaseDisplaySymbol, true
}

// HasBaseDisplaySymbol returns a boolean if a field has been set.
func (o *GetProductResponse) HasBaseDisplaySymbol() bool {
	if o != nil && !isNil(o.BaseDisplaySymbol) {
		return true
	}

	return false
}

// SetBaseDisplaySymbol gets a reference to the given string and assigns it to the BaseDisplaySymbol field.
func (o *GetProductResponse) SetBaseDisplaySymbol(v string) {
	o.BaseDisplaySymbol = &v
}

// GetQuoteDisplaySymbol returns the QuoteDisplaySymbol field value if set, zero value otherwise.
func (o *GetProductResponse) GetQuoteDisplaySymbol() string {
	if o == nil || isNil(o.QuoteDisplaySymbol) {
		var ret string
		return ret
	}
	return *o.QuoteDisplaySymbol
}

// GetQuoteDisplaySymbolOk returns a tuple with the QuoteDisplaySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetQuoteDisplaySymbolOk() (*string, bool) {
	if o == nil || isNil(o.QuoteDisplaySymbol) {
    return nil, false
	}
	return o.QuoteDisplaySymbol, true
}

// HasQuoteDisplaySymbol returns a boolean if a field has been set.
func (o *GetProductResponse) HasQuoteDisplaySymbol() bool {
	if o != nil && !isNil(o.QuoteDisplaySymbol) {
		return true
	}

	return false
}

// SetQuoteDisplaySymbol gets a reference to the given string and assigns it to the QuoteDisplaySymbol field.
func (o *GetProductResponse) SetQuoteDisplaySymbol(v string) {
	o.QuoteDisplaySymbol = &v
}

func (o GetProductResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !isNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !isNil(o.PricePercentageChange24h) {
		toSerialize["price_percentage_change_24h"] = o.PricePercentageChange24h
	}
	if !isNil(o.Volume24h) {
		toSerialize["volume_24h"] = o.Volume24h
	}
	if !isNil(o.VolumePercentageChange24h) {
		toSerialize["volume_percentage_change_24h"] = o.VolumePercentageChange24h
	}
	if !isNil(o.BaseIncrement) {
		toSerialize["base_increment"] = o.BaseIncrement
	}
	if !isNil(o.QuoteIncrement) {
		toSerialize["quote_increment"] = o.QuoteIncrement
	}
	if !isNil(o.QuoteMinSize) {
		toSerialize["quote_min_size"] = o.QuoteMinSize
	}
	if !isNil(o.QuoteMaxSize) {
		toSerialize["quote_max_size"] = o.QuoteMaxSize
	}
	if !isNil(o.BaseMinSize) {
		toSerialize["base_min_size"] = o.BaseMinSize
	}
	if !isNil(o.BaseMaxSize) {
		toSerialize["base_max_size"] = o.BaseMaxSize
	}
	if !isNil(o.BaseName) {
		toSerialize["base_name"] = o.BaseName
	}
	if !isNil(o.QuoteName) {
		toSerialize["quote_name"] = o.QuoteName
	}
	if !isNil(o.Watched) {
		toSerialize["watched"] = o.Watched
	}
	if !isNil(o.IsDisabled) {
		toSerialize["is_disabled"] = o.IsDisabled
	}
	if !isNil(o.New) {
		toSerialize["new"] = o.New
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.CancelOnly) {
		toSerialize["cancel_only"] = o.CancelOnly
	}
	if !isNil(o.LimitOnly) {
		toSerialize["limit_only"] = o.LimitOnly
	}
	if !isNil(o.PostOnly) {
		toSerialize["post_only"] = o.PostOnly
	}
	if !isNil(o.TradingDisabled) {
		toSerialize["trading_disabled"] = o.TradingDisabled
	}
	if !isNil(o.AuctionMode) {
		toSerialize["auction_mode"] = o.AuctionMode
	}
	if !isNil(o.ProductType) {
		toSerialize["product_type"] = o.ProductType
	}
	if !isNil(o.QuoteCurrencyId) {
		toSerialize["quote_currency_id"] = o.QuoteCurrencyId
	}
	if !isNil(o.BaseCurrencyId) {
		toSerialize["base_currency_id"] = o.BaseCurrencyId
	}
	if !isNil(o.FcmTradingSessionDetails) {
		toSerialize["fcm_trading_session_details"] = o.FcmTradingSessionDetails
	}
	if !isNil(o.MidMarketPrice) {
		toSerialize["mid_market_price"] = o.MidMarketPrice
	}
	if !isNil(o.BaseDisplaySymbol) {
		toSerialize["base_display_symbol"] = o.BaseDisplaySymbol
	}
	if !isNil(o.QuoteDisplaySymbol) {
		toSerialize["quote_display_symbol"] = o.QuoteDisplaySymbol
	}
	return json.Marshal(toSerialize)
}

type NullableGetProductResponse struct {
	value *GetProductResponse
	isSet bool
}

func (v NullableGetProductResponse) Get() *GetProductResponse {
	return v.value
}

func (v *NullableGetProductResponse) Set(val *GetProductResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductResponse(val *GetProductResponse) *NullableGetProductResponse {
	return &NullableGetProductResponse{value: val, isSet: true}
}

func (v NullableGetProductResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


