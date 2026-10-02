# DomainCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | Domain with tld, ex: example.com | 
**Nameservers** | Pointer to **string** | List of 2-5 name-servers separated by comma. | [optional] 
**Years** | Pointer to **int32** |  | [optional] [default to 1]

## Methods

### NewDomainCreateRequest

`func NewDomainCreateRequest(domain string, ) *DomainCreateRequest`

NewDomainCreateRequest instantiates a new DomainCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainCreateRequestWithDefaults

`func NewDomainCreateRequestWithDefaults() *DomainCreateRequest`

NewDomainCreateRequestWithDefaults instantiates a new DomainCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DomainCreateRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainCreateRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainCreateRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetNameservers

`func (o *DomainCreateRequest) GetNameservers() string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *DomainCreateRequest) GetNameserversOk() (*string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *DomainCreateRequest) SetNameservers(v string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *DomainCreateRequest) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetYears

`func (o *DomainCreateRequest) GetYears() int32`

GetYears returns the Years field if non-nil, zero value otherwise.

### GetYearsOk

`func (o *DomainCreateRequest) GetYearsOk() (*int32, bool)`

GetYearsOk returns a tuple with the Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYears

`func (o *DomainCreateRequest) SetYears(v int32)`

SetYears sets Years field to given value.

### HasYears

`func (o *DomainCreateRequest) HasYears() bool`

HasYears returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


