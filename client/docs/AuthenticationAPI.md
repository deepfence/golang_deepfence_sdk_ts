# \AuthenticationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthToken**](AuthenticationAPI.md#AuthToken) | **Post** /deepfence/auth/token | Get Access Token for API Token
[**AuthTokenRefresh**](AuthenticationAPI.md#AuthTokenRefresh) | **Post** /deepfence/auth/token/refresh | Refresh access token
[**CreateSSOProvider**](AuthenticationAPI.md#CreateSSOProvider) | **Post** /deepfence/single-sign-on | Configure Single sign-on
[**DeleteSSOProvider**](AuthenticationAPI.md#DeleteSSOProvider) | **Delete** /deepfence/single-sign-on/{id} | Delete Single sign-on configuration
[**GetSSOProviders**](AuthenticationAPI.md#GetSSOProviders) | **Get** /deepfence/single-sign-on | Get Single sign-on configurations
[**Login**](AuthenticationAPI.md#Login) | **Post** /deepfence/user/login | Login API
[**Logout**](AuthenticationAPI.md#Logout) | **Post** /deepfence/user/logout | Logout API
[**SsoInitiateLogin**](AuthenticationAPI.md#SsoInitiateLogin) | **Get** /deepfence/sso/login/{namespace} | SSO Login
[**SsoLogin**](AuthenticationAPI.md#SsoLogin) | **Get** /deepfence/sso/login | SSO Login
[**UpdateSSOProvider**](AuthenticationAPI.md#UpdateSSOProvider) | **Put** /deepfence/single-sign-on/{id} | Update Single sign-on
[**VerifySSOAuth**](AuthenticationAPI.md#VerifySSOAuth) | **Post** /deepfence/sso/verify | Verify SSO auth code



## AuthToken

> ModelResponseAccessToken AuthToken(ctx).ModelAPIAuthRequest(modelAPIAuthRequest).Execute()

Get Access Token for API Token



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
	modelAPIAuthRequest := *openapiclient.NewModelAPIAuthRequest("ApiToken_example") // ModelAPIAuthRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).ModelAPIAuthRequest(modelAPIAuthRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AuthToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthToken`: ModelResponseAccessToken
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AuthToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelAPIAuthRequest** | [**ModelAPIAuthRequest**](ModelAPIAuthRequest.md) |  | 

### Return type

[**ModelResponseAccessToken**](ModelResponseAccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenRefresh

> ModelResponseAccessToken AuthTokenRefresh(ctx).Execute()

Refresh access token



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
	resp, r, err := apiClient.AuthenticationAPI.AuthTokenRefresh(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AuthTokenRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthTokenRefresh`: ModelResponseAccessToken
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AuthTokenRefresh`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenRefreshRequest struct via the builder pattern


### Return type

[**ModelResponseAccessToken**](ModelResponseAccessToken.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSSOProvider

> CreateSSOProvider(ctx).SinglesignonSSOProviderConfig(singlesignonSSOProviderConfig).Execute()

Configure Single sign-on



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
	singlesignonSSOProviderConfig := *openapiclient.NewSinglesignonSSOProviderConfig("ClientId_example", "ClientSecret_example", false, "SsoProviderType_example") // SinglesignonSSOProviderConfig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.CreateSSOProvider(context.Background()).SinglesignonSSOProviderConfig(singlesignonSSOProviderConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.CreateSSOProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSSOProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **singlesignonSSOProviderConfig** | [**SinglesignonSSOProviderConfig**](SinglesignonSSOProviderConfig.md) |  | 

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


## DeleteSSOProvider

> DeleteSSOProvider(ctx, id).Execute()

Delete Single sign-on configuration



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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.DeleteSSOProvider(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.DeleteSSOProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSSOProviderRequest struct via the builder pattern


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


## GetSSOProviders

> SinglesignonGetSingleSignOnResponse GetSSOProviders(ctx).Execute()

Get Single sign-on configurations



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
	resp, r, err := apiClient.AuthenticationAPI.GetSSOProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetSSOProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSSOProviders`: SinglesignonGetSingleSignOnResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetSSOProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSSOProvidersRequest struct via the builder pattern


### Return type

[**SinglesignonGetSingleSignOnResponse**](SinglesignonGetSingleSignOnResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> ModelLoginResponse Login(ctx).ModelLoginRequest(modelLoginRequest).Execute()

Login API



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
	modelLoginRequest := *openapiclient.NewModelLoginRequest("Email_example", "Password_example") // ModelLoginRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.Login(context.Background()).ModelLoginRequest(modelLoginRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Login``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Login`: ModelLoginResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelLoginRequest** | [**ModelLoginRequest**](ModelLoginRequest.md) |  | 

### Return type

[**ModelLoginResponse**](ModelLoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logout

> Logout(ctx).Execute()

Logout API



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
	r, err := apiClient.AuthenticationAPI.Logout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Logout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


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


## SsoInitiateLogin

> SsoInitiateLogin(ctx, namespace).Execute()

SSO Login



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
	namespace := "namespace_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.SsoInitiateLogin(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.SsoInitiateLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSsoInitiateLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SsoLogin

> SsoLogin(ctx).Email(email).Execute()

SSO Login



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
	email := "email_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.SsoLogin(context.Background()).Email(email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.SsoLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSsoLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSSOProvider

> SinglesignonSSOResponse UpdateSSOProvider(ctx, id).SinglesignonUpdateSSOProviderConfig(singlesignonUpdateSSOProviderConfig).Execute()

Update Single sign-on



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
	id := int32(56) // int32 | 
	singlesignonUpdateSSOProviderConfig := *openapiclient.NewSinglesignonUpdateSSOProviderConfig("ClientId_example", false) // SinglesignonUpdateSSOProviderConfig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.UpdateSSOProvider(context.Background(), id).SinglesignonUpdateSSOProviderConfig(singlesignonUpdateSSOProviderConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.UpdateSSOProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSSOProvider`: SinglesignonSSOResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.UpdateSSOProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSSOProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **singlesignonUpdateSSOProviderConfig** | [**SinglesignonUpdateSSOProviderConfig**](SinglesignonUpdateSSOProviderConfig.md) |  | 

### Return type

[**SinglesignonSSOResponse**](SinglesignonSSOResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifySSOAuth

> ModelLoginResponse VerifySSOAuth(ctx).SinglesignonVerifySSOAuthRequest(singlesignonVerifySSOAuthRequest).Execute()

Verify SSO auth code



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
	singlesignonVerifySSOAuthRequest := *openapiclient.NewSinglesignonVerifySSOAuthRequest("Code_example", "Namespace_example", int32(123)) // SinglesignonVerifySSOAuthRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.VerifySSOAuth(context.Background()).SinglesignonVerifySSOAuthRequest(singlesignonVerifySSOAuthRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.VerifySSOAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifySSOAuth`: ModelLoginResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.VerifySSOAuth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifySSOAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **singlesignonVerifySSOAuthRequest** | [**SinglesignonVerifySSOAuthRequest**](SinglesignonVerifySSOAuthRequest.md) |  | 

### Return type

[**ModelLoginResponse**](ModelLoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

