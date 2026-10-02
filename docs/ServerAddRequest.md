# ServerAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | **string** | ID or slug | 
**Package** | **string** | ID or slug | 
**Hostname** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**SshPubKey** | Pointer to **string** | New SSH key | [optional] 
**SshPubKeyId** | Pointer to **string** | ID or fingerprint | [optional] 
**UserData** | Pointer to **string** | Optional startup script. Must be an executable script with a shebang. Stored in cleartext. | [optional] 
**PublicIp** | Pointer to **string** | ID or address of an IPv4 you already own. Aliases accepted: ipv4, public_ipv4. | [optional] 
**NewIpv4** | Pointer to **bool** | Allocate a new IPv4 and attach it before first boot, so the server boots reachable. | [optional] 
**PublicIpv6** | Pointer to **string** | ID or address of an IPv6 you already own. Alias accepted: ipv6. | [optional] 
**NewIpv6** | Pointer to **bool** | Allocate a new IPv6 and attach it before first boot. | [optional] 
**FwRulesSet** | Pointer to **string** | ID or slug | [optional] 
**FwPolicyIn** | Pointer to [**FwPolicyOutEnum**](FwPolicyOutEnum.md) |  | [optional] [default to FWPOLICYOUTENUM_ACCEPT]
**FwPolicyOut** | Pointer to [**FwPolicyOutEnum**](FwPolicyOutEnum.md) |  | [optional] [default to FWPOLICYOUTENUM_ACCEPT]
**PrivateNetwork** | Pointer to **string** | ID or slug | [optional] 
**PrivateAddress** | Pointer to **string** | Leave empty for auto-assign | [optional] 
**ExtraVolumeProduct** | Pointer to **string** | ID or slug | [optional] 
**ExtraVolumeSize** | Pointer to **int32** |  | [optional] [default to 0]
**NoNetworkAcknowledged** | Pointer to **bool** |  | [optional] 
**EnableHa** | Pointer to **bool** |  | [optional] [default to false]
**Generation** | Pointer to **string** |  | [optional] 

## Methods

### NewServerAddRequest

`func NewServerAddRequest(image string, package_ string, ) *ServerAddRequest`

NewServerAddRequest instantiates a new ServerAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerAddRequestWithDefaults

`func NewServerAddRequestWithDefaults() *ServerAddRequest`

NewServerAddRequestWithDefaults instantiates a new ServerAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *ServerAddRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ServerAddRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ServerAddRequest) SetImage(v string)`

SetImage sets Image field to given value.


### GetPackage

`func (o *ServerAddRequest) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *ServerAddRequest) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *ServerAddRequest) SetPackage(v string)`

SetPackage sets Package field to given value.


### GetHostname

`func (o *ServerAddRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ServerAddRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ServerAddRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ServerAddRequest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetProject

`func (o *ServerAddRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ServerAddRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ServerAddRequest) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ServerAddRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPassword

`func (o *ServerAddRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ServerAddRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ServerAddRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ServerAddRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSshPubKey

`func (o *ServerAddRequest) GetSshPubKey() string`

GetSshPubKey returns the SshPubKey field if non-nil, zero value otherwise.

### GetSshPubKeyOk

`func (o *ServerAddRequest) GetSshPubKeyOk() (*string, bool)`

GetSshPubKeyOk returns a tuple with the SshPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPubKey

`func (o *ServerAddRequest) SetSshPubKey(v string)`

SetSshPubKey sets SshPubKey field to given value.

### HasSshPubKey

`func (o *ServerAddRequest) HasSshPubKey() bool`

HasSshPubKey returns a boolean if a field has been set.

### GetSshPubKeyId

`func (o *ServerAddRequest) GetSshPubKeyId() string`

GetSshPubKeyId returns the SshPubKeyId field if non-nil, zero value otherwise.

### GetSshPubKeyIdOk

`func (o *ServerAddRequest) GetSshPubKeyIdOk() (*string, bool)`

GetSshPubKeyIdOk returns a tuple with the SshPubKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPubKeyId

`func (o *ServerAddRequest) SetSshPubKeyId(v string)`

SetSshPubKeyId sets SshPubKeyId field to given value.

### HasSshPubKeyId

`func (o *ServerAddRequest) HasSshPubKeyId() bool`

HasSshPubKeyId returns a boolean if a field has been set.

### GetUserData

`func (o *ServerAddRequest) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *ServerAddRequest) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *ServerAddRequest) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *ServerAddRequest) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetPublicIp

`func (o *ServerAddRequest) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *ServerAddRequest) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *ServerAddRequest) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *ServerAddRequest) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetNewIpv4

`func (o *ServerAddRequest) GetNewIpv4() bool`

GetNewIpv4 returns the NewIpv4 field if non-nil, zero value otherwise.

### GetNewIpv4Ok

`func (o *ServerAddRequest) GetNewIpv4Ok() (*bool, bool)`

GetNewIpv4Ok returns a tuple with the NewIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIpv4

`func (o *ServerAddRequest) SetNewIpv4(v bool)`

SetNewIpv4 sets NewIpv4 field to given value.

### HasNewIpv4

`func (o *ServerAddRequest) HasNewIpv4() bool`

HasNewIpv4 returns a boolean if a field has been set.

### GetPublicIpv6

`func (o *ServerAddRequest) GetPublicIpv6() string`

GetPublicIpv6 returns the PublicIpv6 field if non-nil, zero value otherwise.

### GetPublicIpv6Ok

`func (o *ServerAddRequest) GetPublicIpv6Ok() (*string, bool)`

GetPublicIpv6Ok returns a tuple with the PublicIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpv6

`func (o *ServerAddRequest) SetPublicIpv6(v string)`

SetPublicIpv6 sets PublicIpv6 field to given value.

### HasPublicIpv6

`func (o *ServerAddRequest) HasPublicIpv6() bool`

HasPublicIpv6 returns a boolean if a field has been set.

### GetNewIpv6

`func (o *ServerAddRequest) GetNewIpv6() bool`

GetNewIpv6 returns the NewIpv6 field if non-nil, zero value otherwise.

### GetNewIpv6Ok

`func (o *ServerAddRequest) GetNewIpv6Ok() (*bool, bool)`

GetNewIpv6Ok returns a tuple with the NewIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIpv6

`func (o *ServerAddRequest) SetNewIpv6(v bool)`

SetNewIpv6 sets NewIpv6 field to given value.

### HasNewIpv6

`func (o *ServerAddRequest) HasNewIpv6() bool`

HasNewIpv6 returns a boolean if a field has been set.

### GetFwRulesSet

`func (o *ServerAddRequest) GetFwRulesSet() string`

GetFwRulesSet returns the FwRulesSet field if non-nil, zero value otherwise.

### GetFwRulesSetOk

`func (o *ServerAddRequest) GetFwRulesSetOk() (*string, bool)`

GetFwRulesSetOk returns a tuple with the FwRulesSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwRulesSet

`func (o *ServerAddRequest) SetFwRulesSet(v string)`

SetFwRulesSet sets FwRulesSet field to given value.

### HasFwRulesSet

`func (o *ServerAddRequest) HasFwRulesSet() bool`

HasFwRulesSet returns a boolean if a field has been set.

### GetFwPolicyIn

`func (o *ServerAddRequest) GetFwPolicyIn() FwPolicyOutEnum`

GetFwPolicyIn returns the FwPolicyIn field if non-nil, zero value otherwise.

### GetFwPolicyInOk

`func (o *ServerAddRequest) GetFwPolicyInOk() (*FwPolicyOutEnum, bool)`

GetFwPolicyInOk returns a tuple with the FwPolicyIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwPolicyIn

`func (o *ServerAddRequest) SetFwPolicyIn(v FwPolicyOutEnum)`

SetFwPolicyIn sets FwPolicyIn field to given value.

### HasFwPolicyIn

`func (o *ServerAddRequest) HasFwPolicyIn() bool`

HasFwPolicyIn returns a boolean if a field has been set.

### GetFwPolicyOut

`func (o *ServerAddRequest) GetFwPolicyOut() FwPolicyOutEnum`

GetFwPolicyOut returns the FwPolicyOut field if non-nil, zero value otherwise.

### GetFwPolicyOutOk

`func (o *ServerAddRequest) GetFwPolicyOutOk() (*FwPolicyOutEnum, bool)`

GetFwPolicyOutOk returns a tuple with the FwPolicyOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwPolicyOut

`func (o *ServerAddRequest) SetFwPolicyOut(v FwPolicyOutEnum)`

SetFwPolicyOut sets FwPolicyOut field to given value.

### HasFwPolicyOut

`func (o *ServerAddRequest) HasFwPolicyOut() bool`

HasFwPolicyOut returns a boolean if a field has been set.

### GetPrivateNetwork

`func (o *ServerAddRequest) GetPrivateNetwork() string`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *ServerAddRequest) GetPrivateNetworkOk() (*string, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *ServerAddRequest) SetPrivateNetwork(v string)`

SetPrivateNetwork sets PrivateNetwork field to given value.

### HasPrivateNetwork

`func (o *ServerAddRequest) HasPrivateNetwork() bool`

HasPrivateNetwork returns a boolean if a field has been set.

### GetPrivateAddress

`func (o *ServerAddRequest) GetPrivateAddress() string`

GetPrivateAddress returns the PrivateAddress field if non-nil, zero value otherwise.

### GetPrivateAddressOk

`func (o *ServerAddRequest) GetPrivateAddressOk() (*string, bool)`

GetPrivateAddressOk returns a tuple with the PrivateAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateAddress

`func (o *ServerAddRequest) SetPrivateAddress(v string)`

SetPrivateAddress sets PrivateAddress field to given value.

### HasPrivateAddress

`func (o *ServerAddRequest) HasPrivateAddress() bool`

HasPrivateAddress returns a boolean if a field has been set.

### GetExtraVolumeProduct

`func (o *ServerAddRequest) GetExtraVolumeProduct() string`

GetExtraVolumeProduct returns the ExtraVolumeProduct field if non-nil, zero value otherwise.

### GetExtraVolumeProductOk

`func (o *ServerAddRequest) GetExtraVolumeProductOk() (*string, bool)`

GetExtraVolumeProductOk returns a tuple with the ExtraVolumeProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraVolumeProduct

`func (o *ServerAddRequest) SetExtraVolumeProduct(v string)`

SetExtraVolumeProduct sets ExtraVolumeProduct field to given value.

### HasExtraVolumeProduct

`func (o *ServerAddRequest) HasExtraVolumeProduct() bool`

HasExtraVolumeProduct returns a boolean if a field has been set.

### GetExtraVolumeSize

`func (o *ServerAddRequest) GetExtraVolumeSize() int32`

GetExtraVolumeSize returns the ExtraVolumeSize field if non-nil, zero value otherwise.

### GetExtraVolumeSizeOk

`func (o *ServerAddRequest) GetExtraVolumeSizeOk() (*int32, bool)`

GetExtraVolumeSizeOk returns a tuple with the ExtraVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraVolumeSize

`func (o *ServerAddRequest) SetExtraVolumeSize(v int32)`

SetExtraVolumeSize sets ExtraVolumeSize field to given value.

### HasExtraVolumeSize

`func (o *ServerAddRequest) HasExtraVolumeSize() bool`

HasExtraVolumeSize returns a boolean if a field has been set.

### GetNoNetworkAcknowledged

`func (o *ServerAddRequest) GetNoNetworkAcknowledged() bool`

GetNoNetworkAcknowledged returns the NoNetworkAcknowledged field if non-nil, zero value otherwise.

### GetNoNetworkAcknowledgedOk

`func (o *ServerAddRequest) GetNoNetworkAcknowledgedOk() (*bool, bool)`

GetNoNetworkAcknowledgedOk returns a tuple with the NoNetworkAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoNetworkAcknowledged

`func (o *ServerAddRequest) SetNoNetworkAcknowledged(v bool)`

SetNoNetworkAcknowledged sets NoNetworkAcknowledged field to given value.

### HasNoNetworkAcknowledged

`func (o *ServerAddRequest) HasNoNetworkAcknowledged() bool`

HasNoNetworkAcknowledged returns a boolean if a field has been set.

### GetEnableHa

`func (o *ServerAddRequest) GetEnableHa() bool`

GetEnableHa returns the EnableHa field if non-nil, zero value otherwise.

### GetEnableHaOk

`func (o *ServerAddRequest) GetEnableHaOk() (*bool, bool)`

GetEnableHaOk returns a tuple with the EnableHa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHa

`func (o *ServerAddRequest) SetEnableHa(v bool)`

SetEnableHa sets EnableHa field to given value.

### HasEnableHa

`func (o *ServerAddRequest) HasEnableHa() bool`

HasEnableHa returns a boolean if a field has been set.

### GetGeneration

`func (o *ServerAddRequest) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *ServerAddRequest) GetGenerationOk() (*string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *ServerAddRequest) SetGeneration(v string)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *ServerAddRequest) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


