# DNSRecordCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Record hostname (use &#39;@&#39; or leave empty for zone apex). | 
**Ttl** | **int32** | Time to live in seconds. | 
**Type** | [**DNSRecordCreateTypeEnum**](DNSRecordCreateTypeEnum.md) | DNS record type.  * &#x60;A&#x60; - A * &#x60;AAAA&#x60; - AAAA * &#x60;TYPE257&#x60; - TYPE257 * &#x60;CNAME&#x60; - CNAME * &#x60;MX&#x60; - MX * &#x60;SRV&#x60; - SRV * &#x60;TXT&#x60; - TXT | 
**Address** | Pointer to **string** | IPv4/IPv6 address (A/AAAA). | [optional] 
**Cname** | Pointer to **string** | Canonical name (CNAME). | [optional] 
**Exchange** | Pointer to **string** | Mail exchange host (MX). | [optional] 
**Preference** | Pointer to **int32** | MX preference / priority. | [optional] 
**Txtdata** | Pointer to **string** | TXT record data. | [optional] 
**Unencoded** | Pointer to **string** | Unencoded TXT value. | [optional] 
**Target** | Pointer to **string** | SRV target host. | [optional] 
**Priority** | Pointer to **int32** | SRV priority. | [optional] 
**Weight** | Pointer to **int32** | SRV weight. | [optional] 
**Port** | Pointer to **int32** | SRV port. | [optional] 
**Flag** | Pointer to **int32** | CAA flag (TYPE257). | [optional] 
**Tag** | Pointer to **string** | CAA tag (TYPE257). | [optional] 
**Value** | Pointer to **string** | CAA value (TYPE257). | [optional] 
**Line** | Pointer to **int32** | Line number of existing record to edit. Omit to add a new record. | [optional] 

## Methods

### NewDNSRecordCreateRequest

`func NewDNSRecordCreateRequest(name string, ttl int32, type_ DNSRecordCreateTypeEnum, ) *DNSRecordCreateRequest`

NewDNSRecordCreateRequest instantiates a new DNSRecordCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSRecordCreateRequestWithDefaults

`func NewDNSRecordCreateRequestWithDefaults() *DNSRecordCreateRequest`

NewDNSRecordCreateRequestWithDefaults instantiates a new DNSRecordCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DNSRecordCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DNSRecordCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DNSRecordCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTtl

`func (o *DNSRecordCreateRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DNSRecordCreateRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DNSRecordCreateRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### GetType

`func (o *DNSRecordCreateRequest) GetType() DNSRecordCreateTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSRecordCreateRequest) GetTypeOk() (*DNSRecordCreateTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSRecordCreateRequest) SetType(v DNSRecordCreateTypeEnum)`

SetType sets Type field to given value.


### GetAddress

`func (o *DNSRecordCreateRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DNSRecordCreateRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DNSRecordCreateRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DNSRecordCreateRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCname

`func (o *DNSRecordCreateRequest) GetCname() string`

GetCname returns the Cname field if non-nil, zero value otherwise.

### GetCnameOk

`func (o *DNSRecordCreateRequest) GetCnameOk() (*string, bool)`

GetCnameOk returns a tuple with the Cname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCname

`func (o *DNSRecordCreateRequest) SetCname(v string)`

SetCname sets Cname field to given value.

### HasCname

`func (o *DNSRecordCreateRequest) HasCname() bool`

HasCname returns a boolean if a field has been set.

### GetExchange

`func (o *DNSRecordCreateRequest) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *DNSRecordCreateRequest) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *DNSRecordCreateRequest) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *DNSRecordCreateRequest) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetPreference

`func (o *DNSRecordCreateRequest) GetPreference() int32`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *DNSRecordCreateRequest) GetPreferenceOk() (*int32, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *DNSRecordCreateRequest) SetPreference(v int32)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *DNSRecordCreateRequest) HasPreference() bool`

HasPreference returns a boolean if a field has been set.

### GetTxtdata

`func (o *DNSRecordCreateRequest) GetTxtdata() string`

GetTxtdata returns the Txtdata field if non-nil, zero value otherwise.

### GetTxtdataOk

`func (o *DNSRecordCreateRequest) GetTxtdataOk() (*string, bool)`

GetTxtdataOk returns a tuple with the Txtdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxtdata

`func (o *DNSRecordCreateRequest) SetTxtdata(v string)`

SetTxtdata sets Txtdata field to given value.

### HasTxtdata

`func (o *DNSRecordCreateRequest) HasTxtdata() bool`

HasTxtdata returns a boolean if a field has been set.

### GetUnencoded

`func (o *DNSRecordCreateRequest) GetUnencoded() string`

GetUnencoded returns the Unencoded field if non-nil, zero value otherwise.

### GetUnencodedOk

`func (o *DNSRecordCreateRequest) GetUnencodedOk() (*string, bool)`

GetUnencodedOk returns a tuple with the Unencoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnencoded

`func (o *DNSRecordCreateRequest) SetUnencoded(v string)`

SetUnencoded sets Unencoded field to given value.

### HasUnencoded

`func (o *DNSRecordCreateRequest) HasUnencoded() bool`

HasUnencoded returns a boolean if a field has been set.

### GetTarget

`func (o *DNSRecordCreateRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DNSRecordCreateRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DNSRecordCreateRequest) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *DNSRecordCreateRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetPriority

`func (o *DNSRecordCreateRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DNSRecordCreateRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DNSRecordCreateRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DNSRecordCreateRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetWeight

`func (o *DNSRecordCreateRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DNSRecordCreateRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DNSRecordCreateRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *DNSRecordCreateRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetPort

`func (o *DNSRecordCreateRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DNSRecordCreateRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DNSRecordCreateRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *DNSRecordCreateRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetFlag

`func (o *DNSRecordCreateRequest) GetFlag() int32`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *DNSRecordCreateRequest) GetFlagOk() (*int32, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *DNSRecordCreateRequest) SetFlag(v int32)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *DNSRecordCreateRequest) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetTag

`func (o *DNSRecordCreateRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DNSRecordCreateRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DNSRecordCreateRequest) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *DNSRecordCreateRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetValue

`func (o *DNSRecordCreateRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DNSRecordCreateRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DNSRecordCreateRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DNSRecordCreateRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLine

`func (o *DNSRecordCreateRequest) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *DNSRecordCreateRequest) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *DNSRecordCreateRequest) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *DNSRecordCreateRequest) HasLine() bool`

HasLine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


