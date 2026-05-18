# \FreednsAPI

All URIs are relative to *https://www.pidginhost.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FreednsDnsActivateCreate**](FreednsAPI.md#FreednsDnsActivateCreate) | **Post** /api/freedns/dns/activate/ | 
[**FreednsDnsAddRecordCreate**](FreednsAPI.md#FreednsDnsAddRecordCreate) | **Post** /api/freedns/dns/add-record/ | 
[**FreednsDnsDeactivateCreate**](FreednsAPI.md#FreednsDnsDeactivateCreate) | **Post** /api/freedns/dns/deactivate/ | 
[**FreednsDnsDeleteRecordCreate**](FreednsAPI.md#FreednsDnsDeleteRecordCreate) | **Post** /api/freedns/dns/delete-record/ | 
[**FreednsDnsList**](FreednsAPI.md#FreednsDnsList) | **Get** /api/freedns/dns/ | 
[**FreednsDnsRecordsList**](FreednsAPI.md#FreednsDnsRecordsList) | **Get** /api/freedns/dns/records/ | 



## FreednsDnsActivateCreate

> ActivateFreeDNSResponse FreednsDnsActivateCreate(ctx).ActivateFreeDNS(activateFreeDNS).Execute()





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
	activateFreeDNS := *openapiclient.NewActivateFreeDNS("Domain_example", openapiclient.SourceEnum("internal"), "Ip_example") // ActivateFreeDNS | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FreednsAPI.FreednsDnsActivateCreate(context.Background()).ActivateFreeDNS(activateFreeDNS).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FreednsAPI.FreednsDnsActivateCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FreednsDnsActivateCreate`: ActivateFreeDNSResponse
	fmt.Fprintf(os.Stdout, "Response from `FreednsAPI.FreednsDnsActivateCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFreednsDnsActivateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activateFreeDNS** | [**ActivateFreeDNS**](ActivateFreeDNS.md) |  | 

### Return type

[**ActivateFreeDNSResponse**](ActivateFreeDNSResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FreednsDnsAddRecordCreate

> DNSRecordMutateResponse FreednsDnsAddRecordCreate(ctx).Domain(domain).Source(source).DNSRecordCreate(dNSRecordCreate).Execute()





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
	domain := "domain_example" // string | Domain name or PK.
	source := "source_example" // string | 'internal' or 'external'.
	dNSRecordCreate := *openapiclient.NewDNSRecordCreate("Name_example", int32(123), openapiclient.DNSRecordCreateTypeEnum("A")) // DNSRecordCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FreednsAPI.FreednsDnsAddRecordCreate(context.Background()).Domain(domain).Source(source).DNSRecordCreate(dNSRecordCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FreednsAPI.FreednsDnsAddRecordCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FreednsDnsAddRecordCreate`: DNSRecordMutateResponse
	fmt.Fprintf(os.Stdout, "Response from `FreednsAPI.FreednsDnsAddRecordCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFreednsDnsAddRecordCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | Domain name or PK. | 
 **source** | **string** | &#39;internal&#39; or &#39;external&#39;. | 
 **dNSRecordCreate** | [**DNSRecordCreate**](DNSRecordCreate.md) |  | 

### Return type

[**DNSRecordMutateResponse**](DNSRecordMutateResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FreednsDnsDeactivateCreate

> DeactivateFreeDNSResponse FreednsDnsDeactivateCreate(ctx).DeactivateFreeDNS(deactivateFreeDNS).Execute()





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
	deactivateFreeDNS := *openapiclient.NewDeactivateFreeDNS("Domain_example", openapiclient.SourceEnum("internal")) // DeactivateFreeDNS | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FreednsAPI.FreednsDnsDeactivateCreate(context.Background()).DeactivateFreeDNS(deactivateFreeDNS).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FreednsAPI.FreednsDnsDeactivateCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FreednsDnsDeactivateCreate`: DeactivateFreeDNSResponse
	fmt.Fprintf(os.Stdout, "Response from `FreednsAPI.FreednsDnsDeactivateCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFreednsDnsDeactivateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deactivateFreeDNS** | [**DeactivateFreeDNS**](DeactivateFreeDNS.md) |  | 

### Return type

[**DeactivateFreeDNSResponse**](DeactivateFreeDNSResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FreednsDnsDeleteRecordCreate

> DeleteRecordResponse FreednsDnsDeleteRecordCreate(ctx).Domain(domain).Source(source).DeleteRecord(deleteRecord).Execute()





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
	domain := "domain_example" // string | Domain name or PK.
	source := "source_example" // string | 'internal' or 'external'.
	deleteRecord := *openapiclient.NewDeleteRecord(int32(123)) // DeleteRecord | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FreednsAPI.FreednsDnsDeleteRecordCreate(context.Background()).Domain(domain).Source(source).DeleteRecord(deleteRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FreednsAPI.FreednsDnsDeleteRecordCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FreednsDnsDeleteRecordCreate`: DeleteRecordResponse
	fmt.Fprintf(os.Stdout, "Response from `FreednsAPI.FreednsDnsDeleteRecordCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFreednsDnsDeleteRecordCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | Domain name or PK. | 
 **source** | **string** | &#39;internal&#39; or &#39;external&#39;. | 
 **deleteRecord** | [**DeleteRecord**](DeleteRecord.md) |  | 

### Return type

[**DeleteRecordResponse**](DeleteRecordResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FreednsDnsList

> []FreeDNSDomain FreednsDnsList(ctx).Execute()





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
	resp, r, err := apiClient.FreednsAPI.FreednsDnsList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FreednsAPI.FreednsDnsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FreednsDnsList`: []FreeDNSDomain
	fmt.Fprintf(os.Stdout, "Response from `FreednsAPI.FreednsDnsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFreednsDnsListRequest struct via the builder pattern


### Return type

[**[]FreeDNSDomain**](FreeDNSDomain.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FreednsDnsRecordsList

> []DNSRecord FreednsDnsRecordsList(ctx).Domain(domain).Source(source).Execute()





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
	domain := "domain_example" // string | Domain name or PK.
	source := "source_example" // string | 'internal' or 'external'.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FreednsAPI.FreednsDnsRecordsList(context.Background()).Domain(domain).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FreednsAPI.FreednsDnsRecordsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FreednsDnsRecordsList`: []DNSRecord
	fmt.Fprintf(os.Stdout, "Response from `FreednsAPI.FreednsDnsRecordsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFreednsDnsRecordsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | Domain name or PK. | 
 **source** | **string** | &#39;internal&#39; or &#39;external&#39;. | 

### Return type

[**[]DNSRecord**](DNSRecord.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

