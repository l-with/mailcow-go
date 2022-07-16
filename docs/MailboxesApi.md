# \MailboxesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMailbox**](MailboxesApi.md#CreateMailbox) | **Post** /api/v1/add/mailbox | Create mailbox
[**DeleteMailbox**](MailboxesApi.md#DeleteMailbox) | **Post** /api/v1/delete/mailbox | Delete mailbox
[**DeleteMailboxTags**](MailboxesApi.md#DeleteMailboxTags) | **Post** /api/v1/delete/mailbox/tag/{mailbox} | Delete mailbox tags
[**EditMailboxSpamFilterScore**](MailboxesApi.md#EditMailboxSpamFilterScore) | **Post** /api/v1/edit/spam-score/ | Edit mailbox spam filter score
[**GetMailboxes**](MailboxesApi.md#GetMailboxes) | **Get** /api/v1/get/mailbox/{id} | Get mailboxes
[**QuarantineNotifications**](MailboxesApi.md#QuarantineNotifications) | **Post** /api/v1/edit/quarantine_notification | Quarantine Notifications
[**UpdateMailbox**](MailboxesApi.md#UpdateMailbox) | **Post** /api/v1/edit/mailbox | Update mailbox
[**UpdateMailboxACL**](MailboxesApi.md#UpdateMailboxACL) | **Post** /api/v1/edit/user-acl | Update mailbox ACL
[**UpdatePushoverSettings**](MailboxesApi.md#UpdatePushoverSettings) | **Post** /api/v1/edit/pushover | Update Pushover settings



## CreateMailbox

> []CreateAlias200Response CreateMailbox(ctx).CreateMailboxRequest(createMailboxRequest).Execute()

Create mailbox



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
    createMailboxRequest := *openapiclient.NewCreateMailboxRequest() // CreateMailboxRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailboxesApi.CreateMailbox(context.Background()).CreateMailboxRequest(createMailboxRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailboxesApi.CreateMailbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMailbox`: []CreateAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `MailboxesApi.CreateMailbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMailboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMailboxRequest** | [**CreateMailboxRequest**](CreateMailboxRequest.md) |  | 

### Return type

[**[]CreateAlias200Response**](CreateAlias200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMailbox

> []map[string]interface{} DeleteMailbox(ctx).DeleteMailboxRequest(deleteMailboxRequest).Execute()

Delete mailbox



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
    deleteMailboxRequest := *openapiclient.NewDeleteMailboxRequest() // DeleteMailboxRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailboxesApi.DeleteMailbox(context.Background()).DeleteMailboxRequest(deleteMailboxRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailboxesApi.DeleteMailbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMailbox`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MailboxesApi.DeleteMailbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMailboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteMailboxRequest** | [**DeleteMailboxRequest**](DeleteMailboxRequest.md) |  | 

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


## DeleteMailboxTags

> CreateAlias200Response DeleteMailboxTags(ctx, mailbox).DeleteMailboxTagsRequest(deleteMailboxTagsRequest).Execute()

Delete mailbox tags



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
    mailbox := "info@domain.tld" // string | name of mailbox
    deleteMailboxTagsRequest := *openapiclient.NewDeleteMailboxTagsRequest() // DeleteMailboxTagsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailboxesApi.DeleteMailboxTags(context.Background(), mailbox).DeleteMailboxTagsRequest(deleteMailboxTagsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailboxesApi.DeleteMailboxTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMailboxTags`: CreateAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `MailboxesApi.DeleteMailboxTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailbox** | **string** | name of mailbox | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMailboxTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteMailboxTagsRequest** | [**DeleteMailboxTagsRequest**](DeleteMailboxTagsRequest.md) |  | 

### Return type

[**CreateAlias200Response**](CreateAlias200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditMailboxSpamFilterScore

> CreateAlias200Response EditMailboxSpamFilterScore(ctx).Body(body).Execute()

Edit mailbox spam filter score



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
    body := interface{}(987) // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailboxesApi.EditMailboxSpamFilterScore(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailboxesApi.EditMailboxSpamFilterScore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditMailboxSpamFilterScore`: CreateAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `MailboxesApi.EditMailboxSpamFilterScore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditMailboxSpamFilterScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **interface{}** |  | 

### Return type

[**CreateAlias200Response**](CreateAlias200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMailboxes

> GetMailboxes(ctx, id).Tags(tags).XAPIKey(xAPIKey).Execute()

Get mailboxes



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
    tags := "tag1,tag2" // string | comma seperated list of tags to filter by (optional)
    xAPIKey := "api-key-string" // string | e.g. api-key-string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailboxesApi.GetMailboxes(context.Background(), id).Tags(tags).XAPIKey(xAPIKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailboxesApi.GetMailboxes``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetMailboxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tags** | **string** | comma seperated list of tags to filter by | 
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


## QuarantineNotifications

> QuarantineNotifications(ctx).QuarantineNotificationsRequest(quarantineNotificationsRequest).Execute()

Quarantine Notifications



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
    quarantineNotificationsRequest := *openapiclient.NewQuarantineNotificationsRequest() // QuarantineNotificationsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailboxesApi.QuarantineNotifications(context.Background()).QuarantineNotificationsRequest(quarantineNotificationsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailboxesApi.QuarantineNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuarantineNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quarantineNotificationsRequest** | [**QuarantineNotificationsRequest**](QuarantineNotificationsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMailbox

> []CreateAlias200Response UpdateMailbox(ctx).UpdateMailboxRequest(updateMailboxRequest).Execute()

Update mailbox



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
    updateMailboxRequest := *openapiclient.NewUpdateMailboxRequest() // UpdateMailboxRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailboxesApi.UpdateMailbox(context.Background()).UpdateMailboxRequest(updateMailboxRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailboxesApi.UpdateMailbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMailbox`: []CreateAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `MailboxesApi.UpdateMailbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMailboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateMailboxRequest** | [**UpdateMailboxRequest**](UpdateMailboxRequest.md) |  | 

### Return type

[**[]CreateAlias200Response**](CreateAlias200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMailboxACL

> CreateAlias200Response UpdateMailboxACL(ctx).UpdateMailboxACLRequest(updateMailboxACLRequest).Execute()

Update mailbox ACL



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
    updateMailboxACLRequest := *openapiclient.NewUpdateMailboxACLRequest() // UpdateMailboxACLRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailboxesApi.UpdateMailboxACL(context.Background()).UpdateMailboxACLRequest(updateMailboxACLRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailboxesApi.UpdateMailboxACL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMailboxACL`: CreateAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `MailboxesApi.UpdateMailboxACL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMailboxACLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateMailboxACLRequest** | [**UpdateMailboxACLRequest**](UpdateMailboxACLRequest.md) |  | 

### Return type

[**CreateAlias200Response**](CreateAlias200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePushoverSettings

> CreateAlias200Response UpdatePushoverSettings(ctx).UpdatePushoverSettingsRequest(updatePushoverSettingsRequest).Execute()

Update Pushover settings



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
    updatePushoverSettingsRequest := *openapiclient.NewUpdatePushoverSettingsRequest() // UpdatePushoverSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailboxesApi.UpdatePushoverSettings(context.Background()).UpdatePushoverSettingsRequest(updatePushoverSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailboxesApi.UpdatePushoverSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePushoverSettings`: CreateAlias200Response
    fmt.Fprintf(os.Stdout, "Response from `MailboxesApi.UpdatePushoverSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePushoverSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatePushoverSettingsRequest** | [**UpdatePushoverSettingsRequest**](UpdatePushoverSettingsRequest.md) |  | 

### Return type

[**CreateAlias200Response**](CreateAlias200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

