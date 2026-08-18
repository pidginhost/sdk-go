# InboundRoute

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

## Methods

### NewInboundRoute

`func NewInboundRoute(id int32, domain int32, pattern string, mode ModeEnum, createdAt string, ) *InboundRoute`

NewInboundRoute instantiates a new InboundRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundRouteWithDefaults

`func NewInboundRouteWithDefaults() *InboundRoute`

NewInboundRouteWithDefaults instantiates a new InboundRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InboundRoute) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InboundRoute) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InboundRoute) SetId(v int32)`

SetId sets Id field to given value.


### GetDomain

`func (o *InboundRoute) GetDomain() int32`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *InboundRoute) GetDomainOk() (*int32, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *InboundRoute) SetDomain(v int32)`

SetDomain sets Domain field to given value.


### GetPattern

`func (o *InboundRoute) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *InboundRoute) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *InboundRoute) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetMode

`func (o *InboundRoute) GetMode() ModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InboundRoute) GetModeOk() (*ModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InboundRoute) SetMode(v ModeEnum)`

SetMode sets Mode field to given value.


### GetWebhookUrl

`func (o *InboundRoute) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *InboundRoute) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *InboundRoute) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *InboundRoute) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetForwardTo

`func (o *InboundRoute) GetForwardTo() string`

GetForwardTo returns the ForwardTo field if non-nil, zero value otherwise.

### GetForwardToOk

`func (o *InboundRoute) GetForwardToOk() (*string, bool)`

GetForwardToOk returns a tuple with the ForwardTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardTo

`func (o *InboundRoute) SetForwardTo(v string)`

SetForwardTo sets ForwardTo field to given value.

### HasForwardTo

`func (o *InboundRoute) HasForwardTo() bool`

HasForwardTo returns a boolean if a field has been set.

### GetActive

`func (o *InboundRoute) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *InboundRoute) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *InboundRoute) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *InboundRoute) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InboundRoute) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InboundRoute) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InboundRoute) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


