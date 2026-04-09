# SnapshotCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Letters, numbers, \&quot;.\&quot;, \&quot;_\&quot; and \&quot;-\&quot; (max 63 characters). | 
**Description** | Pointer to **string** |  | [optional] 
**IncludeMemory** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSnapshotCreate

`func NewSnapshotCreate(name string, ) *SnapshotCreate`

NewSnapshotCreate instantiates a new SnapshotCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotCreateWithDefaults

`func NewSnapshotCreateWithDefaults() *SnapshotCreate`

NewSnapshotCreateWithDefaults instantiates a new SnapshotCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SnapshotCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnapshotCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnapshotCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SnapshotCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SnapshotCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SnapshotCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SnapshotCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIncludeMemory

`func (o *SnapshotCreate) GetIncludeMemory() bool`

GetIncludeMemory returns the IncludeMemory field if non-nil, zero value otherwise.

### GetIncludeMemoryOk

`func (o *SnapshotCreate) GetIncludeMemoryOk() (*bool, bool)`

GetIncludeMemoryOk returns a tuple with the IncludeMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMemory

`func (o *SnapshotCreate) SetIncludeMemory(v bool)`

SetIncludeMemory sets IncludeMemory field to given value.

### HasIncludeMemory

`func (o *SnapshotCreate) HasIncludeMemory() bool`

HasIncludeMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


