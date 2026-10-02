# DedicatedServerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**StatusText** | **string** |  | 

## Methods

### NewDedicatedServerStatus

`func NewDedicatedServerStatus(statusText string, ) *DedicatedServerStatus`

NewDedicatedServerStatus instantiates a new DedicatedServerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedServerStatusWithDefaults

`func NewDedicatedServerStatusWithDefaults() *DedicatedServerStatus`

NewDedicatedServerStatusWithDefaults instantiates a new DedicatedServerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DedicatedServerStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DedicatedServerStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DedicatedServerStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DedicatedServerStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusText

`func (o *DedicatedServerStatus) GetStatusText() string`

GetStatusText returns the StatusText field if non-nil, zero value otherwise.

### GetStatusTextOk

`func (o *DedicatedServerStatus) GetStatusTextOk() (*string, bool)`

GetStatusTextOk returns a tuple with the StatusText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusText

`func (o *DedicatedServerStatus) SetStatusText(v string)`

SetStatusText sets StatusText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


