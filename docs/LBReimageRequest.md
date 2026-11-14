# LBReimageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | **string** | The plan_id returned by GET | 
**Confirmation** | **string** | The cluster&#39;s name, typed to confirm | 

## Methods

### NewLBReimageRequest

`func NewLBReimageRequest(planId string, confirmation string, ) *LBReimageRequest`

NewLBReimageRequest instantiates a new LBReimageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLBReimageRequestWithDefaults

`func NewLBReimageRequestWithDefaults() *LBReimageRequest`

NewLBReimageRequestWithDefaults instantiates a new LBReimageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *LBReimageRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *LBReimageRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *LBReimageRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetConfirmation

`func (o *LBReimageRequest) GetConfirmation() string`

GetConfirmation returns the Confirmation field if non-nil, zero value otherwise.

### GetConfirmationOk

`func (o *LBReimageRequest) GetConfirmationOk() (*string, bool)`

GetConfirmationOk returns a tuple with the Confirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmation

`func (o *LBReimageRequest) SetConfirmation(v string)`

SetConfirmation sets Confirmation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


