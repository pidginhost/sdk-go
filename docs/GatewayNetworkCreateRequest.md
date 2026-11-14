# GatewayNetworkCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **int32** | A private network of the service. | 
**Role** | [**GatewayNetworkRoleEnum**](GatewayNetworkRoleEnum.md) |  | 

## Methods

### NewGatewayNetworkCreateRequest

`func NewGatewayNetworkCreateRequest(network int32, role GatewayNetworkRoleEnum, ) *GatewayNetworkCreateRequest`

NewGatewayNetworkCreateRequest instantiates a new GatewayNetworkCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayNetworkCreateRequestWithDefaults

`func NewGatewayNetworkCreateRequestWithDefaults() *GatewayNetworkCreateRequest`

NewGatewayNetworkCreateRequestWithDefaults instantiates a new GatewayNetworkCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *GatewayNetworkCreateRequest) GetNetwork() int32`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GatewayNetworkCreateRequest) GetNetworkOk() (*int32, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GatewayNetworkCreateRequest) SetNetwork(v int32)`

SetNetwork sets Network field to given value.


### GetRole

`func (o *GatewayNetworkCreateRequest) GetRole() GatewayNetworkRoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GatewayNetworkCreateRequest) GetRoleOk() (*GatewayNetworkRoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GatewayNetworkCreateRequest) SetRole(v GatewayNetworkRoleEnum)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


