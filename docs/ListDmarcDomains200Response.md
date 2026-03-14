# ListDmarcDomains200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Domains** | Pointer to [**[]DmarcDomain**](DmarcDomain.md) |  | [optional] 

## Methods

### NewListDmarcDomains200Response

`func NewListDmarcDomains200Response() *ListDmarcDomains200Response`

NewListDmarcDomains200Response instantiates a new ListDmarcDomains200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDmarcDomains200ResponseWithDefaults

`func NewListDmarcDomains200ResponseWithDefaults() *ListDmarcDomains200Response`

NewListDmarcDomains200ResponseWithDefaults instantiates a new ListDmarcDomains200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *ListDmarcDomains200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ListDmarcDomains200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ListDmarcDomains200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ListDmarcDomains200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *ListDmarcDomains200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListDmarcDomains200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListDmarcDomains200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ListDmarcDomains200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetDomains

`func (o *ListDmarcDomains200Response) GetDomains() []DmarcDomain`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *ListDmarcDomains200Response) GetDomainsOk() (*[]DmarcDomain, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *ListDmarcDomains200Response) SetDomains(v []DmarcDomain)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *ListDmarcDomains200Response) HasDomains() bool`

HasDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


