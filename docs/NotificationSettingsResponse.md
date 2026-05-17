# NotificationSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**ThresholdType** | **string** |  | 
**ThresholdAmount** | **NullableFloat64** |  | 
**ThresholdDays** | **NullableInt32** |  | 

## Methods

### NewNotificationSettingsResponse

`func NewNotificationSettingsResponse(message string, thresholdType string, thresholdAmount NullableFloat64, thresholdDays NullableInt32, ) *NotificationSettingsResponse`

NewNotificationSettingsResponse instantiates a new NotificationSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSettingsResponseWithDefaults

`func NewNotificationSettingsResponseWithDefaults() *NotificationSettingsResponse`

NewNotificationSettingsResponseWithDefaults instantiates a new NotificationSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *NotificationSettingsResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NotificationSettingsResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NotificationSettingsResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetThresholdType

`func (o *NotificationSettingsResponse) GetThresholdType() string`

GetThresholdType returns the ThresholdType field if non-nil, zero value otherwise.

### GetThresholdTypeOk

`func (o *NotificationSettingsResponse) GetThresholdTypeOk() (*string, bool)`

GetThresholdTypeOk returns a tuple with the ThresholdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdType

`func (o *NotificationSettingsResponse) SetThresholdType(v string)`

SetThresholdType sets ThresholdType field to given value.


### GetThresholdAmount

`func (o *NotificationSettingsResponse) GetThresholdAmount() float64`

GetThresholdAmount returns the ThresholdAmount field if non-nil, zero value otherwise.

### GetThresholdAmountOk

`func (o *NotificationSettingsResponse) GetThresholdAmountOk() (*float64, bool)`

GetThresholdAmountOk returns a tuple with the ThresholdAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdAmount

`func (o *NotificationSettingsResponse) SetThresholdAmount(v float64)`

SetThresholdAmount sets ThresholdAmount field to given value.


### SetThresholdAmountNil

`func (o *NotificationSettingsResponse) SetThresholdAmountNil(b bool)`

 SetThresholdAmountNil sets the value for ThresholdAmount to be an explicit nil

### UnsetThresholdAmount
`func (o *NotificationSettingsResponse) UnsetThresholdAmount()`

UnsetThresholdAmount ensures that no value is present for ThresholdAmount, not even an explicit nil
### GetThresholdDays

`func (o *NotificationSettingsResponse) GetThresholdDays() int32`

GetThresholdDays returns the ThresholdDays field if non-nil, zero value otherwise.

### GetThresholdDaysOk

`func (o *NotificationSettingsResponse) GetThresholdDaysOk() (*int32, bool)`

GetThresholdDaysOk returns a tuple with the ThresholdDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdDays

`func (o *NotificationSettingsResponse) SetThresholdDays(v int32)`

SetThresholdDays sets ThresholdDays field to given value.


### SetThresholdDaysNil

`func (o *NotificationSettingsResponse) SetThresholdDaysNil(b bool)`

 SetThresholdDaysNil sets the value for ThresholdDays to be an explicit nil

### UnsetThresholdDays
`func (o *NotificationSettingsResponse) UnsetThresholdDays()`

UnsetThresholdDays ensures that no value is present for ThresholdDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


