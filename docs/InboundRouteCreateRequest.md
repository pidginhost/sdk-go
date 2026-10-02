# InboundRouteCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pattern** | **string** |  | 
**Mode** | [**ModeEnum**](ModeEnum.md) |  | 
**WebhookUrl** | Pointer to **string** |  | [optional] 
**ForwardTo** | Pointer to **string** |  | [optional] 

## Methods

### NewInboundRouteCreateRequest

`func NewInboundRouteCreateRequest(pattern string, mode ModeEnum, ) *InboundRouteCreateRequest`

NewInboundRouteCreateRequest instantiates a new InboundRouteCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundRouteCreateRequestWithDefaults

`func NewInboundRouteCreateRequestWithDefaults() *InboundRouteCreateRequest`

NewInboundRouteCreateRequestWithDefaults instantiates a new InboundRouteCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPattern

`func (o *InboundRouteCreateRequest) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *InboundRouteCreateRequest) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *InboundRouteCreateRequest) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetMode

`func (o *InboundRouteCreateRequest) GetMode() ModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InboundRouteCreateRequest) GetModeOk() (*ModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InboundRouteCreateRequest) SetMode(v ModeEnum)`

SetMode sets Mode field to given value.


### GetWebhookUrl

`func (o *InboundRouteCreateRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *InboundRouteCreateRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *InboundRouteCreateRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *InboundRouteCreateRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetForwardTo

`func (o *InboundRouteCreateRequest) GetForwardTo() string`

GetForwardTo returns the ForwardTo field if non-nil, zero value otherwise.

### GetForwardToOk

`func (o *InboundRouteCreateRequest) GetForwardToOk() (*string, bool)`

GetForwardToOk returns a tuple with the ForwardTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardTo

`func (o *InboundRouteCreateRequest) SetForwardTo(v string)`

SetForwardTo sets ForwardTo field to given value.

### HasForwardTo

`func (o *InboundRouteCreateRequest) HasForwardTo() bool`

HasForwardTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


