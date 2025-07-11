# \LookupAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAPIEndpoints**](LookupAPI.md#GetAPIEndpoints) | **Post** /deepfence/lookup/api-endpoints | Retrieve API endpoint data
[**GetCloudCompliances**](LookupAPI.md#GetCloudCompliances) | **Post** /deepfence/lookup/cloud-compliances | Retrieve Cloud Compliances data
[**GetCloudResources**](LookupAPI.md#GetCloudResources) | **Post** /deepfence/lookup/cloud-resources | Get Cloud Resources
[**GetComplianceControls**](LookupAPI.md#GetComplianceControls) | **Post** /deepfence/lookup/compliance-controls | Retrieve Cloud Compliances Control data
[**GetCompliances**](LookupAPI.md#GetCompliances) | **Post** /deepfence/lookup/compliances | Retrieve Compliances data
[**GetContainerImages**](LookupAPI.md#GetContainerImages) | **Post** /deepfence/lookup/containerimages | Retrieve Container Images data
[**GetContainers**](LookupAPI.md#GetContainers) | **Post** /deepfence/lookup/containers | Retrieve Containers data
[**GetExploitableAlerts**](LookupAPI.md#GetExploitableAlerts) | **Post** /deepfence/lookup/exploitable-alerts | Get Exploitable Alerts
[**GetFileAlerts**](LookupAPI.md#GetFileAlerts) | **Post** /deepfence/lookup/file-alerts | Get File Alerts
[**GetFilesystemAlertRules**](LookupAPI.md#GetFilesystemAlertRules) | **Post** /deepfence/lookup/file-alert-rules | Get File Alert Rules
[**GetHosts**](LookupAPI.md#GetHosts) | **Post** /deepfence/lookup/hosts | Retrieve Hosts data
[**GetKubernetesClusters**](LookupAPI.md#GetKubernetesClusters) | **Post** /deepfence/lookup/kubernetesclusters | Retrieve K8s data
[**GetMaliciousConnectionAlerts**](LookupAPI.md#GetMaliciousConnectionAlerts) | **Post** /deepfence/lookup/malicious-connection-alerts | Get Malicious Connection Alerts
[**GetMalwareRules**](LookupAPI.md#GetMalwareRules) | **Post** /deepfence/lookup/malware-rules | Get Malware Rules
[**GetMalwares**](LookupAPI.md#GetMalwares) | **Post** /deepfence/lookup/malwares | Retrieve Malwares data
[**GetNetworkAlertRules**](LookupAPI.md#GetNetworkAlertRules) | **Post** /deepfence/lookup/network-alert-rules | Get Network Alert Rules
[**GetNetworkAlerts**](LookupAPI.md#GetNetworkAlerts) | **Post** /deepfence/lookup/network-alerts | Get Network Alerts
[**GetNetworkViolations**](LookupAPI.md#GetNetworkViolations) | **Post** /deepfence/lookup/network-violations | Get Network Violations
[**GetPods**](LookupAPI.md#GetPods) | **Post** /deepfence/lookup/pods | Retrieve Pods data
[**GetProcessAlertRules**](LookupAPI.md#GetProcessAlertRules) | **Post** /deepfence/lookup/process-alert-rules | Get Process Alert Rules
[**GetProcessAlerts**](LookupAPI.md#GetProcessAlerts) | **Post** /deepfence/lookup/process-alerts | Get Process Alerts
[**GetProcesses**](LookupAPI.md#GetProcesses) | **Post** /deepfence/lookup/processes | Retrieve Processes data
[**GetQuarantineViolations**](LookupAPI.md#GetQuarantineViolations) | **Post** /deepfence/lookup/quarantine-violations | Get Quarantine Violations
[**GetRegistryAccount**](LookupAPI.md#GetRegistryAccount) | **Post** /deepfence/lookup/registryaccount | Get Images in Registry
[**GetSecretRules**](LookupAPI.md#GetSecretRules) | **Post** /deepfence/lookup/secret-rules | Get Secret Rules
[**GetSecrets**](LookupAPI.md#GetSecrets) | **Post** /deepfence/lookup/secrets | Retrieve Secrets data
[**GetVulnerabilities**](LookupAPI.md#GetVulnerabilities) | **Post** /deepfence/lookup/vulnerabilities | Retrieve Vulnerabilities data
[**GetVulnerabilityRules**](LookupAPI.md#GetVulnerabilityRules) | **Post** /deepfence/lookup/vulnerability-rules | Get Vulnerability Rules



## GetAPIEndpoints

> []ModelAPIEndpoint GetAPIEndpoints(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve API endpoint data



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetAPIEndpoints(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetAPIEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAPIEndpoints`: []ModelAPIEndpoint
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetAPIEndpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAPIEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelAPIEndpoint**](ModelAPIEndpoint.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudCompliances

> []ModelCloudCompliance GetCloudCompliances(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve Cloud Compliances data



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetCloudCompliances(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetCloudCompliances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudCompliances`: []ModelCloudCompliance
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetCloudCompliances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudCompliancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelCloudCompliance**](ModelCloudCompliance.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudResources

> []ModelCloudResource GetCloudResources(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Get Cloud Resources



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetCloudResources(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetCloudResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudResources`: []ModelCloudResource
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetCloudResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelCloudResource**](ModelCloudResource.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComplianceControls

> []ModelCloudComplianceControl GetComplianceControls(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve Cloud Compliances Control data



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetComplianceControls(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetComplianceControls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComplianceControls`: []ModelCloudComplianceControl
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetComplianceControls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetComplianceControlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelCloudComplianceControl**](ModelCloudComplianceControl.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompliances

> []ModelCompliance GetCompliances(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve Compliances data



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetCompliances(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetCompliances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompliances`: []ModelCompliance
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetCompliances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompliancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelCompliance**](ModelCompliance.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerImages

> []ModelContainerImage GetContainerImages(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve Container Images data



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetContainerImages(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetContainerImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerImages`: []ModelContainerImage
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetContainerImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelContainerImage**](ModelContainerImage.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainers

> []ModelContainer GetContainers(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve Containers data



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetContainers(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetContainers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainers`: []ModelContainer
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelContainer**](ModelContainer.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExploitableAlerts

> []ModelExploitableAlert GetExploitableAlerts(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Get Exploitable Alerts



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetExploitableAlerts(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetExploitableAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExploitableAlerts`: []ModelExploitableAlert
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetExploitableAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExploitableAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelExploitableAlert**](ModelExploitableAlert.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileAlerts

> []ModelFileAlert GetFileAlerts(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Get File Alerts



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetFileAlerts(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetFileAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileAlerts`: []ModelFileAlert
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetFileAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFileAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelFileAlert**](ModelFileAlert.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilesystemAlertRules

> []ModelFilesystemAlertRule GetFilesystemAlertRules(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Get File Alert Rules



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetFilesystemAlertRules(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetFilesystemAlertRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilesystemAlertRules`: []ModelFilesystemAlertRule
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetFilesystemAlertRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesystemAlertRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelFilesystemAlertRule**](ModelFilesystemAlertRule.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHosts

> []ModelHost GetHosts(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve Hosts data



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetHosts(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHosts`: []ModelHost
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelHost**](ModelHost.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesClusters

> []ModelKubernetesCluster GetKubernetesClusters(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve K8s data



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetKubernetesClusters(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetKubernetesClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesClusters`: []ModelKubernetesCluster
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetKubernetesClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelKubernetesCluster**](ModelKubernetesCluster.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaliciousConnectionAlerts

> []ModelMaliciousConnectionAlert GetMaliciousConnectionAlerts(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Get Malicious Connection Alerts



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetMaliciousConnectionAlerts(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetMaliciousConnectionAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMaliciousConnectionAlerts`: []ModelMaliciousConnectionAlert
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetMaliciousConnectionAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMaliciousConnectionAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelMaliciousConnectionAlert**](ModelMaliciousConnectionAlert.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMalwareRules

> []ModelMalwareRule GetMalwareRules(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Get Malware Rules



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetMalwareRules(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetMalwareRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMalwareRules`: []ModelMalwareRule
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetMalwareRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMalwareRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelMalwareRule**](ModelMalwareRule.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMalwares

> []ModelMalware GetMalwares(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve Malwares data



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetMalwares(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetMalwares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMalwares`: []ModelMalware
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetMalwares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMalwaresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelMalware**](ModelMalware.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAlertRules

> []ModelNetworkAlertRule GetNetworkAlertRules(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Get Network Alert Rules



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetNetworkAlertRules(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetNetworkAlertRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkAlertRules`: []ModelNetworkAlertRule
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetNetworkAlertRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAlertRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelNetworkAlertRule**](ModelNetworkAlertRule.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAlerts

> []ModelNetworkAlert GetNetworkAlerts(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Get Network Alerts



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetNetworkAlerts(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetNetworkAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkAlerts`: []ModelNetworkAlert
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetNetworkAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelNetworkAlert**](ModelNetworkAlert.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkViolations

> []ModelNetworkViolation GetNetworkViolations(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Get Network Violations



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetNetworkViolations(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetNetworkViolations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkViolations`: []ModelNetworkViolation
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetNetworkViolations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkViolationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelNetworkViolation**](ModelNetworkViolation.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPods

> []ModelPod GetPods(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve Pods data



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetPods(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetPods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPods`: []ModelPod
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetPods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelPod**](ModelPod.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessAlertRules

> []ModelProcessAlertRule GetProcessAlertRules(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Get Process Alert Rules



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetProcessAlertRules(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetProcessAlertRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessAlertRules`: []ModelProcessAlertRule
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetProcessAlertRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessAlertRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelProcessAlertRule**](ModelProcessAlertRule.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcessAlerts

> []ModelProcessAlert GetProcessAlerts(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Get Process Alerts



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetProcessAlerts(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetProcessAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcessAlerts`: []ModelProcessAlert
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetProcessAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelProcessAlert**](ModelProcessAlert.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcesses

> []ModelProcess GetProcesses(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve Processes data



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetProcesses(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetProcesses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcesses`: []ModelProcess
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetProcesses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelProcess**](ModelProcess.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuarantineViolations

> []ModelQuarantineViolation GetQuarantineViolations(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Get Quarantine Violations



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetQuarantineViolations(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetQuarantineViolations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuarantineViolations`: []ModelQuarantineViolation
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetQuarantineViolations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuarantineViolationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelQuarantineViolation**](ModelQuarantineViolation.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistryAccount

> []ModelRegistryAccount GetRegistryAccount(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Get Images in Registry



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetRegistryAccount(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetRegistryAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegistryAccount`: []ModelRegistryAccount
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetRegistryAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistryAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelRegistryAccount**](ModelRegistryAccount.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretRules

> []ModelSecretRule GetSecretRules(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Get Secret Rules



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetSecretRules(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetSecretRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecretRules`: []ModelSecretRule
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetSecretRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelSecretRule**](ModelSecretRule.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecrets

> []ModelSecret GetSecrets(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve Secrets data



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetSecrets(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecrets`: []ModelSecret
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetSecrets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelSecret**](ModelSecret.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVulnerabilities

> []ModelVulnerability GetVulnerabilities(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve Vulnerabilities data



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetVulnerabilities(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetVulnerabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVulnerabilities`: []ModelVulnerability
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetVulnerabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVulnerabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelVulnerability**](ModelVulnerability.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVulnerabilityRules

> []ModelVulnerabilityRule GetVulnerabilityRules(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Get Vulnerability Rules



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
	lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LookupAPI.GetVulnerabilityRules(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetVulnerabilityRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVulnerabilityRules`: []ModelVulnerabilityRule
	fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetVulnerabilityRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVulnerabilityRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelVulnerabilityRule**](ModelVulnerabilityRule.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

