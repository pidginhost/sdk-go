# HostingUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FetchedAt** | **string** |  | 
**Package** | **string** |  | 
**Metrics** | [**[]HostingUsageMetric**](HostingUsageMetric.md) |  | 
**LiveAvailable** | **bool** |  | 
**LveAvailable** | **bool** |  | 
**LvePeriod** | [**[]HostingLVEPeriod**](HostingLVEPeriod.md) |  | 
**PackageLimits** | **map[string]int32** |  | 
**PressureDetected** | **bool** |  | 

## Methods

### NewHostingUsage

`func NewHostingUsage(fetchedAt string, package_ string, metrics []HostingUsageMetric, liveAvailable bool, lveAvailable bool, lvePeriod []HostingLVEPeriod, packageLimits map[string]int32, pressureDetected bool, ) *HostingUsage`

NewHostingUsage instantiates a new HostingUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostingUsageWithDefaults

`func NewHostingUsageWithDefaults() *HostingUsage`

NewHostingUsageWithDefaults instantiates a new HostingUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFetchedAt

`func (o *HostingUsage) GetFetchedAt() string`

GetFetchedAt returns the FetchedAt field if non-nil, zero value otherwise.

### GetFetchedAtOk

`func (o *HostingUsage) GetFetchedAtOk() (*string, bool)`

GetFetchedAtOk returns a tuple with the FetchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchedAt

`func (o *HostingUsage) SetFetchedAt(v string)`

SetFetchedAt sets FetchedAt field to given value.


### GetPackage

`func (o *HostingUsage) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *HostingUsage) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *HostingUsage) SetPackage(v string)`

SetPackage sets Package field to given value.


### GetMetrics

`func (o *HostingUsage) GetMetrics() []HostingUsageMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *HostingUsage) GetMetricsOk() (*[]HostingUsageMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *HostingUsage) SetMetrics(v []HostingUsageMetric)`

SetMetrics sets Metrics field to given value.


### GetLiveAvailable

`func (o *HostingUsage) GetLiveAvailable() bool`

GetLiveAvailable returns the LiveAvailable field if non-nil, zero value otherwise.

### GetLiveAvailableOk

`func (o *HostingUsage) GetLiveAvailableOk() (*bool, bool)`

GetLiveAvailableOk returns a tuple with the LiveAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveAvailable

`func (o *HostingUsage) SetLiveAvailable(v bool)`

SetLiveAvailable sets LiveAvailable field to given value.


### GetLveAvailable

`func (o *HostingUsage) GetLveAvailable() bool`

GetLveAvailable returns the LveAvailable field if non-nil, zero value otherwise.

### GetLveAvailableOk

`func (o *HostingUsage) GetLveAvailableOk() (*bool, bool)`

GetLveAvailableOk returns a tuple with the LveAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLveAvailable

`func (o *HostingUsage) SetLveAvailable(v bool)`

SetLveAvailable sets LveAvailable field to given value.


### GetLvePeriod

`func (o *HostingUsage) GetLvePeriod() []HostingLVEPeriod`

GetLvePeriod returns the LvePeriod field if non-nil, zero value otherwise.

### GetLvePeriodOk

`func (o *HostingUsage) GetLvePeriodOk() (*[]HostingLVEPeriod, bool)`

GetLvePeriodOk returns a tuple with the LvePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvePeriod

`func (o *HostingUsage) SetLvePeriod(v []HostingLVEPeriod)`

SetLvePeriod sets LvePeriod field to given value.


### GetPackageLimits

`func (o *HostingUsage) GetPackageLimits() map[string]int32`

GetPackageLimits returns the PackageLimits field if non-nil, zero value otherwise.

### GetPackageLimitsOk

`func (o *HostingUsage) GetPackageLimitsOk() (*map[string]int32, bool)`

GetPackageLimitsOk returns a tuple with the PackageLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLimits

`func (o *HostingUsage) SetPackageLimits(v map[string]int32)`

SetPackageLimits sets PackageLimits field to given value.


### GetPressureDetected

`func (o *HostingUsage) GetPressureDetected() bool`

GetPressureDetected returns the PressureDetected field if non-nil, zero value otherwise.

### GetPressureDetectedOk

`func (o *HostingUsage) GetPressureDetectedOk() (*bool, bool)`

GetPressureDetectedOk returns a tuple with the PressureDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPressureDetected

`func (o *HostingUsage) SetPressureDetected(v bool)`

SetPressureDetected sets PressureDetected field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


