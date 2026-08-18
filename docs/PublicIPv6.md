# PublicIPv6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Slug** | **string** |  | [readonly] 
**Address** | **string** |  | [readonly] 
**Gateway** | **string** |  | [readonly] 
**Prefix** | **int32** |  | [readonly] 
**Attached** | **bool** |  | [readonly] 
**Server** | **string** |  | [readonly] 

## Methods

### NewPublicIPv6

`func NewPublicIPv6(id int32, slug string, address string, gateway string, prefix int32, attached bool, server string, ) *PublicIPv6`

NewPublicIPv6 instantiates a new PublicIPv6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicIPv6WithDefaults

`func NewPublicIPv6WithDefaults() *PublicIPv6`

NewPublicIPv6WithDefaults instantiates a new PublicIPv6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PublicIPv6) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicIPv6) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicIPv6) SetId(v int32)`

SetId sets Id field to given value.


### GetSlug

`func (o *PublicIPv6) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PublicIPv6) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PublicIPv6) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetAddress

`func (o *PublicIPv6) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PublicIPv6) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PublicIPv6) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetGateway

`func (o *PublicIPv6) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *PublicIPv6) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *PublicIPv6) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetPrefix

`func (o *PublicIPv6) GetPrefix() int32`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PublicIPv6) GetPrefixOk() (*int32, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PublicIPv6) SetPrefix(v int32)`

SetPrefix sets Prefix field to given value.


### GetAttached

`func (o *PublicIPv6) GetAttached() bool`

GetAttached returns the Attached field if non-nil, zero value otherwise.

### GetAttachedOk

`func (o *PublicIPv6) GetAttachedOk() (*bool, bool)`

GetAttachedOk returns a tuple with the Attached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttached

`func (o *PublicIPv6) SetAttached(v bool)`

SetAttached sets Attached field to given value.


### GetServer

`func (o *PublicIPv6) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *PublicIPv6) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *PublicIPv6) SetServer(v string)`

SetServer sets Server field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


