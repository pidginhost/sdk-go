# ServerDetailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**SshPubKey** | Pointer to **string** | Public key to apply for SSH login. Applying a non-empty key regenerates cloud-init and reboots a running server. Clearing removes the key from future cloud-init data, but does not revoke keys already in the guest. | [optional] 

## Methods

### NewServerDetailRequest

`func NewServerDetailRequest() *ServerDetailRequest`

NewServerDetailRequest instantiates a new ServerDetailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerDetailRequestWithDefaults

`func NewServerDetailRequestWithDefaults() *ServerDetailRequest`

NewServerDetailRequestWithDefaults instantiates a new ServerDetailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *ServerDetailRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ServerDetailRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ServerDetailRequest) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ServerDetailRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPassword

`func (o *ServerDetailRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ServerDetailRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ServerDetailRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ServerDetailRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSshPubKey

`func (o *ServerDetailRequest) GetSshPubKey() string`

GetSshPubKey returns the SshPubKey field if non-nil, zero value otherwise.

### GetSshPubKeyOk

`func (o *ServerDetailRequest) GetSshPubKeyOk() (*string, bool)`

GetSshPubKeyOk returns a tuple with the SshPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPubKey

`func (o *ServerDetailRequest) SetSshPubKey(v string)`

SetSshPubKey sets SshPubKey field to given value.

### HasSshPubKey

`func (o *ServerDetailRequest) HasSshPubKey() bool`

HasSshPubKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


