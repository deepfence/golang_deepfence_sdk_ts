# \AlertAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DailyAlertCounts**](AlertAPI.md#DailyAlertCounts) | **Get** /deepfence/statistics/alerts-daily-counts | Get Alerts daily counts
[**DeleteAlert**](AlertAPI.md#DeleteAlert) | **Patch** /deepfence/alerts/action/delete | Delete Alerts
[**GetMitreAttackMatrix**](AlertAPI.md#GetMitreAttackMatrix) | **Post** /deepfence/alerts/mitre-attack-matrix | Get Mitre Attack Matrix
[**MaskAlert**](AlertAPI.md#MaskAlert) | **Post** /deepfence/alerts/action/mask | Mask Alerts
[**NotifyAlert**](AlertAPI.md#NotifyAlert) | **Post** /deepfence/alerts/action/notify | Notify Alerts
[**UnmaskAlert**](AlertAPI.md#UnmaskAlert) | **Post** /deepfence/alerts/action/unmask | Unmask Alerts



## DailyAlertCounts

> ReportersDailySevCounts DailyAlertCounts(ctx).Execute()

Get Alerts daily counts



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertAPI.DailyAlertCounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.DailyAlertCounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DailyAlertCounts`: ReportersDailySevCounts
    fmt.Fprintf(os.Stdout, "Response from `AlertAPI.DailyAlertCounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDailyAlertCountsRequest struct via the builder pattern


### Return type

[**ReportersDailySevCounts**](ReportersDailySevCounts.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## GetMitreAttackMatrix

> ModelMitreAttackMatrix GetMitreAttackMatrix(ctx).ModelMitreAttackMatrixRequest(modelMitreAttackMatrixRequest).Execute()

Get Mitre Attack Matrix



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
    modelMitreAttackMatrixRequest := *openapiclient.NewModelMitreAttackMatrixRequest() // ModelMitreAttackMatrixRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertAPI.GetMitreAttackMatrix(context.Background()).ModelMitreAttackMatrixRequest(modelMitreAttackMatrixRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.GetMitreAttackMatrix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMitreAttackMatrix`: ModelMitreAttackMatrix
    fmt.Fprintf(os.Stdout, "Response from `AlertAPI.GetMitreAttackMatrix`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMitreAttackMatrixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelMitreAttackMatrixRequest** | [**ModelMitreAttackMatrixRequest**](ModelMitreAttackMatrixRequest.md) |  | 

### Return type

[**ModelMitreAttackMatrix**](ModelMitreAttackMatrix.md)

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

