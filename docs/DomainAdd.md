# DomainAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DnsSource** | [**DnsSourceEnum**](DnsSourceEnum.md) |  | 
**ManagedDomain** | Pointer to **NullableInt32** |  | [optional] 
**ManagedExternalDomain** | Pointer to **NullableInt32** |  | [optional] 
**UseInbound** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewDomainAdd

`func NewDomainAdd(name string, dnsSource DnsSourceEnum, ) *DomainAdd`

NewDomainAdd instantiates a new DomainAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainAddWithDefaults

`func NewDomainAddWithDefaults() *DomainAdd`

NewDomainAddWithDefaults instantiates a new DomainAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DomainAdd) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainAdd) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainAdd) SetName(v string)`

SetName sets Name field to given value.


### GetDnsSource

`func (o *DomainAdd) GetDnsSource() DnsSourceEnum`

GetDnsSource returns the DnsSource field if non-nil, zero value otherwise.

### GetDnsSourceOk

`func (o *DomainAdd) GetDnsSourceOk() (*DnsSourceEnum, bool)`

GetDnsSourceOk returns a tuple with the DnsSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSource

`func (o *DomainAdd) SetDnsSource(v DnsSourceEnum)`

SetDnsSource sets DnsSource field to given value.


### GetManagedDomain

`func (o *DomainAdd) GetManagedDomain() int32`

GetManagedDomain returns the ManagedDomain field if non-nil, zero value otherwise.

### GetManagedDomainOk

`func (o *DomainAdd) GetManagedDomainOk() (*int32, bool)`

GetManagedDomainOk returns a tuple with the ManagedDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedDomain

`func (o *DomainAdd) SetManagedDomain(v int32)`

SetManagedDomain sets ManagedDomain field to given value.

### HasManagedDomain

`func (o *DomainAdd) HasManagedDomain() bool`

HasManagedDomain returns a boolean if a field has been set.

### SetManagedDomainNil

`func (o *DomainAdd) SetManagedDomainNil(b bool)`

 SetManagedDomainNil sets the value for ManagedDomain to be an explicit nil

### UnsetManagedDomain
`func (o *DomainAdd) UnsetManagedDomain()`

UnsetManagedDomain ensures that no value is present for ManagedDomain, not even an explicit nil
### GetManagedExternalDomain

`func (o *DomainAdd) GetManagedExternalDomain() int32`

GetManagedExternalDomain returns the ManagedExternalDomain field if non-nil, zero value otherwise.

### GetManagedExternalDomainOk

`func (o *DomainAdd) GetManagedExternalDomainOk() (*int32, bool)`

GetManagedExternalDomainOk returns a tuple with the ManagedExternalDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedExternalDomain

`func (o *DomainAdd) SetManagedExternalDomain(v int32)`

SetManagedExternalDomain sets ManagedExternalDomain field to given value.

### HasManagedExternalDomain

`func (o *DomainAdd) HasManagedExternalDomain() bool`

HasManagedExternalDomain returns a boolean if a field has been set.

### SetManagedExternalDomainNil

`func (o *DomainAdd) SetManagedExternalDomainNil(b bool)`

 SetManagedExternalDomainNil sets the value for ManagedExternalDomain to be an explicit nil

### UnsetManagedExternalDomain
`func (o *DomainAdd) UnsetManagedExternalDomain()`

UnsetManagedExternalDomain ensures that no value is present for ManagedExternalDomain, not even an explicit nil
### GetUseInbound

`func (o *DomainAdd) GetUseInbound() bool`

GetUseInbound returns the UseInbound field if non-nil, zero value otherwise.

### GetUseInboundOk

`func (o *DomainAdd) GetUseInboundOk() (*bool, bool)`

GetUseInboundOk returns a tuple with the UseInbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseInbound

`func (o *DomainAdd) SetUseInbound(v bool)`

SetUseInbound sets UseInbound field to given value.

### HasUseInbound

`func (o *DomainAdd) HasUseInbound() bool`

HasUseInbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


