# GatewayRouteOutcome

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | queued, running, applied or failed. | 
**CompletedAt** | **NullableString** |  | 
**Reason** | **string** | Why the server failed, or why its applied route is not acknowledged. | 

## Methods

### NewGatewayRouteOutcome

`func NewGatewayRouteOutcome(status string, completedAt NullableString, reason string, ) *GatewayRouteOutcome`

NewGatewayRouteOutcome instantiates a new GatewayRouteOutcome object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayRouteOutcomeWithDefaults

`func NewGatewayRouteOutcomeWithDefaults() *GatewayRouteOutcome`

NewGatewayRouteOutcomeWithDefaults instantiates a new GatewayRouteOutcome object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GatewayRouteOutcome) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GatewayRouteOutcome) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GatewayRouteOutcome) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCompletedAt

`func (o *GatewayRouteOutcome) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GatewayRouteOutcome) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GatewayRouteOutcome) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *GatewayRouteOutcome) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *GatewayRouteOutcome) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetReason

`func (o *GatewayRouteOutcome) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GatewayRouteOutcome) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GatewayRouteOutcome) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


