# \AgentFilesystemAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableFilesystemTracer**](AgentFilesystemAPI.md#DisableFilesystemTracer) | **Post** /deepfence/filesystem/tracer/disable | Disable filesystem tracer
[**EnableFilesystemTracer**](AgentFilesystemAPI.md#EnableFilesystemTracer) | **Post** /deepfence/filesystem/tracer/enable | Enable filesystem tracer



## DisableFilesystemTracer

> DisableFilesystemTracer(ctx).ModelDisableTracerReq(modelDisableTracerReq).Execute()

Disable filesystem tracer



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
	r, err := apiClient.AgentFilesystemAPI.DisableFilesystemTracer(context.Background()).ModelDisableTracerReq(modelDisableTracerReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentFilesystemAPI.DisableFilesystemTracer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableFilesystemTracerRequest struct via the builder pattern


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


## EnableFilesystemTracer

> EnableFilesystemTracer(ctx).ModelEnableTracerReq(modelEnableTracerReq).Execute()

Enable filesystem tracer



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
	r, err := apiClient.AgentFilesystemAPI.EnableFilesystemTracer(context.Background()).ModelEnableTracerReq(modelEnableTracerReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentFilesystemAPI.EnableFilesystemTracer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableFilesystemTracerRequest struct via the builder pattern


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

