# PatchedPrivateNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Slug** | Pointer to **string** | CIDR format | [optional] [readonly] 
**Address** | Pointer to **string** | CIDR format | [optional] 
**Gateway** | Pointer to **NullableString** |  | [optional] 
**Provisioned** | Pointer to **bool** |  | [optional] [readonly] 
**Servers** | Pointer to **[]map[string]string** |  | [optional] [readonly] 

## Methods

### NewPatchedPrivateNetwork

`func NewPatchedPrivateNetwork() *PatchedPrivateNetwork`

NewPatchedPrivateNetwork instantiates a new PatchedPrivateNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedPrivateNetworkWithDefaults

`func NewPatchedPrivateNetworkWithDefaults() *PatchedPrivateNetwork`

NewPatchedPrivateNetworkWithDefaults instantiates a new PatchedPrivateNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedPrivateNetwork) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedPrivateNetwork) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedPrivateNetwork) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedPrivateNetwork) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedPrivateNetwork) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedPrivateNetwork) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedPrivateNetwork) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedPrivateNetwork) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetAddress

`func (o *PatchedPrivateNetwork) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PatchedPrivateNetwork) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PatchedPrivateNetwork) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PatchedPrivateNetwork) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGateway

`func (o *PatchedPrivateNetwork) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *PatchedPrivateNetwork) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *PatchedPrivateNetwork) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *PatchedPrivateNetwork) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *PatchedPrivateNetwork) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *PatchedPrivateNetwork) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetProvisioned

`func (o *PatchedPrivateNetwork) GetProvisioned() bool`

GetProvisioned returns the Provisioned field if non-nil, zero value otherwise.

### GetProvisionedOk

`func (o *PatchedPrivateNetwork) GetProvisionedOk() (*bool, bool)`

GetProvisionedOk returns a tuple with the Provisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioned

`func (o *PatchedPrivateNetwork) SetProvisioned(v bool)`

SetProvisioned sets Provisioned field to given value.

### HasProvisioned

`func (o *PatchedPrivateNetwork) HasProvisioned() bool`

HasProvisioned returns a boolean if a field has been set.

### GetServers

`func (o *PatchedPrivateNetwork) GetServers() []map[string]string`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *PatchedPrivateNetwork) GetServersOk() (*[]map[string]string, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *PatchedPrivateNetwork) SetServers(v []map[string]string)`

SetServers sets Servers field to given value.

### HasServers

`func (o *PatchedPrivateNetwork) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


