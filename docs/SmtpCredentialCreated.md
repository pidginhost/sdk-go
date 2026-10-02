# SmtpCredentialCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | [**SmtpCredential**](SmtpCredential.md) |  | 
**Password** | **string** |  | [readonly] 

## Methods

### NewSmtpCredentialCreated

`func NewSmtpCredentialCreated(credential SmtpCredential, password string, ) *SmtpCredentialCreated`

NewSmtpCredentialCreated instantiates a new SmtpCredentialCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpCredentialCreatedWithDefaults

`func NewSmtpCredentialCreatedWithDefaults() *SmtpCredentialCreated`

NewSmtpCredentialCreatedWithDefaults instantiates a new SmtpCredentialCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredential

`func (o *SmtpCredentialCreated) GetCredential() SmtpCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *SmtpCredentialCreated) GetCredentialOk() (*SmtpCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *SmtpCredentialCreated) SetCredential(v SmtpCredential)`

SetCredential sets Credential field to given value.


### GetPassword

`func (o *SmtpCredentialCreated) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SmtpCredentialCreated) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SmtpCredentialCreated) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


