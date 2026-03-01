# SuppressionAuditResponseEntriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**EventType** | Pointer to **string** | Audit event type | [optional] 
**EventCategory** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **map[string]interface{}** | Event-specific details | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSuppressionAuditResponseEntriesInner

`func NewSuppressionAuditResponseEntriesInner() *SuppressionAuditResponseEntriesInner`

NewSuppressionAuditResponseEntriesInner instantiates a new SuppressionAuditResponseEntriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuppressionAuditResponseEntriesInnerWithDefaults

`func NewSuppressionAuditResponseEntriesInnerWithDefaults() *SuppressionAuditResponseEntriesInner`

NewSuppressionAuditResponseEntriesInnerWithDefaults instantiates a new SuppressionAuditResponseEntriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SuppressionAuditResponseEntriesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SuppressionAuditResponseEntriesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SuppressionAuditResponseEntriesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SuppressionAuditResponseEntriesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEventType

`func (o *SuppressionAuditResponseEntriesInner) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SuppressionAuditResponseEntriesInner) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *SuppressionAuditResponseEntriesInner) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *SuppressionAuditResponseEntriesInner) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventCategory

`func (o *SuppressionAuditResponseEntriesInner) GetEventCategory() string`

GetEventCategory returns the EventCategory field if non-nil, zero value otherwise.

### GetEventCategoryOk

`func (o *SuppressionAuditResponseEntriesInner) GetEventCategoryOk() (*string, bool)`

GetEventCategoryOk returns a tuple with the EventCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCategory

`func (o *SuppressionAuditResponseEntriesInner) SetEventCategory(v string)`

SetEventCategory sets EventCategory field to given value.

### HasEventCategory

`func (o *SuppressionAuditResponseEntriesInner) HasEventCategory() bool`

HasEventCategory returns a boolean if a field has been set.

### GetDetails

`func (o *SuppressionAuditResponseEntriesInner) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SuppressionAuditResponseEntriesInner) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SuppressionAuditResponseEntriesInner) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *SuppressionAuditResponseEntriesInner) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SuppressionAuditResponseEntriesInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SuppressionAuditResponseEntriesInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SuppressionAuditResponseEntriesInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SuppressionAuditResponseEntriesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


