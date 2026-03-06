# DNSRecordMutateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Record** | Pointer to [**DNSRecord**](DNSRecord.md) |  | [optional] 

## Methods

### NewDNSRecordMutateResponse

`func NewDNSRecordMutateResponse(message string, ) *DNSRecordMutateResponse`

NewDNSRecordMutateResponse instantiates a new DNSRecordMutateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSRecordMutateResponseWithDefaults

`func NewDNSRecordMutateResponseWithDefaults() *DNSRecordMutateResponse`

NewDNSRecordMutateResponseWithDefaults instantiates a new DNSRecordMutateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DNSRecordMutateResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DNSRecordMutateResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DNSRecordMutateResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRecord

`func (o *DNSRecordMutateResponse) GetRecord() DNSRecord`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *DNSRecordMutateResponse) GetRecordOk() (*DNSRecord, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *DNSRecordMutateResponse) SetRecord(v DNSRecord)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *DNSRecordMutateResponse) HasRecord() bool`

HasRecord returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


