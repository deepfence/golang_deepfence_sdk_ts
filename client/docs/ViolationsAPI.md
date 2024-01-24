# \ViolationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNetworkViolations**](ViolationsAPI.md#DeleteNetworkViolations) | **Post** /deepfence/remove/network-violations | Remove network violations
[**DeleteQuarantineViolations**](ViolationsAPI.md#DeleteQuarantineViolations) | **Post** /deepfence/remove/quarantine-violations | Remove quarantine violations



## DeleteNetworkViolations

> DeleteNetworkViolations(ctx).ModelDeleteFilter(modelDeleteFilter).Execute()

Remove network violations



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
	modelDeleteFilter := *openapiclient.NewModelDeleteFilter([]string{"NodeIds_example"}) // ModelDeleteFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ViolationsAPI.DeleteNetworkViolations(context.Background()).ModelDeleteFilter(modelDeleteFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViolationsAPI.DeleteNetworkViolations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkViolationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelDeleteFilter** | [**ModelDeleteFilter**](ModelDeleteFilter.md) |  | 

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


## DeleteQuarantineViolations

> DeleteQuarantineViolations(ctx).ModelDeleteFilter(modelDeleteFilter).Execute()

Remove quarantine violations



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
	modelDeleteFilter := *openapiclient.NewModelDeleteFilter([]string{"NodeIds_example"}) // ModelDeleteFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ViolationsAPI.DeleteQuarantineViolations(context.Background()).ModelDeleteFilter(modelDeleteFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViolationsAPI.DeleteQuarantineViolations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQuarantineViolationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelDeleteFilter** | [**ModelDeleteFilter**](ModelDeleteFilter.md) |  | 

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

