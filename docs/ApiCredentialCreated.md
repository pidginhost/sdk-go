# ApiCredentialCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | [**ApiCredential**](ApiCredential.md) |  | 
**Key** | **string** |  | [readonly] 

## Methods

### NewApiCredentialCreated

`func NewApiCredentialCreated(credential ApiCredential, key string, ) *ApiCredentialCreated`

NewApiCredentialCreated instantiates a new ApiCredentialCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCredentialCreatedWithDefaults

`func NewApiCredentialCreatedWithDefaults() *ApiCredentialCreated`

NewApiCredentialCreatedWithDefaults instantiates a new ApiCredentialCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredential

`func (o *ApiCredentialCreated) GetCredential() ApiCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ApiCredentialCreated) GetCredentialOk() (*ApiCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ApiCredentialCreated) SetCredential(v ApiCredential)`

SetCredential sets Credential field to given value.


### GetKey

`func (o *ApiCredentialCreated) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApiCredentialCreated) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApiCredentialCreated) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


