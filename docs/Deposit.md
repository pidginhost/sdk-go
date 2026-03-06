# Deposit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Status** | [**DepositStatusEnum**](DepositStatusEnum.md) |  | [readonly] 
**Amount** | **float64** | Fara TVA | [readonly] 
**VatValue** | **float64** | TVA | [readonly] 
**VatPercentage** | **int32** |  | [readonly] 
**Total** | **float64** | Cu TVA | [readonly] 
**Created** | **time.Time** |  | [readonly] 

## Methods

### NewDeposit

`func NewDeposit(id int32, status DepositStatusEnum, amount float64, vatValue float64, vatPercentage int32, total float64, created time.Time, ) *Deposit`

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

`func (o *Deposit) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Deposit) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Deposit) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetVatValue

`func (o *Deposit) GetVatValue() float64`

GetVatValue returns the VatValue field if non-nil, zero value otherwise.

### GetVatValueOk

`func (o *Deposit) GetVatValueOk() (*float64, bool)`

GetVatValueOk returns a tuple with the VatValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatValue

`func (o *Deposit) SetVatValue(v float64)`

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

`func (o *Deposit) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Deposit) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Deposit) SetTotal(v float64)`

SetTotal sets Total field to given value.


### GetCreated

`func (o *Deposit) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Deposit) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Deposit) SetCreated(v time.Time)`

SetCreated sets Created field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


