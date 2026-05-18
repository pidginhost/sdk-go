# ServerDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Hostname** | **string** |  | [readonly] 
**Project** | Pointer to **string** |  | [optional] 
**Image** | **string** |  | [readonly] 
**Package** | **string** |  | [readonly] 
**Cpus** | **int32** |  | [readonly] 
**Memory** | **int32** |  | [readonly] 
**DiskSize** | **int32** |  | [readonly] 
**Generation** | **string** |  | [readonly] 
**Machine** | **map[string]interface{}** |  | [readonly] 
**Volumes** | [**[]Volume**](Volume.md) |  | [readonly] 
**Networks** | **map[string]interface{}** |  | [readonly] 
**Password** | Pointer to **string** |  | [optional] 
**Status** | [**StatusA57Enum**](StatusA57Enum.md) |  | [readonly] 
**Username** | **string** |  | [readonly] 
**DestroyProtection** | **bool** | Prevents the server from being destroyed until disabled. | [readonly] 
**HaEnabled** | **bool** | Enables Proxmox HA — automatic restart and migration on node failure. | [readonly] 

## Methods

### NewServerDetail

`func NewServerDetail(id int32, hostname string, image string, package_ string, cpus int32, memory int32, diskSize int32, generation string, machine map[string]interface{}, volumes []Volume, networks map[string]interface{}, status StatusA57Enum, username string, destroyProtection bool, haEnabled bool, ) *ServerDetail`

NewServerDetail instantiates a new ServerDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerDetailWithDefaults

`func NewServerDetailWithDefaults() *ServerDetail`

NewServerDetailWithDefaults instantiates a new ServerDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerDetail) SetId(v int32)`

SetId sets Id field to given value.


### GetHostname

`func (o *ServerDetail) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ServerDetail) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ServerDetail) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetProject

`func (o *ServerDetail) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ServerDetail) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ServerDetail) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ServerDetail) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetImage

`func (o *ServerDetail) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ServerDetail) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ServerDetail) SetImage(v string)`

SetImage sets Image field to given value.


### GetPackage

`func (o *ServerDetail) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *ServerDetail) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *ServerDetail) SetPackage(v string)`

SetPackage sets Package field to given value.


### GetCpus

`func (o *ServerDetail) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *ServerDetail) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *ServerDetail) SetCpus(v int32)`

SetCpus sets Cpus field to given value.


### GetMemory

`func (o *ServerDetail) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ServerDetail) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ServerDetail) SetMemory(v int32)`

SetMemory sets Memory field to given value.


### GetDiskSize

`func (o *ServerDetail) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *ServerDetail) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *ServerDetail) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.


### GetGeneration

`func (o *ServerDetail) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *ServerDetail) GetGenerationOk() (*string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *ServerDetail) SetGeneration(v string)`

SetGeneration sets Generation field to given value.


### GetMachine

`func (o *ServerDetail) GetMachine() map[string]interface{}`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *ServerDetail) GetMachineOk() (*map[string]interface{}, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *ServerDetail) SetMachine(v map[string]interface{})`

SetMachine sets Machine field to given value.


### GetVolumes

`func (o *ServerDetail) GetVolumes() []Volume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ServerDetail) GetVolumesOk() (*[]Volume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ServerDetail) SetVolumes(v []Volume)`

SetVolumes sets Volumes field to given value.


### GetNetworks

`func (o *ServerDetail) GetNetworks() map[string]interface{}`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *ServerDetail) GetNetworksOk() (*map[string]interface{}, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *ServerDetail) SetNetworks(v map[string]interface{})`

SetNetworks sets Networks field to given value.


### GetPassword

`func (o *ServerDetail) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ServerDetail) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ServerDetail) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ServerDetail) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetStatus

`func (o *ServerDetail) GetStatus() StatusA57Enum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerDetail) GetStatusOk() (*StatusA57Enum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerDetail) SetStatus(v StatusA57Enum)`

SetStatus sets Status field to given value.


### GetUsername

`func (o *ServerDetail) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ServerDetail) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ServerDetail) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetDestroyProtection

`func (o *ServerDetail) GetDestroyProtection() bool`

GetDestroyProtection returns the DestroyProtection field if non-nil, zero value otherwise.

### GetDestroyProtectionOk

`func (o *ServerDetail) GetDestroyProtectionOk() (*bool, bool)`

GetDestroyProtectionOk returns a tuple with the DestroyProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyProtection

`func (o *ServerDetail) SetDestroyProtection(v bool)`

SetDestroyProtection sets DestroyProtection field to given value.


### GetHaEnabled

`func (o *ServerDetail) GetHaEnabled() bool`

GetHaEnabled returns the HaEnabled field if non-nil, zero value otherwise.

### GetHaEnabledOk

`func (o *ServerDetail) GetHaEnabledOk() (*bool, bool)`

GetHaEnabledOk returns a tuple with the HaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaEnabled

`func (o *ServerDetail) SetHaEnabled(v bool)`

SetHaEnabled sets HaEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


