# TLDCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tlds** | [**[]TLD**](TLD.md) |  | 
**DefaultNameservers** | **string** |  | 
**Currency** | **string** |  | 

## Methods

### NewTLDCatalog

`func NewTLDCatalog(tlds []TLD, defaultNameservers string, currency string, ) *TLDCatalog`

NewTLDCatalog instantiates a new TLDCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLDCatalogWithDefaults

`func NewTLDCatalogWithDefaults() *TLDCatalog`

NewTLDCatalogWithDefaults instantiates a new TLDCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTlds

`func (o *TLDCatalog) GetTlds() []TLD`

GetTlds returns the Tlds field if non-nil, zero value otherwise.

### GetTldsOk

`func (o *TLDCatalog) GetTldsOk() (*[]TLD, bool)`

GetTldsOk returns a tuple with the Tlds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlds

`func (o *TLDCatalog) SetTlds(v []TLD)`

SetTlds sets Tlds field to given value.


### GetDefaultNameservers

`func (o *TLDCatalog) GetDefaultNameservers() string`

GetDefaultNameservers returns the DefaultNameservers field if non-nil, zero value otherwise.

### GetDefaultNameserversOk

`func (o *TLDCatalog) GetDefaultNameserversOk() (*string, bool)`

GetDefaultNameserversOk returns a tuple with the DefaultNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNameservers

`func (o *TLDCatalog) SetDefaultNameservers(v string)`

SetDefaultNameservers sets DefaultNameservers field to given value.


### GetCurrency

`func (o *TLDCatalog) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TLDCatalog) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TLDCatalog) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


