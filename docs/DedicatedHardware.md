# DedicatedHardware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Processor** | **string** |  | 
**ProcessorCount** | **int32** |  | 
**MemoryGb** | **int32** |  | 
**Disks** | **string** |  | 
**Raid** | **string** |  | 
**Network** | **string** |  | 
**Traffic** | **string** |  | 
**Gpu** | **string** |  | 
**ExtraHardware** | **[]string** |  | 
**Software** | **string** |  | 

## Methods

### NewDedicatedHardware

`func NewDedicatedHardware(title string, processor string, processorCount int32, memoryGb int32, disks string, raid string, network string, traffic string, gpu string, extraHardware []string, software string, ) *DedicatedHardware`

NewDedicatedHardware instantiates a new DedicatedHardware object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedHardwareWithDefaults

`func NewDedicatedHardwareWithDefaults() *DedicatedHardware`

NewDedicatedHardwareWithDefaults instantiates a new DedicatedHardware object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *DedicatedHardware) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DedicatedHardware) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DedicatedHardware) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetProcessor

`func (o *DedicatedHardware) GetProcessor() string`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *DedicatedHardware) GetProcessorOk() (*string, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *DedicatedHardware) SetProcessor(v string)`

SetProcessor sets Processor field to given value.


### GetProcessorCount

`func (o *DedicatedHardware) GetProcessorCount() int32`

GetProcessorCount returns the ProcessorCount field if non-nil, zero value otherwise.

### GetProcessorCountOk

`func (o *DedicatedHardware) GetProcessorCountOk() (*int32, bool)`

GetProcessorCountOk returns a tuple with the ProcessorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCount

`func (o *DedicatedHardware) SetProcessorCount(v int32)`

SetProcessorCount sets ProcessorCount field to given value.


### GetMemoryGb

`func (o *DedicatedHardware) GetMemoryGb() int32`

GetMemoryGb returns the MemoryGb field if non-nil, zero value otherwise.

### GetMemoryGbOk

`func (o *DedicatedHardware) GetMemoryGbOk() (*int32, bool)`

GetMemoryGbOk returns a tuple with the MemoryGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryGb

`func (o *DedicatedHardware) SetMemoryGb(v int32)`

SetMemoryGb sets MemoryGb field to given value.


### GetDisks

`func (o *DedicatedHardware) GetDisks() string`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *DedicatedHardware) GetDisksOk() (*string, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *DedicatedHardware) SetDisks(v string)`

SetDisks sets Disks field to given value.


### GetRaid

`func (o *DedicatedHardware) GetRaid() string`

GetRaid returns the Raid field if non-nil, zero value otherwise.

### GetRaidOk

`func (o *DedicatedHardware) GetRaidOk() (*string, bool)`

GetRaidOk returns a tuple with the Raid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaid

`func (o *DedicatedHardware) SetRaid(v string)`

SetRaid sets Raid field to given value.


### GetNetwork

`func (o *DedicatedHardware) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *DedicatedHardware) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *DedicatedHardware) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetTraffic

`func (o *DedicatedHardware) GetTraffic() string`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *DedicatedHardware) GetTrafficOk() (*string, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *DedicatedHardware) SetTraffic(v string)`

SetTraffic sets Traffic field to given value.


### GetGpu

`func (o *DedicatedHardware) GetGpu() string`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *DedicatedHardware) GetGpuOk() (*string, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *DedicatedHardware) SetGpu(v string)`

SetGpu sets Gpu field to given value.


### GetExtraHardware

`func (o *DedicatedHardware) GetExtraHardware() []string`

GetExtraHardware returns the ExtraHardware field if non-nil, zero value otherwise.

### GetExtraHardwareOk

`func (o *DedicatedHardware) GetExtraHardwareOk() (*[]string, bool)`

GetExtraHardwareOk returns a tuple with the ExtraHardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHardware

`func (o *DedicatedHardware) SetExtraHardware(v []string)`

SetExtraHardware sets ExtraHardware field to given value.


### GetSoftware

`func (o *DedicatedHardware) GetSoftware() string`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *DedicatedHardware) GetSoftwareOk() (*string, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *DedicatedHardware) SetSoftware(v string)`

SetSoftware sets Software field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


