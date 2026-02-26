# SSHKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Alias** | Pointer to **string** |  | [optional] 
**Fingerprint** | **string** |  | [readonly] 
**Key** | **string** |  | [readonly] 

## Methods

### NewSSHKey

`func NewSSHKey(id int32, fingerprint string, key string, ) *SSHKey`

NewSSHKey instantiates a new SSHKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSHKeyWithDefaults

`func NewSSHKeyWithDefaults() *SSHKey`

NewSSHKeyWithDefaults instantiates a new SSHKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SSHKey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SSHKey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SSHKey) SetId(v int32)`

SetId sets Id field to given value.


### GetAlias

`func (o *SSHKey) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *SSHKey) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *SSHKey) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *SSHKey) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetFingerprint

`func (o *SSHKey) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *SSHKey) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *SSHKey) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.


### GetKey

`func (o *SSHKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SSHKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SSHKey) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


