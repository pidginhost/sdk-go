# GatewayCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Ipv4** | Pointer to **NullableString** |  | [optional] 
**Ipv6** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewGatewayCreateRequest

`func NewGatewayCreateRequest(name string, ) *GatewayCreateRequest`

NewGatewayCreateRequest instantiates a new GatewayCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateRequestWithDefaults

`func NewGatewayCreateRequestWithDefaults() *GatewayCreateRequest`

NewGatewayCreateRequestWithDefaults instantiates a new GatewayCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GatewayCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIpv4

`func (o *GatewayCreateRequest) GetIpv4() string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *GatewayCreateRequest) GetIpv4Ok() (*string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *GatewayCreateRequest) SetIpv4(v string)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *GatewayCreateRequest) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### SetIpv4Nil

`func (o *GatewayCreateRequest) SetIpv4Nil(b bool)`

 SetIpv4Nil sets the value for Ipv4 to be an explicit nil

### UnsetIpv4
`func (o *GatewayCreateRequest) UnsetIpv4()`

UnsetIpv4 ensures that no value is present for Ipv4, not even an explicit nil
### GetIpv6

`func (o *GatewayCreateRequest) GetIpv6() bool`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *GatewayCreateRequest) GetIpv6Ok() (*bool, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *GatewayCreateRequest) SetIpv6(v bool)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *GatewayCreateRequest) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


