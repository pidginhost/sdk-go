# UDPRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Namespace** | Pointer to **string** |  | [optional] 
**Port** | **int32** | External port to expose | 
**BackendServiceName** | **string** | Name of the backend Kubernetes Service | 
**BackendServicePort** | **int32** | Port of the backend Service | 
**BackendNamespace** | Pointer to **string** | Namespace of the backend Service | [optional] [default to "default"]

## Methods

### NewUDPRouteRequest

`func NewUDPRouteRequest(name string, port int32, backendServiceName string, backendServicePort int32, ) *UDPRouteRequest`

NewUDPRouteRequest instantiates a new UDPRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUDPRouteRequestWithDefaults

`func NewUDPRouteRequestWithDefaults() *UDPRouteRequest`

NewUDPRouteRequestWithDefaults instantiates a new UDPRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UDPRouteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UDPRouteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UDPRouteRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *UDPRouteRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *UDPRouteRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *UDPRouteRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *UDPRouteRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPort

`func (o *UDPRouteRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UDPRouteRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UDPRouteRequest) SetPort(v int32)`

SetPort sets Port field to given value.


### GetBackendServiceName

`func (o *UDPRouteRequest) GetBackendServiceName() string`

GetBackendServiceName returns the BackendServiceName field if non-nil, zero value otherwise.

### GetBackendServiceNameOk

`func (o *UDPRouteRequest) GetBackendServiceNameOk() (*string, bool)`

GetBackendServiceNameOk returns a tuple with the BackendServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendServiceName

`func (o *UDPRouteRequest) SetBackendServiceName(v string)`

SetBackendServiceName sets BackendServiceName field to given value.


### GetBackendServicePort

`func (o *UDPRouteRequest) GetBackendServicePort() int32`

GetBackendServicePort returns the BackendServicePort field if non-nil, zero value otherwise.

### GetBackendServicePortOk

`func (o *UDPRouteRequest) GetBackendServicePortOk() (*int32, bool)`

GetBackendServicePortOk returns a tuple with the BackendServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendServicePort

`func (o *UDPRouteRequest) SetBackendServicePort(v int32)`

SetBackendServicePort sets BackendServicePort field to given value.


### GetBackendNamespace

`func (o *UDPRouteRequest) GetBackendNamespace() string`

GetBackendNamespace returns the BackendNamespace field if non-nil, zero value otherwise.

### GetBackendNamespaceOk

`func (o *UDPRouteRequest) GetBackendNamespaceOk() (*string, bool)`

GetBackendNamespaceOk returns a tuple with the BackendNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendNamespace

`func (o *UDPRouteRequest) SetBackendNamespace(v string)`

SetBackendNamespace sets BackendNamespace field to given value.

### HasBackendNamespace

`func (o *UDPRouteRequest) HasBackendNamespace() bool`

HasBackendNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


