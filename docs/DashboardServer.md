# DashboardServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Hostname** | **string** |  | 
**Status** | **string** |  | 
**Project** | **string** |  | 
**Cpus** | **NullableInt32** |  | 
**Memory** | **NullableInt32** |  | 

## Methods

### NewDashboardServer

`func NewDashboardServer(id int32, hostname string, status string, project string, cpus NullableInt32, memory NullableInt32, ) *DashboardServer`

NewDashboardServer instantiates a new DashboardServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardServerWithDefaults

`func NewDashboardServerWithDefaults() *DashboardServer`

NewDashboardServerWithDefaults instantiates a new DashboardServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DashboardServer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardServer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardServer) SetId(v int32)`

SetId sets Id field to given value.


### GetHostname

`func (o *DashboardServer) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DashboardServer) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DashboardServer) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetStatus

`func (o *DashboardServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DashboardServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DashboardServer) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetProject

`func (o *DashboardServer) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *DashboardServer) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *DashboardServer) SetProject(v string)`

SetProject sets Project field to given value.


### GetCpus

`func (o *DashboardServer) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *DashboardServer) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *DashboardServer) SetCpus(v int32)`

SetCpus sets Cpus field to given value.


### SetCpusNil

`func (o *DashboardServer) SetCpusNil(b bool)`

 SetCpusNil sets the value for Cpus to be an explicit nil

### UnsetCpus
`func (o *DashboardServer) UnsetCpus()`

UnsetCpus ensures that no value is present for Cpus, not even an explicit nil
### GetMemory

`func (o *DashboardServer) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *DashboardServer) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *DashboardServer) SetMemory(v int32)`

SetMemory sets Memory field to given value.


### SetMemoryNil

`func (o *DashboardServer) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *DashboardServer) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


