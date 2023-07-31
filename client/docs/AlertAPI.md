# \AlertAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAlert**](AlertAPI.md#DeleteAlert) | **Patch** /deepfence/alerts/action/delete | Delete Alerts
[**MaskAlert**](AlertAPI.md#MaskAlert) | **Post** /deepfence/alerts/action/mask | Mask Alerts
[**NotifyAlert**](AlertAPI.md#NotifyAlert) | **Post** /deepfence/alerts/action/notify | Notify Alerts
[**UnmaskAlert**](AlertAPI.md#UnmaskAlert) | **Post** /deepfence/alerts/action/unmask | Unmask Alerts



## DeleteAlert

> DeleteAlert(ctx).ModelAlertsActionRequest(modelAlertsActionRequest).Execute()

Delete Alerts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/deepfence/golang_deepfence_sdk/client"
)

func main() {
    modelAlertsActionRequest := *openapiclient.NewModelAlertsActionRequest([]string{"NodeIds_example"}) // ModelAlertsActionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AlertAPI.DeleteAlert(context.Background()).ModelAlertsActionRequest(modelAlertsActionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.DeleteAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelAlertsActionRequest** | [**ModelAlertsActionRequest**](ModelAlertsActionRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaskAlert

> MaskAlert(ctx).ModelAlertsActionRequest(modelAlertsActionRequest).Execute()

Mask Alerts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/deepfence/golang_deepfence_sdk/client"
)

func main() {
    modelAlertsActionRequest := *openapiclient.NewModelAlertsActionRequest([]string{"NodeIds_example"}) // ModelAlertsActionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AlertAPI.MaskAlert(context.Background()).ModelAlertsActionRequest(modelAlertsActionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.MaskAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMaskAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelAlertsActionRequest** | [**ModelAlertsActionRequest**](ModelAlertsActionRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotifyAlert

> NotifyAlert(ctx).ModelAlertsActionRequest(modelAlertsActionRequest).Execute()

Notify Alerts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/deepfence/golang_deepfence_sdk/client"
)

func main() {
    modelAlertsActionRequest := *openapiclient.NewModelAlertsActionRequest([]string{"NodeIds_example"}) // ModelAlertsActionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AlertAPI.NotifyAlert(context.Background()).ModelAlertsActionRequest(modelAlertsActionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.NotifyAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotifyAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelAlertsActionRequest** | [**ModelAlertsActionRequest**](ModelAlertsActionRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnmaskAlert

> UnmaskAlert(ctx).ModelAlertsActionRequest(modelAlertsActionRequest).Execute()

Unmask Alerts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/deepfence/golang_deepfence_sdk/client"
)

func main() {
    modelAlertsActionRequest := *openapiclient.NewModelAlertsActionRequest([]string{"NodeIds_example"}) // ModelAlertsActionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AlertAPI.UnmaskAlert(context.Background()).ModelAlertsActionRequest(modelAlertsActionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.UnmaskAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnmaskAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelAlertsActionRequest** | [**ModelAlertsActionRequest**](ModelAlertsActionRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

