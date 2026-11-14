# GatewayStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**GatewayStateEnum**](GatewayStateEnum.md) |  | 
**StateReason** | **string** |  | 
**StateChangedAt** | **NullableString** |  | 
**DesiredRevision** | **int32** |  | 
**AppliedRevision** | **NullableInt32** |  | 
**LastStatusAt** | **NullableString** |  | 
**AgentVersion** | **NullableString** |  | 
**EnvoyVersion** | **NullableString** |  | 
**ImageVersion** | **NullableString** |  | 
**AttemptedRevision** | **NullableInt32** |  | 
**LocalReady** | **NullableBool** |  | 
**ClockSynchronized** | **NullableBool** |  | 
**ApplyError** | [**NullableGatewayApplyError**](GatewayApplyError.md) |  | 
**Interfaces** | [**[]GatewayInterface**](GatewayInterface.md) |  | 
**Conntrack** | [**NullableGatewayConntrack**](GatewayConntrack.md) |  | 
**EgressProbe** | [**NullableGatewayEgressProbe**](GatewayEgressProbe.md) |  | 
**SecurityUpdates** | **NullableInt32** |  | 
**RebootRequired** | **NullableBool** |  | 
**UptimeS** | **NullableInt32** |  | 

## Methods

### NewGatewayStatus

`func NewGatewayStatus(state GatewayStateEnum, stateReason string, stateChangedAt NullableString, desiredRevision int32, appliedRevision NullableInt32, lastStatusAt NullableString, agentVersion NullableString, envoyVersion NullableString, imageVersion NullableString, attemptedRevision NullableInt32, localReady NullableBool, clockSynchronized NullableBool, applyError NullableGatewayApplyError, interfaces []GatewayInterface, conntrack NullableGatewayConntrack, egressProbe NullableGatewayEgressProbe, securityUpdates NullableInt32, rebootRequired NullableBool, uptimeS NullableInt32, ) *GatewayStatus`

NewGatewayStatus instantiates a new GatewayStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayStatusWithDefaults

`func NewGatewayStatusWithDefaults() *GatewayStatus`

NewGatewayStatusWithDefaults instantiates a new GatewayStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *GatewayStatus) GetState() GatewayStateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GatewayStatus) GetStateOk() (*GatewayStateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GatewayStatus) SetState(v GatewayStateEnum)`

SetState sets State field to given value.


### GetStateReason

`func (o *GatewayStatus) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *GatewayStatus) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *GatewayStatus) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.


### GetStateChangedAt

`func (o *GatewayStatus) GetStateChangedAt() string`

GetStateChangedAt returns the StateChangedAt field if non-nil, zero value otherwise.

### GetStateChangedAtOk

`func (o *GatewayStatus) GetStateChangedAtOk() (*string, bool)`

GetStateChangedAtOk returns a tuple with the StateChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChangedAt

`func (o *GatewayStatus) SetStateChangedAt(v string)`

SetStateChangedAt sets StateChangedAt field to given value.


### SetStateChangedAtNil

`func (o *GatewayStatus) SetStateChangedAtNil(b bool)`

 SetStateChangedAtNil sets the value for StateChangedAt to be an explicit nil

### UnsetStateChangedAt
`func (o *GatewayStatus) UnsetStateChangedAt()`

UnsetStateChangedAt ensures that no value is present for StateChangedAt, not even an explicit nil
### GetDesiredRevision

`func (o *GatewayStatus) GetDesiredRevision() int32`

GetDesiredRevision returns the DesiredRevision field if non-nil, zero value otherwise.

### GetDesiredRevisionOk

`func (o *GatewayStatus) GetDesiredRevisionOk() (*int32, bool)`

GetDesiredRevisionOk returns a tuple with the DesiredRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredRevision

`func (o *GatewayStatus) SetDesiredRevision(v int32)`

SetDesiredRevision sets DesiredRevision field to given value.


### GetAppliedRevision

`func (o *GatewayStatus) GetAppliedRevision() int32`

GetAppliedRevision returns the AppliedRevision field if non-nil, zero value otherwise.

### GetAppliedRevisionOk

`func (o *GatewayStatus) GetAppliedRevisionOk() (*int32, bool)`

GetAppliedRevisionOk returns a tuple with the AppliedRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedRevision

`func (o *GatewayStatus) SetAppliedRevision(v int32)`

SetAppliedRevision sets AppliedRevision field to given value.


### SetAppliedRevisionNil

`func (o *GatewayStatus) SetAppliedRevisionNil(b bool)`

 SetAppliedRevisionNil sets the value for AppliedRevision to be an explicit nil

### UnsetAppliedRevision
`func (o *GatewayStatus) UnsetAppliedRevision()`

UnsetAppliedRevision ensures that no value is present for AppliedRevision, not even an explicit nil
### GetLastStatusAt

`func (o *GatewayStatus) GetLastStatusAt() string`

GetLastStatusAt returns the LastStatusAt field if non-nil, zero value otherwise.

### GetLastStatusAtOk

`func (o *GatewayStatus) GetLastStatusAtOk() (*string, bool)`

GetLastStatusAtOk returns a tuple with the LastStatusAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusAt

`func (o *GatewayStatus) SetLastStatusAt(v string)`

SetLastStatusAt sets LastStatusAt field to given value.


### SetLastStatusAtNil

`func (o *GatewayStatus) SetLastStatusAtNil(b bool)`

 SetLastStatusAtNil sets the value for LastStatusAt to be an explicit nil

### UnsetLastStatusAt
`func (o *GatewayStatus) UnsetLastStatusAt()`

UnsetLastStatusAt ensures that no value is present for LastStatusAt, not even an explicit nil
### GetAgentVersion

`func (o *GatewayStatus) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *GatewayStatus) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *GatewayStatus) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.


### SetAgentVersionNil

`func (o *GatewayStatus) SetAgentVersionNil(b bool)`

 SetAgentVersionNil sets the value for AgentVersion to be an explicit nil

### UnsetAgentVersion
`func (o *GatewayStatus) UnsetAgentVersion()`

UnsetAgentVersion ensures that no value is present for AgentVersion, not even an explicit nil
### GetEnvoyVersion

`func (o *GatewayStatus) GetEnvoyVersion() string`

GetEnvoyVersion returns the EnvoyVersion field if non-nil, zero value otherwise.

### GetEnvoyVersionOk

`func (o *GatewayStatus) GetEnvoyVersionOk() (*string, bool)`

GetEnvoyVersionOk returns a tuple with the EnvoyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvoyVersion

`func (o *GatewayStatus) SetEnvoyVersion(v string)`

SetEnvoyVersion sets EnvoyVersion field to given value.


### SetEnvoyVersionNil

`func (o *GatewayStatus) SetEnvoyVersionNil(b bool)`

 SetEnvoyVersionNil sets the value for EnvoyVersion to be an explicit nil

### UnsetEnvoyVersion
`func (o *GatewayStatus) UnsetEnvoyVersion()`

UnsetEnvoyVersion ensures that no value is present for EnvoyVersion, not even an explicit nil
### GetImageVersion

`func (o *GatewayStatus) GetImageVersion() string`

GetImageVersion returns the ImageVersion field if non-nil, zero value otherwise.

### GetImageVersionOk

`func (o *GatewayStatus) GetImageVersionOk() (*string, bool)`

GetImageVersionOk returns a tuple with the ImageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersion

`func (o *GatewayStatus) SetImageVersion(v string)`

SetImageVersion sets ImageVersion field to given value.


### SetImageVersionNil

`func (o *GatewayStatus) SetImageVersionNil(b bool)`

 SetImageVersionNil sets the value for ImageVersion to be an explicit nil

### UnsetImageVersion
`func (o *GatewayStatus) UnsetImageVersion()`

UnsetImageVersion ensures that no value is present for ImageVersion, not even an explicit nil
### GetAttemptedRevision

`func (o *GatewayStatus) GetAttemptedRevision() int32`

GetAttemptedRevision returns the AttemptedRevision field if non-nil, zero value otherwise.

### GetAttemptedRevisionOk

`func (o *GatewayStatus) GetAttemptedRevisionOk() (*int32, bool)`

GetAttemptedRevisionOk returns a tuple with the AttemptedRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptedRevision

`func (o *GatewayStatus) SetAttemptedRevision(v int32)`

SetAttemptedRevision sets AttemptedRevision field to given value.


### SetAttemptedRevisionNil

`func (o *GatewayStatus) SetAttemptedRevisionNil(b bool)`

 SetAttemptedRevisionNil sets the value for AttemptedRevision to be an explicit nil

### UnsetAttemptedRevision
`func (o *GatewayStatus) UnsetAttemptedRevision()`

UnsetAttemptedRevision ensures that no value is present for AttemptedRevision, not even an explicit nil
### GetLocalReady

`func (o *GatewayStatus) GetLocalReady() bool`

GetLocalReady returns the LocalReady field if non-nil, zero value otherwise.

### GetLocalReadyOk

`func (o *GatewayStatus) GetLocalReadyOk() (*bool, bool)`

GetLocalReadyOk returns a tuple with the LocalReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalReady

`func (o *GatewayStatus) SetLocalReady(v bool)`

SetLocalReady sets LocalReady field to given value.


### SetLocalReadyNil

`func (o *GatewayStatus) SetLocalReadyNil(b bool)`

 SetLocalReadyNil sets the value for LocalReady to be an explicit nil

### UnsetLocalReady
`func (o *GatewayStatus) UnsetLocalReady()`

UnsetLocalReady ensures that no value is present for LocalReady, not even an explicit nil
### GetClockSynchronized

`func (o *GatewayStatus) GetClockSynchronized() bool`

GetClockSynchronized returns the ClockSynchronized field if non-nil, zero value otherwise.

### GetClockSynchronizedOk

`func (o *GatewayStatus) GetClockSynchronizedOk() (*bool, bool)`

GetClockSynchronizedOk returns a tuple with the ClockSynchronized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockSynchronized

`func (o *GatewayStatus) SetClockSynchronized(v bool)`

SetClockSynchronized sets ClockSynchronized field to given value.


### SetClockSynchronizedNil

`func (o *GatewayStatus) SetClockSynchronizedNil(b bool)`

 SetClockSynchronizedNil sets the value for ClockSynchronized to be an explicit nil

### UnsetClockSynchronized
`func (o *GatewayStatus) UnsetClockSynchronized()`

UnsetClockSynchronized ensures that no value is present for ClockSynchronized, not even an explicit nil
### GetApplyError

`func (o *GatewayStatus) GetApplyError() GatewayApplyError`

GetApplyError returns the ApplyError field if non-nil, zero value otherwise.

### GetApplyErrorOk

`func (o *GatewayStatus) GetApplyErrorOk() (*GatewayApplyError, bool)`

GetApplyErrorOk returns a tuple with the ApplyError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyError

`func (o *GatewayStatus) SetApplyError(v GatewayApplyError)`

SetApplyError sets ApplyError field to given value.


### SetApplyErrorNil

`func (o *GatewayStatus) SetApplyErrorNil(b bool)`

 SetApplyErrorNil sets the value for ApplyError to be an explicit nil

### UnsetApplyError
`func (o *GatewayStatus) UnsetApplyError()`

UnsetApplyError ensures that no value is present for ApplyError, not even an explicit nil
### GetInterfaces

`func (o *GatewayStatus) GetInterfaces() []GatewayInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *GatewayStatus) GetInterfacesOk() (*[]GatewayInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *GatewayStatus) SetInterfaces(v []GatewayInterface)`

SetInterfaces sets Interfaces field to given value.


### SetInterfacesNil

`func (o *GatewayStatus) SetInterfacesNil(b bool)`

 SetInterfacesNil sets the value for Interfaces to be an explicit nil

### UnsetInterfaces
`func (o *GatewayStatus) UnsetInterfaces()`

UnsetInterfaces ensures that no value is present for Interfaces, not even an explicit nil
### GetConntrack

`func (o *GatewayStatus) GetConntrack() GatewayConntrack`

GetConntrack returns the Conntrack field if non-nil, zero value otherwise.

### GetConntrackOk

`func (o *GatewayStatus) GetConntrackOk() (*GatewayConntrack, bool)`

GetConntrackOk returns a tuple with the Conntrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConntrack

`func (o *GatewayStatus) SetConntrack(v GatewayConntrack)`

SetConntrack sets Conntrack field to given value.


### SetConntrackNil

`func (o *GatewayStatus) SetConntrackNil(b bool)`

 SetConntrackNil sets the value for Conntrack to be an explicit nil

### UnsetConntrack
`func (o *GatewayStatus) UnsetConntrack()`

UnsetConntrack ensures that no value is present for Conntrack, not even an explicit nil
### GetEgressProbe

`func (o *GatewayStatus) GetEgressProbe() GatewayEgressProbe`

GetEgressProbe returns the EgressProbe field if non-nil, zero value otherwise.

### GetEgressProbeOk

`func (o *GatewayStatus) GetEgressProbeOk() (*GatewayEgressProbe, bool)`

GetEgressProbeOk returns a tuple with the EgressProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressProbe

`func (o *GatewayStatus) SetEgressProbe(v GatewayEgressProbe)`

SetEgressProbe sets EgressProbe field to given value.


### SetEgressProbeNil

`func (o *GatewayStatus) SetEgressProbeNil(b bool)`

 SetEgressProbeNil sets the value for EgressProbe to be an explicit nil

### UnsetEgressProbe
`func (o *GatewayStatus) UnsetEgressProbe()`

UnsetEgressProbe ensures that no value is present for EgressProbe, not even an explicit nil
### GetSecurityUpdates

`func (o *GatewayStatus) GetSecurityUpdates() int32`

GetSecurityUpdates returns the SecurityUpdates field if non-nil, zero value otherwise.

### GetSecurityUpdatesOk

`func (o *GatewayStatus) GetSecurityUpdatesOk() (*int32, bool)`

GetSecurityUpdatesOk returns a tuple with the SecurityUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityUpdates

`func (o *GatewayStatus) SetSecurityUpdates(v int32)`

SetSecurityUpdates sets SecurityUpdates field to given value.


### SetSecurityUpdatesNil

`func (o *GatewayStatus) SetSecurityUpdatesNil(b bool)`

 SetSecurityUpdatesNil sets the value for SecurityUpdates to be an explicit nil

### UnsetSecurityUpdates
`func (o *GatewayStatus) UnsetSecurityUpdates()`

UnsetSecurityUpdates ensures that no value is present for SecurityUpdates, not even an explicit nil
### GetRebootRequired

`func (o *GatewayStatus) GetRebootRequired() bool`

GetRebootRequired returns the RebootRequired field if non-nil, zero value otherwise.

### GetRebootRequiredOk

`func (o *GatewayStatus) GetRebootRequiredOk() (*bool, bool)`

GetRebootRequiredOk returns a tuple with the RebootRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootRequired

`func (o *GatewayStatus) SetRebootRequired(v bool)`

SetRebootRequired sets RebootRequired field to given value.


### SetRebootRequiredNil

`func (o *GatewayStatus) SetRebootRequiredNil(b bool)`

 SetRebootRequiredNil sets the value for RebootRequired to be an explicit nil

### UnsetRebootRequired
`func (o *GatewayStatus) UnsetRebootRequired()`

UnsetRebootRequired ensures that no value is present for RebootRequired, not even an explicit nil
### GetUptimeS

`func (o *GatewayStatus) GetUptimeS() int32`

GetUptimeS returns the UptimeS field if non-nil, zero value otherwise.

### GetUptimeSOk

`func (o *GatewayStatus) GetUptimeSOk() (*int32, bool)`

GetUptimeSOk returns a tuple with the UptimeS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeS

`func (o *GatewayStatus) SetUptimeS(v int32)`

SetUptimeS sets UptimeS field to given value.


### SetUptimeSNil

`func (o *GatewayStatus) SetUptimeSNil(b bool)`

 SetUptimeSNil sets the value for UptimeS to be an explicit nil

### UnsetUptimeS
`func (o *GatewayStatus) UnsetUptimeS()`

UnsetUptimeS ensures that no value is present for UptimeS, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


