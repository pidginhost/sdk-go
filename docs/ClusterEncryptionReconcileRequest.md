# ClusterEncryptionReconcileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | [**EncryptionModeEnum**](EncryptionModeEnum.md) | Target encryption mode: no encryption, or WireGuard.  * &#x60;none&#x60; - none * &#x60;wireguard&#x60; - wireguard | 
**AcknowledgeWorkloadRestart** | Pointer to **bool** | Confirms the caller accepts that workloads must be restarted after the change. | [optional] [default to false]
**OverrideUnverifiable** | Pointer to **bool** | Record this mode even if verification refuses, together with what was observed. Only the unencrypted mode can be asserted this way: an encrypted state always requires positive per-node evidence. | [optional] [default to false]

## Methods

### NewClusterEncryptionReconcileRequest

`func NewClusterEncryptionReconcileRequest(mode EncryptionModeEnum, ) *ClusterEncryptionReconcileRequest`

NewClusterEncryptionReconcileRequest instantiates a new ClusterEncryptionReconcileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterEncryptionReconcileRequestWithDefaults

`func NewClusterEncryptionReconcileRequestWithDefaults() *ClusterEncryptionReconcileRequest`

NewClusterEncryptionReconcileRequestWithDefaults instantiates a new ClusterEncryptionReconcileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *ClusterEncryptionReconcileRequest) GetMode() EncryptionModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ClusterEncryptionReconcileRequest) GetModeOk() (*EncryptionModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ClusterEncryptionReconcileRequest) SetMode(v EncryptionModeEnum)`

SetMode sets Mode field to given value.


### GetAcknowledgeWorkloadRestart

`func (o *ClusterEncryptionReconcileRequest) GetAcknowledgeWorkloadRestart() bool`

GetAcknowledgeWorkloadRestart returns the AcknowledgeWorkloadRestart field if non-nil, zero value otherwise.

### GetAcknowledgeWorkloadRestartOk

`func (o *ClusterEncryptionReconcileRequest) GetAcknowledgeWorkloadRestartOk() (*bool, bool)`

GetAcknowledgeWorkloadRestartOk returns a tuple with the AcknowledgeWorkloadRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgeWorkloadRestart

`func (o *ClusterEncryptionReconcileRequest) SetAcknowledgeWorkloadRestart(v bool)`

SetAcknowledgeWorkloadRestart sets AcknowledgeWorkloadRestart field to given value.

### HasAcknowledgeWorkloadRestart

`func (o *ClusterEncryptionReconcileRequest) HasAcknowledgeWorkloadRestart() bool`

HasAcknowledgeWorkloadRestart returns a boolean if a field has been set.

### GetOverrideUnverifiable

`func (o *ClusterEncryptionReconcileRequest) GetOverrideUnverifiable() bool`

GetOverrideUnverifiable returns the OverrideUnverifiable field if non-nil, zero value otherwise.

### GetOverrideUnverifiableOk

`func (o *ClusterEncryptionReconcileRequest) GetOverrideUnverifiableOk() (*bool, bool)`

GetOverrideUnverifiableOk returns a tuple with the OverrideUnverifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideUnverifiable

`func (o *ClusterEncryptionReconcileRequest) SetOverrideUnverifiable(v bool)`

SetOverrideUnverifiable sets OverrideUnverifiable field to given value.

### HasOverrideUnverifiable

`func (o *ClusterEncryptionReconcileRequest) HasOverrideUnverifiable() bool`

HasOverrideUnverifiable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


