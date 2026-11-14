# GatewayLBListener

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Protocol** | [**GatewayForwardProtocolEnum**](GatewayForwardProtocolEnum.md) |  | [readonly] 
**Port** | **int32** |  | 
**Mode** | [**GatewayLBListenerModeEnum**](GatewayLBListenerModeEnum.md) |  | 
**Pool** | **NullableInt32** | The backend pool of a tcp or udp listener; null for http. | 
**AllowedSources** | **[]string** | Empty means any source. | 
**Enabled** | **bool** |  | [readonly] 
**Comment** | **string** |  | [readonly] 
**ChangedRevision** | **int32** |  | [readonly] 
**Pending** | **bool** |  | [readonly] 

## Methods

### NewGatewayLBListener

`func NewGatewayLBListener(id int32, protocol GatewayForwardProtocolEnum, port int32, mode GatewayLBListenerModeEnum, pool NullableInt32, allowedSources []string, enabled bool, comment string, changedRevision int32, pending bool, ) *GatewayLBListener`

NewGatewayLBListener instantiates a new GatewayLBListener object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayLBListenerWithDefaults

`func NewGatewayLBListenerWithDefaults() *GatewayLBListener`

NewGatewayLBListenerWithDefaults instantiates a new GatewayLBListener object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GatewayLBListener) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayLBListener) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayLBListener) SetId(v int32)`

SetId sets Id field to given value.


### GetProtocol

`func (o *GatewayLBListener) GetProtocol() GatewayForwardProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GatewayLBListener) GetProtocolOk() (*GatewayForwardProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GatewayLBListener) SetProtocol(v GatewayForwardProtocolEnum)`

SetProtocol sets Protocol field to given value.


### GetPort

`func (o *GatewayLBListener) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GatewayLBListener) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GatewayLBListener) SetPort(v int32)`

SetPort sets Port field to given value.


### GetMode

`func (o *GatewayLBListener) GetMode() GatewayLBListenerModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GatewayLBListener) GetModeOk() (*GatewayLBListenerModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GatewayLBListener) SetMode(v GatewayLBListenerModeEnum)`

SetMode sets Mode field to given value.


### GetPool

`func (o *GatewayLBListener) GetPool() int32`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *GatewayLBListener) GetPoolOk() (*int32, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *GatewayLBListener) SetPool(v int32)`

SetPool sets Pool field to given value.


### SetPoolNil

`func (o *GatewayLBListener) SetPoolNil(b bool)`

 SetPoolNil sets the value for Pool to be an explicit nil

### UnsetPool
`func (o *GatewayLBListener) UnsetPool()`

UnsetPool ensures that no value is present for Pool, not even an explicit nil
### GetAllowedSources

`func (o *GatewayLBListener) GetAllowedSources() []string`

GetAllowedSources returns the AllowedSources field if non-nil, zero value otherwise.

### GetAllowedSourcesOk

`func (o *GatewayLBListener) GetAllowedSourcesOk() (*[]string, bool)`

GetAllowedSourcesOk returns a tuple with the AllowedSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSources

`func (o *GatewayLBListener) SetAllowedSources(v []string)`

SetAllowedSources sets AllowedSources field to given value.


### GetEnabled

`func (o *GatewayLBListener) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GatewayLBListener) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GatewayLBListener) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetComment

`func (o *GatewayLBListener) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GatewayLBListener) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GatewayLBListener) SetComment(v string)`

SetComment sets Comment field to given value.


### GetChangedRevision

`func (o *GatewayLBListener) GetChangedRevision() int32`

GetChangedRevision returns the ChangedRevision field if non-nil, zero value otherwise.

### GetChangedRevisionOk

`func (o *GatewayLBListener) GetChangedRevisionOk() (*int32, bool)`

GetChangedRevisionOk returns a tuple with the ChangedRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedRevision

`func (o *GatewayLBListener) SetChangedRevision(v int32)`

SetChangedRevision sets ChangedRevision field to given value.


### GetPending

`func (o *GatewayLBListener) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *GatewayLBListener) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *GatewayLBListener) SetPending(v bool)`

SetPending sets Pending field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


