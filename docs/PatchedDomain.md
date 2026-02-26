# PatchedDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Domain** | Pointer to **string** |  | [optional] [readonly] 
**Tld** | Pointer to [**TLD**](TLD.md) |  | [optional] [readonly] 
**Idna** | Pointer to **bool** | Domain name is encoded with IDN | [optional] [readonly] 
**Nameservers** | Pointer to **string** | List of 2-5 name-servers separated by comma. | [optional] 
**ExpirationDate** | Pointer to **string** |  | [optional] [readonly] 
**RegistrationDate** | Pointer to **NullableString** |  | [optional] [readonly] 
**Service** | Pointer to [**Service**](Service.md) |  | [optional] [readonly] 
**IdnaName** | Pointer to **string** |  | [optional] [readonly] 
**MaxRenewYears** | Pointer to **int32** |  | [optional] [readonly] 
**ServiceStatus** | Pointer to **string** | Service status | [optional] [readonly] 
**Contacts** | Pointer to **interface{}** |  | [optional] [readonly] 

## Methods

### NewPatchedDomain

`func NewPatchedDomain() *PatchedDomain`

NewPatchedDomain instantiates a new PatchedDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDomainWithDefaults

`func NewPatchedDomainWithDefaults() *PatchedDomain`

NewPatchedDomainWithDefaults instantiates a new PatchedDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedDomain) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedDomain) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedDomain) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedDomain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDomain

`func (o *PatchedDomain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PatchedDomain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PatchedDomain) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PatchedDomain) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetTld

`func (o *PatchedDomain) GetTld() TLD`

GetTld returns the Tld field if non-nil, zero value otherwise.

### GetTldOk

`func (o *PatchedDomain) GetTldOk() (*TLD, bool)`

GetTldOk returns a tuple with the Tld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTld

`func (o *PatchedDomain) SetTld(v TLD)`

SetTld sets Tld field to given value.

### HasTld

`func (o *PatchedDomain) HasTld() bool`

HasTld returns a boolean if a field has been set.

### GetIdna

`func (o *PatchedDomain) GetIdna() bool`

GetIdna returns the Idna field if non-nil, zero value otherwise.

### GetIdnaOk

`func (o *PatchedDomain) GetIdnaOk() (*bool, bool)`

GetIdnaOk returns a tuple with the Idna field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdna

`func (o *PatchedDomain) SetIdna(v bool)`

SetIdna sets Idna field to given value.

### HasIdna

`func (o *PatchedDomain) HasIdna() bool`

HasIdna returns a boolean if a field has been set.

### GetNameservers

`func (o *PatchedDomain) GetNameservers() string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *PatchedDomain) GetNameserversOk() (*string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *PatchedDomain) SetNameservers(v string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *PatchedDomain) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetExpirationDate

`func (o *PatchedDomain) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *PatchedDomain) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *PatchedDomain) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *PatchedDomain) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetRegistrationDate

`func (o *PatchedDomain) GetRegistrationDate() string`

GetRegistrationDate returns the RegistrationDate field if non-nil, zero value otherwise.

### GetRegistrationDateOk

`func (o *PatchedDomain) GetRegistrationDateOk() (*string, bool)`

GetRegistrationDateOk returns a tuple with the RegistrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationDate

`func (o *PatchedDomain) SetRegistrationDate(v string)`

SetRegistrationDate sets RegistrationDate field to given value.

### HasRegistrationDate

`func (o *PatchedDomain) HasRegistrationDate() bool`

HasRegistrationDate returns a boolean if a field has been set.

### SetRegistrationDateNil

`func (o *PatchedDomain) SetRegistrationDateNil(b bool)`

 SetRegistrationDateNil sets the value for RegistrationDate to be an explicit nil

### UnsetRegistrationDate
`func (o *PatchedDomain) UnsetRegistrationDate()`

UnsetRegistrationDate ensures that no value is present for RegistrationDate, not even an explicit nil
### GetService

`func (o *PatchedDomain) GetService() Service`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *PatchedDomain) GetServiceOk() (*Service, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *PatchedDomain) SetService(v Service)`

SetService sets Service field to given value.

### HasService

`func (o *PatchedDomain) HasService() bool`

HasService returns a boolean if a field has been set.

### GetIdnaName

`func (o *PatchedDomain) GetIdnaName() string`

GetIdnaName returns the IdnaName field if non-nil, zero value otherwise.

### GetIdnaNameOk

`func (o *PatchedDomain) GetIdnaNameOk() (*string, bool)`

GetIdnaNameOk returns a tuple with the IdnaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdnaName

`func (o *PatchedDomain) SetIdnaName(v string)`

SetIdnaName sets IdnaName field to given value.

### HasIdnaName

`func (o *PatchedDomain) HasIdnaName() bool`

HasIdnaName returns a boolean if a field has been set.

### GetMaxRenewYears

`func (o *PatchedDomain) GetMaxRenewYears() int32`

GetMaxRenewYears returns the MaxRenewYears field if non-nil, zero value otherwise.

### GetMaxRenewYearsOk

`func (o *PatchedDomain) GetMaxRenewYearsOk() (*int32, bool)`

GetMaxRenewYearsOk returns a tuple with the MaxRenewYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRenewYears

`func (o *PatchedDomain) SetMaxRenewYears(v int32)`

SetMaxRenewYears sets MaxRenewYears field to given value.

### HasMaxRenewYears

`func (o *PatchedDomain) HasMaxRenewYears() bool`

HasMaxRenewYears returns a boolean if a field has been set.

### GetServiceStatus

`func (o *PatchedDomain) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *PatchedDomain) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *PatchedDomain) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.

### HasServiceStatus

`func (o *PatchedDomain) HasServiceStatus() bool`

HasServiceStatus returns a boolean if a field has been set.

### GetContacts

`func (o *PatchedDomain) GetContacts() interface{}`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *PatchedDomain) GetContactsOk() (*interface{}, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *PatchedDomain) SetContacts(v interface{})`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *PatchedDomain) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContactsNil

`func (o *PatchedDomain) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *PatchedDomain) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


