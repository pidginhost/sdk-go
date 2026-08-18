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
**Generation** | Pointer to **string** |  | [optional] [readonly] 
**Machine** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Volumes** | Pointer to [**[]Volume**](Volume.md) |  | [optional] [readonly] 
**Networks** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**FloatingIps** | Pointer to [**[]FloatingIPSummary**](FloatingIPSummary.md) |  | [optional] [readonly] 
**Password** | Pointer to **string** |  | [optional] 
**SshPubKey** | Pointer to **string** | Public key to apply for SSH login. Applying a non-empty key regenerates cloud-init and reboots a running server. Clearing removes the key from future cloud-init data, but does not revoke keys already in the guest. | [optional] 
**Status** | Pointer to [**ResourceStatusEnum**](ResourceStatusEnum.md) |  | [optional] [readonly] 
**Username** | Pointer to **string** |  | [optional] [readonly] 
**DestroyProtection** | Pointer to **bool** | Prevents the server from being destroyed until disabled. | [optional] [readonly] 
**HaEnabled** | Pointer to **bool** | Enables Proxmox HA — automatic restart and migration on node failure. | [optional] [readonly] 
**CustomOs** | Pointer to **bool** | Customer installed their own OS from an ISO; cloud-init features no longer apply | [optional] [readonly] 
**RescueMode** | Pointer to **bool** |  | [optional] [readonly] 
**BootIso** | Pointer to **NullableString** |  | [optional] [readonly] 
**RescueSupported** | Pointer to **bool** |  | [optional] [readonly] 

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

### GetGeneration

`func (o *PatchedServerDetail) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *PatchedServerDetail) GetGenerationOk() (*string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *PatchedServerDetail) SetGeneration(v string)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *PatchedServerDetail) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

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

### GetFloatingIps

`func (o *PatchedServerDetail) GetFloatingIps() []FloatingIPSummary`

GetFloatingIps returns the FloatingIps field if non-nil, zero value otherwise.

### GetFloatingIpsOk

`func (o *PatchedServerDetail) GetFloatingIpsOk() (*[]FloatingIPSummary, bool)`

GetFloatingIpsOk returns a tuple with the FloatingIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIps

`func (o *PatchedServerDetail) SetFloatingIps(v []FloatingIPSummary)`

SetFloatingIps sets FloatingIps field to given value.

### HasFloatingIps

`func (o *PatchedServerDetail) HasFloatingIps() bool`

HasFloatingIps returns a boolean if a field has been set.

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

### GetSshPubKey

`func (o *PatchedServerDetail) GetSshPubKey() string`

GetSshPubKey returns the SshPubKey field if non-nil, zero value otherwise.

### GetSshPubKeyOk

`func (o *PatchedServerDetail) GetSshPubKeyOk() (*string, bool)`

GetSshPubKeyOk returns a tuple with the SshPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPubKey

`func (o *PatchedServerDetail) SetSshPubKey(v string)`

SetSshPubKey sets SshPubKey field to given value.

### HasSshPubKey

`func (o *PatchedServerDetail) HasSshPubKey() bool`

HasSshPubKey returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedServerDetail) GetStatus() ResourceStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedServerDetail) GetStatusOk() (*ResourceStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedServerDetail) SetStatus(v ResourceStatusEnum)`

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

### GetCustomOs

`func (o *PatchedServerDetail) GetCustomOs() bool`

GetCustomOs returns the CustomOs field if non-nil, zero value otherwise.

### GetCustomOsOk

`func (o *PatchedServerDetail) GetCustomOsOk() (*bool, bool)`

GetCustomOsOk returns a tuple with the CustomOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOs

`func (o *PatchedServerDetail) SetCustomOs(v bool)`

SetCustomOs sets CustomOs field to given value.

### HasCustomOs

`func (o *PatchedServerDetail) HasCustomOs() bool`

HasCustomOs returns a boolean if a field has been set.

### GetRescueMode

`func (o *PatchedServerDetail) GetRescueMode() bool`

GetRescueMode returns the RescueMode field if non-nil, zero value otherwise.

### GetRescueModeOk

`func (o *PatchedServerDetail) GetRescueModeOk() (*bool, bool)`

GetRescueModeOk returns a tuple with the RescueMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRescueMode

`func (o *PatchedServerDetail) SetRescueMode(v bool)`

SetRescueMode sets RescueMode field to given value.

### HasRescueMode

`func (o *PatchedServerDetail) HasRescueMode() bool`

HasRescueMode returns a boolean if a field has been set.

### GetBootIso

`func (o *PatchedServerDetail) GetBootIso() string`

GetBootIso returns the BootIso field if non-nil, zero value otherwise.

### GetBootIsoOk

`func (o *PatchedServerDetail) GetBootIsoOk() (*string, bool)`

GetBootIsoOk returns a tuple with the BootIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootIso

`func (o *PatchedServerDetail) SetBootIso(v string)`

SetBootIso sets BootIso field to given value.

### HasBootIso

`func (o *PatchedServerDetail) HasBootIso() bool`

HasBootIso returns a boolean if a field has been set.

### SetBootIsoNil

`func (o *PatchedServerDetail) SetBootIsoNil(b bool)`

 SetBootIsoNil sets the value for BootIso to be an explicit nil

### UnsetBootIso
`func (o *PatchedServerDetail) UnsetBootIso()`

UnsetBootIso ensures that no value is present for BootIso, not even an explicit nil
### GetRescueSupported

`func (o *PatchedServerDetail) GetRescueSupported() bool`

GetRescueSupported returns the RescueSupported field if non-nil, zero value otherwise.

### GetRescueSupportedOk

`func (o *PatchedServerDetail) GetRescueSupportedOk() (*bool, bool)`

GetRescueSupportedOk returns a tuple with the RescueSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRescueSupported

`func (o *PatchedServerDetail) SetRescueSupported(v bool)`

SetRescueSupported sets RescueSupported field to given value.

### HasRescueSupported

`func (o *PatchedServerDetail) HasRescueSupported() bool`

HasRescueSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


