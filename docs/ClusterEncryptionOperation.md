# ClusterEncryptionOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Kind** | **string** |  | [readonly] 
**Status** | **string** |  | [readonly] 
**RequestedMode** | **string** |  | [readonly] 
**PreviousMode** | **string** |  | [readonly] 
**Reason** | **string** |  | [readonly] 
**Message** | **string** |  | [readonly] 
**OverrideUnverifiable** | **bool** |  | [readonly] 
**RequestId** | **string** |  | [readonly] 
**CreatedAt** | **string** |  | [readonly] 
**FinishedAt** | **NullableString** |  | [readonly] 

## Methods

### NewClusterEncryptionOperation

`func NewClusterEncryptionOperation(id int32, kind string, status string, requestedMode string, previousMode string, reason string, message string, overrideUnverifiable bool, requestId string, createdAt string, finishedAt NullableString, ) *ClusterEncryptionOperation`

NewClusterEncryptionOperation instantiates a new ClusterEncryptionOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterEncryptionOperationWithDefaults

`func NewClusterEncryptionOperationWithDefaults() *ClusterEncryptionOperation`

NewClusterEncryptionOperationWithDefaults instantiates a new ClusterEncryptionOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterEncryptionOperation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterEncryptionOperation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterEncryptionOperation) SetId(v int32)`

SetId sets Id field to given value.


### GetKind

`func (o *ClusterEncryptionOperation) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterEncryptionOperation) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterEncryptionOperation) SetKind(v string)`

SetKind sets Kind field to given value.


### GetStatus

`func (o *ClusterEncryptionOperation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterEncryptionOperation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterEncryptionOperation) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRequestedMode

`func (o *ClusterEncryptionOperation) GetRequestedMode() string`

GetRequestedMode returns the RequestedMode field if non-nil, zero value otherwise.

### GetRequestedModeOk

`func (o *ClusterEncryptionOperation) GetRequestedModeOk() (*string, bool)`

GetRequestedModeOk returns a tuple with the RequestedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedMode

`func (o *ClusterEncryptionOperation) SetRequestedMode(v string)`

SetRequestedMode sets RequestedMode field to given value.


### GetPreviousMode

`func (o *ClusterEncryptionOperation) GetPreviousMode() string`

GetPreviousMode returns the PreviousMode field if non-nil, zero value otherwise.

### GetPreviousModeOk

`func (o *ClusterEncryptionOperation) GetPreviousModeOk() (*string, bool)`

GetPreviousModeOk returns a tuple with the PreviousMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousMode

`func (o *ClusterEncryptionOperation) SetPreviousMode(v string)`

SetPreviousMode sets PreviousMode field to given value.


### GetReason

`func (o *ClusterEncryptionOperation) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ClusterEncryptionOperation) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ClusterEncryptionOperation) SetReason(v string)`

SetReason sets Reason field to given value.


### GetMessage

`func (o *ClusterEncryptionOperation) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterEncryptionOperation) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterEncryptionOperation) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOverrideUnverifiable

`func (o *ClusterEncryptionOperation) GetOverrideUnverifiable() bool`

GetOverrideUnverifiable returns the OverrideUnverifiable field if non-nil, zero value otherwise.

### GetOverrideUnverifiableOk

`func (o *ClusterEncryptionOperation) GetOverrideUnverifiableOk() (*bool, bool)`

GetOverrideUnverifiableOk returns a tuple with the OverrideUnverifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideUnverifiable

`func (o *ClusterEncryptionOperation) SetOverrideUnverifiable(v bool)`

SetOverrideUnverifiable sets OverrideUnverifiable field to given value.


### GetRequestId

`func (o *ClusterEncryptionOperation) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ClusterEncryptionOperation) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ClusterEncryptionOperation) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetCreatedAt

`func (o *ClusterEncryptionOperation) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClusterEncryptionOperation) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClusterEncryptionOperation) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetFinishedAt

`func (o *ClusterEncryptionOperation) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ClusterEncryptionOperation) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ClusterEncryptionOperation) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.


### SetFinishedAtNil

`func (o *ClusterEncryptionOperation) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *ClusterEncryptionOperation) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


