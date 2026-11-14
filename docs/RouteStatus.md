# RouteStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **int32** |  | 
**Gateway** | **NullableInt32** | The gateway that routes for the network, or null. | 
**Servers** | [**[]ServerRoute**](ServerRoute.md) |  | 
**Operations** | [**[]GatewayOperation**](GatewayOperation.md) | The network&#39;s route applies, newest first. | 

## Methods

### NewRouteStatus

`func NewRouteStatus(network int32, gateway NullableInt32, servers []ServerRoute, operations []GatewayOperation, ) *RouteStatus`

NewRouteStatus instantiates a new RouteStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteStatusWithDefaults

`func NewRouteStatusWithDefaults() *RouteStatus`

NewRouteStatusWithDefaults instantiates a new RouteStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *RouteStatus) GetNetwork() int32`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *RouteStatus) GetNetworkOk() (*int32, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *RouteStatus) SetNetwork(v int32)`

SetNetwork sets Network field to given value.


### GetGateway

`func (o *RouteStatus) GetGateway() int32`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *RouteStatus) GetGatewayOk() (*int32, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *RouteStatus) SetGateway(v int32)`

SetGateway sets Gateway field to given value.


### SetGatewayNil

`func (o *RouteStatus) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *RouteStatus) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetServers

`func (o *RouteStatus) GetServers() []ServerRoute`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *RouteStatus) GetServersOk() (*[]ServerRoute, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *RouteStatus) SetServers(v []ServerRoute)`

SetServers sets Servers field to given value.


### GetOperations

`func (o *RouteStatus) GetOperations() []GatewayOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *RouteStatus) GetOperationsOk() (*[]GatewayOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *RouteStatus) SetOperations(v []GatewayOperation)`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


