# NodeOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Kind** | [**NodeOperationKindEnum**](NodeOperationKindEnum.md) |  | [readonly] 
**Source** | [**NodeOperationSourceEnum**](NodeOperationSourceEnum.md) |  | [readonly] 
**TargetHostname** | **string** |  | [readonly] 
**Status** | [**NodeOperationStatusEnum**](NodeOperationStatusEnum.md) |  | [readonly] 
**Reason** | **string** |  | [readonly] 
**Message** | **string** |  | [readonly] 
**BypassPdb** | **bool** |  | [readonly] 
**DeleteUnmanagedPods** | **bool** |  | [readonly] 
**LocalDataLossAccepted** | **bool** |  | [readonly] 
**BypassPdbConfirmedAt** | **NullableString** |  | [readonly] 
**UnmanagedPodsConfirmedAt** | **NullableString** |  | [readonly] 
**ActorLabel** | **string** | Who requested the operation (user email or staff name). Never token material. | [readonly] 
**CreatedAt** | **string** |  | [readonly] 
**UpdatedAt** | **string** |  | [readonly] 
**FinishedAt** | **NullableString** |  | [readonly] 
**AllowedActions** | **[]string** |  | [readonly] 

## Methods

### NewNodeOperation

`func NewNodeOperation(id int32, kind NodeOperationKindEnum, source NodeOperationSourceEnum, targetHostname string, status NodeOperationStatusEnum, reason string, message string, bypassPdb bool, deleteUnmanagedPods bool, localDataLossAccepted bool, bypassPdbConfirmedAt NullableString, unmanagedPodsConfirmedAt NullableString, actorLabel string, createdAt string, updatedAt string, finishedAt NullableString, allowedActions []string, ) *NodeOperation`

NewNodeOperation instantiates a new NodeOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeOperationWithDefaults

`func NewNodeOperationWithDefaults() *NodeOperation`

NewNodeOperationWithDefaults instantiates a new NodeOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodeOperation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeOperation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeOperation) SetId(v int32)`

SetId sets Id field to given value.


### GetKind

`func (o *NodeOperation) GetKind() NodeOperationKindEnum`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NodeOperation) GetKindOk() (*NodeOperationKindEnum, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NodeOperation) SetKind(v NodeOperationKindEnum)`

SetKind sets Kind field to given value.


### GetSource

`func (o *NodeOperation) GetSource() NodeOperationSourceEnum`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NodeOperation) GetSourceOk() (*NodeOperationSourceEnum, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NodeOperation) SetSource(v NodeOperationSourceEnum)`

SetSource sets Source field to given value.


### GetTargetHostname

`func (o *NodeOperation) GetTargetHostname() string`

GetTargetHostname returns the TargetHostname field if non-nil, zero value otherwise.

### GetTargetHostnameOk

`func (o *NodeOperation) GetTargetHostnameOk() (*string, bool)`

GetTargetHostnameOk returns a tuple with the TargetHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetHostname

`func (o *NodeOperation) SetTargetHostname(v string)`

SetTargetHostname sets TargetHostname field to given value.


### GetStatus

`func (o *NodeOperation) GetStatus() NodeOperationStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeOperation) GetStatusOk() (*NodeOperationStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeOperation) SetStatus(v NodeOperationStatusEnum)`

SetStatus sets Status field to given value.


### GetReason

`func (o *NodeOperation) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *NodeOperation) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *NodeOperation) SetReason(v string)`

SetReason sets Reason field to given value.


### GetMessage

`func (o *NodeOperation) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NodeOperation) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NodeOperation) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetBypassPdb

`func (o *NodeOperation) GetBypassPdb() bool`

GetBypassPdb returns the BypassPdb field if non-nil, zero value otherwise.

### GetBypassPdbOk

`func (o *NodeOperation) GetBypassPdbOk() (*bool, bool)`

GetBypassPdbOk returns a tuple with the BypassPdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassPdb

`func (o *NodeOperation) SetBypassPdb(v bool)`

SetBypassPdb sets BypassPdb field to given value.


### GetDeleteUnmanagedPods

`func (o *NodeOperation) GetDeleteUnmanagedPods() bool`

GetDeleteUnmanagedPods returns the DeleteUnmanagedPods field if non-nil, zero value otherwise.

### GetDeleteUnmanagedPodsOk

`func (o *NodeOperation) GetDeleteUnmanagedPodsOk() (*bool, bool)`

GetDeleteUnmanagedPodsOk returns a tuple with the DeleteUnmanagedPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteUnmanagedPods

`func (o *NodeOperation) SetDeleteUnmanagedPods(v bool)`

SetDeleteUnmanagedPods sets DeleteUnmanagedPods field to given value.


### GetLocalDataLossAccepted

`func (o *NodeOperation) GetLocalDataLossAccepted() bool`

GetLocalDataLossAccepted returns the LocalDataLossAccepted field if non-nil, zero value otherwise.

### GetLocalDataLossAcceptedOk

`func (o *NodeOperation) GetLocalDataLossAcceptedOk() (*bool, bool)`

GetLocalDataLossAcceptedOk returns a tuple with the LocalDataLossAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDataLossAccepted

`func (o *NodeOperation) SetLocalDataLossAccepted(v bool)`

SetLocalDataLossAccepted sets LocalDataLossAccepted field to given value.


### GetBypassPdbConfirmedAt

`func (o *NodeOperation) GetBypassPdbConfirmedAt() string`

GetBypassPdbConfirmedAt returns the BypassPdbConfirmedAt field if non-nil, zero value otherwise.

### GetBypassPdbConfirmedAtOk

`func (o *NodeOperation) GetBypassPdbConfirmedAtOk() (*string, bool)`

GetBypassPdbConfirmedAtOk returns a tuple with the BypassPdbConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassPdbConfirmedAt

`func (o *NodeOperation) SetBypassPdbConfirmedAt(v string)`

SetBypassPdbConfirmedAt sets BypassPdbConfirmedAt field to given value.


### SetBypassPdbConfirmedAtNil

`func (o *NodeOperation) SetBypassPdbConfirmedAtNil(b bool)`

 SetBypassPdbConfirmedAtNil sets the value for BypassPdbConfirmedAt to be an explicit nil

### UnsetBypassPdbConfirmedAt
`func (o *NodeOperation) UnsetBypassPdbConfirmedAt()`

UnsetBypassPdbConfirmedAt ensures that no value is present for BypassPdbConfirmedAt, not even an explicit nil
### GetUnmanagedPodsConfirmedAt

`func (o *NodeOperation) GetUnmanagedPodsConfirmedAt() string`

GetUnmanagedPodsConfirmedAt returns the UnmanagedPodsConfirmedAt field if non-nil, zero value otherwise.

### GetUnmanagedPodsConfirmedAtOk

`func (o *NodeOperation) GetUnmanagedPodsConfirmedAtOk() (*string, bool)`

GetUnmanagedPodsConfirmedAtOk returns a tuple with the UnmanagedPodsConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmanagedPodsConfirmedAt

`func (o *NodeOperation) SetUnmanagedPodsConfirmedAt(v string)`

SetUnmanagedPodsConfirmedAt sets UnmanagedPodsConfirmedAt field to given value.


### SetUnmanagedPodsConfirmedAtNil

`func (o *NodeOperation) SetUnmanagedPodsConfirmedAtNil(b bool)`

 SetUnmanagedPodsConfirmedAtNil sets the value for UnmanagedPodsConfirmedAt to be an explicit nil

### UnsetUnmanagedPodsConfirmedAt
`func (o *NodeOperation) UnsetUnmanagedPodsConfirmedAt()`

UnsetUnmanagedPodsConfirmedAt ensures that no value is present for UnmanagedPodsConfirmedAt, not even an explicit nil
### GetActorLabel

`func (o *NodeOperation) GetActorLabel() string`

GetActorLabel returns the ActorLabel field if non-nil, zero value otherwise.

### GetActorLabelOk

`func (o *NodeOperation) GetActorLabelOk() (*string, bool)`

GetActorLabelOk returns a tuple with the ActorLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorLabel

`func (o *NodeOperation) SetActorLabel(v string)`

SetActorLabel sets ActorLabel field to given value.


### GetCreatedAt

`func (o *NodeOperation) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NodeOperation) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NodeOperation) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NodeOperation) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NodeOperation) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NodeOperation) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetFinishedAt

`func (o *NodeOperation) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *NodeOperation) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *NodeOperation) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.


### SetFinishedAtNil

`func (o *NodeOperation) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *NodeOperation) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetAllowedActions

`func (o *NodeOperation) GetAllowedActions() []string`

GetAllowedActions returns the AllowedActions field if non-nil, zero value otherwise.

### GetAllowedActionsOk

`func (o *NodeOperation) GetAllowedActionsOk() (*[]string, bool)`

GetAllowedActionsOk returns a tuple with the AllowedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedActions

`func (o *NodeOperation) SetAllowedActions(v []string)`

SetAllowedActions sets AllowedActions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


