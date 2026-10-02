# PatchedTCPRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** | External port to expose (blocked: 22, 6443, 50000, 50001) | [optional] 
**BackendServiceName** | Pointer to **string** | Name of the backend Kubernetes Service | [optional] 
**BackendServicePort** | Pointer to **int32** | Port of the backend Service | [optional] 
**BackendNamespace** | Pointer to **string** | Namespace of the backend Service | [optional] [default to "default"]

## Methods

### NewPatchedTCPRouteRequest

`func NewPatchedTCPRouteRequest() *PatchedTCPRouteRequest`

NewPatchedTCPRouteRequest instantiates a new PatchedTCPRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedTCPRouteRequestWithDefaults

`func NewPatchedTCPRouteRequestWithDefaults() *PatchedTCPRouteRequest`

NewPatchedTCPRouteRequestWithDefaults instantiates a new PatchedTCPRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedTCPRouteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedTCPRouteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedTCPRouteRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedTCPRouteRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *PatchedTCPRouteRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PatchedTCPRouteRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PatchedTCPRouteRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PatchedTCPRouteRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPort

`func (o *PatchedTCPRouteRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PatchedTCPRouteRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PatchedTCPRouteRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PatchedTCPRouteRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetBackendServiceName

`func (o *PatchedTCPRouteRequest) GetBackendServiceName() string`

GetBackendServiceName returns the BackendServiceName field if non-nil, zero value otherwise.

### GetBackendServiceNameOk

`func (o *PatchedTCPRouteRequest) GetBackendServiceNameOk() (*string, bool)`

GetBackendServiceNameOk returns a tuple with the BackendServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendServiceName

`func (o *PatchedTCPRouteRequest) SetBackendServiceName(v string)`

SetBackendServiceName sets BackendServiceName field to given value.

### HasBackendServiceName

`func (o *PatchedTCPRouteRequest) HasBackendServiceName() bool`

HasBackendServiceName returns a boolean if a field has been set.

### GetBackendServicePort

`func (o *PatchedTCPRouteRequest) GetBackendServicePort() int32`

GetBackendServicePort returns the BackendServicePort field if non-nil, zero value otherwise.

### GetBackendServicePortOk

`func (o *PatchedTCPRouteRequest) GetBackendServicePortOk() (*int32, bool)`

GetBackendServicePortOk returns a tuple with the BackendServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendServicePort

`func (o *PatchedTCPRouteRequest) SetBackendServicePort(v int32)`

SetBackendServicePort sets BackendServicePort field to given value.

### HasBackendServicePort

`func (o *PatchedTCPRouteRequest) HasBackendServicePort() bool`

HasBackendServicePort returns a boolean if a field has been set.

### GetBackendNamespace

`func (o *PatchedTCPRouteRequest) GetBackendNamespace() string`

GetBackendNamespace returns the BackendNamespace field if non-nil, zero value otherwise.

### GetBackendNamespaceOk

`func (o *PatchedTCPRouteRequest) GetBackendNamespaceOk() (*string, bool)`

GetBackendNamespaceOk returns a tuple with the BackendNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendNamespace

`func (o *PatchedTCPRouteRequest) SetBackendNamespace(v string)`

SetBackendNamespace sets BackendNamespace field to given value.

### HasBackendNamespace

`func (o *PatchedTCPRouteRequest) HasBackendNamespace() bool`

HasBackendNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


