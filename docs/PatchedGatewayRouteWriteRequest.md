# PatchedGatewayRouteWriteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostnames** | Pointer to **[]string** | Up to 16 names; empty matches every host. | [optional] 
**PathPrefix** | Pointer to **string** | Literal and case-sensitive. | [optional] 
**Pool** | Pointer to **int32** | A backend pool&#39;s id. | [optional] 
**RedirectHttps** | Pointer to **bool** | Only false is accepted for now. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedGatewayRouteWriteRequest

`func NewPatchedGatewayRouteWriteRequest() *PatchedGatewayRouteWriteRequest`

NewPatchedGatewayRouteWriteRequest instantiates a new PatchedGatewayRouteWriteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedGatewayRouteWriteRequestWithDefaults

`func NewPatchedGatewayRouteWriteRequestWithDefaults() *PatchedGatewayRouteWriteRequest`

NewPatchedGatewayRouteWriteRequestWithDefaults instantiates a new PatchedGatewayRouteWriteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostnames

`func (o *PatchedGatewayRouteWriteRequest) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *PatchedGatewayRouteWriteRequest) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *PatchedGatewayRouteWriteRequest) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.

### HasHostnames

`func (o *PatchedGatewayRouteWriteRequest) HasHostnames() bool`

HasHostnames returns a boolean if a field has been set.

### GetPathPrefix

`func (o *PatchedGatewayRouteWriteRequest) GetPathPrefix() string`

GetPathPrefix returns the PathPrefix field if non-nil, zero value otherwise.

### GetPathPrefixOk

`func (o *PatchedGatewayRouteWriteRequest) GetPathPrefixOk() (*string, bool)`

GetPathPrefixOk returns a tuple with the PathPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPrefix

`func (o *PatchedGatewayRouteWriteRequest) SetPathPrefix(v string)`

SetPathPrefix sets PathPrefix field to given value.

### HasPathPrefix

`func (o *PatchedGatewayRouteWriteRequest) HasPathPrefix() bool`

HasPathPrefix returns a boolean if a field has been set.

### GetPool

`func (o *PatchedGatewayRouteWriteRequest) GetPool() int32`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *PatchedGatewayRouteWriteRequest) GetPoolOk() (*int32, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *PatchedGatewayRouteWriteRequest) SetPool(v int32)`

SetPool sets Pool field to given value.

### HasPool

`func (o *PatchedGatewayRouteWriteRequest) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetRedirectHttps

`func (o *PatchedGatewayRouteWriteRequest) GetRedirectHttps() bool`

GetRedirectHttps returns the RedirectHttps field if non-nil, zero value otherwise.

### GetRedirectHttpsOk

`func (o *PatchedGatewayRouteWriteRequest) GetRedirectHttpsOk() (*bool, bool)`

GetRedirectHttpsOk returns a tuple with the RedirectHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectHttps

`func (o *PatchedGatewayRouteWriteRequest) SetRedirectHttps(v bool)`

SetRedirectHttps sets RedirectHttps field to given value.

### HasRedirectHttps

`func (o *PatchedGatewayRouteWriteRequest) HasRedirectHttps() bool`

HasRedirectHttps returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedGatewayRouteWriteRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedGatewayRouteWriteRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedGatewayRouteWriteRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedGatewayRouteWriteRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


