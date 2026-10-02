# ClusterEncryption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** |  | [readonly] 
**Status** | [**ClusterEncryptionStatusEnum**](ClusterEncryptionStatusEnum.md) |  | [readonly] 
**ChangedAt** | **NullableString** |  | [readonly] 
**VerifiedAt** | **NullableString** |  | [readonly] 
**RestartRequired** | **bool** |  | [readonly] 
**RestartRequiredAt** | **NullableString** |  | [readonly] 
**RestartCheckedAt** | **NullableString** |  | [readonly] 
**StalePodCount** | **NullableInt32** |  | [readonly] 
**Reason** | **string** |  | [readonly] 
**Error** | **string** |  | [readonly] 
**PerNode** | **map[string]interface{}** | Evidence from the NEWEST operation, which may not have any yet.  A freshly queued operation carries an empty &#x60;&#x60;verification_result&#x60;&#x60;, so this blanks the moment a toggle is admitted while &#x60;&#x60;mode&#x60;&#x60; and &#x60;&#x60;status&#x60;&#x60; still describe the last verified state. An empty map therefore means \&quot;no evidence from the current operation\&quot;, NEVER \&quot;verification failed\&quot; -- read &#x60;&#x60;operation.status&#x60;&#x60; to tell them apart. Showing the previous operation&#39;s rows instead would label evidence for one mode as evidence for another. | [readonly] 
**Operation** | [**NullableClusterEncryptionOperation**](ClusterEncryptionOperation.md) |  | [readonly] 

## Methods

### NewClusterEncryption

`func NewClusterEncryption(mode string, status ClusterEncryptionStatusEnum, changedAt NullableString, verifiedAt NullableString, restartRequired bool, restartRequiredAt NullableString, restartCheckedAt NullableString, stalePodCount NullableInt32, reason string, error_ string, perNode map[string]interface{}, operation NullableClusterEncryptionOperation, ) *ClusterEncryption`

NewClusterEncryption instantiates a new ClusterEncryption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterEncryptionWithDefaults

`func NewClusterEncryptionWithDefaults() *ClusterEncryption`

NewClusterEncryptionWithDefaults instantiates a new ClusterEncryption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *ClusterEncryption) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ClusterEncryption) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ClusterEncryption) SetMode(v string)`

SetMode sets Mode field to given value.


### GetStatus

`func (o *ClusterEncryption) GetStatus() ClusterEncryptionStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterEncryption) GetStatusOk() (*ClusterEncryptionStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterEncryption) SetStatus(v ClusterEncryptionStatusEnum)`

SetStatus sets Status field to given value.


### GetChangedAt

`func (o *ClusterEncryption) GetChangedAt() string`

GetChangedAt returns the ChangedAt field if non-nil, zero value otherwise.

### GetChangedAtOk

`func (o *ClusterEncryption) GetChangedAtOk() (*string, bool)`

GetChangedAtOk returns a tuple with the ChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedAt

`func (o *ClusterEncryption) SetChangedAt(v string)`

SetChangedAt sets ChangedAt field to given value.


### SetChangedAtNil

`func (o *ClusterEncryption) SetChangedAtNil(b bool)`

 SetChangedAtNil sets the value for ChangedAt to be an explicit nil

### UnsetChangedAt
`func (o *ClusterEncryption) UnsetChangedAt()`

UnsetChangedAt ensures that no value is present for ChangedAt, not even an explicit nil
### GetVerifiedAt

`func (o *ClusterEncryption) GetVerifiedAt() string`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *ClusterEncryption) GetVerifiedAtOk() (*string, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *ClusterEncryption) SetVerifiedAt(v string)`

SetVerifiedAt sets VerifiedAt field to given value.


### SetVerifiedAtNil

`func (o *ClusterEncryption) SetVerifiedAtNil(b bool)`

 SetVerifiedAtNil sets the value for VerifiedAt to be an explicit nil

### UnsetVerifiedAt
`func (o *ClusterEncryption) UnsetVerifiedAt()`

UnsetVerifiedAt ensures that no value is present for VerifiedAt, not even an explicit nil
### GetRestartRequired

`func (o *ClusterEncryption) GetRestartRequired() bool`

GetRestartRequired returns the RestartRequired field if non-nil, zero value otherwise.

### GetRestartRequiredOk

`func (o *ClusterEncryption) GetRestartRequiredOk() (*bool, bool)`

GetRestartRequiredOk returns a tuple with the RestartRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartRequired

`func (o *ClusterEncryption) SetRestartRequired(v bool)`

SetRestartRequired sets RestartRequired field to given value.


### GetRestartRequiredAt

`func (o *ClusterEncryption) GetRestartRequiredAt() string`

GetRestartRequiredAt returns the RestartRequiredAt field if non-nil, zero value otherwise.

### GetRestartRequiredAtOk

`func (o *ClusterEncryption) GetRestartRequiredAtOk() (*string, bool)`

GetRestartRequiredAtOk returns a tuple with the RestartRequiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartRequiredAt

`func (o *ClusterEncryption) SetRestartRequiredAt(v string)`

SetRestartRequiredAt sets RestartRequiredAt field to given value.


### SetRestartRequiredAtNil

`func (o *ClusterEncryption) SetRestartRequiredAtNil(b bool)`

 SetRestartRequiredAtNil sets the value for RestartRequiredAt to be an explicit nil

### UnsetRestartRequiredAt
`func (o *ClusterEncryption) UnsetRestartRequiredAt()`

UnsetRestartRequiredAt ensures that no value is present for RestartRequiredAt, not even an explicit nil
### GetRestartCheckedAt

`func (o *ClusterEncryption) GetRestartCheckedAt() string`

GetRestartCheckedAt returns the RestartCheckedAt field if non-nil, zero value otherwise.

### GetRestartCheckedAtOk

`func (o *ClusterEncryption) GetRestartCheckedAtOk() (*string, bool)`

GetRestartCheckedAtOk returns a tuple with the RestartCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCheckedAt

`func (o *ClusterEncryption) SetRestartCheckedAt(v string)`

SetRestartCheckedAt sets RestartCheckedAt field to given value.


### SetRestartCheckedAtNil

`func (o *ClusterEncryption) SetRestartCheckedAtNil(b bool)`

 SetRestartCheckedAtNil sets the value for RestartCheckedAt to be an explicit nil

### UnsetRestartCheckedAt
`func (o *ClusterEncryption) UnsetRestartCheckedAt()`

UnsetRestartCheckedAt ensures that no value is present for RestartCheckedAt, not even an explicit nil
### GetStalePodCount

`func (o *ClusterEncryption) GetStalePodCount() int32`

GetStalePodCount returns the StalePodCount field if non-nil, zero value otherwise.

### GetStalePodCountOk

`func (o *ClusterEncryption) GetStalePodCountOk() (*int32, bool)`

GetStalePodCountOk returns a tuple with the StalePodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStalePodCount

`func (o *ClusterEncryption) SetStalePodCount(v int32)`

SetStalePodCount sets StalePodCount field to given value.


### SetStalePodCountNil

`func (o *ClusterEncryption) SetStalePodCountNil(b bool)`

 SetStalePodCountNil sets the value for StalePodCount to be an explicit nil

### UnsetStalePodCount
`func (o *ClusterEncryption) UnsetStalePodCount()`

UnsetStalePodCount ensures that no value is present for StalePodCount, not even an explicit nil
### GetReason

`func (o *ClusterEncryption) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ClusterEncryption) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ClusterEncryption) SetReason(v string)`

SetReason sets Reason field to given value.


### GetError

`func (o *ClusterEncryption) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ClusterEncryption) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ClusterEncryption) SetError(v string)`

SetError sets Error field to given value.


### GetPerNode

`func (o *ClusterEncryption) GetPerNode() map[string]interface{}`

GetPerNode returns the PerNode field if non-nil, zero value otherwise.

### GetPerNodeOk

`func (o *ClusterEncryption) GetPerNodeOk() (*map[string]interface{}, bool)`

GetPerNodeOk returns a tuple with the PerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerNode

`func (o *ClusterEncryption) SetPerNode(v map[string]interface{})`

SetPerNode sets PerNode field to given value.


### GetOperation

`func (o *ClusterEncryption) GetOperation() ClusterEncryptionOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ClusterEncryption) GetOperationOk() (*ClusterEncryptionOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ClusterEncryption) SetOperation(v ClusterEncryptionOperation)`

SetOperation sets Operation field to given value.


### SetOperationNil

`func (o *ClusterEncryption) SetOperationNil(b bool)`

 SetOperationNil sets the value for Operation to be an explicit nil

### UnsetOperation
`func (o *ClusterEncryption) UnsetOperation()`

UnsetOperation ensures that no value is present for Operation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


