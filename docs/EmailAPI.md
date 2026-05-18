# \EmailAPI

All URIs are relative to *https://www.pidginhost.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmailApiCredentialsCreate**](EmailAPI.md#EmailApiCredentialsCreate) | **Post** /api/email/api_credentials/ | 
[**EmailApiCredentialsDestroy**](EmailAPI.md#EmailApiCredentialsDestroy) | **Delete** /api/email/api_credentials/{id}/ | 
[**EmailApiCredentialsList**](EmailAPI.md#EmailApiCredentialsList) | **Get** /api/email/api_credentials/ | 
[**EmailApiCredentialsRetrieve**](EmailAPI.md#EmailApiCredentialsRetrieve) | **Get** /api/email/api_credentials/{id}/ | 
[**EmailDomainsCreate**](EmailAPI.md#EmailDomainsCreate) | **Post** /api/email/domains/ | 
[**EmailDomainsInboundRoutesCreate**](EmailAPI.md#EmailDomainsInboundRoutesCreate) | **Post** /api/email/domains/{domain_pk}/inbound_routes/ | 
[**EmailDomainsInboundRoutesList**](EmailAPI.md#EmailDomainsInboundRoutesList) | **Get** /api/email/domains/{domain_pk}/inbound_routes/ | 
[**EmailDomainsList**](EmailAPI.md#EmailDomainsList) | **Get** /api/email/domains/ | 
[**EmailDomainsRetrieve**](EmailAPI.md#EmailDomainsRetrieve) | **Get** /api/email/domains/{id}/ | 
[**EmailDomainsRotateDkimCreate**](EmailAPI.md#EmailDomainsRotateDkimCreate) | **Post** /api/email/domains/{id}/rotate_dkim/ | 
[**EmailDomainsToggleInboundCreate**](EmailAPI.md#EmailDomainsToggleInboundCreate) | **Post** /api/email/domains/{id}/toggle_inbound/ | 
[**EmailDomainsVerifyCreate**](EmailAPI.md#EmailDomainsVerifyCreate) | **Post** /api/email/domains/{id}/verify/ | 
[**EmailInboundRoutesCreate**](EmailAPI.md#EmailInboundRoutesCreate) | **Post** /api/email/inbound_routes/ | 
[**EmailInboundRoutesDestroy**](EmailAPI.md#EmailInboundRoutesDestroy) | **Delete** /api/email/inbound_routes/{id}/ | 
[**EmailInboundRoutesList**](EmailAPI.md#EmailInboundRoutesList) | **Get** /api/email/inbound_routes/ | 
[**EmailInboundRoutesPartialUpdate**](EmailAPI.md#EmailInboundRoutesPartialUpdate) | **Patch** /api/email/inbound_routes/{id}/ | 
[**EmailInboundRoutesRetrieve**](EmailAPI.md#EmailInboundRoutesRetrieve) | **Get** /api/email/inbound_routes/{id}/ | 
[**EmailMessagesRetrieve**](EmailAPI.md#EmailMessagesRetrieve) | **Get** /api/email/messages/{message_id}/ | 
[**EmailSandboxAddressesCreate**](EmailAPI.md#EmailSandboxAddressesCreate) | **Post** /api/email/sandbox_addresses/ | 
[**EmailSandboxAddressesDestroy**](EmailAPI.md#EmailSandboxAddressesDestroy) | **Delete** /api/email/sandbox_addresses/{id}/ | 
[**EmailSandboxAddressesList**](EmailAPI.md#EmailSandboxAddressesList) | **Get** /api/email/sandbox_addresses/ | 
[**EmailSandboxAddressesRetrieve**](EmailAPI.md#EmailSandboxAddressesRetrieve) | **Get** /api/email/sandbox_addresses/{id}/ | 
[**EmailSendCreate**](EmailAPI.md#EmailSendCreate) | **Post** /api/email/send/ | 
[**EmailServicesApiCredentialsCreate**](EmailAPI.md#EmailServicesApiCredentialsCreate) | **Post** /api/email/services/{service_pk}/api_credentials/ | 
[**EmailServicesApiCredentialsList**](EmailAPI.md#EmailServicesApiCredentialsList) | **Get** /api/email/services/{service_pk}/api_credentials/ | 
[**EmailServicesCancelCreate**](EmailAPI.md#EmailServicesCancelCreate) | **Post** /api/email/services/{id}/cancel/ | 
[**EmailServicesChangeTierPartialUpdate**](EmailAPI.md#EmailServicesChangeTierPartialUpdate) | **Patch** /api/email/services/{id}/change_tier/ | 
[**EmailServicesCreate**](EmailAPI.md#EmailServicesCreate) | **Post** /api/email/services/ | 
[**EmailServicesDedicatedIpCreate**](EmailAPI.md#EmailServicesDedicatedIpCreate) | **Post** /api/email/services/{id}/dedicated_ip/ | 
[**EmailServicesDedicatedIpDestroy**](EmailAPI.md#EmailServicesDedicatedIpDestroy) | **Delete** /api/email/services/{id}/dedicated_ip/ | 
[**EmailServicesDestroy**](EmailAPI.md#EmailServicesDestroy) | **Delete** /api/email/services/{id}/ | 
[**EmailServicesDomainsCreate**](EmailAPI.md#EmailServicesDomainsCreate) | **Post** /api/email/services/{service_pk}/domains/ | 
[**EmailServicesDomainsList**](EmailAPI.md#EmailServicesDomainsList) | **Get** /api/email/services/{service_pk}/domains/ | 
[**EmailServicesList**](EmailAPI.md#EmailServicesList) | **Get** /api/email/services/ | 
[**EmailServicesMessagesRetrieve**](EmailAPI.md#EmailServicesMessagesRetrieve) | **Get** /api/email/services/{service_pk}/messages/ | 
[**EmailServicesPartialUpdate**](EmailAPI.md#EmailServicesPartialUpdate) | **Patch** /api/email/services/{id}/ | 
[**EmailServicesRestoreCreate**](EmailAPI.md#EmailServicesRestoreCreate) | **Post** /api/email/services/{id}/restore/ | 
[**EmailServicesRetrieve**](EmailAPI.md#EmailServicesRetrieve) | **Get** /api/email/services/{id}/ | 
[**EmailServicesSandboxAddressesCreate**](EmailAPI.md#EmailServicesSandboxAddressesCreate) | **Post** /api/email/services/{service_pk}/sandbox_addresses/ | 
[**EmailServicesSandboxAddressesList**](EmailAPI.md#EmailServicesSandboxAddressesList) | **Get** /api/email/services/{service_pk}/sandbox_addresses/ | 
[**EmailServicesSmtpCredentialsCreate**](EmailAPI.md#EmailServicesSmtpCredentialsCreate) | **Post** /api/email/services/{service_pk}/smtp_credentials/ | 
[**EmailServicesSmtpCredentialsList**](EmailAPI.md#EmailServicesSmtpCredentialsList) | **Get** /api/email/services/{service_pk}/smtp_credentials/ | 
[**EmailServicesStatsRetrieve**](EmailAPI.md#EmailServicesStatsRetrieve) | **Get** /api/email/services/{service_pk}/stats/ | 
[**EmailServicesSuppressionsCreate**](EmailAPI.md#EmailServicesSuppressionsCreate) | **Post** /api/email/services/{service_pk}/suppressions/ | 
[**EmailServicesSuppressionsList**](EmailAPI.md#EmailServicesSuppressionsList) | **Get** /api/email/services/{service_pk}/suppressions/ | 
[**EmailSmtpCredentialsCreate**](EmailAPI.md#EmailSmtpCredentialsCreate) | **Post** /api/email/smtp_credentials/ | 
[**EmailSmtpCredentialsDestroy**](EmailAPI.md#EmailSmtpCredentialsDestroy) | **Delete** /api/email/smtp_credentials/{id}/ | 
[**EmailSmtpCredentialsList**](EmailAPI.md#EmailSmtpCredentialsList) | **Get** /api/email/smtp_credentials/ | 
[**EmailSmtpCredentialsRetrieve**](EmailAPI.md#EmailSmtpCredentialsRetrieve) | **Get** /api/email/smtp_credentials/{id}/ | 
[**EmailSuppressionsCreate**](EmailAPI.md#EmailSuppressionsCreate) | **Post** /api/email/suppressions/ | 
[**EmailSuppressionsDestroy**](EmailAPI.md#EmailSuppressionsDestroy) | **Delete** /api/email/suppressions/{id}/ | 
[**EmailSuppressionsList**](EmailAPI.md#EmailSuppressionsList) | **Get** /api/email/suppressions/ | 
[**EmailSuppressionsRetrieve**](EmailAPI.md#EmailSuppressionsRetrieve) | **Get** /api/email/suppressions/{id}/ | 



## EmailApiCredentialsCreate

> ApiCredential EmailApiCredentialsCreate(ctx).ApiCredential(apiCredential).Execute()





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
	apiCredential := *openapiclient.NewApiCredential(int32(123), "Label_example", "KeyPrefix_example", "LastUsedAt_example", false, "CreatedAt_example", "RevokedAt_example") // ApiCredential |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailApiCredentialsCreate(context.Background()).ApiCredential(apiCredential).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailApiCredentialsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailApiCredentialsCreate`: ApiCredential
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailApiCredentialsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailApiCredentialsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiCredential** | [**ApiCredential**](ApiCredential.md) |  | 

### Return type

[**ApiCredential**](ApiCredential.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailApiCredentialsDestroy

> EmailApiCredentialsDestroy(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this api credential.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmailAPI.EmailApiCredentialsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailApiCredentialsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this api credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailApiCredentialsDestroyRequest struct via the builder pattern


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


## EmailApiCredentialsList

> PaginatedApiCredentialList EmailApiCredentialsList(ctx).Page(page).Execute()





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
	resp, r, err := apiClient.EmailAPI.EmailApiCredentialsList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailApiCredentialsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailApiCredentialsList`: PaginatedApiCredentialList
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailApiCredentialsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailApiCredentialsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedApiCredentialList**](PaginatedApiCredentialList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailApiCredentialsRetrieve

> ApiCredential EmailApiCredentialsRetrieve(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this api credential.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailApiCredentialsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailApiCredentialsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailApiCredentialsRetrieve`: ApiCredential
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailApiCredentialsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this api credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailApiCredentialsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiCredential**](ApiCredential.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailDomainsCreate

> SendingDomain EmailDomainsCreate(ctx).DomainAdd(domainAdd).Execute()





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
	domainAdd := *openapiclient.NewDomainAdd("Name_example", openapiclient.DnsSourceEnum("manual")) // DomainAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailDomainsCreate(context.Background()).DomainAdd(domainAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailDomainsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailDomainsCreate`: SendingDomain
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailDomainsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailDomainsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainAdd** | [**DomainAdd**](DomainAdd.md) |  | 

### Return type

[**SendingDomain**](SendingDomain.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailDomainsInboundRoutesCreate

> InboundRoute EmailDomainsInboundRoutesCreate(ctx, domainPk).InboundRoute(inboundRoute).Execute()





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
	domainPk := int32(56) // int32 | 
	inboundRoute := *openapiclient.NewInboundRoute(int32(123), int32(123), "Pattern_example", openapiclient.ModeEnum("webhook"), "CreatedAt_example") // InboundRoute | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailDomainsInboundRoutesCreate(context.Background(), domainPk).InboundRoute(inboundRoute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailDomainsInboundRoutesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailDomainsInboundRoutesCreate`: InboundRoute
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailDomainsInboundRoutesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainPk** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailDomainsInboundRoutesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inboundRoute** | [**InboundRoute**](InboundRoute.md) |  | 

### Return type

[**InboundRoute**](InboundRoute.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailDomainsInboundRoutesList

> PaginatedInboundRouteList EmailDomainsInboundRoutesList(ctx, domainPk).Page(page).Execute()





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
	domainPk := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailDomainsInboundRoutesList(context.Background(), domainPk).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailDomainsInboundRoutesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailDomainsInboundRoutesList`: PaginatedInboundRouteList
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailDomainsInboundRoutesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainPk** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailDomainsInboundRoutesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedInboundRouteList**](PaginatedInboundRouteList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailDomainsList

> PaginatedSendingDomainList EmailDomainsList(ctx).Page(page).Execute()





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
	resp, r, err := apiClient.EmailAPI.EmailDomainsList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailDomainsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailDomainsList`: PaginatedSendingDomainList
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailDomainsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailDomainsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedSendingDomainList**](PaginatedSendingDomainList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailDomainsRetrieve

> SendingDomain EmailDomainsRetrieve(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this sending domain.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailDomainsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailDomainsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailDomainsRetrieve`: SendingDomain
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailDomainsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this sending domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailDomainsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SendingDomain**](SendingDomain.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailDomainsRotateDkimCreate

> SendingDomain EmailDomainsRotateDkimCreate(ctx, id).SendingDomain(sendingDomain).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this sending domain.
	sendingDomain := *openapiclient.NewSendingDomain(int32(123), "Name_example", openapiclient.SendingDomainStatusEnum("pending"), openapiclient.DnsSourceEnum("manual"), false, "DkimSelector_example", "DkimRecord_example", "SpfRecord_example", "DmarcRecord_example", "VerifiedAt_example", "LastCheckAt_example", interface{}(123)) // SendingDomain |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailDomainsRotateDkimCreate(context.Background(), id).SendingDomain(sendingDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailDomainsRotateDkimCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailDomainsRotateDkimCreate`: SendingDomain
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailDomainsRotateDkimCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this sending domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailDomainsRotateDkimCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendingDomain** | [**SendingDomain**](SendingDomain.md) |  | 

### Return type

[**SendingDomain**](SendingDomain.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailDomainsToggleInboundCreate

> SendingDomain EmailDomainsToggleInboundCreate(ctx, id).SendingDomain(sendingDomain).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this sending domain.
	sendingDomain := *openapiclient.NewSendingDomain(int32(123), "Name_example", openapiclient.SendingDomainStatusEnum("pending"), openapiclient.DnsSourceEnum("manual"), false, "DkimSelector_example", "DkimRecord_example", "SpfRecord_example", "DmarcRecord_example", "VerifiedAt_example", "LastCheckAt_example", interface{}(123)) // SendingDomain |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailDomainsToggleInboundCreate(context.Background(), id).SendingDomain(sendingDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailDomainsToggleInboundCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailDomainsToggleInboundCreate`: SendingDomain
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailDomainsToggleInboundCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this sending domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailDomainsToggleInboundCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendingDomain** | [**SendingDomain**](SendingDomain.md) |  | 

### Return type

[**SendingDomain**](SendingDomain.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailDomainsVerifyCreate

> SendingDomain EmailDomainsVerifyCreate(ctx, id).SendingDomain(sendingDomain).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this sending domain.
	sendingDomain := *openapiclient.NewSendingDomain(int32(123), "Name_example", openapiclient.SendingDomainStatusEnum("pending"), openapiclient.DnsSourceEnum("manual"), false, "DkimSelector_example", "DkimRecord_example", "SpfRecord_example", "DmarcRecord_example", "VerifiedAt_example", "LastCheckAt_example", interface{}(123)) // SendingDomain |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailDomainsVerifyCreate(context.Background(), id).SendingDomain(sendingDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailDomainsVerifyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailDomainsVerifyCreate`: SendingDomain
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailDomainsVerifyCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this sending domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailDomainsVerifyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendingDomain** | [**SendingDomain**](SendingDomain.md) |  | 

### Return type

[**SendingDomain**](SendingDomain.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailInboundRoutesCreate

> InboundRoute EmailInboundRoutesCreate(ctx).InboundRoute(inboundRoute).Execute()





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
	inboundRoute := *openapiclient.NewInboundRoute(int32(123), int32(123), "Pattern_example", openapiclient.ModeEnum("webhook"), "CreatedAt_example") // InboundRoute | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailInboundRoutesCreate(context.Background()).InboundRoute(inboundRoute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailInboundRoutesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailInboundRoutesCreate`: InboundRoute
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailInboundRoutesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailInboundRoutesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboundRoute** | [**InboundRoute**](InboundRoute.md) |  | 

### Return type

[**InboundRoute**](InboundRoute.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailInboundRoutesDestroy

> EmailInboundRoutesDestroy(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this inbound route.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmailAPI.EmailInboundRoutesDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailInboundRoutesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this inbound route. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailInboundRoutesDestroyRequest struct via the builder pattern


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


## EmailInboundRoutesList

> PaginatedInboundRouteList EmailInboundRoutesList(ctx).Page(page).Execute()





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
	resp, r, err := apiClient.EmailAPI.EmailInboundRoutesList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailInboundRoutesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailInboundRoutesList`: PaginatedInboundRouteList
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailInboundRoutesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailInboundRoutesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedInboundRouteList**](PaginatedInboundRouteList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailInboundRoutesPartialUpdate

> InboundRoute EmailInboundRoutesPartialUpdate(ctx, id).PatchedInboundRoute(patchedInboundRoute).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this inbound route.
	patchedInboundRoute := *openapiclient.NewPatchedInboundRoute() // PatchedInboundRoute |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailInboundRoutesPartialUpdate(context.Background(), id).PatchedInboundRoute(patchedInboundRoute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailInboundRoutesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailInboundRoutesPartialUpdate`: InboundRoute
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailInboundRoutesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this inbound route. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailInboundRoutesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedInboundRoute** | [**PatchedInboundRoute**](PatchedInboundRoute.md) |  | 

### Return type

[**InboundRoute**](InboundRoute.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailInboundRoutesRetrieve

> InboundRoute EmailInboundRoutesRetrieve(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this inbound route.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailInboundRoutesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailInboundRoutesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailInboundRoutesRetrieve`: InboundRoute
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailInboundRoutesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this inbound route. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailInboundRoutesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InboundRoute**](InboundRoute.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailMessagesRetrieve

> EmailMessagesRetrieve(ctx, messageId).Execute()





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
	messageId := "messageId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmailAPI.EmailMessagesRetrieve(context.Background(), messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailMessagesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailMessagesRetrieveRequest struct via the builder pattern


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


## EmailSandboxAddressesCreate

> SandboxAddress EmailSandboxAddressesCreate(ctx).SandboxAddress(sandboxAddress).Execute()





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
	sandboxAddress := *openapiclient.NewSandboxAddress(int32(123), "Address_example", "VerifiedAt_example", "CreatedAt_example") // SandboxAddress | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailSandboxAddressesCreate(context.Background()).SandboxAddress(sandboxAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailSandboxAddressesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailSandboxAddressesCreate`: SandboxAddress
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailSandboxAddressesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailSandboxAddressesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxAddress** | [**SandboxAddress**](SandboxAddress.md) |  | 

### Return type

[**SandboxAddress**](SandboxAddress.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailSandboxAddressesDestroy

> EmailSandboxAddressesDestroy(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this sandbox verified address.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmailAPI.EmailSandboxAddressesDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailSandboxAddressesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this sandbox verified address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailSandboxAddressesDestroyRequest struct via the builder pattern


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


## EmailSandboxAddressesList

> PaginatedSandboxAddressList EmailSandboxAddressesList(ctx).Page(page).Execute()





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
	resp, r, err := apiClient.EmailAPI.EmailSandboxAddressesList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailSandboxAddressesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailSandboxAddressesList`: PaginatedSandboxAddressList
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailSandboxAddressesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailSandboxAddressesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedSandboxAddressList**](PaginatedSandboxAddressList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailSandboxAddressesRetrieve

> SandboxAddress EmailSandboxAddressesRetrieve(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this sandbox verified address.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailSandboxAddressesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailSandboxAddressesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailSandboxAddressesRetrieve`: SandboxAddress
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailSandboxAddressesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this sandbox verified address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailSandboxAddressesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SandboxAddress**](SandboxAddress.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailSendCreate

> EmailSendCreate(ctx).Execute()



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
	r, err := apiClient.EmailAPI.EmailSendCreate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailSendCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEmailSendCreateRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailServicesApiCredentialsCreate

> ApiCredential EmailServicesApiCredentialsCreate(ctx, servicePk).ApiCredential(apiCredential).Execute()





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
	servicePk := int32(56) // int32 | 
	apiCredential := *openapiclient.NewApiCredential(int32(123), "Label_example", "KeyPrefix_example", "LastUsedAt_example", false, "CreatedAt_example", "RevokedAt_example") // ApiCredential |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailServicesApiCredentialsCreate(context.Background(), servicePk).ApiCredential(apiCredential).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailServicesApiCredentialsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailServicesApiCredentialsCreate`: ApiCredential
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailServicesApiCredentialsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePk** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailServicesApiCredentialsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiCredential** | [**ApiCredential**](ApiCredential.md) |  | 

### Return type

[**ApiCredential**](ApiCredential.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailServicesApiCredentialsList

> PaginatedApiCredentialList EmailServicesApiCredentialsList(ctx, servicePk).Page(page).Execute()





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
	servicePk := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailServicesApiCredentialsList(context.Background(), servicePk).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailServicesApiCredentialsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailServicesApiCredentialsList`: PaginatedApiCredentialList
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailServicesApiCredentialsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePk** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailServicesApiCredentialsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedApiCredentialList**](PaginatedApiCredentialList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailServicesCancelCreate

> EmailService EmailServicesCancelCreate(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this email service.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailServicesCancelCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailServicesCancelCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailServicesCancelCreate`: EmailService
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailServicesCancelCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this email service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailServicesCancelCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailService**](EmailService.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailServicesChangeTierPartialUpdate

> EmailService EmailServicesChangeTierPartialUpdate(ctx, id).PatchedSubscribe(patchedSubscribe).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this email service.
	patchedSubscribe := *openapiclient.NewPatchedSubscribe() // PatchedSubscribe |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailServicesChangeTierPartialUpdate(context.Background(), id).PatchedSubscribe(patchedSubscribe).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailServicesChangeTierPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailServicesChangeTierPartialUpdate`: EmailService
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailServicesChangeTierPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this email service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailServicesChangeTierPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedSubscribe** | [**PatchedSubscribe**](PatchedSubscribe.md) |  | 

### Return type

[**EmailService**](EmailService.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailServicesCreate

> EmailService EmailServicesCreate(ctx).Subscribe(subscribe).Execute()





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
	subscribe := *openapiclient.NewSubscribe(openapiclient.TierEnum("starter")) // Subscribe | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailServicesCreate(context.Background()).Subscribe(subscribe).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailServicesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailServicesCreate`: EmailService
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailServicesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailServicesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscribe** | [**Subscribe**](Subscribe.md) |  | 

### Return type

[**EmailService**](EmailService.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailServicesDedicatedIpCreate

> EmailService EmailServicesDedicatedIpCreate(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this email service.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailServicesDedicatedIpCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailServicesDedicatedIpCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailServicesDedicatedIpCreate`: EmailService
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailServicesDedicatedIpCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this email service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailServicesDedicatedIpCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailService**](EmailService.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailServicesDedicatedIpDestroy

> EmailServicesDedicatedIpDestroy(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this email service.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmailAPI.EmailServicesDedicatedIpDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailServicesDedicatedIpDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this email service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailServicesDedicatedIpDestroyRequest struct via the builder pattern


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


## EmailServicesDestroy

> EmailServicesDestroy(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this email service.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmailAPI.EmailServicesDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailServicesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this email service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailServicesDestroyRequest struct via the builder pattern


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


## EmailServicesDomainsCreate

> SendingDomain EmailServicesDomainsCreate(ctx, servicePk).DomainAdd(domainAdd).Execute()





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
	servicePk := int32(56) // int32 | 
	domainAdd := *openapiclient.NewDomainAdd("Name_example", openapiclient.DnsSourceEnum("manual")) // DomainAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailServicesDomainsCreate(context.Background(), servicePk).DomainAdd(domainAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailServicesDomainsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailServicesDomainsCreate`: SendingDomain
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailServicesDomainsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePk** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailServicesDomainsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domainAdd** | [**DomainAdd**](DomainAdd.md) |  | 

### Return type

[**SendingDomain**](SendingDomain.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailServicesDomainsList

> PaginatedSendingDomainList EmailServicesDomainsList(ctx, servicePk).Page(page).Execute()





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
	servicePk := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailServicesDomainsList(context.Background(), servicePk).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailServicesDomainsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailServicesDomainsList`: PaginatedSendingDomainList
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailServicesDomainsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePk** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailServicesDomainsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedSendingDomainList**](PaginatedSendingDomainList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailServicesList

> PaginatedEmailServiceList EmailServicesList(ctx).Page(page).Execute()





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
	resp, r, err := apiClient.EmailAPI.EmailServicesList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailServicesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailServicesList`: PaginatedEmailServiceList
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailServicesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailServicesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedEmailServiceList**](PaginatedEmailServiceList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailServicesMessagesRetrieve

> EmailServicesMessagesRetrieve(ctx, servicePk).Execute()





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
	servicePk := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmailAPI.EmailServicesMessagesRetrieve(context.Background(), servicePk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailServicesMessagesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePk** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailServicesMessagesRetrieveRequest struct via the builder pattern


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


## EmailServicesPartialUpdate

> EmailService EmailServicesPartialUpdate(ctx, id).PatchedEmailService(patchedEmailService).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this email service.
	patchedEmailService := *openapiclient.NewPatchedEmailService() // PatchedEmailService |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailServicesPartialUpdate(context.Background(), id).PatchedEmailService(patchedEmailService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailServicesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailServicesPartialUpdate`: EmailService
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailServicesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this email service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailServicesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedEmailService** | [**PatchedEmailService**](PatchedEmailService.md) |  | 

### Return type

[**EmailService**](EmailService.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailServicesRestoreCreate

> EmailService EmailServicesRestoreCreate(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this email service.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailServicesRestoreCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailServicesRestoreCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailServicesRestoreCreate`: EmailService
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailServicesRestoreCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this email service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailServicesRestoreCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailService**](EmailService.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailServicesRetrieve

> EmailService EmailServicesRetrieve(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this email service.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailServicesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailServicesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailServicesRetrieve`: EmailService
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailServicesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this email service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailServicesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailService**](EmailService.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailServicesSandboxAddressesCreate

> SandboxAddress EmailServicesSandboxAddressesCreate(ctx, servicePk).SandboxAddress(sandboxAddress).Execute()





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
	servicePk := int32(56) // int32 | 
	sandboxAddress := *openapiclient.NewSandboxAddress(int32(123), "Address_example", "VerifiedAt_example", "CreatedAt_example") // SandboxAddress | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailServicesSandboxAddressesCreate(context.Background(), servicePk).SandboxAddress(sandboxAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailServicesSandboxAddressesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailServicesSandboxAddressesCreate`: SandboxAddress
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailServicesSandboxAddressesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePk** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailServicesSandboxAddressesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sandboxAddress** | [**SandboxAddress**](SandboxAddress.md) |  | 

### Return type

[**SandboxAddress**](SandboxAddress.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailServicesSandboxAddressesList

> PaginatedSandboxAddressList EmailServicesSandboxAddressesList(ctx, servicePk).Page(page).Execute()





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
	servicePk := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailServicesSandboxAddressesList(context.Background(), servicePk).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailServicesSandboxAddressesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailServicesSandboxAddressesList`: PaginatedSandboxAddressList
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailServicesSandboxAddressesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePk** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailServicesSandboxAddressesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedSandboxAddressList**](PaginatedSandboxAddressList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailServicesSmtpCredentialsCreate

> SmtpCredential EmailServicesSmtpCredentialsCreate(ctx, servicePk).SmtpCredential(smtpCredential).Execute()





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
	servicePk := int32(56) // int32 | 
	smtpCredential := *openapiclient.NewSmtpCredential(int32(123), "Label_example", "Username_example", false, "CreatedAt_example", "RevokedAt_example") // SmtpCredential |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailServicesSmtpCredentialsCreate(context.Background(), servicePk).SmtpCredential(smtpCredential).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailServicesSmtpCredentialsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailServicesSmtpCredentialsCreate`: SmtpCredential
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailServicesSmtpCredentialsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePk** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailServicesSmtpCredentialsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smtpCredential** | [**SmtpCredential**](SmtpCredential.md) |  | 

### Return type

[**SmtpCredential**](SmtpCredential.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailServicesSmtpCredentialsList

> PaginatedSmtpCredentialList EmailServicesSmtpCredentialsList(ctx, servicePk).Page(page).Execute()





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
	servicePk := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailServicesSmtpCredentialsList(context.Background(), servicePk).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailServicesSmtpCredentialsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailServicesSmtpCredentialsList`: PaginatedSmtpCredentialList
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailServicesSmtpCredentialsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePk** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailServicesSmtpCredentialsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedSmtpCredentialList**](PaginatedSmtpCredentialList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailServicesStatsRetrieve

> EmailServicesStatsRetrieve(ctx, servicePk).Execute()





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
	servicePk := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmailAPI.EmailServicesStatsRetrieve(context.Background(), servicePk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailServicesStatsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePk** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailServicesStatsRetrieveRequest struct via the builder pattern


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


## EmailServicesSuppressionsCreate

> SuppressionEntry EmailServicesSuppressionsCreate(ctx, servicePk).SuppressionEntry(suppressionEntry).Execute()





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
	servicePk := int32(56) // int32 | 
	suppressionEntry := *openapiclient.NewSuppressionEntry(int32(123), "Address_example", openapiclient.ReasonEnum("hard_bounce"), "Detail_example", "CreatedAt_example") // SuppressionEntry |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailServicesSuppressionsCreate(context.Background(), servicePk).SuppressionEntry(suppressionEntry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailServicesSuppressionsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailServicesSuppressionsCreate`: SuppressionEntry
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailServicesSuppressionsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePk** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailServicesSuppressionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suppressionEntry** | [**SuppressionEntry**](SuppressionEntry.md) |  | 

### Return type

[**SuppressionEntry**](SuppressionEntry.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailServicesSuppressionsList

> PaginatedSuppressionEntryList EmailServicesSuppressionsList(ctx, servicePk).Page(page).Execute()





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
	servicePk := int32(56) // int32 | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailServicesSuppressionsList(context.Background(), servicePk).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailServicesSuppressionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailServicesSuppressionsList`: PaginatedSuppressionEntryList
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailServicesSuppressionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePk** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailServicesSuppressionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedSuppressionEntryList**](PaginatedSuppressionEntryList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailSmtpCredentialsCreate

> SmtpCredential EmailSmtpCredentialsCreate(ctx).SmtpCredential(smtpCredential).Execute()





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
	smtpCredential := *openapiclient.NewSmtpCredential(int32(123), "Label_example", "Username_example", false, "CreatedAt_example", "RevokedAt_example") // SmtpCredential |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailSmtpCredentialsCreate(context.Background()).SmtpCredential(smtpCredential).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailSmtpCredentialsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailSmtpCredentialsCreate`: SmtpCredential
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailSmtpCredentialsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailSmtpCredentialsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smtpCredential** | [**SmtpCredential**](SmtpCredential.md) |  | 

### Return type

[**SmtpCredential**](SmtpCredential.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailSmtpCredentialsDestroy

> EmailSmtpCredentialsDestroy(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this smtp credential.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmailAPI.EmailSmtpCredentialsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailSmtpCredentialsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this smtp credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailSmtpCredentialsDestroyRequest struct via the builder pattern


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


## EmailSmtpCredentialsList

> PaginatedSmtpCredentialList EmailSmtpCredentialsList(ctx).Page(page).Execute()





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
	resp, r, err := apiClient.EmailAPI.EmailSmtpCredentialsList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailSmtpCredentialsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailSmtpCredentialsList`: PaginatedSmtpCredentialList
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailSmtpCredentialsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailSmtpCredentialsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedSmtpCredentialList**](PaginatedSmtpCredentialList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailSmtpCredentialsRetrieve

> SmtpCredential EmailSmtpCredentialsRetrieve(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this smtp credential.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailSmtpCredentialsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailSmtpCredentialsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailSmtpCredentialsRetrieve`: SmtpCredential
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailSmtpCredentialsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this smtp credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailSmtpCredentialsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SmtpCredential**](SmtpCredential.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailSuppressionsCreate

> SuppressionEntry EmailSuppressionsCreate(ctx).SuppressionEntry(suppressionEntry).Execute()





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
	suppressionEntry := *openapiclient.NewSuppressionEntry(int32(123), "Address_example", openapiclient.ReasonEnum("hard_bounce"), "Detail_example", "CreatedAt_example") // SuppressionEntry |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailSuppressionsCreate(context.Background()).SuppressionEntry(suppressionEntry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailSuppressionsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailSuppressionsCreate`: SuppressionEntry
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailSuppressionsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailSuppressionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **suppressionEntry** | [**SuppressionEntry**](SuppressionEntry.md) |  | 

### Return type

[**SuppressionEntry**](SuppressionEntry.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailSuppressionsDestroy

> EmailSuppressionsDestroy(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this suppression entry.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmailAPI.EmailSuppressionsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailSuppressionsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this suppression entry. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailSuppressionsDestroyRequest struct via the builder pattern


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


## EmailSuppressionsList

> PaginatedSuppressionEntryList EmailSuppressionsList(ctx).Page(page).Execute()





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
	resp, r, err := apiClient.EmailAPI.EmailSuppressionsList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailSuppressionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailSuppressionsList`: PaginatedSuppressionEntryList
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailSuppressionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailSuppressionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 

### Return type

[**PaginatedSuppressionEntryList**](PaginatedSuppressionEntryList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailSuppressionsRetrieve

> SuppressionEntry EmailSuppressionsRetrieve(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this suppression entry.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.EmailSuppressionsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.EmailSuppressionsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailSuppressionsRetrieve`: SuppressionEntry
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.EmailSuppressionsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this suppression entry. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailSuppressionsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SuppressionEntry**](SuppressionEntry.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

