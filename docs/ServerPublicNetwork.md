# ServerPublicNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | Pointer to **string** |  | [optional] 
**Ipv4** | Pointer to **string** |  | [optional] 
**Ipv6** | Pointer to **string** |  | [optional] 
**Interfaces** | Pointer to [**[]ServerPublicInterface**](ServerPublicInterface.md) |  | [optional] 

## Methods

### NewServerPublicNetwork

`func NewServerPublicNetwork() *ServerPublicNetwork`

NewServerPublicNetwork instantiates a new ServerPublicNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerPublicNetworkWithDefaults

`func NewServerPublicNetworkWithDefaults() *ServerPublicNetwork`

NewServerPublicNetworkWithDefaults instantiates a new ServerPublicNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *ServerPublicNetwork) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *ServerPublicNetwork) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *ServerPublicNetwork) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *ServerPublicNetwork) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetIpv4

`func (o *ServerPublicNetwork) GetIpv4() string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *ServerPublicNetwork) GetIpv4Ok() (*string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *ServerPublicNetwork) SetIpv4(v string)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *ServerPublicNetwork) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *ServerPublicNetwork) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *ServerPublicNetwork) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *ServerPublicNetwork) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *ServerPublicNetwork) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetInterfaces

`func (o *ServerPublicNetwork) GetInterfaces() []ServerPublicInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *ServerPublicNetwork) GetInterfacesOk() (*[]ServerPublicInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *ServerPublicNetwork) SetInterfaces(v []ServerPublicInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *ServerPublicNetwork) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


