# PatchedSSHKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Alias** | Pointer to **string** |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] [readonly] 
**Key** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPatchedSSHKey

`func NewPatchedSSHKey() *PatchedSSHKey`

NewPatchedSSHKey instantiates a new PatchedSSHKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSSHKeyWithDefaults

`func NewPatchedSSHKeyWithDefaults() *PatchedSSHKey`

NewPatchedSSHKeyWithDefaults instantiates a new PatchedSSHKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedSSHKey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedSSHKey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedSSHKey) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedSSHKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlias

`func (o *PatchedSSHKey) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *PatchedSSHKey) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *PatchedSSHKey) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *PatchedSSHKey) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetFingerprint

`func (o *PatchedSSHKey) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *PatchedSSHKey) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *PatchedSSHKey) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *PatchedSSHKey) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetKey

`func (o *PatchedSSHKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PatchedSSHKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PatchedSSHKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PatchedSSHKey) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


