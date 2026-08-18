# BucketCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**QuotaGb** | **int32** |  | 
**PublicRead** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewBucketCreate

`func NewBucketCreate(name string, quotaGb int32, ) *BucketCreate`

NewBucketCreate instantiates a new BucketCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketCreateWithDefaults

`func NewBucketCreateWithDefaults() *BucketCreate`

NewBucketCreateWithDefaults instantiates a new BucketCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BucketCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BucketCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BucketCreate) SetName(v string)`

SetName sets Name field to given value.


### GetQuotaGb

`func (o *BucketCreate) GetQuotaGb() int32`

GetQuotaGb returns the QuotaGb field if non-nil, zero value otherwise.

### GetQuotaGbOk

`func (o *BucketCreate) GetQuotaGbOk() (*int32, bool)`

GetQuotaGbOk returns a tuple with the QuotaGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaGb

`func (o *BucketCreate) SetQuotaGb(v int32)`

SetQuotaGb sets QuotaGb field to given value.


### GetPublicRead

`func (o *BucketCreate) GetPublicRead() bool`

GetPublicRead returns the PublicRead field if non-nil, zero value otherwise.

### GetPublicReadOk

`func (o *BucketCreate) GetPublicReadOk() (*bool, bool)`

GetPublicReadOk returns a tuple with the PublicRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicRead

`func (o *BucketCreate) SetPublicRead(v bool)`

SetPublicRead sets PublicRead field to given value.

### HasPublicRead

`func (o *BucketCreate) HasPublicRead() bool`

HasPublicRead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


