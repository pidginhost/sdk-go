# ResourcePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Package** | **string** |  | [readonly] 
**Size** | **string** |  | [readonly] 
**Nodes** | [**[]ResourcePoolNode**](ResourcePoolNode.md) |  | [readonly] 
**NewSize** | Pointer to **int32** |  | [optional] 

## Methods

### NewResourcePool

`func NewResourcePool(id int32, package_ string, size string, nodes []ResourcePoolNode, ) *ResourcePool`

NewResourcePool instantiates a new ResourcePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePoolWithDefaults

`func NewResourcePoolWithDefaults() *ResourcePool`

NewResourcePoolWithDefaults instantiates a new ResourcePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourcePool) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourcePool) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourcePool) SetId(v int32)`

SetId sets Id field to given value.


### GetPackage

`func (o *ResourcePool) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *ResourcePool) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *ResourcePool) SetPackage(v string)`

SetPackage sets Package field to given value.


### GetSize

`func (o *ResourcePool) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ResourcePool) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ResourcePool) SetSize(v string)`

SetSize sets Size field to given value.


### GetNodes

`func (o *ResourcePool) GetNodes() []ResourcePoolNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ResourcePool) GetNodesOk() (*[]ResourcePoolNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ResourcePool) SetNodes(v []ResourcePoolNode)`

SetNodes sets Nodes field to given value.


### GetNewSize

`func (o *ResourcePool) GetNewSize() int32`

GetNewSize returns the NewSize field if non-nil, zero value otherwise.

### GetNewSizeOk

`func (o *ResourcePool) GetNewSizeOk() (*int32, bool)`

GetNewSizeOk returns a tuple with the NewSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSize

`func (o *ResourcePool) SetNewSize(v int32)`

SetNewSize sets NewSize field to given value.

### HasNewSize

`func (o *ResourcePool) HasNewSize() bool`

HasNewSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


