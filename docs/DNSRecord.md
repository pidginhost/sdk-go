# DNSRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line** | **int32** |  | [readonly] 
**Name** | **string** |  | [readonly] 
**Ttl** | **int32** |  | [readonly] 
**Type** | **string** |  | [readonly] 
**Address** | **string** |  | [readonly] 
**Cname** | **string** |  | [readonly] 
**Exchange** | **string** |  | [readonly] 
**Preference** | **int32** |  | [readonly] 
**Txtdata** | **string** |  | [readonly] 
**Unencoded** | **string** |  | [readonly] 
**Target** | **string** |  | [readonly] 
**Priority** | **int32** |  | [readonly] 
**Weight** | **int32** |  | [readonly] 
**Port** | **int32** |  | [readonly] 
**Flag** | **int32** |  | [readonly] 
**Tag** | **string** |  | [readonly] 
**Value** | **string** |  | [readonly] 

## Methods

### NewDNSRecord

`func NewDNSRecord(line int32, name string, ttl int32, type_ string, address string, cname string, exchange string, preference int32, txtdata string, unencoded string, target string, priority int32, weight int32, port int32, flag int32, tag string, value string, ) *DNSRecord`

NewDNSRecord instantiates a new DNSRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSRecordWithDefaults

`func NewDNSRecordWithDefaults() *DNSRecord`

NewDNSRecordWithDefaults instantiates a new DNSRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine

`func (o *DNSRecord) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *DNSRecord) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *DNSRecord) SetLine(v int32)`

SetLine sets Line field to given value.


### GetName

`func (o *DNSRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DNSRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DNSRecord) SetName(v string)`

SetName sets Name field to given value.


### GetTtl

`func (o *DNSRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DNSRecord) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DNSRecord) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### GetType

`func (o *DNSRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSRecord) SetType(v string)`

SetType sets Type field to given value.


### GetAddress

`func (o *DNSRecord) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DNSRecord) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DNSRecord) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCname

`func (o *DNSRecord) GetCname() string`

GetCname returns the Cname field if non-nil, zero value otherwise.

### GetCnameOk

`func (o *DNSRecord) GetCnameOk() (*string, bool)`

GetCnameOk returns a tuple with the Cname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCname

`func (o *DNSRecord) SetCname(v string)`

SetCname sets Cname field to given value.


### GetExchange

`func (o *DNSRecord) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *DNSRecord) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *DNSRecord) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetPreference

`func (o *DNSRecord) GetPreference() int32`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *DNSRecord) GetPreferenceOk() (*int32, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *DNSRecord) SetPreference(v int32)`

SetPreference sets Preference field to given value.


### GetTxtdata

`func (o *DNSRecord) GetTxtdata() string`

GetTxtdata returns the Txtdata field if non-nil, zero value otherwise.

### GetTxtdataOk

`func (o *DNSRecord) GetTxtdataOk() (*string, bool)`

GetTxtdataOk returns a tuple with the Txtdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxtdata

`func (o *DNSRecord) SetTxtdata(v string)`

SetTxtdata sets Txtdata field to given value.


### GetUnencoded

`func (o *DNSRecord) GetUnencoded() string`

GetUnencoded returns the Unencoded field if non-nil, zero value otherwise.

### GetUnencodedOk

`func (o *DNSRecord) GetUnencodedOk() (*string, bool)`

GetUnencodedOk returns a tuple with the Unencoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnencoded

`func (o *DNSRecord) SetUnencoded(v string)`

SetUnencoded sets Unencoded field to given value.


### GetTarget

`func (o *DNSRecord) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DNSRecord) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DNSRecord) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetPriority

`func (o *DNSRecord) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DNSRecord) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DNSRecord) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetWeight

`func (o *DNSRecord) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DNSRecord) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DNSRecord) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetPort

`func (o *DNSRecord) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DNSRecord) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DNSRecord) SetPort(v int32)`

SetPort sets Port field to given value.


### GetFlag

`func (o *DNSRecord) GetFlag() int32`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *DNSRecord) GetFlagOk() (*int32, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *DNSRecord) SetFlag(v int32)`

SetFlag sets Flag field to given value.


### GetTag

`func (o *DNSRecord) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DNSRecord) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DNSRecord) SetTag(v string)`

SetTag sets Tag field to given value.


### GetValue

`func (o *DNSRecord) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DNSRecord) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DNSRecord) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


