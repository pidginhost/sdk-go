# Server

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Hostname** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Image** | **string** |  | [readonly] 
**Package** | **string** |  | [readonly] 
**Cpus** | **string** |  | [readonly] 
**Memory** | **string** |  | [readonly] 
**DiskSize** | **string** |  | [readonly] 
**Status** | Pointer to [**StatusA57Enum**](StatusA57Enum.md) |  | [optional] 
**DestroyProtection** | **bool** | Prevents the server from being destroyed until disabled. | [readonly] 
**HaEnabled** | **bool** | Enables Proxmox HA — automatic restart and migration on node failure. | [readonly] 
**Networks** | **map[string]interface{}** |  | [readonly] 

## Methods

### NewServer

`func NewServer(id int32, image string, package_ string, cpus string, memory string, diskSize string, destroyProtection bool, haEnabled bool, networks map[string]interface{}, ) *Server`

NewServer instantiates a new Server object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerWithDefaults

`func NewServerWithDefaults() *Server`

NewServerWithDefaults instantiates a new Server object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Server) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Server) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Server) SetId(v int32)`

SetId sets Id field to given value.


### GetHostname

`func (o *Server) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Server) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Server) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *Server) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetProject

`func (o *Server) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Server) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Server) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Server) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetImage

`func (o *Server) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Server) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Server) SetImage(v string)`

SetImage sets Image field to given value.


### GetPackage

`func (o *Server) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *Server) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *Server) SetPackage(v string)`

SetPackage sets Package field to given value.


### GetCpus

`func (o *Server) GetCpus() string`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *Server) GetCpusOk() (*string, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *Server) SetCpus(v string)`

SetCpus sets Cpus field to given value.


### GetMemory

`func (o *Server) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Server) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Server) SetMemory(v string)`

SetMemory sets Memory field to given value.


### GetDiskSize

`func (o *Server) GetDiskSize() string`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *Server) GetDiskSizeOk() (*string, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *Server) SetDiskSize(v string)`

SetDiskSize sets DiskSize field to given value.


### GetStatus

`func (o *Server) GetStatus() StatusA57Enum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Server) GetStatusOk() (*StatusA57Enum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Server) SetStatus(v StatusA57Enum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Server) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDestroyProtection

`func (o *Server) GetDestroyProtection() bool`

GetDestroyProtection returns the DestroyProtection field if non-nil, zero value otherwise.

### GetDestroyProtectionOk

`func (o *Server) GetDestroyProtectionOk() (*bool, bool)`

GetDestroyProtectionOk returns a tuple with the DestroyProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyProtection

`func (o *Server) SetDestroyProtection(v bool)`

SetDestroyProtection sets DestroyProtection field to given value.


### GetHaEnabled

`func (o *Server) GetHaEnabled() bool`

GetHaEnabled returns the HaEnabled field if non-nil, zero value otherwise.

### GetHaEnabledOk

`func (o *Server) GetHaEnabledOk() (*bool, bool)`

GetHaEnabledOk returns a tuple with the HaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaEnabled

`func (o *Server) SetHaEnabled(v bool)`

SetHaEnabled sets HaEnabled field to given value.


### GetNetworks

`func (o *Server) GetNetworks() map[string]interface{}`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *Server) GetNetworksOk() (*map[string]interface{}, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *Server) SetNetworks(v map[string]interface{})`

SetNetworks sets Networks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


