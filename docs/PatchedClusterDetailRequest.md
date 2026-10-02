# PatchedClusterDetailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**PricePerMonth** | Pointer to **string** |  | [optional] 
**Features** | Pointer to [**[]FeaturesEnum**](FeaturesEnum.md) |  | [optional] 
**Protected** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedClusterDetailRequest

`func NewPatchedClusterDetailRequest() *PatchedClusterDetailRequest`

NewPatchedClusterDetailRequest instantiates a new PatchedClusterDetailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedClusterDetailRequestWithDefaults

`func NewPatchedClusterDetailRequestWithDefaults() *PatchedClusterDetailRequest`

NewPatchedClusterDetailRequestWithDefaults instantiates a new PatchedClusterDetailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedClusterDetailRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedClusterDetailRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedClusterDetailRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedClusterDetailRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPricePerMonth

`func (o *PatchedClusterDetailRequest) GetPricePerMonth() string`

GetPricePerMonth returns the PricePerMonth field if non-nil, zero value otherwise.

### GetPricePerMonthOk

`func (o *PatchedClusterDetailRequest) GetPricePerMonthOk() (*string, bool)`

GetPricePerMonthOk returns a tuple with the PricePerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerMonth

`func (o *PatchedClusterDetailRequest) SetPricePerMonth(v string)`

SetPricePerMonth sets PricePerMonth field to given value.

### HasPricePerMonth

`func (o *PatchedClusterDetailRequest) HasPricePerMonth() bool`

HasPricePerMonth returns a boolean if a field has been set.

### GetFeatures

`func (o *PatchedClusterDetailRequest) GetFeatures() []FeaturesEnum`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *PatchedClusterDetailRequest) GetFeaturesOk() (*[]FeaturesEnum, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *PatchedClusterDetailRequest) SetFeatures(v []FeaturesEnum)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *PatchedClusterDetailRequest) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetProtected

`func (o *PatchedClusterDetailRequest) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *PatchedClusterDetailRequest) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *PatchedClusterDetailRequest) SetProtected(v bool)`

SetProtected sets Protected field to given value.

### HasProtected

`func (o *PatchedClusterDetailRequest) HasProtected() bool`

HasProtected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


