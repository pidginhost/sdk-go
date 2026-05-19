# FundsBalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | **float64** |  | 
**ThresholdType** | **string** |  | 
**ThresholdAmount** | **NullableFloat64** |  | 
**ThresholdDays** | **NullableInt32** |  | 

## Methods

### NewFundsBalanceResponse

`func NewFundsBalanceResponse(balance float64, thresholdType string, thresholdAmount NullableFloat64, thresholdDays NullableInt32, ) *FundsBalanceResponse`

NewFundsBalanceResponse instantiates a new FundsBalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundsBalanceResponseWithDefaults

`func NewFundsBalanceResponseWithDefaults() *FundsBalanceResponse`

NewFundsBalanceResponseWithDefaults instantiates a new FundsBalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *FundsBalanceResponse) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *FundsBalanceResponse) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *FundsBalanceResponse) SetBalance(v float64)`

SetBalance sets Balance field to given value.


### GetThresholdType

`func (o *FundsBalanceResponse) GetThresholdType() string`

GetThresholdType returns the ThresholdType field if non-nil, zero value otherwise.

### GetThresholdTypeOk

`func (o *FundsBalanceResponse) GetThresholdTypeOk() (*string, bool)`

GetThresholdTypeOk returns a tuple with the ThresholdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdType

`func (o *FundsBalanceResponse) SetThresholdType(v string)`

SetThresholdType sets ThresholdType field to given value.


### GetThresholdAmount

`func (o *FundsBalanceResponse) GetThresholdAmount() float64`

GetThresholdAmount returns the ThresholdAmount field if non-nil, zero value otherwise.

### GetThresholdAmountOk

`func (o *FundsBalanceResponse) GetThresholdAmountOk() (*float64, bool)`

GetThresholdAmountOk returns a tuple with the ThresholdAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdAmount

`func (o *FundsBalanceResponse) SetThresholdAmount(v float64)`

SetThresholdAmount sets ThresholdAmount field to given value.


### SetThresholdAmountNil

`func (o *FundsBalanceResponse) SetThresholdAmountNil(b bool)`

 SetThresholdAmountNil sets the value for ThresholdAmount to be an explicit nil

### UnsetThresholdAmount
`func (o *FundsBalanceResponse) UnsetThresholdAmount()`

UnsetThresholdAmount ensures that no value is present for ThresholdAmount, not even an explicit nil
### GetThresholdDays

`func (o *FundsBalanceResponse) GetThresholdDays() int32`

GetThresholdDays returns the ThresholdDays field if non-nil, zero value otherwise.

### GetThresholdDaysOk

`func (o *FundsBalanceResponse) GetThresholdDaysOk() (*int32, bool)`

GetThresholdDaysOk returns a tuple with the ThresholdDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdDays

`func (o *FundsBalanceResponse) SetThresholdDays(v int32)`

SetThresholdDays sets ThresholdDays field to given value.


### SetThresholdDaysNil

`func (o *FundsBalanceResponse) SetThresholdDaysNil(b bool)`

 SetThresholdDaysNil sets the value for ThresholdDays to be an explicit nil

### UnsetThresholdDays
`func (o *FundsBalanceResponse) UnsetThresholdDays()`

UnsetThresholdDays ensures that no value is present for ThresholdDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


