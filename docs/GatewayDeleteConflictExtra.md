# GatewayDeleteConflictExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | **map[string][]string** |  | 
**Preview** | Pointer to [**GatewayDeletePreview**](GatewayDeletePreview.md) | Present when the confirmation is stale. | [optional] 

## Methods

### NewGatewayDeleteConflictExtra

`func NewGatewayDeleteConflictExtra(fields map[string][]string, ) *GatewayDeleteConflictExtra`

NewGatewayDeleteConflictExtra instantiates a new GatewayDeleteConflictExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayDeleteConflictExtraWithDefaults

`func NewGatewayDeleteConflictExtraWithDefaults() *GatewayDeleteConflictExtra`

NewGatewayDeleteConflictExtraWithDefaults instantiates a new GatewayDeleteConflictExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *GatewayDeleteConflictExtra) GetFields() map[string][]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *GatewayDeleteConflictExtra) GetFieldsOk() (*map[string][]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *GatewayDeleteConflictExtra) SetFields(v map[string][]string)`

SetFields sets Fields field to given value.


### GetPreview

`func (o *GatewayDeleteConflictExtra) GetPreview() GatewayDeletePreview`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *GatewayDeleteConflictExtra) GetPreviewOk() (*GatewayDeletePreview, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *GatewayDeleteConflictExtra) SetPreview(v GatewayDeletePreview)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *GatewayDeleteConflictExtra) HasPreview() bool`

HasPreview returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


