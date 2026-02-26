# ClusterType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**TypeEnum**](TypeEnum.md) |  | 
**WorkerNodesCountMin** | Pointer to **int32** |  | [optional] 
**WorkerNodesCountMax** | Pointer to **int32** |  | [optional] 
**WorkerNodePackages** | [**[]ClusterPackage**](ClusterPackage.md) |  | 

## Methods

### NewClusterType

`func NewClusterType(type_ TypeEnum, workerNodePackages []ClusterPackage, ) *ClusterType`

NewClusterType instantiates a new ClusterType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterTypeWithDefaults

`func NewClusterTypeWithDefaults() *ClusterType`

NewClusterTypeWithDefaults instantiates a new ClusterType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ClusterType) GetType() TypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterType) GetTypeOk() (*TypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterType) SetType(v TypeEnum)`

SetType sets Type field to given value.


### GetWorkerNodesCountMin

`func (o *ClusterType) GetWorkerNodesCountMin() int32`

GetWorkerNodesCountMin returns the WorkerNodesCountMin field if non-nil, zero value otherwise.

### GetWorkerNodesCountMinOk

`func (o *ClusterType) GetWorkerNodesCountMinOk() (*int32, bool)`

GetWorkerNodesCountMinOk returns a tuple with the WorkerNodesCountMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerNodesCountMin

`func (o *ClusterType) SetWorkerNodesCountMin(v int32)`

SetWorkerNodesCountMin sets WorkerNodesCountMin field to given value.

### HasWorkerNodesCountMin

`func (o *ClusterType) HasWorkerNodesCountMin() bool`

HasWorkerNodesCountMin returns a boolean if a field has been set.

### GetWorkerNodesCountMax

`func (o *ClusterType) GetWorkerNodesCountMax() int32`

GetWorkerNodesCountMax returns the WorkerNodesCountMax field if non-nil, zero value otherwise.

### GetWorkerNodesCountMaxOk

`func (o *ClusterType) GetWorkerNodesCountMaxOk() (*int32, bool)`

GetWorkerNodesCountMaxOk returns a tuple with the WorkerNodesCountMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerNodesCountMax

`func (o *ClusterType) SetWorkerNodesCountMax(v int32)`

SetWorkerNodesCountMax sets WorkerNodesCountMax field to given value.

### HasWorkerNodesCountMax

`func (o *ClusterType) HasWorkerNodesCountMax() bool`

HasWorkerNodesCountMax returns a boolean if a field has been set.

### GetWorkerNodePackages

`func (o *ClusterType) GetWorkerNodePackages() []ClusterPackage`

GetWorkerNodePackages returns the WorkerNodePackages field if non-nil, zero value otherwise.

### GetWorkerNodePackagesOk

`func (o *ClusterType) GetWorkerNodePackagesOk() (*[]ClusterPackage, bool)`

GetWorkerNodePackagesOk returns a tuple with the WorkerNodePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerNodePackages

`func (o *ClusterType) SetWorkerNodePackages(v []ClusterPackage)`

SetWorkerNodePackages sets WorkerNodePackages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


