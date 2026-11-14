# Gateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | [readonly] 
**State** | [**GatewayStateEnum**](GatewayStateEnum.md) |  | [readonly] 
**StateReason** | **string** |  | [readonly] 
**Status** | [**ResourceStatusEnum**](ResourceStatusEnum.md) |  | [readonly] 
**DesiredRevision** | **int32** |  | [readonly] 
**AppliedRevision** | **NullableInt32** |  | [readonly] 
**Ipv4** | **NullableString** |  | [readonly] 
**Ipv6** | **NullableString** |  | [readonly] 
**MaintenanceWeekday** | **int32** |  | [readonly] 
**MaintenanceHour** | **int32** |  | [readonly] 
**HaEnabled** | **bool** |  | [readonly] 
**CurrentOperation** | **NullableString** |  | [readonly] 
**Created** | **string** |  | [readonly] 

## Methods

### NewGateway

`func NewGateway(id int32, name string, state GatewayStateEnum, stateReason string, status ResourceStatusEnum, desiredRevision int32, appliedRevision NullableInt32, ipv4 NullableString, ipv6 NullableString, maintenanceWeekday int32, maintenanceHour int32, haEnabled bool, currentOperation NullableString, created string, ) *Gateway`

NewGateway instantiates a new Gateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayWithDefaults

`func NewGatewayWithDefaults() *Gateway`

NewGatewayWithDefaults instantiates a new Gateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Gateway) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Gateway) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Gateway) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Gateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Gateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Gateway) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *Gateway) GetState() GatewayStateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Gateway) GetStateOk() (*GatewayStateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Gateway) SetState(v GatewayStateEnum)`

SetState sets State field to given value.


### GetStateReason

`func (o *Gateway) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *Gateway) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *Gateway) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.


### GetStatus

`func (o *Gateway) GetStatus() ResourceStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Gateway) GetStatusOk() (*ResourceStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Gateway) SetStatus(v ResourceStatusEnum)`

SetStatus sets Status field to given value.


### GetDesiredRevision

`func (o *Gateway) GetDesiredRevision() int32`

GetDesiredRevision returns the DesiredRevision field if non-nil, zero value otherwise.

### GetDesiredRevisionOk

`func (o *Gateway) GetDesiredRevisionOk() (*int32, bool)`

GetDesiredRevisionOk returns a tuple with the DesiredRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredRevision

`func (o *Gateway) SetDesiredRevision(v int32)`

SetDesiredRevision sets DesiredRevision field to given value.


### GetAppliedRevision

`func (o *Gateway) GetAppliedRevision() int32`

GetAppliedRevision returns the AppliedRevision field if non-nil, zero value otherwise.

### GetAppliedRevisionOk

`func (o *Gateway) GetAppliedRevisionOk() (*int32, bool)`

GetAppliedRevisionOk returns a tuple with the AppliedRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedRevision

`func (o *Gateway) SetAppliedRevision(v int32)`

SetAppliedRevision sets AppliedRevision field to given value.


### SetAppliedRevisionNil

`func (o *Gateway) SetAppliedRevisionNil(b bool)`

 SetAppliedRevisionNil sets the value for AppliedRevision to be an explicit nil

### UnsetAppliedRevision
`func (o *Gateway) UnsetAppliedRevision()`

UnsetAppliedRevision ensures that no value is present for AppliedRevision, not even an explicit nil
### GetIpv4

`func (o *Gateway) GetIpv4() string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *Gateway) GetIpv4Ok() (*string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *Gateway) SetIpv4(v string)`

SetIpv4 sets Ipv4 field to given value.


### SetIpv4Nil

`func (o *Gateway) SetIpv4Nil(b bool)`

 SetIpv4Nil sets the value for Ipv4 to be an explicit nil

### UnsetIpv4
`func (o *Gateway) UnsetIpv4()`

UnsetIpv4 ensures that no value is present for Ipv4, not even an explicit nil
### GetIpv6

`func (o *Gateway) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *Gateway) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *Gateway) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.


### SetIpv6Nil

`func (o *Gateway) SetIpv6Nil(b bool)`

 SetIpv6Nil sets the value for Ipv6 to be an explicit nil

### UnsetIpv6
`func (o *Gateway) UnsetIpv6()`

UnsetIpv6 ensures that no value is present for Ipv6, not even an explicit nil
### GetMaintenanceWeekday

`func (o *Gateway) GetMaintenanceWeekday() int32`

GetMaintenanceWeekday returns the MaintenanceWeekday field if non-nil, zero value otherwise.

### GetMaintenanceWeekdayOk

`func (o *Gateway) GetMaintenanceWeekdayOk() (*int32, bool)`

GetMaintenanceWeekdayOk returns a tuple with the MaintenanceWeekday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWeekday

`func (o *Gateway) SetMaintenanceWeekday(v int32)`

SetMaintenanceWeekday sets MaintenanceWeekday field to given value.


### GetMaintenanceHour

`func (o *Gateway) GetMaintenanceHour() int32`

GetMaintenanceHour returns the MaintenanceHour field if non-nil, zero value otherwise.

### GetMaintenanceHourOk

`func (o *Gateway) GetMaintenanceHourOk() (*int32, bool)`

GetMaintenanceHourOk returns a tuple with the MaintenanceHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceHour

`func (o *Gateway) SetMaintenanceHour(v int32)`

SetMaintenanceHour sets MaintenanceHour field to given value.


### GetHaEnabled

`func (o *Gateway) GetHaEnabled() bool`

GetHaEnabled returns the HaEnabled field if non-nil, zero value otherwise.

### GetHaEnabledOk

`func (o *Gateway) GetHaEnabledOk() (*bool, bool)`

GetHaEnabledOk returns a tuple with the HaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaEnabled

`func (o *Gateway) SetHaEnabled(v bool)`

SetHaEnabled sets HaEnabled field to given value.


### GetCurrentOperation

`func (o *Gateway) GetCurrentOperation() string`

GetCurrentOperation returns the CurrentOperation field if non-nil, zero value otherwise.

### GetCurrentOperationOk

`func (o *Gateway) GetCurrentOperationOk() (*string, bool)`

GetCurrentOperationOk returns a tuple with the CurrentOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentOperation

`func (o *Gateway) SetCurrentOperation(v string)`

SetCurrentOperation sets CurrentOperation field to given value.


### SetCurrentOperationNil

`func (o *Gateway) SetCurrentOperationNil(b bool)`

 SetCurrentOperationNil sets the value for CurrentOperation to be an explicit nil

### UnsetCurrentOperation
`func (o *Gateway) UnsetCurrentOperation()`

UnsetCurrentOperation ensures that no value is present for CurrentOperation, not even an explicit nil
### GetCreated

`func (o *Gateway) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Gateway) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Gateway) SetCreated(v string)`

SetCreated sets Created field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


