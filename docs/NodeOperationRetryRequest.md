# NodeOperationRetryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BypassPdb** | Pointer to **NullableBool** |  | [optional] 
**DeleteUnmanagedPods** | Pointer to **NullableBool** |  | [optional] 
**AcknowledgePdbBypass** | Pointer to **bool** |  | [optional] [default to false]
**AcknowledgeUnmanagedPodDeletion** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewNodeOperationRetryRequest

`func NewNodeOperationRetryRequest() *NodeOperationRetryRequest`

NewNodeOperationRetryRequest instantiates a new NodeOperationRetryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeOperationRetryRequestWithDefaults

`func NewNodeOperationRetryRequestWithDefaults() *NodeOperationRetryRequest`

NewNodeOperationRetryRequestWithDefaults instantiates a new NodeOperationRetryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBypassPdb

`func (o *NodeOperationRetryRequest) GetBypassPdb() bool`

GetBypassPdb returns the BypassPdb field if non-nil, zero value otherwise.

### GetBypassPdbOk

`func (o *NodeOperationRetryRequest) GetBypassPdbOk() (*bool, bool)`

GetBypassPdbOk returns a tuple with the BypassPdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassPdb

`func (o *NodeOperationRetryRequest) SetBypassPdb(v bool)`

SetBypassPdb sets BypassPdb field to given value.

### HasBypassPdb

`func (o *NodeOperationRetryRequest) HasBypassPdb() bool`

HasBypassPdb returns a boolean if a field has been set.

### SetBypassPdbNil

`func (o *NodeOperationRetryRequest) SetBypassPdbNil(b bool)`

 SetBypassPdbNil sets the value for BypassPdb to be an explicit nil

### UnsetBypassPdb
`func (o *NodeOperationRetryRequest) UnsetBypassPdb()`

UnsetBypassPdb ensures that no value is present for BypassPdb, not even an explicit nil
### GetDeleteUnmanagedPods

`func (o *NodeOperationRetryRequest) GetDeleteUnmanagedPods() bool`

GetDeleteUnmanagedPods returns the DeleteUnmanagedPods field if non-nil, zero value otherwise.

### GetDeleteUnmanagedPodsOk

`func (o *NodeOperationRetryRequest) GetDeleteUnmanagedPodsOk() (*bool, bool)`

GetDeleteUnmanagedPodsOk returns a tuple with the DeleteUnmanagedPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteUnmanagedPods

`func (o *NodeOperationRetryRequest) SetDeleteUnmanagedPods(v bool)`

SetDeleteUnmanagedPods sets DeleteUnmanagedPods field to given value.

### HasDeleteUnmanagedPods

`func (o *NodeOperationRetryRequest) HasDeleteUnmanagedPods() bool`

HasDeleteUnmanagedPods returns a boolean if a field has been set.

### SetDeleteUnmanagedPodsNil

`func (o *NodeOperationRetryRequest) SetDeleteUnmanagedPodsNil(b bool)`

 SetDeleteUnmanagedPodsNil sets the value for DeleteUnmanagedPods to be an explicit nil

### UnsetDeleteUnmanagedPods
`func (o *NodeOperationRetryRequest) UnsetDeleteUnmanagedPods()`

UnsetDeleteUnmanagedPods ensures that no value is present for DeleteUnmanagedPods, not even an explicit nil
### GetAcknowledgePdbBypass

`func (o *NodeOperationRetryRequest) GetAcknowledgePdbBypass() bool`

GetAcknowledgePdbBypass returns the AcknowledgePdbBypass field if non-nil, zero value otherwise.

### GetAcknowledgePdbBypassOk

`func (o *NodeOperationRetryRequest) GetAcknowledgePdbBypassOk() (*bool, bool)`

GetAcknowledgePdbBypassOk returns a tuple with the AcknowledgePdbBypass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgePdbBypass

`func (o *NodeOperationRetryRequest) SetAcknowledgePdbBypass(v bool)`

SetAcknowledgePdbBypass sets AcknowledgePdbBypass field to given value.

### HasAcknowledgePdbBypass

`func (o *NodeOperationRetryRequest) HasAcknowledgePdbBypass() bool`

HasAcknowledgePdbBypass returns a boolean if a field has been set.

### GetAcknowledgeUnmanagedPodDeletion

`func (o *NodeOperationRetryRequest) GetAcknowledgeUnmanagedPodDeletion() bool`

GetAcknowledgeUnmanagedPodDeletion returns the AcknowledgeUnmanagedPodDeletion field if non-nil, zero value otherwise.

### GetAcknowledgeUnmanagedPodDeletionOk

`func (o *NodeOperationRetryRequest) GetAcknowledgeUnmanagedPodDeletionOk() (*bool, bool)`

GetAcknowledgeUnmanagedPodDeletionOk returns a tuple with the AcknowledgeUnmanagedPodDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgeUnmanagedPodDeletion

`func (o *NodeOperationRetryRequest) SetAcknowledgeUnmanagedPodDeletion(v bool)`

SetAcknowledgeUnmanagedPodDeletion sets AcknowledgeUnmanagedPodDeletion field to given value.

### HasAcknowledgeUnmanagedPodDeletion

`func (o *NodeOperationRetryRequest) HasAcknowledgeUnmanagedPodDeletion() bool`

HasAcknowledgeUnmanagedPodDeletion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


