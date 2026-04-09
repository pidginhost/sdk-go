# DomainCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | Domain with tld, ex: example.com | 
**Nameservers** | Pointer to **string** | List of 2-5 name-servers separated by comma. | [optional] 
**Years** | Pointer to **int32** |  | [optional] [default to 1]

## Methods

### NewDomainCreate

`func NewDomainCreate(domain string, ) *DomainCreate`

NewDomainCreate instantiates a new DomainCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainCreateWithDefaults

`func NewDomainCreateWithDefaults() *DomainCreate`

NewDomainCreateWithDefaults instantiates a new DomainCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DomainCreate) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainCreate) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainCreate) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetNameservers

`func (o *DomainCreate) GetNameservers() string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *DomainCreate) GetNameserversOk() (*string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *DomainCreate) SetNameservers(v string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *DomainCreate) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetYears

`func (o *DomainCreate) GetYears() int32`

GetYears returns the Years field if non-nil, zero value otherwise.

### GetYearsOk

`func (o *DomainCreate) GetYearsOk() (*int32, bool)`

GetYearsOk returns a tuple with the Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYears

`func (o *DomainCreate) SetYears(v int32)`

SetYears sets Years field to given value.

### HasYears

`func (o *DomainCreate) HasYears() bool`

HasYears returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


