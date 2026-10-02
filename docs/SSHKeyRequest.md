# SSHKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** |  | [optional] 
**Key** | **string** |  | 

## Methods

### NewSSHKeyRequest

`func NewSSHKeyRequest(key string, ) *SSHKeyRequest`

NewSSHKeyRequest instantiates a new SSHKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSHKeyRequestWithDefaults

`func NewSSHKeyRequestWithDefaults() *SSHKeyRequest`

NewSSHKeyRequestWithDefaults instantiates a new SSHKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *SSHKeyRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *SSHKeyRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *SSHKeyRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *SSHKeyRequest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetKey

`func (o *SSHKeyRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SSHKeyRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SSHKeyRequest) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


