# Volume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Project** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Size** | **int32** | GB | 
**Product** | **string** | ID or slug | 
**Attached** | **bool** |  | [readonly] 
**Server** | **string** |  | [readonly] [default to ""]

## Methods

### NewVolume

`func NewVolume(id int32, size int32, product string, attached bool, server string, ) *Volume`

NewVolume instantiates a new Volume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeWithDefaults

`func NewVolumeWithDefaults() *Volume`

NewVolumeWithDefaults instantiates a new Volume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Volume) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Volume) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Volume) SetId(v int32)`

SetId sets Id field to given value.


### GetProject

`func (o *Volume) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Volume) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Volume) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Volume) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetAlias

`func (o *Volume) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *Volume) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *Volume) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *Volume) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetSize

`func (o *Volume) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Volume) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Volume) SetSize(v int32)`

SetSize sets Size field to given value.


### GetProduct

`func (o *Volume) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Volume) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Volume) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetAttached

`func (o *Volume) GetAttached() bool`

GetAttached returns the Attached field if non-nil, zero value otherwise.

### GetAttachedOk

`func (o *Volume) GetAttachedOk() (*bool, bool)`

GetAttachedOk returns a tuple with the Attached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttached

`func (o *Volume) SetAttached(v bool)`

SetAttached sets Attached field to given value.


### GetServer

`func (o *Volume) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *Volume) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *Volume) SetServer(v string)`

SetServer sets Server field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


