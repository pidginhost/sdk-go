# SendingDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | [readonly] 
**Status** | [**SendingDomainStatusEnum**](SendingDomainStatusEnum.md) |  | [readonly] 
**DnsSource** | [**DnsSourceEnum**](DnsSourceEnum.md) |  | [readonly] 
**UseInbound** | **bool** |  | [readonly] 
**DkimSelector** | **string** |  | [readonly] 
**DkimRecord** | **string** |  | [readonly] 
**SpfRecord** | **string** |  | [readonly] 
**DmarcRecord** | **string** |  | [readonly] 
**VerifiedAt** | **NullableString** |  | [readonly] 
**LastCheckAt** | **NullableString** |  | [readonly] 
**LastCheckErrors** | **interface{}** |  | [readonly] 

## Methods

### NewSendingDomain

`func NewSendingDomain(id int32, name string, status SendingDomainStatusEnum, dnsSource DnsSourceEnum, useInbound bool, dkimSelector string, dkimRecord string, spfRecord string, dmarcRecord string, verifiedAt NullableString, lastCheckAt NullableString, lastCheckErrors interface{}, ) *SendingDomain`

NewSendingDomain instantiates a new SendingDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendingDomainWithDefaults

`func NewSendingDomainWithDefaults() *SendingDomain`

NewSendingDomainWithDefaults instantiates a new SendingDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SendingDomain) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SendingDomain) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SendingDomain) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *SendingDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SendingDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SendingDomain) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *SendingDomain) GetStatus() SendingDomainStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SendingDomain) GetStatusOk() (*SendingDomainStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SendingDomain) SetStatus(v SendingDomainStatusEnum)`

SetStatus sets Status field to given value.


### GetDnsSource

`func (o *SendingDomain) GetDnsSource() DnsSourceEnum`

GetDnsSource returns the DnsSource field if non-nil, zero value otherwise.

### GetDnsSourceOk

`func (o *SendingDomain) GetDnsSourceOk() (*DnsSourceEnum, bool)`

GetDnsSourceOk returns a tuple with the DnsSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSource

`func (o *SendingDomain) SetDnsSource(v DnsSourceEnum)`

SetDnsSource sets DnsSource field to given value.


### GetUseInbound

`func (o *SendingDomain) GetUseInbound() bool`

GetUseInbound returns the UseInbound field if non-nil, zero value otherwise.

### GetUseInboundOk

`func (o *SendingDomain) GetUseInboundOk() (*bool, bool)`

GetUseInboundOk returns a tuple with the UseInbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseInbound

`func (o *SendingDomain) SetUseInbound(v bool)`

SetUseInbound sets UseInbound field to given value.


### GetDkimSelector

`func (o *SendingDomain) GetDkimSelector() string`

GetDkimSelector returns the DkimSelector field if non-nil, zero value otherwise.

### GetDkimSelectorOk

`func (o *SendingDomain) GetDkimSelectorOk() (*string, bool)`

GetDkimSelectorOk returns a tuple with the DkimSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkimSelector

`func (o *SendingDomain) SetDkimSelector(v string)`

SetDkimSelector sets DkimSelector field to given value.


### GetDkimRecord

`func (o *SendingDomain) GetDkimRecord() string`

GetDkimRecord returns the DkimRecord field if non-nil, zero value otherwise.

### GetDkimRecordOk

`func (o *SendingDomain) GetDkimRecordOk() (*string, bool)`

GetDkimRecordOk returns a tuple with the DkimRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkimRecord

`func (o *SendingDomain) SetDkimRecord(v string)`

SetDkimRecord sets DkimRecord field to given value.


### GetSpfRecord

`func (o *SendingDomain) GetSpfRecord() string`

GetSpfRecord returns the SpfRecord field if non-nil, zero value otherwise.

### GetSpfRecordOk

`func (o *SendingDomain) GetSpfRecordOk() (*string, bool)`

GetSpfRecordOk returns a tuple with the SpfRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpfRecord

`func (o *SendingDomain) SetSpfRecord(v string)`

SetSpfRecord sets SpfRecord field to given value.


### GetDmarcRecord

`func (o *SendingDomain) GetDmarcRecord() string`

GetDmarcRecord returns the DmarcRecord field if non-nil, zero value otherwise.

### GetDmarcRecordOk

`func (o *SendingDomain) GetDmarcRecordOk() (*string, bool)`

GetDmarcRecordOk returns a tuple with the DmarcRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmarcRecord

`func (o *SendingDomain) SetDmarcRecord(v string)`

SetDmarcRecord sets DmarcRecord field to given value.


### GetVerifiedAt

`func (o *SendingDomain) GetVerifiedAt() string`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *SendingDomain) GetVerifiedAtOk() (*string, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *SendingDomain) SetVerifiedAt(v string)`

SetVerifiedAt sets VerifiedAt field to given value.


### SetVerifiedAtNil

`func (o *SendingDomain) SetVerifiedAtNil(b bool)`

 SetVerifiedAtNil sets the value for VerifiedAt to be an explicit nil

### UnsetVerifiedAt
`func (o *SendingDomain) UnsetVerifiedAt()`

UnsetVerifiedAt ensures that no value is present for VerifiedAt, not even an explicit nil
### GetLastCheckAt

`func (o *SendingDomain) GetLastCheckAt() string`

GetLastCheckAt returns the LastCheckAt field if non-nil, zero value otherwise.

### GetLastCheckAtOk

`func (o *SendingDomain) GetLastCheckAtOk() (*string, bool)`

GetLastCheckAtOk returns a tuple with the LastCheckAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckAt

`func (o *SendingDomain) SetLastCheckAt(v string)`

SetLastCheckAt sets LastCheckAt field to given value.


### SetLastCheckAtNil

`func (o *SendingDomain) SetLastCheckAtNil(b bool)`

 SetLastCheckAtNil sets the value for LastCheckAt to be an explicit nil

### UnsetLastCheckAt
`func (o *SendingDomain) UnsetLastCheckAt()`

UnsetLastCheckAt ensures that no value is present for LastCheckAt, not even an explicit nil
### GetLastCheckErrors

`func (o *SendingDomain) GetLastCheckErrors() interface{}`

GetLastCheckErrors returns the LastCheckErrors field if non-nil, zero value otherwise.

### GetLastCheckErrorsOk

`func (o *SendingDomain) GetLastCheckErrorsOk() (*interface{}, bool)`

GetLastCheckErrorsOk returns a tuple with the LastCheckErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckErrors

`func (o *SendingDomain) SetLastCheckErrors(v interface{})`

SetLastCheckErrors sets LastCheckErrors field to given value.


### SetLastCheckErrorsNil

`func (o *SendingDomain) SetLastCheckErrorsNil(b bool)`

 SetLastCheckErrorsNil sets the value for LastCheckErrors to be an explicit nil

### UnsetLastCheckErrors
`func (o *SendingDomain) UnsetLastCheckErrors()`

UnsetLastCheckErrors ensures that no value is present for LastCheckErrors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


