# HTTPRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Namespace** | Pointer to **string** |  | [optional] 
**Hostnames** | **[]string** | List of hostnames to route (e.g., [\&quot;example.com\&quot;, \&quot;www.example.com\&quot;]) | 
**BackendServiceName** | **string** | Name of the backend Kubernetes Service | 
**BackendServicePort** | **int32** | Port of the backend Service | 
**BackendNamespace** | Pointer to **string** | Namespace of the backend Service | [optional] [default to "default"]
**PathPrefix** | Pointer to **string** | Path prefix to match (default: /) | [optional] [default to "/"]
**EnableTls** | Pointer to **bool** | Enable TLS termination with automatic certificate issuance | [optional] [default to true]

## Methods

### NewHTTPRouteRequest

`func NewHTTPRouteRequest(name string, hostnames []string, backendServiceName string, backendServicePort int32, ) *HTTPRouteRequest`

NewHTTPRouteRequest instantiates a new HTTPRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHTTPRouteRequestWithDefaults

`func NewHTTPRouteRequestWithDefaults() *HTTPRouteRequest`

NewHTTPRouteRequestWithDefaults instantiates a new HTTPRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HTTPRouteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HTTPRouteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HTTPRouteRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *HTTPRouteRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *HTTPRouteRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *HTTPRouteRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *HTTPRouteRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetHostnames

`func (o *HTTPRouteRequest) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *HTTPRouteRequest) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *HTTPRouteRequest) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.


### GetBackendServiceName

`func (o *HTTPRouteRequest) GetBackendServiceName() string`

GetBackendServiceName returns the BackendServiceName field if non-nil, zero value otherwise.

### GetBackendServiceNameOk

`func (o *HTTPRouteRequest) GetBackendServiceNameOk() (*string, bool)`

GetBackendServiceNameOk returns a tuple with the BackendServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendServiceName

`func (o *HTTPRouteRequest) SetBackendServiceName(v string)`

SetBackendServiceName sets BackendServiceName field to given value.


### GetBackendServicePort

`func (o *HTTPRouteRequest) GetBackendServicePort() int32`

GetBackendServicePort returns the BackendServicePort field if non-nil, zero value otherwise.

### GetBackendServicePortOk

`func (o *HTTPRouteRequest) GetBackendServicePortOk() (*int32, bool)`

GetBackendServicePortOk returns a tuple with the BackendServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendServicePort

`func (o *HTTPRouteRequest) SetBackendServicePort(v int32)`

SetBackendServicePort sets BackendServicePort field to given value.


### GetBackendNamespace

`func (o *HTTPRouteRequest) GetBackendNamespace() string`

GetBackendNamespace returns the BackendNamespace field if non-nil, zero value otherwise.

### GetBackendNamespaceOk

`func (o *HTTPRouteRequest) GetBackendNamespaceOk() (*string, bool)`

GetBackendNamespaceOk returns a tuple with the BackendNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendNamespace

`func (o *HTTPRouteRequest) SetBackendNamespace(v string)`

SetBackendNamespace sets BackendNamespace field to given value.

### HasBackendNamespace

`func (o *HTTPRouteRequest) HasBackendNamespace() bool`

HasBackendNamespace returns a boolean if a field has been set.

### GetPathPrefix

`func (o *HTTPRouteRequest) GetPathPrefix() string`

GetPathPrefix returns the PathPrefix field if non-nil, zero value otherwise.

### GetPathPrefixOk

`func (o *HTTPRouteRequest) GetPathPrefixOk() (*string, bool)`

GetPathPrefixOk returns a tuple with the PathPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPrefix

`func (o *HTTPRouteRequest) SetPathPrefix(v string)`

SetPathPrefix sets PathPrefix field to given value.

### HasPathPrefix

`func (o *HTTPRouteRequest) HasPathPrefix() bool`

HasPathPrefix returns a boolean if a field has been set.

### GetEnableTls

`func (o *HTTPRouteRequest) GetEnableTls() bool`

GetEnableTls returns the EnableTls field if non-nil, zero value otherwise.

### GetEnableTlsOk

`func (o *HTTPRouteRequest) GetEnableTlsOk() (*bool, bool)`

GetEnableTlsOk returns a tuple with the EnableTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTls

`func (o *HTTPRouteRequest) SetEnableTls(v bool)`

SetEnableTls sets EnableTls field to given value.

### HasEnableTls

`func (o *HTTPRouteRequest) HasEnableTls() bool`

HasEnableTls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


