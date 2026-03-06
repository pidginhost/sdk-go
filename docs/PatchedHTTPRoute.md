# PatchedHTTPRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Hostnames** | Pointer to **[]string** | List of hostnames to route (e.g., [\&quot;example.com\&quot;, \&quot;www.example.com\&quot;]) | [optional] 
**BackendServiceName** | Pointer to **string** | Name of the backend Kubernetes Service | [optional] 
**BackendServicePort** | Pointer to **int32** | Port of the backend Service | [optional] 
**BackendNamespace** | Pointer to **string** | Namespace of the backend Service | [optional] [default to "default"]
**PathPrefix** | Pointer to **string** | Path prefix to match (default: /) | [optional] [default to "/"]
**EnableTls** | Pointer to **bool** | Enable TLS termination with automatic certificate issuance | [optional] [default to true]
**StatusReady** | Pointer to **NullableBool** |  | [optional] [readonly] 
**StatusMessage** | Pointer to **string** |  | [optional] [readonly] 
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**Updated** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewPatchedHTTPRoute

`func NewPatchedHTTPRoute() *PatchedHTTPRoute`

NewPatchedHTTPRoute instantiates a new PatchedHTTPRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedHTTPRouteWithDefaults

`func NewPatchedHTTPRouteWithDefaults() *PatchedHTTPRoute`

NewPatchedHTTPRouteWithDefaults instantiates a new PatchedHTTPRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedHTTPRoute) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedHTTPRoute) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedHTTPRoute) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedHTTPRoute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedHTTPRoute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedHTTPRoute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedHTTPRoute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedHTTPRoute) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *PatchedHTTPRoute) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PatchedHTTPRoute) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PatchedHTTPRoute) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PatchedHTTPRoute) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetHostnames

`func (o *PatchedHTTPRoute) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *PatchedHTTPRoute) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *PatchedHTTPRoute) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.

### HasHostnames

`func (o *PatchedHTTPRoute) HasHostnames() bool`

HasHostnames returns a boolean if a field has been set.

### GetBackendServiceName

`func (o *PatchedHTTPRoute) GetBackendServiceName() string`

GetBackendServiceName returns the BackendServiceName field if non-nil, zero value otherwise.

### GetBackendServiceNameOk

`func (o *PatchedHTTPRoute) GetBackendServiceNameOk() (*string, bool)`

GetBackendServiceNameOk returns a tuple with the BackendServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendServiceName

`func (o *PatchedHTTPRoute) SetBackendServiceName(v string)`

SetBackendServiceName sets BackendServiceName field to given value.

### HasBackendServiceName

`func (o *PatchedHTTPRoute) HasBackendServiceName() bool`

HasBackendServiceName returns a boolean if a field has been set.

### GetBackendServicePort

`func (o *PatchedHTTPRoute) GetBackendServicePort() int32`

GetBackendServicePort returns the BackendServicePort field if non-nil, zero value otherwise.

### GetBackendServicePortOk

`func (o *PatchedHTTPRoute) GetBackendServicePortOk() (*int32, bool)`

GetBackendServicePortOk returns a tuple with the BackendServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendServicePort

`func (o *PatchedHTTPRoute) SetBackendServicePort(v int32)`

SetBackendServicePort sets BackendServicePort field to given value.

### HasBackendServicePort

`func (o *PatchedHTTPRoute) HasBackendServicePort() bool`

HasBackendServicePort returns a boolean if a field has been set.

### GetBackendNamespace

`func (o *PatchedHTTPRoute) GetBackendNamespace() string`

GetBackendNamespace returns the BackendNamespace field if non-nil, zero value otherwise.

### GetBackendNamespaceOk

`func (o *PatchedHTTPRoute) GetBackendNamespaceOk() (*string, bool)`

GetBackendNamespaceOk returns a tuple with the BackendNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendNamespace

`func (o *PatchedHTTPRoute) SetBackendNamespace(v string)`

SetBackendNamespace sets BackendNamespace field to given value.

### HasBackendNamespace

`func (o *PatchedHTTPRoute) HasBackendNamespace() bool`

HasBackendNamespace returns a boolean if a field has been set.

### GetPathPrefix

`func (o *PatchedHTTPRoute) GetPathPrefix() string`

GetPathPrefix returns the PathPrefix field if non-nil, zero value otherwise.

### GetPathPrefixOk

`func (o *PatchedHTTPRoute) GetPathPrefixOk() (*string, bool)`

GetPathPrefixOk returns a tuple with the PathPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPrefix

`func (o *PatchedHTTPRoute) SetPathPrefix(v string)`

SetPathPrefix sets PathPrefix field to given value.

### HasPathPrefix

`func (o *PatchedHTTPRoute) HasPathPrefix() bool`

HasPathPrefix returns a boolean if a field has been set.

### GetEnableTls

`func (o *PatchedHTTPRoute) GetEnableTls() bool`

GetEnableTls returns the EnableTls field if non-nil, zero value otherwise.

### GetEnableTlsOk

`func (o *PatchedHTTPRoute) GetEnableTlsOk() (*bool, bool)`

GetEnableTlsOk returns a tuple with the EnableTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTls

`func (o *PatchedHTTPRoute) SetEnableTls(v bool)`

SetEnableTls sets EnableTls field to given value.

### HasEnableTls

`func (o *PatchedHTTPRoute) HasEnableTls() bool`

HasEnableTls returns a boolean if a field has been set.

### GetStatusReady

`func (o *PatchedHTTPRoute) GetStatusReady() bool`

GetStatusReady returns the StatusReady field if non-nil, zero value otherwise.

### GetStatusReadyOk

`func (o *PatchedHTTPRoute) GetStatusReadyOk() (*bool, bool)`

GetStatusReadyOk returns a tuple with the StatusReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReady

`func (o *PatchedHTTPRoute) SetStatusReady(v bool)`

SetStatusReady sets StatusReady field to given value.

### HasStatusReady

`func (o *PatchedHTTPRoute) HasStatusReady() bool`

HasStatusReady returns a boolean if a field has been set.

### SetStatusReadyNil

`func (o *PatchedHTTPRoute) SetStatusReadyNil(b bool)`

 SetStatusReadyNil sets the value for StatusReady to be an explicit nil

### UnsetStatusReady
`func (o *PatchedHTTPRoute) UnsetStatusReady()`

UnsetStatusReady ensures that no value is present for StatusReady, not even an explicit nil
### GetStatusMessage

`func (o *PatchedHTTPRoute) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *PatchedHTTPRoute) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *PatchedHTTPRoute) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *PatchedHTTPRoute) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetCreated

`func (o *PatchedHTTPRoute) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PatchedHTTPRoute) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PatchedHTTPRoute) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PatchedHTTPRoute) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *PatchedHTTPRoute) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PatchedHTTPRoute) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PatchedHTTPRoute) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PatchedHTTPRoute) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


