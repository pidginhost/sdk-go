# Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Domain** | **string** |  | [readonly] 
**Tld** | [**TLD**](TLD.md) |  | [readonly] 
**Idna** | **bool** | Domain name is encoded with IDN | [readonly] 
**Nameservers** | Pointer to **string** | List of 2-5 name-servers separated by comma. | [optional] 
**ExpirationDate** | **string** |  | [readonly] 
**RegistrationDate** | **NullableString** |  | [readonly] 
**Service** | [**Service**](Service.md) |  | [readonly] 
**IdnaName** | **string** |  | [readonly] 
**MaxRenewYears** | **int32** |  | [readonly] 
**ServiceStatus** | **string** | Service status | [readonly] 
**Contacts** | **interface{}** |  | [readonly] 

## Methods

### NewDomain

`func NewDomain(id int32, domain string, tld TLD, idna bool, expirationDate string, registrationDate NullableString, service Service, idnaName string, maxRenewYears int32, serviceStatus string, contacts interface{}, ) *Domain`

NewDomain instantiates a new Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWithDefaults

`func NewDomainWithDefaults() *Domain`

NewDomainWithDefaults instantiates a new Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Domain) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Domain) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Domain) SetId(v int32)`

SetId sets Id field to given value.


### GetDomain

`func (o *Domain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Domain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Domain) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetTld

`func (o *Domain) GetTld() TLD`

GetTld returns the Tld field if non-nil, zero value otherwise.

### GetTldOk

`func (o *Domain) GetTldOk() (*TLD, bool)`

GetTldOk returns a tuple with the Tld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTld

`func (o *Domain) SetTld(v TLD)`

SetTld sets Tld field to given value.


### GetIdna

`func (o *Domain) GetIdna() bool`

GetIdna returns the Idna field if non-nil, zero value otherwise.

### GetIdnaOk

`func (o *Domain) GetIdnaOk() (*bool, bool)`

GetIdnaOk returns a tuple with the Idna field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdna

`func (o *Domain) SetIdna(v bool)`

SetIdna sets Idna field to given value.


### GetNameservers

`func (o *Domain) GetNameservers() string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *Domain) GetNameserversOk() (*string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *Domain) SetNameservers(v string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *Domain) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetExpirationDate

`func (o *Domain) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *Domain) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *Domain) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.


### GetRegistrationDate

`func (o *Domain) GetRegistrationDate() string`

GetRegistrationDate returns the RegistrationDate field if non-nil, zero value otherwise.

### GetRegistrationDateOk

`func (o *Domain) GetRegistrationDateOk() (*string, bool)`

GetRegistrationDateOk returns a tuple with the RegistrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationDate

`func (o *Domain) SetRegistrationDate(v string)`

SetRegistrationDate sets RegistrationDate field to given value.


### SetRegistrationDateNil

`func (o *Domain) SetRegistrationDateNil(b bool)`

 SetRegistrationDateNil sets the value for RegistrationDate to be an explicit nil

### UnsetRegistrationDate
`func (o *Domain) UnsetRegistrationDate()`

UnsetRegistrationDate ensures that no value is present for RegistrationDate, not even an explicit nil
### GetService

`func (o *Domain) GetService() Service`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *Domain) GetServiceOk() (*Service, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *Domain) SetService(v Service)`

SetService sets Service field to given value.


### GetIdnaName

`func (o *Domain) GetIdnaName() string`

GetIdnaName returns the IdnaName field if non-nil, zero value otherwise.

### GetIdnaNameOk

`func (o *Domain) GetIdnaNameOk() (*string, bool)`

GetIdnaNameOk returns a tuple with the IdnaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdnaName

`func (o *Domain) SetIdnaName(v string)`

SetIdnaName sets IdnaName field to given value.


### GetMaxRenewYears

`func (o *Domain) GetMaxRenewYears() int32`

GetMaxRenewYears returns the MaxRenewYears field if non-nil, zero value otherwise.

### GetMaxRenewYearsOk

`func (o *Domain) GetMaxRenewYearsOk() (*int32, bool)`

GetMaxRenewYearsOk returns a tuple with the MaxRenewYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRenewYears

`func (o *Domain) SetMaxRenewYears(v int32)`

SetMaxRenewYears sets MaxRenewYears field to given value.


### GetServiceStatus

`func (o *Domain) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *Domain) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *Domain) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetContacts

`func (o *Domain) GetContacts() interface{}`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *Domain) GetContactsOk() (*interface{}, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *Domain) SetContacts(v interface{})`

SetContacts sets Contacts field to given value.


### SetContactsNil

`func (o *Domain) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *Domain) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


