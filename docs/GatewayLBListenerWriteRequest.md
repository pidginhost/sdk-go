# GatewayLBListenerWriteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to [**GatewayForwardProtocolEnum**](GatewayForwardProtocolEnum.md) |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Mode** | Pointer to [**GatewayLBListenerModeEnum**](GatewayLBListenerModeEnum.md) | tcp, udp or http; https is not available yet.  * &#x60;tcp&#x60; - tcp * &#x60;udp&#x60; - udp * &#x60;http&#x60; - http * &#x60;https&#x60; - https | [optional] 
**Pool** | Pointer to **NullableInt32** | A backend pool&#39;s id; tcp and udp only. | [optional] 
**AllowedSources** | Pointer to **[]string** | IPv4 and IPv6 networks; empty means any source. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewGatewayLBListenerWriteRequest

`func NewGatewayLBListenerWriteRequest() *GatewayLBListenerWriteRequest`

NewGatewayLBListenerWriteRequest instantiates a new GatewayLBListenerWriteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayLBListenerWriteRequestWithDefaults

`func NewGatewayLBListenerWriteRequestWithDefaults() *GatewayLBListenerWriteRequest`

NewGatewayLBListenerWriteRequestWithDefaults instantiates a new GatewayLBListenerWriteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *GatewayLBListenerWriteRequest) GetProtocol() GatewayForwardProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GatewayLBListenerWriteRequest) GetProtocolOk() (*GatewayForwardProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GatewayLBListenerWriteRequest) SetProtocol(v GatewayForwardProtocolEnum)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GatewayLBListenerWriteRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPort

`func (o *GatewayLBListenerWriteRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GatewayLBListenerWriteRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GatewayLBListenerWriteRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GatewayLBListenerWriteRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetMode

`func (o *GatewayLBListenerWriteRequest) GetMode() GatewayLBListenerModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GatewayLBListenerWriteRequest) GetModeOk() (*GatewayLBListenerModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GatewayLBListenerWriteRequest) SetMode(v GatewayLBListenerModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GatewayLBListenerWriteRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPool

`func (o *GatewayLBListenerWriteRequest) GetPool() int32`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *GatewayLBListenerWriteRequest) GetPoolOk() (*int32, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *GatewayLBListenerWriteRequest) SetPool(v int32)`

SetPool sets Pool field to given value.

### HasPool

`func (o *GatewayLBListenerWriteRequest) HasPool() bool`

HasPool returns a boolean if a field has been set.

### SetPoolNil

`func (o *GatewayLBListenerWriteRequest) SetPoolNil(b bool)`

 SetPoolNil sets the value for Pool to be an explicit nil

### UnsetPool
`func (o *GatewayLBListenerWriteRequest) UnsetPool()`

UnsetPool ensures that no value is present for Pool, not even an explicit nil
### GetAllowedSources

`func (o *GatewayLBListenerWriteRequest) GetAllowedSources() []string`

GetAllowedSources returns the AllowedSources field if non-nil, zero value otherwise.

### GetAllowedSourcesOk

`func (o *GatewayLBListenerWriteRequest) GetAllowedSourcesOk() (*[]string, bool)`

GetAllowedSourcesOk returns a tuple with the AllowedSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSources

`func (o *GatewayLBListenerWriteRequest) SetAllowedSources(v []string)`

SetAllowedSources sets AllowedSources field to given value.

### HasAllowedSources

`func (o *GatewayLBListenerWriteRequest) HasAllowedSources() bool`

HasAllowedSources returns a boolean if a field has been set.

### GetEnabled

`func (o *GatewayLBListenerWriteRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GatewayLBListenerWriteRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GatewayLBListenerWriteRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GatewayLBListenerWriteRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetComment

`func (o *GatewayLBListenerWriteRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GatewayLBListenerWriteRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GatewayLBListenerWriteRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GatewayLBListenerWriteRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


