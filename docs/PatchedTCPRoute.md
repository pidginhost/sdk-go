# PatchedTCPRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** | External port to expose (blocked: 22, 6443, 50000, 50001) | [optional] 
**BackendServiceName** | Pointer to **string** | Name of the backend Kubernetes Service | [optional] 
**BackendServicePort** | Pointer to **int32** | Port of the backend Service | [optional] 
**BackendNamespace** | Pointer to **string** | Namespace of the backend Service | [optional] [default to "default"]
**StatusReady** | Pointer to **NullableBool** |  | [optional] [readonly] 
**StatusMessage** | Pointer to **string** |  | [optional] [readonly] 
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**Updated** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewPatchedTCPRoute

`func NewPatchedTCPRoute() *PatchedTCPRoute`

NewPatchedTCPRoute instantiates a new PatchedTCPRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedTCPRouteWithDefaults

`func NewPatchedTCPRouteWithDefaults() *PatchedTCPRoute`

NewPatchedTCPRouteWithDefaults instantiates a new PatchedTCPRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedTCPRoute) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedTCPRoute) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedTCPRoute) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedTCPRoute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedTCPRoute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedTCPRoute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedTCPRoute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedTCPRoute) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *PatchedTCPRoute) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PatchedTCPRoute) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PatchedTCPRoute) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PatchedTCPRoute) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPort

`func (o *PatchedTCPRoute) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PatchedTCPRoute) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PatchedTCPRoute) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PatchedTCPRoute) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetBackendServiceName

`func (o *PatchedTCPRoute) GetBackendServiceName() string`

GetBackendServiceName returns the BackendServiceName field if non-nil, zero value otherwise.

### GetBackendServiceNameOk

`func (o *PatchedTCPRoute) GetBackendServiceNameOk() (*string, bool)`

GetBackendServiceNameOk returns a tuple with the BackendServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendServiceName

`func (o *PatchedTCPRoute) SetBackendServiceName(v string)`

SetBackendServiceName sets BackendServiceName field to given value.

### HasBackendServiceName

`func (o *PatchedTCPRoute) HasBackendServiceName() bool`

HasBackendServiceName returns a boolean if a field has been set.

### GetBackendServicePort

`func (o *PatchedTCPRoute) GetBackendServicePort() int32`

GetBackendServicePort returns the BackendServicePort field if non-nil, zero value otherwise.

### GetBackendServicePortOk

`func (o *PatchedTCPRoute) GetBackendServicePortOk() (*int32, bool)`

GetBackendServicePortOk returns a tuple with the BackendServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendServicePort

`func (o *PatchedTCPRoute) SetBackendServicePort(v int32)`

SetBackendServicePort sets BackendServicePort field to given value.

### HasBackendServicePort

`func (o *PatchedTCPRoute) HasBackendServicePort() bool`

HasBackendServicePort returns a boolean if a field has been set.

### GetBackendNamespace

`func (o *PatchedTCPRoute) GetBackendNamespace() string`

GetBackendNamespace returns the BackendNamespace field if non-nil, zero value otherwise.

### GetBackendNamespaceOk

`func (o *PatchedTCPRoute) GetBackendNamespaceOk() (*string, bool)`

GetBackendNamespaceOk returns a tuple with the BackendNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendNamespace

`func (o *PatchedTCPRoute) SetBackendNamespace(v string)`

SetBackendNamespace sets BackendNamespace field to given value.

### HasBackendNamespace

`func (o *PatchedTCPRoute) HasBackendNamespace() bool`

HasBackendNamespace returns a boolean if a field has been set.

### GetStatusReady

`func (o *PatchedTCPRoute) GetStatusReady() bool`

GetStatusReady returns the StatusReady field if non-nil, zero value otherwise.

### GetStatusReadyOk

`func (o *PatchedTCPRoute) GetStatusReadyOk() (*bool, bool)`

GetStatusReadyOk returns a tuple with the StatusReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReady

`func (o *PatchedTCPRoute) SetStatusReady(v bool)`

SetStatusReady sets StatusReady field to given value.

### HasStatusReady

`func (o *PatchedTCPRoute) HasStatusReady() bool`

HasStatusReady returns a boolean if a field has been set.

### SetStatusReadyNil

`func (o *PatchedTCPRoute) SetStatusReadyNil(b bool)`

 SetStatusReadyNil sets the value for StatusReady to be an explicit nil

### UnsetStatusReady
`func (o *PatchedTCPRoute) UnsetStatusReady()`

UnsetStatusReady ensures that no value is present for StatusReady, not even an explicit nil
### GetStatusMessage

`func (o *PatchedTCPRoute) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *PatchedTCPRoute) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *PatchedTCPRoute) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *PatchedTCPRoute) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetCreated

`func (o *PatchedTCPRoute) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PatchedTCPRoute) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PatchedTCPRoute) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PatchedTCPRoute) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *PatchedTCPRoute) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PatchedTCPRoute) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PatchedTCPRoute) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PatchedTCPRoute) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


