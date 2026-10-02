# DomainAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DnsSource** | [**DnsSourceEnum**](DnsSourceEnum.md) |  | 
**ManagedDomain** | Pointer to **NullableInt32** |  | [optional] 
**ManagedExternalDomain** | Pointer to **NullableInt32** |  | [optional] 
**UseInbound** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewDomainAddRequest

`func NewDomainAddRequest(name string, dnsSource DnsSourceEnum, ) *DomainAddRequest`

NewDomainAddRequest instantiates a new DomainAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainAddRequestWithDefaults

`func NewDomainAddRequestWithDefaults() *DomainAddRequest`

NewDomainAddRequestWithDefaults instantiates a new DomainAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DomainAddRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainAddRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainAddRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDnsSource

`func (o *DomainAddRequest) GetDnsSource() DnsSourceEnum`

GetDnsSource returns the DnsSource field if non-nil, zero value otherwise.

### GetDnsSourceOk

`func (o *DomainAddRequest) GetDnsSourceOk() (*DnsSourceEnum, bool)`

GetDnsSourceOk returns a tuple with the DnsSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSource

`func (o *DomainAddRequest) SetDnsSource(v DnsSourceEnum)`

SetDnsSource sets DnsSource field to given value.


### GetManagedDomain

`func (o *DomainAddRequest) GetManagedDomain() int32`

GetManagedDomain returns the ManagedDomain field if non-nil, zero value otherwise.

### GetManagedDomainOk

`func (o *DomainAddRequest) GetManagedDomainOk() (*int32, bool)`

GetManagedDomainOk returns a tuple with the ManagedDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedDomain

`func (o *DomainAddRequest) SetManagedDomain(v int32)`

SetManagedDomain sets ManagedDomain field to given value.

### HasManagedDomain

`func (o *DomainAddRequest) HasManagedDomain() bool`

HasManagedDomain returns a boolean if a field has been set.

### SetManagedDomainNil

`func (o *DomainAddRequest) SetManagedDomainNil(b bool)`

 SetManagedDomainNil sets the value for ManagedDomain to be an explicit nil

### UnsetManagedDomain
`func (o *DomainAddRequest) UnsetManagedDomain()`

UnsetManagedDomain ensures that no value is present for ManagedDomain, not even an explicit nil
### GetManagedExternalDomain

`func (o *DomainAddRequest) GetManagedExternalDomain() int32`

GetManagedExternalDomain returns the ManagedExternalDomain field if non-nil, zero value otherwise.

### GetManagedExternalDomainOk

`func (o *DomainAddRequest) GetManagedExternalDomainOk() (*int32, bool)`

GetManagedExternalDomainOk returns a tuple with the ManagedExternalDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedExternalDomain

`func (o *DomainAddRequest) SetManagedExternalDomain(v int32)`

SetManagedExternalDomain sets ManagedExternalDomain field to given value.

### HasManagedExternalDomain

`func (o *DomainAddRequest) HasManagedExternalDomain() bool`

HasManagedExternalDomain returns a boolean if a field has been set.

### SetManagedExternalDomainNil

`func (o *DomainAddRequest) SetManagedExternalDomainNil(b bool)`

 SetManagedExternalDomainNil sets the value for ManagedExternalDomain to be an explicit nil

### UnsetManagedExternalDomain
`func (o *DomainAddRequest) UnsetManagedExternalDomain()`

UnsetManagedExternalDomain ensures that no value is present for ManagedExternalDomain, not even an explicit nil
### GetUseInbound

`func (o *DomainAddRequest) GetUseInbound() bool`

GetUseInbound returns the UseInbound field if non-nil, zero value otherwise.

### GetUseInboundOk

`func (o *DomainAddRequest) GetUseInboundOk() (*bool, bool)`

GetUseInboundOk returns a tuple with the UseInbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseInbound

`func (o *DomainAddRequest) SetUseInbound(v bool)`

SetUseInbound sets UseInbound field to given value.

### HasUseInbound

`func (o *DomainAddRequest) HasUseInbound() bool`

HasUseInbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


