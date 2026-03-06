# \DedicatedAPI

All URIs are relative to *https://www.pidginhost.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DedicatedServersList**](DedicatedAPI.md#DedicatedServersList) | **Get** /api/dedicated/servers/ | 
[**DedicatedServersList2**](DedicatedAPI.md#DedicatedServersList2) | **Get** /api/v1/dedicated/servers/ | 
[**DedicatedServersPowerCreate**](DedicatedAPI.md#DedicatedServersPowerCreate) | **Post** /api/dedicated/servers/{id}/power/ | 
[**DedicatedServersPowerCreate2**](DedicatedAPI.md#DedicatedServersPowerCreate2) | **Post** /api/v1/dedicated/servers/{id}/power/ | 
[**DedicatedServersRdnsCreate**](DedicatedAPI.md#DedicatedServersRdnsCreate) | **Post** /api/dedicated/servers/{id}/rdns/ | 
[**DedicatedServersRdnsCreate2**](DedicatedAPI.md#DedicatedServersRdnsCreate2) | **Post** /api/v1/dedicated/servers/{id}/rdns/ | 
[**DedicatedServersReinstallCreate**](DedicatedAPI.md#DedicatedServersReinstallCreate) | **Post** /api/dedicated/servers/{id}/reinstall/ | 
[**DedicatedServersReinstallCreate2**](DedicatedAPI.md#DedicatedServersReinstallCreate2) | **Post** /api/v1/dedicated/servers/{id}/reinstall/ | 
[**DedicatedServersRetrieve**](DedicatedAPI.md#DedicatedServersRetrieve) | **Get** /api/dedicated/servers/{id}/ | 
[**DedicatedServersRetrieve2**](DedicatedAPI.md#DedicatedServersRetrieve2) | **Get** /api/v1/dedicated/servers/{id}/ | 



## DedicatedServersList

> PaginatedDedicatedServerList DedicatedServersList(ctx).Page(page).Execute()





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
	resp, r, err := apiClient.DedicatedAPI.DedicatedServersList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedAPI.DedicatedServersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DedicatedServersList`: PaginatedDedicatedServerList
	fmt.Fprintf(os.Stdout, "Response from `DedicatedAPI.DedicatedServersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDedicatedServersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedDedicatedServerList**](PaginatedDedicatedServerList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DedicatedServersList2

> PaginatedDedicatedServerList DedicatedServersList2(ctx).Page(page).Execute()





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
	resp, r, err := apiClient.DedicatedAPI.DedicatedServersList2(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedAPI.DedicatedServersList2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DedicatedServersList2`: PaginatedDedicatedServerList
	fmt.Fprintf(os.Stdout, "Response from `DedicatedAPI.DedicatedServersList2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDedicatedServersList2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedDedicatedServerList**](PaginatedDedicatedServerList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DedicatedServersPowerCreate

> PowerActionResponse DedicatedServersPowerCreate(ctx, id).PowerAction(powerAction).Execute()





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
	powerAction := *openapiclient.NewPowerAction(openapiclient.PowerActionActionEnum("start")) // PowerAction | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedAPI.DedicatedServersPowerCreate(context.Background(), id).PowerAction(powerAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedAPI.DedicatedServersPowerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DedicatedServersPowerCreate`: PowerActionResponse
	fmt.Fprintf(os.Stdout, "Response from `DedicatedAPI.DedicatedServersPowerCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDedicatedServersPowerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **powerAction** | [**PowerAction**](PowerAction.md) |  | 

### Return type

[**PowerActionResponse**](PowerActionResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DedicatedServersPowerCreate2

> PowerActionResponse DedicatedServersPowerCreate2(ctx, id).PowerAction(powerAction).Execute()





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
	powerAction := *openapiclient.NewPowerAction(openapiclient.PowerActionActionEnum("start")) // PowerAction | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedAPI.DedicatedServersPowerCreate2(context.Background(), id).PowerAction(powerAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedAPI.DedicatedServersPowerCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DedicatedServersPowerCreate2`: PowerActionResponse
	fmt.Fprintf(os.Stdout, "Response from `DedicatedAPI.DedicatedServersPowerCreate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDedicatedServersPowerCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **powerAction** | [**PowerAction**](PowerAction.md) |  | 

### Return type

[**PowerActionResponse**](PowerActionResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DedicatedServersRdnsCreate

> RDNSUpdateResponse DedicatedServersRdnsCreate(ctx, id).DedicatedRDNS(dedicatedRDNS).Execute()





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
	dedicatedRDNS := *openapiclient.NewDedicatedRDNS(int32(123), "ReverseDns_example") // DedicatedRDNS | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedAPI.DedicatedServersRdnsCreate(context.Background(), id).DedicatedRDNS(dedicatedRDNS).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedAPI.DedicatedServersRdnsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DedicatedServersRdnsCreate`: RDNSUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `DedicatedAPI.DedicatedServersRdnsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDedicatedServersRdnsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dedicatedRDNS** | [**DedicatedRDNS**](DedicatedRDNS.md) |  | 

### Return type

[**RDNSUpdateResponse**](RDNSUpdateResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DedicatedServersRdnsCreate2

> RDNSUpdateResponse DedicatedServersRdnsCreate2(ctx, id).DedicatedRDNS(dedicatedRDNS).Execute()





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
	dedicatedRDNS := *openapiclient.NewDedicatedRDNS(int32(123), "ReverseDns_example") // DedicatedRDNS | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedAPI.DedicatedServersRdnsCreate2(context.Background(), id).DedicatedRDNS(dedicatedRDNS).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedAPI.DedicatedServersRdnsCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DedicatedServersRdnsCreate2`: RDNSUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `DedicatedAPI.DedicatedServersRdnsCreate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDedicatedServersRdnsCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dedicatedRDNS** | [**DedicatedRDNS**](DedicatedRDNS.md) |  | 

### Return type

[**RDNSUpdateResponse**](RDNSUpdateResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DedicatedServersReinstallCreate

> ReinstallResponse DedicatedServersReinstallCreate(ctx, id).Reinstall(reinstall).Execute()





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
	reinstall := *openapiclient.NewReinstall(int32(123)) // Reinstall | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedAPI.DedicatedServersReinstallCreate(context.Background(), id).Reinstall(reinstall).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedAPI.DedicatedServersReinstallCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DedicatedServersReinstallCreate`: ReinstallResponse
	fmt.Fprintf(os.Stdout, "Response from `DedicatedAPI.DedicatedServersReinstallCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDedicatedServersReinstallCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reinstall** | [**Reinstall**](Reinstall.md) |  | 

### Return type

[**ReinstallResponse**](ReinstallResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DedicatedServersReinstallCreate2

> ReinstallResponse DedicatedServersReinstallCreate2(ctx, id).Reinstall(reinstall).Execute()





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
	reinstall := *openapiclient.NewReinstall(int32(123)) // Reinstall | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedAPI.DedicatedServersReinstallCreate2(context.Background(), id).Reinstall(reinstall).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedAPI.DedicatedServersReinstallCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DedicatedServersReinstallCreate2`: ReinstallResponse
	fmt.Fprintf(os.Stdout, "Response from `DedicatedAPI.DedicatedServersReinstallCreate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDedicatedServersReinstallCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reinstall** | [**Reinstall**](Reinstall.md) |  | 

### Return type

[**ReinstallResponse**](ReinstallResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DedicatedServersRetrieve

> DedicatedServer DedicatedServersRetrieve(ctx, id).Execute()





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
	resp, r, err := apiClient.DedicatedAPI.DedicatedServersRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedAPI.DedicatedServersRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DedicatedServersRetrieve`: DedicatedServer
	fmt.Fprintf(os.Stdout, "Response from `DedicatedAPI.DedicatedServersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDedicatedServersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DedicatedServer**](DedicatedServer.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DedicatedServersRetrieve2

> DedicatedServer DedicatedServersRetrieve2(ctx, id).Execute()





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
	resp, r, err := apiClient.DedicatedAPI.DedicatedServersRetrieve2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedAPI.DedicatedServersRetrieve2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DedicatedServersRetrieve2`: DedicatedServer
	fmt.Fprintf(os.Stdout, "Response from `DedicatedAPI.DedicatedServersRetrieve2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDedicatedServersRetrieve2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DedicatedServer**](DedicatedServer.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

