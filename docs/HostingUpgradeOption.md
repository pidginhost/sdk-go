# HostingUpgradeOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**BasePrice** | **float64** |  | 
**Currency** | **string** |  | 
**TaxIncluded** | **bool** |  | 
**Limits** | **map[string]interface{}** | Package display limits; disk_quota and traffic are in GB. | 
**ReviewUrl** | **string** |  | 

## Methods

### NewHostingUpgradeOption

`func NewHostingUpgradeOption(id int32, name string, basePrice float64, currency string, taxIncluded bool, limits map[string]interface{}, reviewUrl string, ) *HostingUpgradeOption`

NewHostingUpgradeOption instantiates a new HostingUpgradeOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostingUpgradeOptionWithDefaults

`func NewHostingUpgradeOptionWithDefaults() *HostingUpgradeOption`

NewHostingUpgradeOptionWithDefaults instantiates a new HostingUpgradeOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HostingUpgradeOption) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostingUpgradeOption) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostingUpgradeOption) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *HostingUpgradeOption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HostingUpgradeOption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HostingUpgradeOption) SetName(v string)`

SetName sets Name field to given value.


### GetBasePrice

`func (o *HostingUpgradeOption) GetBasePrice() float64`

GetBasePrice returns the BasePrice field if non-nil, zero value otherwise.

### GetBasePriceOk

`func (o *HostingUpgradeOption) GetBasePriceOk() (*float64, bool)`

GetBasePriceOk returns a tuple with the BasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePrice

`func (o *HostingUpgradeOption) SetBasePrice(v float64)`

SetBasePrice sets BasePrice field to given value.


### GetCurrency

`func (o *HostingUpgradeOption) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *HostingUpgradeOption) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *HostingUpgradeOption) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTaxIncluded

`func (o *HostingUpgradeOption) GetTaxIncluded() bool`

GetTaxIncluded returns the TaxIncluded field if non-nil, zero value otherwise.

### GetTaxIncludedOk

`func (o *HostingUpgradeOption) GetTaxIncludedOk() (*bool, bool)`

GetTaxIncludedOk returns a tuple with the TaxIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIncluded

`func (o *HostingUpgradeOption) SetTaxIncluded(v bool)`

SetTaxIncluded sets TaxIncluded field to given value.


### GetLimits

`func (o *HostingUpgradeOption) GetLimits() map[string]interface{}`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *HostingUpgradeOption) GetLimitsOk() (*map[string]interface{}, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *HostingUpgradeOption) SetLimits(v map[string]interface{})`

SetLimits sets Limits field to given value.


### GetReviewUrl

`func (o *HostingUpgradeOption) GetReviewUrl() string`

GetReviewUrl returns the ReviewUrl field if non-nil, zero value otherwise.

### GetReviewUrlOk

`func (o *HostingUpgradeOption) GetReviewUrlOk() (*string, bool)`

GetReviewUrlOk returns a tuple with the ReviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewUrl

`func (o *HostingUpgradeOption) SetReviewUrl(v string)`

SetReviewUrl sets ReviewUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


