# ClusterDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Status** | [**StatusA57Enum**](StatusA57Enum.md) |  | [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**ClusterType** | **string** |  | [readonly] 
**KubeVersion** | **string** |  | [readonly] 
**PricePerMonth** | **float64** |  | 
**PricePerHour** | **string** |  | [readonly] 
**Features** | Pointer to [**[]FeaturesEnum**](FeaturesEnum.md) |  | [optional] 
**FeaturesReady** | **bool** |  | [readonly] 
**KubeconfigValidUntil** | **string** |  | [readonly] 
**Ipv4Address** | **string** |  | [readonly] 
**Protected** | Pointer to **bool** |  | [optional] 
**TalosVersion** | **string** |  | [readonly] 
**TalosUpgradeAvailable** | **string** |  | [readonly] 
**TalosNextVersion** | **string** |  | [readonly] 

## Methods

### NewClusterDetail

`func NewClusterDetail(id int32, status StatusA57Enum, clusterType string, kubeVersion string, pricePerMonth float64, pricePerHour string, featuresReady bool, kubeconfigValidUntil string, ipv4Address string, talosVersion string, talosUpgradeAvailable string, talosNextVersion string, ) *ClusterDetail`

NewClusterDetail instantiates a new ClusterDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDetailWithDefaults

`func NewClusterDetailWithDefaults() *ClusterDetail`

NewClusterDetailWithDefaults instantiates a new ClusterDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterDetail) SetId(v int32)`

SetId sets Id field to given value.


### GetStatus

`func (o *ClusterDetail) GetStatus() StatusA57Enum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterDetail) GetStatusOk() (*StatusA57Enum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterDetail) SetStatus(v StatusA57Enum)`

SetStatus sets Status field to given value.


### GetName

`func (o *ClusterDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClusterType

`func (o *ClusterDetail) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *ClusterDetail) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *ClusterDetail) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.


### GetKubeVersion

`func (o *ClusterDetail) GetKubeVersion() string`

GetKubeVersion returns the KubeVersion field if non-nil, zero value otherwise.

### GetKubeVersionOk

`func (o *ClusterDetail) GetKubeVersionOk() (*string, bool)`

GetKubeVersionOk returns a tuple with the KubeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeVersion

`func (o *ClusterDetail) SetKubeVersion(v string)`

SetKubeVersion sets KubeVersion field to given value.


### GetPricePerMonth

`func (o *ClusterDetail) GetPricePerMonth() float64`

GetPricePerMonth returns the PricePerMonth field if non-nil, zero value otherwise.

### GetPricePerMonthOk

`func (o *ClusterDetail) GetPricePerMonthOk() (*float64, bool)`

GetPricePerMonthOk returns a tuple with the PricePerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerMonth

`func (o *ClusterDetail) SetPricePerMonth(v float64)`

SetPricePerMonth sets PricePerMonth field to given value.


### GetPricePerHour

`func (o *ClusterDetail) GetPricePerHour() string`

GetPricePerHour returns the PricePerHour field if non-nil, zero value otherwise.

### GetPricePerHourOk

`func (o *ClusterDetail) GetPricePerHourOk() (*string, bool)`

GetPricePerHourOk returns a tuple with the PricePerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerHour

`func (o *ClusterDetail) SetPricePerHour(v string)`

SetPricePerHour sets PricePerHour field to given value.


### GetFeatures

`func (o *ClusterDetail) GetFeatures() []FeaturesEnum`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ClusterDetail) GetFeaturesOk() (*[]FeaturesEnum, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ClusterDetail) SetFeatures(v []FeaturesEnum)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ClusterDetail) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFeaturesReady

`func (o *ClusterDetail) GetFeaturesReady() bool`

GetFeaturesReady returns the FeaturesReady field if non-nil, zero value otherwise.

### GetFeaturesReadyOk

`func (o *ClusterDetail) GetFeaturesReadyOk() (*bool, bool)`

GetFeaturesReadyOk returns a tuple with the FeaturesReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesReady

`func (o *ClusterDetail) SetFeaturesReady(v bool)`

SetFeaturesReady sets FeaturesReady field to given value.


### GetKubeconfigValidUntil

`func (o *ClusterDetail) GetKubeconfigValidUntil() string`

GetKubeconfigValidUntil returns the KubeconfigValidUntil field if non-nil, zero value otherwise.

### GetKubeconfigValidUntilOk

`func (o *ClusterDetail) GetKubeconfigValidUntilOk() (*string, bool)`

GetKubeconfigValidUntilOk returns a tuple with the KubeconfigValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeconfigValidUntil

`func (o *ClusterDetail) SetKubeconfigValidUntil(v string)`

SetKubeconfigValidUntil sets KubeconfigValidUntil field to given value.


### GetIpv4Address

`func (o *ClusterDetail) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *ClusterDetail) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *ClusterDetail) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.


### GetProtected

`func (o *ClusterDetail) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *ClusterDetail) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *ClusterDetail) SetProtected(v bool)`

SetProtected sets Protected field to given value.

### HasProtected

`func (o *ClusterDetail) HasProtected() bool`

HasProtected returns a boolean if a field has been set.

### GetTalosVersion

`func (o *ClusterDetail) GetTalosVersion() string`

GetTalosVersion returns the TalosVersion field if non-nil, zero value otherwise.

### GetTalosVersionOk

`func (o *ClusterDetail) GetTalosVersionOk() (*string, bool)`

GetTalosVersionOk returns a tuple with the TalosVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTalosVersion

`func (o *ClusterDetail) SetTalosVersion(v string)`

SetTalosVersion sets TalosVersion field to given value.


### GetTalosUpgradeAvailable

`func (o *ClusterDetail) GetTalosUpgradeAvailable() string`

GetTalosUpgradeAvailable returns the TalosUpgradeAvailable field if non-nil, zero value otherwise.

### GetTalosUpgradeAvailableOk

`func (o *ClusterDetail) GetTalosUpgradeAvailableOk() (*string, bool)`

GetTalosUpgradeAvailableOk returns a tuple with the TalosUpgradeAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTalosUpgradeAvailable

`func (o *ClusterDetail) SetTalosUpgradeAvailable(v string)`

SetTalosUpgradeAvailable sets TalosUpgradeAvailable field to given value.


### GetTalosNextVersion

`func (o *ClusterDetail) GetTalosNextVersion() string`

GetTalosNextVersion returns the TalosNextVersion field if non-nil, zero value otherwise.

### GetTalosNextVersionOk

`func (o *ClusterDetail) GetTalosNextVersionOk() (*string, bool)`

GetTalosNextVersionOk returns a tuple with the TalosNextVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTalosNextVersion

`func (o *ClusterDetail) SetTalosNextVersion(v string)`

SetTalosNextVersion sets TalosNextVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


