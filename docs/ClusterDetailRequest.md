# ClusterDetailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**PricePerMonth** | **string** |  | 
**Features** | Pointer to [**[]FeaturesEnum**](FeaturesEnum.md) |  | [optional] 
**Protected** | Pointer to **bool** |  | [optional] 

## Methods

### NewClusterDetailRequest

`func NewClusterDetailRequest(pricePerMonth string, ) *ClusterDetailRequest`

NewClusterDetailRequest instantiates a new ClusterDetailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDetailRequestWithDefaults

`func NewClusterDetailRequestWithDefaults() *ClusterDetailRequest`

NewClusterDetailRequestWithDefaults instantiates a new ClusterDetailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterDetailRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterDetailRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterDetailRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterDetailRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPricePerMonth

`func (o *ClusterDetailRequest) GetPricePerMonth() string`

GetPricePerMonth returns the PricePerMonth field if non-nil, zero value otherwise.

### GetPricePerMonthOk

`func (o *ClusterDetailRequest) GetPricePerMonthOk() (*string, bool)`

GetPricePerMonthOk returns a tuple with the PricePerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerMonth

`func (o *ClusterDetailRequest) SetPricePerMonth(v string)`

SetPricePerMonth sets PricePerMonth field to given value.


### GetFeatures

`func (o *ClusterDetailRequest) GetFeatures() []FeaturesEnum`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ClusterDetailRequest) GetFeaturesOk() (*[]FeaturesEnum, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ClusterDetailRequest) SetFeatures(v []FeaturesEnum)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ClusterDetailRequest) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetProtected

`func (o *ClusterDetailRequest) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *ClusterDetailRequest) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *ClusterDetailRequest) SetProtected(v bool)`

SetProtected sets Protected field to given value.

### HasProtected

`func (o *ClusterDetailRequest) HasProtected() bool`

HasProtected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


