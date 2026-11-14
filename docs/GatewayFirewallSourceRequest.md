# GatewayFirewallSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to **int32** | A private network the gateway is attached to. | [optional] 
**Cidr** | Pointer to **string** | An IPv4 network, such as 10.0.0.0/24. | [optional] 

## Methods

### NewGatewayFirewallSourceRequest

`func NewGatewayFirewallSourceRequest() *GatewayFirewallSourceRequest`

NewGatewayFirewallSourceRequest instantiates a new GatewayFirewallSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayFirewallSourceRequestWithDefaults

`func NewGatewayFirewallSourceRequestWithDefaults() *GatewayFirewallSourceRequest`

NewGatewayFirewallSourceRequestWithDefaults instantiates a new GatewayFirewallSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *GatewayFirewallSourceRequest) GetNetwork() int32`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GatewayFirewallSourceRequest) GetNetworkOk() (*int32, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GatewayFirewallSourceRequest) SetNetwork(v int32)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GatewayFirewallSourceRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCidr

`func (o *GatewayFirewallSourceRequest) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *GatewayFirewallSourceRequest) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *GatewayFirewallSourceRequest) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *GatewayFirewallSourceRequest) HasCidr() bool`

HasCidr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


