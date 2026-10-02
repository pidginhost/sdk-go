# \KubernetesAPI

All URIs are relative to *https://www.pidginhost.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KubernetesClusterTypesList**](KubernetesAPI.md#KubernetesClusterTypesList) | **Get** /api/kubernetes/cluster-types/ | 
[**KubernetesClustersConnectVmCreate**](KubernetesAPI.md#KubernetesClustersConnectVmCreate) | **Post** /api/kubernetes/clusters/{id}/connect-vm/ | 
[**KubernetesClustersConnectedVmsRetrieve**](KubernetesAPI.md#KubernetesClustersConnectedVmsRetrieve) | **Get** /api/kubernetes/clusters/{id}/connected-vms/ | 
[**KubernetesClustersCreate**](KubernetesAPI.md#KubernetesClustersCreate) | **Post** /api/kubernetes/clusters/ | 
[**KubernetesClustersDestroy**](KubernetesAPI.md#KubernetesClustersDestroy) | **Delete** /api/kubernetes/clusters/{id}/ | 
[**KubernetesClustersDisconnectVmCreate**](KubernetesAPI.md#KubernetesClustersDisconnectVmCreate) | **Post** /api/kubernetes/clusters/{id}/disconnect-vm/ | 
[**KubernetesClustersEligibleVmsRetrieve**](KubernetesAPI.md#KubernetesClustersEligibleVmsRetrieve) | **Get** /api/kubernetes/clusters/{id}/eligible-vms/ | 
[**KubernetesClustersEncryptionCreate**](KubernetesAPI.md#KubernetesClustersEncryptionCreate) | **Post** /api/kubernetes/clusters/{id}/encryption/ | 
[**KubernetesClustersEncryptionRecheckCreate**](KubernetesAPI.md#KubernetesClustersEncryptionRecheckCreate) | **Post** /api/kubernetes/clusters/{id}/encryption/recheck/ | 
[**KubernetesClustersEncryptionReconcileCreate**](KubernetesAPI.md#KubernetesClustersEncryptionReconcileCreate) | **Post** /api/kubernetes/clusters/{id}/encryption/reconcile/ | 
[**KubernetesClustersEncryptionRetrieve**](KubernetesAPI.md#KubernetesClustersEncryptionRetrieve) | **Get** /api/kubernetes/clusters/{id}/encryption/ | 
[**KubernetesClustersHttproutesCreate**](KubernetesAPI.md#KubernetesClustersHttproutesCreate) | **Post** /api/kubernetes/clusters/{cluster_id}/httproutes/ | 
[**KubernetesClustersHttproutesDestroy**](KubernetesAPI.md#KubernetesClustersHttproutesDestroy) | **Delete** /api/kubernetes/clusters/{cluster_id}/httproutes/{id}/ | 
[**KubernetesClustersHttproutesList**](KubernetesAPI.md#KubernetesClustersHttproutesList) | **Get** /api/kubernetes/clusters/{cluster_id}/httproutes/ | 
[**KubernetesClustersHttproutesPartialUpdate**](KubernetesAPI.md#KubernetesClustersHttproutesPartialUpdate) | **Patch** /api/kubernetes/clusters/{cluster_id}/httproutes/{id}/ | 
[**KubernetesClustersHttproutesRetrieve**](KubernetesAPI.md#KubernetesClustersHttproutesRetrieve) | **Get** /api/kubernetes/clusters/{cluster_id}/httproutes/{id}/ | 
[**KubernetesClustersHttproutesUpdate**](KubernetesAPI.md#KubernetesClustersHttproutesUpdate) | **Put** /api/kubernetes/clusters/{cluster_id}/httproutes/{id}/ | 
[**KubernetesClustersKubeVersionUpgradeCreate**](KubernetesAPI.md#KubernetesClustersKubeVersionUpgradeCreate) | **Post** /api/kubernetes/clusters/{id}/kube-version-upgrade/ | 
[**KubernetesClustersKubeconfigCreate**](KubernetesAPI.md#KubernetesClustersKubeconfigCreate) | **Post** /api/kubernetes/clusters/{id}/kubeconfig/ | 
[**KubernetesClustersKubeconfigRetrieve**](KubernetesAPI.md#KubernetesClustersKubeconfigRetrieve) | **Get** /api/kubernetes/clusters/{id}/kubeconfig/ | 
[**KubernetesClustersLbFirewallCreate**](KubernetesAPI.md#KubernetesClustersLbFirewallCreate) | **Post** /api/kubernetes/clusters/{cluster_id}/lb-firewall/ | 
[**KubernetesClustersLbFirewallDestroy**](KubernetesAPI.md#KubernetesClustersLbFirewallDestroy) | **Delete** /api/kubernetes/clusters/{cluster_id}/lb-firewall/{id}/ | 
[**KubernetesClustersLbFirewallList**](KubernetesAPI.md#KubernetesClustersLbFirewallList) | **Get** /api/kubernetes/clusters/{cluster_id}/lb-firewall/ | 
[**KubernetesClustersLbFirewallPartialUpdate**](KubernetesAPI.md#KubernetesClustersLbFirewallPartialUpdate) | **Patch** /api/kubernetes/clusters/{cluster_id}/lb-firewall/{id}/ | 
[**KubernetesClustersLbFirewallRetrieve**](KubernetesAPI.md#KubernetesClustersLbFirewallRetrieve) | **Get** /api/kubernetes/clusters/{cluster_id}/lb-firewall/{id}/ | 
[**KubernetesClustersLbFirewallUpdate**](KubernetesAPI.md#KubernetesClustersLbFirewallUpdate) | **Put** /api/kubernetes/clusters/{cluster_id}/lb-firewall/{id}/ | 
[**KubernetesClustersList**](KubernetesAPI.md#KubernetesClustersList) | **Get** /api/kubernetes/clusters/ | 
[**KubernetesClustersNodeOperationsCancelCreate**](KubernetesAPI.md#KubernetesClustersNodeOperationsCancelCreate) | **Post** /api/kubernetes/clusters/{cluster_id}/node-operations/{id}/cancel/ | 
[**KubernetesClustersNodeOperationsList**](KubernetesAPI.md#KubernetesClustersNodeOperationsList) | **Get** /api/kubernetes/clusters/{cluster_id}/node-operations/ | 
[**KubernetesClustersNodeOperationsResumeCreate**](KubernetesAPI.md#KubernetesClustersNodeOperationsResumeCreate) | **Post** /api/kubernetes/clusters/{cluster_id}/node-operations/{id}/resume/ | 
[**KubernetesClustersNodeOperationsRetrieve**](KubernetesAPI.md#KubernetesClustersNodeOperationsRetrieve) | **Get** /api/kubernetes/clusters/{cluster_id}/node-operations/{id}/ | 
[**KubernetesClustersNodeOperationsRetryCreate**](KubernetesAPI.md#KubernetesClustersNodeOperationsRetryCreate) | **Post** /api/kubernetes/clusters/{cluster_id}/node-operations/{id}/retry/ | 
[**KubernetesClustersPartialUpdate**](KubernetesAPI.md#KubernetesClustersPartialUpdate) | **Patch** /api/kubernetes/clusters/{id}/ | 
[**KubernetesClustersPoolRemovalJournalsList**](KubernetesAPI.md#KubernetesClustersPoolRemovalJournalsList) | **Get** /api/kubernetes/clusters/{cluster_id}/pool-removal-journals/ | 
[**KubernetesClustersPoolRemovalJournalsResumeCreate**](KubernetesAPI.md#KubernetesClustersPoolRemovalJournalsResumeCreate) | **Post** /api/kubernetes/clusters/{cluster_id}/pool-removal-journals/{id}/resume/ | 
[**KubernetesClustersPoolRemovalJournalsRetrieve**](KubernetesAPI.md#KubernetesClustersPoolRemovalJournalsRetrieve) | **Get** /api/kubernetes/clusters/{cluster_id}/pool-removal-journals/{id}/ | 
[**KubernetesClustersPortForwardsCreate**](KubernetesAPI.md#KubernetesClustersPortForwardsCreate) | **Post** /api/kubernetes/clusters/{cluster_id}/port-forwards/ | 
[**KubernetesClustersPortForwardsDestroy**](KubernetesAPI.md#KubernetesClustersPortForwardsDestroy) | **Delete** /api/kubernetes/clusters/{cluster_id}/port-forwards/{id}/ | 
[**KubernetesClustersPortForwardsList**](KubernetesAPI.md#KubernetesClustersPortForwardsList) | **Get** /api/kubernetes/clusters/{cluster_id}/port-forwards/ | 
[**KubernetesClustersPortForwardsPartialUpdate**](KubernetesAPI.md#KubernetesClustersPortForwardsPartialUpdate) | **Patch** /api/kubernetes/clusters/{cluster_id}/port-forwards/{id}/ | 
[**KubernetesClustersPortForwardsRetrieve**](KubernetesAPI.md#KubernetesClustersPortForwardsRetrieve) | **Get** /api/kubernetes/clusters/{cluster_id}/port-forwards/{id}/ | 
[**KubernetesClustersPortForwardsUpdate**](KubernetesAPI.md#KubernetesClustersPortForwardsUpdate) | **Put** /api/kubernetes/clusters/{cluster_id}/port-forwards/{id}/ | 
[**KubernetesClustersResourcePoolsCreate**](KubernetesAPI.md#KubernetesClustersResourcePoolsCreate) | **Post** /api/kubernetes/clusters/{cluster_id}/resource-pools/ | 
[**KubernetesClustersResourcePoolsDestroy**](KubernetesAPI.md#KubernetesClustersResourcePoolsDestroy) | **Delete** /api/kubernetes/clusters/{cluster_id}/resource-pools/{id}/ | 
[**KubernetesClustersResourcePoolsList**](KubernetesAPI.md#KubernetesClustersResourcePoolsList) | **Get** /api/kubernetes/clusters/{cluster_id}/resource-pools/ | 
[**KubernetesClustersResourcePoolsNodesDestroy**](KubernetesAPI.md#KubernetesClustersResourcePoolsNodesDestroy) | **Delete** /api/kubernetes/clusters/{cluster_id}/resource-pools/{pool_id}/nodes/{id}/ | 
[**KubernetesClustersResourcePoolsNodesList**](KubernetesAPI.md#KubernetesClustersResourcePoolsNodesList) | **Get** /api/kubernetes/clusters/{cluster_id}/resource-pools/{pool_id}/nodes/ | 
[**KubernetesClustersResourcePoolsNodesMetricsRetrieve**](KubernetesAPI.md#KubernetesClustersResourcePoolsNodesMetricsRetrieve) | **Get** /api/kubernetes/clusters/{cluster_id}/resource-pools/{pool_id}/nodes/{id}/metrics/ | 
[**KubernetesClustersResourcePoolsNodesRebootCreate**](KubernetesAPI.md#KubernetesClustersResourcePoolsNodesRebootCreate) | **Post** /api/kubernetes/clusters/{cluster_id}/resource-pools/{pool_id}/nodes/{id}/reboot/ | 
[**KubernetesClustersResourcePoolsNodesRetrieve**](KubernetesAPI.md#KubernetesClustersResourcePoolsNodesRetrieve) | **Get** /api/kubernetes/clusters/{cluster_id}/resource-pools/{pool_id}/nodes/{id}/ | 
[**KubernetesClustersResourcePoolsNodesRrdRetrieve**](KubernetesAPI.md#KubernetesClustersResourcePoolsNodesRrdRetrieve) | **Get** /api/kubernetes/clusters/{cluster_id}/resource-pools/{pool_id}/nodes/{id}/rrd/ | 
[**KubernetesClustersResourcePoolsPartialUpdate**](KubernetesAPI.md#KubernetesClustersResourcePoolsPartialUpdate) | **Patch** /api/kubernetes/clusters/{cluster_id}/resource-pools/{id}/ | 
[**KubernetesClustersResourcePoolsRetrieve**](KubernetesAPI.md#KubernetesClustersResourcePoolsRetrieve) | **Get** /api/kubernetes/clusters/{cluster_id}/resource-pools/{id}/ | 
[**KubernetesClustersResourcePoolsUpdate**](KubernetesAPI.md#KubernetesClustersResourcePoolsUpdate) | **Put** /api/kubernetes/clusters/{cluster_id}/resource-pools/{id}/ | 
[**KubernetesClustersRetrieve**](KubernetesAPI.md#KubernetesClustersRetrieve) | **Get** /api/kubernetes/clusters/{id}/ | 
[**KubernetesClustersTalosVersionUpgradeCreate**](KubernetesAPI.md#KubernetesClustersTalosVersionUpgradeCreate) | **Post** /api/kubernetes/clusters/{id}/talos-version-upgrade/ | 
[**KubernetesClustersTcproutesCreate**](KubernetesAPI.md#KubernetesClustersTcproutesCreate) | **Post** /api/kubernetes/clusters/{cluster_id}/tcproutes/ | 
[**KubernetesClustersTcproutesDestroy**](KubernetesAPI.md#KubernetesClustersTcproutesDestroy) | **Delete** /api/kubernetes/clusters/{cluster_id}/tcproutes/{id}/ | 
[**KubernetesClustersTcproutesList**](KubernetesAPI.md#KubernetesClustersTcproutesList) | **Get** /api/kubernetes/clusters/{cluster_id}/tcproutes/ | 
[**KubernetesClustersTcproutesPartialUpdate**](KubernetesAPI.md#KubernetesClustersTcproutesPartialUpdate) | **Patch** /api/kubernetes/clusters/{cluster_id}/tcproutes/{id}/ | 
[**KubernetesClustersTcproutesRetrieve**](KubernetesAPI.md#KubernetesClustersTcproutesRetrieve) | **Get** /api/kubernetes/clusters/{cluster_id}/tcproutes/{id}/ | 
[**KubernetesClustersTcproutesUpdate**](KubernetesAPI.md#KubernetesClustersTcproutesUpdate) | **Put** /api/kubernetes/clusters/{cluster_id}/tcproutes/{id}/ | 
[**KubernetesClustersToggleCloudVmAccessCreate**](KubernetesAPI.md#KubernetesClustersToggleCloudVmAccessCreate) | **Post** /api/kubernetes/clusters/{id}/toggle-cloud-vm-access/ | 
[**KubernetesClustersUdproutesCreate**](KubernetesAPI.md#KubernetesClustersUdproutesCreate) | **Post** /api/kubernetes/clusters/{cluster_id}/udproutes/ | 
[**KubernetesClustersUdproutesDestroy**](KubernetesAPI.md#KubernetesClustersUdproutesDestroy) | **Delete** /api/kubernetes/clusters/{cluster_id}/udproutes/{id}/ | 
[**KubernetesClustersUdproutesList**](KubernetesAPI.md#KubernetesClustersUdproutesList) | **Get** /api/kubernetes/clusters/{cluster_id}/udproutes/ | 
[**KubernetesClustersUdproutesPartialUpdate**](KubernetesAPI.md#KubernetesClustersUdproutesPartialUpdate) | **Patch** /api/kubernetes/clusters/{cluster_id}/udproutes/{id}/ | 
[**KubernetesClustersUdproutesRetrieve**](KubernetesAPI.md#KubernetesClustersUdproutesRetrieve) | **Get** /api/kubernetes/clusters/{cluster_id}/udproutes/{id}/ | 
[**KubernetesClustersUdproutesUpdate**](KubernetesAPI.md#KubernetesClustersUdproutesUpdate) | **Put** /api/kubernetes/clusters/{cluster_id}/udproutes/{id}/ | 
[**KubernetesClustersUpdate**](KubernetesAPI.md#KubernetesClustersUpdate) | **Put** /api/kubernetes/clusters/{id}/ | 
[**KubernetesClustersUpgradeFeatureCreate**](KubernetesAPI.md#KubernetesClustersUpgradeFeatureCreate) | **Post** /api/kubernetes/clusters/{id}/upgrade-feature/ | 
[**KubernetesClustersUpgradeLbCreate**](KubernetesAPI.md#KubernetesClustersUpgradeLbCreate) | **Post** /api/kubernetes/clusters/{id}/upgrade-lb/ | 



## KubernetesClusterTypesList

> PaginatedClusterTypeList KubernetesClusterTypesList(ctx).Page(page).Execute()





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
	resp, r, err := apiClient.KubernetesAPI.KubernetesClusterTypesList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClusterTypesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClusterTypesList`: PaginatedClusterTypeList
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClusterTypesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClusterTypesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedClusterTypeList**](PaginatedClusterTypeList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersConnectVmCreate

> ConnectVMResponse KubernetesClustersConnectVmCreate(ctx, id).ConnectVMRequest(connectVMRequest).Execute()





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
	id := "id_example" // string | 
	connectVMRequest := *openapiclient.NewConnectVMRequest(int32(123)) // ConnectVMRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersConnectVmCreate(context.Background(), id).ConnectVMRequest(connectVMRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersConnectVmCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersConnectVmCreate`: ConnectVMResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersConnectVmCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersConnectVmCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectVMRequest** | [**ConnectVMRequest**](ConnectVMRequest.md) |  | 

### Return type

[**ConnectVMResponse**](ConnectVMResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersConnectedVmsRetrieve

> ConnectedVMsResponse KubernetesClustersConnectedVmsRetrieve(ctx, id).Execute()





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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersConnectedVmsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersConnectedVmsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersConnectedVmsRetrieve`: ConnectedVMsResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersConnectedVmsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersConnectedVmsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectedVMsResponse**](ConnectedVMsResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersCreate

> ClusterAddResponse KubernetesClustersCreate(ctx).ClusterAddRequest(clusterAddRequest).Execute()





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
	clusterAddRequest := *openapiclient.NewClusterAddRequest(openapiclient.ClusterTypeEnum("dev"), "ResourcePoolPackage_example") // ClusterAddRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersCreate(context.Background()).ClusterAddRequest(clusterAddRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersCreate`: ClusterAddResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterAddRequest** | [**ClusterAddRequest**](ClusterAddRequest.md) |  | 

### Return type

[**ClusterAddResponse**](ClusterAddResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersDestroy

> KubernetesClustersDestroy(ctx, id).Execute()





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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubernetesAPI.KubernetesClustersDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersDisconnectVmCreate

> DisconnectVMResponse KubernetesClustersDisconnectVmCreate(ctx, id).DisconnectVMRequest(disconnectVMRequest).Execute()





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
	id := "id_example" // string | 
	disconnectVMRequest := *openapiclient.NewDisconnectVMRequest(int32(123)) // DisconnectVMRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersDisconnectVmCreate(context.Background(), id).DisconnectVMRequest(disconnectVMRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersDisconnectVmCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersDisconnectVmCreate`: DisconnectVMResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersDisconnectVmCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersDisconnectVmCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disconnectVMRequest** | [**DisconnectVMRequest**](DisconnectVMRequest.md) |  | 

### Return type

[**DisconnectVMResponse**](DisconnectVMResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersEligibleVmsRetrieve

> EligibleVMsResponse KubernetesClustersEligibleVmsRetrieve(ctx, id).Execute()





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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersEligibleVmsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersEligibleVmsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersEligibleVmsRetrieve`: EligibleVMsResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersEligibleVmsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersEligibleVmsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EligibleVMsResponse**](EligibleVMsResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersEncryptionCreate

> ClusterEncryptionOperation KubernetesClustersEncryptionCreate(ctx, id).ClusterEncryptionRequest(clusterEncryptionRequest).Execute()





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
	id := "id_example" // string | 
	clusterEncryptionRequest := *openapiclient.NewClusterEncryptionRequest(openapiclient.EncryptionModeEnum("none")) // ClusterEncryptionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersEncryptionCreate(context.Background(), id).ClusterEncryptionRequest(clusterEncryptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersEncryptionCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersEncryptionCreate`: ClusterEncryptionOperation
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersEncryptionCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersEncryptionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterEncryptionRequest** | [**ClusterEncryptionRequest**](ClusterEncryptionRequest.md) |  | 

### Return type

[**ClusterEncryptionOperation**](ClusterEncryptionOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersEncryptionRecheckCreate

> ClusterEncryption KubernetesClustersEncryptionRecheckCreate(ctx, id).Execute()





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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersEncryptionRecheckCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersEncryptionRecheckCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersEncryptionRecheckCreate`: ClusterEncryption
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersEncryptionRecheckCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersEncryptionRecheckCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterEncryption**](ClusterEncryption.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersEncryptionReconcileCreate

> ClusterEncryptionOperation KubernetesClustersEncryptionReconcileCreate(ctx, id).ClusterEncryptionReconcileRequest(clusterEncryptionReconcileRequest).Execute()





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
	id := "id_example" // string | 
	clusterEncryptionReconcileRequest := *openapiclient.NewClusterEncryptionReconcileRequest(openapiclient.EncryptionModeEnum("none")) // ClusterEncryptionReconcileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersEncryptionReconcileCreate(context.Background(), id).ClusterEncryptionReconcileRequest(clusterEncryptionReconcileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersEncryptionReconcileCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersEncryptionReconcileCreate`: ClusterEncryptionOperation
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersEncryptionReconcileCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersEncryptionReconcileCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterEncryptionReconcileRequest** | [**ClusterEncryptionReconcileRequest**](ClusterEncryptionReconcileRequest.md) |  | 

### Return type

[**ClusterEncryptionOperation**](ClusterEncryptionOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersEncryptionRetrieve

> ClusterEncryption KubernetesClustersEncryptionRetrieve(ctx, id).Execute()





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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersEncryptionRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersEncryptionRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersEncryptionRetrieve`: ClusterEncryption
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersEncryptionRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersEncryptionRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterEncryption**](ClusterEncryption.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersHttproutesCreate

> HTTPRoute KubernetesClustersHttproutesCreate(ctx, clusterId).HTTPRouteRequest(hTTPRouteRequest).Execute()





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
	clusterId := int32(56) // int32 | 
	hTTPRouteRequest := *openapiclient.NewHTTPRouteRequest("Name_example", []string{"Hostnames_example"}, "BackendServiceName_example", int32(123)) // HTTPRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersHttproutesCreate(context.Background(), clusterId).HTTPRouteRequest(hTTPRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersHttproutesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersHttproutesCreate`: HTTPRoute
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersHttproutesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersHttproutesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hTTPRouteRequest** | [**HTTPRouteRequest**](HTTPRouteRequest.md) |  | 

### Return type

[**HTTPRoute**](HTTPRoute.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersHttproutesDestroy

> KubernetesClustersHttproutesDestroy(ctx, clusterId, id).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubernetesAPI.KubernetesClustersHttproutesDestroy(context.Background(), clusterId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersHttproutesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersHttproutesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersHttproutesList

> PaginatedHTTPRouteList KubernetesClustersHttproutesList(ctx, clusterId).Page(page).Execute()





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
	clusterId := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersHttproutesList(context.Background(), clusterId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersHttproutesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersHttproutesList`: PaginatedHTTPRouteList
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersHttproutesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersHttproutesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedHTTPRouteList**](PaginatedHTTPRouteList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersHttproutesPartialUpdate

> HTTPRoute KubernetesClustersHttproutesPartialUpdate(ctx, clusterId, id).PatchedHTTPRouteRequest(patchedHTTPRouteRequest).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 
	patchedHTTPRouteRequest := *openapiclient.NewPatchedHTTPRouteRequest() // PatchedHTTPRouteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersHttproutesPartialUpdate(context.Background(), clusterId, id).PatchedHTTPRouteRequest(patchedHTTPRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersHttproutesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersHttproutesPartialUpdate`: HTTPRoute
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersHttproutesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersHttproutesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedHTTPRouteRequest** | [**PatchedHTTPRouteRequest**](PatchedHTTPRouteRequest.md) |  | 

### Return type

[**HTTPRoute**](HTTPRoute.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersHttproutesRetrieve

> HTTPRoute KubernetesClustersHttproutesRetrieve(ctx, clusterId, id).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersHttproutesRetrieve(context.Background(), clusterId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersHttproutesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersHttproutesRetrieve`: HTTPRoute
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersHttproutesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersHttproutesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HTTPRoute**](HTTPRoute.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersHttproutesUpdate

> HTTPRoute KubernetesClustersHttproutesUpdate(ctx, clusterId, id).HTTPRouteRequest(hTTPRouteRequest).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 
	hTTPRouteRequest := *openapiclient.NewHTTPRouteRequest("Name_example", []string{"Hostnames_example"}, "BackendServiceName_example", int32(123)) // HTTPRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersHttproutesUpdate(context.Background(), clusterId, id).HTTPRouteRequest(hTTPRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersHttproutesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersHttproutesUpdate`: HTTPRoute
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersHttproutesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersHttproutesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hTTPRouteRequest** | [**HTTPRouteRequest**](HTTPRouteRequest.md) |  | 

### Return type

[**HTTPRoute**](HTTPRoute.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersKubeVersionUpgradeCreate

> KubeUpgradeResponse KubernetesClustersKubeVersionUpgradeCreate(ctx, id).Execute()





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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersKubeVersionUpgradeCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersKubeVersionUpgradeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersKubeVersionUpgradeCreate`: KubeUpgradeResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersKubeVersionUpgradeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersKubeVersionUpgradeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KubeUpgradeResponse**](KubeUpgradeResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersKubeconfigCreate

> string KubernetesClustersKubeconfigCreate(ctx, id).Execute()





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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersKubeconfigCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersKubeconfigCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersKubeconfigCreate`: string
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersKubeconfigCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersKubeconfigCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersKubeconfigRetrieve

> string KubernetesClustersKubeconfigRetrieve(ctx, id).Execute()





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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersKubeconfigRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersKubeconfigRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersKubeconfigRetrieve`: string
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersKubeconfigRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersKubeconfigRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersLbFirewallCreate

> LBFirewallRule KubernetesClustersLbFirewallCreate(ctx, clusterId).LBFirewallRuleRequest(lBFirewallRuleRequest).Execute()





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
	clusterId := int32(56) // int32 | 
	lBFirewallRuleRequest := *openapiclient.NewLBFirewallRuleRequest() // LBFirewallRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersLbFirewallCreate(context.Background(), clusterId).LBFirewallRuleRequest(lBFirewallRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersLbFirewallCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersLbFirewallCreate`: LBFirewallRule
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersLbFirewallCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersLbFirewallCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lBFirewallRuleRequest** | [**LBFirewallRuleRequest**](LBFirewallRuleRequest.md) |  | 

### Return type

[**LBFirewallRule**](LBFirewallRule.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersLbFirewallDestroy

> KubernetesClustersLbFirewallDestroy(ctx, clusterId, id).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubernetesAPI.KubernetesClustersLbFirewallDestroy(context.Background(), clusterId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersLbFirewallDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersLbFirewallDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersLbFirewallList

> PaginatedLBFirewallRuleList KubernetesClustersLbFirewallList(ctx, clusterId).Page(page).Execute()





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
	clusterId := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersLbFirewallList(context.Background(), clusterId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersLbFirewallList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersLbFirewallList`: PaginatedLBFirewallRuleList
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersLbFirewallList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersLbFirewallListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedLBFirewallRuleList**](PaginatedLBFirewallRuleList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersLbFirewallPartialUpdate

> LBFirewallRule KubernetesClustersLbFirewallPartialUpdate(ctx, clusterId, id).PatchedLBFirewallRuleRequest(patchedLBFirewallRuleRequest).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 
	patchedLBFirewallRuleRequest := *openapiclient.NewPatchedLBFirewallRuleRequest() // PatchedLBFirewallRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersLbFirewallPartialUpdate(context.Background(), clusterId, id).PatchedLBFirewallRuleRequest(patchedLBFirewallRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersLbFirewallPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersLbFirewallPartialUpdate`: LBFirewallRule
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersLbFirewallPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersLbFirewallPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedLBFirewallRuleRequest** | [**PatchedLBFirewallRuleRequest**](PatchedLBFirewallRuleRequest.md) |  | 

### Return type

[**LBFirewallRule**](LBFirewallRule.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersLbFirewallRetrieve

> LBFirewallRule KubernetesClustersLbFirewallRetrieve(ctx, clusterId, id).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersLbFirewallRetrieve(context.Background(), clusterId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersLbFirewallRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersLbFirewallRetrieve`: LBFirewallRule
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersLbFirewallRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersLbFirewallRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LBFirewallRule**](LBFirewallRule.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersLbFirewallUpdate

> LBFirewallRule KubernetesClustersLbFirewallUpdate(ctx, clusterId, id).LBFirewallRuleRequest(lBFirewallRuleRequest).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 
	lBFirewallRuleRequest := *openapiclient.NewLBFirewallRuleRequest() // LBFirewallRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersLbFirewallUpdate(context.Background(), clusterId, id).LBFirewallRuleRequest(lBFirewallRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersLbFirewallUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersLbFirewallUpdate`: LBFirewallRule
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersLbFirewallUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersLbFirewallUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lBFirewallRuleRequest** | [**LBFirewallRuleRequest**](LBFirewallRuleRequest.md) |  | 

### Return type

[**LBFirewallRule**](LBFirewallRule.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersList

> PaginatedClusterDetailList KubernetesClustersList(ctx).Page(page).Execute()





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
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersList`: PaginatedClusterDetailList
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedClusterDetailList**](PaginatedClusterDetailList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersNodeOperationsCancelCreate

> NodeOperation KubernetesClustersNodeOperationsCancelCreate(ctx, clusterId, id).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersNodeOperationsCancelCreate(context.Background(), clusterId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersNodeOperationsCancelCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersNodeOperationsCancelCreate`: NodeOperation
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersNodeOperationsCancelCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersNodeOperationsCancelCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NodeOperation**](NodeOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersNodeOperationsList

> PaginatedNodeOperationList KubernetesClustersNodeOperationsList(ctx, clusterId).Page(page).Execute()





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
	clusterId := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersNodeOperationsList(context.Background(), clusterId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersNodeOperationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersNodeOperationsList`: PaginatedNodeOperationList
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersNodeOperationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersNodeOperationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedNodeOperationList**](PaginatedNodeOperationList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersNodeOperationsResumeCreate

> NodeOperation KubernetesClustersNodeOperationsResumeCreate(ctx, clusterId, id).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersNodeOperationsResumeCreate(context.Background(), clusterId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersNodeOperationsResumeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersNodeOperationsResumeCreate`: NodeOperation
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersNodeOperationsResumeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersNodeOperationsResumeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NodeOperation**](NodeOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersNodeOperationsRetrieve

> NodeOperation KubernetesClustersNodeOperationsRetrieve(ctx, clusterId, id).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersNodeOperationsRetrieve(context.Background(), clusterId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersNodeOperationsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersNodeOperationsRetrieve`: NodeOperation
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersNodeOperationsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersNodeOperationsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NodeOperation**](NodeOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersNodeOperationsRetryCreate

> NodeOperation KubernetesClustersNodeOperationsRetryCreate(ctx, clusterId, id).NodeOperationRetryRequest(nodeOperationRetryRequest).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 
	nodeOperationRetryRequest := *openapiclient.NewNodeOperationRetryRequest() // NodeOperationRetryRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersNodeOperationsRetryCreate(context.Background(), clusterId, id).NodeOperationRetryRequest(nodeOperationRetryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersNodeOperationsRetryCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersNodeOperationsRetryCreate`: NodeOperation
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersNodeOperationsRetryCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersNodeOperationsRetryCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nodeOperationRetryRequest** | [**NodeOperationRetryRequest**](NodeOperationRetryRequest.md) |  | 

### Return type

[**NodeOperation**](NodeOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersPartialUpdate

> ClusterDetail KubernetesClustersPartialUpdate(ctx, id).PatchedClusterDetailRequest(patchedClusterDetailRequest).Execute()





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
	id := "id_example" // string | 
	patchedClusterDetailRequest := *openapiclient.NewPatchedClusterDetailRequest() // PatchedClusterDetailRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersPartialUpdate(context.Background(), id).PatchedClusterDetailRequest(patchedClusterDetailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersPartialUpdate`: ClusterDetail
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedClusterDetailRequest** | [**PatchedClusterDetailRequest**](PatchedClusterDetailRequest.md) |  | 

### Return type

[**ClusterDetail**](ClusterDetail.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersPoolRemovalJournalsList

> PaginatedPoolRemovalJournalList KubernetesClustersPoolRemovalJournalsList(ctx, clusterId).Page(page).Execute()





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
	clusterId := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersPoolRemovalJournalsList(context.Background(), clusterId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersPoolRemovalJournalsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersPoolRemovalJournalsList`: PaginatedPoolRemovalJournalList
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersPoolRemovalJournalsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersPoolRemovalJournalsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedPoolRemovalJournalList**](PaginatedPoolRemovalJournalList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersPoolRemovalJournalsResumeCreate

> PoolRemovalJournal KubernetesClustersPoolRemovalJournalsResumeCreate(ctx, clusterId, id).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersPoolRemovalJournalsResumeCreate(context.Background(), clusterId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersPoolRemovalJournalsResumeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersPoolRemovalJournalsResumeCreate`: PoolRemovalJournal
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersPoolRemovalJournalsResumeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersPoolRemovalJournalsResumeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PoolRemovalJournal**](PoolRemovalJournal.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersPoolRemovalJournalsRetrieve

> PoolRemovalJournal KubernetesClustersPoolRemovalJournalsRetrieve(ctx, clusterId, id).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersPoolRemovalJournalsRetrieve(context.Background(), clusterId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersPoolRemovalJournalsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersPoolRemovalJournalsRetrieve`: PoolRemovalJournal
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersPoolRemovalJournalsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersPoolRemovalJournalsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PoolRemovalJournal**](PoolRemovalJournal.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersPortForwardsCreate

> K8sPortForward KubernetesClustersPortForwardsCreate(ctx, clusterId).K8sPortForwardRequest(k8sPortForwardRequest).Execute()





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
	clusterId := int32(56) // int32 | 
	k8sPortForwardRequest := *openapiclient.NewK8sPortForwardRequest("InternalIp_example", int32(123), openapiclient.ProtocolEnum("tcp")) // K8sPortForwardRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersPortForwardsCreate(context.Background(), clusterId).K8sPortForwardRequest(k8sPortForwardRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersPortForwardsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersPortForwardsCreate`: K8sPortForward
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersPortForwardsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersPortForwardsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **k8sPortForwardRequest** | [**K8sPortForwardRequest**](K8sPortForwardRequest.md) |  | 

### Return type

[**K8sPortForward**](K8sPortForward.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersPortForwardsDestroy

> KubernetesClustersPortForwardsDestroy(ctx, clusterId, id).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubernetesAPI.KubernetesClustersPortForwardsDestroy(context.Background(), clusterId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersPortForwardsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersPortForwardsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersPortForwardsList

> PaginatedK8sPortForwardList KubernetesClustersPortForwardsList(ctx, clusterId).Page(page).Execute()





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
	clusterId := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersPortForwardsList(context.Background(), clusterId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersPortForwardsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersPortForwardsList`: PaginatedK8sPortForwardList
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersPortForwardsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersPortForwardsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedK8sPortForwardList**](PaginatedK8sPortForwardList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersPortForwardsPartialUpdate

> K8sPortForward KubernetesClustersPortForwardsPartialUpdate(ctx, clusterId, id).PatchedK8sPortForwardRequest(patchedK8sPortForwardRequest).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 
	patchedK8sPortForwardRequest := *openapiclient.NewPatchedK8sPortForwardRequest() // PatchedK8sPortForwardRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersPortForwardsPartialUpdate(context.Background(), clusterId, id).PatchedK8sPortForwardRequest(patchedK8sPortForwardRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersPortForwardsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersPortForwardsPartialUpdate`: K8sPortForward
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersPortForwardsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersPortForwardsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedK8sPortForwardRequest** | [**PatchedK8sPortForwardRequest**](PatchedK8sPortForwardRequest.md) |  | 

### Return type

[**K8sPortForward**](K8sPortForward.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersPortForwardsRetrieve

> K8sPortForward KubernetesClustersPortForwardsRetrieve(ctx, clusterId, id).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersPortForwardsRetrieve(context.Background(), clusterId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersPortForwardsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersPortForwardsRetrieve`: K8sPortForward
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersPortForwardsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersPortForwardsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**K8sPortForward**](K8sPortForward.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersPortForwardsUpdate

> K8sPortForward KubernetesClustersPortForwardsUpdate(ctx, clusterId, id).K8sPortForwardRequest(k8sPortForwardRequest).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 
	k8sPortForwardRequest := *openapiclient.NewK8sPortForwardRequest("InternalIp_example", int32(123), openapiclient.ProtocolEnum("tcp")) // K8sPortForwardRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersPortForwardsUpdate(context.Background(), clusterId, id).K8sPortForwardRequest(k8sPortForwardRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersPortForwardsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersPortForwardsUpdate`: K8sPortForward
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersPortForwardsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersPortForwardsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **k8sPortForwardRequest** | [**K8sPortForwardRequest**](K8sPortForwardRequest.md) |  | 

### Return type

[**K8sPortForward**](K8sPortForward.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersResourcePoolsCreate

> ResourcePoolAddResponse KubernetesClustersResourcePoolsCreate(ctx, clusterId).ResourcePoolAddRequest(resourcePoolAddRequest).Execute()





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
	clusterId := int32(56) // int32 | 
	resourcePoolAddRequest := *openapiclient.NewResourcePoolAddRequest("ResourcePoolPackage_example", int32(123)) // ResourcePoolAddRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersResourcePoolsCreate(context.Background(), clusterId).ResourcePoolAddRequest(resourcePoolAddRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersResourcePoolsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersResourcePoolsCreate`: ResourcePoolAddResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersResourcePoolsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersResourcePoolsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourcePoolAddRequest** | [**ResourcePoolAddRequest**](ResourcePoolAddRequest.md) |  | 

### Return type

[**ResourcePoolAddResponse**](ResourcePoolAddResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersResourcePoolsDestroy

> KubernetesClustersResourcePoolsDestroy(ctx, clusterId, id).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubernetesAPI.KubernetesClustersResourcePoolsDestroy(context.Background(), clusterId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersResourcePoolsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersResourcePoolsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersResourcePoolsList

> PaginatedResourcePoolList KubernetesClustersResourcePoolsList(ctx, clusterId).Page(page).Execute()





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
	clusterId := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersResourcePoolsList(context.Background(), clusterId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersResourcePoolsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersResourcePoolsList`: PaginatedResourcePoolList
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersResourcePoolsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersResourcePoolsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedResourcePoolList**](PaginatedResourcePoolList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersResourcePoolsNodesDestroy

> NodeOperation KubernetesClustersResourcePoolsNodesDestroy(ctx, clusterId, id, poolId).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 
	poolId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersResourcePoolsNodesDestroy(context.Background(), clusterId, id, poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersResourcePoolsNodesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersResourcePoolsNodesDestroy`: NodeOperation
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersResourcePoolsNodesDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 
**poolId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersResourcePoolsNodesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**NodeOperation**](NodeOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersResourcePoolsNodesList

> PaginatedResourcePoolNodeList KubernetesClustersResourcePoolsNodesList(ctx, clusterId, poolId).Page(page).Execute()





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
	clusterId := int32(56) // int32 | 
	poolId := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersResourcePoolsNodesList(context.Background(), clusterId, poolId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersResourcePoolsNodesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersResourcePoolsNodesList`: PaginatedResourcePoolNodeList
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersResourcePoolsNodesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**poolId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersResourcePoolsNodesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedResourcePoolNodeList**](PaginatedResourcePoolNodeList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersResourcePoolsNodesMetricsRetrieve

> NodeMetricsResponse KubernetesClustersResourcePoolsNodesMetricsRetrieve(ctx, clusterId, id, poolId).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 
	poolId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersResourcePoolsNodesMetricsRetrieve(context.Background(), clusterId, id, poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersResourcePoolsNodesMetricsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersResourcePoolsNodesMetricsRetrieve`: NodeMetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersResourcePoolsNodesMetricsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 
**poolId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersResourcePoolsNodesMetricsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**NodeMetricsResponse**](NodeMetricsResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersResourcePoolsNodesRebootCreate

> NodeOperation KubernetesClustersResourcePoolsNodesRebootCreate(ctx, clusterId, id, poolId).NodeOperationRebootRequest(nodeOperationRebootRequest).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 
	poolId := int32(56) // int32 | 
	nodeOperationRebootRequest := *openapiclient.NewNodeOperationRebootRequest() // NodeOperationRebootRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersResourcePoolsNodesRebootCreate(context.Background(), clusterId, id, poolId).NodeOperationRebootRequest(nodeOperationRebootRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersResourcePoolsNodesRebootCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersResourcePoolsNodesRebootCreate`: NodeOperation
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersResourcePoolsNodesRebootCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 
**poolId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersResourcePoolsNodesRebootCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **nodeOperationRebootRequest** | [**NodeOperationRebootRequest**](NodeOperationRebootRequest.md) |  | 

### Return type

[**NodeOperation**](NodeOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersResourcePoolsNodesRetrieve

> ResourcePoolNode KubernetesClustersResourcePoolsNodesRetrieve(ctx, clusterId, id, poolId).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 
	poolId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersResourcePoolsNodesRetrieve(context.Background(), clusterId, id, poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersResourcePoolsNodesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersResourcePoolsNodesRetrieve`: ResourcePoolNode
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersResourcePoolsNodesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 
**poolId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersResourcePoolsNodesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ResourcePoolNode**](ResourcePoolNode.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersResourcePoolsNodesRrdRetrieve

> NodeRRDResponse KubernetesClustersResourcePoolsNodesRrdRetrieve(ctx, clusterId, id, poolId).Timeframe(timeframe).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 
	poolId := int32(56) // int32 | 
	timeframe := "timeframe_example" // string | Window of recorded data to return. (optional) (default to "hour")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersResourcePoolsNodesRrdRetrieve(context.Background(), clusterId, id, poolId).Timeframe(timeframe).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersResourcePoolsNodesRrdRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersResourcePoolsNodesRrdRetrieve`: NodeRRDResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersResourcePoolsNodesRrdRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 
**poolId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersResourcePoolsNodesRrdRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **timeframe** | **string** | Window of recorded data to return. | [default to &quot;hour&quot;]

### Return type

[**NodeRRDResponse**](NodeRRDResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersResourcePoolsPartialUpdate

> ResourcePool KubernetesClustersResourcePoolsPartialUpdate(ctx, clusterId, id).PatchedResourcePoolRequest(patchedResourcePoolRequest).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 
	patchedResourcePoolRequest := *openapiclient.NewPatchedResourcePoolRequest() // PatchedResourcePoolRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersResourcePoolsPartialUpdate(context.Background(), clusterId, id).PatchedResourcePoolRequest(patchedResourcePoolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersResourcePoolsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersResourcePoolsPartialUpdate`: ResourcePool
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersResourcePoolsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersResourcePoolsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedResourcePoolRequest** | [**PatchedResourcePoolRequest**](PatchedResourcePoolRequest.md) |  | 

### Return type

[**ResourcePool**](ResourcePool.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersResourcePoolsRetrieve

> ResourcePool KubernetesClustersResourcePoolsRetrieve(ctx, clusterId, id).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersResourcePoolsRetrieve(context.Background(), clusterId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersResourcePoolsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersResourcePoolsRetrieve`: ResourcePool
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersResourcePoolsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersResourcePoolsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResourcePool**](ResourcePool.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersResourcePoolsUpdate

> ResourcePool KubernetesClustersResourcePoolsUpdate(ctx, clusterId, id).ResourcePoolRequest(resourcePoolRequest).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 
	resourcePoolRequest := *openapiclient.NewResourcePoolRequest() // ResourcePoolRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersResourcePoolsUpdate(context.Background(), clusterId, id).ResourcePoolRequest(resourcePoolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersResourcePoolsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersResourcePoolsUpdate`: ResourcePool
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersResourcePoolsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersResourcePoolsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourcePoolRequest** | [**ResourcePoolRequest**](ResourcePoolRequest.md) |  | 

### Return type

[**ResourcePool**](ResourcePool.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersRetrieve

> ClusterDetail KubernetesClustersRetrieve(ctx, id).Execute()





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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersRetrieve`: ClusterDetail
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterDetail**](ClusterDetail.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersTalosVersionUpgradeCreate

> TalosUpgradeResponse KubernetesClustersTalosVersionUpgradeCreate(ctx, id).Execute()





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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersTalosVersionUpgradeCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersTalosVersionUpgradeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersTalosVersionUpgradeCreate`: TalosUpgradeResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersTalosVersionUpgradeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersTalosVersionUpgradeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TalosUpgradeResponse**](TalosUpgradeResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersTcproutesCreate

> TCPRoute KubernetesClustersTcproutesCreate(ctx, clusterId).TCPRouteRequest(tCPRouteRequest).Execute()





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
	clusterId := int32(56) // int32 | 
	tCPRouteRequest := *openapiclient.NewTCPRouteRequest("Name_example", int32(123), "BackendServiceName_example", int32(123)) // TCPRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersTcproutesCreate(context.Background(), clusterId).TCPRouteRequest(tCPRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersTcproutesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersTcproutesCreate`: TCPRoute
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersTcproutesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersTcproutesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tCPRouteRequest** | [**TCPRouteRequest**](TCPRouteRequest.md) |  | 

### Return type

[**TCPRoute**](TCPRoute.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersTcproutesDestroy

> KubernetesClustersTcproutesDestroy(ctx, clusterId, id).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubernetesAPI.KubernetesClustersTcproutesDestroy(context.Background(), clusterId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersTcproutesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersTcproutesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersTcproutesList

> PaginatedTCPRouteList KubernetesClustersTcproutesList(ctx, clusterId).Page(page).Execute()





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
	clusterId := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersTcproutesList(context.Background(), clusterId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersTcproutesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersTcproutesList`: PaginatedTCPRouteList
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersTcproutesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersTcproutesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedTCPRouteList**](PaginatedTCPRouteList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersTcproutesPartialUpdate

> TCPRoute KubernetesClustersTcproutesPartialUpdate(ctx, clusterId, id).PatchedTCPRouteRequest(patchedTCPRouteRequest).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 
	patchedTCPRouteRequest := *openapiclient.NewPatchedTCPRouteRequest() // PatchedTCPRouteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersTcproutesPartialUpdate(context.Background(), clusterId, id).PatchedTCPRouteRequest(patchedTCPRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersTcproutesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersTcproutesPartialUpdate`: TCPRoute
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersTcproutesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersTcproutesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedTCPRouteRequest** | [**PatchedTCPRouteRequest**](PatchedTCPRouteRequest.md) |  | 

### Return type

[**TCPRoute**](TCPRoute.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersTcproutesRetrieve

> TCPRoute KubernetesClustersTcproutesRetrieve(ctx, clusterId, id).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersTcproutesRetrieve(context.Background(), clusterId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersTcproutesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersTcproutesRetrieve`: TCPRoute
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersTcproutesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersTcproutesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TCPRoute**](TCPRoute.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersTcproutesUpdate

> TCPRoute KubernetesClustersTcproutesUpdate(ctx, clusterId, id).TCPRouteRequest(tCPRouteRequest).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 
	tCPRouteRequest := *openapiclient.NewTCPRouteRequest("Name_example", int32(123), "BackendServiceName_example", int32(123)) // TCPRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersTcproutesUpdate(context.Background(), clusterId, id).TCPRouteRequest(tCPRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersTcproutesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersTcproutesUpdate`: TCPRoute
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersTcproutesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersTcproutesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tCPRouteRequest** | [**TCPRouteRequest**](TCPRouteRequest.md) |  | 

### Return type

[**TCPRoute**](TCPRoute.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersToggleCloudVmAccessCreate

> ToggleCloudVMAccessResponse KubernetesClustersToggleCloudVmAccessCreate(ctx, id).Execute()





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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersToggleCloudVmAccessCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersToggleCloudVmAccessCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersToggleCloudVmAccessCreate`: ToggleCloudVMAccessResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersToggleCloudVmAccessCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersToggleCloudVmAccessCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ToggleCloudVMAccessResponse**](ToggleCloudVMAccessResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersUdproutesCreate

> UDPRoute KubernetesClustersUdproutesCreate(ctx, clusterId).UDPRouteRequest(uDPRouteRequest).Execute()





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
	clusterId := int32(56) // int32 | 
	uDPRouteRequest := *openapiclient.NewUDPRouteRequest("Name_example", int32(123), "BackendServiceName_example", int32(123)) // UDPRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersUdproutesCreate(context.Background(), clusterId).UDPRouteRequest(uDPRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersUdproutesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersUdproutesCreate`: UDPRoute
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersUdproutesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersUdproutesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uDPRouteRequest** | [**UDPRouteRequest**](UDPRouteRequest.md) |  | 

### Return type

[**UDPRoute**](UDPRoute.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersUdproutesDestroy

> KubernetesClustersUdproutesDestroy(ctx, clusterId, id).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KubernetesAPI.KubernetesClustersUdproutesDestroy(context.Background(), clusterId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersUdproutesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersUdproutesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersUdproutesList

> PaginatedUDPRouteList KubernetesClustersUdproutesList(ctx, clusterId).Page(page).Execute()





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
	clusterId := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersUdproutesList(context.Background(), clusterId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersUdproutesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersUdproutesList`: PaginatedUDPRouteList
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersUdproutesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersUdproutesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedUDPRouteList**](PaginatedUDPRouteList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersUdproutesPartialUpdate

> UDPRoute KubernetesClustersUdproutesPartialUpdate(ctx, clusterId, id).PatchedUDPRouteRequest(patchedUDPRouteRequest).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 
	patchedUDPRouteRequest := *openapiclient.NewPatchedUDPRouteRequest() // PatchedUDPRouteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersUdproutesPartialUpdate(context.Background(), clusterId, id).PatchedUDPRouteRequest(patchedUDPRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersUdproutesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersUdproutesPartialUpdate`: UDPRoute
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersUdproutesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersUdproutesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedUDPRouteRequest** | [**PatchedUDPRouteRequest**](PatchedUDPRouteRequest.md) |  | 

### Return type

[**UDPRoute**](UDPRoute.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersUdproutesRetrieve

> UDPRoute KubernetesClustersUdproutesRetrieve(ctx, clusterId, id).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersUdproutesRetrieve(context.Background(), clusterId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersUdproutesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersUdproutesRetrieve`: UDPRoute
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersUdproutesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersUdproutesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UDPRoute**](UDPRoute.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersUdproutesUpdate

> UDPRoute KubernetesClustersUdproutesUpdate(ctx, clusterId, id).UDPRouteRequest(uDPRouteRequest).Execute()





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
	clusterId := int32(56) // int32 | 
	id := "id_example" // string | 
	uDPRouteRequest := *openapiclient.NewUDPRouteRequest("Name_example", int32(123), "BackendServiceName_example", int32(123)) // UDPRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersUdproutesUpdate(context.Background(), clusterId, id).UDPRouteRequest(uDPRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersUdproutesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersUdproutesUpdate`: UDPRoute
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersUdproutesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersUdproutesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uDPRouteRequest** | [**UDPRouteRequest**](UDPRouteRequest.md) |  | 

### Return type

[**UDPRoute**](UDPRoute.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersUpdate

> ClusterDetail KubernetesClustersUpdate(ctx, id).ClusterDetailRequest(clusterDetailRequest).Execute()





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
	id := "id_example" // string | 
	clusterDetailRequest := *openapiclient.NewClusterDetailRequest("PricePerMonth_example") // ClusterDetailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersUpdate(context.Background(), id).ClusterDetailRequest(clusterDetailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersUpdate`: ClusterDetail
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterDetailRequest** | [**ClusterDetailRequest**](ClusterDetailRequest.md) |  | 

### Return type

[**ClusterDetail**](ClusterDetail.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersUpgradeFeatureCreate

> FeatureUpgradeResponse KubernetesClustersUpgradeFeatureCreate(ctx, id).FeatureUpgradeRequest(featureUpgradeRequest).Execute()





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
	id := "id_example" // string | 
	featureUpgradeRequest := *openapiclient.NewFeatureUpgradeRequest("FeatureName_example") // FeatureUpgradeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersUpgradeFeatureCreate(context.Background(), id).FeatureUpgradeRequest(featureUpgradeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersUpgradeFeatureCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersUpgradeFeatureCreate`: FeatureUpgradeResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersUpgradeFeatureCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersUpgradeFeatureCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **featureUpgradeRequest** | [**FeatureUpgradeRequest**](FeatureUpgradeRequest.md) |  | 

### Return type

[**FeatureUpgradeResponse**](FeatureUpgradeResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClustersUpgradeLbCreate

> LBUpgradePlanResponse KubernetesClustersUpgradeLbCreate(ctx, id).LBUpgradeRequest(lBUpgradeRequest).Execute()





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
	id := "id_example" // string | 
	lBUpgradeRequest := *openapiclient.NewLBUpgradeRequest() // LBUpgradeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubernetesAPI.KubernetesClustersUpgradeLbCreate(context.Background(), id).LBUpgradeRequest(lBUpgradeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.KubernetesClustersUpgradeLbCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KubernetesClustersUpgradeLbCreate`: LBUpgradePlanResponse
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.KubernetesClustersUpgradeLbCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKubernetesClustersUpgradeLbCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lBUpgradeRequest** | [**LBUpgradeRequest**](LBUpgradeRequest.md) |  | 

### Return type

[**LBUpgradePlanResponse**](LBUpgradePlanResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

