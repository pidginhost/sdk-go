# InboundRouteWriteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Domain** | **int32** |  | [readonly] 
**Pattern** | **string** |  | 
**Mode** | [**ModeEnum**](ModeEnum.md) |  | 
**WebhookUrl** | Pointer to **string** |  | [optional] 
**ForwardTo** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**CreatedAt** | **string** |  | [readonly] 
**WebhookSecret** | Pointer to **string** |  | [optional] 

## Methods

### NewInboundRouteWriteResponse

`func NewInboundRouteWriteResponse(id int32, domain int32, pattern string, mode ModeEnum, createdAt string, ) *InboundRouteWriteResponse`

NewInboundRouteWriteResponse instantiates a new InboundRouteWriteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundRouteWriteResponseWithDefaults

`func NewInboundRouteWriteResponseWithDefaults() *InboundRouteWriteResponse`

NewInboundRouteWriteResponseWithDefaults instantiates a new InboundRouteWriteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InboundRouteWriteResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InboundRouteWriteResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InboundRouteWriteResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetDomain

`func (o *InboundRouteWriteResponse) GetDomain() int32`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *InboundRouteWriteResponse) GetDomainOk() (*int32, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *InboundRouteWriteResponse) SetDomain(v int32)`

SetDomain sets Domain field to given value.


### GetPattern

`func (o *InboundRouteWriteResponse) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *InboundRouteWriteResponse) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *InboundRouteWriteResponse) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetMode

`func (o *InboundRouteWriteResponse) GetMode() ModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InboundRouteWriteResponse) GetModeOk() (*ModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InboundRouteWriteResponse) SetMode(v ModeEnum)`

SetMode sets Mode field to given value.


### GetWebhookUrl

`func (o *InboundRouteWriteResponse) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *InboundRouteWriteResponse) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *InboundRouteWriteResponse) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *InboundRouteWriteResponse) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetForwardTo

`func (o *InboundRouteWriteResponse) GetForwardTo() string`

GetForwardTo returns the ForwardTo field if non-nil, zero value otherwise.

### GetForwardToOk

`func (o *InboundRouteWriteResponse) GetForwardToOk() (*string, bool)`

GetForwardToOk returns a tuple with the ForwardTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardTo

`func (o *InboundRouteWriteResponse) SetForwardTo(v string)`

SetForwardTo sets ForwardTo field to given value.

### HasForwardTo

`func (o *InboundRouteWriteResponse) HasForwardTo() bool`

HasForwardTo returns a boolean if a field has been set.

### GetActive

`func (o *InboundRouteWriteResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *InboundRouteWriteResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *InboundRouteWriteResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *InboundRouteWriteResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InboundRouteWriteResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InboundRouteWriteResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InboundRouteWriteResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetWebhookSecret

`func (o *InboundRouteWriteResponse) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *InboundRouteWriteResponse) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *InboundRouteWriteResponse) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.

### HasWebhookSecret

`func (o *InboundRouteWriteResponse) HasWebhookSecret() bool`

HasWebhookSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


