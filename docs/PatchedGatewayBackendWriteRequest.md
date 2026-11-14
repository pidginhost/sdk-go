# PatchedGatewayBackendWriteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to **int32** | An attached private network&#39;s id. | [optional] 
**Address** | Pointer to **string** | A host of that network, other than the gateway. | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedGatewayBackendWriteRequest

`func NewPatchedGatewayBackendWriteRequest() *PatchedGatewayBackendWriteRequest`

NewPatchedGatewayBackendWriteRequest instantiates a new PatchedGatewayBackendWriteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedGatewayBackendWriteRequestWithDefaults

`func NewPatchedGatewayBackendWriteRequestWithDefaults() *PatchedGatewayBackendWriteRequest`

NewPatchedGatewayBackendWriteRequestWithDefaults instantiates a new PatchedGatewayBackendWriteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *PatchedGatewayBackendWriteRequest) GetNetwork() int32`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *PatchedGatewayBackendWriteRequest) GetNetworkOk() (*int32, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *PatchedGatewayBackendWriteRequest) SetNetwork(v int32)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *PatchedGatewayBackendWriteRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetAddress

`func (o *PatchedGatewayBackendWriteRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PatchedGatewayBackendWriteRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PatchedGatewayBackendWriteRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PatchedGatewayBackendWriteRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPort

`func (o *PatchedGatewayBackendWriteRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PatchedGatewayBackendWriteRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PatchedGatewayBackendWriteRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PatchedGatewayBackendWriteRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetWeight

`func (o *PatchedGatewayBackendWriteRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedGatewayBackendWriteRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedGatewayBackendWriteRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedGatewayBackendWriteRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedGatewayBackendWriteRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedGatewayBackendWriteRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedGatewayBackendWriteRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedGatewayBackendWriteRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


