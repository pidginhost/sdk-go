# PatchedServerDetailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**SshPubKey** | Pointer to **string** | Public key to apply for SSH login. Applying a non-empty key regenerates cloud-init and reboots a running server. Clearing removes the key from future cloud-init data, but does not revoke keys already in the guest. | [optional] 

## Methods

### NewPatchedServerDetailRequest

`func NewPatchedServerDetailRequest() *PatchedServerDetailRequest`

NewPatchedServerDetailRequest instantiates a new PatchedServerDetailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedServerDetailRequestWithDefaults

`func NewPatchedServerDetailRequestWithDefaults() *PatchedServerDetailRequest`

NewPatchedServerDetailRequestWithDefaults instantiates a new PatchedServerDetailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *PatchedServerDetailRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PatchedServerDetailRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PatchedServerDetailRequest) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PatchedServerDetailRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPassword

`func (o *PatchedServerDetailRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatchedServerDetailRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatchedServerDetailRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatchedServerDetailRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSshPubKey

`func (o *PatchedServerDetailRequest) GetSshPubKey() string`

GetSshPubKey returns the SshPubKey field if non-nil, zero value otherwise.

### GetSshPubKeyOk

`func (o *PatchedServerDetailRequest) GetSshPubKeyOk() (*string, bool)`

GetSshPubKeyOk returns a tuple with the SshPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPubKey

`func (o *PatchedServerDetailRequest) SetSshPubKey(v string)`

SetSshPubKey sets SshPubKey field to given value.

### HasSshPubKey

`func (o *PatchedServerDetailRequest) HasSshPubKey() bool`

HasSshPubKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


