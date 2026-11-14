# GatewayForwardTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **int32** | An attached private network&#39;s id. | 
**Address** | **string** | A host of that network, other than the gateway. | 
**PortStart** | **int32** | The first public port goes here, the next ones follow in order. | 

## Methods

### NewGatewayForwardTarget

`func NewGatewayForwardTarget(network int32, address string, portStart int32, ) *GatewayForwardTarget`

NewGatewayForwardTarget instantiates a new GatewayForwardTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayForwardTargetWithDefaults

`func NewGatewayForwardTargetWithDefaults() *GatewayForwardTarget`

NewGatewayForwardTargetWithDefaults instantiates a new GatewayForwardTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *GatewayForwardTarget) GetNetwork() int32`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GatewayForwardTarget) GetNetworkOk() (*int32, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GatewayForwardTarget) SetNetwork(v int32)`

SetNetwork sets Network field to given value.


### GetAddress

`func (o *GatewayForwardTarget) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GatewayForwardTarget) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GatewayForwardTarget) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPortStart

`func (o *GatewayForwardTarget) GetPortStart() int32`

GetPortStart returns the PortStart field if non-nil, zero value otherwise.

### GetPortStartOk

`func (o *GatewayForwardTarget) GetPortStartOk() (*int32, bool)`

GetPortStartOk returns a tuple with the PortStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStart

`func (o *GatewayForwardTarget) SetPortStart(v int32)`

SetPortStart sets PortStart field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


