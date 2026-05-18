# DeactivateFreeDNS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | Domain name or primary key of the domain to deactivate. | 
**Source** | [**SourceEnum**](SourceEnum.md) |  | 

## Methods

### NewDeactivateFreeDNS

`func NewDeactivateFreeDNS(domain string, source SourceEnum, ) *DeactivateFreeDNS`

NewDeactivateFreeDNS instantiates a new DeactivateFreeDNS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeactivateFreeDNSWithDefaults

`func NewDeactivateFreeDNSWithDefaults() *DeactivateFreeDNS`

NewDeactivateFreeDNSWithDefaults instantiates a new DeactivateFreeDNS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DeactivateFreeDNS) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DeactivateFreeDNS) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DeactivateFreeDNS) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetSource

`func (o *DeactivateFreeDNS) GetSource() SourceEnum`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DeactivateFreeDNS) GetSourceOk() (*SourceEnum, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DeactivateFreeDNS) SetSource(v SourceEnum)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


