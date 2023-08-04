# \CloudResourcesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IngestCloudResources**](CloudResourcesAPI.md#IngestCloudResources) | **Post** /deepfence/ingest/cloud-resources | Ingest Cloud resources
[**IngestViolations**](CloudResourcesAPI.md#IngestViolations) | **Post** /deepfence/ingest/violations | Ingest Violations
[**IngestWAFRules**](CloudResourcesAPI.md#IngestWAFRules) | **Post** /deepfence/ingest/waf-rules | Ingest WAF Rules



## IngestCloudResources

> IngestCloudResources(ctx).IngestersCloudResource(ingestersCloudResource).Execute()

Ingest Cloud resources



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
    ingestersCloudResource := []openapiclient.IngestersCloudResource{*openapiclient.NewIngestersCloudResource()} // []IngestersCloudResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CloudResourcesAPI.IngestCloudResources(context.Background()).IngestersCloudResource(ingestersCloudResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudResourcesAPI.IngestCloudResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestCloudResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestersCloudResource** | [**[]IngestersCloudResource**](IngestersCloudResource.md) |  | 

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


## IngestViolations

> IngestViolations(ctx).IngestersNetworkViolation(ingestersNetworkViolation).Execute()

Ingest Violations



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
    ingestersNetworkViolation := []openapiclient.IngestersNetworkViolation{*openapiclient.NewIngestersNetworkViolation()} // []IngestersNetworkViolation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CloudResourcesAPI.IngestViolations(context.Background()).IngestersNetworkViolation(ingestersNetworkViolation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudResourcesAPI.IngestViolations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestViolationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestersNetworkViolation** | [**[]IngestersNetworkViolation**](IngestersNetworkViolation.md) |  | 

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


## IngestWAFRules

> IngestWAFRules(ctx).IngestersWAFRule(ingestersWAFRule).Execute()

Ingest WAF Rules



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
    ingestersWAFRule := []openapiclient.IngestersWAFRule{*openapiclient.NewIngestersWAFRule()} // []IngestersWAFRule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CloudResourcesAPI.IngestWAFRules(context.Background()).IngestersWAFRule(ingestersWAFRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudResourcesAPI.IngestWAFRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestWAFRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestersWAFRule** | [**[]IngestersWAFRule**](IngestersWAFRule.md) |  | 

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

