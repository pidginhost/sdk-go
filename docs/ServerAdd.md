# ServerAdd

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
**PublicIp** | Pointer to **string** | ID or slug | [optional] 
**NewIpv4** | Pointer to **bool** |  | [optional] 
**PublicIpv6** | Pointer to **string** | ID or slug | [optional] 
**NewIpv6** | Pointer to **bool** |  | [optional] 
**FwRulesSet** | Pointer to **string** | ID or slug | [optional] 
**FwPolicyIn** | Pointer to [**FwPolicyOutEnum**](FwPolicyOutEnum.md) |  | [optional] [default to FWPOLICYOUTENUM_ACCEPT]
**FwPolicyOut** | Pointer to [**FwPolicyOutEnum**](FwPolicyOutEnum.md) |  | [optional] [default to FWPOLICYOUTENUM_ACCEPT]
**PrivateNetwork** | Pointer to **string** | ID or slug | [optional] 
**PrivateAddress** | Pointer to **string** | Leave empty for auto-assign | [optional] 
**ExtraVolumeProduct** | Pointer to **string** | ID or slug | [optional] 
**ExtraVolumeSize** | Pointer to **int32** |  | [optional] [default to 0]
**NoNetworkAcknowledged** | Pointer to **bool** |  | [optional] 
**EnableHa** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewServerAdd

`func NewServerAdd(image string, package_ string, ) *ServerAdd`

NewServerAdd instantiates a new ServerAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerAddWithDefaults

`func NewServerAddWithDefaults() *ServerAdd`

NewServerAddWithDefaults instantiates a new ServerAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *ServerAdd) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ServerAdd) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ServerAdd) SetImage(v string)`

SetImage sets Image field to given value.


### GetPackage

`func (o *ServerAdd) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *ServerAdd) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *ServerAdd) SetPackage(v string)`

SetPackage sets Package field to given value.


### GetHostname

`func (o *ServerAdd) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ServerAdd) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ServerAdd) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ServerAdd) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetProject

`func (o *ServerAdd) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ServerAdd) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ServerAdd) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ServerAdd) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPassword

`func (o *ServerAdd) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ServerAdd) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ServerAdd) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ServerAdd) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSshPubKey

`func (o *ServerAdd) GetSshPubKey() string`

GetSshPubKey returns the SshPubKey field if non-nil, zero value otherwise.

### GetSshPubKeyOk

`func (o *ServerAdd) GetSshPubKeyOk() (*string, bool)`

GetSshPubKeyOk returns a tuple with the SshPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPubKey

`func (o *ServerAdd) SetSshPubKey(v string)`

SetSshPubKey sets SshPubKey field to given value.

### HasSshPubKey

`func (o *ServerAdd) HasSshPubKey() bool`

HasSshPubKey returns a boolean if a field has been set.

### GetSshPubKeyId

`func (o *ServerAdd) GetSshPubKeyId() string`

GetSshPubKeyId returns the SshPubKeyId field if non-nil, zero value otherwise.

### GetSshPubKeyIdOk

`func (o *ServerAdd) GetSshPubKeyIdOk() (*string, bool)`

GetSshPubKeyIdOk returns a tuple with the SshPubKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPubKeyId

`func (o *ServerAdd) SetSshPubKeyId(v string)`

SetSshPubKeyId sets SshPubKeyId field to given value.

### HasSshPubKeyId

`func (o *ServerAdd) HasSshPubKeyId() bool`

HasSshPubKeyId returns a boolean if a field has been set.

### GetPublicIp

`func (o *ServerAdd) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *ServerAdd) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *ServerAdd) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *ServerAdd) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetNewIpv4

`func (o *ServerAdd) GetNewIpv4() bool`

GetNewIpv4 returns the NewIpv4 field if non-nil, zero value otherwise.

### GetNewIpv4Ok

`func (o *ServerAdd) GetNewIpv4Ok() (*bool, bool)`

GetNewIpv4Ok returns a tuple with the NewIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIpv4

`func (o *ServerAdd) SetNewIpv4(v bool)`

SetNewIpv4 sets NewIpv4 field to given value.

### HasNewIpv4

`func (o *ServerAdd) HasNewIpv4() bool`

HasNewIpv4 returns a boolean if a field has been set.

### GetPublicIpv6

`func (o *ServerAdd) GetPublicIpv6() string`

GetPublicIpv6 returns the PublicIpv6 field if non-nil, zero value otherwise.

### GetPublicIpv6Ok

`func (o *ServerAdd) GetPublicIpv6Ok() (*string, bool)`

GetPublicIpv6Ok returns a tuple with the PublicIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpv6

`func (o *ServerAdd) SetPublicIpv6(v string)`

SetPublicIpv6 sets PublicIpv6 field to given value.

### HasPublicIpv6

`func (o *ServerAdd) HasPublicIpv6() bool`

HasPublicIpv6 returns a boolean if a field has been set.

### GetNewIpv6

`func (o *ServerAdd) GetNewIpv6() bool`

GetNewIpv6 returns the NewIpv6 field if non-nil, zero value otherwise.

### GetNewIpv6Ok

`func (o *ServerAdd) GetNewIpv6Ok() (*bool, bool)`

GetNewIpv6Ok returns a tuple with the NewIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIpv6

`func (o *ServerAdd) SetNewIpv6(v bool)`

SetNewIpv6 sets NewIpv6 field to given value.

### HasNewIpv6

`func (o *ServerAdd) HasNewIpv6() bool`

HasNewIpv6 returns a boolean if a field has been set.

### GetFwRulesSet

`func (o *ServerAdd) GetFwRulesSet() string`

GetFwRulesSet returns the FwRulesSet field if non-nil, zero value otherwise.

### GetFwRulesSetOk

`func (o *ServerAdd) GetFwRulesSetOk() (*string, bool)`

GetFwRulesSetOk returns a tuple with the FwRulesSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwRulesSet

`func (o *ServerAdd) SetFwRulesSet(v string)`

SetFwRulesSet sets FwRulesSet field to given value.

### HasFwRulesSet

`func (o *ServerAdd) HasFwRulesSet() bool`

HasFwRulesSet returns a boolean if a field has been set.

### GetFwPolicyIn

`func (o *ServerAdd) GetFwPolicyIn() FwPolicyOutEnum`

GetFwPolicyIn returns the FwPolicyIn field if non-nil, zero value otherwise.

### GetFwPolicyInOk

`func (o *ServerAdd) GetFwPolicyInOk() (*FwPolicyOutEnum, bool)`

GetFwPolicyInOk returns a tuple with the FwPolicyIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwPolicyIn

`func (o *ServerAdd) SetFwPolicyIn(v FwPolicyOutEnum)`

SetFwPolicyIn sets FwPolicyIn field to given value.

### HasFwPolicyIn

`func (o *ServerAdd) HasFwPolicyIn() bool`

HasFwPolicyIn returns a boolean if a field has been set.

### GetFwPolicyOut

`func (o *ServerAdd) GetFwPolicyOut() FwPolicyOutEnum`

GetFwPolicyOut returns the FwPolicyOut field if non-nil, zero value otherwise.

### GetFwPolicyOutOk

`func (o *ServerAdd) GetFwPolicyOutOk() (*FwPolicyOutEnum, bool)`

GetFwPolicyOutOk returns a tuple with the FwPolicyOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwPolicyOut

`func (o *ServerAdd) SetFwPolicyOut(v FwPolicyOutEnum)`

SetFwPolicyOut sets FwPolicyOut field to given value.

### HasFwPolicyOut

`func (o *ServerAdd) HasFwPolicyOut() bool`

HasFwPolicyOut returns a boolean if a field has been set.

### GetPrivateNetwork

`func (o *ServerAdd) GetPrivateNetwork() string`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *ServerAdd) GetPrivateNetworkOk() (*string, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *ServerAdd) SetPrivateNetwork(v string)`

SetPrivateNetwork sets PrivateNetwork field to given value.

### HasPrivateNetwork

`func (o *ServerAdd) HasPrivateNetwork() bool`

HasPrivateNetwork returns a boolean if a field has been set.

### GetPrivateAddress

`func (o *ServerAdd) GetPrivateAddress() string`

GetPrivateAddress returns the PrivateAddress field if non-nil, zero value otherwise.

### GetPrivateAddressOk

`func (o *ServerAdd) GetPrivateAddressOk() (*string, bool)`

GetPrivateAddressOk returns a tuple with the PrivateAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateAddress

`func (o *ServerAdd) SetPrivateAddress(v string)`

SetPrivateAddress sets PrivateAddress field to given value.

### HasPrivateAddress

`func (o *ServerAdd) HasPrivateAddress() bool`

HasPrivateAddress returns a boolean if a field has been set.

### GetExtraVolumeProduct

`func (o *ServerAdd) GetExtraVolumeProduct() string`

GetExtraVolumeProduct returns the ExtraVolumeProduct field if non-nil, zero value otherwise.

### GetExtraVolumeProductOk

`func (o *ServerAdd) GetExtraVolumeProductOk() (*string, bool)`

GetExtraVolumeProductOk returns a tuple with the ExtraVolumeProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraVolumeProduct

`func (o *ServerAdd) SetExtraVolumeProduct(v string)`

SetExtraVolumeProduct sets ExtraVolumeProduct field to given value.

### HasExtraVolumeProduct

`func (o *ServerAdd) HasExtraVolumeProduct() bool`

HasExtraVolumeProduct returns a boolean if a field has been set.

### GetExtraVolumeSize

`func (o *ServerAdd) GetExtraVolumeSize() int32`

GetExtraVolumeSize returns the ExtraVolumeSize field if non-nil, zero value otherwise.

### GetExtraVolumeSizeOk

`func (o *ServerAdd) GetExtraVolumeSizeOk() (*int32, bool)`

GetExtraVolumeSizeOk returns a tuple with the ExtraVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraVolumeSize

`func (o *ServerAdd) SetExtraVolumeSize(v int32)`

SetExtraVolumeSize sets ExtraVolumeSize field to given value.

### HasExtraVolumeSize

`func (o *ServerAdd) HasExtraVolumeSize() bool`

HasExtraVolumeSize returns a boolean if a field has been set.

### GetNoNetworkAcknowledged

`func (o *ServerAdd) GetNoNetworkAcknowledged() bool`

GetNoNetworkAcknowledged returns the NoNetworkAcknowledged field if non-nil, zero value otherwise.

### GetNoNetworkAcknowledgedOk

`func (o *ServerAdd) GetNoNetworkAcknowledgedOk() (*bool, bool)`

GetNoNetworkAcknowledgedOk returns a tuple with the NoNetworkAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoNetworkAcknowledged

`func (o *ServerAdd) SetNoNetworkAcknowledged(v bool)`

SetNoNetworkAcknowledged sets NoNetworkAcknowledged field to given value.

### HasNoNetworkAcknowledged

`func (o *ServerAdd) HasNoNetworkAcknowledged() bool`

HasNoNetworkAcknowledged returns a boolean if a field has been set.

### GetEnableHa

`func (o *ServerAdd) GetEnableHa() bool`

GetEnableHa returns the EnableHa field if non-nil, zero value otherwise.

### GetEnableHaOk

`func (o *ServerAdd) GetEnableHaOk() (*bool, bool)`

GetEnableHaOk returns a tuple with the EnableHa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHa

`func (o *ServerAdd) SetEnableHa(v bool)`

SetEnableHa sets EnableHa field to given value.

### HasEnableHa

`func (o *ServerAdd) HasEnableHa() bool`

HasEnableHa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


