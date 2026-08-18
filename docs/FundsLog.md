# FundsLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Date** | **string** |  | [readonly] 
**Operation** | [**OperationEnum**](OperationEnum.md) |  | [readonly] 
**Amount** | **string** | + or - sum | [readonly] 
**Balance** | **string** | Balance after operation | [readonly] 
**ForDate** | **NullableString** | Used for cloud service payments | [readonly] 

## Methods

### NewFundsLog

`func NewFundsLog(id int32, date string, operation OperationEnum, amount string, balance string, forDate NullableString, ) *FundsLog`

NewFundsLog instantiates a new FundsLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundsLogWithDefaults

`func NewFundsLogWithDefaults() *FundsLog`

NewFundsLogWithDefaults instantiates a new FundsLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FundsLog) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FundsLog) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FundsLog) SetId(v int32)`

SetId sets Id field to given value.


### GetDate

`func (o *FundsLog) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *FundsLog) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *FundsLog) SetDate(v string)`

SetDate sets Date field to given value.


### GetOperation

`func (o *FundsLog) GetOperation() OperationEnum`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *FundsLog) GetOperationOk() (*OperationEnum, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *FundsLog) SetOperation(v OperationEnum)`

SetOperation sets Operation field to given value.


### GetAmount

`func (o *FundsLog) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FundsLog) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FundsLog) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetBalance

`func (o *FundsLog) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *FundsLog) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *FundsLog) SetBalance(v string)`

SetBalance sets Balance field to given value.


### GetForDate

`func (o *FundsLog) GetForDate() string`

GetForDate returns the ForDate field if non-nil, zero value otherwise.

### GetForDateOk

`func (o *FundsLog) GetForDateOk() (*string, bool)`

GetForDateOk returns a tuple with the ForDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForDate

`func (o *FundsLog) SetForDate(v string)`

SetForDate sets ForDate field to given value.


### SetForDateNil

`func (o *FundsLog) SetForDateNil(b bool)`

 SetForDateNil sets the value for ForDate to be an explicit nil

### UnsetForDate
`func (o *FundsLog) UnsetForDate()`

UnsetForDate ensures that no value is present for ForDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


