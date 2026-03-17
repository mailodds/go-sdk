# \ContactListsAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddContact**](ContactListsAPI.md#AddContact) | **Post** /v1/contact-lists/{list_id}/contacts | Add contact to list
[**AppendToContactList**](ContactListsAPI.md#AppendToContactList) | **Post** /v1/contact-lists/{list_id}/append | Append to contact list
[**CreateContactList**](ContactListsAPI.md#CreateContactList) | **Post** /v1/contact-lists | Create contact list
[**DeleteContact**](ContactListsAPI.md#DeleteContact) | **Delete** /v1/contact-lists/{list_id}/contacts/{contact_id} | Delete contact
[**DeleteContactList**](ContactListsAPI.md#DeleteContactList) | **Delete** /v1/contact-lists/{list_id} | Delete a contact list
[**ExportContactList**](ContactListsAPI.md#ExportContactList) | **Get** /v1/contact-lists/{list_id}/export | Export contact list
[**GetInactiveContactsReport**](ContactListsAPI.md#GetInactiveContactsReport) | **Get** /v1/contacts/inactive-report | Get inactive contacts report
[**ImportContactList**](ContactListsAPI.md#ImportContactList) | **Post** /v1/contact-lists/{list_id}/import | Import contacts from CSV
[**ListContactLists**](ContactListsAPI.md#ListContactLists) | **Get** /v1/contact-lists | List contact lists
[**QueryContactList**](ContactListsAPI.md#QueryContactList) | **Post** /v1/contact-lists/{list_id}/query | Query contact list
[**UpdateContact**](ContactListsAPI.md#UpdateContact) | **Patch** /v1/contact-lists/{list_id}/contacts/{contact_id} | Update contact



## AddContact

> AddContact201Response AddContact(ctx, listId).AddContactRequest(addContactRequest).Execute()

Add contact to list



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
	listId := "listId_example" // string | Contact list ID
	addContactRequest := *openapiclient.NewAddContactRequest("Email_example") // AddContactRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactListsAPI.AddContact(context.Background(), listId).AddContactRequest(addContactRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.AddContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddContact`: AddContact201Response
	fmt.Fprintf(os.Stdout, "Response from `ContactListsAPI.AddContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | Contact list ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addContactRequest** | [**AddContactRequest**](AddContactRequest.md) |  | 

### Return type

[**AddContact201Response**](AddContact201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppendToContactList

> AppendToContactList200Response AppendToContactList(ctx, listId).AppendToContactListRequest(appendToContactListRequest).Execute()

Append to contact list



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
	listId := "listId_example" // string | Contact list UUID
	appendToContactListRequest := *openapiclient.NewAppendToContactListRequest([]string{"SourceJobIds_example"}) // AppendToContactListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactListsAPI.AppendToContactList(context.Background(), listId).AppendToContactListRequest(appendToContactListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.AppendToContactList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppendToContactList`: AppendToContactList200Response
	fmt.Fprintf(os.Stdout, "Response from `ContactListsAPI.AppendToContactList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | Contact list UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppendToContactListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appendToContactListRequest** | [**AppendToContactListRequest**](AppendToContactListRequest.md) |  | 

### Return type

[**AppendToContactList200Response**](AppendToContactList200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContactList

> CreateContactList201Response CreateContactList(ctx).CreateContactListRequest(createContactListRequest).Execute()

Create contact list



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
	createContactListRequest := *openapiclient.NewCreateContactListRequest("Name_example") // CreateContactListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactListsAPI.CreateContactList(context.Background()).CreateContactListRequest(createContactListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.CreateContactList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContactList`: CreateContactList201Response
	fmt.Fprintf(os.Stdout, "Response from `ContactListsAPI.CreateContactList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContactListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createContactListRequest** | [**CreateContactListRequest**](CreateContactListRequest.md) |  | 

### Return type

[**CreateContactList201Response**](CreateContactList201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContact

> DeletePolicyRule200Response DeleteContact(ctx, listId, contactId).Execute()

Delete contact



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
	listId := "listId_example" // string | Contact list ID
	contactId := "contactId_example" // string | Contact ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactListsAPI.DeleteContact(context.Background(), listId, contactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.DeleteContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteContact`: DeletePolicyRule200Response
	fmt.Fprintf(os.Stdout, "Response from `ContactListsAPI.DeleteContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | Contact list ID | 
**contactId** | **string** | Contact ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeletePolicyRule200Response**](DeletePolicyRule200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContactList

> DeletePolicyRule200Response DeleteContactList(ctx, listId).Execute()

Delete a contact list



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
	listId := "listId_example" // string | Contact list UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactListsAPI.DeleteContactList(context.Background(), listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.DeleteContactList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteContactList`: DeletePolicyRule200Response
	fmt.Fprintf(os.Stdout, "Response from `ContactListsAPI.DeleteContactList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | Contact list UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContactListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeletePolicyRule200Response**](DeletePolicyRule200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportContactList

> string ExportContactList(ctx, listId).Execute()

Export contact list



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
	listId := "listId_example" // string | Contact list ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactListsAPI.ExportContactList(context.Background(), listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.ExportContactList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportContactList`: string
	fmt.Fprintf(os.Stdout, "Response from `ContactListsAPI.ExportContactList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | Contact list ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportContactListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInactiveContactsReport

> GetInactiveContactsReport200Response GetInactiveContactsReport(ctx).Days(days).Execute()

Get inactive contacts report



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
	days := int32(56) // int32 | Inactivity threshold in days (optional) (default to 90)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactListsAPI.GetInactiveContactsReport(context.Background()).Days(days).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.GetInactiveContactsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInactiveContactsReport`: GetInactiveContactsReport200Response
	fmt.Fprintf(os.Stdout, "Response from `ContactListsAPI.GetInactiveContactsReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInactiveContactsReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **int32** | Inactivity threshold in days | [default to 90]

### Return type

[**GetInactiveContactsReport200Response**](GetInactiveContactsReport200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportContactList

> ImportContactList200Response ImportContactList(ctx, listId).File(file).ColumnMapping(columnMapping).ConsentSource(consentSource).Tags(tags).Execute()

Import contacts from CSV



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
	listId := "listId_example" // string | Contact list ID
	file := os.NewFile(1234, "some_file") // *os.File | CSV file (max 10MB)
	columnMapping := "columnMapping_example" // string | JSON mapping of CSV columns to contact fields (optional)
	consentSource := "consentSource_example" // string | Source of consent for imported contacts (optional)
	tags := "tags_example" // string | JSON array of tags to apply (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactListsAPI.ImportContactList(context.Background(), listId).File(file).ColumnMapping(columnMapping).ConsentSource(consentSource).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.ImportContactList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportContactList`: ImportContactList200Response
	fmt.Fprintf(os.Stdout, "Response from `ContactListsAPI.ImportContactList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | Contact list ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportContactListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | CSV file (max 10MB) | 
 **columnMapping** | **string** | JSON mapping of CSV columns to contact fields | 
 **consentSource** | **string** | Source of consent for imported contacts | 
 **tags** | **string** | JSON array of tags to apply | 

### Return type

[**ImportContactList200Response**](ImportContactList200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContactLists

> ListContactLists200Response ListContactLists(ctx).Page(page).PerPage(perPage).Execute()

List contact lists



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
	page := int32(56) // int32 |  (optional) (default to 1)
	perPage := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactListsAPI.ListContactLists(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.ListContactLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContactLists`: ListContactLists200Response
	fmt.Fprintf(os.Stdout, "Response from `ContactListsAPI.ListContactLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListContactListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **perPage** | **int32** |  | [default to 20]

### Return type

[**ListContactLists200Response**](ListContactLists200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryContactList

> QueryContactList200Response QueryContactList(ctx, listId).QueryContactListRequest(queryContactListRequest).Execute()

Query contact list



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
	listId := "listId_example" // string | Contact list UUID
	queryContactListRequest := *openapiclient.NewQueryContactListRequest() // QueryContactListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactListsAPI.QueryContactList(context.Background(), listId).QueryContactListRequest(queryContactListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.QueryContactList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryContactList`: QueryContactList200Response
	fmt.Fprintf(os.Stdout, "Response from `ContactListsAPI.QueryContactList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | Contact list UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryContactListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queryContactListRequest** | [**QueryContactListRequest**](QueryContactListRequest.md) |  | 

### Return type

[**QueryContactList200Response**](QueryContactList200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContact

> AddContact201Response UpdateContact(ctx, listId, contactId).UpdateContactRequest(updateContactRequest).Execute()

Update contact



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
	listId := "listId_example" // string | Contact list ID
	contactId := "contactId_example" // string | Contact ID
	updateContactRequest := *openapiclient.NewUpdateContactRequest() // UpdateContactRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactListsAPI.UpdateContact(context.Background(), listId, contactId).UpdateContactRequest(updateContactRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.UpdateContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContact`: AddContact201Response
	fmt.Fprintf(os.Stdout, "Response from `ContactListsAPI.UpdateContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | Contact list ID | 
**contactId** | **string** | Contact ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateContactRequest** | [**UpdateContactRequest**](UpdateContactRequest.md) |  | 

### Return type

[**AddContact201Response**](AddContact201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

