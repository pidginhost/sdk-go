# PatchedDomainRegistrantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **NullableString** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] 
**Country** | Pointer to [**CountryEnum**](CountryEnum.md) |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**CifCnp** | Pointer to **NullableString** |  | [optional] 
**RegCom** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPatchedDomainRegistrantRequest

`func NewPatchedDomainRegistrantRequest() *PatchedDomainRegistrantRequest`

NewPatchedDomainRegistrantRequest instantiates a new PatchedDomainRegistrantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDomainRegistrantRequestWithDefaults

`func NewPatchedDomainRegistrantRequestWithDefaults() *PatchedDomainRegistrantRequest`

NewPatchedDomainRegistrantRequestWithDefaults instantiates a new PatchedDomainRegistrantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *PatchedDomainRegistrantRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PatchedDomainRegistrantRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PatchedDomainRegistrantRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PatchedDomainRegistrantRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *PatchedDomainRegistrantRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PatchedDomainRegistrantRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PatchedDomainRegistrantRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PatchedDomainRegistrantRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetCompany

`func (o *PatchedDomainRegistrantRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PatchedDomainRegistrantRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PatchedDomainRegistrantRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PatchedDomainRegistrantRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *PatchedDomainRegistrantRequest) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *PatchedDomainRegistrantRequest) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetAddress

`func (o *PatchedDomainRegistrantRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PatchedDomainRegistrantRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PatchedDomainRegistrantRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PatchedDomainRegistrantRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCity

`func (o *PatchedDomainRegistrantRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PatchedDomainRegistrantRequest) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PatchedDomainRegistrantRequest) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *PatchedDomainRegistrantRequest) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetRegion

`func (o *PatchedDomainRegistrantRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PatchedDomainRegistrantRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PatchedDomainRegistrantRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PatchedDomainRegistrantRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetPostalCode

`func (o *PatchedDomainRegistrantRequest) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PatchedDomainRegistrantRequest) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PatchedDomainRegistrantRequest) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *PatchedDomainRegistrantRequest) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountry

`func (o *PatchedDomainRegistrantRequest) GetCountry() CountryEnum`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PatchedDomainRegistrantRequest) GetCountryOk() (*CountryEnum, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PatchedDomainRegistrantRequest) SetCountry(v CountryEnum)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PatchedDomainRegistrantRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *PatchedDomainRegistrantRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PatchedDomainRegistrantRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PatchedDomainRegistrantRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PatchedDomainRegistrantRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *PatchedDomainRegistrantRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PatchedDomainRegistrantRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PatchedDomainRegistrantRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *PatchedDomainRegistrantRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCifCnp

`func (o *PatchedDomainRegistrantRequest) GetCifCnp() string`

GetCifCnp returns the CifCnp field if non-nil, zero value otherwise.

### GetCifCnpOk

`func (o *PatchedDomainRegistrantRequest) GetCifCnpOk() (*string, bool)`

GetCifCnpOk returns a tuple with the CifCnp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCifCnp

`func (o *PatchedDomainRegistrantRequest) SetCifCnp(v string)`

SetCifCnp sets CifCnp field to given value.

### HasCifCnp

`func (o *PatchedDomainRegistrantRequest) HasCifCnp() bool`

HasCifCnp returns a boolean if a field has been set.

### SetCifCnpNil

`func (o *PatchedDomainRegistrantRequest) SetCifCnpNil(b bool)`

 SetCifCnpNil sets the value for CifCnp to be an explicit nil

### UnsetCifCnp
`func (o *PatchedDomainRegistrantRequest) UnsetCifCnp()`

UnsetCifCnp ensures that no value is present for CifCnp, not even an explicit nil
### GetRegCom

`func (o *PatchedDomainRegistrantRequest) GetRegCom() string`

GetRegCom returns the RegCom field if non-nil, zero value otherwise.

### GetRegComOk

`func (o *PatchedDomainRegistrantRequest) GetRegComOk() (*string, bool)`

GetRegComOk returns a tuple with the RegCom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegCom

`func (o *PatchedDomainRegistrantRequest) SetRegCom(v string)`

SetRegCom sets RegCom field to given value.

### HasRegCom

`func (o *PatchedDomainRegistrantRequest) HasRegCom() bool`

HasRegCom returns a boolean if a field has been set.

### SetRegComNil

`func (o *PatchedDomainRegistrantRequest) SetRegComNil(b bool)`

 SetRegComNil sets the value for RegCom to be an explicit nil

### UnsetRegCom
`func (o *PatchedDomainRegistrantRequest) UnsetRegCom()`

UnsetRegCom ensures that no value is present for RegCom, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


