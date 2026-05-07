# PatchedVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Project** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** | GB | [optional] 
**Product** | Pointer to **string** | ID or slug | [optional] 
**Attached** | Pointer to **bool** |  | [optional] [readonly] 
**Server** | Pointer to **string** |  | [optional] [readonly] [default to ""]

## Methods

### NewPatchedVolume

`func NewPatchedVolume() *PatchedVolume`

NewPatchedVolume instantiates a new PatchedVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedVolumeWithDefaults

`func NewPatchedVolumeWithDefaults() *PatchedVolume`

NewPatchedVolumeWithDefaults instantiates a new PatchedVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedVolume) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedVolume) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedVolume) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedVolume) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProject

`func (o *PatchedVolume) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PatchedVolume) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PatchedVolume) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PatchedVolume) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetAlias

`func (o *PatchedVolume) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *PatchedVolume) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *PatchedVolume) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *PatchedVolume) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetSize

`func (o *PatchedVolume) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PatchedVolume) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PatchedVolume) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PatchedVolume) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetProduct

`func (o *PatchedVolume) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PatchedVolume) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *PatchedVolume) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *PatchedVolume) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetAttached

`func (o *PatchedVolume) GetAttached() bool`

GetAttached returns the Attached field if non-nil, zero value otherwise.

### GetAttachedOk

`func (o *PatchedVolume) GetAttachedOk() (*bool, bool)`

GetAttachedOk returns a tuple with the Attached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttached

`func (o *PatchedVolume) SetAttached(v bool)`

SetAttached sets Attached field to given value.

### HasAttached

`func (o *PatchedVolume) HasAttached() bool`

HasAttached returns a boolean if a field has been set.

### GetServer

`func (o *PatchedVolume) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *PatchedVolume) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *PatchedVolume) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *PatchedVolume) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


