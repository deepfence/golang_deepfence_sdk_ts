# \AgentNetworkAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableCloudNetworkTracer**](AgentNetworkAPI.md#DisableCloudNetworkTracer) | **Post** /deepfence/network/cloud-tracer/disable | Disable cloud network tracer
[**DisableNetworkTracer**](AgentNetworkAPI.md#DisableNetworkTracer) | **Post** /deepfence/network/tracer/disable | Disable network tracer
[**EnableCloudNetworkTracer**](AgentNetworkAPI.md#EnableCloudNetworkTracer) | **Post** /deepfence/network/cloud-tracer/enable | Enable cloud network tracer
[**EnableNetworkTracer**](AgentNetworkAPI.md#EnableNetworkTracer) | **Post** /deepfence/network/tracer/enable | Enable network tracer



## DisableCloudNetworkTracer

> DisableCloudNetworkTracer(ctx).ModelDisableCloudTracerReq(modelDisableCloudTracerReq).Execute()

Disable cloud network tracer



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
	modelDisableCloudTracerReq := *openapiclient.NewModelDisableCloudTracerReq([]openapiclient.ModelAgentID{*openapiclient.NewModelAgentID(int32(123), "NodeId_example")}) // ModelDisableCloudTracerReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentNetworkAPI.DisableCloudNetworkTracer(context.Background()).ModelDisableCloudTracerReq(modelDisableCloudTracerReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentNetworkAPI.DisableCloudNetworkTracer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableCloudNetworkTracerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelDisableCloudTracerReq** | [**ModelDisableCloudTracerReq**](ModelDisableCloudTracerReq.md) |  | 

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


## DisableNetworkTracer

> DisableNetworkTracer(ctx).ModelDisableTracerReq(modelDisableTracerReq).Execute()

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
	modelDisableTracerReq := *openapiclient.NewModelDisableTracerReq([]openapiclient.ModelAgentID{*openapiclient.NewModelAgentID(int32(123), "NodeId_example")}) // ModelDisableTracerReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentNetworkAPI.DisableNetworkTracer(context.Background()).ModelDisableTracerReq(modelDisableTracerReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentNetworkAPI.DisableNetworkTracer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableNetworkTracerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelDisableTracerReq** | [**ModelDisableTracerReq**](ModelDisableTracerReq.md) |  | 

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


## EnableCloudNetworkTracer

> EnableCloudNetworkTracer(ctx).ModelEnableCloudTracerReq(modelEnableCloudTracerReq).Execute()

Enable cloud network tracer



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
	modelEnableCloudTracerReq := *openapiclient.NewModelEnableCloudTracerReq([]openapiclient.ModelAgentID{*openapiclient.NewModelAgentID(int32(123), "NodeId_example")}, []string{"AwsS3Bucket_example"}) // ModelEnableCloudTracerReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentNetworkAPI.EnableCloudNetworkTracer(context.Background()).ModelEnableCloudTracerReq(modelEnableCloudTracerReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentNetworkAPI.EnableCloudNetworkTracer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableCloudNetworkTracerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelEnableCloudTracerReq** | [**ModelEnableCloudTracerReq**](ModelEnableCloudTracerReq.md) |  | 

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


## EnableNetworkTracer

> EnableNetworkTracer(ctx).ModelEnableTracerReq(modelEnableTracerReq).Execute()

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
	modelEnableTracerReq := *openapiclient.NewModelEnableTracerReq([]openapiclient.ModelAgentID{*openapiclient.NewModelAgentID(int32(123), "NodeId_example")}) // ModelEnableTracerReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentNetworkAPI.EnableNetworkTracer(context.Background()).ModelEnableTracerReq(modelEnableTracerReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentNetworkAPI.EnableNetworkTracer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableNetworkTracerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelEnableTracerReq** | [**ModelEnableTracerReq**](ModelEnableTracerReq.md) |  | 

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

