# \AliasesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlias**](AliasesApi.md#CreateAlias) | **Post** /api/v1/add/alias | Create alias
[**CreateTimeLimitedAlias**](AliasesApi.md#CreateTimeLimitedAlias) | **Post** /api/v1/add/time_limited_alias | Create time limited alias
[**DeleteAlias**](AliasesApi.md#DeleteAlias) | **Post** /api/v1/delete/alias | Delete alias
[**GetAliases**](AliasesApi.md#GetAliases) | **Get** /api/v1/get/alias/{id} | Get aliases
[**GetTimeLimitedAliases**](AliasesApi.md#GetTimeLimitedAliases) | **Get** /api/v1/get/time_limited_aliases/{mailbox} | Get time limited aliases
[**UpdateAlias**](AliasesApi.md#UpdateAlias) | **Post** /api/v1/edit/alias | Update alias



## CreateAlias

> []map[string]interface{} CreateAlias(ctx).CreateAliasRequest(createAliasRequest).Execute()

Create alias



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
    createAliasRequest := *openapiclient.NewCreateAliasRequest() // CreateAliasRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AliasesApi.CreateAlias(context.Background()).CreateAliasRequest(createAliasRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasesApi.CreateAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlias`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AliasesApi.CreateAlias`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAliasRequest** | [**CreateAliasRequest**](CreateAliasRequest.md) |  | 

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


## CreateTimeLimitedAlias

> CreateTimeLimitedAlias200Response CreateTimeLimitedAlias(ctx).CreateTimeLimitedAliasRequest(createTimeLimitedAliasRequest).Execute()

Create time limited alias



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
    createTimeLimitedAliasRequest := *openapiclient.NewCreateTimeLimitedAliasRequest() // CreateTimeLimitedAliasRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AliasesApi.CreateTimeLimitedAlias(context.Background()).CreateTimeLimitedAliasRequest(createTimeLimitedAliasRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasesApi.CreateTimeLimitedAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTimeLimitedAlias`: CreateTimeLimitedAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `AliasesApi.CreateTimeLimitedAlias`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTimeLimitedAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTimeLimitedAliasRequest** | [**CreateTimeLimitedAliasRequest**](CreateTimeLimitedAliasRequest.md) |  | 

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


## DeleteAlias

> []map[string]interface{} DeleteAlias(ctx).DeleteAliasRequest(deleteAliasRequest).Execute()

Delete alias



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
    deleteAliasRequest := *openapiclient.NewDeleteAliasRequest() // DeleteAliasRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AliasesApi.DeleteAlias(context.Background()).DeleteAliasRequest(deleteAliasRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasesApi.DeleteAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAlias`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AliasesApi.DeleteAlias`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteAliasRequest** | [**DeleteAliasRequest**](DeleteAliasRequest.md) |  | 

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


## GetAliases

> GetAliases(ctx, id).XAPIKey(xAPIKey).Execute()

Get aliases



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
    resp, r, err := apiClient.AliasesApi.GetAliases(context.Background(), id).XAPIKey(xAPIKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasesApi.GetAliases``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetAliasesRequest struct via the builder pattern


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


## GetTimeLimitedAliases

> GetTimeLimitedAliases(ctx, mailbox).XAPIKey(xAPIKey).Execute()

Get time limited aliases



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
    mailbox := "domain.tld" // string | mailbox you want to get aliasses from
    xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AliasesApi.GetTimeLimitedAliases(context.Background(), mailbox).XAPIKey(xAPIKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasesApi.GetTimeLimitedAliases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailbox** | **string** | mailbox you want to get aliasses from | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeLimitedAliasesRequest struct via the builder pattern


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


## UpdateAlias

> []CreateTimeLimitedAlias200Response UpdateAlias(ctx).UpdateAliasRequest(updateAliasRequest).Execute()

Update alias



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
    updateAliasRequest := *openapiclient.NewUpdateAliasRequest() // UpdateAliasRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AliasesApi.UpdateAlias(context.Background()).UpdateAliasRequest(updateAliasRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasesApi.UpdateAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAlias`: []CreateTimeLimitedAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `AliasesApi.UpdateAlias`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAliasRequest** | [**UpdateAliasRequest**](UpdateAliasRequest.md) |  | 

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

