# GatewayPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **int32** |  | 
**Servers** | [**[]GatewayPreviewServer**](GatewayPreviewServer.md) |  | 

## Methods

### NewGatewayPreview

`func NewGatewayPreview(revision int32, servers []GatewayPreviewServer, ) *GatewayPreview`

NewGatewayPreview instantiates a new GatewayPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayPreviewWithDefaults

`func NewGatewayPreviewWithDefaults() *GatewayPreview`

NewGatewayPreviewWithDefaults instantiates a new GatewayPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *GatewayPreview) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *GatewayPreview) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *GatewayPreview) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetServers

`func (o *GatewayPreview) GetServers() []GatewayPreviewServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GatewayPreview) GetServersOk() (*[]GatewayPreviewServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GatewayPreview) SetServers(v []GatewayPreviewServer)`

SetServers sets Servers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


