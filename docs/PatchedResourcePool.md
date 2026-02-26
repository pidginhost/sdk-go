# PatchedResourcePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Package** | Pointer to **string** |  | [optional] [readonly] 
**Size** | Pointer to **string** |  | [optional] [readonly] 
**Nodes** | Pointer to [**[]ResourcePoolNode**](ResourcePoolNode.md) |  | [optional] [readonly] 
**NewSize** | Pointer to **int32** |  | [optional] 

## Methods

### NewPatchedResourcePool

`func NewPatchedResourcePool() *PatchedResourcePool`

NewPatchedResourcePool instantiates a new PatchedResourcePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedResourcePoolWithDefaults

`func NewPatchedResourcePoolWithDefaults() *PatchedResourcePool`

NewPatchedResourcePoolWithDefaults instantiates a new PatchedResourcePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedResourcePool) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedResourcePool) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedResourcePool) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedResourcePool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPackage

`func (o *PatchedResourcePool) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *PatchedResourcePool) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *PatchedResourcePool) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *PatchedResourcePool) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetSize

`func (o *PatchedResourcePool) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PatchedResourcePool) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PatchedResourcePool) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *PatchedResourcePool) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetNodes

`func (o *PatchedResourcePool) GetNodes() []ResourcePoolNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *PatchedResourcePool) GetNodesOk() (*[]ResourcePoolNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *PatchedResourcePool) SetNodes(v []ResourcePoolNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *PatchedResourcePool) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetNewSize

`func (o *PatchedResourcePool) GetNewSize() int32`

GetNewSize returns the NewSize field if non-nil, zero value otherwise.

### GetNewSizeOk

`func (o *PatchedResourcePool) GetNewSizeOk() (*int32, bool)`

GetNewSizeOk returns a tuple with the NewSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSize

`func (o *PatchedResourcePool) SetNewSize(v int32)`

SetNewSize sets NewSize field to given value.

### HasNewSize

`func (o *PatchedResourcePool) HasNewSize() bool`

HasNewSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


