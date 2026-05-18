# LowBalanceSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThresholdType** | [**ThresholdTypeEnum**](ThresholdTypeEnum.md) |  | 
**ThresholdAmount** | Pointer to **NullableFloat64** |  | [optional] 
**ThresholdDays** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewLowBalanceSettings

`func NewLowBalanceSettings(thresholdType ThresholdTypeEnum, ) *LowBalanceSettings`

NewLowBalanceSettings instantiates a new LowBalanceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLowBalanceSettingsWithDefaults

`func NewLowBalanceSettingsWithDefaults() *LowBalanceSettings`

NewLowBalanceSettingsWithDefaults instantiates a new LowBalanceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThresholdType

`func (o *LowBalanceSettings) GetThresholdType() ThresholdTypeEnum`

GetThresholdType returns the ThresholdType field if non-nil, zero value otherwise.

### GetThresholdTypeOk

`func (o *LowBalanceSettings) GetThresholdTypeOk() (*ThresholdTypeEnum, bool)`

GetThresholdTypeOk returns a tuple with the ThresholdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdType

`func (o *LowBalanceSettings) SetThresholdType(v ThresholdTypeEnum)`

SetThresholdType sets ThresholdType field to given value.


### GetThresholdAmount

`func (o *LowBalanceSettings) GetThresholdAmount() float64`

GetThresholdAmount returns the ThresholdAmount field if non-nil, zero value otherwise.

### GetThresholdAmountOk

`func (o *LowBalanceSettings) GetThresholdAmountOk() (*float64, bool)`

GetThresholdAmountOk returns a tuple with the ThresholdAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdAmount

`func (o *LowBalanceSettings) SetThresholdAmount(v float64)`

SetThresholdAmount sets ThresholdAmount field to given value.

### HasThresholdAmount

`func (o *LowBalanceSettings) HasThresholdAmount() bool`

HasThresholdAmount returns a boolean if a field has been set.

### SetThresholdAmountNil

`func (o *LowBalanceSettings) SetThresholdAmountNil(b bool)`

 SetThresholdAmountNil sets the value for ThresholdAmount to be an explicit nil

### UnsetThresholdAmount
`func (o *LowBalanceSettings) UnsetThresholdAmount()`

UnsetThresholdAmount ensures that no value is present for ThresholdAmount, not even an explicit nil
### GetThresholdDays

`func (o *LowBalanceSettings) GetThresholdDays() int32`

GetThresholdDays returns the ThresholdDays field if non-nil, zero value otherwise.

### GetThresholdDaysOk

`func (o *LowBalanceSettings) GetThresholdDaysOk() (*int32, bool)`

GetThresholdDaysOk returns a tuple with the ThresholdDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdDays

`func (o *LowBalanceSettings) SetThresholdDays(v int32)`

SetThresholdDays sets ThresholdDays field to given value.

### HasThresholdDays

`func (o *LowBalanceSettings) HasThresholdDays() bool`

HasThresholdDays returns a boolean if a field has been set.

### SetThresholdDaysNil

`func (o *LowBalanceSettings) SetThresholdDaysNil(b bool)`

 SetThresholdDaysNil sets the value for ThresholdDays to be an explicit nil

### UnsetThresholdDays
`func (o *LowBalanceSettings) UnsetThresholdDays()`

UnsetThresholdDays ensures that no value is present for ThresholdDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


