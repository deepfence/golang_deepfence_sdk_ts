# \AgentConfigAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachAgentConfig**](AgentConfigAPI.md#AttachAgentConfig) | **Post** /configs/agent/attach | Register Agent config
[**GetAgentConfig**](AgentConfigAPI.md#GetAgentConfig) | **Get** /configs/agent/ | Register Agent config
[**RegisterAgentConfig**](AgentConfigAPI.md#RegisterAgentConfig) | **Post** /configs/agent/ | Register Agent config



## AttachAgentConfig

> AttachAgentConfig(ctx).ModelAttachAgentConfigReq(modelAttachAgentConfigReq).Execute()

Register Agent config



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
    modelAttachAgentConfigReq := *openapiclient.NewModelAttachAgentConfigReq([]openapiclient.ModelAgentId{*openapiclient.NewModelAgentId(int32(123), "NodeId_example")}, "Name_example") // ModelAttachAgentConfigReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AgentConfigAPI.AttachAgentConfig(context.Background()).ModelAttachAgentConfigReq(modelAttachAgentConfigReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.AttachAgentConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachAgentConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelAttachAgentConfigReq** | [**ModelAttachAgentConfigReq**](ModelAttachAgentConfigReq.md) |  | 

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


## GetAgentConfig

> []ControlsAgentConfig GetAgentConfig(ctx).Execute()

Register Agent config



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
    resp, r, err := apiClient.AgentConfigAPI.GetAgentConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.GetAgentConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentConfig`: []ControlsAgentConfig
    fmt.Fprintf(os.Stdout, "Response from `AgentConfigAPI.GetAgentConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentConfigRequest struct via the builder pattern


### Return type

[**[]ControlsAgentConfig**](ControlsAgentConfig.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterAgentConfig

> RegisterAgentConfig(ctx).ControlsAgentConfig(controlsAgentConfig).Execute()

Register Agent config



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
    controlsAgentConfig := *openapiclient.NewControlsAgentConfig(*openapiclient.NewControlsFilesystemTracerConfig("Hash_example", []openapiclient.ControlsMonitoredFilesConfig{*openapiclient.NewControlsMonitoredFilesConfig([]string{"AccessTypes_example"}, "Path_example", "Wight_example")}, []openapiclient.ControlsProcessEventConfig{*openapiclient.NewControlsProcessEventConfig("EventName_example", "Wight_example")}), "Name_example", *openapiclient.NewControlsNetworkTracerConfig("Hash_example", *openapiclient.NewControlsNetworkRules([]string{"Inbound_example"}, []string{"Outbound_example"}), *openapiclient.NewControlsNetworkRules([]string{"Inbound_example"}, []string{"Outbound_example"}), *openapiclient.NewControlsNetworkRules([]string{"Inbound_example"}, []string{"Outbound_example"})), *openapiclient.NewControlsPolicyFilterConfig("EventType_example", "Hash_example", map[string]interface{}(123))) // ControlsAgentConfig |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AgentConfigAPI.RegisterAgentConfig(context.Background()).ControlsAgentConfig(controlsAgentConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.RegisterAgentConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterAgentConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **controlsAgentConfig** | [**ControlsAgentConfig**](ControlsAgentConfig.md) |  | 

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

