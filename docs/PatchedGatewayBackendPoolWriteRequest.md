# PatchedGatewayBackendPoolWriteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Policy** | Pointer to [**GatewayBackendPoolPolicyEnum**](GatewayBackendPoolPolicyEnum.md) |  | [optional] 
**ProxyProtocol** | Pointer to **bool** | PROXY protocol v2 to the servers; tcp only. | [optional] 
**HealthCheck** | Pointer to [**GatewayHealthCheckWriteRequest**](GatewayHealthCheckWriteRequest.md) |  | [optional] 

## Methods

### NewPatchedGatewayBackendPoolWriteRequest

`func NewPatchedGatewayBackendPoolWriteRequest() *PatchedGatewayBackendPoolWriteRequest`

NewPatchedGatewayBackendPoolWriteRequest instantiates a new PatchedGatewayBackendPoolWriteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedGatewayBackendPoolWriteRequestWithDefaults

`func NewPatchedGatewayBackendPoolWriteRequestWithDefaults() *PatchedGatewayBackendPoolWriteRequest`

NewPatchedGatewayBackendPoolWriteRequestWithDefaults instantiates a new PatchedGatewayBackendPoolWriteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedGatewayBackendPoolWriteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedGatewayBackendPoolWriteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedGatewayBackendPoolWriteRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedGatewayBackendPoolWriteRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicy

`func (o *PatchedGatewayBackendPoolWriteRequest) GetPolicy() GatewayBackendPoolPolicyEnum`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PatchedGatewayBackendPoolWriteRequest) GetPolicyOk() (*GatewayBackendPoolPolicyEnum, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PatchedGatewayBackendPoolWriteRequest) SetPolicy(v GatewayBackendPoolPolicyEnum)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PatchedGatewayBackendPoolWriteRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetProxyProtocol

`func (o *PatchedGatewayBackendPoolWriteRequest) GetProxyProtocol() bool`

GetProxyProtocol returns the ProxyProtocol field if non-nil, zero value otherwise.

### GetProxyProtocolOk

`func (o *PatchedGatewayBackendPoolWriteRequest) GetProxyProtocolOk() (*bool, bool)`

GetProxyProtocolOk returns a tuple with the ProxyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyProtocol

`func (o *PatchedGatewayBackendPoolWriteRequest) SetProxyProtocol(v bool)`

SetProxyProtocol sets ProxyProtocol field to given value.

### HasProxyProtocol

`func (o *PatchedGatewayBackendPoolWriteRequest) HasProxyProtocol() bool`

HasProxyProtocol returns a boolean if a field has been set.

### GetHealthCheck

`func (o *PatchedGatewayBackendPoolWriteRequest) GetHealthCheck() GatewayHealthCheckWriteRequest`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *PatchedGatewayBackendPoolWriteRequest) GetHealthCheckOk() (*GatewayHealthCheckWriteRequest, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *PatchedGatewayBackendPoolWriteRequest) SetHealthCheck(v GatewayHealthCheckWriteRequest)`

SetHealthCheck sets HealthCheck field to given value.

### HasHealthCheck

`func (o *PatchedGatewayBackendPoolWriteRequest) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


