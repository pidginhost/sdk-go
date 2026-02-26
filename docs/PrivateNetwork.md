# PrivateNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Slug** | **string** | CIDR format | [readonly] 
**Address** | **string** | CIDR format | 
**Gateway** | Pointer to **NullableString** |  | [optional] 
**Provisioned** | **bool** |  | [readonly] 
**Servers** | **[]map[string]string** |  | [readonly] 

## Methods

### NewPrivateNetwork

`func NewPrivateNetwork(id int32, slug string, address string, provisioned bool, servers []map[string]string, ) *PrivateNetwork`

NewPrivateNetwork instantiates a new PrivateNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateNetworkWithDefaults

`func NewPrivateNetworkWithDefaults() *PrivateNetwork`

NewPrivateNetworkWithDefaults instantiates a new PrivateNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateNetwork) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateNetwork) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateNetwork) SetId(v int32)`

SetId sets Id field to given value.


### GetSlug

`func (o *PrivateNetwork) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PrivateNetwork) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PrivateNetwork) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetAddress

`func (o *PrivateNetwork) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PrivateNetwork) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PrivateNetwork) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetGateway

`func (o *PrivateNetwork) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *PrivateNetwork) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *PrivateNetwork) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *PrivateNetwork) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *PrivateNetwork) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *PrivateNetwork) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetProvisioned

`func (o *PrivateNetwork) GetProvisioned() bool`

GetProvisioned returns the Provisioned field if non-nil, zero value otherwise.

### GetProvisionedOk

`func (o *PrivateNetwork) GetProvisionedOk() (*bool, bool)`

GetProvisionedOk returns a tuple with the Provisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioned

`func (o *PrivateNetwork) SetProvisioned(v bool)`

SetProvisioned sets Provisioned field to given value.


### GetServers

`func (o *PrivateNetwork) GetServers() []map[string]string`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *PrivateNetwork) GetServersOk() (*[]map[string]string, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *PrivateNetwork) SetServers(v []map[string]string)`

SetServers sets Servers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


