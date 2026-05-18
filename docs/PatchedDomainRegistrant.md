# PatchedDomainRegistrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
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

### NewPatchedDomainRegistrant

`func NewPatchedDomainRegistrant() *PatchedDomainRegistrant`

NewPatchedDomainRegistrant instantiates a new PatchedDomainRegistrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDomainRegistrantWithDefaults

`func NewPatchedDomainRegistrantWithDefaults() *PatchedDomainRegistrant`

NewPatchedDomainRegistrantWithDefaults instantiates a new PatchedDomainRegistrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedDomainRegistrant) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedDomainRegistrant) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedDomainRegistrant) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedDomainRegistrant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstName

`func (o *PatchedDomainRegistrant) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PatchedDomainRegistrant) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PatchedDomainRegistrant) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PatchedDomainRegistrant) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *PatchedDomainRegistrant) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PatchedDomainRegistrant) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PatchedDomainRegistrant) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PatchedDomainRegistrant) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetCompany

`func (o *PatchedDomainRegistrant) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PatchedDomainRegistrant) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PatchedDomainRegistrant) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PatchedDomainRegistrant) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *PatchedDomainRegistrant) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *PatchedDomainRegistrant) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetAddress

`func (o *PatchedDomainRegistrant) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PatchedDomainRegistrant) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PatchedDomainRegistrant) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PatchedDomainRegistrant) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCity

`func (o *PatchedDomainRegistrant) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PatchedDomainRegistrant) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PatchedDomainRegistrant) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *PatchedDomainRegistrant) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetRegion

`func (o *PatchedDomainRegistrant) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PatchedDomainRegistrant) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PatchedDomainRegistrant) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PatchedDomainRegistrant) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetPostalCode

`func (o *PatchedDomainRegistrant) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PatchedDomainRegistrant) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PatchedDomainRegistrant) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *PatchedDomainRegistrant) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountry

`func (o *PatchedDomainRegistrant) GetCountry() CountryEnum`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PatchedDomainRegistrant) GetCountryOk() (*CountryEnum, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PatchedDomainRegistrant) SetCountry(v CountryEnum)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PatchedDomainRegistrant) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *PatchedDomainRegistrant) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PatchedDomainRegistrant) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PatchedDomainRegistrant) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PatchedDomainRegistrant) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *PatchedDomainRegistrant) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PatchedDomainRegistrant) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PatchedDomainRegistrant) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *PatchedDomainRegistrant) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCifCnp

`func (o *PatchedDomainRegistrant) GetCifCnp() string`

GetCifCnp returns the CifCnp field if non-nil, zero value otherwise.

### GetCifCnpOk

`func (o *PatchedDomainRegistrant) GetCifCnpOk() (*string, bool)`

GetCifCnpOk returns a tuple with the CifCnp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCifCnp

`func (o *PatchedDomainRegistrant) SetCifCnp(v string)`

SetCifCnp sets CifCnp field to given value.

### HasCifCnp

`func (o *PatchedDomainRegistrant) HasCifCnp() bool`

HasCifCnp returns a boolean if a field has been set.

### SetCifCnpNil

`func (o *PatchedDomainRegistrant) SetCifCnpNil(b bool)`

 SetCifCnpNil sets the value for CifCnp to be an explicit nil

### UnsetCifCnp
`func (o *PatchedDomainRegistrant) UnsetCifCnp()`

UnsetCifCnp ensures that no value is present for CifCnp, not even an explicit nil
### GetRegCom

`func (o *PatchedDomainRegistrant) GetRegCom() string`

GetRegCom returns the RegCom field if non-nil, zero value otherwise.

### GetRegComOk

`func (o *PatchedDomainRegistrant) GetRegComOk() (*string, bool)`

GetRegComOk returns a tuple with the RegCom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegCom

`func (o *PatchedDomainRegistrant) SetRegCom(v string)`

SetRegCom sets RegCom field to given value.

### HasRegCom

`func (o *PatchedDomainRegistrant) HasRegCom() bool`

HasRegCom returns a boolean if a field has been set.

### SetRegComNil

`func (o *PatchedDomainRegistrant) SetRegComNil(b bool)`

 SetRegComNil sets the value for RegCom to be an explicit nil

### UnsetRegCom
`func (o *PatchedDomainRegistrant) UnsetRegCom()`

UnsetRegCom ensures that no value is present for RegCom, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


