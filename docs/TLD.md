# TLD

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Tld** | **string** |  | 
**Price** | **float64** | Euro without TVA | 
**RenewPrice** | Pointer to **NullableFloat64** | Euro without TVA | [optional] 
**TransferPrice** | Pointer to **NullableFloat64** | Euro without TVA | [optional] 
**Registrar** | **string** |  | 
**ApiRenewable** | Pointer to **bool** |  | [optional] 
**IdnaSupport** | Pointer to **bool** |  | [optional] 
**WgSupport** | Pointer to **bool** |  | [optional] 
**ReactivateMaxDays** | Pointer to **int32** |  | [optional] 

## Methods

### NewTLD

`func NewTLD(id int32, tld string, price float64, registrar string, ) *TLD`

NewTLD instantiates a new TLD object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLDWithDefaults

`func NewTLDWithDefaults() *TLD`

NewTLDWithDefaults instantiates a new TLD object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TLD) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TLD) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TLD) SetId(v int32)`

SetId sets Id field to given value.


### GetTld

`func (o *TLD) GetTld() string`

GetTld returns the Tld field if non-nil, zero value otherwise.

### GetTldOk

`func (o *TLD) GetTldOk() (*string, bool)`

GetTldOk returns a tuple with the Tld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTld

`func (o *TLD) SetTld(v string)`

SetTld sets Tld field to given value.


### GetPrice

`func (o *TLD) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TLD) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *TLD) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetRenewPrice

`func (o *TLD) GetRenewPrice() float64`

GetRenewPrice returns the RenewPrice field if non-nil, zero value otherwise.

### GetRenewPriceOk

`func (o *TLD) GetRenewPriceOk() (*float64, bool)`

GetRenewPriceOk returns a tuple with the RenewPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewPrice

`func (o *TLD) SetRenewPrice(v float64)`

SetRenewPrice sets RenewPrice field to given value.

### HasRenewPrice

`func (o *TLD) HasRenewPrice() bool`

HasRenewPrice returns a boolean if a field has been set.

### SetRenewPriceNil

`func (o *TLD) SetRenewPriceNil(b bool)`

 SetRenewPriceNil sets the value for RenewPrice to be an explicit nil

### UnsetRenewPrice
`func (o *TLD) UnsetRenewPrice()`

UnsetRenewPrice ensures that no value is present for RenewPrice, not even an explicit nil
### GetTransferPrice

`func (o *TLD) GetTransferPrice() float64`

GetTransferPrice returns the TransferPrice field if non-nil, zero value otherwise.

### GetTransferPriceOk

`func (o *TLD) GetTransferPriceOk() (*float64, bool)`

GetTransferPriceOk returns a tuple with the TransferPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferPrice

`func (o *TLD) SetTransferPrice(v float64)`

SetTransferPrice sets TransferPrice field to given value.

### HasTransferPrice

`func (o *TLD) HasTransferPrice() bool`

HasTransferPrice returns a boolean if a field has been set.

### SetTransferPriceNil

`func (o *TLD) SetTransferPriceNil(b bool)`

 SetTransferPriceNil sets the value for TransferPrice to be an explicit nil

### UnsetTransferPrice
`func (o *TLD) UnsetTransferPrice()`

UnsetTransferPrice ensures that no value is present for TransferPrice, not even an explicit nil
### GetRegistrar

`func (o *TLD) GetRegistrar() string`

GetRegistrar returns the Registrar field if non-nil, zero value otherwise.

### GetRegistrarOk

`func (o *TLD) GetRegistrarOk() (*string, bool)`

GetRegistrarOk returns a tuple with the Registrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrar

`func (o *TLD) SetRegistrar(v string)`

SetRegistrar sets Registrar field to given value.


### GetApiRenewable

`func (o *TLD) GetApiRenewable() bool`

GetApiRenewable returns the ApiRenewable field if non-nil, zero value otherwise.

### GetApiRenewableOk

`func (o *TLD) GetApiRenewableOk() (*bool, bool)`

GetApiRenewableOk returns a tuple with the ApiRenewable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiRenewable

`func (o *TLD) SetApiRenewable(v bool)`

SetApiRenewable sets ApiRenewable field to given value.

### HasApiRenewable

`func (o *TLD) HasApiRenewable() bool`

HasApiRenewable returns a boolean if a field has been set.

### GetIdnaSupport

`func (o *TLD) GetIdnaSupport() bool`

GetIdnaSupport returns the IdnaSupport field if non-nil, zero value otherwise.

### GetIdnaSupportOk

`func (o *TLD) GetIdnaSupportOk() (*bool, bool)`

GetIdnaSupportOk returns a tuple with the IdnaSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdnaSupport

`func (o *TLD) SetIdnaSupport(v bool)`

SetIdnaSupport sets IdnaSupport field to given value.

### HasIdnaSupport

`func (o *TLD) HasIdnaSupport() bool`

HasIdnaSupport returns a boolean if a field has been set.

### GetWgSupport

`func (o *TLD) GetWgSupport() bool`

GetWgSupport returns the WgSupport field if non-nil, zero value otherwise.

### GetWgSupportOk

`func (o *TLD) GetWgSupportOk() (*bool, bool)`

GetWgSupportOk returns a tuple with the WgSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWgSupport

`func (o *TLD) SetWgSupport(v bool)`

SetWgSupport sets WgSupport field to given value.

### HasWgSupport

`func (o *TLD) HasWgSupport() bool`

HasWgSupport returns a boolean if a field has been set.

### GetReactivateMaxDays

`func (o *TLD) GetReactivateMaxDays() int32`

GetReactivateMaxDays returns the ReactivateMaxDays field if non-nil, zero value otherwise.

### GetReactivateMaxDaysOk

`func (o *TLD) GetReactivateMaxDaysOk() (*int32, bool)`

GetReactivateMaxDaysOk returns a tuple with the ReactivateMaxDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactivateMaxDays

`func (o *TLD) SetReactivateMaxDays(v int32)`

SetReactivateMaxDays sets ReactivateMaxDays field to given value.

### HasReactivateMaxDays

`func (o *TLD) HasReactivateMaxDays() bool`

HasReactivateMaxDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


