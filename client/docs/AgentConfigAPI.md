# \AgentConfigAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachAgentFilesystemConfig**](AgentConfigAPI.md#AttachAgentFilesystemConfig) | **Post** /deepfence/configs/agent/filemon/attach | Register Agent Filesystem config
[**AttachAgentNetworkConfig**](AgentConfigAPI.md#AttachAgentNetworkConfig) | **Post** /deepfence/configs/agent/network/attach | Register Agent Network config
[**AttachAgentPolicyConfig**](AgentConfigAPI.md#AttachAgentPolicyConfig) | **Post** /deepfence/configs/agent/policy/attach | Register Agent Policy config
[**AttachAgentQuarantineConfig**](AgentConfigAPI.md#AttachAgentQuarantineConfig) | **Post** /deepfence/configs/agent/quarantine/attach | Register Agent Policy config
[**GetAgentFilesystemConfig**](AgentConfigAPI.md#GetAgentFilesystemConfig) | **Post** /deepfence/configs/agent/filemon/list | Register Agent Filesystem config
[**GetAgentNetworkConfig**](AgentConfigAPI.md#GetAgentNetworkConfig) | **Post** /deepfence/configs/agent/network/list | Register Agent Network config
[**GetAgentPolicyConfig**](AgentConfigAPI.md#GetAgentPolicyConfig) | **Post** /deepfence/configs/agent/policy/list | Register Agent Policy config
[**GetAgentQuarantineConfig**](AgentConfigAPI.md#GetAgentQuarantineConfig) | **Post** /deepfence/configs/agent/quarantine/list | Register Agent Policy config
[**GetNetworkRules**](AgentConfigAPI.md#GetNetworkRules) | **Get** /deepfence/configs/agent/network/rules | Get Network Rules
[**RegisterAgentFilesystemConfig**](AgentConfigAPI.md#RegisterAgentFilesystemConfig) | **Post** /deepfence/configs/agent/filemon/ | Register Agent Filesystem config
[**RegisterAgentNetworkConfig**](AgentConfigAPI.md#RegisterAgentNetworkConfig) | **Post** /deepfence/configs/agent/network/ | Register Agent Network config
[**RegisterAgentPolicyConfig**](AgentConfigAPI.md#RegisterAgentPolicyConfig) | **Post** /deepfence/configs/agent/policy/ | Register Agent Policy config
[**RegisterAgentQuarantineConfig**](AgentConfigAPI.md#RegisterAgentQuarantineConfig) | **Post** /deepfence/configs/agent/quarantine/ | Register Agent Policy config



## AttachAgentFilesystemConfig

> AttachAgentFilesystemConfig(ctx).ModelAttachAgentConfigReq(modelAttachAgentConfigReq).Execute()

Register Agent Filesystem config



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
    modelAttachAgentConfigReq := *openapiclient.NewModelAttachAgentConfigReq([]openapiclient.ModelAgentId{*openapiclient.NewModelAgentId(int32(123), "NodeId_example")}, "ConfigId_example") // ModelAttachAgentConfigReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AgentConfigAPI.AttachAgentFilesystemConfig(context.Background()).ModelAttachAgentConfigReq(modelAttachAgentConfigReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.AttachAgentFilesystemConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachAgentFilesystemConfigRequest struct via the builder pattern


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


## AttachAgentNetworkConfig

> AttachAgentNetworkConfig(ctx).ModelAttachAgentConfigReq(modelAttachAgentConfigReq).Execute()

Register Agent Network config



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
    modelAttachAgentConfigReq := *openapiclient.NewModelAttachAgentConfigReq([]openapiclient.ModelAgentId{*openapiclient.NewModelAgentId(int32(123), "NodeId_example")}, "ConfigId_example") // ModelAttachAgentConfigReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AgentConfigAPI.AttachAgentNetworkConfig(context.Background()).ModelAttachAgentConfigReq(modelAttachAgentConfigReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.AttachAgentNetworkConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachAgentNetworkConfigRequest struct via the builder pattern


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


## AttachAgentPolicyConfig

> AttachAgentPolicyConfig(ctx).ModelAttachAgentConfigReq(modelAttachAgentConfigReq).Execute()

Register Agent Policy config



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
    modelAttachAgentConfigReq := *openapiclient.NewModelAttachAgentConfigReq([]openapiclient.ModelAgentId{*openapiclient.NewModelAgentId(int32(123), "NodeId_example")}, "ConfigId_example") // ModelAttachAgentConfigReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AgentConfigAPI.AttachAgentPolicyConfig(context.Background()).ModelAttachAgentConfigReq(modelAttachAgentConfigReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.AttachAgentPolicyConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachAgentPolicyConfigRequest struct via the builder pattern


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


## AttachAgentQuarantineConfig

> AttachAgentQuarantineConfig(ctx).ModelAttachAgentConfigReq(modelAttachAgentConfigReq).Execute()

Register Agent Policy config



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
    modelAttachAgentConfigReq := *openapiclient.NewModelAttachAgentConfigReq([]openapiclient.ModelAgentId{*openapiclient.NewModelAgentId(int32(123), "NodeId_example")}, "ConfigId_example") // ModelAttachAgentConfigReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AgentConfigAPI.AttachAgentQuarantineConfig(context.Background()).ModelAttachAgentConfigReq(modelAttachAgentConfigReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.AttachAgentQuarantineConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachAgentQuarantineConfigRequest struct via the builder pattern


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


## GetAgentFilesystemConfig

> []ControlsFilesystemTracerConfig GetAgentFilesystemConfig(ctx).ModelGetAgentConfigReq(modelGetAgentConfigReq).Execute()

Register Agent Filesystem config



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
    modelGetAgentConfigReq := *openapiclient.NewModelGetAgentConfigReq([]string{"ConfigIds_example"}) // ModelGetAgentConfigReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentConfigAPI.GetAgentFilesystemConfig(context.Background()).ModelGetAgentConfigReq(modelGetAgentConfigReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.GetAgentFilesystemConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentFilesystemConfig`: []ControlsFilesystemTracerConfig
    fmt.Fprintf(os.Stdout, "Response from `AgentConfigAPI.GetAgentFilesystemConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentFilesystemConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelGetAgentConfigReq** | [**ModelGetAgentConfigReq**](ModelGetAgentConfigReq.md) |  | 

### Return type

[**[]ControlsFilesystemTracerConfig**](ControlsFilesystemTracerConfig.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentNetworkConfig

> []ControlsNetworkTracerConfig GetAgentNetworkConfig(ctx).ModelGetAgentConfigReq(modelGetAgentConfigReq).Execute()

Register Agent Network config



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
    modelGetAgentConfigReq := *openapiclient.NewModelGetAgentConfigReq([]string{"ConfigIds_example"}) // ModelGetAgentConfigReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentConfigAPI.GetAgentNetworkConfig(context.Background()).ModelGetAgentConfigReq(modelGetAgentConfigReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.GetAgentNetworkConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentNetworkConfig`: []ControlsNetworkTracerConfig
    fmt.Fprintf(os.Stdout, "Response from `AgentConfigAPI.GetAgentNetworkConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentNetworkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelGetAgentConfigReq** | [**ModelGetAgentConfigReq**](ModelGetAgentConfigReq.md) |  | 

### Return type

[**[]ControlsNetworkTracerConfig**](ControlsNetworkTracerConfig.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentPolicyConfig

> []ControlsPolicyFilterConfig GetAgentPolicyConfig(ctx).ModelGetAgentConfigReq(modelGetAgentConfigReq).Execute()

Register Agent Policy config



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
    modelGetAgentConfigReq := *openapiclient.NewModelGetAgentConfigReq([]string{"ConfigIds_example"}) // ModelGetAgentConfigReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentConfigAPI.GetAgentPolicyConfig(context.Background()).ModelGetAgentConfigReq(modelGetAgentConfigReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.GetAgentPolicyConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentPolicyConfig`: []ControlsPolicyFilterConfig
    fmt.Fprintf(os.Stdout, "Response from `AgentConfigAPI.GetAgentPolicyConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentPolicyConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelGetAgentConfigReq** | [**ModelGetAgentConfigReq**](ModelGetAgentConfigReq.md) |  | 

### Return type

[**[]ControlsPolicyFilterConfig**](ControlsPolicyFilterConfig.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentQuarantineConfig

> []ControlsQuarantineConfig GetAgentQuarantineConfig(ctx).ModelGetAgentConfigReq(modelGetAgentConfigReq).Execute()

Register Agent Policy config



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
    modelGetAgentConfigReq := *openapiclient.NewModelGetAgentConfigReq([]string{"ConfigIds_example"}) // ModelGetAgentConfigReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentConfigAPI.GetAgentQuarantineConfig(context.Background()).ModelGetAgentConfigReq(modelGetAgentConfigReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.GetAgentQuarantineConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentQuarantineConfig`: []ControlsQuarantineConfig
    fmt.Fprintf(os.Stdout, "Response from `AgentConfigAPI.GetAgentQuarantineConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentQuarantineConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelGetAgentConfigReq** | [**ModelGetAgentConfigReq**](ModelGetAgentConfigReq.md) |  | 

### Return type

[**[]ControlsQuarantineConfig**](ControlsQuarantineConfig.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRules

> map[string]string GetNetworkRules(ctx).Execute()

Get Network Rules



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
    resp, r, err := apiClient.AgentConfigAPI.GetNetworkRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.GetNetworkRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkRules`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `AgentConfigAPI.GetNetworkRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRulesRequest struct via the builder pattern


### Return type

**map[string]string**

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterAgentFilesystemConfig

> RegisterAgentFilesystemConfig(ctx).ControlsFilesystemTracerConfig(controlsFilesystemTracerConfig).Execute()

Register Agent Filesystem config



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
    controlsFilesystemTracerConfig := *openapiclient.NewControlsFilesystemTracerConfig([]openapiclient.ControlsMonitoredFilesConfig{*openapiclient.NewControlsMonitoredFilesConfig([]string{"AccessTypes_example"}, "Path_example", "Wight_example")}, "NodeId_example", []openapiclient.ControlsProcessEventConfig{*openapiclient.NewControlsProcessEventConfig("EventName_example", "Wight_example")}) // ControlsFilesystemTracerConfig |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AgentConfigAPI.RegisterAgentFilesystemConfig(context.Background()).ControlsFilesystemTracerConfig(controlsFilesystemTracerConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.RegisterAgentFilesystemConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterAgentFilesystemConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **controlsFilesystemTracerConfig** | [**ControlsFilesystemTracerConfig**](ControlsFilesystemTracerConfig.md) |  | 

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


## RegisterAgentNetworkConfig

> RegisterAgentNetworkConfig(ctx).ControlsNetworkTracerConfig(controlsNetworkTracerConfig).Execute()

Register Agent Network config



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
    controlsNetworkTracerConfig := *openapiclient.NewControlsNetworkTracerConfig(*openapiclient.NewControlsNetworkRules([]string{"Inbound_example"}, []string{"Outbound_example"}), *openapiclient.NewControlsNetworkRules([]string{"Inbound_example"}, []string{"Outbound_example"}), "Mode_example", "NodeId_example", []string{"ProcessNames_example"}, *openapiclient.NewControlsNetworkRules([]string{"Inbound_example"}, []string{"Outbound_example"})) // ControlsNetworkTracerConfig |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AgentConfigAPI.RegisterAgentNetworkConfig(context.Background()).ControlsNetworkTracerConfig(controlsNetworkTracerConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.RegisterAgentNetworkConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterAgentNetworkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **controlsNetworkTracerConfig** | [**ControlsNetworkTracerConfig**](ControlsNetworkTracerConfig.md) |  | 

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


## RegisterAgentPolicyConfig

> RegisterAgentPolicyConfig(ctx).ControlsPolicyFilterConfig(controlsPolicyFilterConfig).Execute()

Register Agent Policy config



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
    controlsPolicyFilterConfig := *openapiclient.NewControlsPolicyFilterConfig("NodeId_example", []openapiclient.ControlsPolicy{*openapiclient.NewControlsPolicy(int32(123), int32(123), int32(123), "EventType_example", *openapiclient.NewControlsPolicyAlertMatcher(map[string][]string{"key": []string{"Inner_example"}}), "PolicyId_example")}, false) // ControlsPolicyFilterConfig |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AgentConfigAPI.RegisterAgentPolicyConfig(context.Background()).ControlsPolicyFilterConfig(controlsPolicyFilterConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.RegisterAgentPolicyConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterAgentPolicyConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **controlsPolicyFilterConfig** | [**ControlsPolicyFilterConfig**](ControlsPolicyFilterConfig.md) |  | 

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


## RegisterAgentQuarantineConfig

> RegisterAgentQuarantineConfig(ctx).ControlsQuarantineConfig(controlsQuarantineConfig).Execute()

Register Agent Policy config



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
    controlsQuarantineConfig := *openapiclient.NewControlsQuarantineConfig("NodeId_example", []openapiclient.ControlsPolicy{*openapiclient.NewControlsPolicy(int32(123), int32(123), int32(123), "EventType_example", *openapiclient.NewControlsPolicyAlertMatcher(map[string][]string{"key": []string{"Inner_example"}}), "PolicyId_example")}) // ControlsQuarantineConfig |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AgentConfigAPI.RegisterAgentQuarantineConfig(context.Background()).ControlsQuarantineConfig(controlsQuarantineConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.RegisterAgentQuarantineConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterAgentQuarantineConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **controlsQuarantineConfig** | [**ControlsQuarantineConfig**](ControlsQuarantineConfig.md) |  | 

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

