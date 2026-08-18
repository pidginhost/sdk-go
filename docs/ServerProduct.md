# ServerProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Slug** | **string** |  | 
**Name** | **string** |  | [readonly] 
**Cpus** | **int32** |  | [readonly] 
**Memory** | **int32** |  | [readonly] 
**DiskSize** | **int32** |  | [readonly] 
**Traffic** | **int32** | Included monthly traffic, in TB. | [readonly] 
**AvailableGenerations** | **[]string** |  | [readonly] 

## Methods

### NewServerProduct

`func NewServerProduct(id int32, slug string, name string, cpus int32, memory int32, diskSize int32, traffic int32, availableGenerations []string, ) *ServerProduct`

NewServerProduct instantiates a new ServerProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerProductWithDefaults

`func NewServerProductWithDefaults() *ServerProduct`

NewServerProductWithDefaults instantiates a new ServerProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerProduct) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerProduct) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerProduct) SetId(v int32)`

SetId sets Id field to given value.


### GetSlug

`func (o *ServerProduct) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ServerProduct) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ServerProduct) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetName

`func (o *ServerProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerProduct) SetName(v string)`

SetName sets Name field to given value.


### GetCpus

`func (o *ServerProduct) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *ServerProduct) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *ServerProduct) SetCpus(v int32)`

SetCpus sets Cpus field to given value.


### GetMemory

`func (o *ServerProduct) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ServerProduct) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ServerProduct) SetMemory(v int32)`

SetMemory sets Memory field to given value.


### GetDiskSize

`func (o *ServerProduct) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *ServerProduct) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *ServerProduct) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.


### GetTraffic

`func (o *ServerProduct) GetTraffic() int32`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *ServerProduct) GetTrafficOk() (*int32, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *ServerProduct) SetTraffic(v int32)`

SetTraffic sets Traffic field to given value.


### GetAvailableGenerations

`func (o *ServerProduct) GetAvailableGenerations() []string`

GetAvailableGenerations returns the AvailableGenerations field if non-nil, zero value otherwise.

### GetAvailableGenerationsOk

`func (o *ServerProduct) GetAvailableGenerationsOk() (*[]string, bool)`

GetAvailableGenerationsOk returns a tuple with the AvailableGenerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableGenerations

`func (o *ServerProduct) SetAvailableGenerations(v []string)`

SetAvailableGenerations sets AvailableGenerations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


