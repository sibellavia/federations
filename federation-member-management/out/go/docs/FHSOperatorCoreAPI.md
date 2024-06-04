# \FHSOperatorCoreAPI

All URIs are relative to *https://tbd_fed_server.org/FHSv1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFedAdmin**](FHSOperatorCoreAPI.md#DeleteFedAdmin) | **Delete** /FHSOperator/FedAdmin/{member_id} | Delete a FedAdmin.
[**DeleteFederation**](FHSOperatorCoreAPI.md#DeleteFederation) | **Delete** /FHSOperator/FedInstance/{fed_id} | Delete a federation.
[**ListFedAdmins**](FHSOperatorCoreAPI.md#ListFedAdmins) | **Get** /FHSOperator/FedAdmins | List federation admins
[**ListFedInstances**](FHSOperatorCoreAPI.md#ListFedInstances) | **Get** /FHSOperator/FedInstances | List federation instances
[**NewFedAdmin**](FHSOperatorCoreAPI.md#NewFedAdmin) | **Post** /FHSOperator/NewFedAdmin | Add a new user authorized to create and administer federations
[**UpdateFedAdmin**](FHSOperatorCoreAPI.md#UpdateFedAdmin) | **Put** /FHSOperator/FedAdmin/{member_id} | Update FedAdmin information.
[**UpdateFedInstance**](FHSOperatorCoreAPI.md#UpdateFedInstance) | **Put** /FHSOperator/FedInstance/{fed_id} | Update FedInstance information.



## DeleteFedAdmin

> DeleteFedAdmin(ctx, memberId).Execute()

Delete a FedAdmin.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	memberId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The FedAdmin MemberID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FHSOperatorCoreAPI.DeleteFedAdmin(context.Background(), memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FHSOperatorCoreAPI.DeleteFedAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberId** | **string** | The FedAdmin MemberID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFedAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[fhs_auth](../README.md#fhs_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFederation

> DeleteFederation(ctx, fedId).Execute()

Delete a federation.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	fedId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The federation ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FHSOperatorCoreAPI.DeleteFederation(context.Background(), fedId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FHSOperatorCoreAPI.DeleteFederation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fedId** | **string** | The federation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFederationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[fhs_auth](../README.md#fhs_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFedAdmins

> []FedAdminInfo ListFedAdmins(ctx).Execute()

List federation admins



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FHSOperatorCoreAPI.ListFedAdmins(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FHSOperatorCoreAPI.ListFedAdmins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFedAdmins`: []FedAdminInfo
	fmt.Fprintf(os.Stdout, "Response from `FHSOperatorCoreAPI.ListFedAdmins`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFedAdminsRequest struct via the builder pattern


### Return type

[**[]FedAdminInfo**](FedAdminInfo.md)

### Authorization

[fhs_auth](../README.md#fhs_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFedInstances

> []FedInstanceInfo ListFedInstances(ctx).Execute()

List federation instances



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FHSOperatorCoreAPI.ListFedInstances(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FHSOperatorCoreAPI.ListFedInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFedInstances`: []FedInstanceInfo
	fmt.Fprintf(os.Stdout, "Response from `FHSOperatorCoreAPI.ListFedInstances`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFedInstancesRequest struct via the builder pattern


### Return type

[**[]FedInstanceInfo**](FedInstanceInfo.md)

### Authorization

[fhs_auth](../README.md#fhs_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewFedAdmin

> FedAdminID NewFedAdmin(ctx).NewFedAdmin(newFedAdmin).Execute()

Add a new user authorized to create and administer federations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	newFedAdmin := *openapiclient.NewNewFedAdmin("Name_example", false) // NewFedAdmin | New FedAdmin information

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FHSOperatorCoreAPI.NewFedAdmin(context.Background()).NewFedAdmin(newFedAdmin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FHSOperatorCoreAPI.NewFedAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NewFedAdmin`: FedAdminID
	fmt.Fprintf(os.Stdout, "Response from `FHSOperatorCoreAPI.NewFedAdmin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewFedAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newFedAdmin** | [**NewFedAdmin**](NewFedAdmin.md) | New FedAdmin information | 

### Return type

[**FedAdminID**](FedAdminID.md)

### Authorization

[fhs_auth](../README.md#fhs_auth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFedAdmin

> UpdateFedAdmin(ctx, memberId).FedAdminUpdateInfo(fedAdminUpdateInfo).Execute()

Update FedAdmin information.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	memberId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The FedAdmin MemberID
	fedAdminUpdateInfo := *openapiclient.NewFedAdminUpdateInfo() // FedAdminUpdateInfo | FedAdmin update information

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FHSOperatorCoreAPI.UpdateFedAdmin(context.Background(), memberId).FedAdminUpdateInfo(fedAdminUpdateInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FHSOperatorCoreAPI.UpdateFedAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberId** | **string** | The FedAdmin MemberID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFedAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fedAdminUpdateInfo** | [**FedAdminUpdateInfo**](FedAdminUpdateInfo.md) | FedAdmin update information | 

### Return type

 (empty response body)

### Authorization

[fhs_auth](../README.md#fhs_auth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFedInstance

> UpdateFedInstance(ctx, fedId).FedInstanceUpdateInfo(fedInstanceUpdateInfo).Execute()

Update FedInstance information.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	fedId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Federation ID
	fedInstanceUpdateInfo := *openapiclient.NewFedInstanceUpdateInfo(false) // FedInstanceUpdateInfo | Federation update information

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FHSOperatorCoreAPI.UpdateFedInstance(context.Background(), fedId).FedInstanceUpdateInfo(fedInstanceUpdateInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FHSOperatorCoreAPI.UpdateFedInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fedId** | **string** | The Federation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFedInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fedInstanceUpdateInfo** | [**FedInstanceUpdateInfo**](FedInstanceUpdateInfo.md) | Federation update information | 

### Return type

 (empty response body)

### Authorization

[fhs_auth](../README.md#fhs_auth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

