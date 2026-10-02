# PatchedVolumeUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** | GB | [optional] 

## Methods

### NewPatchedVolumeUpdateRequest

`func NewPatchedVolumeUpdateRequest() *PatchedVolumeUpdateRequest`

NewPatchedVolumeUpdateRequest instantiates a new PatchedVolumeUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedVolumeUpdateRequestWithDefaults

`func NewPatchedVolumeUpdateRequestWithDefaults() *PatchedVolumeUpdateRequest`

NewPatchedVolumeUpdateRequestWithDefaults instantiates a new PatchedVolumeUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *PatchedVolumeUpdateRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PatchedVolumeUpdateRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PatchedVolumeUpdateRequest) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PatchedVolumeUpdateRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetAlias

`func (o *PatchedVolumeUpdateRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *PatchedVolumeUpdateRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *PatchedVolumeUpdateRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *PatchedVolumeUpdateRequest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetSize

`func (o *PatchedVolumeUpdateRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PatchedVolumeUpdateRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PatchedVolumeUpdateRequest) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PatchedVolumeUpdateRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


