# DashboardResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerCount** | **int32** |  | 
**HostingCount** | **int32** |  | 
**Servers** | [**[]DashboardServer**](DashboardServer.md) |  | 
**Hosting** | [**[]HostingService**](HostingService.md) |  | 
**ExpiringDomains** | [**[]ExpiringDomain**](ExpiringDomain.md) |  | 

## Methods

### NewDashboardResources

`func NewDashboardResources(serverCount int32, hostingCount int32, servers []DashboardServer, hosting []HostingService, expiringDomains []ExpiringDomain, ) *DashboardResources`

NewDashboardResources instantiates a new DashboardResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardResourcesWithDefaults

`func NewDashboardResourcesWithDefaults() *DashboardResources`

NewDashboardResourcesWithDefaults instantiates a new DashboardResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerCount

`func (o *DashboardResources) GetServerCount() int32`

GetServerCount returns the ServerCount field if non-nil, zero value otherwise.

### GetServerCountOk

`func (o *DashboardResources) GetServerCountOk() (*int32, bool)`

GetServerCountOk returns a tuple with the ServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCount

`func (o *DashboardResources) SetServerCount(v int32)`

SetServerCount sets ServerCount field to given value.


### GetHostingCount

`func (o *DashboardResources) GetHostingCount() int32`

GetHostingCount returns the HostingCount field if non-nil, zero value otherwise.

### GetHostingCountOk

`func (o *DashboardResources) GetHostingCountOk() (*int32, bool)`

GetHostingCountOk returns a tuple with the HostingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingCount

`func (o *DashboardResources) SetHostingCount(v int32)`

SetHostingCount sets HostingCount field to given value.


### GetServers

`func (o *DashboardResources) GetServers() []DashboardServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *DashboardResources) GetServersOk() (*[]DashboardServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *DashboardResources) SetServers(v []DashboardServer)`

SetServers sets Servers field to given value.


### GetHosting

`func (o *DashboardResources) GetHosting() []HostingService`

GetHosting returns the Hosting field if non-nil, zero value otherwise.

### GetHostingOk

`func (o *DashboardResources) GetHostingOk() (*[]HostingService, bool)`

GetHostingOk returns a tuple with the Hosting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosting

`func (o *DashboardResources) SetHosting(v []HostingService)`

SetHosting sets Hosting field to given value.


### GetExpiringDomains

`func (o *DashboardResources) GetExpiringDomains() []ExpiringDomain`

GetExpiringDomains returns the ExpiringDomains field if non-nil, zero value otherwise.

### GetExpiringDomainsOk

`func (o *DashboardResources) GetExpiringDomainsOk() (*[]ExpiringDomain, bool)`

GetExpiringDomainsOk returns a tuple with the ExpiringDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiringDomains

`func (o *DashboardResources) SetExpiringDomains(v []ExpiringDomain)`

SetExpiringDomains sets ExpiringDomains field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


