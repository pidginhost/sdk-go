# ClusterAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterType** | [**ClusterTypeEnum**](ClusterTypeEnum.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**ResourcePoolPackage** | **string** | ID or slug | 
**ResourcePoolSize** | Pointer to **int32** |  | [optional] 
**KubeVersion** | Pointer to [**KubeVersionEnum**](KubeVersionEnum.md) |  | [optional] [default to KUBEVERSIONENUM__1_35_1]
**Features** | Pointer to [**[]FeaturesEnum**](FeaturesEnum.md) |  | [optional] 
**EnableGatewayApi** | Pointer to **bool** |  | [optional] 

## Methods

### NewClusterAdd

`func NewClusterAdd(clusterType ClusterTypeEnum, resourcePoolPackage string, ) *ClusterAdd`

NewClusterAdd instantiates a new ClusterAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAddWithDefaults

`func NewClusterAddWithDefaults() *ClusterAdd`

NewClusterAddWithDefaults instantiates a new ClusterAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterType

`func (o *ClusterAdd) GetClusterType() ClusterTypeEnum`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *ClusterAdd) GetClusterTypeOk() (*ClusterTypeEnum, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *ClusterAdd) SetClusterType(v ClusterTypeEnum)`

SetClusterType sets ClusterType field to given value.


### GetName

`func (o *ClusterAdd) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterAdd) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterAdd) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterAdd) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourcePoolPackage

`func (o *ClusterAdd) GetResourcePoolPackage() string`

GetResourcePoolPackage returns the ResourcePoolPackage field if non-nil, zero value otherwise.

### GetResourcePoolPackageOk

`func (o *ClusterAdd) GetResourcePoolPackageOk() (*string, bool)`

GetResourcePoolPackageOk returns a tuple with the ResourcePoolPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolPackage

`func (o *ClusterAdd) SetResourcePoolPackage(v string)`

SetResourcePoolPackage sets ResourcePoolPackage field to given value.


### GetResourcePoolSize

`func (o *ClusterAdd) GetResourcePoolSize() int32`

GetResourcePoolSize returns the ResourcePoolSize field if non-nil, zero value otherwise.

### GetResourcePoolSizeOk

`func (o *ClusterAdd) GetResourcePoolSizeOk() (*int32, bool)`

GetResourcePoolSizeOk returns a tuple with the ResourcePoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolSize

`func (o *ClusterAdd) SetResourcePoolSize(v int32)`

SetResourcePoolSize sets ResourcePoolSize field to given value.

### HasResourcePoolSize

`func (o *ClusterAdd) HasResourcePoolSize() bool`

HasResourcePoolSize returns a boolean if a field has been set.

### GetKubeVersion

`func (o *ClusterAdd) GetKubeVersion() KubeVersionEnum`

GetKubeVersion returns the KubeVersion field if non-nil, zero value otherwise.

### GetKubeVersionOk

`func (o *ClusterAdd) GetKubeVersionOk() (*KubeVersionEnum, bool)`

GetKubeVersionOk returns a tuple with the KubeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeVersion

`func (o *ClusterAdd) SetKubeVersion(v KubeVersionEnum)`

SetKubeVersion sets KubeVersion field to given value.

### HasKubeVersion

`func (o *ClusterAdd) HasKubeVersion() bool`

HasKubeVersion returns a boolean if a field has been set.

### GetFeatures

`func (o *ClusterAdd) GetFeatures() []FeaturesEnum`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ClusterAdd) GetFeaturesOk() (*[]FeaturesEnum, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ClusterAdd) SetFeatures(v []FeaturesEnum)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ClusterAdd) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetEnableGatewayApi

`func (o *ClusterAdd) GetEnableGatewayApi() bool`

GetEnableGatewayApi returns the EnableGatewayApi field if non-nil, zero value otherwise.

### GetEnableGatewayApiOk

`func (o *ClusterAdd) GetEnableGatewayApiOk() (*bool, bool)`

GetEnableGatewayApiOk returns a tuple with the EnableGatewayApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGatewayApi

`func (o *ClusterAdd) SetEnableGatewayApi(v bool)`

SetEnableGatewayApi sets EnableGatewayApi field to given value.

### HasEnableGatewayApi

`func (o *ClusterAdd) HasEnableGatewayApi() bool`

HasEnableGatewayApi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


