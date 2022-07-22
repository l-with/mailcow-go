# \OutgoingTLSPolicyMapOverridesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTLSPolicyMap**](OutgoingTLSPolicyMapOverridesApi.md#CreateTLSPolicyMap) | **Post** /api/v1/add/tls-policy-map | Create TLS Policy Map
[**DeleteTLSPolicyMap**](OutgoingTLSPolicyMapOverridesApi.md#DeleteTLSPolicyMap) | **Post** /api/v1/delete/tls-policy-map | Delete TLS Policy Map
[**GetTLSPolicyMap**](OutgoingTLSPolicyMapOverridesApi.md#GetTLSPolicyMap) | **Get** /api/v1/get/tls-policy-map/{id} | Get TLS Policy Map



## CreateTLSPolicyMap

> CreateTimeLimitedAlias200Response CreateTLSPolicyMap(ctx).CreateTLSPolicyMapRequest(createTLSPolicyMapRequest).Execute()

Create TLS Policy Map



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
    createTLSPolicyMapRequest := *openapiclient.NewCreateTLSPolicyMapRequest() // CreateTLSPolicyMapRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OutgoingTLSPolicyMapOverridesApi.CreateTLSPolicyMap(context.Background()).CreateTLSPolicyMapRequest(createTLSPolicyMapRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutgoingTLSPolicyMapOverridesApi.CreateTLSPolicyMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTLSPolicyMap`: CreateTimeLimitedAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `OutgoingTLSPolicyMapOverridesApi.CreateTLSPolicyMap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTLSPolicyMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTLSPolicyMapRequest** | [**CreateTLSPolicyMapRequest**](CreateTLSPolicyMapRequest.md) |  | 

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


## DeleteTLSPolicyMap

> CreateTimeLimitedAlias200Response DeleteTLSPolicyMap(ctx).DeleteTLSPolicyMapRequest(deleteTLSPolicyMapRequest).Execute()

Delete TLS Policy Map



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
    deleteTLSPolicyMapRequest := *openapiclient.NewDeleteTLSPolicyMapRequest() // DeleteTLSPolicyMapRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OutgoingTLSPolicyMapOverridesApi.DeleteTLSPolicyMap(context.Background()).DeleteTLSPolicyMapRequest(deleteTLSPolicyMapRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutgoingTLSPolicyMapOverridesApi.DeleteTLSPolicyMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTLSPolicyMap`: CreateTimeLimitedAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `OutgoingTLSPolicyMapOverridesApi.DeleteTLSPolicyMap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTLSPolicyMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteTLSPolicyMapRequest** | [**DeleteTLSPolicyMapRequest**](DeleteTLSPolicyMapRequest.md) |  | 

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


## GetTLSPolicyMap

> GetTLSPolicyMap(ctx, id).XAPIKey(xAPIKey).Execute()

Get TLS Policy Map



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
    id := "all" // string | id of entry you want to get
    xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OutgoingTLSPolicyMapOverridesApi.GetTLSPolicyMap(context.Background(), id).XAPIKey(xAPIKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutgoingTLSPolicyMapOverridesApi.GetTLSPolicyMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of entry you want to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTLSPolicyMapRequest struct via the builder pattern


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

