# GetDmarcDomain200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to [**GetDmarcDomain200ResponseDomain**](GetDmarcDomain200ResponseDomain.md) |  | [optional] 

## Methods

### NewGetDmarcDomain200Response

`func NewGetDmarcDomain200Response() *GetDmarcDomain200Response`

NewGetDmarcDomain200Response instantiates a new GetDmarcDomain200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDmarcDomain200ResponseWithDefaults

`func NewGetDmarcDomain200ResponseWithDefaults() *GetDmarcDomain200Response`

NewGetDmarcDomain200ResponseWithDefaults instantiates a new GetDmarcDomain200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *GetDmarcDomain200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *GetDmarcDomain200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *GetDmarcDomain200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *GetDmarcDomain200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *GetDmarcDomain200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetDmarcDomain200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetDmarcDomain200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *GetDmarcDomain200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetDomain

`func (o *GetDmarcDomain200Response) GetDomain() GetDmarcDomain200ResponseDomain`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GetDmarcDomain200Response) GetDomainOk() (*GetDmarcDomain200ResponseDomain, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GetDmarcDomain200Response) SetDomain(v GetDmarcDomain200ResponseDomain)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GetDmarcDomain200Response) HasDomain() bool`

HasDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


