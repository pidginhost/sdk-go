# DomainRegistrantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**Company** | Pointer to **NullableString** |  | [optional] 
**Address** | **string** |  | 
**City** | **string** |  | 
**Region** | **string** |  | 
**PostalCode** | **string** |  | 
**Country** | [**CountryEnum**](CountryEnum.md) |  | 
**Email** | **string** |  | 
**Phone** | **string** |  | 
**CifCnp** | Pointer to **NullableString** |  | [optional] 
**RegCom** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDomainRegistrantRequest

`func NewDomainRegistrantRequest(firstName string, lastName string, address string, city string, region string, postalCode string, country CountryEnum, email string, phone string, ) *DomainRegistrantRequest`

NewDomainRegistrantRequest instantiates a new DomainRegistrantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainRegistrantRequestWithDefaults

`func NewDomainRegistrantRequestWithDefaults() *DomainRegistrantRequest`

NewDomainRegistrantRequestWithDefaults instantiates a new DomainRegistrantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *DomainRegistrantRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *DomainRegistrantRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *DomainRegistrantRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *DomainRegistrantRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *DomainRegistrantRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *DomainRegistrantRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetCompany

`func (o *DomainRegistrantRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *DomainRegistrantRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *DomainRegistrantRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *DomainRegistrantRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *DomainRegistrantRequest) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *DomainRegistrantRequest) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetAddress

`func (o *DomainRegistrantRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DomainRegistrantRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DomainRegistrantRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCity

`func (o *DomainRegistrantRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *DomainRegistrantRequest) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *DomainRegistrantRequest) SetCity(v string)`

SetCity sets City field to given value.


### GetRegion

`func (o *DomainRegistrantRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DomainRegistrantRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DomainRegistrantRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetPostalCode

`func (o *DomainRegistrantRequest) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *DomainRegistrantRequest) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *DomainRegistrantRequest) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetCountry

`func (o *DomainRegistrantRequest) GetCountry() CountryEnum`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *DomainRegistrantRequest) GetCountryOk() (*CountryEnum, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *DomainRegistrantRequest) SetCountry(v CountryEnum)`

SetCountry sets Country field to given value.


### GetEmail

`func (o *DomainRegistrantRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DomainRegistrantRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DomainRegistrantRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhone

`func (o *DomainRegistrantRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *DomainRegistrantRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *DomainRegistrantRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetCifCnp

`func (o *DomainRegistrantRequest) GetCifCnp() string`

GetCifCnp returns the CifCnp field if non-nil, zero value otherwise.

### GetCifCnpOk

`func (o *DomainRegistrantRequest) GetCifCnpOk() (*string, bool)`

GetCifCnpOk returns a tuple with the CifCnp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCifCnp

`func (o *DomainRegistrantRequest) SetCifCnp(v string)`

SetCifCnp sets CifCnp field to given value.

### HasCifCnp

`func (o *DomainRegistrantRequest) HasCifCnp() bool`

HasCifCnp returns a boolean if a field has been set.

### SetCifCnpNil

`func (o *DomainRegistrantRequest) SetCifCnpNil(b bool)`

 SetCifCnpNil sets the value for CifCnp to be an explicit nil

### UnsetCifCnp
`func (o *DomainRegistrantRequest) UnsetCifCnp()`

UnsetCifCnp ensures that no value is present for CifCnp, not even an explicit nil
### GetRegCom

`func (o *DomainRegistrantRequest) GetRegCom() string`

GetRegCom returns the RegCom field if non-nil, zero value otherwise.

### GetRegComOk

`func (o *DomainRegistrantRequest) GetRegComOk() (*string, bool)`

GetRegComOk returns a tuple with the RegCom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegCom

`func (o *DomainRegistrantRequest) SetRegCom(v string)`

SetRegCom sets RegCom field to given value.

### HasRegCom

`func (o *DomainRegistrantRequest) HasRegCom() bool`

HasRegCom returns a boolean if a field has been set.

### SetRegComNil

`func (o *DomainRegistrantRequest) SetRegComNil(b bool)`

 SetRegComNil sets the value for RegCom to be an explicit nil

### UnsetRegCom
`func (o *DomainRegistrantRequest) UnsetRegCom()`

UnsetRegCom ensures that no value is present for RegCom, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


