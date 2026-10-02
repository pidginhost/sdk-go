# LBUpgradePlanResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InspectionId** | **string** |  | 
**Convergence** | **string** |  | 
**Level** | **NullableInt32** |  | 
**Actions** | **[]string** |  | 
**Blockers** | **[]string** |  | 
**CanExecute** | **bool** |  | 

## Methods

### NewLBUpgradePlanResponse

`func NewLBUpgradePlanResponse(inspectionId string, convergence string, level NullableInt32, actions []string, blockers []string, canExecute bool, ) *LBUpgradePlanResponse`

NewLBUpgradePlanResponse instantiates a new LBUpgradePlanResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLBUpgradePlanResponseWithDefaults

`func NewLBUpgradePlanResponseWithDefaults() *LBUpgradePlanResponse`

NewLBUpgradePlanResponseWithDefaults instantiates a new LBUpgradePlanResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInspectionId

`func (o *LBUpgradePlanResponse) GetInspectionId() string`

GetInspectionId returns the InspectionId field if non-nil, zero value otherwise.

### GetInspectionIdOk

`func (o *LBUpgradePlanResponse) GetInspectionIdOk() (*string, bool)`

GetInspectionIdOk returns a tuple with the InspectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectionId

`func (o *LBUpgradePlanResponse) SetInspectionId(v string)`

SetInspectionId sets InspectionId field to given value.


### GetConvergence

`func (o *LBUpgradePlanResponse) GetConvergence() string`

GetConvergence returns the Convergence field if non-nil, zero value otherwise.

### GetConvergenceOk

`func (o *LBUpgradePlanResponse) GetConvergenceOk() (*string, bool)`

GetConvergenceOk returns a tuple with the Convergence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvergence

`func (o *LBUpgradePlanResponse) SetConvergence(v string)`

SetConvergence sets Convergence field to given value.


### GetLevel

`func (o *LBUpgradePlanResponse) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *LBUpgradePlanResponse) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *LBUpgradePlanResponse) SetLevel(v int32)`

SetLevel sets Level field to given value.


### SetLevelNil

`func (o *LBUpgradePlanResponse) SetLevelNil(b bool)`

 SetLevelNil sets the value for Level to be an explicit nil

### UnsetLevel
`func (o *LBUpgradePlanResponse) UnsetLevel()`

UnsetLevel ensures that no value is present for Level, not even an explicit nil
### GetActions

`func (o *LBUpgradePlanResponse) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *LBUpgradePlanResponse) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *LBUpgradePlanResponse) SetActions(v []string)`

SetActions sets Actions field to given value.


### GetBlockers

`func (o *LBUpgradePlanResponse) GetBlockers() []string`

GetBlockers returns the Blockers field if non-nil, zero value otherwise.

### GetBlockersOk

`func (o *LBUpgradePlanResponse) GetBlockersOk() (*[]string, bool)`

GetBlockersOk returns a tuple with the Blockers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockers

`func (o *LBUpgradePlanResponse) SetBlockers(v []string)`

SetBlockers sets Blockers field to given value.


### GetCanExecute

`func (o *LBUpgradePlanResponse) GetCanExecute() bool`

GetCanExecute returns the CanExecute field if non-nil, zero value otherwise.

### GetCanExecuteOk

`func (o *LBUpgradePlanResponse) GetCanExecuteOk() (*bool, bool)`

GetCanExecuteOk returns a tuple with the CanExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanExecute

`func (o *LBUpgradePlanResponse) SetCanExecute(v bool)`

SetCanExecute sets CanExecute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


