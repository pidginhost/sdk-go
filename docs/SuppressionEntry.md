# SuppressionEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Address** | **string** |  | [readonly] 
**Reason** | [**ReasonEnum**](ReasonEnum.md) |  | [readonly] 
**Detail** | **string** |  | [readonly] 
**CreatedAt** | **string** |  | [readonly] 

## Methods

### NewSuppressionEntry

`func NewSuppressionEntry(id int32, address string, reason ReasonEnum, detail string, createdAt string, ) *SuppressionEntry`

NewSuppressionEntry instantiates a new SuppressionEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuppressionEntryWithDefaults

`func NewSuppressionEntryWithDefaults() *SuppressionEntry`

NewSuppressionEntryWithDefaults instantiates a new SuppressionEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SuppressionEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SuppressionEntry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SuppressionEntry) SetId(v int32)`

SetId sets Id field to given value.


### GetAddress

`func (o *SuppressionEntry) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SuppressionEntry) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SuppressionEntry) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetReason

`func (o *SuppressionEntry) GetReason() ReasonEnum`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SuppressionEntry) GetReasonOk() (*ReasonEnum, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SuppressionEntry) SetReason(v ReasonEnum)`

SetReason sets Reason field to given value.


### GetDetail

`func (o *SuppressionEntry) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SuppressionEntry) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SuppressionEntry) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetCreatedAt

`func (o *SuppressionEntry) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SuppressionEntry) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SuppressionEntry) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


