# ActivityLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Date** | **string** |  | 

## Methods

### NewActivityLogEntry

`func NewActivityLogEntry(message string, date string, ) *ActivityLogEntry`

NewActivityLogEntry instantiates a new ActivityLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityLogEntryWithDefaults

`func NewActivityLogEntryWithDefaults() *ActivityLogEntry`

NewActivityLogEntryWithDefaults instantiates a new ActivityLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ActivityLogEntry) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ActivityLogEntry) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ActivityLogEntry) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDate

`func (o *ActivityLogEntry) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ActivityLogEntry) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ActivityLogEntry) SetDate(v string)`

SetDate sets Date field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


