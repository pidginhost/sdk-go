# LBReimageStart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eligible** | **bool** |  | 
**Reason** | [**NullableLBReimageAvailabilityEnum**](LBReimageAvailabilityEnum.md) | Why no start is possible now: busy (try later) or unavailable  * &#x60;busy&#x60; - busy * &#x60;unavailable&#x60; - unavailable | 
**PlanId** | **string** |  | 

## Methods

### NewLBReimageStart

`func NewLBReimageStart(eligible bool, reason NullableLBReimageAvailabilityEnum, planId string, ) *LBReimageStart`

NewLBReimageStart instantiates a new LBReimageStart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLBReimageStartWithDefaults

`func NewLBReimageStartWithDefaults() *LBReimageStart`

NewLBReimageStartWithDefaults instantiates a new LBReimageStart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEligible

`func (o *LBReimageStart) GetEligible() bool`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *LBReimageStart) GetEligibleOk() (*bool, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *LBReimageStart) SetEligible(v bool)`

SetEligible sets Eligible field to given value.


### GetReason

`func (o *LBReimageStart) GetReason() LBReimageAvailabilityEnum`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *LBReimageStart) GetReasonOk() (*LBReimageAvailabilityEnum, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *LBReimageStart) SetReason(v LBReimageAvailabilityEnum)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *LBReimageStart) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *LBReimageStart) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetPlanId

`func (o *LBReimageStart) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *LBReimageStart) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *LBReimageStart) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


