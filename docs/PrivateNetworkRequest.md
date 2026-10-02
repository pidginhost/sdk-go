# PrivateNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | CIDR format | 
**Gateway** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPrivateNetworkRequest

`func NewPrivateNetworkRequest(address string, ) *PrivateNetworkRequest`

NewPrivateNetworkRequest instantiates a new PrivateNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateNetworkRequestWithDefaults

`func NewPrivateNetworkRequestWithDefaults() *PrivateNetworkRequest`

NewPrivateNetworkRequestWithDefaults instantiates a new PrivateNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PrivateNetworkRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PrivateNetworkRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PrivateNetworkRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetGateway

`func (o *PrivateNetworkRequest) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *PrivateNetworkRequest) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *PrivateNetworkRequest) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *PrivateNetworkRequest) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *PrivateNetworkRequest) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *PrivateNetworkRequest) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


