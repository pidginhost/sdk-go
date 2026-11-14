# GatewayDeletePreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **int32** |  | 
**Servers** | [**[]GatewayPreviewServer**](GatewayPreviewServer.md) |  | 
**Networks** | [**[]GatewayPreviewNetwork**](GatewayPreviewNetwork.md) |  | 

## Methods

### NewGatewayDeletePreview

`func NewGatewayDeletePreview(revision int32, servers []GatewayPreviewServer, networks []GatewayPreviewNetwork, ) *GatewayDeletePreview`

NewGatewayDeletePreview instantiates a new GatewayDeletePreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayDeletePreviewWithDefaults

`func NewGatewayDeletePreviewWithDefaults() *GatewayDeletePreview`

NewGatewayDeletePreviewWithDefaults instantiates a new GatewayDeletePreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *GatewayDeletePreview) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *GatewayDeletePreview) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *GatewayDeletePreview) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetServers

`func (o *GatewayDeletePreview) GetServers() []GatewayPreviewServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GatewayDeletePreview) GetServersOk() (*[]GatewayPreviewServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GatewayDeletePreview) SetServers(v []GatewayPreviewServer)`

SetServers sets Servers field to given value.


### GetNetworks

`func (o *GatewayDeletePreview) GetNetworks() []GatewayPreviewNetwork`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *GatewayDeletePreview) GetNetworksOk() (*[]GatewayPreviewNetwork, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *GatewayDeletePreview) SetNetworks(v []GatewayPreviewNetwork)`

SetNetworks sets Networks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


