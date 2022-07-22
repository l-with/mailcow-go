# \DKIMApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDKIMKey**](DKIMApi.md#DeleteDKIMKey) | **Post** /api/v1/delete/dkim | Delete DKIM Key
[**DuplicateDKIMKey**](DKIMApi.md#DuplicateDKIMKey) | **Post** /api/v1/add/dkim_duplicate | Duplicate DKIM Key
[**GenerateDKIMKey**](DKIMApi.md#GenerateDKIMKey) | **Post** /api/v1/add/dkim | Generate DKIM Key
[**GetDKIMKey**](DKIMApi.md#GetDKIMKey) | **Get** /api/v1/get/dkim/{domain} | Get DKIM Key



## DeleteDKIMKey

> []CreateTimeLimitedAlias200Response DeleteDKIMKey(ctx).DeleteDKIMKeyRequest(deleteDKIMKeyRequest).Execute()

Delete DKIM Key



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
    deleteDKIMKeyRequest := *openapiclient.NewDeleteDKIMKeyRequest() // DeleteDKIMKeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DKIMApi.DeleteDKIMKey(context.Background()).DeleteDKIMKeyRequest(deleteDKIMKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DKIMApi.DeleteDKIMKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDKIMKey`: []CreateTimeLimitedAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `DKIMApi.DeleteDKIMKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDKIMKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDKIMKeyRequest** | [**DeleteDKIMKeyRequest**](DeleteDKIMKeyRequest.md) |  | 

### Return type

[**[]CreateTimeLimitedAlias200Response**](CreateTimeLimitedAlias200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DuplicateDKIMKey

> CreateTimeLimitedAlias200Response DuplicateDKIMKey(ctx).DuplicateDKIMKeyRequest(duplicateDKIMKeyRequest).Execute()

Duplicate DKIM Key



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
    duplicateDKIMKeyRequest := *openapiclient.NewDuplicateDKIMKeyRequest() // DuplicateDKIMKeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DKIMApi.DuplicateDKIMKey(context.Background()).DuplicateDKIMKeyRequest(duplicateDKIMKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DKIMApi.DuplicateDKIMKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DuplicateDKIMKey`: CreateTimeLimitedAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `DKIMApi.DuplicateDKIMKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDuplicateDKIMKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **duplicateDKIMKeyRequest** | [**DuplicateDKIMKeyRequest**](DuplicateDKIMKeyRequest.md) |  | 

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


## GenerateDKIMKey

> []map[string]interface{} GenerateDKIMKey(ctx).GenerateDKIMKeyRequest(generateDKIMKeyRequest).Execute()

Generate DKIM Key



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
    generateDKIMKeyRequest := *openapiclient.NewGenerateDKIMKeyRequest() // GenerateDKIMKeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DKIMApi.GenerateDKIMKey(context.Background()).GenerateDKIMKeyRequest(generateDKIMKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DKIMApi.GenerateDKIMKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateDKIMKey`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DKIMApi.GenerateDKIMKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateDKIMKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateDKIMKeyRequest** | [**GenerateDKIMKeyRequest**](GenerateDKIMKeyRequest.md) |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDKIMKey

> GetDKIMKey(ctx, domain).XAPIKey(xAPIKey).Execute()

Get DKIM Key



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
    resp, r, err := apiClient.DKIMApi.GetDKIMKey(context.Background(), domain).XAPIKey(xAPIKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DKIMApi.GetDKIMKey``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetDKIMKeyRequest struct via the builder pattern


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

