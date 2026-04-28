# \CloudAPI

All URIs are relative to *https://www.pidginhost.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudFirewallRulesSetCreate**](CloudAPI.md#CloudFirewallRulesSetCreate) | **Post** /api/cloud/firewall-rules-set/ | 
[**CloudFirewallRulesSetDestroy**](CloudAPI.md#CloudFirewallRulesSetDestroy) | **Delete** /api/cloud/firewall-rules-set/{id}/ | 
[**CloudFirewallRulesSetList**](CloudAPI.md#CloudFirewallRulesSetList) | **Get** /api/cloud/firewall-rules-set/ | 
[**CloudFirewallRulesSetPartialUpdate**](CloudAPI.md#CloudFirewallRulesSetPartialUpdate) | **Patch** /api/cloud/firewall-rules-set/{id}/ | 
[**CloudFirewallRulesSetRetrieve**](CloudAPI.md#CloudFirewallRulesSetRetrieve) | **Get** /api/cloud/firewall-rules-set/{id}/ | 
[**CloudFirewallRulesSetRulesCreate**](CloudAPI.md#CloudFirewallRulesSetRulesCreate) | **Post** /api/cloud/firewall-rules-set/{rules_set_id}/rules/ | 
[**CloudFirewallRulesSetRulesDestroy**](CloudAPI.md#CloudFirewallRulesSetRulesDestroy) | **Delete** /api/cloud/firewall-rules-set/{rules_set_id}/rules/{rule_id}/ | 
[**CloudFirewallRulesSetRulesList**](CloudAPI.md#CloudFirewallRulesSetRulesList) | **Get** /api/cloud/firewall-rules-set/{rules_set_id}/rules/ | 
[**CloudFirewallRulesSetRulesPartialUpdate**](CloudAPI.md#CloudFirewallRulesSetRulesPartialUpdate) | **Patch** /api/cloud/firewall-rules-set/{rules_set_id}/rules/{rule_id}/ | 
[**CloudFirewallRulesSetRulesRetrieve**](CloudAPI.md#CloudFirewallRulesSetRulesRetrieve) | **Get** /api/cloud/firewall-rules-set/{rules_set_id}/rules/{rule_id}/ | 
[**CloudFirewallRulesSetRulesUpdate**](CloudAPI.md#CloudFirewallRulesSetRulesUpdate) | **Put** /api/cloud/firewall-rules-set/{rules_set_id}/rules/{rule_id}/ | 
[**CloudFirewallRulesSetUpdate**](CloudAPI.md#CloudFirewallRulesSetUpdate) | **Put** /api/cloud/firewall-rules-set/{id}/ | 
[**CloudGenerationsList**](CloudAPI.md#CloudGenerationsList) | **Get** /api/cloud/generations/ | List hardware generations
[**CloudGenerationsRetrieve**](CloudAPI.md#CloudGenerationsRetrieve) | **Get** /api/cloud/generations/{slug}/ | 
[**CloudImagesList**](CloudAPI.md#CloudImagesList) | **Get** /api/cloud/images/ | 
[**CloudImagesRetrieve**](CloudAPI.md#CloudImagesRetrieve) | **Get** /api/cloud/images/{id}/ | 
[**CloudIpv4Create**](CloudAPI.md#CloudIpv4Create) | **Post** /api/cloud/ipv4/ | 
[**CloudIpv4Destroy**](CloudAPI.md#CloudIpv4Destroy) | **Delete** /api/cloud/ipv4/{id}/ | 
[**CloudIpv4DetachCreate**](CloudAPI.md#CloudIpv4DetachCreate) | **Post** /api/cloud/ipv4/{id}/detach/ | 
[**CloudIpv4List**](CloudAPI.md#CloudIpv4List) | **Get** /api/cloud/ipv4/ | 
[**CloudIpv4RdnsCreate**](CloudAPI.md#CloudIpv4RdnsCreate) | **Post** /api/cloud/ipv4/{id}/rdns/ | 
[**CloudIpv4RdnsRetrieve**](CloudAPI.md#CloudIpv4RdnsRetrieve) | **Get** /api/cloud/ipv4/{id}/rdns/ | 
[**CloudIpv4Retrieve**](CloudAPI.md#CloudIpv4Retrieve) | **Get** /api/cloud/ipv4/{id}/ | 
[**CloudIpv6Create**](CloudAPI.md#CloudIpv6Create) | **Post** /api/cloud/ipv6/ | 
[**CloudIpv6Destroy**](CloudAPI.md#CloudIpv6Destroy) | **Delete** /api/cloud/ipv6/{id}/ | 
[**CloudIpv6DetachCreate**](CloudAPI.md#CloudIpv6DetachCreate) | **Post** /api/cloud/ipv6/{id}/detach/ | 
[**CloudIpv6List**](CloudAPI.md#CloudIpv6List) | **Get** /api/cloud/ipv6/ | 
[**CloudIpv6RdnsCreate**](CloudAPI.md#CloudIpv6RdnsCreate) | **Post** /api/cloud/ipv6/{id}/rdns/ | 
[**CloudIpv6RdnsRetrieve**](CloudAPI.md#CloudIpv6RdnsRetrieve) | **Get** /api/cloud/ipv6/{id}/rdns/ | 
[**CloudIpv6Retrieve**](CloudAPI.md#CloudIpv6Retrieve) | **Get** /api/cloud/ipv6/{id}/ | 
[**CloudPrivateNetworksAddServerCreate**](CloudAPI.md#CloudPrivateNetworksAddServerCreate) | **Post** /api/cloud/private-networks/{id}/add-server/ | 
[**CloudPrivateNetworksCreate**](CloudAPI.md#CloudPrivateNetworksCreate) | **Post** /api/cloud/private-networks/ | 
[**CloudPrivateNetworksDestroy**](CloudAPI.md#CloudPrivateNetworksDestroy) | **Delete** /api/cloud/private-networks/{id}/ | 
[**CloudPrivateNetworksList**](CloudAPI.md#CloudPrivateNetworksList) | **Get** /api/cloud/private-networks/ | 
[**CloudPrivateNetworksPartialUpdate**](CloudAPI.md#CloudPrivateNetworksPartialUpdate) | **Patch** /api/cloud/private-networks/{id}/ | 
[**CloudPrivateNetworksRemoveServerCreate**](CloudAPI.md#CloudPrivateNetworksRemoveServerCreate) | **Post** /api/cloud/private-networks/{id}/remove-server/ | 
[**CloudPrivateNetworksRetrieve**](CloudAPI.md#CloudPrivateNetworksRetrieve) | **Get** /api/cloud/private-networks/{id}/ | 
[**CloudPrivateNetworksUpdate**](CloudAPI.md#CloudPrivateNetworksUpdate) | **Put** /api/cloud/private-networks/{id}/ | 
[**CloudServerPackagesByGenerationRetrieve**](CloudAPI.md#CloudServerPackagesByGenerationRetrieve) | **Get** /api/cloud/server-packages/by-generation/ | 
[**CloudServerPackagesList**](CloudAPI.md#CloudServerPackagesList) | **Get** /api/cloud/server-packages/ | 
[**CloudServerPackagesRetrieve**](CloudAPI.md#CloudServerPackagesRetrieve) | **Get** /api/cloud/server-packages/{id}/ | 
[**CloudServersActivityRetrieve**](CloudAPI.md#CloudServersActivityRetrieve) | **Get** /api/cloud/servers/{id}/activity/ | 
[**CloudServersAttachIpv4Create**](CloudAPI.md#CloudServersAttachIpv4Create) | **Post** /api/cloud/servers/{id}/attach-ipv4/ | 
[**CloudServersAttachIpv6Create**](CloudAPI.md#CloudServersAttachIpv6Create) | **Post** /api/cloud/servers/{id}/attach-ipv6/ | 
[**CloudServersConsoleCreate**](CloudAPI.md#CloudServersConsoleCreate) | **Post** /api/cloud/servers/{id}/console/ | 
[**CloudServersCreate**](CloudAPI.md#CloudServersCreate) | **Post** /api/cloud/servers/ | 
[**CloudServersDestroy**](CloudAPI.md#CloudServersDestroy) | **Delete** /api/cloud/servers/{id}/ | 
[**CloudServersDestroyProtectionCreate**](CloudAPI.md#CloudServersDestroyProtectionCreate) | **Post** /api/cloud/servers/{id}/destroy-protection/ | 
[**CloudServersDetachIpv4Create**](CloudAPI.md#CloudServersDetachIpv4Create) | **Post** /api/cloud/servers/{id}/detach-ipv4/ | 
[**CloudServersDetachIpv6Create**](CloudAPI.md#CloudServersDetachIpv6Create) | **Post** /api/cloud/servers/{id}/detach-ipv6/ | 
[**CloudServersList**](CloudAPI.md#CloudServersList) | **Get** /api/cloud/servers/ | 
[**CloudServersModifyPackageCreate**](CloudAPI.md#CloudServersModifyPackageCreate) | **Post** /api/cloud/servers/{id}/modify-package/ | 
[**CloudServersPartialUpdate**](CloudAPI.md#CloudServersPartialUpdate) | **Patch** /api/cloud/servers/{id}/ | 
[**CloudServersPowerManagementCreate**](CloudAPI.md#CloudServersPowerManagementCreate) | **Post** /api/cloud/servers/{id}/power-management/ | 
[**CloudServersPowerManagementRetrieve**](CloudAPI.md#CloudServersPowerManagementRetrieve) | **Get** /api/cloud/servers/{id}/power-management/ | 
[**CloudServersPublicInterfaceCreate**](CloudAPI.md#CloudServersPublicInterfaceCreate) | **Post** /api/cloud/servers/{id}/public-interface/ | 
[**CloudServersPublicInterfaceDestroy**](CloudAPI.md#CloudServersPublicInterfaceDestroy) | **Delete** /api/cloud/servers/{id}/public-interface/ | 
[**CloudServersPublicInterfaceRetrieve**](CloudAPI.md#CloudServersPublicInterfaceRetrieve) | **Get** /api/cloud/servers/{id}/public-interface/ | 
[**CloudServersRetrieve**](CloudAPI.md#CloudServersRetrieve) | **Get** /api/cloud/servers/{id}/ | 
[**CloudServersRetryProvisionCreate**](CloudAPI.md#CloudServersRetryProvisionCreate) | **Post** /api/cloud/servers/{id}/retry-provision/ | 
[**CloudServersSnapshotsCreate**](CloudAPI.md#CloudServersSnapshotsCreate) | **Post** /api/cloud/servers/{id}/snapshots/ | 
[**CloudServersSnapshotsDestroy**](CloudAPI.md#CloudServersSnapshotsDestroy) | **Delete** /api/cloud/servers/{id}/snapshots/{snapshot_name}/ | 
[**CloudServersSnapshotsList**](CloudAPI.md#CloudServersSnapshotsList) | **Get** /api/cloud/servers/{id}/snapshots/ | 
[**CloudServersSnapshotsRollbackCreate**](CloudAPI.md#CloudServersSnapshotsRollbackCreate) | **Post** /api/cloud/servers/{id}/snapshots/{snapshot_name}/rollback/ | 
[**CloudServersUpdate**](CloudAPI.md#CloudServersUpdate) | **Put** /api/cloud/servers/{id}/ | 
[**CloudServersUsageRetrieve**](CloudAPI.md#CloudServersUsageRetrieve) | **Get** /api/cloud/servers/{id}/usage/ | 
[**CloudServersVolumesCreate**](CloudAPI.md#CloudServersVolumesCreate) | **Post** /api/cloud/servers/{server_id}/volumes/ | 
[**CloudServersVolumesDestroy**](CloudAPI.md#CloudServersVolumesDestroy) | **Delete** /api/cloud/servers/{server_id}/volumes/{volume_id}/ | 
[**CloudServersVolumesList**](CloudAPI.md#CloudServersVolumesList) | **Get** /api/cloud/servers/{server_id}/volumes/ | 
[**CloudServersVolumesPartialUpdate**](CloudAPI.md#CloudServersVolumesPartialUpdate) | **Patch** /api/cloud/servers/{server_id}/volumes/{volume_id}/ | 
[**CloudServersVolumesRetrieve**](CloudAPI.md#CloudServersVolumesRetrieve) | **Get** /api/cloud/servers/{server_id}/volumes/{volume_id}/ | 
[**CloudServersVolumesUpdate**](CloudAPI.md#CloudServersVolumesUpdate) | **Put** /api/cloud/servers/{server_id}/volumes/{volume_id}/ | 
[**CloudStorageProductsList**](CloudAPI.md#CloudStorageProductsList) | **Get** /api/cloud/storage-products/ | 
[**CloudStorageProductsRetrieve**](CloudAPI.md#CloudStorageProductsRetrieve) | **Get** /api/cloud/storage-products/{id}/ | 
[**CloudVolumesAttachCreate**](CloudAPI.md#CloudVolumesAttachCreate) | **Post** /api/cloud/volumes/{id}/attach/ | 
[**CloudVolumesDestroy**](CloudAPI.md#CloudVolumesDestroy) | **Delete** /api/cloud/volumes/{id}/ | 
[**CloudVolumesDetachCreate**](CloudAPI.md#CloudVolumesDetachCreate) | **Post** /api/cloud/volumes/{id}/detach/ | 
[**CloudVolumesList**](CloudAPI.md#CloudVolumesList) | **Get** /api/cloud/volumes/ | 
[**CloudVolumesPartialUpdate**](CloudAPI.md#CloudVolumesPartialUpdate) | **Patch** /api/cloud/volumes/{id}/ | 
[**CloudVolumesRetrieve**](CloudAPI.md#CloudVolumesRetrieve) | **Get** /api/cloud/volumes/{id}/ | 
[**CloudVolumesUpdate**](CloudAPI.md#CloudVolumesUpdate) | **Put** /api/cloud/volumes/{id}/ | 



## CloudFirewallRulesSetCreate

> FirewallRulesSet CloudFirewallRulesSetCreate(ctx).FirewallRulesSet(firewallRulesSet).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	firewallRulesSet := *openapiclient.NewFirewallRulesSet(int32(123), "Name_example", openapiclient.FirewallRulesSetStatusEnum("validated"), []openapiclient.FirewallRule{*openapiclient.NewFirewallRule(int32(123), openapiclient.FirewallRuleDirectionEnum("in"), openapiclient.FwPolicyOutEnum("ACCEPT"), false, "ErrorMessage_example")}, false) // FirewallRulesSet | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFirewallRulesSetCreate(context.Background()).FirewallRulesSet(firewallRulesSet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFirewallRulesSetCreate`: FirewallRulesSet
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFirewallRulesSetCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firewallRulesSet** | [**FirewallRulesSet**](FirewallRulesSet.md) |  | 

### Return type

[**FirewallRulesSet**](FirewallRulesSet.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetDestroy

> CloudFirewallRulesSetDestroy(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this firewall rules set.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudAPI.CloudFirewallRulesSetDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this firewall rules set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetList

> []FirewallRulesSet CloudFirewallRulesSetList(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFirewallRulesSetList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFirewallRulesSetList`: []FirewallRulesSet
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFirewallRulesSetList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetListRequest struct via the builder pattern


### Return type

[**[]FirewallRulesSet**](FirewallRulesSet.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetPartialUpdate

> FirewallRulesSet CloudFirewallRulesSetPartialUpdate(ctx, id).PatchedFirewallRulesSet(patchedFirewallRulesSet).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this firewall rules set.
	patchedFirewallRulesSet := *openapiclient.NewPatchedFirewallRulesSet() // PatchedFirewallRulesSet |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFirewallRulesSetPartialUpdate(context.Background(), id).PatchedFirewallRulesSet(patchedFirewallRulesSet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFirewallRulesSetPartialUpdate`: FirewallRulesSet
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFirewallRulesSetPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this firewall rules set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedFirewallRulesSet** | [**PatchedFirewallRulesSet**](PatchedFirewallRulesSet.md) |  | 

### Return type

[**FirewallRulesSet**](FirewallRulesSet.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetRetrieve

> FirewallRulesSet CloudFirewallRulesSetRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this firewall rules set.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFirewallRulesSetRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFirewallRulesSetRetrieve`: FirewallRulesSet
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFirewallRulesSetRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this firewall rules set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FirewallRulesSet**](FirewallRulesSet.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetRulesCreate

> FirewallRule CloudFirewallRulesSetRulesCreate(ctx, rulesSetId).FirewallRule(firewallRule).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	rulesSetId := "rulesSetId_example" // string | 
	firewallRule := *openapiclient.NewFirewallRule(int32(123), openapiclient.FirewallRuleDirectionEnum("in"), openapiclient.FwPolicyOutEnum("ACCEPT"), false, "ErrorMessage_example") // FirewallRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFirewallRulesSetRulesCreate(context.Background(), rulesSetId).FirewallRule(firewallRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetRulesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFirewallRulesSetRulesCreate`: FirewallRule
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFirewallRulesSetRulesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rulesSetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetRulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **firewallRule** | [**FirewallRule**](FirewallRule.md) |  | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetRulesDestroy

> CloudFirewallRulesSetRulesDestroy(ctx, ruleId, rulesSetId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	ruleId := "ruleId_example" // string | 
	rulesSetId := "rulesSetId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudAPI.CloudFirewallRulesSetRulesDestroy(context.Background(), ruleId, rulesSetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetRulesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 
**rulesSetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetRulesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetRulesList

> []FirewallRule CloudFirewallRulesSetRulesList(ctx, rulesSetId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	rulesSetId := "rulesSetId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFirewallRulesSetRulesList(context.Background(), rulesSetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetRulesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFirewallRulesSetRulesList`: []FirewallRule
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFirewallRulesSetRulesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rulesSetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetRulesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FirewallRule**](FirewallRule.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetRulesPartialUpdate

> FirewallRule CloudFirewallRulesSetRulesPartialUpdate(ctx, ruleId, rulesSetId).PatchedFirewallRule(patchedFirewallRule).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	ruleId := "ruleId_example" // string | 
	rulesSetId := "rulesSetId_example" // string | 
	patchedFirewallRule := *openapiclient.NewPatchedFirewallRule() // PatchedFirewallRule |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFirewallRulesSetRulesPartialUpdate(context.Background(), ruleId, rulesSetId).PatchedFirewallRule(patchedFirewallRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetRulesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFirewallRulesSetRulesPartialUpdate`: FirewallRule
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFirewallRulesSetRulesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 
**rulesSetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetRulesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedFirewallRule** | [**PatchedFirewallRule**](PatchedFirewallRule.md) |  | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetRulesRetrieve

> FirewallRule CloudFirewallRulesSetRulesRetrieve(ctx, ruleId, rulesSetId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	ruleId := "ruleId_example" // string | 
	rulesSetId := "rulesSetId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFirewallRulesSetRulesRetrieve(context.Background(), ruleId, rulesSetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetRulesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFirewallRulesSetRulesRetrieve`: FirewallRule
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFirewallRulesSetRulesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 
**rulesSetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetRulesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetRulesUpdate

> FirewallRule CloudFirewallRulesSetRulesUpdate(ctx, ruleId, rulesSetId).FirewallRule(firewallRule).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	ruleId := "ruleId_example" // string | 
	rulesSetId := "rulesSetId_example" // string | 
	firewallRule := *openapiclient.NewFirewallRule(int32(123), openapiclient.FirewallRuleDirectionEnum("in"), openapiclient.FwPolicyOutEnum("ACCEPT"), false, "ErrorMessage_example") // FirewallRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFirewallRulesSetRulesUpdate(context.Background(), ruleId, rulesSetId).FirewallRule(firewallRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetRulesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFirewallRulesSetRulesUpdate`: FirewallRule
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFirewallRulesSetRulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 
**rulesSetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetRulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **firewallRule** | [**FirewallRule**](FirewallRule.md) |  | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFirewallRulesSetUpdate

> FirewallRulesSet CloudFirewallRulesSetUpdate(ctx, id).FirewallRulesSet(firewallRulesSet).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this firewall rules set.
	firewallRulesSet := *openapiclient.NewFirewallRulesSet(int32(123), "Name_example", openapiclient.FirewallRulesSetStatusEnum("validated"), []openapiclient.FirewallRule{*openapiclient.NewFirewallRule(int32(123), openapiclient.FirewallRuleDirectionEnum("in"), openapiclient.FwPolicyOutEnum("ACCEPT"), false, "ErrorMessage_example")}, false) // FirewallRulesSet | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudFirewallRulesSetUpdate(context.Background(), id).FirewallRulesSet(firewallRulesSet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudFirewallRulesSetUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudFirewallRulesSetUpdate`: FirewallRulesSet
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudFirewallRulesSetUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this firewall rules set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudFirewallRulesSetUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **firewallRulesSet** | [**FirewallRulesSet**](FirewallRulesSet.md) |  | 

### Return type

[**FirewallRulesSet**](FirewallRulesSet.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGenerationsList

> []HardwareGeneration CloudGenerationsList(ctx).Execute()

List hardware generations

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGenerationsList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGenerationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGenerationsList`: []HardwareGeneration
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGenerationsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGenerationsListRequest struct via the builder pattern


### Return type

[**[]HardwareGeneration**](HardwareGeneration.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGenerationsRetrieve

> HardwareGeneration CloudGenerationsRetrieve(ctx, slug).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudGenerationsRetrieve(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudGenerationsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGenerationsRetrieve`: HardwareGeneration
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudGenerationsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGenerationsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HardwareGeneration**](HardwareGeneration.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudImagesList

> PaginatedOSImageList CloudImagesList(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudImagesList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudImagesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudImagesList`: PaginatedOSImageList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudImagesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudImagesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedOSImageList**](PaginatedOSImageList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudImagesRetrieve

> OSImage CloudImagesRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this operating system.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudImagesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudImagesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudImagesRetrieve`: OSImage
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudImagesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this operating system. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudImagesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSImage**](OSImage.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv4Create

> PublicIPv4 CloudIpv4Create(ctx).PublicIPv4(publicIPv4).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	publicIPv4 := *openapiclient.NewPublicIPv4(int32(123), "Slug_example", "Address_example", "Gateway_example", int32(123), false, "Server_example") // PublicIPv4 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv4Create(context.Background()).PublicIPv4(publicIPv4).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv4Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv4Create`: PublicIPv4
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv4Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv4CreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicIPv4** | [**PublicIPv4**](PublicIPv4.md) |  | 

### Return type

[**PublicIPv4**](PublicIPv4.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv4Destroy

> CloudIpv4Destroy(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Public IPv4.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudAPI.CloudIpv4Destroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv4Destroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Public IPv4. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv4DestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv4DetachCreate

> DetachIPv4Response CloudIpv4DetachCreate(ctx, id).PublicIPv4(publicIPv4).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Public IPv4.
	publicIPv4 := *openapiclient.NewPublicIPv4(int32(123), "Slug_example", "Address_example", "Gateway_example", int32(123), false, "Server_example") // PublicIPv4 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv4DetachCreate(context.Background(), id).PublicIPv4(publicIPv4).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv4DetachCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv4DetachCreate`: DetachIPv4Response
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv4DetachCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Public IPv4. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv4DetachCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicIPv4** | [**PublicIPv4**](PublicIPv4.md) |  | 

### Return type

[**DetachIPv4Response**](DetachIPv4Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv4List

> PaginatedPublicIPv4List CloudIpv4List(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv4List(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv4List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv4List`: PaginatedPublicIPv4List
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv4List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv4ListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedPublicIPv4List**](PaginatedPublicIPv4List.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv4RdnsCreate

> ReverseDNS CloudIpv4RdnsCreate(ctx, id).ReverseDNS(reverseDNS).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Public IPv4.
	reverseDNS := *openapiclient.NewReverseDNS("ReverseDns_example") // ReverseDNS | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv4RdnsCreate(context.Background(), id).ReverseDNS(reverseDNS).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv4RdnsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv4RdnsCreate`: ReverseDNS
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv4RdnsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Public IPv4. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv4RdnsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reverseDNS** | [**ReverseDNS**](ReverseDNS.md) |  | 

### Return type

[**ReverseDNS**](ReverseDNS.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv4RdnsRetrieve

> ReverseDNS CloudIpv4RdnsRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Public IPv4.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv4RdnsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv4RdnsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv4RdnsRetrieve`: ReverseDNS
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv4RdnsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Public IPv4. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv4RdnsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReverseDNS**](ReverseDNS.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv4Retrieve

> PublicIPv4 CloudIpv4Retrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Public IPv4.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv4Retrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv4Retrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv4Retrieve`: PublicIPv4
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv4Retrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Public IPv4. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv4RetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicIPv4**](PublicIPv4.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv6Create

> PublicIPv6 CloudIpv6Create(ctx).PublicIPv6(publicIPv6).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	publicIPv6 := *openapiclient.NewPublicIPv6(int32(123), "Slug_example", "Address_example", "Gateway_example", int32(123), false, "Server_example") // PublicIPv6 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv6Create(context.Background()).PublicIPv6(publicIPv6).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv6Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv6Create`: PublicIPv6
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv6Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv6CreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicIPv6** | [**PublicIPv6**](PublicIPv6.md) |  | 

### Return type

[**PublicIPv6**](PublicIPv6.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv6Destroy

> CloudIpv6Destroy(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Public IPv6.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudAPI.CloudIpv6Destroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv6Destroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Public IPv6. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv6DestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv6DetachCreate

> DetachIPv6Response CloudIpv6DetachCreate(ctx, id).PublicIPv6(publicIPv6).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Public IPv6.
	publicIPv6 := *openapiclient.NewPublicIPv6(int32(123), "Slug_example", "Address_example", "Gateway_example", int32(123), false, "Server_example") // PublicIPv6 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv6DetachCreate(context.Background(), id).PublicIPv6(publicIPv6).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv6DetachCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv6DetachCreate`: DetachIPv6Response
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv6DetachCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Public IPv6. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv6DetachCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicIPv6** | [**PublicIPv6**](PublicIPv6.md) |  | 

### Return type

[**DetachIPv6Response**](DetachIPv6Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv6List

> PaginatedPublicIPv6List CloudIpv6List(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv6List(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv6List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv6List`: PaginatedPublicIPv6List
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv6List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv6ListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedPublicIPv6List**](PaginatedPublicIPv6List.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv6RdnsCreate

> ReverseDNS CloudIpv6RdnsCreate(ctx, id).ReverseDNS(reverseDNS).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Public IPv6.
	reverseDNS := *openapiclient.NewReverseDNS("ReverseDns_example") // ReverseDNS | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv6RdnsCreate(context.Background(), id).ReverseDNS(reverseDNS).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv6RdnsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv6RdnsCreate`: ReverseDNS
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv6RdnsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Public IPv6. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv6RdnsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reverseDNS** | [**ReverseDNS**](ReverseDNS.md) |  | 

### Return type

[**ReverseDNS**](ReverseDNS.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv6RdnsRetrieve

> ReverseDNS CloudIpv6RdnsRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Public IPv6.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv6RdnsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv6RdnsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv6RdnsRetrieve`: ReverseDNS
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv6RdnsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Public IPv6. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv6RdnsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReverseDNS**](ReverseDNS.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudIpv6Retrieve

> PublicIPv6 CloudIpv6Retrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Public IPv6.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudIpv6Retrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudIpv6Retrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudIpv6Retrieve`: PublicIPv6
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudIpv6Retrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Public IPv6. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudIpv6RetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicIPv6**](PublicIPv6.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPrivateNetworksAddServerCreate

> AddServerResponse CloudPrivateNetworksAddServerCreate(ctx, id).PrivateNetworkAddHost(privateNetworkAddHost).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this private network.
	privateNetworkAddHost := *openapiclient.NewPrivateNetworkAddHost("Server_example") // PrivateNetworkAddHost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudPrivateNetworksAddServerCreate(context.Background(), id).PrivateNetworkAddHost(privateNetworkAddHost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudPrivateNetworksAddServerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPrivateNetworksAddServerCreate`: AddServerResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudPrivateNetworksAddServerCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this private network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPrivateNetworksAddServerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateNetworkAddHost** | [**PrivateNetworkAddHost**](PrivateNetworkAddHost.md) |  | 

### Return type

[**AddServerResponse**](AddServerResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPrivateNetworksCreate

> PrivateNetwork CloudPrivateNetworksCreate(ctx).PrivateNetwork(privateNetwork).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	privateNetwork := *openapiclient.NewPrivateNetwork(int32(123), "Slug_example", "Address_example", false, []map[string]string{map[string]string{"key": "Inner_example"}}) // PrivateNetwork | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudPrivateNetworksCreate(context.Background()).PrivateNetwork(privateNetwork).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudPrivateNetworksCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPrivateNetworksCreate`: PrivateNetwork
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudPrivateNetworksCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPrivateNetworksCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **privateNetwork** | [**PrivateNetwork**](PrivateNetwork.md) |  | 

### Return type

[**PrivateNetwork**](PrivateNetwork.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPrivateNetworksDestroy

> CloudPrivateNetworksDestroy(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this private network.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudAPI.CloudPrivateNetworksDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudPrivateNetworksDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this private network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPrivateNetworksDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPrivateNetworksList

> PaginatedPrivateNetworkList CloudPrivateNetworksList(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudPrivateNetworksList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudPrivateNetworksList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPrivateNetworksList`: PaginatedPrivateNetworkList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudPrivateNetworksList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPrivateNetworksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedPrivateNetworkList**](PaginatedPrivateNetworkList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPrivateNetworksPartialUpdate

> PrivateNetwork CloudPrivateNetworksPartialUpdate(ctx, id).PatchedPrivateNetwork(patchedPrivateNetwork).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this private network.
	patchedPrivateNetwork := *openapiclient.NewPatchedPrivateNetwork() // PatchedPrivateNetwork |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudPrivateNetworksPartialUpdate(context.Background(), id).PatchedPrivateNetwork(patchedPrivateNetwork).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudPrivateNetworksPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPrivateNetworksPartialUpdate`: PrivateNetwork
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudPrivateNetworksPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this private network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPrivateNetworksPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedPrivateNetwork** | [**PatchedPrivateNetwork**](PatchedPrivateNetwork.md) |  | 

### Return type

[**PrivateNetwork**](PrivateNetwork.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPrivateNetworksRemoveServerCreate

> RemoveServerResponse CloudPrivateNetworksRemoveServerCreate(ctx, id).PrivateNetworkRemoveHost(privateNetworkRemoveHost).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this private network.
	privateNetworkRemoveHost := *openapiclient.NewPrivateNetworkRemoveHost("Server_example") // PrivateNetworkRemoveHost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudPrivateNetworksRemoveServerCreate(context.Background(), id).PrivateNetworkRemoveHost(privateNetworkRemoveHost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudPrivateNetworksRemoveServerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPrivateNetworksRemoveServerCreate`: RemoveServerResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudPrivateNetworksRemoveServerCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this private network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPrivateNetworksRemoveServerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateNetworkRemoveHost** | [**PrivateNetworkRemoveHost**](PrivateNetworkRemoveHost.md) |  | 

### Return type

[**RemoveServerResponse**](RemoveServerResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPrivateNetworksRetrieve

> PrivateNetwork CloudPrivateNetworksRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this private network.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudPrivateNetworksRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudPrivateNetworksRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPrivateNetworksRetrieve`: PrivateNetwork
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudPrivateNetworksRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this private network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPrivateNetworksRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrivateNetwork**](PrivateNetwork.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPrivateNetworksUpdate

> PrivateNetwork CloudPrivateNetworksUpdate(ctx, id).PrivateNetwork(privateNetwork).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this private network.
	privateNetwork := *openapiclient.NewPrivateNetwork(int32(123), "Slug_example", "Address_example", false, []map[string]string{map[string]string{"key": "Inner_example"}}) // PrivateNetwork | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudPrivateNetworksUpdate(context.Background(), id).PrivateNetwork(privateNetwork).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudPrivateNetworksUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPrivateNetworksUpdate`: PrivateNetwork
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudPrivateNetworksUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this private network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPrivateNetworksUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateNetwork** | [**PrivateNetwork**](PrivateNetwork.md) |  | 

### Return type

[**PrivateNetwork**](PrivateNetwork.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServerPackagesByGenerationRetrieve

> ServerProduct CloudServerPackagesByGenerationRetrieve(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServerPackagesByGenerationRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServerPackagesByGenerationRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServerPackagesByGenerationRetrieve`: ServerProduct
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServerPackagesByGenerationRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServerPackagesByGenerationRetrieveRequest struct via the builder pattern


### Return type

[**ServerProduct**](ServerProduct.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServerPackagesList

> PaginatedServerProductList CloudServerPackagesList(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServerPackagesList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServerPackagesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServerPackagesList`: PaginatedServerProductList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServerPackagesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudServerPackagesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedServerProductList**](PaginatedServerProductList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServerPackagesRetrieve

> ServerProduct CloudServerPackagesRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this metered product.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServerPackagesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServerPackagesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServerPackagesRetrieve`: ServerProduct
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServerPackagesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this metered product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServerPackagesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerProduct**](ServerProduct.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersActivityRetrieve

> ActivityLogResponse CloudServersActivityRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersActivityRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersActivityRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersActivityRetrieve`: ActivityLogResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersActivityRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersActivityRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActivityLogResponse**](ActivityLogResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersAttachIpv4Create

> AttachIPv4 CloudServersAttachIpv4Create(ctx, id).AttachIPv4(attachIPv4).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	attachIPv4 := *openapiclient.NewAttachIPv4("Ipv4_example") // AttachIPv4 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersAttachIpv4Create(context.Background(), id).AttachIPv4(attachIPv4).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersAttachIpv4Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersAttachIpv4Create`: AttachIPv4
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersAttachIpv4Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersAttachIpv4CreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachIPv4** | [**AttachIPv4**](AttachIPv4.md) |  | 

### Return type

[**AttachIPv4**](AttachIPv4.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersAttachIpv6Create

> AttachIPv6 CloudServersAttachIpv6Create(ctx, id).AttachIPv6(attachIPv6).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	attachIPv6 := *openapiclient.NewAttachIPv6("Ipv6_example") // AttachIPv6 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersAttachIpv6Create(context.Background(), id).AttachIPv6(attachIPv6).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersAttachIpv6Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersAttachIpv6Create`: AttachIPv6
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersAttachIpv6Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersAttachIpv6CreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachIPv6** | [**AttachIPv6**](AttachIPv6.md) |  | 

### Return type

[**AttachIPv6**](AttachIPv6.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersConsoleCreate

> ConsoleToken CloudServersConsoleCreate(ctx, id).Server(server).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	server := *openapiclient.NewServer(int32(123), "Image_example", "Package_example", int32(123), int32(123), int32(123), "Generation_example", false, false, map[string]interface{}{"key": interface{}(123)}) // Server |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersConsoleCreate(context.Background(), id).Server(server).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersConsoleCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersConsoleCreate`: ConsoleToken
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersConsoleCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersConsoleCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **server** | [**Server**](Server.md) |  | 

### Return type

[**ConsoleToken**](ConsoleToken.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersCreate

> ServerAddResponse CloudServersCreate(ctx).ServerAdd(serverAdd).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	serverAdd := *openapiclient.NewServerAdd("Image_example", "Package_example") // ServerAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersCreate(context.Background()).ServerAdd(serverAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersCreate`: ServerAddResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverAdd** | [**ServerAdd**](ServerAdd.md) |  | 

### Return type

[**ServerAddResponse**](ServerAddResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersDestroy

> CloudServersDestroy(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudAPI.CloudServersDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersDestroyProtectionCreate

> DestroyProtection CloudServersDestroyProtectionCreate(ctx, id).DestroyProtection(destroyProtection).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	destroyProtection := *openapiclient.NewDestroyProtection(false) // DestroyProtection | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersDestroyProtectionCreate(context.Background(), id).DestroyProtection(destroyProtection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersDestroyProtectionCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersDestroyProtectionCreate`: DestroyProtection
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersDestroyProtectionCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersDestroyProtectionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **destroyProtection** | [**DestroyProtection**](DestroyProtection.md) |  | 

### Return type

[**DestroyProtection**](DestroyProtection.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersDetachIpv4Create

> DetachIPv4 CloudServersDetachIpv4Create(ctx, id).Server(server).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	server := *openapiclient.NewServer(int32(123), "Image_example", "Package_example", int32(123), int32(123), int32(123), "Generation_example", false, false, map[string]interface{}{"key": interface{}(123)}) // Server |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersDetachIpv4Create(context.Background(), id).Server(server).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersDetachIpv4Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersDetachIpv4Create`: DetachIPv4
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersDetachIpv4Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersDetachIpv4CreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **server** | [**Server**](Server.md) |  | 

### Return type

[**DetachIPv4**](DetachIPv4.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersDetachIpv6Create

> DetachIPv6 CloudServersDetachIpv6Create(ctx, id).Server(server).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	server := *openapiclient.NewServer(int32(123), "Image_example", "Package_example", int32(123), int32(123), int32(123), "Generation_example", false, false, map[string]interface{}{"key": interface{}(123)}) // Server |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersDetachIpv6Create(context.Background(), id).Server(server).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersDetachIpv6Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersDetachIpv6Create`: DetachIPv6
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersDetachIpv6Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersDetachIpv6CreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **server** | [**Server**](Server.md) |  | 

### Return type

[**DetachIPv6**](DetachIPv6.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersList

> PaginatedServerList CloudServersList(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersList`: PaginatedServerList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedServerList**](PaginatedServerList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersModifyPackageCreate

> ServerUpgradeResponse CloudServersModifyPackageCreate(ctx, id).ServerProductUpgrade(serverProductUpgrade).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	serverProductUpgrade := *openapiclient.NewServerProductUpgrade("Package_example") // ServerProductUpgrade | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersModifyPackageCreate(context.Background(), id).ServerProductUpgrade(serverProductUpgrade).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersModifyPackageCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersModifyPackageCreate`: ServerUpgradeResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersModifyPackageCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersModifyPackageCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverProductUpgrade** | [**ServerProductUpgrade**](ServerProductUpgrade.md) |  | 

### Return type

[**ServerUpgradeResponse**](ServerUpgradeResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersPartialUpdate

> ServerDetail CloudServersPartialUpdate(ctx, id).PatchedServerDetail(patchedServerDetail).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	patchedServerDetail := *openapiclient.NewPatchedServerDetail() // PatchedServerDetail |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersPartialUpdate(context.Background(), id).PatchedServerDetail(patchedServerDetail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersPartialUpdate`: ServerDetail
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedServerDetail** | [**PatchedServerDetail**](PatchedServerDetail.md) |  | 

### Return type

[**ServerDetail**](ServerDetail.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersPowerManagementCreate

> PowerManagement CloudServersPowerManagementCreate(ctx, id).PowerManagementRequest(powerManagementRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	powerManagementRequest := *openapiclient.NewPowerManagementRequest(openapiclient.PowerManagementRequestActionEnum("start")) // PowerManagementRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersPowerManagementCreate(context.Background(), id).PowerManagementRequest(powerManagementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersPowerManagementCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersPowerManagementCreate`: PowerManagement
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersPowerManagementCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersPowerManagementCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **powerManagementRequest** | [**PowerManagementRequest**](PowerManagementRequest.md) |  | 

### Return type

[**PowerManagement**](PowerManagement.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersPowerManagementRetrieve

> PowerManagement CloudServersPowerManagementRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersPowerManagementRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersPowerManagementRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersPowerManagementRetrieve`: PowerManagement
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersPowerManagementRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersPowerManagementRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PowerManagement**](PowerManagement.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersPublicInterfaceCreate

> PublicInterface CloudServersPublicInterfaceCreate(ctx, id).PublicInterface(publicInterface).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	publicInterface := *openapiclient.NewPublicInterface("Interface_example", "Ipv4_example", "Ipv6_example") // PublicInterface |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersPublicInterfaceCreate(context.Background(), id).PublicInterface(publicInterface).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersPublicInterfaceCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersPublicInterfaceCreate`: PublicInterface
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersPublicInterfaceCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersPublicInterfaceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicInterface** | [**PublicInterface**](PublicInterface.md) |  | 

### Return type

[**PublicInterface**](PublicInterface.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersPublicInterfaceDestroy

> CloudServersPublicInterfaceDestroy(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudAPI.CloudServersPublicInterfaceDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersPublicInterfaceDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersPublicInterfaceDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersPublicInterfaceRetrieve

> PublicInterface CloudServersPublicInterfaceRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersPublicInterfaceRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersPublicInterfaceRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersPublicInterfaceRetrieve`: PublicInterface
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersPublicInterfaceRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersPublicInterfaceRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicInterface**](PublicInterface.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersRetrieve

> ServerDetail CloudServersRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersRetrieve`: ServerDetail
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerDetail**](ServerDetail.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersRetryProvisionCreate

> RetryProvision CloudServersRetryProvisionCreate(ctx, id).Server(server).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	server := *openapiclient.NewServer(int32(123), "Image_example", "Package_example", int32(123), int32(123), int32(123), "Generation_example", false, false, map[string]interface{}{"key": interface{}(123)}) // Server |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersRetryProvisionCreate(context.Background(), id).Server(server).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersRetryProvisionCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersRetryProvisionCreate`: RetryProvision
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersRetryProvisionCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersRetryProvisionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **server** | [**Server**](Server.md) |  | 

### Return type

[**RetryProvision**](RetryProvision.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersSnapshotsCreate

> PaginatedSnapshotList CloudServersSnapshotsCreate(ctx, id).SnapshotCreate(snapshotCreate).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	snapshotCreate := *openapiclient.NewSnapshotCreate("Name_example") // SnapshotCreate | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersSnapshotsCreate(context.Background(), id).SnapshotCreate(snapshotCreate).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersSnapshotsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersSnapshotsCreate`: PaginatedSnapshotList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersSnapshotsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersSnapshotsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snapshotCreate** | [**SnapshotCreate**](SnapshotCreate.md) |  | 
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedSnapshotList**](PaginatedSnapshotList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersSnapshotsDestroy

> SnapshotDeleteQueued CloudServersSnapshotsDestroy(ctx, id, snapshotName).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	snapshotName := "snapshotName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersSnapshotsDestroy(context.Background(), id, snapshotName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersSnapshotsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersSnapshotsDestroy`: SnapshotDeleteQueued
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersSnapshotsDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 
**snapshotName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersSnapshotsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SnapshotDeleteQueued**](SnapshotDeleteQueued.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersSnapshotsList

> PaginatedSnapshotList CloudServersSnapshotsList(ctx, id).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersSnapshotsList(context.Background(), id).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersSnapshotsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersSnapshotsList`: PaginatedSnapshotList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersSnapshotsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersSnapshotsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedSnapshotList**](PaginatedSnapshotList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersSnapshotsRollbackCreate

> SnapshotRollbackQueued CloudServersSnapshotsRollbackCreate(ctx, id, snapshotName).Server(server).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	snapshotName := "snapshotName_example" // string | 
	server := *openapiclient.NewServer(int32(123), "Image_example", "Package_example", int32(123), int32(123), int32(123), "Generation_example", false, false, map[string]interface{}{"key": interface{}(123)}) // Server |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersSnapshotsRollbackCreate(context.Background(), id, snapshotName).Server(server).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersSnapshotsRollbackCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersSnapshotsRollbackCreate`: SnapshotRollbackQueued
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersSnapshotsRollbackCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 
**snapshotName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersSnapshotsRollbackCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **server** | [**Server**](Server.md) |  | 

### Return type

[**SnapshotRollbackQueued**](SnapshotRollbackQueued.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersUpdate

> ServerDetail CloudServersUpdate(ctx, id).ServerDetail(serverDetail).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.
	serverDetail := *openapiclient.NewServerDetail(int32(123), "Hostname_example", "Image_example", "Package_example", int32(123), int32(123), int32(123), "Generation_example", map[string]interface{}{"key": interface{}(123)}, []openapiclient.Volume{*openapiclient.NewVolume(int32(123), int32(123), "Product_example", false, "Server_example")}, map[string]interface{}{"key": interface{}(123)}, openapiclient.StatusA57Enum("pending"), "Username_example", false, false) // ServerDetail |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersUpdate(context.Background(), id).ServerDetail(serverDetail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersUpdate`: ServerDetail
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverDetail** | [**ServerDetail**](ServerDetail.md) |  | 

### Return type

[**ServerDetail**](ServerDetail.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersUsageRetrieve

> ServerUsageResponse CloudServersUsageRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this virtual machine.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersUsageRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersUsageRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersUsageRetrieve`: ServerUsageResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersUsageRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this virtual machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersUsageRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerUsageResponse**](ServerUsageResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersVolumesCreate

> Volume CloudServersVolumesCreate(ctx, serverId).Volume(volume).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	serverId := "serverId_example" // string | 
	volume := *openapiclient.NewVolume(int32(123), int32(123), "Product_example", false, "Server_example") // Volume | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersVolumesCreate(context.Background(), serverId).Volume(volume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersVolumesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersVolumesCreate`: Volume
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersVolumesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersVolumesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volume** | [**Volume**](Volume.md) |  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersVolumesDestroy

> CloudServersVolumesDestroy(ctx, serverId, volumeId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	serverId := "serverId_example" // string | 
	volumeId := "volumeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudAPI.CloudServersVolumesDestroy(context.Background(), serverId, volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersVolumesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 
**volumeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersVolumesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersVolumesList

> []Volume CloudServersVolumesList(ctx, serverId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	serverId := "serverId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersVolumesList(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersVolumesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersVolumesList`: []Volume
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersVolumesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersVolumesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Volume**](Volume.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersVolumesPartialUpdate

> Volume CloudServersVolumesPartialUpdate(ctx, serverId, volumeId).PatchedVolume(patchedVolume).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	serverId := "serverId_example" // string | 
	volumeId := "volumeId_example" // string | 
	patchedVolume := *openapiclient.NewPatchedVolume() // PatchedVolume |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersVolumesPartialUpdate(context.Background(), serverId, volumeId).PatchedVolume(patchedVolume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersVolumesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersVolumesPartialUpdate`: Volume
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersVolumesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 
**volumeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersVolumesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedVolume** | [**PatchedVolume**](PatchedVolume.md) |  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersVolumesRetrieve

> Volume CloudServersVolumesRetrieve(ctx, serverId, volumeId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	serverId := "serverId_example" // string | 
	volumeId := "volumeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersVolumesRetrieve(context.Background(), serverId, volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersVolumesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersVolumesRetrieve`: Volume
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersVolumesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 
**volumeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersVolumesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Volume**](Volume.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServersVolumesUpdate

> Volume CloudServersVolumesUpdate(ctx, serverId, volumeId).Volume(volume).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	serverId := "serverId_example" // string | 
	volumeId := "volumeId_example" // string | 
	volume := *openapiclient.NewVolume(int32(123), int32(123), "Product_example", false, "Server_example") // Volume | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudServersVolumesUpdate(context.Background(), serverId, volumeId).Volume(volume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudServersVolumesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudServersVolumesUpdate`: Volume
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudServersVolumesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 
**volumeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServersVolumesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **volume** | [**Volume**](Volume.md) |  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudStorageProductsList

> PaginatedStorageProductList CloudStorageProductsList(ctx).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudStorageProductsList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudStorageProductsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudStorageProductsList`: PaginatedStorageProductList
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudStorageProductsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudStorageProductsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedStorageProductList**](PaginatedStorageProductList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudStorageProductsRetrieve

> StorageProduct CloudStorageProductsRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this metered product.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudStorageProductsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudStorageProductsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudStorageProductsRetrieve`: StorageProduct
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudStorageProductsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this metered product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudStorageProductsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StorageProduct**](StorageProduct.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudVolumesAttachCreate

> AttachVolume CloudVolumesAttachCreate(ctx, id).AttachVolume(attachVolume).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this storage.
	attachVolume := *openapiclient.NewAttachVolume(int32(123)) // AttachVolume | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudVolumesAttachCreate(context.Background(), id).AttachVolume(attachVolume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudVolumesAttachCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudVolumesAttachCreate`: AttachVolume
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudVolumesAttachCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudVolumesAttachCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachVolume** | [**AttachVolume**](AttachVolume.md) |  | 

### Return type

[**AttachVolume**](AttachVolume.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudVolumesDestroy

> CloudVolumesDestroy(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this storage.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudAPI.CloudVolumesDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudVolumesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudVolumesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudVolumesDetachCreate

> DetachVolume CloudVolumesDetachCreate(ctx, id).Volume(volume).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this storage.
	volume := *openapiclient.NewVolume(int32(123), int32(123), "Product_example", false, "Server_example") // Volume | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudVolumesDetachCreate(context.Background(), id).Volume(volume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudVolumesDetachCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudVolumesDetachCreate`: DetachVolume
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudVolumesDetachCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudVolumesDetachCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volume** | [**Volume**](Volume.md) |  | 

### Return type

[**DetachVolume**](DetachVolume.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudVolumesList

> []Volume CloudVolumesList(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudVolumesList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudVolumesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudVolumesList`: []Volume
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudVolumesList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudVolumesListRequest struct via the builder pattern


### Return type

[**[]Volume**](Volume.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudVolumesPartialUpdate

> Volume CloudVolumesPartialUpdate(ctx, id).PatchedVolume(patchedVolume).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this storage.
	patchedVolume := *openapiclient.NewPatchedVolume() // PatchedVolume |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudVolumesPartialUpdate(context.Background(), id).PatchedVolume(patchedVolume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudVolumesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudVolumesPartialUpdate`: Volume
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudVolumesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudVolumesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedVolume** | [**PatchedVolume**](PatchedVolume.md) |  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudVolumesRetrieve

> Volume CloudVolumesRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this storage.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudVolumesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudVolumesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudVolumesRetrieve`: Volume
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudVolumesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudVolumesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Volume**](Volume.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudVolumesUpdate

> Volume CloudVolumesUpdate(ctx, id).Volume(volume).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pidginhost/sdk-go"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this storage.
	volume := *openapiclient.NewVolume(int32(123), int32(123), "Product_example", false, "Server_example") // Volume | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.CloudVolumesUpdate(context.Background(), id).Volume(volume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.CloudVolumesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudVolumesUpdate`: Volume
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.CloudVolumesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudVolumesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volume** | [**Volume**](Volume.md) |  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

