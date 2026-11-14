# GatewayBackendPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | [readonly] 
**Policy** | [**GatewayBackendPoolPolicyEnum**](GatewayBackendPoolPolicyEnum.md) |  | [readonly] 
**ProxyProtocol** | **bool** |  | [readonly] 
**HealthCheck** | [**GatewayHealthCheck**](GatewayHealthCheck.md) |  | [readonly] 
**ChangedRevision** | **int32** |  | [readonly] 
**Pending** | **bool** |  | [readonly] 

## Methods

### NewGatewayBackendPool

`func NewGatewayBackendPool(id int32, name string, policy GatewayBackendPoolPolicyEnum, proxyProtocol bool, healthCheck GatewayHealthCheck, changedRevision int32, pending bool, ) *GatewayBackendPool`

NewGatewayBackendPool instantiates a new GatewayBackendPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayBackendPoolWithDefaults

`func NewGatewayBackendPoolWithDefaults() *GatewayBackendPool`

NewGatewayBackendPoolWithDefaults instantiates a new GatewayBackendPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GatewayBackendPool) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayBackendPool) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayBackendPool) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *GatewayBackendPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayBackendPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayBackendPool) SetName(v string)`

SetName sets Name field to given value.


### GetPolicy

`func (o *GatewayBackendPool) GetPolicy() GatewayBackendPoolPolicyEnum`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GatewayBackendPool) GetPolicyOk() (*GatewayBackendPoolPolicyEnum, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GatewayBackendPool) SetPolicy(v GatewayBackendPoolPolicyEnum)`

SetPolicy sets Policy field to given value.


### GetProxyProtocol

`func (o *GatewayBackendPool) GetProxyProtocol() bool`

GetProxyProtocol returns the ProxyProtocol field if non-nil, zero value otherwise.

### GetProxyProtocolOk

`func (o *GatewayBackendPool) GetProxyProtocolOk() (*bool, bool)`

GetProxyProtocolOk returns a tuple with the ProxyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyProtocol

`func (o *GatewayBackendPool) SetProxyProtocol(v bool)`

SetProxyProtocol sets ProxyProtocol field to given value.


### GetHealthCheck

`func (o *GatewayBackendPool) GetHealthCheck() GatewayHealthCheck`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *GatewayBackendPool) GetHealthCheckOk() (*GatewayHealthCheck, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *GatewayBackendPool) SetHealthCheck(v GatewayHealthCheck)`

SetHealthCheck sets HealthCheck field to given value.


### GetChangedRevision

`func (o *GatewayBackendPool) GetChangedRevision() int32`

GetChangedRevision returns the ChangedRevision field if non-nil, zero value otherwise.

### GetChangedRevisionOk

`func (o *GatewayBackendPool) GetChangedRevisionOk() (*int32, bool)`

GetChangedRevisionOk returns a tuple with the ChangedRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedRevision

`func (o *GatewayBackendPool) SetChangedRevision(v int32)`

SetChangedRevision sets ChangedRevision field to given value.


### GetPending

`func (o *GatewayBackendPool) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *GatewayBackendPool) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *GatewayBackendPool) SetPending(v bool)`

SetPending sets Pending field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


