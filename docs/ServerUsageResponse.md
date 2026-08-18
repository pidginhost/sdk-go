# ServerUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Uptime** | **int32** |  | 
**UptimeText** | **string** |  | 
**Cpu** | **map[string]interface{}** |  | 
**Memory** | **map[string]interface{}** |  | 

## Methods

### NewServerUsageResponse

`func NewServerUsageResponse(status string, uptime int32, uptimeText string, cpu map[string]interface{}, memory map[string]interface{}, ) *ServerUsageResponse`

NewServerUsageResponse instantiates a new ServerUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerUsageResponseWithDefaults

`func NewServerUsageResponseWithDefaults() *ServerUsageResponse`

NewServerUsageResponseWithDefaults instantiates a new ServerUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ServerUsageResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerUsageResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerUsageResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUptime

`func (o *ServerUsageResponse) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *ServerUsageResponse) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *ServerUsageResponse) SetUptime(v int32)`

SetUptime sets Uptime field to given value.


### GetUptimeText

`func (o *ServerUsageResponse) GetUptimeText() string`

GetUptimeText returns the UptimeText field if non-nil, zero value otherwise.

### GetUptimeTextOk

`func (o *ServerUsageResponse) GetUptimeTextOk() (*string, bool)`

GetUptimeTextOk returns a tuple with the UptimeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeText

`func (o *ServerUsageResponse) SetUptimeText(v string)`

SetUptimeText sets UptimeText field to given value.


### GetCpu

`func (o *ServerUsageResponse) GetCpu() map[string]interface{}`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ServerUsageResponse) GetCpuOk() (*map[string]interface{}, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ServerUsageResponse) SetCpu(v map[string]interface{})`

SetCpu sets Cpu field to given value.


### GetMemory

`func (o *ServerUsageResponse) GetMemory() map[string]interface{}`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ServerUsageResponse) GetMemoryOk() (*map[string]interface{}, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ServerUsageResponse) SetMemory(v map[string]interface{})`

SetMemory sets Memory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


