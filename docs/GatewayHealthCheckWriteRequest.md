# GatewayHealthCheckWriteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**GatewayHealthCheckTypeEnum**](GatewayHealthCheckTypeEnum.md) |  | [optional] 
**Path** | Pointer to **string** | http only. | [optional] 
**StatusMin** | Pointer to **int32** | http only. | [optional] 
**StatusMax** | Pointer to **int32** | http only. | [optional] 
**IntervalS** | Pointer to **int32** |  | [optional] 
**TimeoutS** | Pointer to **int32** | Below interval_s. | [optional] 
**HealthyThreshold** | Pointer to **int32** |  | [optional] 
**UnhealthyThreshold** | Pointer to **int32** |  | [optional] 
**Port** | Pointer to **NullableInt32** | Null checks each server&#39;s own port. | [optional] 

## Methods

### NewGatewayHealthCheckWriteRequest

`func NewGatewayHealthCheckWriteRequest() *GatewayHealthCheckWriteRequest`

NewGatewayHealthCheckWriteRequest instantiates a new GatewayHealthCheckWriteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayHealthCheckWriteRequestWithDefaults

`func NewGatewayHealthCheckWriteRequestWithDefaults() *GatewayHealthCheckWriteRequest`

NewGatewayHealthCheckWriteRequestWithDefaults instantiates a new GatewayHealthCheckWriteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GatewayHealthCheckWriteRequest) GetType() GatewayHealthCheckTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayHealthCheckWriteRequest) GetTypeOk() (*GatewayHealthCheckTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayHealthCheckWriteRequest) SetType(v GatewayHealthCheckTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayHealthCheckWriteRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPath

`func (o *GatewayHealthCheckWriteRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GatewayHealthCheckWriteRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GatewayHealthCheckWriteRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GatewayHealthCheckWriteRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetStatusMin

`func (o *GatewayHealthCheckWriteRequest) GetStatusMin() int32`

GetStatusMin returns the StatusMin field if non-nil, zero value otherwise.

### GetStatusMinOk

`func (o *GatewayHealthCheckWriteRequest) GetStatusMinOk() (*int32, bool)`

GetStatusMinOk returns a tuple with the StatusMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMin

`func (o *GatewayHealthCheckWriteRequest) SetStatusMin(v int32)`

SetStatusMin sets StatusMin field to given value.

### HasStatusMin

`func (o *GatewayHealthCheckWriteRequest) HasStatusMin() bool`

HasStatusMin returns a boolean if a field has been set.

### GetStatusMax

`func (o *GatewayHealthCheckWriteRequest) GetStatusMax() int32`

GetStatusMax returns the StatusMax field if non-nil, zero value otherwise.

### GetStatusMaxOk

`func (o *GatewayHealthCheckWriteRequest) GetStatusMaxOk() (*int32, bool)`

GetStatusMaxOk returns a tuple with the StatusMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMax

`func (o *GatewayHealthCheckWriteRequest) SetStatusMax(v int32)`

SetStatusMax sets StatusMax field to given value.

### HasStatusMax

`func (o *GatewayHealthCheckWriteRequest) HasStatusMax() bool`

HasStatusMax returns a boolean if a field has been set.

### GetIntervalS

`func (o *GatewayHealthCheckWriteRequest) GetIntervalS() int32`

GetIntervalS returns the IntervalS field if non-nil, zero value otherwise.

### GetIntervalSOk

`func (o *GatewayHealthCheckWriteRequest) GetIntervalSOk() (*int32, bool)`

GetIntervalSOk returns a tuple with the IntervalS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalS

`func (o *GatewayHealthCheckWriteRequest) SetIntervalS(v int32)`

SetIntervalS sets IntervalS field to given value.

### HasIntervalS

`func (o *GatewayHealthCheckWriteRequest) HasIntervalS() bool`

HasIntervalS returns a boolean if a field has been set.

### GetTimeoutS

`func (o *GatewayHealthCheckWriteRequest) GetTimeoutS() int32`

GetTimeoutS returns the TimeoutS field if non-nil, zero value otherwise.

### GetTimeoutSOk

`func (o *GatewayHealthCheckWriteRequest) GetTimeoutSOk() (*int32, bool)`

GetTimeoutSOk returns a tuple with the TimeoutS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutS

`func (o *GatewayHealthCheckWriteRequest) SetTimeoutS(v int32)`

SetTimeoutS sets TimeoutS field to given value.

### HasTimeoutS

`func (o *GatewayHealthCheckWriteRequest) HasTimeoutS() bool`

HasTimeoutS returns a boolean if a field has been set.

### GetHealthyThreshold

`func (o *GatewayHealthCheckWriteRequest) GetHealthyThreshold() int32`

GetHealthyThreshold returns the HealthyThreshold field if non-nil, zero value otherwise.

### GetHealthyThresholdOk

`func (o *GatewayHealthCheckWriteRequest) GetHealthyThresholdOk() (*int32, bool)`

GetHealthyThresholdOk returns a tuple with the HealthyThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyThreshold

`func (o *GatewayHealthCheckWriteRequest) SetHealthyThreshold(v int32)`

SetHealthyThreshold sets HealthyThreshold field to given value.

### HasHealthyThreshold

`func (o *GatewayHealthCheckWriteRequest) HasHealthyThreshold() bool`

HasHealthyThreshold returns a boolean if a field has been set.

### GetUnhealthyThreshold

`func (o *GatewayHealthCheckWriteRequest) GetUnhealthyThreshold() int32`

GetUnhealthyThreshold returns the UnhealthyThreshold field if non-nil, zero value otherwise.

### GetUnhealthyThresholdOk

`func (o *GatewayHealthCheckWriteRequest) GetUnhealthyThresholdOk() (*int32, bool)`

GetUnhealthyThresholdOk returns a tuple with the UnhealthyThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhealthyThreshold

`func (o *GatewayHealthCheckWriteRequest) SetUnhealthyThreshold(v int32)`

SetUnhealthyThreshold sets UnhealthyThreshold field to given value.

### HasUnhealthyThreshold

`func (o *GatewayHealthCheckWriteRequest) HasUnhealthyThreshold() bool`

HasUnhealthyThreshold returns a boolean if a field has been set.

### GetPort

`func (o *GatewayHealthCheckWriteRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GatewayHealthCheckWriteRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GatewayHealthCheckWriteRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GatewayHealthCheckWriteRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *GatewayHealthCheckWriteRequest) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *GatewayHealthCheckWriteRequest) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


