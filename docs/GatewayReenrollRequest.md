# GatewayReenrollRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | Pointer to **int32** | The gateway&#39;s desired revision from the preview; without it, a preview. | [optional] 
**Servers** | Pointer to **[]int32** | The preview&#39;s server ids. | [optional] 
**Networks** | Pointer to **[]int32** | The preview&#39;s network ids. | [optional] 

## Methods

### NewGatewayReenrollRequest

`func NewGatewayReenrollRequest() *GatewayReenrollRequest`

NewGatewayReenrollRequest instantiates a new GatewayReenrollRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayReenrollRequestWithDefaults

`func NewGatewayReenrollRequestWithDefaults() *GatewayReenrollRequest`

NewGatewayReenrollRequestWithDefaults instantiates a new GatewayReenrollRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *GatewayReenrollRequest) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *GatewayReenrollRequest) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *GatewayReenrollRequest) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *GatewayReenrollRequest) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetServers

`func (o *GatewayReenrollRequest) GetServers() []int32`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GatewayReenrollRequest) GetServersOk() (*[]int32, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GatewayReenrollRequest) SetServers(v []int32)`

SetServers sets Servers field to given value.

### HasServers

`func (o *GatewayReenrollRequest) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetNetworks

`func (o *GatewayReenrollRequest) GetNetworks() []int32`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *GatewayReenrollRequest) GetNetworksOk() (*[]int32, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *GatewayReenrollRequest) SetNetworks(v []int32)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *GatewayReenrollRequest) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


