# PatchedInboundRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Domain** | Pointer to **int32** |  | [optional] [readonly] 
**Pattern** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to [**ModeEnum**](ModeEnum.md) |  | [optional] 
**WebhookUrl** | Pointer to **string** |  | [optional] 
**ForwardTo** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPatchedInboundRoute

`func NewPatchedInboundRoute() *PatchedInboundRoute`

NewPatchedInboundRoute instantiates a new PatchedInboundRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedInboundRouteWithDefaults

`func NewPatchedInboundRouteWithDefaults() *PatchedInboundRoute`

NewPatchedInboundRouteWithDefaults instantiates a new PatchedInboundRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedInboundRoute) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedInboundRoute) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedInboundRoute) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedInboundRoute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDomain

`func (o *PatchedInboundRoute) GetDomain() int32`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PatchedInboundRoute) GetDomainOk() (*int32, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PatchedInboundRoute) SetDomain(v int32)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PatchedInboundRoute) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetPattern

`func (o *PatchedInboundRoute) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *PatchedInboundRoute) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *PatchedInboundRoute) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *PatchedInboundRoute) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetMode

`func (o *PatchedInboundRoute) GetMode() ModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PatchedInboundRoute) GetModeOk() (*ModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PatchedInboundRoute) SetMode(v ModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PatchedInboundRoute) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *PatchedInboundRoute) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *PatchedInboundRoute) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *PatchedInboundRoute) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *PatchedInboundRoute) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetForwardTo

`func (o *PatchedInboundRoute) GetForwardTo() string`

GetForwardTo returns the ForwardTo field if non-nil, zero value otherwise.

### GetForwardToOk

`func (o *PatchedInboundRoute) GetForwardToOk() (*string, bool)`

GetForwardToOk returns a tuple with the ForwardTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardTo

`func (o *PatchedInboundRoute) SetForwardTo(v string)`

SetForwardTo sets ForwardTo field to given value.

### HasForwardTo

`func (o *PatchedInboundRoute) HasForwardTo() bool`

HasForwardTo returns a boolean if a field has been set.

### GetActive

`func (o *PatchedInboundRoute) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedInboundRoute) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedInboundRoute) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedInboundRoute) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedInboundRoute) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedInboundRoute) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedInboundRoute) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedInboundRoute) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


