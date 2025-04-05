# \AgentConfigAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentFilesystemConfigAttachedNodes**](AgentConfigAPI.md#AgentFilesystemConfigAttachedNodes) | **Get** /deepfence/configs/agent/filemon/attached-nodes/{config_id} | Get Agent Filesystem config attached nodes
[**AgentNetworkConfigAttachedNodes**](AgentConfigAPI.md#AgentNetworkConfigAttachedNodes) | **Get** /deepfence/configs/agent/network/attached-nodes/{config_id} | Get Agent Network config attached nodes
[**AgentPolicyConfigAttachedNodes**](AgentConfigAPI.md#AgentPolicyConfigAttachedNodes) | **Get** /deepfence/configs/agent/policy/attached-nodes/{config_id} | Get Agent Policy config attached nodes
[**AgentProcessConfigAttachedNodes**](AgentConfigAPI.md#AgentProcessConfigAttachedNodes) | **Get** /deepfence/configs/agent/procmon/attached-nodes/{config_id} | Get Agent Process config attached nodes
[**AgentQuarantineConfigAttachedNodes**](AgentConfigAPI.md#AgentQuarantineConfigAttachedNodes) | **Get** /deepfence/configs/agent/quarantine/attached-nodes/{config_id} | Get Agent Quarantine config attached nodes
[**AttachAgentFilesystemConfig**](AgentConfigAPI.md#AttachAgentFilesystemConfig) | **Post** /deepfence/configs/agent/filemon/attach | Attach Agent Filesystem config
[**AttachAgentNetworkConfig**](AgentConfigAPI.md#AttachAgentNetworkConfig) | **Post** /deepfence/configs/agent/network/attach | Attach Agent Network config
[**AttachAgentPolicyConfig**](AgentConfigAPI.md#AttachAgentPolicyConfig) | **Post** /deepfence/configs/agent/policy/attach | Attach Agent Policy config
[**AttachAgentProcessConfig**](AgentConfigAPI.md#AttachAgentProcessConfig) | **Post** /deepfence/configs/agent/procmon/attach | Attach Agent Process config
[**AttachAgentQuarantineConfig**](AgentConfigAPI.md#AttachAgentQuarantineConfig) | **Post** /deepfence/configs/agent/quarantine/attach | Attach Agent Quarantine config
[**DeleteAgentFilesystemConfig**](AgentConfigAPI.md#DeleteAgentFilesystemConfig) | **Delete** /deepfence/configs/agent/filemon/{config_id} | Delete Agent Filesystem config
[**DeleteAgentNetworkConfig**](AgentConfigAPI.md#DeleteAgentNetworkConfig) | **Delete** /deepfence/configs/agent/network/{config_id} | Delete Agent Network config
[**DeleteAgentPolicyConfig**](AgentConfigAPI.md#DeleteAgentPolicyConfig) | **Delete** /deepfence/configs/agent/policy/{config_id} | Delete Agent Policy config
[**DeleteAgentProcessConfig**](AgentConfigAPI.md#DeleteAgentProcessConfig) | **Delete** /deepfence/configs/agent/procmon/{config_id} | Delete Agent Process config
[**DeleteAgentQuarantineConfig**](AgentConfigAPI.md#DeleteAgentQuarantineConfig) | **Delete** /deepfence/configs/agent/quarantine/{config_id} | Delete Agent Quarantine config
[**ExportAgentPolicyConfig**](AgentConfigAPI.md#ExportAgentPolicyConfig) | **Get** /deepfence/configs/agent/policy/export/{config_id} | Export Agent Policy config
[**ExportAgentQuarantineConfig**](AgentConfigAPI.md#ExportAgentQuarantineConfig) | **Get** /deepfence/configs/agent/quarantine/export/{config_id} | Export Agent Quarantine config
[**GetAgentFilesystemConfig**](AgentConfigAPI.md#GetAgentFilesystemConfig) | **Post** /deepfence/configs/agent/filemon/list | Get Agent Filesystem config
[**GetAgentNetworkConfig**](AgentConfigAPI.md#GetAgentNetworkConfig) | **Post** /deepfence/configs/agent/network/list | Get Agent Network config
[**GetAgentPolicyConfig**](AgentConfigAPI.md#GetAgentPolicyConfig) | **Post** /deepfence/configs/agent/policy/list | Get Agent Policy config
[**GetAgentProcessConfig**](AgentConfigAPI.md#GetAgentProcessConfig) | **Post** /deepfence/configs/agent/procmon/list | Get Agent Process config
[**GetAgentQuarantineConfig**](AgentConfigAPI.md#GetAgentQuarantineConfig) | **Post** /deepfence/configs/agent/quarantine/list | Get Agent Quarantine config
[**GetAgentThreatIntelConfig**](AgentConfigAPI.md#GetAgentThreatIntelConfig) | **Post** /deepfence/configs/agent/threatintel/list | Get Agent threat intel config
[**GetNetworkRules**](AgentConfigAPI.md#GetNetworkRules) | **Get** /deepfence/configs/agent/network/rules | Get Network Rules
[**ImportAgentPolicyConfig**](AgentConfigAPI.md#ImportAgentPolicyConfig) | **Post** /deepfence/configs/agent/policy/import | Import Agent Policy config
[**ImportAgentQuarantineConfig**](AgentConfigAPI.md#ImportAgentQuarantineConfig) | **Post** /deepfence/configs/agent/quarantine/import | Import Agent Quarantine config
[**RegisterAgentFilesystemConfig**](AgentConfigAPI.md#RegisterAgentFilesystemConfig) | **Post** /deepfence/configs/agent/filemon/ | Register Agent Filesystem config
[**RegisterAgentNetworkConfig**](AgentConfigAPI.md#RegisterAgentNetworkConfig) | **Post** /deepfence/configs/agent/network/ | Register Agent Network config
[**RegisterAgentPolicyConfig**](AgentConfigAPI.md#RegisterAgentPolicyConfig) | **Post** /deepfence/configs/agent/policy/ | Register Agent Policy config
[**RegisterAgentProcessConfig**](AgentConfigAPI.md#RegisterAgentProcessConfig) | **Post** /deepfence/configs/agent/procmon/ | Register Agent Process config
[**RegisterAgentQuarantineConfig**](AgentConfigAPI.md#RegisterAgentQuarantineConfig) | **Post** /deepfence/configs/agent/quarantine/ | Register Agent Quarantine config
[**RegisterThreatIntelConfig**](AgentConfigAPI.md#RegisterThreatIntelConfig) | **Post** /deepfence/configs/agent/threatintel/ | Register threat intel config



## AgentFilesystemConfigAttachedNodes

> ModelGetAttachedNodesResp AgentFilesystemConfigAttachedNodes(ctx, configId).Execute()

Get Agent Filesystem config attached nodes



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
	configId := "configId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentConfigAPI.AgentFilesystemConfigAttachedNodes(context.Background(), configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.AgentFilesystemConfigAttachedNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentFilesystemConfigAttachedNodes`: ModelGetAttachedNodesResp
	fmt.Fprintf(os.Stdout, "Response from `AgentConfigAPI.AgentFilesystemConfigAttachedNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentFilesystemConfigAttachedNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelGetAttachedNodesResp**](ModelGetAttachedNodesResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentNetworkConfigAttachedNodes

> ModelGetAttachedNodesResp AgentNetworkConfigAttachedNodes(ctx, configId).Execute()

Get Agent Network config attached nodes



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
	configId := "configId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentConfigAPI.AgentNetworkConfigAttachedNodes(context.Background(), configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.AgentNetworkConfigAttachedNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentNetworkConfigAttachedNodes`: ModelGetAttachedNodesResp
	fmt.Fprintf(os.Stdout, "Response from `AgentConfigAPI.AgentNetworkConfigAttachedNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentNetworkConfigAttachedNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelGetAttachedNodesResp**](ModelGetAttachedNodesResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentPolicyConfigAttachedNodes

> ModelGetAttachedNodesResp AgentPolicyConfigAttachedNodes(ctx, configId).Execute()

Get Agent Policy config attached nodes



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
	configId := "configId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentConfigAPI.AgentPolicyConfigAttachedNodes(context.Background(), configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.AgentPolicyConfigAttachedNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentPolicyConfigAttachedNodes`: ModelGetAttachedNodesResp
	fmt.Fprintf(os.Stdout, "Response from `AgentConfigAPI.AgentPolicyConfigAttachedNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentPolicyConfigAttachedNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelGetAttachedNodesResp**](ModelGetAttachedNodesResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentProcessConfigAttachedNodes

> ModelGetAttachedNodesResp AgentProcessConfigAttachedNodes(ctx, configId).Execute()

Get Agent Process config attached nodes



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
	configId := "configId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentConfigAPI.AgentProcessConfigAttachedNodes(context.Background(), configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.AgentProcessConfigAttachedNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentProcessConfigAttachedNodes`: ModelGetAttachedNodesResp
	fmt.Fprintf(os.Stdout, "Response from `AgentConfigAPI.AgentProcessConfigAttachedNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentProcessConfigAttachedNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelGetAttachedNodesResp**](ModelGetAttachedNodesResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentQuarantineConfigAttachedNodes

> ModelGetAttachedNodesResp AgentQuarantineConfigAttachedNodes(ctx, configId).Execute()

Get Agent Quarantine config attached nodes



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
	configId := "configId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentConfigAPI.AgentQuarantineConfigAttachedNodes(context.Background(), configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.AgentQuarantineConfigAttachedNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentQuarantineConfigAttachedNodes`: ModelGetAttachedNodesResp
	fmt.Fprintf(os.Stdout, "Response from `AgentConfigAPI.AgentQuarantineConfigAttachedNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentQuarantineConfigAttachedNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelGetAttachedNodesResp**](ModelGetAttachedNodesResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachAgentFilesystemConfig

> AttachAgentFilesystemConfig(ctx).ModelAttachAgentConfigReq(modelAttachAgentConfigReq).Execute()

Attach Agent Filesystem config



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
	modelAttachAgentConfigReq := *openapiclient.NewModelAttachAgentConfigReq([]openapiclient.ModelAgentID{*openapiclient.NewModelAgentID(int32(123), "NodeId_example")}, "ConfigId_example") // ModelAttachAgentConfigReq |  (optional)

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

Attach Agent Network config



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
	modelAttachAgentConfigReq := *openapiclient.NewModelAttachAgentConfigReq([]openapiclient.ModelAgentID{*openapiclient.NewModelAgentID(int32(123), "NodeId_example")}, "ConfigId_example") // ModelAttachAgentConfigReq |  (optional)

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

Attach Agent Policy config



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
	modelAttachAgentConfigReq := *openapiclient.NewModelAttachAgentConfigReq([]openapiclient.ModelAgentID{*openapiclient.NewModelAgentID(int32(123), "NodeId_example")}, "ConfigId_example") // ModelAttachAgentConfigReq |  (optional)

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


## AttachAgentProcessConfig

> AttachAgentProcessConfig(ctx).ModelAttachAgentConfigReq(modelAttachAgentConfigReq).Execute()

Attach Agent Process config



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
	modelAttachAgentConfigReq := *openapiclient.NewModelAttachAgentConfigReq([]openapiclient.ModelAgentID{*openapiclient.NewModelAgentID(int32(123), "NodeId_example")}, "ConfigId_example") // ModelAttachAgentConfigReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentConfigAPI.AttachAgentProcessConfig(context.Background()).ModelAttachAgentConfigReq(modelAttachAgentConfigReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.AttachAgentProcessConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachAgentProcessConfigRequest struct via the builder pattern


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

Attach Agent Quarantine config



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
	modelAttachAgentConfigReq := *openapiclient.NewModelAttachAgentConfigReq([]openapiclient.ModelAgentID{*openapiclient.NewModelAgentID(int32(123), "NodeId_example")}, "ConfigId_example") // ModelAttachAgentConfigReq |  (optional)

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


## DeleteAgentFilesystemConfig

> DeleteAgentFilesystemConfig(ctx, configId).Execute()

Delete Agent Filesystem config



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
	configId := "configId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentConfigAPI.DeleteAgentFilesystemConfig(context.Background(), configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.DeleteAgentFilesystemConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentFilesystemConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgentNetworkConfig

> DeleteAgentNetworkConfig(ctx, configId).Execute()

Delete Agent Network config



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
	configId := "configId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentConfigAPI.DeleteAgentNetworkConfig(context.Background(), configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.DeleteAgentNetworkConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentNetworkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgentPolicyConfig

> DeleteAgentPolicyConfig(ctx, configId).Execute()

Delete Agent Policy config



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
	configId := "configId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentConfigAPI.DeleteAgentPolicyConfig(context.Background(), configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.DeleteAgentPolicyConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentPolicyConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgentProcessConfig

> DeleteAgentProcessConfig(ctx, configId).Execute()

Delete Agent Process config



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
	configId := "configId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentConfigAPI.DeleteAgentProcessConfig(context.Background(), configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.DeleteAgentProcessConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentProcessConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgentQuarantineConfig

> DeleteAgentQuarantineConfig(ctx, configId).Execute()

Delete Agent Quarantine config



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
	configId := "configId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentConfigAPI.DeleteAgentQuarantineConfig(context.Background(), configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.DeleteAgentQuarantineConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentQuarantineConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAgentPolicyConfig

> ControlsPolicyFilterConfig ExportAgentPolicyConfig(ctx, configId).Execute()

Export Agent Policy config



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
	configId := "configId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentConfigAPI.ExportAgentPolicyConfig(context.Background(), configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.ExportAgentPolicyConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAgentPolicyConfig`: ControlsPolicyFilterConfig
	fmt.Fprintf(os.Stdout, "Response from `AgentConfigAPI.ExportAgentPolicyConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportAgentPolicyConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ControlsPolicyFilterConfig**](ControlsPolicyFilterConfig.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAgentQuarantineConfig

> ControlsQuarantineConfig ExportAgentQuarantineConfig(ctx, configId).Execute()

Export Agent Quarantine config



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
	configId := "configId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentConfigAPI.ExportAgentQuarantineConfig(context.Background(), configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.ExportAgentQuarantineConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAgentQuarantineConfig`: ControlsQuarantineConfig
	fmt.Fprintf(os.Stdout, "Response from `AgentConfigAPI.ExportAgentQuarantineConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportAgentQuarantineConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ControlsQuarantineConfig**](ControlsQuarantineConfig.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentFilesystemConfig

> []ControlsFilesystemTracerConfig GetAgentFilesystemConfig(ctx).ModelGetAgentConfigReq(modelGetAgentConfigReq).Execute()

Get Agent Filesystem config



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

Get Agent Network config



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

Get Agent Policy config



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


## GetAgentProcessConfig

> []ControlsProcessTracerConfig GetAgentProcessConfig(ctx).ModelGetAgentConfigReq(modelGetAgentConfigReq).Execute()

Get Agent Process config



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
	resp, r, err := apiClient.AgentConfigAPI.GetAgentProcessConfig(context.Background()).ModelGetAgentConfigReq(modelGetAgentConfigReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.GetAgentProcessConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentProcessConfig`: []ControlsProcessTracerConfig
	fmt.Fprintf(os.Stdout, "Response from `AgentConfigAPI.GetAgentProcessConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentProcessConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelGetAgentConfigReq** | [**ModelGetAgentConfigReq**](ModelGetAgentConfigReq.md) |  | 

### Return type

[**[]ControlsProcessTracerConfig**](ControlsProcessTracerConfig.md)

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

Get Agent Quarantine config



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


## GetAgentThreatIntelConfig

> []ControlsThreatIntelInfo GetAgentThreatIntelConfig(ctx).ModelGetAgentConfigReq(modelGetAgentConfigReq).Execute()

Get Agent threat intel config



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
	resp, r, err := apiClient.AgentConfigAPI.GetAgentThreatIntelConfig(context.Background()).ModelGetAgentConfigReq(modelGetAgentConfigReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.GetAgentThreatIntelConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentThreatIntelConfig`: []ControlsThreatIntelInfo
	fmt.Fprintf(os.Stdout, "Response from `AgentConfigAPI.GetAgentThreatIntelConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentThreatIntelConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelGetAgentConfigReq** | [**ModelGetAgentConfigReq**](ModelGetAgentConfigReq.md) |  | 

### Return type

[**[]ControlsThreatIntelInfo**](ControlsThreatIntelInfo.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRules

> ThreatintelRulesWithDirection GetNetworkRules(ctx).Execute()

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
	// response from `GetNetworkRules`: ThreatintelRulesWithDirection
	fmt.Fprintf(os.Stdout, "Response from `AgentConfigAPI.GetNetworkRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRulesRequest struct via the builder pattern


### Return type

[**ThreatintelRulesWithDirection**](ThreatintelRulesWithDirection.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportAgentPolicyConfig

> ModelMessageResponse ImportAgentPolicyConfig(ctx).ConfigId(configId).NetworkPolicyJson(networkPolicyJson).Execute()

Import Agent Policy config



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
	configId := "configId_example" // string | 
	networkPolicyJson := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentConfigAPI.ImportAgentPolicyConfig(context.Background()).ConfigId(configId).NetworkPolicyJson(networkPolicyJson).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.ImportAgentPolicyConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportAgentPolicyConfig`: ModelMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentConfigAPI.ImportAgentPolicyConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportAgentPolicyConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configId** | **string** |  | 
 **networkPolicyJson** | ***os.File** |  | 

### Return type

[**ModelMessageResponse**](ModelMessageResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportAgentQuarantineConfig

> ModelMessageResponse ImportAgentQuarantineConfig(ctx).ConfigId(configId).QuarantinePolicyJson(quarantinePolicyJson).Execute()

Import Agent Quarantine config



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
	configId := "configId_example" // string | 
	quarantinePolicyJson := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentConfigAPI.ImportAgentQuarantineConfig(context.Background()).ConfigId(configId).QuarantinePolicyJson(quarantinePolicyJson).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.ImportAgentQuarantineConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportAgentQuarantineConfig`: ModelMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentConfigAPI.ImportAgentQuarantineConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportAgentQuarantineConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configId** | **string** |  | 
 **quarantinePolicyJson** | ***os.File** |  | 

### Return type

[**ModelMessageResponse**](ModelMessageResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: multipart/form-data
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
	controlsFilesystemTracerConfig := *openapiclient.NewControlsFilesystemTracerConfig("NodeId_example", int32(123), []openapiclient.ControlsMonitoredFilesConfig{*openapiclient.NewControlsMonitoredFilesConfig([]string{"Accesstypes_example"}, []string{"Exclude_example"}, false, "Root_example", "Severity_example")}) // ControlsFilesystemTracerConfig |  (optional)

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
	controlsNetworkTracerConfig := *openapiclient.NewControlsNetworkTracerConfig(*openapiclient.NewControlsNetworkRules([]string{"Inbound_example"}, []string{"Outbound_example"}), *openapiclient.NewControlsNetworkRules([]string{"Inbound_example"}, []string{"Outbound_example"}), "Mode_example", "NodeId_example", []string{"ProcessNames_example"}, *openapiclient.NewControlsNetworkRules([]string{"Inbound_example"}, []string{"Outbound_example"}), int32(123)) // ControlsNetworkTracerConfig |  (optional)

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
	controlsPolicyFilterConfig := *openapiclient.NewControlsPolicyFilterConfig([]string{"IgnoredRemoteIps_example"}, "NodeId_example", []openapiclient.ControlsNetworkPolicy{*openapiclient.NewControlsNetworkPolicy("Action_example", int32(123), int32(123), int32(123), *openapiclient.NewControlsPolicyAlertMatcher(map[string][]string{"key": []string{"Inner_example"}}), "PolicyId_example", int64(123))}, int64(123), false) // ControlsPolicyFilterConfig |  (optional)

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


## RegisterAgentProcessConfig

> RegisterAgentProcessConfig(ctx).ControlsProcessTracerConfig(controlsProcessTracerConfig).Execute()

Register Agent Process config



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
	controlsProcessTracerConfig := *openapiclient.NewControlsProcessTracerConfig("NodeId_example", int32(123)) // ControlsProcessTracerConfig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentConfigAPI.RegisterAgentProcessConfig(context.Background()).ControlsProcessTracerConfig(controlsProcessTracerConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.RegisterAgentProcessConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterAgentProcessConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **controlsProcessTracerConfig** | [**ControlsProcessTracerConfig**](ControlsProcessTracerConfig.md) |  | 

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

Register Agent Quarantine config



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
	controlsQuarantineConfig := *openapiclient.NewControlsQuarantineConfig("NodeId_example", []openapiclient.ControlsRuncPolicy{*openapiclient.NewControlsRuncPolicy("Action_example", int32(123), int32(123), *openapiclient.NewControlsPolicyAlertMatcher(map[string][]string{"key": []string{"Inner_example"}}), "NodeType_example", "PolicyId_example", int32(123))}, int32(123)) // ControlsQuarantineConfig |  (optional)

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


## RegisterThreatIntelConfig

> RegisterThreatIntelConfig(ctx).ControlsThreatIntelInfo(controlsThreatIntelInfo).Execute()

Register threat intel config



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
	controlsThreatIntelInfo := *openapiclient.NewControlsThreatIntelInfo("CloudPostureControlsHash_example", "CloudPostureControlsUrl_example", []string{"IgnoredAlertRuleIds_example"}, []string{"InternalIps_example"}, "MalwareScannerRulesHash_example", "MalwareScannerRulesUrl_example", "NetworkAlertRulesUrl_example", "RulesHash_example", "SecretScannerRulesHash_example", "SecretScannerRulesUrl_example", int32(123)) // ControlsThreatIntelInfo |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentConfigAPI.RegisterThreatIntelConfig(context.Background()).ControlsThreatIntelInfo(controlsThreatIntelInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentConfigAPI.RegisterThreatIntelConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterThreatIntelConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **controlsThreatIntelInfo** | [**ControlsThreatIntelInfo**](ControlsThreatIntelInfo.md) |  | 

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

