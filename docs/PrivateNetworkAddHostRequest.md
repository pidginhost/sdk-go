# PrivateNetworkAddHostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | **string** | Server hostname | 
**Address** | Pointer to **string** |  | [optional] 

## Methods

### NewPrivateNetworkAddHostRequest

`func NewPrivateNetworkAddHostRequest(server string, ) *PrivateNetworkAddHostRequest`

NewPrivateNetworkAddHostRequest instantiates a new PrivateNetworkAddHostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateNetworkAddHostRequestWithDefaults

`func NewPrivateNetworkAddHostRequestWithDefaults() *PrivateNetworkAddHostRequest`

NewPrivateNetworkAddHostRequestWithDefaults instantiates a new PrivateNetworkAddHostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *PrivateNetworkAddHostRequest) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *PrivateNetworkAddHostRequest) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *PrivateNetworkAddHostRequest) SetServer(v string)`

SetServer sets Server field to given value.


### GetAddress

`func (o *PrivateNetworkAddHostRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PrivateNetworkAddHostRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PrivateNetworkAddHostRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PrivateNetworkAddHostRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


