# GetInactiveContactsReport200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**InactiveCount** | Pointer to **int32** |  | [optional] 
**TotalContacts** | Pointer to **int32** |  | [optional] 
**InactiveRate** | Pointer to **float32** |  | [optional] 
**ByList** | Pointer to [**[]GetInactiveContactsReport200ResponseByListInner**](GetInactiveContactsReport200ResponseByListInner.md) |  | [optional] 

## Methods

### NewGetInactiveContactsReport200Response

`func NewGetInactiveContactsReport200Response() *GetInactiveContactsReport200Response`

NewGetInactiveContactsReport200Response instantiates a new GetInactiveContactsReport200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInactiveContactsReport200ResponseWithDefaults

`func NewGetInactiveContactsReport200ResponseWithDefaults() *GetInactiveContactsReport200Response`

NewGetInactiveContactsReport200ResponseWithDefaults instantiates a new GetInactiveContactsReport200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *GetInactiveContactsReport200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *GetInactiveContactsReport200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *GetInactiveContactsReport200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *GetInactiveContactsReport200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *GetInactiveContactsReport200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetInactiveContactsReport200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetInactiveContactsReport200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *GetInactiveContactsReport200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetInactiveCount

`func (o *GetInactiveContactsReport200Response) GetInactiveCount() int32`

GetInactiveCount returns the InactiveCount field if non-nil, zero value otherwise.

### GetInactiveCountOk

`func (o *GetInactiveContactsReport200Response) GetInactiveCountOk() (*int32, bool)`

GetInactiveCountOk returns a tuple with the InactiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveCount

`func (o *GetInactiveContactsReport200Response) SetInactiveCount(v int32)`

SetInactiveCount sets InactiveCount field to given value.

### HasInactiveCount

`func (o *GetInactiveContactsReport200Response) HasInactiveCount() bool`

HasInactiveCount returns a boolean if a field has been set.

### GetTotalContacts

`func (o *GetInactiveContactsReport200Response) GetTotalContacts() int32`

GetTotalContacts returns the TotalContacts field if non-nil, zero value otherwise.

### GetTotalContactsOk

`func (o *GetInactiveContactsReport200Response) GetTotalContactsOk() (*int32, bool)`

GetTotalContactsOk returns a tuple with the TotalContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalContacts

`func (o *GetInactiveContactsReport200Response) SetTotalContacts(v int32)`

SetTotalContacts sets TotalContacts field to given value.

### HasTotalContacts

`func (o *GetInactiveContactsReport200Response) HasTotalContacts() bool`

HasTotalContacts returns a boolean if a field has been set.

### GetInactiveRate

`func (o *GetInactiveContactsReport200Response) GetInactiveRate() float32`

GetInactiveRate returns the InactiveRate field if non-nil, zero value otherwise.

### GetInactiveRateOk

`func (o *GetInactiveContactsReport200Response) GetInactiveRateOk() (*float32, bool)`

GetInactiveRateOk returns a tuple with the InactiveRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveRate

`func (o *GetInactiveContactsReport200Response) SetInactiveRate(v float32)`

SetInactiveRate sets InactiveRate field to given value.

### HasInactiveRate

`func (o *GetInactiveContactsReport200Response) HasInactiveRate() bool`

HasInactiveRate returns a boolean if a field has been set.

### GetByList

`func (o *GetInactiveContactsReport200Response) GetByList() []GetInactiveContactsReport200ResponseByListInner`

GetByList returns the ByList field if non-nil, zero value otherwise.

### GetByListOk

`func (o *GetInactiveContactsReport200Response) GetByListOk() (*[]GetInactiveContactsReport200ResponseByListInner, bool)`

GetByListOk returns a tuple with the ByList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByList

`func (o *GetInactiveContactsReport200Response) SetByList(v []GetInactiveContactsReport200ResponseByListInner)`

SetByList sets ByList field to given value.

### HasByList

`func (o *GetInactiveContactsReport200Response) HasByList() bool`

HasByList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


