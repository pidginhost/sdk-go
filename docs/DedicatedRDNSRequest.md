# DedicatedRDNSRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpId** | **int32** |  | 
**ReverseDns** | **string** |  | 

## Methods

### NewDedicatedRDNSRequest

`func NewDedicatedRDNSRequest(ipId int32, reverseDns string, ) *DedicatedRDNSRequest`

NewDedicatedRDNSRequest instantiates a new DedicatedRDNSRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedRDNSRequestWithDefaults

`func NewDedicatedRDNSRequestWithDefaults() *DedicatedRDNSRequest`

NewDedicatedRDNSRequestWithDefaults instantiates a new DedicatedRDNSRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpId

`func (o *DedicatedRDNSRequest) GetIpId() int32`

GetIpId returns the IpId field if non-nil, zero value otherwise.

### GetIpIdOk

`func (o *DedicatedRDNSRequest) GetIpIdOk() (*int32, bool)`

GetIpIdOk returns a tuple with the IpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpId

`func (o *DedicatedRDNSRequest) SetIpId(v int32)`

SetIpId sets IpId field to given value.


### GetReverseDns

`func (o *DedicatedRDNSRequest) GetReverseDns() string`

GetReverseDns returns the ReverseDns field if non-nil, zero value otherwise.

### GetReverseDnsOk

`func (o *DedicatedRDNSRequest) GetReverseDnsOk() (*string, bool)`

GetReverseDnsOk returns a tuple with the ReverseDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseDns

`func (o *DedicatedRDNSRequest) SetReverseDns(v string)`

SetReverseDns sets ReverseDns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


