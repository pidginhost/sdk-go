# PatchedServerDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Hostname** | Pointer to **string** |  | [optional] [readonly] 
**Project** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] [readonly] 
**Package** | Pointer to **string** |  | [optional] [readonly] 
**Cpus** | Pointer to **int32** |  | [optional] [readonly] 
**Memory** | Pointer to **int32** |  | [optional] [readonly] 
**DiskSize** | Pointer to **int32** |  | [optional] [readonly] 
**Machine** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Volumes** | Pointer to [**[]Volume**](Volume.md) |  | [optional] [readonly] 
**Networks** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Password** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**StatusA57Enum**](StatusA57Enum.md) |  | [optional] [readonly] 
**Username** | Pointer to **string** |  | [optional] [readonly] 
**DestroyProtection** | Pointer to **bool** | Prevents the server from being destroyed until disabled. | [optional] [readonly] 
**HaEnabled** | Pointer to **bool** | Enables Proxmox HA — automatic restart and migration on node failure. | [optional] [readonly] 

## Methods

### NewPatchedServerDetail

`func NewPatchedServerDetail() *PatchedServerDetail`

NewPatchedServerDetail instantiates a new PatchedServerDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedServerDetailWithDefaults

`func NewPatchedServerDetailWithDefaults() *PatchedServerDetail`

NewPatchedServerDetailWithDefaults instantiates a new PatchedServerDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedServerDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedServerDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedServerDetail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedServerDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHostname

`func (o *PatchedServerDetail) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *PatchedServerDetail) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *PatchedServerDetail) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *PatchedServerDetail) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetProject

`func (o *PatchedServerDetail) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PatchedServerDetail) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PatchedServerDetail) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PatchedServerDetail) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetImage

`func (o *PatchedServerDetail) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PatchedServerDetail) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PatchedServerDetail) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *PatchedServerDetail) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetPackage

`func (o *PatchedServerDetail) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *PatchedServerDetail) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *PatchedServerDetail) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *PatchedServerDetail) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetCpus

`func (o *PatchedServerDetail) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *PatchedServerDetail) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *PatchedServerDetail) SetCpus(v int32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *PatchedServerDetail) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetMemory

`func (o *PatchedServerDetail) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *PatchedServerDetail) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *PatchedServerDetail) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *PatchedServerDetail) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetDiskSize

`func (o *PatchedServerDetail) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *PatchedServerDetail) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *PatchedServerDetail) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *PatchedServerDetail) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetMachine

`func (o *PatchedServerDetail) GetMachine() map[string]interface{}`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *PatchedServerDetail) GetMachineOk() (*map[string]interface{}, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *PatchedServerDetail) SetMachine(v map[string]interface{})`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *PatchedServerDetail) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetVolumes

`func (o *PatchedServerDetail) GetVolumes() []Volume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *PatchedServerDetail) GetVolumesOk() (*[]Volume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *PatchedServerDetail) SetVolumes(v []Volume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *PatchedServerDetail) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetNetworks

`func (o *PatchedServerDetail) GetNetworks() map[string]interface{}`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *PatchedServerDetail) GetNetworksOk() (*map[string]interface{}, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *PatchedServerDetail) SetNetworks(v map[string]interface{})`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *PatchedServerDetail) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetPassword

`func (o *PatchedServerDetail) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatchedServerDetail) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatchedServerDetail) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatchedServerDetail) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedServerDetail) GetStatus() StatusA57Enum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedServerDetail) GetStatusOk() (*StatusA57Enum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedServerDetail) SetStatus(v StatusA57Enum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedServerDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsername

`func (o *PatchedServerDetail) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PatchedServerDetail) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PatchedServerDetail) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PatchedServerDetail) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDestroyProtection

`func (o *PatchedServerDetail) GetDestroyProtection() bool`

GetDestroyProtection returns the DestroyProtection field if non-nil, zero value otherwise.

### GetDestroyProtectionOk

`func (o *PatchedServerDetail) GetDestroyProtectionOk() (*bool, bool)`

GetDestroyProtectionOk returns a tuple with the DestroyProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyProtection

`func (o *PatchedServerDetail) SetDestroyProtection(v bool)`

SetDestroyProtection sets DestroyProtection field to given value.

### HasDestroyProtection

`func (o *PatchedServerDetail) HasDestroyProtection() bool`

HasDestroyProtection returns a boolean if a field has been set.

### GetHaEnabled

`func (o *PatchedServerDetail) GetHaEnabled() bool`

GetHaEnabled returns the HaEnabled field if non-nil, zero value otherwise.

### GetHaEnabledOk

`func (o *PatchedServerDetail) GetHaEnabledOk() (*bool, bool)`

GetHaEnabledOk returns a tuple with the HaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaEnabled

`func (o *PatchedServerDetail) SetHaEnabled(v bool)`

SetHaEnabled sets HaEnabled field to given value.

### HasHaEnabled

`func (o *PatchedServerDetail) HasHaEnabled() bool`

HasHaEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


