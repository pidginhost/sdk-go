# UDPRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**Namespace** | Pointer to **string** |  | [optional] 
**Port** | **int32** | External port to expose | 
**BackendServiceName** | **string** | Name of the backend Kubernetes Service | 
**BackendServicePort** | **int32** | Port of the backend Service | 
**BackendNamespace** | Pointer to **string** | Namespace of the backend Service | [optional] [default to "default"]
**StatusReady** | **NullableBool** |  | [readonly] 
**StatusMessage** | **string** |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 
**Updated** | **time.Time** |  | [readonly] 

## Methods

### NewUDPRoute

`func NewUDPRoute(id int32, name string, port int32, backendServiceName string, backendServicePort int32, statusReady NullableBool, statusMessage string, created time.Time, updated time.Time, ) *UDPRoute`

NewUDPRoute instantiates a new UDPRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUDPRouteWithDefaults

`func NewUDPRouteWithDefaults() *UDPRoute`

NewUDPRouteWithDefaults instantiates a new UDPRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UDPRoute) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UDPRoute) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UDPRoute) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *UDPRoute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UDPRoute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UDPRoute) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *UDPRoute) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *UDPRoute) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *UDPRoute) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *UDPRoute) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPort

`func (o *UDPRoute) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UDPRoute) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UDPRoute) SetPort(v int32)`

SetPort sets Port field to given value.


### GetBackendServiceName

`func (o *UDPRoute) GetBackendServiceName() string`

GetBackendServiceName returns the BackendServiceName field if non-nil, zero value otherwise.

### GetBackendServiceNameOk

`func (o *UDPRoute) GetBackendServiceNameOk() (*string, bool)`

GetBackendServiceNameOk returns a tuple with the BackendServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendServiceName

`func (o *UDPRoute) SetBackendServiceName(v string)`

SetBackendServiceName sets BackendServiceName field to given value.


### GetBackendServicePort

`func (o *UDPRoute) GetBackendServicePort() int32`

GetBackendServicePort returns the BackendServicePort field if non-nil, zero value otherwise.

### GetBackendServicePortOk

`func (o *UDPRoute) GetBackendServicePortOk() (*int32, bool)`

GetBackendServicePortOk returns a tuple with the BackendServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendServicePort

`func (o *UDPRoute) SetBackendServicePort(v int32)`

SetBackendServicePort sets BackendServicePort field to given value.


### GetBackendNamespace

`func (o *UDPRoute) GetBackendNamespace() string`

GetBackendNamespace returns the BackendNamespace field if non-nil, zero value otherwise.

### GetBackendNamespaceOk

`func (o *UDPRoute) GetBackendNamespaceOk() (*string, bool)`

GetBackendNamespaceOk returns a tuple with the BackendNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendNamespace

`func (o *UDPRoute) SetBackendNamespace(v string)`

SetBackendNamespace sets BackendNamespace field to given value.

### HasBackendNamespace

`func (o *UDPRoute) HasBackendNamespace() bool`

HasBackendNamespace returns a boolean if a field has been set.

### GetStatusReady

`func (o *UDPRoute) GetStatusReady() bool`

GetStatusReady returns the StatusReady field if non-nil, zero value otherwise.

### GetStatusReadyOk

`func (o *UDPRoute) GetStatusReadyOk() (*bool, bool)`

GetStatusReadyOk returns a tuple with the StatusReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReady

`func (o *UDPRoute) SetStatusReady(v bool)`

SetStatusReady sets StatusReady field to given value.


### SetStatusReadyNil

`func (o *UDPRoute) SetStatusReadyNil(b bool)`

 SetStatusReadyNil sets the value for StatusReady to be an explicit nil

### UnsetStatusReady
`func (o *UDPRoute) UnsetStatusReady()`

UnsetStatusReady ensures that no value is present for StatusReady, not even an explicit nil
### GetStatusMessage

`func (o *UDPRoute) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *UDPRoute) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *UDPRoute) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.


### GetCreated

`func (o *UDPRoute) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UDPRoute) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UDPRoute) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *UDPRoute) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *UDPRoute) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *UDPRoute) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


