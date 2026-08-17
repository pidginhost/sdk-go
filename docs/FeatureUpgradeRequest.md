# FeatureUpgradeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureName** | **string** | Name of the feature to upgrade (e.g. cert-manager) | 
**Retry** | Pointer to **bool** | Retry a failed install | [optional] [default to false]

## Methods

### NewFeatureUpgradeRequest

`func NewFeatureUpgradeRequest(featureName string, ) *FeatureUpgradeRequest`

NewFeatureUpgradeRequest instantiates a new FeatureUpgradeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureUpgradeRequestWithDefaults

`func NewFeatureUpgradeRequestWithDefaults() *FeatureUpgradeRequest`

NewFeatureUpgradeRequestWithDefaults instantiates a new FeatureUpgradeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureName

`func (o *FeatureUpgradeRequest) GetFeatureName() string`

GetFeatureName returns the FeatureName field if non-nil, zero value otherwise.

### GetFeatureNameOk

`func (o *FeatureUpgradeRequest) GetFeatureNameOk() (*string, bool)`

GetFeatureNameOk returns a tuple with the FeatureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureName

`func (o *FeatureUpgradeRequest) SetFeatureName(v string)`

SetFeatureName sets FeatureName field to given value.


### GetRetry

`func (o *FeatureUpgradeRequest) GetRetry() bool`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *FeatureUpgradeRequest) GetRetryOk() (*bool, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *FeatureUpgradeRequest) SetRetry(v bool)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *FeatureUpgradeRequest) HasRetry() bool`

HasRetry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


