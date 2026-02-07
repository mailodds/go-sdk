# TelemetrySummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**Window** | Pointer to **string** |  | [optional] 
**GeneratedAt** | Pointer to **time.Time** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**Totals** | Pointer to [**TelemetrySummaryTotals**](TelemetrySummaryTotals.md) |  | [optional] 
**Rates** | Pointer to [**TelemetrySummaryRates**](TelemetrySummaryRates.md) |  | [optional] 
**TopReasons** | Pointer to [**[]TelemetrySummaryTopReasonsInner**](TelemetrySummaryTopReasonsInner.md) | Top rejection/status reasons | [optional] 
**TopDomains** | Pointer to [**[]TelemetrySummaryTopDomainsInner**](TelemetrySummaryTopDomainsInner.md) | Top domains by volume | [optional] 

## Methods

### NewTelemetrySummary

`func NewTelemetrySummary() *TelemetrySummary`

NewTelemetrySummary instantiates a new TelemetrySummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetrySummaryWithDefaults

`func NewTelemetrySummaryWithDefaults() *TelemetrySummary`

NewTelemetrySummaryWithDefaults instantiates a new TelemetrySummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *TelemetrySummary) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *TelemetrySummary) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *TelemetrySummary) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *TelemetrySummary) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetWindow

`func (o *TelemetrySummary) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *TelemetrySummary) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *TelemetrySummary) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *TelemetrySummary) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *TelemetrySummary) GetGeneratedAt() time.Time`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *TelemetrySummary) GetGeneratedAtOk() (*time.Time, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *TelemetrySummary) SetGeneratedAt(v time.Time)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *TelemetrySummary) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetTimezone

`func (o *TelemetrySummary) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *TelemetrySummary) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *TelemetrySummary) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *TelemetrySummary) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTotals

`func (o *TelemetrySummary) GetTotals() TelemetrySummaryTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *TelemetrySummary) GetTotalsOk() (*TelemetrySummaryTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *TelemetrySummary) SetTotals(v TelemetrySummaryTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *TelemetrySummary) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### GetRates

`func (o *TelemetrySummary) GetRates() TelemetrySummaryRates`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *TelemetrySummary) GetRatesOk() (*TelemetrySummaryRates, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *TelemetrySummary) SetRates(v TelemetrySummaryRates)`

SetRates sets Rates field to given value.

### HasRates

`func (o *TelemetrySummary) HasRates() bool`

HasRates returns a boolean if a field has been set.

### GetTopReasons

`func (o *TelemetrySummary) GetTopReasons() []TelemetrySummaryTopReasonsInner`

GetTopReasons returns the TopReasons field if non-nil, zero value otherwise.

### GetTopReasonsOk

`func (o *TelemetrySummary) GetTopReasonsOk() (*[]TelemetrySummaryTopReasonsInner, bool)`

GetTopReasonsOk returns a tuple with the TopReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopReasons

`func (o *TelemetrySummary) SetTopReasons(v []TelemetrySummaryTopReasonsInner)`

SetTopReasons sets TopReasons field to given value.

### HasTopReasons

`func (o *TelemetrySummary) HasTopReasons() bool`

HasTopReasons returns a boolean if a field has been set.

### GetTopDomains

`func (o *TelemetrySummary) GetTopDomains() []TelemetrySummaryTopDomainsInner`

GetTopDomains returns the TopDomains field if non-nil, zero value otherwise.

### GetTopDomainsOk

`func (o *TelemetrySummary) GetTopDomainsOk() (*[]TelemetrySummaryTopDomainsInner, bool)`

GetTopDomainsOk returns a tuple with the TopDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopDomains

`func (o *TelemetrySummary) SetTopDomains(v []TelemetrySummaryTopDomainsInner)`

SetTopDomains sets TopDomains field to given value.

### HasTopDomains

`func (o *TelemetrySummary) HasTopDomains() bool`

HasTopDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


