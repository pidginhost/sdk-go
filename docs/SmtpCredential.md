# SmtpCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Label** | **string** |  | [readonly] 
**Username** | **string** |  | [readonly] 
**Active** | **bool** |  | [readonly] 
**CreatedAt** | **string** |  | [readonly] 
**RevokedAt** | **NullableString** |  | [readonly] 

## Methods

### NewSmtpCredential

`func NewSmtpCredential(id int32, label string, username string, active bool, createdAt string, revokedAt NullableString, ) *SmtpCredential`

NewSmtpCredential instantiates a new SmtpCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpCredentialWithDefaults

`func NewSmtpCredentialWithDefaults() *SmtpCredential`

NewSmtpCredentialWithDefaults instantiates a new SmtpCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SmtpCredential) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SmtpCredential) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SmtpCredential) SetId(v int32)`

SetId sets Id field to given value.


### GetLabel

`func (o *SmtpCredential) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SmtpCredential) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SmtpCredential) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetUsername

`func (o *SmtpCredential) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SmtpCredential) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SmtpCredential) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetActive

`func (o *SmtpCredential) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SmtpCredential) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SmtpCredential) SetActive(v bool)`

SetActive sets Active field to given value.


### GetCreatedAt

`func (o *SmtpCredential) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SmtpCredential) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SmtpCredential) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetRevokedAt

`func (o *SmtpCredential) GetRevokedAt() string`

GetRevokedAt returns the RevokedAt field if non-nil, zero value otherwise.

### GetRevokedAtOk

`func (o *SmtpCredential) GetRevokedAtOk() (*string, bool)`

GetRevokedAtOk returns a tuple with the RevokedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAt

`func (o *SmtpCredential) SetRevokedAt(v string)`

SetRevokedAt sets RevokedAt field to given value.


### SetRevokedAtNil

`func (o *SmtpCredential) SetRevokedAtNil(b bool)`

 SetRevokedAtNil sets the value for RevokedAt to be an explicit nil

### UnsetRevokedAt
`func (o *SmtpCredential) UnsetRevokedAt()`

UnsetRevokedAt ensures that no value is present for RevokedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


