# PoolRemovalItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetHostname** | **string** |  | [readonly] 
**CordonedAt** | **NullableString** |  | [readonly] 
**DrainedAt** | **NullableString** |  | [readonly] 
**DetachedAt** | **NullableString** |  | [readonly] 
**ValidatedAt** | **NullableString** |  | [readonly] 
**ResetStartedAt** | **NullableString** |  | [readonly] 
**ResetCompletedAt** | **NullableString** |  | [readonly] 
**NodeDeletedAt** | **NullableString** |  | [readonly] 
**VmDeletedAt** | **NullableString** |  | [readonly] 
**UncordonedAt** | **NullableString** |  | [readonly] 

## Methods

### NewPoolRemovalItem

`func NewPoolRemovalItem(targetHostname string, cordonedAt NullableString, drainedAt NullableString, detachedAt NullableString, validatedAt NullableString, resetStartedAt NullableString, resetCompletedAt NullableString, nodeDeletedAt NullableString, vmDeletedAt NullableString, uncordonedAt NullableString, ) *PoolRemovalItem`

NewPoolRemovalItem instantiates a new PoolRemovalItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolRemovalItemWithDefaults

`func NewPoolRemovalItemWithDefaults() *PoolRemovalItem`

NewPoolRemovalItemWithDefaults instantiates a new PoolRemovalItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetHostname

`func (o *PoolRemovalItem) GetTargetHostname() string`

GetTargetHostname returns the TargetHostname field if non-nil, zero value otherwise.

### GetTargetHostnameOk

`func (o *PoolRemovalItem) GetTargetHostnameOk() (*string, bool)`

GetTargetHostnameOk returns a tuple with the TargetHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetHostname

`func (o *PoolRemovalItem) SetTargetHostname(v string)`

SetTargetHostname sets TargetHostname field to given value.


### GetCordonedAt

`func (o *PoolRemovalItem) GetCordonedAt() string`

GetCordonedAt returns the CordonedAt field if non-nil, zero value otherwise.

### GetCordonedAtOk

`func (o *PoolRemovalItem) GetCordonedAtOk() (*string, bool)`

GetCordonedAtOk returns a tuple with the CordonedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCordonedAt

`func (o *PoolRemovalItem) SetCordonedAt(v string)`

SetCordonedAt sets CordonedAt field to given value.


### SetCordonedAtNil

`func (o *PoolRemovalItem) SetCordonedAtNil(b bool)`

 SetCordonedAtNil sets the value for CordonedAt to be an explicit nil

### UnsetCordonedAt
`func (o *PoolRemovalItem) UnsetCordonedAt()`

UnsetCordonedAt ensures that no value is present for CordonedAt, not even an explicit nil
### GetDrainedAt

`func (o *PoolRemovalItem) GetDrainedAt() string`

GetDrainedAt returns the DrainedAt field if non-nil, zero value otherwise.

### GetDrainedAtOk

`func (o *PoolRemovalItem) GetDrainedAtOk() (*string, bool)`

GetDrainedAtOk returns a tuple with the DrainedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainedAt

`func (o *PoolRemovalItem) SetDrainedAt(v string)`

SetDrainedAt sets DrainedAt field to given value.


### SetDrainedAtNil

`func (o *PoolRemovalItem) SetDrainedAtNil(b bool)`

 SetDrainedAtNil sets the value for DrainedAt to be an explicit nil

### UnsetDrainedAt
`func (o *PoolRemovalItem) UnsetDrainedAt()`

UnsetDrainedAt ensures that no value is present for DrainedAt, not even an explicit nil
### GetDetachedAt

`func (o *PoolRemovalItem) GetDetachedAt() string`

GetDetachedAt returns the DetachedAt field if non-nil, zero value otherwise.

### GetDetachedAtOk

`func (o *PoolRemovalItem) GetDetachedAtOk() (*string, bool)`

GetDetachedAtOk returns a tuple with the DetachedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetachedAt

`func (o *PoolRemovalItem) SetDetachedAt(v string)`

SetDetachedAt sets DetachedAt field to given value.


### SetDetachedAtNil

`func (o *PoolRemovalItem) SetDetachedAtNil(b bool)`

 SetDetachedAtNil sets the value for DetachedAt to be an explicit nil

### UnsetDetachedAt
`func (o *PoolRemovalItem) UnsetDetachedAt()`

UnsetDetachedAt ensures that no value is present for DetachedAt, not even an explicit nil
### GetValidatedAt

`func (o *PoolRemovalItem) GetValidatedAt() string`

GetValidatedAt returns the ValidatedAt field if non-nil, zero value otherwise.

### GetValidatedAtOk

`func (o *PoolRemovalItem) GetValidatedAtOk() (*string, bool)`

GetValidatedAtOk returns a tuple with the ValidatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedAt

`func (o *PoolRemovalItem) SetValidatedAt(v string)`

SetValidatedAt sets ValidatedAt field to given value.


### SetValidatedAtNil

`func (o *PoolRemovalItem) SetValidatedAtNil(b bool)`

 SetValidatedAtNil sets the value for ValidatedAt to be an explicit nil

### UnsetValidatedAt
`func (o *PoolRemovalItem) UnsetValidatedAt()`

UnsetValidatedAt ensures that no value is present for ValidatedAt, not even an explicit nil
### GetResetStartedAt

`func (o *PoolRemovalItem) GetResetStartedAt() string`

GetResetStartedAt returns the ResetStartedAt field if non-nil, zero value otherwise.

### GetResetStartedAtOk

`func (o *PoolRemovalItem) GetResetStartedAtOk() (*string, bool)`

GetResetStartedAtOk returns a tuple with the ResetStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetStartedAt

`func (o *PoolRemovalItem) SetResetStartedAt(v string)`

SetResetStartedAt sets ResetStartedAt field to given value.


### SetResetStartedAtNil

`func (o *PoolRemovalItem) SetResetStartedAtNil(b bool)`

 SetResetStartedAtNil sets the value for ResetStartedAt to be an explicit nil

### UnsetResetStartedAt
`func (o *PoolRemovalItem) UnsetResetStartedAt()`

UnsetResetStartedAt ensures that no value is present for ResetStartedAt, not even an explicit nil
### GetResetCompletedAt

`func (o *PoolRemovalItem) GetResetCompletedAt() string`

GetResetCompletedAt returns the ResetCompletedAt field if non-nil, zero value otherwise.

### GetResetCompletedAtOk

`func (o *PoolRemovalItem) GetResetCompletedAtOk() (*string, bool)`

GetResetCompletedAtOk returns a tuple with the ResetCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetCompletedAt

`func (o *PoolRemovalItem) SetResetCompletedAt(v string)`

SetResetCompletedAt sets ResetCompletedAt field to given value.


### SetResetCompletedAtNil

`func (o *PoolRemovalItem) SetResetCompletedAtNil(b bool)`

 SetResetCompletedAtNil sets the value for ResetCompletedAt to be an explicit nil

### UnsetResetCompletedAt
`func (o *PoolRemovalItem) UnsetResetCompletedAt()`

UnsetResetCompletedAt ensures that no value is present for ResetCompletedAt, not even an explicit nil
### GetNodeDeletedAt

`func (o *PoolRemovalItem) GetNodeDeletedAt() string`

GetNodeDeletedAt returns the NodeDeletedAt field if non-nil, zero value otherwise.

### GetNodeDeletedAtOk

`func (o *PoolRemovalItem) GetNodeDeletedAtOk() (*string, bool)`

GetNodeDeletedAtOk returns a tuple with the NodeDeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeDeletedAt

`func (o *PoolRemovalItem) SetNodeDeletedAt(v string)`

SetNodeDeletedAt sets NodeDeletedAt field to given value.


### SetNodeDeletedAtNil

`func (o *PoolRemovalItem) SetNodeDeletedAtNil(b bool)`

 SetNodeDeletedAtNil sets the value for NodeDeletedAt to be an explicit nil

### UnsetNodeDeletedAt
`func (o *PoolRemovalItem) UnsetNodeDeletedAt()`

UnsetNodeDeletedAt ensures that no value is present for NodeDeletedAt, not even an explicit nil
### GetVmDeletedAt

`func (o *PoolRemovalItem) GetVmDeletedAt() string`

GetVmDeletedAt returns the VmDeletedAt field if non-nil, zero value otherwise.

### GetVmDeletedAtOk

`func (o *PoolRemovalItem) GetVmDeletedAtOk() (*string, bool)`

GetVmDeletedAtOk returns a tuple with the VmDeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmDeletedAt

`func (o *PoolRemovalItem) SetVmDeletedAt(v string)`

SetVmDeletedAt sets VmDeletedAt field to given value.


### SetVmDeletedAtNil

`func (o *PoolRemovalItem) SetVmDeletedAtNil(b bool)`

 SetVmDeletedAtNil sets the value for VmDeletedAt to be an explicit nil

### UnsetVmDeletedAt
`func (o *PoolRemovalItem) UnsetVmDeletedAt()`

UnsetVmDeletedAt ensures that no value is present for VmDeletedAt, not even an explicit nil
### GetUncordonedAt

`func (o *PoolRemovalItem) GetUncordonedAt() string`

GetUncordonedAt returns the UncordonedAt field if non-nil, zero value otherwise.

### GetUncordonedAtOk

`func (o *PoolRemovalItem) GetUncordonedAtOk() (*string, bool)`

GetUncordonedAtOk returns a tuple with the UncordonedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncordonedAt

`func (o *PoolRemovalItem) SetUncordonedAt(v string)`

SetUncordonedAt sets UncordonedAt field to given value.


### SetUncordonedAtNil

`func (o *PoolRemovalItem) SetUncordonedAtNil(b bool)`

 SetUncordonedAtNil sets the value for UncordonedAt to be an explicit nil

### UnsetUncordonedAt
`func (o *PoolRemovalItem) UnsetUncordonedAt()`

UnsetUncordonedAt ensures that no value is present for UncordonedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


