# ActivateFreeDNSRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | Domain name or primary key of the domain to activate. | 
**Source** | [**SourceEnum**](SourceEnum.md) | &#39;internal&#39; for domains purchased on PidginHost, &#39;external&#39; for user-added domains.  * &#x60;internal&#x60; - Internal * &#x60;external&#x60; - External | 
**Ip** | **string** | IPv4 address to use as the default A record for the zone. | 

## Methods

### NewActivateFreeDNSRequest

`func NewActivateFreeDNSRequest(domain string, source SourceEnum, ip string, ) *ActivateFreeDNSRequest`

NewActivateFreeDNSRequest instantiates a new ActivateFreeDNSRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivateFreeDNSRequestWithDefaults

`func NewActivateFreeDNSRequestWithDefaults() *ActivateFreeDNSRequest`

NewActivateFreeDNSRequestWithDefaults instantiates a new ActivateFreeDNSRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ActivateFreeDNSRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ActivateFreeDNSRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ActivateFreeDNSRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetSource

`func (o *ActivateFreeDNSRequest) GetSource() SourceEnum`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ActivateFreeDNSRequest) GetSourceOk() (*SourceEnum, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ActivateFreeDNSRequest) SetSource(v SourceEnum)`

SetSource sets Source field to given value.


### GetIp

`func (o *ActivateFreeDNSRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ActivateFreeDNSRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ActivateFreeDNSRequest) SetIp(v string)`

SetIp sets Ip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


