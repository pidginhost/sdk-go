# GatewayHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**GatewayHealthCheckTypeEnum**](GatewayHealthCheckTypeEnum.md) |  | 
**Path** | **NullableString** |  | 
**StatusMin** | **NullableInt32** |  | 
**StatusMax** | **NullableInt32** |  | 
**IntervalS** | **NullableInt32** |  | 
**TimeoutS** | **NullableInt32** |  | 
**HealthyThreshold** | **NullableInt32** |  | 
**UnhealthyThreshold** | **NullableInt32** |  | 
**Port** | **NullableInt32** | Null checks each server&#39;s own port. | 

## Methods

### NewGatewayHealthCheck

`func NewGatewayHealthCheck(type_ GatewayHealthCheckTypeEnum, path NullableString, statusMin NullableInt32, statusMax NullableInt32, intervalS NullableInt32, timeoutS NullableInt32, healthyThreshold NullableInt32, unhealthyThreshold NullableInt32, port NullableInt32, ) *GatewayHealthCheck`

NewGatewayHealthCheck instantiates a new GatewayHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayHealthCheckWithDefaults

`func NewGatewayHealthCheckWithDefaults() *GatewayHealthCheck`

NewGatewayHealthCheckWithDefaults instantiates a new GatewayHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GatewayHealthCheck) GetType() GatewayHealthCheckTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayHealthCheck) GetTypeOk() (*GatewayHealthCheckTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayHealthCheck) SetType(v GatewayHealthCheckTypeEnum)`

SetType sets Type field to given value.


### GetPath

`func (o *GatewayHealthCheck) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GatewayHealthCheck) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GatewayHealthCheck) SetPath(v string)`

SetPath sets Path field to given value.


### SetPathNil

`func (o *GatewayHealthCheck) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *GatewayHealthCheck) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetStatusMin

`func (o *GatewayHealthCheck) GetStatusMin() int32`

GetStatusMin returns the StatusMin field if non-nil, zero value otherwise.

### GetStatusMinOk

`func (o *GatewayHealthCheck) GetStatusMinOk() (*int32, bool)`

GetStatusMinOk returns a tuple with the StatusMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMin

`func (o *GatewayHealthCheck) SetStatusMin(v int32)`

SetStatusMin sets StatusMin field to given value.


### SetStatusMinNil

`func (o *GatewayHealthCheck) SetStatusMinNil(b bool)`

 SetStatusMinNil sets the value for StatusMin to be an explicit nil

### UnsetStatusMin
`func (o *GatewayHealthCheck) UnsetStatusMin()`

UnsetStatusMin ensures that no value is present for StatusMin, not even an explicit nil
### GetStatusMax

`func (o *GatewayHealthCheck) GetStatusMax() int32`

GetStatusMax returns the StatusMax field if non-nil, zero value otherwise.

### GetStatusMaxOk

`func (o *GatewayHealthCheck) GetStatusMaxOk() (*int32, bool)`

GetStatusMaxOk returns a tuple with the StatusMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMax

`func (o *GatewayHealthCheck) SetStatusMax(v int32)`

SetStatusMax sets StatusMax field to given value.


### SetStatusMaxNil

`func (o *GatewayHealthCheck) SetStatusMaxNil(b bool)`

 SetStatusMaxNil sets the value for StatusMax to be an explicit nil

### UnsetStatusMax
`func (o *GatewayHealthCheck) UnsetStatusMax()`

UnsetStatusMax ensures that no value is present for StatusMax, not even an explicit nil
### GetIntervalS

`func (o *GatewayHealthCheck) GetIntervalS() int32`

GetIntervalS returns the IntervalS field if non-nil, zero value otherwise.

### GetIntervalSOk

`func (o *GatewayHealthCheck) GetIntervalSOk() (*int32, bool)`

GetIntervalSOk returns a tuple with the IntervalS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalS

`func (o *GatewayHealthCheck) SetIntervalS(v int32)`

SetIntervalS sets IntervalS field to given value.


### SetIntervalSNil

`func (o *GatewayHealthCheck) SetIntervalSNil(b bool)`

 SetIntervalSNil sets the value for IntervalS to be an explicit nil

### UnsetIntervalS
`func (o *GatewayHealthCheck) UnsetIntervalS()`

UnsetIntervalS ensures that no value is present for IntervalS, not even an explicit nil
### GetTimeoutS

`func (o *GatewayHealthCheck) GetTimeoutS() int32`

GetTimeoutS returns the TimeoutS field if non-nil, zero value otherwise.

### GetTimeoutSOk

`func (o *GatewayHealthCheck) GetTimeoutSOk() (*int32, bool)`

GetTimeoutSOk returns a tuple with the TimeoutS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutS

`func (o *GatewayHealthCheck) SetTimeoutS(v int32)`

SetTimeoutS sets TimeoutS field to given value.


### SetTimeoutSNil

`func (o *GatewayHealthCheck) SetTimeoutSNil(b bool)`

 SetTimeoutSNil sets the value for TimeoutS to be an explicit nil

### UnsetTimeoutS
`func (o *GatewayHealthCheck) UnsetTimeoutS()`

UnsetTimeoutS ensures that no value is present for TimeoutS, not even an explicit nil
### GetHealthyThreshold

`func (o *GatewayHealthCheck) GetHealthyThreshold() int32`

GetHealthyThreshold returns the HealthyThreshold field if non-nil, zero value otherwise.

### GetHealthyThresholdOk

`func (o *GatewayHealthCheck) GetHealthyThresholdOk() (*int32, bool)`

GetHealthyThresholdOk returns a tuple with the HealthyThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyThreshold

`func (o *GatewayHealthCheck) SetHealthyThreshold(v int32)`

SetHealthyThreshold sets HealthyThreshold field to given value.


### SetHealthyThresholdNil

`func (o *GatewayHealthCheck) SetHealthyThresholdNil(b bool)`

 SetHealthyThresholdNil sets the value for HealthyThreshold to be an explicit nil

### UnsetHealthyThreshold
`func (o *GatewayHealthCheck) UnsetHealthyThreshold()`

UnsetHealthyThreshold ensures that no value is present for HealthyThreshold, not even an explicit nil
### GetUnhealthyThreshold

`func (o *GatewayHealthCheck) GetUnhealthyThreshold() int32`

GetUnhealthyThreshold returns the UnhealthyThreshold field if non-nil, zero value otherwise.

### GetUnhealthyThresholdOk

`func (o *GatewayHealthCheck) GetUnhealthyThresholdOk() (*int32, bool)`

GetUnhealthyThresholdOk returns a tuple with the UnhealthyThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhealthyThreshold

`func (o *GatewayHealthCheck) SetUnhealthyThreshold(v int32)`

SetUnhealthyThreshold sets UnhealthyThreshold field to given value.


### SetUnhealthyThresholdNil

`func (o *GatewayHealthCheck) SetUnhealthyThresholdNil(b bool)`

 SetUnhealthyThresholdNil sets the value for UnhealthyThreshold to be an explicit nil

### UnsetUnhealthyThreshold
`func (o *GatewayHealthCheck) UnsetUnhealthyThreshold()`

UnsetUnhealthyThreshold ensures that no value is present for UnhealthyThreshold, not even an explicit nil
### GetPort

`func (o *GatewayHealthCheck) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GatewayHealthCheck) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GatewayHealthCheck) SetPort(v int32)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *GatewayHealthCheck) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *GatewayHealthCheck) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


