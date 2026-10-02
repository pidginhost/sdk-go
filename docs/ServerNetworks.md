# ServerNetworks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Public** | [**ServerPublicNetwork**](ServerPublicNetwork.md) |  | 
**Private** | [**[]ServerPrivateInterface**](ServerPrivateInterface.md) |  | 

## Methods

### NewServerNetworks

`func NewServerNetworks(public ServerPublicNetwork, private []ServerPrivateInterface, ) *ServerNetworks`

NewServerNetworks instantiates a new ServerNetworks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerNetworksWithDefaults

`func NewServerNetworksWithDefaults() *ServerNetworks`

NewServerNetworksWithDefaults instantiates a new ServerNetworks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublic

`func (o *ServerNetworks) GetPublic() ServerPublicNetwork`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *ServerNetworks) GetPublicOk() (*ServerPublicNetwork, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *ServerNetworks) SetPublic(v ServerPublicNetwork)`

SetPublic sets Public field to given value.


### GetPrivate

`func (o *ServerNetworks) GetPrivate() []ServerPrivateInterface`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *ServerNetworks) GetPrivateOk() (*[]ServerPrivateInterface, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *ServerNetworks) SetPrivate(v []ServerPrivateInterface)`

SetPrivate sets Private field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


