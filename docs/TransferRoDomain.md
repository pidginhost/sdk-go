# TransferRoDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | Domain with tld, ex: example.com | 
**AuthCode** | **string** | Auth code | 

## Methods

### NewTransferRoDomain

`func NewTransferRoDomain(domain string, authCode string, ) *TransferRoDomain`

NewTransferRoDomain instantiates a new TransferRoDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferRoDomainWithDefaults

`func NewTransferRoDomainWithDefaults() *TransferRoDomain`

NewTransferRoDomainWithDefaults instantiates a new TransferRoDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *TransferRoDomain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *TransferRoDomain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *TransferRoDomain) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetAuthCode

`func (o *TransferRoDomain) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *TransferRoDomain) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *TransferRoDomain) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


