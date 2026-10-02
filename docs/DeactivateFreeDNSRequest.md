# DeactivateFreeDNSRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | Domain name or primary key of the domain to deactivate. | 
**Source** | [**SourceEnum**](SourceEnum.md) |  | 

## Methods

### NewDeactivateFreeDNSRequest

`func NewDeactivateFreeDNSRequest(domain string, source SourceEnum, ) *DeactivateFreeDNSRequest`

NewDeactivateFreeDNSRequest instantiates a new DeactivateFreeDNSRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeactivateFreeDNSRequestWithDefaults

`func NewDeactivateFreeDNSRequestWithDefaults() *DeactivateFreeDNSRequest`

NewDeactivateFreeDNSRequestWithDefaults instantiates a new DeactivateFreeDNSRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DeactivateFreeDNSRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DeactivateFreeDNSRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DeactivateFreeDNSRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetSource

`func (o *DeactivateFreeDNSRequest) GetSource() SourceEnum`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DeactivateFreeDNSRequest) GetSourceOk() (*SourceEnum, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DeactivateFreeDNSRequest) SetSource(v SourceEnum)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


