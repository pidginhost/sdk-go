# LBUpgradeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | Return the computed plan without performing it | [optional] [default to false]
**InspectionId** | Pointer to **string** | The plan identifier returned by a dry run | [optional] 

## Methods

### NewLBUpgradeRequest

`func NewLBUpgradeRequest() *LBUpgradeRequest`

NewLBUpgradeRequest instantiates a new LBUpgradeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLBUpgradeRequestWithDefaults

`func NewLBUpgradeRequestWithDefaults() *LBUpgradeRequest`

NewLBUpgradeRequestWithDefaults instantiates a new LBUpgradeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *LBUpgradeRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *LBUpgradeRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *LBUpgradeRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *LBUpgradeRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetInspectionId

`func (o *LBUpgradeRequest) GetInspectionId() string`

GetInspectionId returns the InspectionId field if non-nil, zero value otherwise.

### GetInspectionIdOk

`func (o *LBUpgradeRequest) GetInspectionIdOk() (*string, bool)`

GetInspectionIdOk returns a tuple with the InspectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectionId

`func (o *LBUpgradeRequest) SetInspectionId(v string)`

SetInspectionId sets InspectionId field to given value.

### HasInspectionId

`func (o *LBUpgradeRequest) HasInspectionId() bool`

HasInspectionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


