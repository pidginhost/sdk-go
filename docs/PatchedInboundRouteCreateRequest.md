# PatchedInboundRouteCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pattern** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to [**ModeEnum**](ModeEnum.md) |  | [optional] 
**WebhookUrl** | Pointer to **string** |  | [optional] 
**ForwardTo** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedInboundRouteCreateRequest

`func NewPatchedInboundRouteCreateRequest() *PatchedInboundRouteCreateRequest`

NewPatchedInboundRouteCreateRequest instantiates a new PatchedInboundRouteCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedInboundRouteCreateRequestWithDefaults

`func NewPatchedInboundRouteCreateRequestWithDefaults() *PatchedInboundRouteCreateRequest`

NewPatchedInboundRouteCreateRequestWithDefaults instantiates a new PatchedInboundRouteCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPattern

`func (o *PatchedInboundRouteCreateRequest) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *PatchedInboundRouteCreateRequest) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *PatchedInboundRouteCreateRequest) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *PatchedInboundRouteCreateRequest) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetMode

`func (o *PatchedInboundRouteCreateRequest) GetMode() ModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PatchedInboundRouteCreateRequest) GetModeOk() (*ModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PatchedInboundRouteCreateRequest) SetMode(v ModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PatchedInboundRouteCreateRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *PatchedInboundRouteCreateRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *PatchedInboundRouteCreateRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *PatchedInboundRouteCreateRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *PatchedInboundRouteCreateRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetForwardTo

`func (o *PatchedInboundRouteCreateRequest) GetForwardTo() string`

GetForwardTo returns the ForwardTo field if non-nil, zero value otherwise.

### GetForwardToOk

`func (o *PatchedInboundRouteCreateRequest) GetForwardToOk() (*string, bool)`

GetForwardToOk returns a tuple with the ForwardTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardTo

`func (o *PatchedInboundRouteCreateRequest) SetForwardTo(v string)`

SetForwardTo sets ForwardTo field to given value.

### HasForwardTo

`func (o *PatchedInboundRouteCreateRequest) HasForwardTo() bool`

HasForwardTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


