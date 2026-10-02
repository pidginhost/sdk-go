# ClusterEncryptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | [**EncryptionModeEnum**](EncryptionModeEnum.md) | Target encryption mode: no encryption, or WireGuard.  * &#x60;none&#x60; - none * &#x60;wireguard&#x60; - wireguard | 
**AcknowledgeWorkloadRestart** | Pointer to **bool** | Confirms the caller accepts that workloads must be restarted after the change. | [optional] [default to false]

## Methods

### NewClusterEncryptionRequest

`func NewClusterEncryptionRequest(mode EncryptionModeEnum, ) *ClusterEncryptionRequest`

NewClusterEncryptionRequest instantiates a new ClusterEncryptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterEncryptionRequestWithDefaults

`func NewClusterEncryptionRequestWithDefaults() *ClusterEncryptionRequest`

NewClusterEncryptionRequestWithDefaults instantiates a new ClusterEncryptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *ClusterEncryptionRequest) GetMode() EncryptionModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ClusterEncryptionRequest) GetModeOk() (*EncryptionModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ClusterEncryptionRequest) SetMode(v EncryptionModeEnum)`

SetMode sets Mode field to given value.


### GetAcknowledgeWorkloadRestart

`func (o *ClusterEncryptionRequest) GetAcknowledgeWorkloadRestart() bool`

GetAcknowledgeWorkloadRestart returns the AcknowledgeWorkloadRestart field if non-nil, zero value otherwise.

### GetAcknowledgeWorkloadRestartOk

`func (o *ClusterEncryptionRequest) GetAcknowledgeWorkloadRestartOk() (*bool, bool)`

GetAcknowledgeWorkloadRestartOk returns a tuple with the AcknowledgeWorkloadRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgeWorkloadRestart

`func (o *ClusterEncryptionRequest) SetAcknowledgeWorkloadRestart(v bool)`

SetAcknowledgeWorkloadRestart sets AcknowledgeWorkloadRestart field to given value.

### HasAcknowledgeWorkloadRestart

`func (o *ClusterEncryptionRequest) HasAcknowledgeWorkloadRestart() bool`

HasAcknowledgeWorkloadRestart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


