# \DomainAntispamPoliciesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDomainPolicy**](DomainAntispamPoliciesApi.md#CreateDomainPolicy) | **Post** /api/v1/add/domain-policy | Create domain policy
[**DeleteDomainPolicy**](DomainAntispamPoliciesApi.md#DeleteDomainPolicy) | **Post** /api/v1/delete/domain-policy | Delete domain policy
[**ListBlacklistDomainPolicy**](DomainAntispamPoliciesApi.md#ListBlacklistDomainPolicy) | **Get** /api/v1/get/policy_bl_domain/{domain} | List blacklist domain policy
[**ListWhitelistDomainPolicy**](DomainAntispamPoliciesApi.md#ListWhitelistDomainPolicy) | **Get** /api/v1/get/policy_wl_domain/{domain} | List whitelist domain policy



## CreateDomainPolicy

> CreateTimeLimitedAlias200Response CreateDomainPolicy(ctx).CreateDomainPolicyRequest(createDomainPolicyRequest).Execute()

Create domain policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createDomainPolicyRequest := *openapiclient.NewCreateDomainPolicyRequest() // CreateDomainPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainAntispamPoliciesApi.CreateDomainPolicy(context.Background()).CreateDomainPolicyRequest(createDomainPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAntispamPoliciesApi.CreateDomainPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDomainPolicy`: CreateTimeLimitedAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainAntispamPoliciesApi.CreateDomainPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDomainPolicyRequest** | [**CreateDomainPolicyRequest**](CreateDomainPolicyRequest.md) |  | 

### Return type

[**CreateTimeLimitedAlias200Response**](CreateTimeLimitedAlias200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDomainPolicy

> CreateTimeLimitedAlias200Response DeleteDomainPolicy(ctx).DeleteDomainPolicyRequest(deleteDomainPolicyRequest).Execute()

Delete domain policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deleteDomainPolicyRequest := *openapiclient.NewDeleteDomainPolicyRequest() // DeleteDomainPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainAntispamPoliciesApi.DeleteDomainPolicy(context.Background()).DeleteDomainPolicyRequest(deleteDomainPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAntispamPoliciesApi.DeleteDomainPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDomainPolicy`: CreateTimeLimitedAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainAntispamPoliciesApi.DeleteDomainPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDomainPolicyRequest** | [**DeleteDomainPolicyRequest**](DeleteDomainPolicyRequest.md) |  | 

### Return type

[**CreateTimeLimitedAlias200Response**](CreateTimeLimitedAlias200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBlacklistDomainPolicy

> ListBlacklistDomainPolicy(ctx, domain).XAPIKey(xAPIKey).Execute()

List blacklist domain policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domain := "domain_example" // string | name of domain
    xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainAntispamPoliciesApi.ListBlacklistDomainPolicy(context.Background(), domain).XAPIKey(xAPIKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAntispamPoliciesApi.ListBlacklistDomainPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | name of domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBlacklistDomainPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** | e.g. api-key-string | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWhitelistDomainPolicy

> ListWhitelistDomainPolicy(ctx, domain).XAPIKey(xAPIKey).Execute()

List whitelist domain policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domain := "domain_example" // string | name of domain
    xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainAntispamPoliciesApi.ListWhitelistDomainPolicy(context.Background(), domain).XAPIKey(xAPIKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAntispamPoliciesApi.ListWhitelistDomainPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | name of domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWhitelistDomainPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKey** | **string** | e.g. api-key-string | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

