# NodeMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | **float64** |  | 
**Mem** | **int64** |  | 
**Maxmem** | **int64** |  | 
**Status** | **string** |  | 
**Uptime** | **int32** |  | 
**Netin** | **int64** |  | 
**Netout** | **int64** |  | 
**Disk** | **int64** |  | 
**Maxdisk** | **int64** |  | 

## Methods

### NewNodeMetricsResponse

`func NewNodeMetricsResponse(cpu float64, mem int64, maxmem int64, status string, uptime int32, netin int64, netout int64, disk int64, maxdisk int64, ) *NodeMetricsResponse`

NewNodeMetricsResponse instantiates a new NodeMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeMetricsResponseWithDefaults

`func NewNodeMetricsResponseWithDefaults() *NodeMetricsResponse`

NewNodeMetricsResponseWithDefaults instantiates a new NodeMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *NodeMetricsResponse) GetCpu() float64`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *NodeMetricsResponse) GetCpuOk() (*float64, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *NodeMetricsResponse) SetCpu(v float64)`

SetCpu sets Cpu field to given value.


### GetMem

`func (o *NodeMetricsResponse) GetMem() int64`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *NodeMetricsResponse) GetMemOk() (*int64, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *NodeMetricsResponse) SetMem(v int64)`

SetMem sets Mem field to given value.


### GetMaxmem

`func (o *NodeMetricsResponse) GetMaxmem() int64`

GetMaxmem returns the Maxmem field if non-nil, zero value otherwise.

### GetMaxmemOk

`func (o *NodeMetricsResponse) GetMaxmemOk() (*int64, bool)`

GetMaxmemOk returns a tuple with the Maxmem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxmem

`func (o *NodeMetricsResponse) SetMaxmem(v int64)`

SetMaxmem sets Maxmem field to given value.


### GetStatus

`func (o *NodeMetricsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeMetricsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeMetricsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUptime

`func (o *NodeMetricsResponse) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *NodeMetricsResponse) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *NodeMetricsResponse) SetUptime(v int32)`

SetUptime sets Uptime field to given value.


### GetNetin

`func (o *NodeMetricsResponse) GetNetin() int64`

GetNetin returns the Netin field if non-nil, zero value otherwise.

### GetNetinOk

`func (o *NodeMetricsResponse) GetNetinOk() (*int64, bool)`

GetNetinOk returns a tuple with the Netin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetin

`func (o *NodeMetricsResponse) SetNetin(v int64)`

SetNetin sets Netin field to given value.


### GetNetout

`func (o *NodeMetricsResponse) GetNetout() int64`

GetNetout returns the Netout field if non-nil, zero value otherwise.

### GetNetoutOk

`func (o *NodeMetricsResponse) GetNetoutOk() (*int64, bool)`

GetNetoutOk returns a tuple with the Netout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetout

`func (o *NodeMetricsResponse) SetNetout(v int64)`

SetNetout sets Netout field to given value.


### GetDisk

`func (o *NodeMetricsResponse) GetDisk() int64`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *NodeMetricsResponse) GetDiskOk() (*int64, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *NodeMetricsResponse) SetDisk(v int64)`

SetDisk sets Disk field to given value.


### GetMaxdisk

`func (o *NodeMetricsResponse) GetMaxdisk() int64`

GetMaxdisk returns the Maxdisk field if non-nil, zero value otherwise.

### GetMaxdiskOk

`func (o *NodeMetricsResponse) GetMaxdiskOk() (*int64, bool)`

GetMaxdiskOk returns a tuple with the Maxdisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxdisk

`func (o *NodeMetricsResponse) SetMaxdisk(v int64)`

SetMaxdisk sets Maxdisk field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


