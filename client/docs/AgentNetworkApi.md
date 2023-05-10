# \AgentNetworkApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableNetworkTracer**](AgentNetworkApi.md#DisableNetworkTracer) | **Post** /deepfence/network/tracer/disable | Disable network tracer
[**EnableNetworkTracer**](AgentNetworkApi.md#EnableNetworkTracer) | **Post** /deepfence/network/tracer/enable | Enable network tracer



## DisableNetworkTracer

> DisableNetworkTracer(ctx).ModelDisableNetworkTracerReq(modelDisableNetworkTracerReq).Execute()

Disable network tracer



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
    modelDisableNetworkTracerReq := *openapiclient.NewModelDisableNetworkTracerReq([]openapiclient.ModelAgentId{*openapiclient.NewModelAgentId(int32(123), "NodeId_example")}) // ModelDisableNetworkTracerReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AgentNetworkApi.DisableNetworkTracer(context.Background()).ModelDisableNetworkTracerReq(modelDisableNetworkTracerReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentNetworkApi.DisableNetworkTracer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableNetworkTracerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelDisableNetworkTracerReq** | [**ModelDisableNetworkTracerReq**](ModelDisableNetworkTracerReq.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableNetworkTracer

> EnableNetworkTracer(ctx).ModelEnableNetworkTracerReq(modelEnableNetworkTracerReq).Execute()

Enable network tracer



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
    modelEnableNetworkTracerReq := *openapiclient.NewModelEnableNetworkTracerReq([]openapiclient.ModelAgentId{*openapiclient.NewModelAgentId(int32(123), "NodeId_example")}, *openapiclient.NewControlsNetworkRules([]string{"Inbound_example"}, []string{"Outbound_example"}), *openapiclient.NewControlsNetworkRules([]string{"Inbound_example"}, []string{"Outbound_example"}), *openapiclient.NewControlsNetworkRules([]string{"Inbound_example"}, []string{"Outbound_example"})) // ModelEnableNetworkTracerReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AgentNetworkApi.EnableNetworkTracer(context.Background()).ModelEnableNetworkTracerReq(modelEnableNetworkTracerReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentNetworkApi.EnableNetworkTracer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableNetworkTracerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelEnableNetworkTracerReq** | [**ModelEnableNetworkTracerReq**](ModelEnableNetworkTracerReq.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

