# GatewayBackendPoolWriteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Policy** | Pointer to [**GatewayBackendPoolPolicyEnum**](GatewayBackendPoolPolicyEnum.md) |  | [optional] 
**ProxyProtocol** | Pointer to **bool** | PROXY protocol v2 to the servers; tcp only. | [optional] 
**HealthCheck** | Pointer to [**GatewayHealthCheckWriteRequest**](GatewayHealthCheckWriteRequest.md) |  | [optional] 

## Methods

### NewGatewayBackendPoolWriteRequest

`func NewGatewayBackendPoolWriteRequest() *GatewayBackendPoolWriteRequest`

NewGatewayBackendPoolWriteRequest instantiates a new GatewayBackendPoolWriteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayBackendPoolWriteRequestWithDefaults

`func NewGatewayBackendPoolWriteRequestWithDefaults() *GatewayBackendPoolWriteRequest`

NewGatewayBackendPoolWriteRequestWithDefaults instantiates a new GatewayBackendPoolWriteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GatewayBackendPoolWriteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayBackendPoolWriteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayBackendPoolWriteRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayBackendPoolWriteRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicy

`func (o *GatewayBackendPoolWriteRequest) GetPolicy() GatewayBackendPoolPolicyEnum`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GatewayBackendPoolWriteRequest) GetPolicyOk() (*GatewayBackendPoolPolicyEnum, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GatewayBackendPoolWriteRequest) SetPolicy(v GatewayBackendPoolPolicyEnum)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *GatewayBackendPoolWriteRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetProxyProtocol

`func (o *GatewayBackendPoolWriteRequest) GetProxyProtocol() bool`

GetProxyProtocol returns the ProxyProtocol field if non-nil, zero value otherwise.

### GetProxyProtocolOk

`func (o *GatewayBackendPoolWriteRequest) GetProxyProtocolOk() (*bool, bool)`

GetProxyProtocolOk returns a tuple with the ProxyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyProtocol

`func (o *GatewayBackendPoolWriteRequest) SetProxyProtocol(v bool)`

SetProxyProtocol sets ProxyProtocol field to given value.

### HasProxyProtocol

`func (o *GatewayBackendPoolWriteRequest) HasProxyProtocol() bool`

HasProxyProtocol returns a boolean if a field has been set.

### GetHealthCheck

`func (o *GatewayBackendPoolWriteRequest) GetHealthCheck() GatewayHealthCheckWriteRequest`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *GatewayBackendPoolWriteRequest) GetHealthCheckOk() (*GatewayHealthCheckWriteRequest, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *GatewayBackendPoolWriteRequest) SetHealthCheck(v GatewayHealthCheckWriteRequest)`

SetHealthCheck sets HealthCheck field to given value.

### HasHealthCheck

`func (o *GatewayBackendPoolWriteRequest) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


