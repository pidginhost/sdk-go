# BootISO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Category** | Pointer to [**CategoryEnum**](CategoryEnum.md) |  | [optional] 
**MinDiskSize** | Pointer to **int32** | GB, 0 &#x3D; no minimum | [optional] 
**MinRam** | Pointer to **float64** | GB, 0 &#x3D; no minimum | [optional] 
**WipesDisk** | Pointer to **bool** | Installer media — show the data-loss warning | [optional] 
**Compatible** | **bool** |  | [readonly] 

## Methods

### NewBootISO

`func NewBootISO(slug string, name string, compatible bool, ) *BootISO`

NewBootISO instantiates a new BootISO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootISOWithDefaults

`func NewBootISOWithDefaults() *BootISO`

NewBootISOWithDefaults instantiates a new BootISO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlug

`func (o *BootISO) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BootISO) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BootISO) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetName

`func (o *BootISO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BootISO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BootISO) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BootISO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BootISO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BootISO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BootISO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *BootISO) GetCategory() CategoryEnum`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *BootISO) GetCategoryOk() (*CategoryEnum, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *BootISO) SetCategory(v CategoryEnum)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *BootISO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetMinDiskSize

`func (o *BootISO) GetMinDiskSize() int32`

GetMinDiskSize returns the MinDiskSize field if non-nil, zero value otherwise.

### GetMinDiskSizeOk

`func (o *BootISO) GetMinDiskSizeOk() (*int32, bool)`

GetMinDiskSizeOk returns a tuple with the MinDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDiskSize

`func (o *BootISO) SetMinDiskSize(v int32)`

SetMinDiskSize sets MinDiskSize field to given value.

### HasMinDiskSize

`func (o *BootISO) HasMinDiskSize() bool`

HasMinDiskSize returns a boolean if a field has been set.

### GetMinRam

`func (o *BootISO) GetMinRam() float64`

GetMinRam returns the MinRam field if non-nil, zero value otherwise.

### GetMinRamOk

`func (o *BootISO) GetMinRamOk() (*float64, bool)`

GetMinRamOk returns a tuple with the MinRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRam

`func (o *BootISO) SetMinRam(v float64)`

SetMinRam sets MinRam field to given value.

### HasMinRam

`func (o *BootISO) HasMinRam() bool`

HasMinRam returns a boolean if a field has been set.

### GetWipesDisk

`func (o *BootISO) GetWipesDisk() bool`

GetWipesDisk returns the WipesDisk field if non-nil, zero value otherwise.

### GetWipesDiskOk

`func (o *BootISO) GetWipesDiskOk() (*bool, bool)`

GetWipesDiskOk returns a tuple with the WipesDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWipesDisk

`func (o *BootISO) SetWipesDisk(v bool)`

SetWipesDisk sets WipesDisk field to given value.

### HasWipesDisk

`func (o *BootISO) HasWipesDisk() bool`

HasWipesDisk returns a boolean if a field has been set.

### GetCompatible

`func (o *BootISO) GetCompatible() bool`

GetCompatible returns the Compatible field if non-nil, zero value otherwise.

### GetCompatibleOk

`func (o *BootISO) GetCompatibleOk() (*bool, bool)`

GetCompatibleOk returns a tuple with the Compatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatible

`func (o *BootISO) SetCompatible(v bool)`

SetCompatible sets Compatible field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


