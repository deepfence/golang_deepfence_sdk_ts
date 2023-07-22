# \AgentFilesystemAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnableFilesystemTracer**](AgentFilesystemAPI.md#EnableFilesystemTracer) | **Post** /deepfence/filesystem/tracer/enable | Enable filesystem tracer



## EnableFilesystemTracer

> EnableFilesystemTracer(ctx).ModelEnableFilesystemTracerReq(modelEnableFilesystemTracerReq).Execute()

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
    modelEnableFilesystemTracerReq := *openapiclient.NewModelEnableFilesystemTracerReq([]openapiclient.ModelAgentId{*openapiclient.NewModelAgentId(int32(123), "NodeId_example")}, "Path_example") // ModelEnableFilesystemTracerReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AgentFilesystemAPI.EnableFilesystemTracer(context.Background()).ModelEnableFilesystemTracerReq(modelEnableFilesystemTracerReq).Execute()
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
 **modelEnableFilesystemTracerReq** | [**ModelEnableFilesystemTracerReq**](ModelEnableFilesystemTracerReq.md) |  | 

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

