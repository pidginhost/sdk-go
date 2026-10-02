# PoolRemovalJournal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Kind** | [**PoolRemovalJournalKindEnum**](PoolRemovalJournalKindEnum.md) |  | [readonly] 
**Status** | [**PoolRemovalJournalStatusEnum**](PoolRemovalJournalStatusEnum.md) |  | [readonly] 
**Reason** | **string** |  | [readonly] 
**Message** | **string** |  | [readonly] 
**RequestedPoolSize** | **NullableInt32** |  | [readonly] 
**LocalDataLossAccepted** | **bool** |  | [readonly] 
**ActorLabel** | **string** | Who requested the removal (user email or staff name). Never token material. | [readonly] 
**CreatedAt** | **string** |  | [readonly] 
**UpdatedAt** | **string** |  | [readonly] 
**FinishedAt** | **NullableString** |  | [readonly] 
**Items** | [**[]PoolRemovalItem**](PoolRemovalItem.md) |  | [readonly] 
**AllowedActions** | **[]string** |  | [readonly] 

## Methods

### NewPoolRemovalJournal

`func NewPoolRemovalJournal(id int32, kind PoolRemovalJournalKindEnum, status PoolRemovalJournalStatusEnum, reason string, message string, requestedPoolSize NullableInt32, localDataLossAccepted bool, actorLabel string, createdAt string, updatedAt string, finishedAt NullableString, items []PoolRemovalItem, allowedActions []string, ) *PoolRemovalJournal`

NewPoolRemovalJournal instantiates a new PoolRemovalJournal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolRemovalJournalWithDefaults

`func NewPoolRemovalJournalWithDefaults() *PoolRemovalJournal`

NewPoolRemovalJournalWithDefaults instantiates a new PoolRemovalJournal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PoolRemovalJournal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoolRemovalJournal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoolRemovalJournal) SetId(v int32)`

SetId sets Id field to given value.


### GetKind

`func (o *PoolRemovalJournal) GetKind() PoolRemovalJournalKindEnum`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PoolRemovalJournal) GetKindOk() (*PoolRemovalJournalKindEnum, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PoolRemovalJournal) SetKind(v PoolRemovalJournalKindEnum)`

SetKind sets Kind field to given value.


### GetStatus

`func (o *PoolRemovalJournal) GetStatus() PoolRemovalJournalStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PoolRemovalJournal) GetStatusOk() (*PoolRemovalJournalStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PoolRemovalJournal) SetStatus(v PoolRemovalJournalStatusEnum)`

SetStatus sets Status field to given value.


### GetReason

`func (o *PoolRemovalJournal) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PoolRemovalJournal) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PoolRemovalJournal) SetReason(v string)`

SetReason sets Reason field to given value.


### GetMessage

`func (o *PoolRemovalJournal) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PoolRemovalJournal) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PoolRemovalJournal) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRequestedPoolSize

`func (o *PoolRemovalJournal) GetRequestedPoolSize() int32`

GetRequestedPoolSize returns the RequestedPoolSize field if non-nil, zero value otherwise.

### GetRequestedPoolSizeOk

`func (o *PoolRemovalJournal) GetRequestedPoolSizeOk() (*int32, bool)`

GetRequestedPoolSizeOk returns a tuple with the RequestedPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedPoolSize

`func (o *PoolRemovalJournal) SetRequestedPoolSize(v int32)`

SetRequestedPoolSize sets RequestedPoolSize field to given value.


### SetRequestedPoolSizeNil

`func (o *PoolRemovalJournal) SetRequestedPoolSizeNil(b bool)`

 SetRequestedPoolSizeNil sets the value for RequestedPoolSize to be an explicit nil

### UnsetRequestedPoolSize
`func (o *PoolRemovalJournal) UnsetRequestedPoolSize()`

UnsetRequestedPoolSize ensures that no value is present for RequestedPoolSize, not even an explicit nil
### GetLocalDataLossAccepted

`func (o *PoolRemovalJournal) GetLocalDataLossAccepted() bool`

GetLocalDataLossAccepted returns the LocalDataLossAccepted field if non-nil, zero value otherwise.

### GetLocalDataLossAcceptedOk

`func (o *PoolRemovalJournal) GetLocalDataLossAcceptedOk() (*bool, bool)`

GetLocalDataLossAcceptedOk returns a tuple with the LocalDataLossAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDataLossAccepted

`func (o *PoolRemovalJournal) SetLocalDataLossAccepted(v bool)`

SetLocalDataLossAccepted sets LocalDataLossAccepted field to given value.


### GetActorLabel

`func (o *PoolRemovalJournal) GetActorLabel() string`

GetActorLabel returns the ActorLabel field if non-nil, zero value otherwise.

### GetActorLabelOk

`func (o *PoolRemovalJournal) GetActorLabelOk() (*string, bool)`

GetActorLabelOk returns a tuple with the ActorLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorLabel

`func (o *PoolRemovalJournal) SetActorLabel(v string)`

SetActorLabel sets ActorLabel field to given value.


### GetCreatedAt

`func (o *PoolRemovalJournal) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PoolRemovalJournal) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PoolRemovalJournal) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PoolRemovalJournal) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PoolRemovalJournal) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PoolRemovalJournal) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetFinishedAt

`func (o *PoolRemovalJournal) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *PoolRemovalJournal) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *PoolRemovalJournal) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.


### SetFinishedAtNil

`func (o *PoolRemovalJournal) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *PoolRemovalJournal) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetItems

`func (o *PoolRemovalJournal) GetItems() []PoolRemovalItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PoolRemovalJournal) GetItemsOk() (*[]PoolRemovalItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PoolRemovalJournal) SetItems(v []PoolRemovalItem)`

SetItems sets Items field to given value.


### GetAllowedActions

`func (o *PoolRemovalJournal) GetAllowedActions() []string`

GetAllowedActions returns the AllowedActions field if non-nil, zero value otherwise.

### GetAllowedActionsOk

`func (o *PoolRemovalJournal) GetAllowedActionsOk() (*[]string, bool)`

GetAllowedActionsOk returns a tuple with the AllowedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedActions

`func (o *PoolRemovalJournal) SetAllowedActions(v []string)`

SetAllowedActions sets AllowedActions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


