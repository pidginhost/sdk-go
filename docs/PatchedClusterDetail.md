# PatchedClusterDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Status** | Pointer to [**StatusA57Enum**](StatusA57Enum.md) |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**ClusterType** | Pointer to **string** |  | [optional] [readonly] 
**KubeVersion** | Pointer to **string** |  | [optional] [readonly] 
**PricePerMonth** | Pointer to **float64** |  | [optional] 
**PricePerHour** | Pointer to **string** |  | [optional] [readonly] 
**Features** | Pointer to [**[]FeaturesEnum**](FeaturesEnum.md) |  | [optional] 
**FeaturesReady** | Pointer to **bool** |  | [optional] [readonly] 
**KubeconfigValidUntil** | Pointer to **string** |  | [optional] [readonly] 
**Ipv4Address** | Pointer to **string** |  | [optional] [readonly] 
**Protected** | Pointer to **bool** |  | [optional] 
**TalosVersion** | Pointer to **string** |  | [optional] [readonly] 
**TalosUpgradeAvailable** | Pointer to **string** |  | [optional] [readonly] 
**TalosNextVersion** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPatchedClusterDetail

`func NewPatchedClusterDetail() *PatchedClusterDetail`

NewPatchedClusterDetail instantiates a new PatchedClusterDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedClusterDetailWithDefaults

`func NewPatchedClusterDetailWithDefaults() *PatchedClusterDetail`

NewPatchedClusterDetailWithDefaults instantiates a new PatchedClusterDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedClusterDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedClusterDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedClusterDetail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedClusterDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedClusterDetail) GetStatus() StatusA57Enum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedClusterDetail) GetStatusOk() (*StatusA57Enum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedClusterDetail) SetStatus(v StatusA57Enum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedClusterDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetName

`func (o *PatchedClusterDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedClusterDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedClusterDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedClusterDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClusterType

`func (o *PatchedClusterDetail) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *PatchedClusterDetail) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *PatchedClusterDetail) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *PatchedClusterDetail) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.

### GetKubeVersion

`func (o *PatchedClusterDetail) GetKubeVersion() string`

GetKubeVersion returns the KubeVersion field if non-nil, zero value otherwise.

### GetKubeVersionOk

`func (o *PatchedClusterDetail) GetKubeVersionOk() (*string, bool)`

GetKubeVersionOk returns a tuple with the KubeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeVersion

`func (o *PatchedClusterDetail) SetKubeVersion(v string)`

SetKubeVersion sets KubeVersion field to given value.

### HasKubeVersion

`func (o *PatchedClusterDetail) HasKubeVersion() bool`

HasKubeVersion returns a boolean if a field has been set.

### GetPricePerMonth

`func (o *PatchedClusterDetail) GetPricePerMonth() float64`

GetPricePerMonth returns the PricePerMonth field if non-nil, zero value otherwise.

### GetPricePerMonthOk

`func (o *PatchedClusterDetail) GetPricePerMonthOk() (*float64, bool)`

GetPricePerMonthOk returns a tuple with the PricePerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerMonth

`func (o *PatchedClusterDetail) SetPricePerMonth(v float64)`

SetPricePerMonth sets PricePerMonth field to given value.

### HasPricePerMonth

`func (o *PatchedClusterDetail) HasPricePerMonth() bool`

HasPricePerMonth returns a boolean if a field has been set.

### GetPricePerHour

`func (o *PatchedClusterDetail) GetPricePerHour() string`

GetPricePerHour returns the PricePerHour field if non-nil, zero value otherwise.

### GetPricePerHourOk

`func (o *PatchedClusterDetail) GetPricePerHourOk() (*string, bool)`

GetPricePerHourOk returns a tuple with the PricePerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerHour

`func (o *PatchedClusterDetail) SetPricePerHour(v string)`

SetPricePerHour sets PricePerHour field to given value.

### HasPricePerHour

`func (o *PatchedClusterDetail) HasPricePerHour() bool`

HasPricePerHour returns a boolean if a field has been set.

### GetFeatures

`func (o *PatchedClusterDetail) GetFeatures() []FeaturesEnum`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *PatchedClusterDetail) GetFeaturesOk() (*[]FeaturesEnum, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *PatchedClusterDetail) SetFeatures(v []FeaturesEnum)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *PatchedClusterDetail) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFeaturesReady

`func (o *PatchedClusterDetail) GetFeaturesReady() bool`

GetFeaturesReady returns the FeaturesReady field if non-nil, zero value otherwise.

### GetFeaturesReadyOk

`func (o *PatchedClusterDetail) GetFeaturesReadyOk() (*bool, bool)`

GetFeaturesReadyOk returns a tuple with the FeaturesReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesReady

`func (o *PatchedClusterDetail) SetFeaturesReady(v bool)`

SetFeaturesReady sets FeaturesReady field to given value.

### HasFeaturesReady

`func (o *PatchedClusterDetail) HasFeaturesReady() bool`

HasFeaturesReady returns a boolean if a field has been set.

### GetKubeconfigValidUntil

`func (o *PatchedClusterDetail) GetKubeconfigValidUntil() string`

GetKubeconfigValidUntil returns the KubeconfigValidUntil field if non-nil, zero value otherwise.

### GetKubeconfigValidUntilOk

`func (o *PatchedClusterDetail) GetKubeconfigValidUntilOk() (*string, bool)`

GetKubeconfigValidUntilOk returns a tuple with the KubeconfigValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeconfigValidUntil

`func (o *PatchedClusterDetail) SetKubeconfigValidUntil(v string)`

SetKubeconfigValidUntil sets KubeconfigValidUntil field to given value.

### HasKubeconfigValidUntil

`func (o *PatchedClusterDetail) HasKubeconfigValidUntil() bool`

HasKubeconfigValidUntil returns a boolean if a field has been set.

### GetIpv4Address

`func (o *PatchedClusterDetail) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *PatchedClusterDetail) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *PatchedClusterDetail) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *PatchedClusterDetail) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetProtected

`func (o *PatchedClusterDetail) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *PatchedClusterDetail) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *PatchedClusterDetail) SetProtected(v bool)`

SetProtected sets Protected field to given value.

### HasProtected

`func (o *PatchedClusterDetail) HasProtected() bool`

HasProtected returns a boolean if a field has been set.

### GetTalosVersion

`func (o *PatchedClusterDetail) GetTalosVersion() string`

GetTalosVersion returns the TalosVersion field if non-nil, zero value otherwise.

### GetTalosVersionOk

`func (o *PatchedClusterDetail) GetTalosVersionOk() (*string, bool)`

GetTalosVersionOk returns a tuple with the TalosVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTalosVersion

`func (o *PatchedClusterDetail) SetTalosVersion(v string)`

SetTalosVersion sets TalosVersion field to given value.

### HasTalosVersion

`func (o *PatchedClusterDetail) HasTalosVersion() bool`

HasTalosVersion returns a boolean if a field has been set.

### GetTalosUpgradeAvailable

`func (o *PatchedClusterDetail) GetTalosUpgradeAvailable() string`

GetTalosUpgradeAvailable returns the TalosUpgradeAvailable field if non-nil, zero value otherwise.

### GetTalosUpgradeAvailableOk

`func (o *PatchedClusterDetail) GetTalosUpgradeAvailableOk() (*string, bool)`

GetTalosUpgradeAvailableOk returns a tuple with the TalosUpgradeAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTalosUpgradeAvailable

`func (o *PatchedClusterDetail) SetTalosUpgradeAvailable(v string)`

SetTalosUpgradeAvailable sets TalosUpgradeAvailable field to given value.

### HasTalosUpgradeAvailable

`func (o *PatchedClusterDetail) HasTalosUpgradeAvailable() bool`

HasTalosUpgradeAvailable returns a boolean if a field has been set.

### GetTalosNextVersion

`func (o *PatchedClusterDetail) GetTalosNextVersion() string`

GetTalosNextVersion returns the TalosNextVersion field if non-nil, zero value otherwise.

### GetTalosNextVersionOk

`func (o *PatchedClusterDetail) GetTalosNextVersionOk() (*string, bool)`

GetTalosNextVersionOk returns a tuple with the TalosNextVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTalosNextVersion

`func (o *PatchedClusterDetail) SetTalosNextVersion(v string)`

SetTalosNextVersion sets TalosNextVersion field to given value.

### HasTalosNextVersion

`func (o *PatchedClusterDetail) HasTalosNextVersion() bool`

HasTalosNextVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


