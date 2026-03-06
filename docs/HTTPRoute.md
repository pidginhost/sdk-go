# HTTPRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**Namespace** | Pointer to **string** |  | [optional] 
**Hostnames** | **[]string** | List of hostnames to route (e.g., [\&quot;example.com\&quot;, \&quot;www.example.com\&quot;]) | 
**BackendServiceName** | **string** | Name of the backend Kubernetes Service | 
**BackendServicePort** | **int32** | Port of the backend Service | 
**BackendNamespace** | Pointer to **string** | Namespace of the backend Service | [optional] [default to "default"]
**PathPrefix** | Pointer to **string** | Path prefix to match (default: /) | [optional] [default to "/"]
**EnableTls** | Pointer to **bool** | Enable TLS termination with automatic certificate issuance | [optional] [default to true]
**StatusReady** | **NullableBool** |  | [readonly] 
**StatusMessage** | **string** |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 
**Updated** | **time.Time** |  | [readonly] 

## Methods

### NewHTTPRoute

`func NewHTTPRoute(id int32, name string, hostnames []string, backendServiceName string, backendServicePort int32, statusReady NullableBool, statusMessage string, created time.Time, updated time.Time, ) *HTTPRoute`

NewHTTPRoute instantiates a new HTTPRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHTTPRouteWithDefaults

`func NewHTTPRouteWithDefaults() *HTTPRoute`

NewHTTPRouteWithDefaults instantiates a new HTTPRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HTTPRoute) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HTTPRoute) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HTTPRoute) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *HTTPRoute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HTTPRoute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HTTPRoute) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *HTTPRoute) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *HTTPRoute) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *HTTPRoute) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *HTTPRoute) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetHostnames

`func (o *HTTPRoute) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *HTTPRoute) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *HTTPRoute) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.


### GetBackendServiceName

`func (o *HTTPRoute) GetBackendServiceName() string`

GetBackendServiceName returns the BackendServiceName field if non-nil, zero value otherwise.

### GetBackendServiceNameOk

`func (o *HTTPRoute) GetBackendServiceNameOk() (*string, bool)`

GetBackendServiceNameOk returns a tuple with the BackendServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendServiceName

`func (o *HTTPRoute) SetBackendServiceName(v string)`

SetBackendServiceName sets BackendServiceName field to given value.


### GetBackendServicePort

`func (o *HTTPRoute) GetBackendServicePort() int32`

GetBackendServicePort returns the BackendServicePort field if non-nil, zero value otherwise.

### GetBackendServicePortOk

`func (o *HTTPRoute) GetBackendServicePortOk() (*int32, bool)`

GetBackendServicePortOk returns a tuple with the BackendServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendServicePort

`func (o *HTTPRoute) SetBackendServicePort(v int32)`

SetBackendServicePort sets BackendServicePort field to given value.


### GetBackendNamespace

`func (o *HTTPRoute) GetBackendNamespace() string`

GetBackendNamespace returns the BackendNamespace field if non-nil, zero value otherwise.

### GetBackendNamespaceOk

`func (o *HTTPRoute) GetBackendNamespaceOk() (*string, bool)`

GetBackendNamespaceOk returns a tuple with the BackendNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendNamespace

`func (o *HTTPRoute) SetBackendNamespace(v string)`

SetBackendNamespace sets BackendNamespace field to given value.

### HasBackendNamespace

`func (o *HTTPRoute) HasBackendNamespace() bool`

HasBackendNamespace returns a boolean if a field has been set.

### GetPathPrefix

`func (o *HTTPRoute) GetPathPrefix() string`

GetPathPrefix returns the PathPrefix field if non-nil, zero value otherwise.

### GetPathPrefixOk

`func (o *HTTPRoute) GetPathPrefixOk() (*string, bool)`

GetPathPrefixOk returns a tuple with the PathPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPrefix

`func (o *HTTPRoute) SetPathPrefix(v string)`

SetPathPrefix sets PathPrefix field to given value.

### HasPathPrefix

`func (o *HTTPRoute) HasPathPrefix() bool`

HasPathPrefix returns a boolean if a field has been set.

### GetEnableTls

`func (o *HTTPRoute) GetEnableTls() bool`

GetEnableTls returns the EnableTls field if non-nil, zero value otherwise.

### GetEnableTlsOk

`func (o *HTTPRoute) GetEnableTlsOk() (*bool, bool)`

GetEnableTlsOk returns a tuple with the EnableTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTls

`func (o *HTTPRoute) SetEnableTls(v bool)`

SetEnableTls sets EnableTls field to given value.

### HasEnableTls

`func (o *HTTPRoute) HasEnableTls() bool`

HasEnableTls returns a boolean if a field has been set.

### GetStatusReady

`func (o *HTTPRoute) GetStatusReady() bool`

GetStatusReady returns the StatusReady field if non-nil, zero value otherwise.

### GetStatusReadyOk

`func (o *HTTPRoute) GetStatusReadyOk() (*bool, bool)`

GetStatusReadyOk returns a tuple with the StatusReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReady

`func (o *HTTPRoute) SetStatusReady(v bool)`

SetStatusReady sets StatusReady field to given value.


### SetStatusReadyNil

`func (o *HTTPRoute) SetStatusReadyNil(b bool)`

 SetStatusReadyNil sets the value for StatusReady to be an explicit nil

### UnsetStatusReady
`func (o *HTTPRoute) UnsetStatusReady()`

UnsetStatusReady ensures that no value is present for StatusReady, not even an explicit nil
### GetStatusMessage

`func (o *HTTPRoute) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *HTTPRoute) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *HTTPRoute) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.


### GetCreated

`func (o *HTTPRoute) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *HTTPRoute) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *HTTPRoute) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *HTTPRoute) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *HTTPRoute) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *HTTPRoute) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


