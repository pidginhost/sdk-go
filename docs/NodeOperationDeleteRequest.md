# NodeOperationDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalDataLossAccepted** | Pointer to **bool** | Acknowledge that data kept on the node itself is destroyed. The drain always deletes emptyDir. | [optional] [default to false]
**ConfirmHostname** | Pointer to **string** | The node&#39;s hostname, typed back exactly. | [optional] [default to ""]

## Methods

### NewNodeOperationDeleteRequest

`func NewNodeOperationDeleteRequest() *NodeOperationDeleteRequest`

NewNodeOperationDeleteRequest instantiates a new NodeOperationDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeOperationDeleteRequestWithDefaults

`func NewNodeOperationDeleteRequestWithDefaults() *NodeOperationDeleteRequest`

NewNodeOperationDeleteRequestWithDefaults instantiates a new NodeOperationDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalDataLossAccepted

`func (o *NodeOperationDeleteRequest) GetLocalDataLossAccepted() bool`

GetLocalDataLossAccepted returns the LocalDataLossAccepted field if non-nil, zero value otherwise.

### GetLocalDataLossAcceptedOk

`func (o *NodeOperationDeleteRequest) GetLocalDataLossAcceptedOk() (*bool, bool)`

GetLocalDataLossAcceptedOk returns a tuple with the LocalDataLossAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDataLossAccepted

`func (o *NodeOperationDeleteRequest) SetLocalDataLossAccepted(v bool)`

SetLocalDataLossAccepted sets LocalDataLossAccepted field to given value.

### HasLocalDataLossAccepted

`func (o *NodeOperationDeleteRequest) HasLocalDataLossAccepted() bool`

HasLocalDataLossAccepted returns a boolean if a field has been set.

### GetConfirmHostname

`func (o *NodeOperationDeleteRequest) GetConfirmHostname() string`

GetConfirmHostname returns the ConfirmHostname field if non-nil, zero value otherwise.

### GetConfirmHostnameOk

`func (o *NodeOperationDeleteRequest) GetConfirmHostnameOk() (*string, bool)`

GetConfirmHostnameOk returns a tuple with the ConfirmHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmHostname

`func (o *NodeOperationDeleteRequest) SetConfirmHostname(v string)`

SetConfirmHostname sets ConfirmHostname field to given value.

### HasConfirmHostname

`func (o *NodeOperationDeleteRequest) HasConfirmHostname() bool`

HasConfirmHostname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


