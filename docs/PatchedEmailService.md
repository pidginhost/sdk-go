# PatchedEmailService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Tier** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to [**StatusA57Enum**](StatusA57Enum.md) |  | [optional] [readonly] 
**SandboxMode** | Pointer to **bool** |  | [optional] [readonly] 
**AutoSuspended** | Pointer to **bool** |  | [optional] [readonly] 
**AutoSuspendReason** | Pointer to **string** |  | [optional] [readonly] 
**MsgsSent24h** | Pointer to **int32** |  | [optional] [readonly] 
**MsgsSent30d** | Pointer to **int32** |  | [optional] [readonly] 
**BounceRatePct** | Pointer to **float64** |  | [optional] [readonly] 
**ComplaintRatePct** | Pointer to **float64** |  | [optional] [readonly] 
**DedicatedIpAddon** | Pointer to **bool** |  | [optional] [readonly] 
**QuotaMonthly** | Pointer to **string** |  | [optional] [readonly] 
**PriceMonthlyEur** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPatchedEmailService

`func NewPatchedEmailService() *PatchedEmailService`

NewPatchedEmailService instantiates a new PatchedEmailService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEmailServiceWithDefaults

`func NewPatchedEmailServiceWithDefaults() *PatchedEmailService`

NewPatchedEmailServiceWithDefaults instantiates a new PatchedEmailService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedEmailService) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedEmailService) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedEmailService) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedEmailService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTier

`func (o *PatchedEmailService) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *PatchedEmailService) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *PatchedEmailService) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *PatchedEmailService) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedEmailService) GetStatus() StatusA57Enum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedEmailService) GetStatusOk() (*StatusA57Enum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedEmailService) SetStatus(v StatusA57Enum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedEmailService) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSandboxMode

`func (o *PatchedEmailService) GetSandboxMode() bool`

GetSandboxMode returns the SandboxMode field if non-nil, zero value otherwise.

### GetSandboxModeOk

`func (o *PatchedEmailService) GetSandboxModeOk() (*bool, bool)`

GetSandboxModeOk returns a tuple with the SandboxMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxMode

`func (o *PatchedEmailService) SetSandboxMode(v bool)`

SetSandboxMode sets SandboxMode field to given value.

### HasSandboxMode

`func (o *PatchedEmailService) HasSandboxMode() bool`

HasSandboxMode returns a boolean if a field has been set.

### GetAutoSuspended

`func (o *PatchedEmailService) GetAutoSuspended() bool`

GetAutoSuspended returns the AutoSuspended field if non-nil, zero value otherwise.

### GetAutoSuspendedOk

`func (o *PatchedEmailService) GetAutoSuspendedOk() (*bool, bool)`

GetAutoSuspendedOk returns a tuple with the AutoSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSuspended

`func (o *PatchedEmailService) SetAutoSuspended(v bool)`

SetAutoSuspended sets AutoSuspended field to given value.

### HasAutoSuspended

`func (o *PatchedEmailService) HasAutoSuspended() bool`

HasAutoSuspended returns a boolean if a field has been set.

### GetAutoSuspendReason

`func (o *PatchedEmailService) GetAutoSuspendReason() string`

GetAutoSuspendReason returns the AutoSuspendReason field if non-nil, zero value otherwise.

### GetAutoSuspendReasonOk

`func (o *PatchedEmailService) GetAutoSuspendReasonOk() (*string, bool)`

GetAutoSuspendReasonOk returns a tuple with the AutoSuspendReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSuspendReason

`func (o *PatchedEmailService) SetAutoSuspendReason(v string)`

SetAutoSuspendReason sets AutoSuspendReason field to given value.

### HasAutoSuspendReason

`func (o *PatchedEmailService) HasAutoSuspendReason() bool`

HasAutoSuspendReason returns a boolean if a field has been set.

### GetMsgsSent24h

`func (o *PatchedEmailService) GetMsgsSent24h() int32`

GetMsgsSent24h returns the MsgsSent24h field if non-nil, zero value otherwise.

### GetMsgsSent24hOk

`func (o *PatchedEmailService) GetMsgsSent24hOk() (*int32, bool)`

GetMsgsSent24hOk returns a tuple with the MsgsSent24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgsSent24h

`func (o *PatchedEmailService) SetMsgsSent24h(v int32)`

SetMsgsSent24h sets MsgsSent24h field to given value.

### HasMsgsSent24h

`func (o *PatchedEmailService) HasMsgsSent24h() bool`

HasMsgsSent24h returns a boolean if a field has been set.

### GetMsgsSent30d

`func (o *PatchedEmailService) GetMsgsSent30d() int32`

GetMsgsSent30d returns the MsgsSent30d field if non-nil, zero value otherwise.

### GetMsgsSent30dOk

`func (o *PatchedEmailService) GetMsgsSent30dOk() (*int32, bool)`

GetMsgsSent30dOk returns a tuple with the MsgsSent30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgsSent30d

`func (o *PatchedEmailService) SetMsgsSent30d(v int32)`

SetMsgsSent30d sets MsgsSent30d field to given value.

### HasMsgsSent30d

`func (o *PatchedEmailService) HasMsgsSent30d() bool`

HasMsgsSent30d returns a boolean if a field has been set.

### GetBounceRatePct

`func (o *PatchedEmailService) GetBounceRatePct() float64`

GetBounceRatePct returns the BounceRatePct field if non-nil, zero value otherwise.

### GetBounceRatePctOk

`func (o *PatchedEmailService) GetBounceRatePctOk() (*float64, bool)`

GetBounceRatePctOk returns a tuple with the BounceRatePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounceRatePct

`func (o *PatchedEmailService) SetBounceRatePct(v float64)`

SetBounceRatePct sets BounceRatePct field to given value.

### HasBounceRatePct

`func (o *PatchedEmailService) HasBounceRatePct() bool`

HasBounceRatePct returns a boolean if a field has been set.

### GetComplaintRatePct

`func (o *PatchedEmailService) GetComplaintRatePct() float64`

GetComplaintRatePct returns the ComplaintRatePct field if non-nil, zero value otherwise.

### GetComplaintRatePctOk

`func (o *PatchedEmailService) GetComplaintRatePctOk() (*float64, bool)`

GetComplaintRatePctOk returns a tuple with the ComplaintRatePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplaintRatePct

`func (o *PatchedEmailService) SetComplaintRatePct(v float64)`

SetComplaintRatePct sets ComplaintRatePct field to given value.

### HasComplaintRatePct

`func (o *PatchedEmailService) HasComplaintRatePct() bool`

HasComplaintRatePct returns a boolean if a field has been set.

### GetDedicatedIpAddon

`func (o *PatchedEmailService) GetDedicatedIpAddon() bool`

GetDedicatedIpAddon returns the DedicatedIpAddon field if non-nil, zero value otherwise.

### GetDedicatedIpAddonOk

`func (o *PatchedEmailService) GetDedicatedIpAddonOk() (*bool, bool)`

GetDedicatedIpAddonOk returns a tuple with the DedicatedIpAddon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedIpAddon

`func (o *PatchedEmailService) SetDedicatedIpAddon(v bool)`

SetDedicatedIpAddon sets DedicatedIpAddon field to given value.

### HasDedicatedIpAddon

`func (o *PatchedEmailService) HasDedicatedIpAddon() bool`

HasDedicatedIpAddon returns a boolean if a field has been set.

### GetQuotaMonthly

`func (o *PatchedEmailService) GetQuotaMonthly() string`

GetQuotaMonthly returns the QuotaMonthly field if non-nil, zero value otherwise.

### GetQuotaMonthlyOk

`func (o *PatchedEmailService) GetQuotaMonthlyOk() (*string, bool)`

GetQuotaMonthlyOk returns a tuple with the QuotaMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaMonthly

`func (o *PatchedEmailService) SetQuotaMonthly(v string)`

SetQuotaMonthly sets QuotaMonthly field to given value.

### HasQuotaMonthly

`func (o *PatchedEmailService) HasQuotaMonthly() bool`

HasQuotaMonthly returns a boolean if a field has been set.

### GetPriceMonthlyEur

`func (o *PatchedEmailService) GetPriceMonthlyEur() string`

GetPriceMonthlyEur returns the PriceMonthlyEur field if non-nil, zero value otherwise.

### GetPriceMonthlyEurOk

`func (o *PatchedEmailService) GetPriceMonthlyEurOk() (*string, bool)`

GetPriceMonthlyEurOk returns a tuple with the PriceMonthlyEur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMonthlyEur

`func (o *PatchedEmailService) SetPriceMonthlyEur(v string)`

SetPriceMonthlyEur sets PriceMonthlyEur field to given value.

### HasPriceMonthlyEur

`func (o *PatchedEmailService) HasPriceMonthlyEur() bool`

HasPriceMonthlyEur returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


