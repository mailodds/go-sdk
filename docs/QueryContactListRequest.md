# QueryContactListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**[]QueryContactListRequestFiltersInner**](QueryContactListRequestFiltersInner.md) | Array of filter conditions | [optional] 
**Page** | Pointer to **int32** |  | [optional] [default to 1]
**PerPage** | Pointer to **int32** |  | [optional] [default to 100]

## Methods

### NewQueryContactListRequest

`func NewQueryContactListRequest() *QueryContactListRequest`

NewQueryContactListRequest instantiates a new QueryContactListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryContactListRequestWithDefaults

`func NewQueryContactListRequestWithDefaults() *QueryContactListRequest`

NewQueryContactListRequestWithDefaults instantiates a new QueryContactListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *QueryContactListRequest) GetFilters() []QueryContactListRequestFiltersInner`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *QueryContactListRequest) GetFiltersOk() (*[]QueryContactListRequestFiltersInner, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *QueryContactListRequest) SetFilters(v []QueryContactListRequestFiltersInner)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *QueryContactListRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetPage

`func (o *QueryContactListRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *QueryContactListRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *QueryContactListRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *QueryContactListRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerPage

`func (o *QueryContactListRequest) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *QueryContactListRequest) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *QueryContactListRequest) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *QueryContactListRequest) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


