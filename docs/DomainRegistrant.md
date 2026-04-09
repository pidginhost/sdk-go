# DomainRegistrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
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

### NewDomainRegistrant

`func NewDomainRegistrant(id int32, firstName string, lastName string, address string, city string, region string, postalCode string, country CountryEnum, email string, phone string, ) *DomainRegistrant`

NewDomainRegistrant instantiates a new DomainRegistrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainRegistrantWithDefaults

`func NewDomainRegistrantWithDefaults() *DomainRegistrant`

NewDomainRegistrantWithDefaults instantiates a new DomainRegistrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DomainRegistrant) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainRegistrant) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainRegistrant) SetId(v int32)`

SetId sets Id field to given value.


### GetFirstName

`func (o *DomainRegistrant) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *DomainRegistrant) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *DomainRegistrant) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *DomainRegistrant) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *DomainRegistrant) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *DomainRegistrant) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetCompany

`func (o *DomainRegistrant) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *DomainRegistrant) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *DomainRegistrant) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *DomainRegistrant) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *DomainRegistrant) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *DomainRegistrant) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetAddress

`func (o *DomainRegistrant) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DomainRegistrant) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DomainRegistrant) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCity

`func (o *DomainRegistrant) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *DomainRegistrant) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *DomainRegistrant) SetCity(v string)`

SetCity sets City field to given value.


### GetRegion

`func (o *DomainRegistrant) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DomainRegistrant) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DomainRegistrant) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetPostalCode

`func (o *DomainRegistrant) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *DomainRegistrant) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *DomainRegistrant) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetCountry

`func (o *DomainRegistrant) GetCountry() CountryEnum`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *DomainRegistrant) GetCountryOk() (*CountryEnum, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *DomainRegistrant) SetCountry(v CountryEnum)`

SetCountry sets Country field to given value.


### GetEmail

`func (o *DomainRegistrant) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DomainRegistrant) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DomainRegistrant) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhone

`func (o *DomainRegistrant) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *DomainRegistrant) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *DomainRegistrant) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetCifCnp

`func (o *DomainRegistrant) GetCifCnp() string`

GetCifCnp returns the CifCnp field if non-nil, zero value otherwise.

### GetCifCnpOk

`func (o *DomainRegistrant) GetCifCnpOk() (*string, bool)`

GetCifCnpOk returns a tuple with the CifCnp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCifCnp

`func (o *DomainRegistrant) SetCifCnp(v string)`

SetCifCnp sets CifCnp field to given value.

### HasCifCnp

`func (o *DomainRegistrant) HasCifCnp() bool`

HasCifCnp returns a boolean if a field has been set.

### SetCifCnpNil

`func (o *DomainRegistrant) SetCifCnpNil(b bool)`

 SetCifCnpNil sets the value for CifCnp to be an explicit nil

### UnsetCifCnp
`func (o *DomainRegistrant) UnsetCifCnp()`

UnsetCifCnp ensures that no value is present for CifCnp, not even an explicit nil
### GetRegCom

`func (o *DomainRegistrant) GetRegCom() string`

GetRegCom returns the RegCom field if non-nil, zero value otherwise.

### GetRegComOk

`func (o *DomainRegistrant) GetRegComOk() (*string, bool)`

GetRegComOk returns a tuple with the RegCom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegCom

`func (o *DomainRegistrant) SetRegCom(v string)`

SetRegCom sets RegCom field to given value.

### HasRegCom

`func (o *DomainRegistrant) HasRegCom() bool`

HasRegCom returns a boolean if a field has been set.

### SetRegComNil

`func (o *DomainRegistrant) SetRegComNil(b bool)`

 SetRegComNil sets the value for RegCom to be an explicit nil

### UnsetRegCom
`func (o *DomainRegistrant) UnsetRegCom()`

UnsetRegCom ensures that no value is present for RegCom, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


