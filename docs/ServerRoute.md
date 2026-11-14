# ServerRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | **int32** |  | 
**Hostname** | **string** |  | 
**Power** | [**RoutePowerEnum**](RoutePowerEnum.md) |  | 
**Config** | [**RouteConfigEnum**](RouteConfigEnum.md) | The spec 6.2 rule over current config.  * &#x60;ok&#x60; - ok * &#x60;missing&#x60; - missing * &#x60;duplicate&#x60; - duplicate * &#x60;obsolete&#x60; - obsolete * &#x60;unknown&#x60; - unknown | 
**Pending** | [**NullableRouteConfigEnum**](RouteConfigEnum.md) | The rule over the config a reboot applies; null without one.  * &#x60;ok&#x60; - ok * &#x60;missing&#x60; - missing * &#x60;duplicate&#x60; - duplicate * &#x60;obsolete&#x60; - obsolete * &#x60;unknown&#x60; - unknown | 
**RoutesHere** | **bool** | The rule gives this network&#39;s interface the default route. | 
**Acknowledged** | **bool** | The last route apply&#39;s config and boot still hold. | 

## Methods

### NewServerRoute

`func NewServerRoute(server int32, hostname string, power RoutePowerEnum, config RouteConfigEnum, pending NullableRouteConfigEnum, routesHere bool, acknowledged bool, ) *ServerRoute`

NewServerRoute instantiates a new ServerRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerRouteWithDefaults

`func NewServerRouteWithDefaults() *ServerRoute`

NewServerRouteWithDefaults instantiates a new ServerRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *ServerRoute) GetServer() int32`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ServerRoute) GetServerOk() (*int32, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ServerRoute) SetServer(v int32)`

SetServer sets Server field to given value.


### GetHostname

`func (o *ServerRoute) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ServerRoute) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ServerRoute) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetPower

`func (o *ServerRoute) GetPower() RoutePowerEnum`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *ServerRoute) GetPowerOk() (*RoutePowerEnum, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *ServerRoute) SetPower(v RoutePowerEnum)`

SetPower sets Power field to given value.


### GetConfig

`func (o *ServerRoute) GetConfig() RouteConfigEnum`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ServerRoute) GetConfigOk() (*RouteConfigEnum, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ServerRoute) SetConfig(v RouteConfigEnum)`

SetConfig sets Config field to given value.


### GetPending

`func (o *ServerRoute) GetPending() RouteConfigEnum`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *ServerRoute) GetPendingOk() (*RouteConfigEnum, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *ServerRoute) SetPending(v RouteConfigEnum)`

SetPending sets Pending field to given value.


### SetPendingNil

`func (o *ServerRoute) SetPendingNil(b bool)`

 SetPendingNil sets the value for Pending to be an explicit nil

### UnsetPending
`func (o *ServerRoute) UnsetPending()`

UnsetPending ensures that no value is present for Pending, not even an explicit nil
### GetRoutesHere

`func (o *ServerRoute) GetRoutesHere() bool`

GetRoutesHere returns the RoutesHere field if non-nil, zero value otherwise.

### GetRoutesHereOk

`func (o *ServerRoute) GetRoutesHereOk() (*bool, bool)`

GetRoutesHereOk returns a tuple with the RoutesHere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutesHere

`func (o *ServerRoute) SetRoutesHere(v bool)`

SetRoutesHere sets RoutesHere field to given value.


### GetAcknowledged

`func (o *ServerRoute) GetAcknowledged() bool`

GetAcknowledged returns the Acknowledged field if non-nil, zero value otherwise.

### GetAcknowledgedOk

`func (o *ServerRoute) GetAcknowledgedOk() (*bool, bool)`

GetAcknowledgedOk returns a tuple with the Acknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledged

`func (o *ServerRoute) SetAcknowledged(v bool)`

SetAcknowledged sets Acknowledged field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


