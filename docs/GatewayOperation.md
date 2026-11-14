# GatewayOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Kind** | [**GatewayOperationKindEnum**](GatewayOperationKindEnum.md) |  | [readonly] 
**Status** | [**GatewayOperationStatusEnum**](GatewayOperationStatusEnum.md) |  | [readonly] 
**Network** | **NullableInt32** |  | [readonly] 
**Revision** | **NullableInt32** |  | [readonly] 
**StartedAt** | **string** |  | [readonly] 
**DeadlineAt** | **NullableString** |  | [readonly] 
**CompletedAt** | **NullableString** |  | [readonly] 
**Servers** | **[]int32** |  | [readonly] 
**Outcomes** | [**map[string]GatewayRouteOutcome**](GatewayRouteOutcome.md) |  | [readonly] 
**Reason** | **string** | Why the operation failed; empty unless it did. | [readonly] 

## Methods

### NewGatewayOperation

`func NewGatewayOperation(id string, kind GatewayOperationKindEnum, status GatewayOperationStatusEnum, network NullableInt32, revision NullableInt32, startedAt string, deadlineAt NullableString, completedAt NullableString, servers []int32, outcomes map[string]GatewayRouteOutcome, reason string, ) *GatewayOperation`

NewGatewayOperation instantiates a new GatewayOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayOperationWithDefaults

`func NewGatewayOperationWithDefaults() *GatewayOperation`

NewGatewayOperationWithDefaults instantiates a new GatewayOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GatewayOperation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayOperation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayOperation) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *GatewayOperation) GetKind() GatewayOperationKindEnum`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GatewayOperation) GetKindOk() (*GatewayOperationKindEnum, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GatewayOperation) SetKind(v GatewayOperationKindEnum)`

SetKind sets Kind field to given value.


### GetStatus

`func (o *GatewayOperation) GetStatus() GatewayOperationStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GatewayOperation) GetStatusOk() (*GatewayOperationStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GatewayOperation) SetStatus(v GatewayOperationStatusEnum)`

SetStatus sets Status field to given value.


### GetNetwork

`func (o *GatewayOperation) GetNetwork() int32`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GatewayOperation) GetNetworkOk() (*int32, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GatewayOperation) SetNetwork(v int32)`

SetNetwork sets Network field to given value.


### SetNetworkNil

`func (o *GatewayOperation) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *GatewayOperation) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil
### GetRevision

`func (o *GatewayOperation) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *GatewayOperation) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *GatewayOperation) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### SetRevisionNil

`func (o *GatewayOperation) SetRevisionNil(b bool)`

 SetRevisionNil sets the value for Revision to be an explicit nil

### UnsetRevision
`func (o *GatewayOperation) UnsetRevision()`

UnsetRevision ensures that no value is present for Revision, not even an explicit nil
### GetStartedAt

`func (o *GatewayOperation) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GatewayOperation) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GatewayOperation) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.


### GetDeadlineAt

`func (o *GatewayOperation) GetDeadlineAt() string`

GetDeadlineAt returns the DeadlineAt field if non-nil, zero value otherwise.

### GetDeadlineAtOk

`func (o *GatewayOperation) GetDeadlineAtOk() (*string, bool)`

GetDeadlineAtOk returns a tuple with the DeadlineAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadlineAt

`func (o *GatewayOperation) SetDeadlineAt(v string)`

SetDeadlineAt sets DeadlineAt field to given value.


### SetDeadlineAtNil

`func (o *GatewayOperation) SetDeadlineAtNil(b bool)`

 SetDeadlineAtNil sets the value for DeadlineAt to be an explicit nil

### UnsetDeadlineAt
`func (o *GatewayOperation) UnsetDeadlineAt()`

UnsetDeadlineAt ensures that no value is present for DeadlineAt, not even an explicit nil
### GetCompletedAt

`func (o *GatewayOperation) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GatewayOperation) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GatewayOperation) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *GatewayOperation) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *GatewayOperation) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetServers

`func (o *GatewayOperation) GetServers() []int32`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GatewayOperation) GetServersOk() (*[]int32, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GatewayOperation) SetServers(v []int32)`

SetServers sets Servers field to given value.


### GetOutcomes

`func (o *GatewayOperation) GetOutcomes() map[string]GatewayRouteOutcome`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *GatewayOperation) GetOutcomesOk() (*map[string]GatewayRouteOutcome, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomes

`func (o *GatewayOperation) SetOutcomes(v map[string]GatewayRouteOutcome)`

SetOutcomes sets Outcomes field to given value.


### GetReason

`func (o *GatewayOperation) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GatewayOperation) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GatewayOperation) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


