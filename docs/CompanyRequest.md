# CompanyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**CifVat** | Pointer to **string** |  | [optional] 
**Reg** | Pointer to **string** |  | [optional] 
**Iban** | Pointer to **string** |  | [optional] 
**Bank** | Pointer to **string** |  | [optional] 
**ContactName** | Pointer to **string** |  | [optional] 
**ContactEmail** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**AddressRequest**](AddressRequest.md) |  | [optional] 

## Methods

### NewCompanyRequest

`func NewCompanyRequest(name string, ) *CompanyRequest`

NewCompanyRequest instantiates a new CompanyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyRequestWithDefaults

`func NewCompanyRequestWithDefaults() *CompanyRequest`

NewCompanyRequestWithDefaults instantiates a new CompanyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CompanyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCifVat

`func (o *CompanyRequest) GetCifVat() string`

GetCifVat returns the CifVat field if non-nil, zero value otherwise.

### GetCifVatOk

`func (o *CompanyRequest) GetCifVatOk() (*string, bool)`

GetCifVatOk returns a tuple with the CifVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCifVat

`func (o *CompanyRequest) SetCifVat(v string)`

SetCifVat sets CifVat field to given value.

### HasCifVat

`func (o *CompanyRequest) HasCifVat() bool`

HasCifVat returns a boolean if a field has been set.

### GetReg

`func (o *CompanyRequest) GetReg() string`

GetReg returns the Reg field if non-nil, zero value otherwise.

### GetRegOk

`func (o *CompanyRequest) GetRegOk() (*string, bool)`

GetRegOk returns a tuple with the Reg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReg

`func (o *CompanyRequest) SetReg(v string)`

SetReg sets Reg field to given value.

### HasReg

`func (o *CompanyRequest) HasReg() bool`

HasReg returns a boolean if a field has been set.

### GetIban

`func (o *CompanyRequest) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *CompanyRequest) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *CompanyRequest) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *CompanyRequest) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetBank

`func (o *CompanyRequest) GetBank() string`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *CompanyRequest) GetBankOk() (*string, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *CompanyRequest) SetBank(v string)`

SetBank sets Bank field to given value.

### HasBank

`func (o *CompanyRequest) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetContactName

`func (o *CompanyRequest) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *CompanyRequest) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *CompanyRequest) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *CompanyRequest) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetContactEmail

`func (o *CompanyRequest) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *CompanyRequest) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *CompanyRequest) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *CompanyRequest) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetAddress

`func (o *CompanyRequest) GetAddress() AddressRequest`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CompanyRequest) GetAddressOk() (*AddressRequest, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CompanyRequest) SetAddress(v AddressRequest)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CompanyRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


