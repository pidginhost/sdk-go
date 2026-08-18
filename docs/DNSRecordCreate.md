# DNSRecordCreate

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

### NewDNSRecordCreate

`func NewDNSRecordCreate(name string, ttl int32, type_ DNSRecordCreateTypeEnum, ) *DNSRecordCreate`

NewDNSRecordCreate instantiates a new DNSRecordCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSRecordCreateWithDefaults

`func NewDNSRecordCreateWithDefaults() *DNSRecordCreate`

NewDNSRecordCreateWithDefaults instantiates a new DNSRecordCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DNSRecordCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DNSRecordCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DNSRecordCreate) SetName(v string)`

SetName sets Name field to given value.


### GetTtl

`func (o *DNSRecordCreate) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DNSRecordCreate) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DNSRecordCreate) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### GetType

`func (o *DNSRecordCreate) GetType() DNSRecordCreateTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSRecordCreate) GetTypeOk() (*DNSRecordCreateTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSRecordCreate) SetType(v DNSRecordCreateTypeEnum)`

SetType sets Type field to given value.


### GetAddress

`func (o *DNSRecordCreate) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DNSRecordCreate) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DNSRecordCreate) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DNSRecordCreate) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCname

`func (o *DNSRecordCreate) GetCname() string`

GetCname returns the Cname field if non-nil, zero value otherwise.

### GetCnameOk

`func (o *DNSRecordCreate) GetCnameOk() (*string, bool)`

GetCnameOk returns a tuple with the Cname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCname

`func (o *DNSRecordCreate) SetCname(v string)`

SetCname sets Cname field to given value.

### HasCname

`func (o *DNSRecordCreate) HasCname() bool`

HasCname returns a boolean if a field has been set.

### GetExchange

`func (o *DNSRecordCreate) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *DNSRecordCreate) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *DNSRecordCreate) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *DNSRecordCreate) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetPreference

`func (o *DNSRecordCreate) GetPreference() int32`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *DNSRecordCreate) GetPreferenceOk() (*int32, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *DNSRecordCreate) SetPreference(v int32)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *DNSRecordCreate) HasPreference() bool`

HasPreference returns a boolean if a field has been set.

### GetTxtdata

`func (o *DNSRecordCreate) GetTxtdata() string`

GetTxtdata returns the Txtdata field if non-nil, zero value otherwise.

### GetTxtdataOk

`func (o *DNSRecordCreate) GetTxtdataOk() (*string, bool)`

GetTxtdataOk returns a tuple with the Txtdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxtdata

`func (o *DNSRecordCreate) SetTxtdata(v string)`

SetTxtdata sets Txtdata field to given value.

### HasTxtdata

`func (o *DNSRecordCreate) HasTxtdata() bool`

HasTxtdata returns a boolean if a field has been set.

### GetUnencoded

`func (o *DNSRecordCreate) GetUnencoded() string`

GetUnencoded returns the Unencoded field if non-nil, zero value otherwise.

### GetUnencodedOk

`func (o *DNSRecordCreate) GetUnencodedOk() (*string, bool)`

GetUnencodedOk returns a tuple with the Unencoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnencoded

`func (o *DNSRecordCreate) SetUnencoded(v string)`

SetUnencoded sets Unencoded field to given value.

### HasUnencoded

`func (o *DNSRecordCreate) HasUnencoded() bool`

HasUnencoded returns a boolean if a field has been set.

### GetTarget

`func (o *DNSRecordCreate) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DNSRecordCreate) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DNSRecordCreate) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *DNSRecordCreate) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetPriority

`func (o *DNSRecordCreate) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DNSRecordCreate) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DNSRecordCreate) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DNSRecordCreate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetWeight

`func (o *DNSRecordCreate) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DNSRecordCreate) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DNSRecordCreate) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *DNSRecordCreate) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetPort

`func (o *DNSRecordCreate) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DNSRecordCreate) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DNSRecordCreate) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *DNSRecordCreate) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetFlag

`func (o *DNSRecordCreate) GetFlag() int32`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *DNSRecordCreate) GetFlagOk() (*int32, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *DNSRecordCreate) SetFlag(v int32)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *DNSRecordCreate) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetTag

`func (o *DNSRecordCreate) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DNSRecordCreate) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DNSRecordCreate) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *DNSRecordCreate) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetValue

`func (o *DNSRecordCreate) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DNSRecordCreate) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DNSRecordCreate) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DNSRecordCreate) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLine

`func (o *DNSRecordCreate) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *DNSRecordCreate) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *DNSRecordCreate) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *DNSRecordCreate) HasLine() bool`

HasLine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


