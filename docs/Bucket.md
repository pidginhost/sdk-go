# Bucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | [readonly] 
**FullName** | **string** |  | [readonly] 
**QuotaGb** | **int32** |  | [readonly] 
**UsedBytes** | **int32** |  | [readonly] 
**ObjectCount** | **int32** |  | [readonly] 
**PublicRead** | **bool** |  | [readonly] 
**Status** | [**StatusA57Enum**](StatusA57Enum.md) |  | [readonly] 
**Endpoint** | **string** |  | [readonly] 
**Region** | **string** |  | [readonly] 
**Created** | **string** |  | [readonly] 

## Methods

### NewBucket

`func NewBucket(id int32, name string, fullName string, quotaGb int32, usedBytes int32, objectCount int32, publicRead bool, status StatusA57Enum, endpoint string, region string, created string, ) *Bucket`

NewBucket instantiates a new Bucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketWithDefaults

`func NewBucketWithDefaults() *Bucket`

NewBucketWithDefaults instantiates a new Bucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Bucket) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bucket) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bucket) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Bucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Bucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Bucket) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *Bucket) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *Bucket) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *Bucket) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetQuotaGb

`func (o *Bucket) GetQuotaGb() int32`

GetQuotaGb returns the QuotaGb field if non-nil, zero value otherwise.

### GetQuotaGbOk

`func (o *Bucket) GetQuotaGbOk() (*int32, bool)`

GetQuotaGbOk returns a tuple with the QuotaGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaGb

`func (o *Bucket) SetQuotaGb(v int32)`

SetQuotaGb sets QuotaGb field to given value.


### GetUsedBytes

`func (o *Bucket) GetUsedBytes() int32`

GetUsedBytes returns the UsedBytes field if non-nil, zero value otherwise.

### GetUsedBytesOk

`func (o *Bucket) GetUsedBytesOk() (*int32, bool)`

GetUsedBytesOk returns a tuple with the UsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBytes

`func (o *Bucket) SetUsedBytes(v int32)`

SetUsedBytes sets UsedBytes field to given value.


### GetObjectCount

`func (o *Bucket) GetObjectCount() int32`

GetObjectCount returns the ObjectCount field if non-nil, zero value otherwise.

### GetObjectCountOk

`func (o *Bucket) GetObjectCountOk() (*int32, bool)`

GetObjectCountOk returns a tuple with the ObjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectCount

`func (o *Bucket) SetObjectCount(v int32)`

SetObjectCount sets ObjectCount field to given value.


### GetPublicRead

`func (o *Bucket) GetPublicRead() bool`

GetPublicRead returns the PublicRead field if non-nil, zero value otherwise.

### GetPublicReadOk

`func (o *Bucket) GetPublicReadOk() (*bool, bool)`

GetPublicReadOk returns a tuple with the PublicRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicRead

`func (o *Bucket) SetPublicRead(v bool)`

SetPublicRead sets PublicRead field to given value.


### GetStatus

`func (o *Bucket) GetStatus() StatusA57Enum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Bucket) GetStatusOk() (*StatusA57Enum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Bucket) SetStatus(v StatusA57Enum)`

SetStatus sets Status field to given value.


### GetEndpoint

`func (o *Bucket) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *Bucket) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *Bucket) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetRegion

`func (o *Bucket) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Bucket) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Bucket) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetCreated

`func (o *Bucket) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Bucket) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Bucket) SetCreated(v string)`

SetCreated sets Created field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


