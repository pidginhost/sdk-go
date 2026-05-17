# FloatingIPAuthorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**ServerId** | **int32** |  | [readonly] 
**ServerHostname** | **string** |  | [readonly] 
**CreatedAt** | **string** |  | [readonly] 

## Methods

### NewFloatingIPAuthorization

`func NewFloatingIPAuthorization(id int32, serverId int32, serverHostname string, createdAt string, ) *FloatingIPAuthorization`

NewFloatingIPAuthorization instantiates a new FloatingIPAuthorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloatingIPAuthorizationWithDefaults

`func NewFloatingIPAuthorizationWithDefaults() *FloatingIPAuthorization`

NewFloatingIPAuthorizationWithDefaults instantiates a new FloatingIPAuthorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FloatingIPAuthorization) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FloatingIPAuthorization) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FloatingIPAuthorization) SetId(v int32)`

SetId sets Id field to given value.


### GetServerId

`func (o *FloatingIPAuthorization) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *FloatingIPAuthorization) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *FloatingIPAuthorization) SetServerId(v int32)`

SetServerId sets ServerId field to given value.


### GetServerHostname

`func (o *FloatingIPAuthorization) GetServerHostname() string`

GetServerHostname returns the ServerHostname field if non-nil, zero value otherwise.

### GetServerHostnameOk

`func (o *FloatingIPAuthorization) GetServerHostnameOk() (*string, bool)`

GetServerHostnameOk returns a tuple with the ServerHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostname

`func (o *FloatingIPAuthorization) SetServerHostname(v string)`

SetServerHostname sets ServerHostname field to given value.


### GetCreatedAt

`func (o *FloatingIPAuthorization) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FloatingIPAuthorization) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FloatingIPAuthorization) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


