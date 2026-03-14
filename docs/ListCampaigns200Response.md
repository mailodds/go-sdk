# ListCampaigns200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** | Unique request identifier | [optional] 
**Campaigns** | Pointer to [**[]Campaign**](Campaign.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**PerPage** | Pointer to **int32** |  | [optional] 
**Pages** | Pointer to **int32** |  | [optional] 

## Methods

### NewListCampaigns200Response

`func NewListCampaigns200Response() *ListCampaigns200Response`

NewListCampaigns200Response instantiates a new ListCampaigns200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCampaigns200ResponseWithDefaults

`func NewListCampaigns200ResponseWithDefaults() *ListCampaigns200Response`

NewListCampaigns200ResponseWithDefaults instantiates a new ListCampaigns200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *ListCampaigns200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ListCampaigns200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ListCampaigns200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ListCampaigns200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *ListCampaigns200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListCampaigns200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListCampaigns200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ListCampaigns200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetCampaigns

`func (o *ListCampaigns200Response) GetCampaigns() []Campaign`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *ListCampaigns200Response) GetCampaignsOk() (*[]Campaign, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaigns

`func (o *ListCampaigns200Response) SetCampaigns(v []Campaign)`

SetCampaigns sets Campaigns field to given value.

### HasCampaigns

`func (o *ListCampaigns200Response) HasCampaigns() bool`

HasCampaigns returns a boolean if a field has been set.

### GetTotal

`func (o *ListCampaigns200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListCampaigns200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListCampaigns200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListCampaigns200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPage

`func (o *ListCampaigns200Response) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListCampaigns200Response) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListCampaigns200Response) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListCampaigns200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerPage

`func (o *ListCampaigns200Response) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *ListCampaigns200Response) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *ListCampaigns200Response) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *ListCampaigns200Response) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetPages

`func (o *ListCampaigns200Response) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *ListCampaigns200Response) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *ListCampaigns200Response) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *ListCampaigns200Response) HasPages() bool`

HasPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


