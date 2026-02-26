# ResourcePoolAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePoolPackage** | **string** | ID or slug | 
**ResourcePoolSize** | **int32** |  | 

## Methods

### NewResourcePoolAdd

`func NewResourcePoolAdd(resourcePoolPackage string, resourcePoolSize int32, ) *ResourcePoolAdd`

NewResourcePoolAdd instantiates a new ResourcePoolAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePoolAddWithDefaults

`func NewResourcePoolAddWithDefaults() *ResourcePoolAdd`

NewResourcePoolAddWithDefaults instantiates a new ResourcePoolAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePoolPackage

`func (o *ResourcePoolAdd) GetResourcePoolPackage() string`

GetResourcePoolPackage returns the ResourcePoolPackage field if non-nil, zero value otherwise.

### GetResourcePoolPackageOk

`func (o *ResourcePoolAdd) GetResourcePoolPackageOk() (*string, bool)`

GetResourcePoolPackageOk returns a tuple with the ResourcePoolPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolPackage

`func (o *ResourcePoolAdd) SetResourcePoolPackage(v string)`

SetResourcePoolPackage sets ResourcePoolPackage field to given value.


### GetResourcePoolSize

`func (o *ResourcePoolAdd) GetResourcePoolSize() int32`

GetResourcePoolSize returns the ResourcePoolSize field if non-nil, zero value otherwise.

### GetResourcePoolSizeOk

`func (o *ResourcePoolAdd) GetResourcePoolSizeOk() (*int32, bool)`

GetResourcePoolSizeOk returns a tuple with the ResourcePoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolSize

`func (o *ResourcePoolAdd) SetResourcePoolSize(v int32)`

SetResourcePoolSize sets ResourcePoolSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


