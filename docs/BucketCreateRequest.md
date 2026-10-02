# BucketCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**QuotaGb** | **int32** |  | 
**PublicRead** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewBucketCreateRequest

`func NewBucketCreateRequest(name string, quotaGb int32, ) *BucketCreateRequest`

NewBucketCreateRequest instantiates a new BucketCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketCreateRequestWithDefaults

`func NewBucketCreateRequestWithDefaults() *BucketCreateRequest`

NewBucketCreateRequestWithDefaults instantiates a new BucketCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BucketCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BucketCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BucketCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetQuotaGb

`func (o *BucketCreateRequest) GetQuotaGb() int32`

GetQuotaGb returns the QuotaGb field if non-nil, zero value otherwise.

### GetQuotaGbOk

`func (o *BucketCreateRequest) GetQuotaGbOk() (*int32, bool)`

GetQuotaGbOk returns a tuple with the QuotaGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaGb

`func (o *BucketCreateRequest) SetQuotaGb(v int32)`

SetQuotaGb sets QuotaGb field to given value.


### GetPublicRead

`func (o *BucketCreateRequest) GetPublicRead() bool`

GetPublicRead returns the PublicRead field if non-nil, zero value otherwise.

### GetPublicReadOk

`func (o *BucketCreateRequest) GetPublicReadOk() (*bool, bool)`

GetPublicReadOk returns a tuple with the PublicRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicRead

`func (o *BucketCreateRequest) SetPublicRead(v bool)`

SetPublicRead sets PublicRead field to given value.

### HasPublicRead

`func (o *BucketCreateRequest) HasPublicRead() bool`

HasPublicRead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


