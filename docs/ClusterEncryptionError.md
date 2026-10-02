# ClusterEncryptionError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Reason** | Pointer to [**EncryptionReasonCodeEnum**](EncryptionReasonCodeEnum.md) |  | [optional] 
**Extra** | Pointer to **map[string]interface{}** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 

## Methods

### NewClusterEncryptionError

`func NewClusterEncryptionError(message string, ) *ClusterEncryptionError`

NewClusterEncryptionError instantiates a new ClusterEncryptionError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterEncryptionErrorWithDefaults

`func NewClusterEncryptionErrorWithDefaults() *ClusterEncryptionError`

NewClusterEncryptionErrorWithDefaults instantiates a new ClusterEncryptionError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ClusterEncryptionError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterEncryptionError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterEncryptionError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetReason

`func (o *ClusterEncryptionError) GetReason() EncryptionReasonCodeEnum`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ClusterEncryptionError) GetReasonOk() (*EncryptionReasonCodeEnum, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ClusterEncryptionError) SetReason(v EncryptionReasonCodeEnum)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ClusterEncryptionError) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetExtra

`func (o *ClusterEncryptionError) GetExtra() map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *ClusterEncryptionError) GetExtraOk() (*map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *ClusterEncryptionError) SetExtra(v map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *ClusterEncryptionError) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetCode

`func (o *ClusterEncryptionError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ClusterEncryptionError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ClusterEncryptionError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ClusterEncryptionError) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


