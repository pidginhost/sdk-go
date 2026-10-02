# DNSGlueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | only subdomain part | 
**Ip** | **string** |  | 
**Ip2** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDNSGlueRequest

`func NewDNSGlueRequest(name string, ip string, ) *DNSGlueRequest`

NewDNSGlueRequest instantiates a new DNSGlueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSGlueRequestWithDefaults

`func NewDNSGlueRequestWithDefaults() *DNSGlueRequest`

NewDNSGlueRequestWithDefaults instantiates a new DNSGlueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DNSGlueRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DNSGlueRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DNSGlueRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIp

`func (o *DNSGlueRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DNSGlueRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DNSGlueRequest) SetIp(v string)`

SetIp sets Ip field to given value.


### GetIp2

`func (o *DNSGlueRequest) GetIp2() string`

GetIp2 returns the Ip2 field if non-nil, zero value otherwise.

### GetIp2Ok

`func (o *DNSGlueRequest) GetIp2Ok() (*string, bool)`

GetIp2Ok returns a tuple with the Ip2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp2

`func (o *DNSGlueRequest) SetIp2(v string)`

SetIp2 sets Ip2 field to given value.

### HasIp2

`func (o *DNSGlueRequest) HasIp2() bool`

HasIp2 returns a boolean if a field has been set.

### SetIp2Nil

`func (o *DNSGlueRequest) SetIp2Nil(b bool)`

 SetIp2Nil sets the value for Ip2 to be an explicit nil

### UnsetIp2
`func (o *DNSGlueRequest) UnsetIp2()`

UnsetIp2 ensures that no value is present for Ip2, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


