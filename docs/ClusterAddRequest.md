# ClusterAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterType** | [**ClusterTypeEnum**](ClusterTypeEnum.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**ResourcePoolPackage** | **string** | ID or slug | 
**ResourcePoolSize** | Pointer to **int32** |  | [optional] 
**KubeVersion** | Pointer to [**KubeVersionEnum**](KubeVersionEnum.md) |  | [optional] [default to KUBEVERSIONENUM__1_36_3]
**Features** | Pointer to [**[]FeaturesEnum**](FeaturesEnum.md) |  | [optional] 
**EnableGatewayApi** | Pointer to **bool** |  | [optional] 
**DualStack** | Pointer to **bool** | Enable IPv6 dual-stack for pods, services, and the cluster private network. Available only when the platform has K8S_DUAL_STACK_ENABLED. Cannot be changed after provisioning. | [optional] [default to false]
**Generation** | Pointer to **string** |  | [optional] 

## Methods

### NewClusterAddRequest

`func NewClusterAddRequest(clusterType ClusterTypeEnum, resourcePoolPackage string, ) *ClusterAddRequest`

NewClusterAddRequest instantiates a new ClusterAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAddRequestWithDefaults

`func NewClusterAddRequestWithDefaults() *ClusterAddRequest`

NewClusterAddRequestWithDefaults instantiates a new ClusterAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterType

`func (o *ClusterAddRequest) GetClusterType() ClusterTypeEnum`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *ClusterAddRequest) GetClusterTypeOk() (*ClusterTypeEnum, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *ClusterAddRequest) SetClusterType(v ClusterTypeEnum)`

SetClusterType sets ClusterType field to given value.


### GetName

`func (o *ClusterAddRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterAddRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterAddRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterAddRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourcePoolPackage

`func (o *ClusterAddRequest) GetResourcePoolPackage() string`

GetResourcePoolPackage returns the ResourcePoolPackage field if non-nil, zero value otherwise.

### GetResourcePoolPackageOk

`func (o *ClusterAddRequest) GetResourcePoolPackageOk() (*string, bool)`

GetResourcePoolPackageOk returns a tuple with the ResourcePoolPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolPackage

`func (o *ClusterAddRequest) SetResourcePoolPackage(v string)`

SetResourcePoolPackage sets ResourcePoolPackage field to given value.


### GetResourcePoolSize

`func (o *ClusterAddRequest) GetResourcePoolSize() int32`

GetResourcePoolSize returns the ResourcePoolSize field if non-nil, zero value otherwise.

### GetResourcePoolSizeOk

`func (o *ClusterAddRequest) GetResourcePoolSizeOk() (*int32, bool)`

GetResourcePoolSizeOk returns a tuple with the ResourcePoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolSize

`func (o *ClusterAddRequest) SetResourcePoolSize(v int32)`

SetResourcePoolSize sets ResourcePoolSize field to given value.

### HasResourcePoolSize

`func (o *ClusterAddRequest) HasResourcePoolSize() bool`

HasResourcePoolSize returns a boolean if a field has been set.

### GetKubeVersion

`func (o *ClusterAddRequest) GetKubeVersion() KubeVersionEnum`

GetKubeVersion returns the KubeVersion field if non-nil, zero value otherwise.

### GetKubeVersionOk

`func (o *ClusterAddRequest) GetKubeVersionOk() (*KubeVersionEnum, bool)`

GetKubeVersionOk returns a tuple with the KubeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeVersion

`func (o *ClusterAddRequest) SetKubeVersion(v KubeVersionEnum)`

SetKubeVersion sets KubeVersion field to given value.

### HasKubeVersion

`func (o *ClusterAddRequest) HasKubeVersion() bool`

HasKubeVersion returns a boolean if a field has been set.

### GetFeatures

`func (o *ClusterAddRequest) GetFeatures() []FeaturesEnum`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ClusterAddRequest) GetFeaturesOk() (*[]FeaturesEnum, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ClusterAddRequest) SetFeatures(v []FeaturesEnum)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ClusterAddRequest) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetEnableGatewayApi

`func (o *ClusterAddRequest) GetEnableGatewayApi() bool`

GetEnableGatewayApi returns the EnableGatewayApi field if non-nil, zero value otherwise.

### GetEnableGatewayApiOk

`func (o *ClusterAddRequest) GetEnableGatewayApiOk() (*bool, bool)`

GetEnableGatewayApiOk returns a tuple with the EnableGatewayApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGatewayApi

`func (o *ClusterAddRequest) SetEnableGatewayApi(v bool)`

SetEnableGatewayApi sets EnableGatewayApi field to given value.

### HasEnableGatewayApi

`func (o *ClusterAddRequest) HasEnableGatewayApi() bool`

HasEnableGatewayApi returns a boolean if a field has been set.

### GetDualStack

`func (o *ClusterAddRequest) GetDualStack() bool`

GetDualStack returns the DualStack field if non-nil, zero value otherwise.

### GetDualStackOk

`func (o *ClusterAddRequest) GetDualStackOk() (*bool, bool)`

GetDualStackOk returns a tuple with the DualStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDualStack

`func (o *ClusterAddRequest) SetDualStack(v bool)`

SetDualStack sets DualStack field to given value.

### HasDualStack

`func (o *ClusterAddRequest) HasDualStack() bool`

HasDualStack returns a boolean if a field has been set.

### GetGeneration

`func (o *ClusterAddRequest) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *ClusterAddRequest) GetGenerationOk() (*string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *ClusterAddRequest) SetGeneration(v string)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *ClusterAddRequest) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


