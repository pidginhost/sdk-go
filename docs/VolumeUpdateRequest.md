# VolumeUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Size** | **int32** | GB | 

## Methods

### NewVolumeUpdateRequest

`func NewVolumeUpdateRequest(size int32, ) *VolumeUpdateRequest`

NewVolumeUpdateRequest instantiates a new VolumeUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeUpdateRequestWithDefaults

`func NewVolumeUpdateRequestWithDefaults() *VolumeUpdateRequest`

NewVolumeUpdateRequestWithDefaults instantiates a new VolumeUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *VolumeUpdateRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *VolumeUpdateRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *VolumeUpdateRequest) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *VolumeUpdateRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetAlias

`func (o *VolumeUpdateRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *VolumeUpdateRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *VolumeUpdateRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *VolumeUpdateRequest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetSize

`func (o *VolumeUpdateRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VolumeUpdateRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VolumeUpdateRequest) SetSize(v int32)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


