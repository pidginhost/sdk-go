# ResourcePoolDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalDataLossAccepted** | Pointer to **bool** | Acknowledge that data kept on the pool&#39;s nodes is destroyed. Not needed for an empty pool. | [optional] [default to false]

## Methods

### NewResourcePoolDeleteRequest

`func NewResourcePoolDeleteRequest() *ResourcePoolDeleteRequest`

NewResourcePoolDeleteRequest instantiates a new ResourcePoolDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePoolDeleteRequestWithDefaults

`func NewResourcePoolDeleteRequestWithDefaults() *ResourcePoolDeleteRequest`

NewResourcePoolDeleteRequestWithDefaults instantiates a new ResourcePoolDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalDataLossAccepted

`func (o *ResourcePoolDeleteRequest) GetLocalDataLossAccepted() bool`

GetLocalDataLossAccepted returns the LocalDataLossAccepted field if non-nil, zero value otherwise.

### GetLocalDataLossAcceptedOk

`func (o *ResourcePoolDeleteRequest) GetLocalDataLossAcceptedOk() (*bool, bool)`

GetLocalDataLossAcceptedOk returns a tuple with the LocalDataLossAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDataLossAccepted

`func (o *ResourcePoolDeleteRequest) SetLocalDataLossAccepted(v bool)`

SetLocalDataLossAccepted sets LocalDataLossAccepted field to given value.

### HasLocalDataLossAccepted

`func (o *ResourcePoolDeleteRequest) HasLocalDataLossAccepted() bool`

HasLocalDataLossAccepted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


