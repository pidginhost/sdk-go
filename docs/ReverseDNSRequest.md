# ReverseDNSRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReverseDns** | **string** | Fully-qualified domain name for PTR record (e.g., host.example.com) | 

## Methods

### NewReverseDNSRequest

`func NewReverseDNSRequest(reverseDns string, ) *ReverseDNSRequest`

NewReverseDNSRequest instantiates a new ReverseDNSRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReverseDNSRequestWithDefaults

`func NewReverseDNSRequestWithDefaults() *ReverseDNSRequest`

NewReverseDNSRequestWithDefaults instantiates a new ReverseDNSRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReverseDns

`func (o *ReverseDNSRequest) GetReverseDns() string`

GetReverseDns returns the ReverseDns field if non-nil, zero value otherwise.

### GetReverseDnsOk

`func (o *ReverseDNSRequest) GetReverseDnsOk() (*string, bool)`

GetReverseDnsOk returns a tuple with the ReverseDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseDns

`func (o *ReverseDNSRequest) SetReverseDns(v string)`

SetReverseDns sets ReverseDns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


