# PatchedHTTPRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Hostnames** | Pointer to **[]string** | List of hostnames to route (e.g., [\&quot;example.com\&quot;, \&quot;www.example.com\&quot;]) | [optional] 
**BackendServiceName** | Pointer to **string** | Name of the backend Kubernetes Service | [optional] 
**BackendServicePort** | Pointer to **int32** | Port of the backend Service | [optional] 
**BackendNamespace** | Pointer to **string** | Namespace of the backend Service | [optional] [default to "default"]
**PathPrefix** | Pointer to **string** | Path prefix to match (default: /) | [optional] [default to "/"]
**EnableTls** | Pointer to **bool** | Enable TLS termination with automatic certificate issuance | [optional] [default to true]

## Methods

### NewPatchedHTTPRouteRequest

`func NewPatchedHTTPRouteRequest() *PatchedHTTPRouteRequest`

NewPatchedHTTPRouteRequest instantiates a new PatchedHTTPRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedHTTPRouteRequestWithDefaults

`func NewPatchedHTTPRouteRequestWithDefaults() *PatchedHTTPRouteRequest`

NewPatchedHTTPRouteRequestWithDefaults instantiates a new PatchedHTTPRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedHTTPRouteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedHTTPRouteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedHTTPRouteRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedHTTPRouteRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *PatchedHTTPRouteRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PatchedHTTPRouteRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PatchedHTTPRouteRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PatchedHTTPRouteRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetHostnames

`func (o *PatchedHTTPRouteRequest) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *PatchedHTTPRouteRequest) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *PatchedHTTPRouteRequest) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.

### HasHostnames

`func (o *PatchedHTTPRouteRequest) HasHostnames() bool`

HasHostnames returns a boolean if a field has been set.

### GetBackendServiceName

`func (o *PatchedHTTPRouteRequest) GetBackendServiceName() string`

GetBackendServiceName returns the BackendServiceName field if non-nil, zero value otherwise.

### GetBackendServiceNameOk

`func (o *PatchedHTTPRouteRequest) GetBackendServiceNameOk() (*string, bool)`

GetBackendServiceNameOk returns a tuple with the BackendServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendServiceName

`func (o *PatchedHTTPRouteRequest) SetBackendServiceName(v string)`

SetBackendServiceName sets BackendServiceName field to given value.

### HasBackendServiceName

`func (o *PatchedHTTPRouteRequest) HasBackendServiceName() bool`

HasBackendServiceName returns a boolean if a field has been set.

### GetBackendServicePort

`func (o *PatchedHTTPRouteRequest) GetBackendServicePort() int32`

GetBackendServicePort returns the BackendServicePort field if non-nil, zero value otherwise.

### GetBackendServicePortOk

`func (o *PatchedHTTPRouteRequest) GetBackendServicePortOk() (*int32, bool)`

GetBackendServicePortOk returns a tuple with the BackendServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendServicePort

`func (o *PatchedHTTPRouteRequest) SetBackendServicePort(v int32)`

SetBackendServicePort sets BackendServicePort field to given value.

### HasBackendServicePort

`func (o *PatchedHTTPRouteRequest) HasBackendServicePort() bool`

HasBackendServicePort returns a boolean if a field has been set.

### GetBackendNamespace

`func (o *PatchedHTTPRouteRequest) GetBackendNamespace() string`

GetBackendNamespace returns the BackendNamespace field if non-nil, zero value otherwise.

### GetBackendNamespaceOk

`func (o *PatchedHTTPRouteRequest) GetBackendNamespaceOk() (*string, bool)`

GetBackendNamespaceOk returns a tuple with the BackendNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendNamespace

`func (o *PatchedHTTPRouteRequest) SetBackendNamespace(v string)`

SetBackendNamespace sets BackendNamespace field to given value.

### HasBackendNamespace

`func (o *PatchedHTTPRouteRequest) HasBackendNamespace() bool`

HasBackendNamespace returns a boolean if a field has been set.

### GetPathPrefix

`func (o *PatchedHTTPRouteRequest) GetPathPrefix() string`

GetPathPrefix returns the PathPrefix field if non-nil, zero value otherwise.

### GetPathPrefixOk

`func (o *PatchedHTTPRouteRequest) GetPathPrefixOk() (*string, bool)`

GetPathPrefixOk returns a tuple with the PathPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPrefix

`func (o *PatchedHTTPRouteRequest) SetPathPrefix(v string)`

SetPathPrefix sets PathPrefix field to given value.

### HasPathPrefix

`func (o *PatchedHTTPRouteRequest) HasPathPrefix() bool`

HasPathPrefix returns a boolean if a field has been set.

### GetEnableTls

`func (o *PatchedHTTPRouteRequest) GetEnableTls() bool`

GetEnableTls returns the EnableTls field if non-nil, zero value otherwise.

### GetEnableTlsOk

`func (o *PatchedHTTPRouteRequest) GetEnableTlsOk() (*bool, bool)`

GetEnableTlsOk returns a tuple with the EnableTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTls

`func (o *PatchedHTTPRouteRequest) SetEnableTls(v bool)`

SetEnableTls sets EnableTls field to given value.

### HasEnableTls

`func (o *PatchedHTTPRouteRequest) HasEnableTls() bool`

HasEnableTls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


