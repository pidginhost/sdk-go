# FloatingIPv6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Address** | **string** |  | [readonly] 
**Gateway** | **string** |  | [readonly] 
**Prefix** | **int32** |  | [readonly] 
**Label** | Pointer to **string** |  | [optional] 
**AuthorizedVmCount** | **int32** |  | [readonly] 
**CreatedAt** | **string** |  | [readonly] 

## Methods

### NewFloatingIPv6

`func NewFloatingIPv6(id int32, address string, gateway string, prefix int32, authorizedVmCount int32, createdAt string, ) *FloatingIPv6`

NewFloatingIPv6 instantiates a new FloatingIPv6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloatingIPv6WithDefaults

`func NewFloatingIPv6WithDefaults() *FloatingIPv6`

NewFloatingIPv6WithDefaults instantiates a new FloatingIPv6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FloatingIPv6) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FloatingIPv6) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FloatingIPv6) SetId(v int32)`

SetId sets Id field to given value.


### GetAddress

`func (o *FloatingIPv6) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FloatingIPv6) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FloatingIPv6) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetGateway

`func (o *FloatingIPv6) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *FloatingIPv6) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *FloatingIPv6) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetPrefix

`func (o *FloatingIPv6) GetPrefix() int32`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *FloatingIPv6) GetPrefixOk() (*int32, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *FloatingIPv6) SetPrefix(v int32)`

SetPrefix sets Prefix field to given value.


### GetLabel

`func (o *FloatingIPv6) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FloatingIPv6) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FloatingIPv6) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FloatingIPv6) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetAuthorizedVmCount

`func (o *FloatingIPv6) GetAuthorizedVmCount() int32`

GetAuthorizedVmCount returns the AuthorizedVmCount field if non-nil, zero value otherwise.

### GetAuthorizedVmCountOk

`func (o *FloatingIPv6) GetAuthorizedVmCountOk() (*int32, bool)`

GetAuthorizedVmCountOk returns a tuple with the AuthorizedVmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedVmCount

`func (o *FloatingIPv6) SetAuthorizedVmCount(v int32)`

SetAuthorizedVmCount sets AuthorizedVmCount field to given value.


### GetCreatedAt

`func (o *FloatingIPv6) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FloatingIPv6) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FloatingIPv6) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


