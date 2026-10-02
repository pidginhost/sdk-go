# ResourcePoolAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePoolPackage** | **string** | ID or slug | 
**ResourcePoolSize** | **int32** |  | 
**Generation** | Pointer to **string** |  | [optional] 

## Methods

### NewResourcePoolAddRequest

`func NewResourcePoolAddRequest(resourcePoolPackage string, resourcePoolSize int32, ) *ResourcePoolAddRequest`

NewResourcePoolAddRequest instantiates a new ResourcePoolAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePoolAddRequestWithDefaults

`func NewResourcePoolAddRequestWithDefaults() *ResourcePoolAddRequest`

NewResourcePoolAddRequestWithDefaults instantiates a new ResourcePoolAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePoolPackage

`func (o *ResourcePoolAddRequest) GetResourcePoolPackage() string`

GetResourcePoolPackage returns the ResourcePoolPackage field if non-nil, zero value otherwise.

### GetResourcePoolPackageOk

`func (o *ResourcePoolAddRequest) GetResourcePoolPackageOk() (*string, bool)`

GetResourcePoolPackageOk returns a tuple with the ResourcePoolPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolPackage

`func (o *ResourcePoolAddRequest) SetResourcePoolPackage(v string)`

SetResourcePoolPackage sets ResourcePoolPackage field to given value.


### GetResourcePoolSize

`func (o *ResourcePoolAddRequest) GetResourcePoolSize() int32`

GetResourcePoolSize returns the ResourcePoolSize field if non-nil, zero value otherwise.

### GetResourcePoolSizeOk

`func (o *ResourcePoolAddRequest) GetResourcePoolSizeOk() (*int32, bool)`

GetResourcePoolSizeOk returns a tuple with the ResourcePoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolSize

`func (o *ResourcePoolAddRequest) SetResourcePoolSize(v int32)`

SetResourcePoolSize sets ResourcePoolSize field to given value.


### GetGeneration

`func (o *ResourcePoolAddRequest) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *ResourcePoolAddRequest) GetGenerationOk() (*string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *ResourcePoolAddRequest) SetGeneration(v string)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *ResourcePoolAddRequest) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


