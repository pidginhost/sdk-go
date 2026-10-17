# HostingUsageMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Label** | **string** |  | 
**Used** | **NullableFloat64** |  | 
**Limit** | **NullableFloat64** |  | 
**Unit** | **string** |  | 
**Percentage** | **NullableFloat64** |  | 
**Faults** | Pointer to **NullableFloat64** |  | [optional] 
**Source** | **string** |  | 
**NearLimit** | **bool** |  | 

## Methods

### NewHostingUsageMetric

`func NewHostingUsageMetric(key string, label string, used NullableFloat64, limit NullableFloat64, unit string, percentage NullableFloat64, source string, nearLimit bool, ) *HostingUsageMetric`

NewHostingUsageMetric instantiates a new HostingUsageMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostingUsageMetricWithDefaults

`func NewHostingUsageMetricWithDefaults() *HostingUsageMetric`

NewHostingUsageMetricWithDefaults instantiates a new HostingUsageMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *HostingUsageMetric) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *HostingUsageMetric) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *HostingUsageMetric) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *HostingUsageMetric) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *HostingUsageMetric) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *HostingUsageMetric) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetUsed

`func (o *HostingUsageMetric) GetUsed() float64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *HostingUsageMetric) GetUsedOk() (*float64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *HostingUsageMetric) SetUsed(v float64)`

SetUsed sets Used field to given value.


### SetUsedNil

`func (o *HostingUsageMetric) SetUsedNil(b bool)`

 SetUsedNil sets the value for Used to be an explicit nil

### UnsetUsed
`func (o *HostingUsageMetric) UnsetUsed()`

UnsetUsed ensures that no value is present for Used, not even an explicit nil
### GetLimit

`func (o *HostingUsageMetric) GetLimit() float64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *HostingUsageMetric) GetLimitOk() (*float64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *HostingUsageMetric) SetLimit(v float64)`

SetLimit sets Limit field to given value.


### SetLimitNil

`func (o *HostingUsageMetric) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *HostingUsageMetric) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetUnit

`func (o *HostingUsageMetric) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *HostingUsageMetric) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *HostingUsageMetric) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetPercentage

`func (o *HostingUsageMetric) GetPercentage() float64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *HostingUsageMetric) GetPercentageOk() (*float64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *HostingUsageMetric) SetPercentage(v float64)`

SetPercentage sets Percentage field to given value.


### SetPercentageNil

`func (o *HostingUsageMetric) SetPercentageNil(b bool)`

 SetPercentageNil sets the value for Percentage to be an explicit nil

### UnsetPercentage
`func (o *HostingUsageMetric) UnsetPercentage()`

UnsetPercentage ensures that no value is present for Percentage, not even an explicit nil
### GetFaults

`func (o *HostingUsageMetric) GetFaults() float64`

GetFaults returns the Faults field if non-nil, zero value otherwise.

### GetFaultsOk

`func (o *HostingUsageMetric) GetFaultsOk() (*float64, bool)`

GetFaultsOk returns a tuple with the Faults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaults

`func (o *HostingUsageMetric) SetFaults(v float64)`

SetFaults sets Faults field to given value.

### HasFaults

`func (o *HostingUsageMetric) HasFaults() bool`

HasFaults returns a boolean if a field has been set.

### SetFaultsNil

`func (o *HostingUsageMetric) SetFaultsNil(b bool)`

 SetFaultsNil sets the value for Faults to be an explicit nil

### UnsetFaults
`func (o *HostingUsageMetric) UnsetFaults()`

UnsetFaults ensures that no value is present for Faults, not even an explicit nil
### GetSource

`func (o *HostingUsageMetric) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *HostingUsageMetric) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *HostingUsageMetric) SetSource(v string)`

SetSource sets Source field to given value.


### GetNearLimit

`func (o *HostingUsageMetric) GetNearLimit() bool`

GetNearLimit returns the NearLimit field if non-nil, zero value otherwise.

### GetNearLimitOk

`func (o *HostingUsageMetric) GetNearLimitOk() (*bool, bool)`

GetNearLimitOk returns a tuple with the NearLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNearLimit

`func (o *HostingUsageMetric) SetNearLimit(v bool)`

SetNearLimit sets NearLimit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


