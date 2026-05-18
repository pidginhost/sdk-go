# StorageProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Slug** | **string** |  | 
**Name** | **string** |  | [readonly] 
**Type** | **string** |  | [readonly] 
**Unit** | **string** |  | [readonly] 
**Price** | **float64** | price per quantity units per month (if applicable) | 
**MinSize** | **string** |  | [readonly] 
**MaxSize** | **string** |  | [readonly] 

## Methods

### NewStorageProduct

`func NewStorageProduct(id int32, slug string, name string, type_ string, unit string, price float64, minSize string, maxSize string, ) *StorageProduct`

NewStorageProduct instantiates a new StorageProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageProductWithDefaults

`func NewStorageProductWithDefaults() *StorageProduct`

NewStorageProductWithDefaults instantiates a new StorageProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageProduct) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageProduct) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageProduct) SetId(v int32)`

SetId sets Id field to given value.


### GetSlug

`func (o *StorageProduct) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *StorageProduct) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *StorageProduct) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetName

`func (o *StorageProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageProduct) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *StorageProduct) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageProduct) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageProduct) SetType(v string)`

SetType sets Type field to given value.


### GetUnit

`func (o *StorageProduct) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *StorageProduct) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *StorageProduct) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetPrice

`func (o *StorageProduct) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *StorageProduct) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *StorageProduct) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetMinSize

`func (o *StorageProduct) GetMinSize() string`

GetMinSize returns the MinSize field if non-nil, zero value otherwise.

### GetMinSizeOk

`func (o *StorageProduct) GetMinSizeOk() (*string, bool)`

GetMinSizeOk returns a tuple with the MinSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSize

`func (o *StorageProduct) SetMinSize(v string)`

SetMinSize sets MinSize field to given value.


### GetMaxSize

`func (o *StorageProduct) GetMaxSize() string`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *StorageProduct) GetMaxSizeOk() (*string, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *StorageProduct) SetMaxSize(v string)`

SetMaxSize sets MaxSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


