# ResourcePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Package** | **string** |  | [readonly] 
**Generation** | **string** |  | [readonly] 
**Size** | **int32** |  | [readonly] 
**Nodes** | [**[]ResourcePoolNode**](ResourcePoolNode.md) |  | [readonly] 

## Methods

### NewResourcePool

`func NewResourcePool(id int32, package_ string, generation string, size int32, nodes []ResourcePoolNode, ) *ResourcePool`

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


### GetGeneration

`func (o *ResourcePool) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *ResourcePool) GetGenerationOk() (*string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *ResourcePool) SetGeneration(v string)`

SetGeneration sets Generation field to given value.


### GetSize

`func (o *ResourcePool) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ResourcePool) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ResourcePool) SetSize(v int32)`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


