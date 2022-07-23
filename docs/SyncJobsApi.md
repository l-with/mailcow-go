# \SyncJobsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSyncJob**](SyncJobsApi.md#CreateSyncJob) | **Post** /api/v1/add/syncjob | Create sync job
[**DeleteSyncJob**](SyncJobsApi.md#DeleteSyncJob) | **Post** /api/v1/delete/syncjob | Delete sync job
[**GetSyncJobs**](SyncJobsApi.md#GetSyncJobs) | **Get** /api/v1/get/syncjobs/{id} | Get sync jobs
[**UpdateSyncJob**](SyncJobsApi.md#UpdateSyncJob) | **Post** /api/v1/edit/syncjob | Update sync job



## CreateSyncJob

> []map[string]interface{} CreateSyncJob(ctx).CreateSyncJobRequest(createSyncJobRequest).Execute()

Create sync job



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
    createSyncJobRequest := *openapiclient.NewCreateSyncJobRequest() // CreateSyncJobRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyncJobsApi.CreateSyncJob(context.Background()).CreateSyncJobRequest(createSyncJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyncJobsApi.CreateSyncJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSyncJob`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SyncJobsApi.CreateSyncJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSyncJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSyncJobRequest** | [**CreateSyncJobRequest**](CreateSyncJobRequest.md) |  | 

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


## DeleteSyncJob

> []map[string]interface{} DeleteSyncJob(ctx).DeleteSyncJobRequest(deleteSyncJobRequest).Execute()

Delete sync job



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
    deleteSyncJobRequest := *openapiclient.NewDeleteSyncJobRequest() // DeleteSyncJobRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyncJobsApi.DeleteSyncJob(context.Background()).DeleteSyncJobRequest(deleteSyncJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyncJobsApi.DeleteSyncJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSyncJob`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SyncJobsApi.DeleteSyncJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSyncJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSyncJobRequest** | [**DeleteSyncJobRequest**](DeleteSyncJobRequest.md) |  | 

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


## GetSyncJobs

> GetSyncJobs(ctx, id).Execute()

Get sync jobs



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
    id := "id_example" // string | id of entry you want to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyncJobsApi.GetSyncJobs(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyncJobsApi.GetSyncJobs``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetSyncJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UpdateSyncJob

> []CreateTimeLimitedAlias200Response UpdateSyncJob(ctx).UpdateSyncJobRequest(updateSyncJobRequest).Execute()

Update sync job



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
    updateSyncJobRequest := *openapiclient.NewUpdateSyncJobRequest() // UpdateSyncJobRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyncJobsApi.UpdateSyncJob(context.Background()).UpdateSyncJobRequest(updateSyncJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyncJobsApi.UpdateSyncJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSyncJob`: []CreateTimeLimitedAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `SyncJobsApi.UpdateSyncJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSyncJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSyncJobRequest** | [**UpdateSyncJobRequest**](UpdateSyncJobRequest.md) |  | 

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

