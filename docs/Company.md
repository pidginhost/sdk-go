# Company

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**CifVat** | Pointer to **string** |  | [optional] 
**Reg** | Pointer to **string** |  | [optional] 
**Iban** | Pointer to **string** |  | [optional] 
**Bank** | Pointer to **string** |  | [optional] 

## Methods

### NewCompany

`func NewCompany(id int32, name string, ) *Company`

NewCompany instantiates a new Company object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyWithDefaults

`func NewCompanyWithDefaults() *Company`

NewCompanyWithDefaults instantiates a new Company object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Company) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Company) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Company) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Company) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Company) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Company) SetName(v string)`

SetName sets Name field to given value.


### GetCifVat

`func (o *Company) GetCifVat() string`

GetCifVat returns the CifVat field if non-nil, zero value otherwise.

### GetCifVatOk

`func (o *Company) GetCifVatOk() (*string, bool)`

GetCifVatOk returns a tuple with the CifVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCifVat

`func (o *Company) SetCifVat(v string)`

SetCifVat sets CifVat field to given value.

### HasCifVat

`func (o *Company) HasCifVat() bool`

HasCifVat returns a boolean if a field has been set.

### GetReg

`func (o *Company) GetReg() string`

GetReg returns the Reg field if non-nil, zero value otherwise.

### GetRegOk

`func (o *Company) GetRegOk() (*string, bool)`

GetRegOk returns a tuple with the Reg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReg

`func (o *Company) SetReg(v string)`

SetReg sets Reg field to given value.

### HasReg

`func (o *Company) HasReg() bool`

HasReg returns a boolean if a field has been set.

### GetIban

`func (o *Company) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *Company) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *Company) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *Company) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetBank

`func (o *Company) GetBank() string`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *Company) GetBankOk() (*string, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *Company) SetBank(v string)`

SetBank sets Bank field to given value.

### HasBank

`func (o *Company) HasBank() bool`

HasBank returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


