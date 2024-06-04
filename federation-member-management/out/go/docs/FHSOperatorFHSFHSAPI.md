# \FHSOperatorFHSFHSAPI

All URIs are relative to *https://tbd_fed_server.org/FHSv1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllowConnect**](FHSOperatorFHSFHSAPI.md#AllowConnect) | **Post** /FHSOperator/AllowConnection | Allow a connection from another Fed Service
[**DeleteConnection**](FHSOperatorFHSFHSAPI.md#DeleteConnection) | **Delete** /FHSOperator/Connection/{conn_id} | Delete a connection.
[**ListConnections**](FHSOperatorFHSFHSAPI.md#ListConnections) | **Get** /FHSOperator/Connections | List Connections
[**RequestConnect**](FHSOperatorFHSFHSAPI.md#RequestConnect) | **Post** /FHSOperator/Connect | Request connection to another Fed Service



## AllowConnect

> ConnectionInfo AllowConnect(ctx).AllowedConnectInfo(allowedConnectInfo).Execute()

Allow a connection from another Fed Service



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
	allowedConnectInfo := *openapiclient.NewAllowedConnectInfo("FHS_URL_example") // AllowedConnectInfo | Allowable connect information

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FHSOperatorFHSFHSAPI.AllowConnect(context.Background()).AllowedConnectInfo(allowedConnectInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FHSOperatorFHSFHSAPI.AllowConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AllowConnect`: ConnectionInfo
	fmt.Fprintf(os.Stdout, "Response from `FHSOperatorFHSFHSAPI.AllowConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllowConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allowedConnectInfo** | [**AllowedConnectInfo**](AllowedConnectInfo.md) | Allowable connect information | 

### Return type

[**ConnectionInfo**](ConnectionInfo.md)

### Authorization

[fhs_auth](../README.md#fhs_auth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnection

> DeleteConnection(ctx, connId).Execute()

Delete a connection.



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
	connId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Connection ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FHSOperatorFHSFHSAPI.DeleteConnection(context.Background(), connId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FHSOperatorFHSFHSAPI.DeleteConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connId** | **string** | The Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectionRequest struct via the builder pattern


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


## ListConnections

> []ConnectionInfo ListConnections(ctx).Execute()

List Connections



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
	resp, r, err := apiClient.FHSOperatorFHSFHSAPI.ListConnections(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FHSOperatorFHSFHSAPI.ListConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConnections`: []ConnectionInfo
	fmt.Fprintf(os.Stdout, "Response from `FHSOperatorFHSFHSAPI.ListConnections`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectionsRequest struct via the builder pattern


### Return type

[**[]ConnectionInfo**](ConnectionInfo.md)

### Authorization

[fhs_auth](../README.md#fhs_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestConnect

> ConnectionInfo RequestConnect(ctx).ConnectReq(connectReq).Execute()

Request connection to another Fed Service



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
	connectReq := *openapiclient.NewConnectReq("FHS_URL_example") // ConnectReq | Connection request information

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FHSOperatorFHSFHSAPI.RequestConnect(context.Background()).ConnectReq(connectReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FHSOperatorFHSFHSAPI.RequestConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestConnect`: ConnectionInfo
	fmt.Fprintf(os.Stdout, "Response from `FHSOperatorFHSFHSAPI.RequestConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectReq** | [**ConnectReq**](ConnectReq.md) | Connection request information | 

### Return type

[**ConnectionInfo**](ConnectionInfo.md)

### Authorization

[fhs_auth](../README.md#fhs_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

