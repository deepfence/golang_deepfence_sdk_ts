# \AgentProcessAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableProcessTracer**](AgentProcessAPI.md#DisableProcessTracer) | **Post** /deepfence/process/tracer/disable | Disable process tracer
[**EnableProcessTracer**](AgentProcessAPI.md#EnableProcessTracer) | **Post** /deepfence/process/tracer/enable | Enable process tracer



## DisableProcessTracer

> DisableProcessTracer(ctx).ModelDisableTracerReq(modelDisableTracerReq).Execute()

Disable process tracer



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
	r, err := apiClient.AgentProcessAPI.DisableProcessTracer(context.Background()).ModelDisableTracerReq(modelDisableTracerReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentProcessAPI.DisableProcessTracer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableProcessTracerRequest struct via the builder pattern


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


## EnableProcessTracer

> EnableProcessTracer(ctx).ModelEnableTracerReq(modelEnableTracerReq).Execute()

Enable process tracer



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
	r, err := apiClient.AgentProcessAPI.EnableProcessTracer(context.Background()).ModelEnableTracerReq(modelEnableTracerReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentProcessAPI.EnableProcessTracer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableProcessTracerRequest struct via the builder pattern


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

