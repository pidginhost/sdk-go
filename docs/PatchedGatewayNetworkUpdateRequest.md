# PatchedGatewayNetworkUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snat** | Pointer to **bool** | A router network&#39;s internet egress through the gateway. | [optional] 

## Methods

### NewPatchedGatewayNetworkUpdateRequest

`func NewPatchedGatewayNetworkUpdateRequest() *PatchedGatewayNetworkUpdateRequest`

NewPatchedGatewayNetworkUpdateRequest instantiates a new PatchedGatewayNetworkUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedGatewayNetworkUpdateRequestWithDefaults

`func NewPatchedGatewayNetworkUpdateRequestWithDefaults() *PatchedGatewayNetworkUpdateRequest`

NewPatchedGatewayNetworkUpdateRequestWithDefaults instantiates a new PatchedGatewayNetworkUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnat

`func (o *PatchedGatewayNetworkUpdateRequest) GetSnat() bool`

GetSnat returns the Snat field if non-nil, zero value otherwise.

### GetSnatOk

`func (o *PatchedGatewayNetworkUpdateRequest) GetSnatOk() (*bool, bool)`

GetSnatOk returns a tuple with the Snat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnat

`func (o *PatchedGatewayNetworkUpdateRequest) SetSnat(v bool)`

SetSnat sets Snat field to given value.

### HasSnat

`func (o *PatchedGatewayNetworkUpdateRequest) HasSnat() bool`

HasSnat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


