# \AppPasswordsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppPassword**](AppPasswordsApi.md#CreateAppPassword) | **Post** /api/v1/add/app-passwd | Create App Password
[**DeleteAppPassword**](AppPasswordsApi.md#DeleteAppPassword) | **Post** /api/v1/delete/app-passwd | Delete App Password
[**GetAppPassword**](AppPasswordsApi.md#GetAppPassword) | **Get** /api/v1/get/app-passwd/all/{mailbox} | Get App Password



## CreateAppPassword

> CreateAlias200ResponseInner CreateAppPassword(ctx).CreateAppPasswordRequest(createAppPasswordRequest).Execute()

Create App Password



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
    createAppPasswordRequest := *openapiclient.NewCreateAppPasswordRequest() // CreateAppPasswordRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppPasswordsApi.CreateAppPassword(context.Background()).CreateAppPasswordRequest(createAppPasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppPasswordsApi.CreateAppPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAppPassword`: CreateAlias200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `AppPasswordsApi.CreateAppPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAppPasswordRequest** | [**CreateAppPasswordRequest**](CreateAppPasswordRequest.md) |  | 

### Return type

[**CreateAlias200ResponseInner**](CreateAlias200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppPassword

> CreateAlias200ResponseInner DeleteAppPassword(ctx).DeleteAppPasswordRequest(deleteAppPasswordRequest).Execute()

Delete App Password



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
    deleteAppPasswordRequest := *openapiclient.NewDeleteAppPasswordRequest() // DeleteAppPasswordRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppPasswordsApi.DeleteAppPassword(context.Background()).DeleteAppPasswordRequest(deleteAppPasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppPasswordsApi.DeleteAppPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAppPassword`: CreateAlias200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `AppPasswordsApi.DeleteAppPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteAppPasswordRequest** | [**DeleteAppPasswordRequest**](DeleteAppPasswordRequest.md) |  | 

### Return type

[**CreateAlias200ResponseInner**](CreateAlias200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppPassword

> GetAppPassword(ctx, mailbox).XAPIKey(xAPIKey).Execute()

Get App Password



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
    mailbox := "hello@mailcow.email" // string | mailbox of entry you want to get
    xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppPasswordsApi.GetAppPassword(context.Background(), mailbox).XAPIKey(xAPIKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppPasswordsApi.GetAppPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailbox** | **string** | mailbox of entry you want to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppPasswordRequest struct via the builder pattern


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

