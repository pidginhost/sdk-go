# Deposit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Status** | [**DepositStatusEnum**](DepositStatusEnum.md) |  | [readonly] 
**Amount** | **string** | Fara TVA | [readonly] 
**VatValue** | **string** | TVA | [readonly] 
**VatPercentage** | **int32** |  | [readonly] 
**Total** | **string** | Cu TVA | [readonly] 
**Created** | **string** |  | [readonly] 

## Methods

### NewDeposit

`func NewDeposit(id int32, status DepositStatusEnum, amount string, vatValue string, vatPercentage int32, total string, created string, ) *Deposit`

NewDeposit instantiates a new Deposit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositWithDefaults

`func NewDepositWithDefaults() *Deposit`

NewDepositWithDefaults instantiates a new Deposit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Deposit) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Deposit) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Deposit) SetId(v int32)`

SetId sets Id field to given value.


### GetStatus

`func (o *Deposit) GetStatus() DepositStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Deposit) GetStatusOk() (*DepositStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Deposit) SetStatus(v DepositStatusEnum)`

SetStatus sets Status field to given value.


### GetAmount

`func (o *Deposit) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Deposit) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Deposit) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetVatValue

`func (o *Deposit) GetVatValue() string`

GetVatValue returns the VatValue field if non-nil, zero value otherwise.

### GetVatValueOk

`func (o *Deposit) GetVatValueOk() (*string, bool)`

GetVatValueOk returns a tuple with the VatValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatValue

`func (o *Deposit) SetVatValue(v string)`

SetVatValue sets VatValue field to given value.


### GetVatPercentage

`func (o *Deposit) GetVatPercentage() int32`

GetVatPercentage returns the VatPercentage field if non-nil, zero value otherwise.

### GetVatPercentageOk

`func (o *Deposit) GetVatPercentageOk() (*int32, bool)`

GetVatPercentageOk returns a tuple with the VatPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatPercentage

`func (o *Deposit) SetVatPercentage(v int32)`

SetVatPercentage sets VatPercentage field to given value.


### GetTotal

`func (o *Deposit) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Deposit) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Deposit) SetTotal(v string)`

SetTotal sets Total field to given value.


### GetCreated

`func (o *Deposit) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Deposit) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Deposit) SetCreated(v string)`

SetCreated sets Created field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


