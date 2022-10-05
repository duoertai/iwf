# \DefaultApi

All URIs are relative to *http://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1WorkflowGetPost**](DefaultApi.md#ApiV1WorkflowGetPost) | **Post** /api/v1/workflow/get | get a workflow&#39;s status and results(if completed &amp; requested)
[**ApiV1WorkflowGetWithLongWaitPost**](DefaultApi.md#ApiV1WorkflowGetWithLongWaitPost) | **Post** /api/v1/workflow/getWithLongWait | get a workflow&#39;s status and results(if completed &amp; requested), wait if the workflow is still running
[**ApiV1WorkflowQueryPost**](DefaultApi.md#ApiV1WorkflowQueryPost) | **Post** /api/v1/workflow/query | query a workflow
[**ApiV1WorkflowSearchPost**](DefaultApi.md#ApiV1WorkflowSearchPost) | **Post** /api/v1/workflow/search | search for workflows by a search attribute query
[**ApiV1WorkflowSignalPost**](DefaultApi.md#ApiV1WorkflowSignalPost) | **Post** /api/v1/workflow/signal | signal a workflow
[**ApiV1WorkflowStartPost**](DefaultApi.md#ApiV1WorkflowStartPost) | **Post** /api/v1/workflow/start | start a workflow
[**ApiV1WorkflowStateDecidePost**](DefaultApi.md#ApiV1WorkflowStateDecidePost) | **Post** /api/v1/workflowState/decide | for invoking WorkflowState.decide API
[**ApiV1WorkflowStateStartPost**](DefaultApi.md#ApiV1WorkflowStateStartPost) | **Post** /api/v1/workflowState/start | for invoking WorkflowState.start API



## ApiV1WorkflowGetPost

> WorkflowGetResponse ApiV1WorkflowGetPost(ctx).WorkflowGetRequest(workflowGetRequest).Execute()

get a workflow's status and results(if completed & requested)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowGetRequest := *openapiclient.NewWorkflowGetRequest("WorkflowId_example") // WorkflowGetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiV1WorkflowGetPost(context.Background()).WorkflowGetRequest(workflowGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1WorkflowGetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1WorkflowGetPost`: WorkflowGetResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV1WorkflowGetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowGetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowGetRequest** | [**WorkflowGetRequest**](WorkflowGetRequest.md) |  | 

### Return type

[**WorkflowGetResponse**](WorkflowGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowGetWithLongWaitPost

> WorkflowGetResponse ApiV1WorkflowGetWithLongWaitPost(ctx).WorkflowGetRequest(workflowGetRequest).Execute()

get a workflow's status and results(if completed & requested), wait if the workflow is still running

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowGetRequest := *openapiclient.NewWorkflowGetRequest("WorkflowId_example") // WorkflowGetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiV1WorkflowGetWithLongWaitPost(context.Background()).WorkflowGetRequest(workflowGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1WorkflowGetWithLongWaitPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1WorkflowGetWithLongWaitPost`: WorkflowGetResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV1WorkflowGetWithLongWaitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowGetWithLongWaitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowGetRequest** | [**WorkflowGetRequest**](WorkflowGetRequest.md) |  | 

### Return type

[**WorkflowGetResponse**](WorkflowGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowQueryPost

> WorkflowQueryResponse ApiV1WorkflowQueryPost(ctx).WorkflowQueryRequest(workflowQueryRequest).Execute()

query a workflow

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowQueryRequest := *openapiclient.NewWorkflowQueryRequest("WorkflowId_example") // WorkflowQueryRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiV1WorkflowQueryPost(context.Background()).WorkflowQueryRequest(workflowQueryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1WorkflowQueryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1WorkflowQueryPost`: WorkflowQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV1WorkflowQueryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowQueryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowQueryRequest** | [**WorkflowQueryRequest**](WorkflowQueryRequest.md) |  | 

### Return type

[**WorkflowQueryResponse**](WorkflowQueryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowSearchPost

> WorkflowSearchResponse ApiV1WorkflowSearchPost(ctx).WorkflowSearchRequest(workflowSearchRequest).Execute()

search for workflows by a search attribute query

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowSearchRequest := *openapiclient.NewWorkflowSearchRequest("Query_example") // WorkflowSearchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiV1WorkflowSearchPost(context.Background()).WorkflowSearchRequest(workflowSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1WorkflowSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1WorkflowSearchPost`: WorkflowSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV1WorkflowSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowSearchRequest** | [**WorkflowSearchRequest**](WorkflowSearchRequest.md) |  | 

### Return type

[**WorkflowSearchResponse**](WorkflowSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowSignalPost

> map[string]interface{} ApiV1WorkflowSignalPost(ctx).WorkflowSignalRequest(workflowSignalRequest).Execute()

signal a workflow

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowSignalRequest := *openapiclient.NewWorkflowSignalRequest("WorkflowId_example", "SignalName_example") // WorkflowSignalRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiV1WorkflowSignalPost(context.Background()).WorkflowSignalRequest(workflowSignalRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1WorkflowSignalPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1WorkflowSignalPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV1WorkflowSignalPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowSignalPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowSignalRequest** | [**WorkflowSignalRequest**](WorkflowSignalRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowStartPost

> WorkflowStartResponse ApiV1WorkflowStartPost(ctx).WorkflowStartRequest(workflowStartRequest).Execute()

start a workflow

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowStartRequest := *openapiclient.NewWorkflowStartRequest("WorkflowId_example", "IwfWorkflowType_example", int32(123), "IwfWorkerUrl_example", "StartStateId_example") // WorkflowStartRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiV1WorkflowStartPost(context.Background()).WorkflowStartRequest(workflowStartRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1WorkflowStartPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1WorkflowStartPost`: WorkflowStartResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV1WorkflowStartPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowStartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowStartRequest** | [**WorkflowStartRequest**](WorkflowStartRequest.md) |  | 

### Return type

[**WorkflowStartResponse**](WorkflowStartResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowStateDecidePost

> WorkflowStateDecideResponse ApiV1WorkflowStateDecidePost(ctx).WorkflowStateDecideRequest(workflowStateDecideRequest).Execute()

for invoking WorkflowState.decide API

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowStateDecideRequest := *openapiclient.NewWorkflowStateDecideRequest(*openapiclient.NewContext("WorkflowId_example", "WorkflowRunId_example", int64(123), "StateExecutionId_example"), "WorkflowType_example", "WorkflowStateId_example") // WorkflowStateDecideRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiV1WorkflowStateDecidePost(context.Background()).WorkflowStateDecideRequest(workflowStateDecideRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1WorkflowStateDecidePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1WorkflowStateDecidePost`: WorkflowStateDecideResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV1WorkflowStateDecidePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowStateDecidePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowStateDecideRequest** | [**WorkflowStateDecideRequest**](WorkflowStateDecideRequest.md) |  | 

### Return type

[**WorkflowStateDecideResponse**](WorkflowStateDecideResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WorkflowStateStartPost

> WorkflowStateStartResponse ApiV1WorkflowStateStartPost(ctx).WorkflowStateStartRequest(workflowStateStartRequest).Execute()

for invoking WorkflowState.start API

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowStateStartRequest := *openapiclient.NewWorkflowStateStartRequest(*openapiclient.NewContext("WorkflowId_example", "WorkflowRunId_example", int64(123), "StateExecutionId_example"), "WorkflowType_example", "WorkflowStateId_example") // WorkflowStateStartRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiV1WorkflowStateStartPost(context.Background()).WorkflowStateStartRequest(workflowStateStartRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV1WorkflowStateStartPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1WorkflowStateStartPost`: WorkflowStateStartResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV1WorkflowStateStartPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WorkflowStateStartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowStateStartRequest** | [**WorkflowStateStartRequest**](WorkflowStateStartRequest.md) |  | 

### Return type

[**WorkflowStateStartResponse**](WorkflowStateStartResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
