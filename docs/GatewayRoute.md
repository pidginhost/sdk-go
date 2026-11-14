# GatewayRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Hostnames** | **[]string** | Empty matches every host. | 
**PathPrefix** | **string** |  | [readonly] 
**Pool** | **int32** |  | 
**RedirectHttps** | **bool** | Always false: HTTPS redirects are not available yet. | [readonly] 
**Enabled** | **bool** |  | [readonly] 
**ChangedRevision** | **int32** |  | [readonly] 
**Pending** | **bool** |  | [readonly] 

## Methods

### NewGatewayRoute

`func NewGatewayRoute(id int32, hostnames []string, pathPrefix string, pool int32, redirectHttps bool, enabled bool, changedRevision int32, pending bool, ) *GatewayRoute`

NewGatewayRoute instantiates a new GatewayRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayRouteWithDefaults

`func NewGatewayRouteWithDefaults() *GatewayRoute`

NewGatewayRouteWithDefaults instantiates a new GatewayRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GatewayRoute) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayRoute) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayRoute) SetId(v int32)`

SetId sets Id field to given value.


### GetHostnames

`func (o *GatewayRoute) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *GatewayRoute) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *GatewayRoute) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.


### GetPathPrefix

`func (o *GatewayRoute) GetPathPrefix() string`

GetPathPrefix returns the PathPrefix field if non-nil, zero value otherwise.

### GetPathPrefixOk

`func (o *GatewayRoute) GetPathPrefixOk() (*string, bool)`

GetPathPrefixOk returns a tuple with the PathPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPrefix

`func (o *GatewayRoute) SetPathPrefix(v string)`

SetPathPrefix sets PathPrefix field to given value.


### GetPool

`func (o *GatewayRoute) GetPool() int32`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *GatewayRoute) GetPoolOk() (*int32, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *GatewayRoute) SetPool(v int32)`

SetPool sets Pool field to given value.


### GetRedirectHttps

`func (o *GatewayRoute) GetRedirectHttps() bool`

GetRedirectHttps returns the RedirectHttps field if non-nil, zero value otherwise.

### GetRedirectHttpsOk

`func (o *GatewayRoute) GetRedirectHttpsOk() (*bool, bool)`

GetRedirectHttpsOk returns a tuple with the RedirectHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectHttps

`func (o *GatewayRoute) SetRedirectHttps(v bool)`

SetRedirectHttps sets RedirectHttps field to given value.


### GetEnabled

`func (o *GatewayRoute) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GatewayRoute) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GatewayRoute) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetChangedRevision

`func (o *GatewayRoute) GetChangedRevision() int32`

GetChangedRevision returns the ChangedRevision field if non-nil, zero value otherwise.

### GetChangedRevisionOk

`func (o *GatewayRoute) GetChangedRevisionOk() (*int32, bool)`

GetChangedRevisionOk returns a tuple with the ChangedRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedRevision

`func (o *GatewayRoute) SetChangedRevision(v int32)`

SetChangedRevision sets ChangedRevision field to given value.


### GetPending

`func (o *GatewayRoute) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *GatewayRoute) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *GatewayRoute) SetPending(v bool)`

SetPending sets Pending field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


