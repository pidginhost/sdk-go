# ClusterEncryptionRefusal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | [**EncryptionReasonCodeEnum**](EncryptionReasonCodeEnum.md) |  | 
**Message** | **string** |  | 

## Methods

### NewClusterEncryptionRefusal

`func NewClusterEncryptionRefusal(reason EncryptionReasonCodeEnum, message string, ) *ClusterEncryptionRefusal`

NewClusterEncryptionRefusal instantiates a new ClusterEncryptionRefusal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterEncryptionRefusalWithDefaults

`func NewClusterEncryptionRefusalWithDefaults() *ClusterEncryptionRefusal`

NewClusterEncryptionRefusalWithDefaults instantiates a new ClusterEncryptionRefusal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *ClusterEncryptionRefusal) GetReason() EncryptionReasonCodeEnum`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ClusterEncryptionRefusal) GetReasonOk() (*EncryptionReasonCodeEnum, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ClusterEncryptionRefusal) SetReason(v EncryptionReasonCodeEnum)`

SetReason sets Reason field to given value.


### GetMessage

`func (o *ClusterEncryptionRefusal) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterEncryptionRefusal) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterEncryptionRefusal) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


