# \BulkValidationAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelJob**](BulkValidationAPI.md#CancelJob) | **Post** /v1/jobs/{job_id}/cancel | Cancel a job
[**CreateJob**](BulkValidationAPI.md#CreateJob) | **Post** /v1/jobs | Create bulk validation job (JSON)
[**CreateJobFromS3**](BulkValidationAPI.md#CreateJobFromS3) | **Post** /v1/jobs/upload/s3 | Create job from S3 upload
[**CreateJobUpload**](BulkValidationAPI.md#CreateJobUpload) | **Post** /v1/jobs/upload | Create bulk validation job (file upload)
[**DeleteJob**](BulkValidationAPI.md#DeleteJob) | **Delete** /v1/jobs/{job_id} | Delete a job
[**GetJob**](BulkValidationAPI.md#GetJob) | **Get** /v1/jobs/{job_id} | Get job status
[**GetJobResults**](BulkValidationAPI.md#GetJobResults) | **Get** /v1/jobs/{job_id}/results | Get job results
[**GetPresignedUpload**](BulkValidationAPI.md#GetPresignedUpload) | **Post** /v1/jobs/upload/presigned | Get S3 presigned upload URL
[**ListJobs**](BulkValidationAPI.md#ListJobs) | **Get** /v1/jobs | List validation jobs



## CancelJob

> JobResponse CancelJob(ctx, jobId).Execute()

Cancel a job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mailodds/go-sdk"
)

func main() {
	jobId := "jobId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkValidationAPI.CancelJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkValidationAPI.CancelJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelJob`: JobResponse
	fmt.Fprintf(os.Stdout, "Response from `BulkValidationAPI.CancelJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobResponse**](JobResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJob

> JobResponse CreateJob(ctx).CreateJobRequest(createJobRequest).Execute()

Create bulk validation job (JSON)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mailodds/go-sdk"
)

func main() {
	createJobRequest := *openapiclient.NewCreateJobRequest([]string{"Emails_example"}) // CreateJobRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkValidationAPI.CreateJob(context.Background()).CreateJobRequest(createJobRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkValidationAPI.CreateJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJob`: JobResponse
	fmt.Fprintf(os.Stdout, "Response from `BulkValidationAPI.CreateJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createJobRequest** | [**CreateJobRequest**](CreateJobRequest.md) |  | 

### Return type

[**JobResponse**](JobResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJobFromS3

> JobResponse CreateJobFromS3(ctx).CreateJobFromS3Request(createJobFromS3Request).Execute()

Create job from S3 upload



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mailodds/go-sdk"
)

func main() {
	createJobFromS3Request := *openapiclient.NewCreateJobFromS3Request("S3Key_example") // CreateJobFromS3Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkValidationAPI.CreateJobFromS3(context.Background()).CreateJobFromS3Request(createJobFromS3Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkValidationAPI.CreateJobFromS3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJobFromS3`: JobResponse
	fmt.Fprintf(os.Stdout, "Response from `BulkValidationAPI.CreateJobFromS3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobFromS3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createJobFromS3Request** | [**CreateJobFromS3Request**](CreateJobFromS3Request.md) |  | 

### Return type

[**JobResponse**](JobResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJobUpload

> JobResponse CreateJobUpload(ctx).File(file).Dedup(dedup).Metadata(metadata).Execute()

Create bulk validation job (file upload)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mailodds/go-sdk"
)

func main() {
	file := os.NewFile(1234, "some_file") // *os.File | CSV, Excel (.xlsx, .xls), ODS, or TXT file
	dedup := true // bool | Remove duplicate emails (optional) (default to false)
	metadata := "metadata_example" // string | JSON metadata for the job (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkValidationAPI.CreateJobUpload(context.Background()).File(file).Dedup(dedup).Metadata(metadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkValidationAPI.CreateJobUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJobUpload`: JobResponse
	fmt.Fprintf(os.Stdout, "Response from `BulkValidationAPI.CreateJobUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | CSV, Excel (.xlsx, .xls), ODS, or TXT file | 
 **dedup** | **bool** | Remove duplicate emails | [default to false]
 **metadata** | **string** | JSON metadata for the job | 

### Return type

[**JobResponse**](JobResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteJob

> DeleteJob200Response DeleteJob(ctx, jobId).Execute()

Delete a job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mailodds/go-sdk"
)

func main() {
	jobId := "jobId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkValidationAPI.DeleteJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkValidationAPI.DeleteJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteJob`: DeleteJob200Response
	fmt.Fprintf(os.Stdout, "Response from `BulkValidationAPI.DeleteJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteJob200Response**](DeleteJob200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJob

> JobResponse GetJob(ctx, jobId).Execute()

Get job status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mailodds/go-sdk"
)

func main() {
	jobId := "jobId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkValidationAPI.GetJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkValidationAPI.GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJob`: JobResponse
	fmt.Fprintf(os.Stdout, "Response from `BulkValidationAPI.GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobResponse**](JobResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobResults

> ResultsResponse GetJobResults(ctx, jobId).Format(format).Filter(filter).Page(page).PerPage(perPage).Execute()

Get job results



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mailodds/go-sdk"
)

func main() {
	jobId := "jobId_example" // string | 
	format := "format_example" // string |  (optional) (default to "json")
	filter := "filter_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	perPage := int32(56) // int32 |  (optional) (default to 1000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkValidationAPI.GetJobResults(context.Background(), jobId).Format(format).Filter(filter).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkValidationAPI.GetJobResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobResults`: ResultsResponse
	fmt.Fprintf(os.Stdout, "Response from `BulkValidationAPI.GetJobResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** |  | [default to &quot;json&quot;]
 **filter** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **perPage** | **int32** |  | [default to 1000]

### Return type

[**ResultsResponse**](ResultsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPresignedUpload

> PresignedUploadResponse GetPresignedUpload(ctx).GetPresignedUploadRequest(getPresignedUploadRequest).Execute()

Get S3 presigned upload URL



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mailodds/go-sdk"
)

func main() {
	getPresignedUploadRequest := *openapiclient.NewGetPresignedUploadRequest("Filename_example") // GetPresignedUploadRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkValidationAPI.GetPresignedUpload(context.Background()).GetPresignedUploadRequest(getPresignedUploadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkValidationAPI.GetPresignedUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPresignedUpload`: PresignedUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `BulkValidationAPI.GetPresignedUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPresignedUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPresignedUploadRequest** | [**GetPresignedUploadRequest**](GetPresignedUploadRequest.md) |  | 

### Return type

[**PresignedUploadResponse**](PresignedUploadResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListJobs

> JobListResponse ListJobs(ctx).Page(page).PerPage(perPage).Status(status).Execute()

List validation jobs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mailodds/go-sdk"
)

func main() {
	page := int32(56) // int32 |  (optional) (default to 1)
	perPage := int32(56) // int32 |  (optional) (default to 20)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkValidationAPI.ListJobs(context.Background()).Page(page).PerPage(perPage).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkValidationAPI.ListJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListJobs`: JobListResponse
	fmt.Fprintf(os.Stdout, "Response from `BulkValidationAPI.ListJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **perPage** | **int32** |  | [default to 20]
 **status** | **string** |  | 

### Return type

[**JobListResponse**](JobListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

