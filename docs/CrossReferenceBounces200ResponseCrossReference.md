# CrossReferenceBounces200ResponseCrossReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalBounced** | Pointer to **int32** |  | [optional] 
**Matched** | Pointer to **int32** | Bounced emails found in validation logs | [optional] 
**Unmatched** | Pointer to **int32** | Bounced emails not in validation logs | [optional] 
**MatchRate** | Pointer to **float32** |  | [optional] 
**Entries** | Pointer to [**[]CrossReferenceBounces200ResponseCrossReferenceEntriesInner**](CrossReferenceBounces200ResponseCrossReferenceEntriesInner.md) |  | [optional] 

## Methods

### NewCrossReferenceBounces200ResponseCrossReference

`func NewCrossReferenceBounces200ResponseCrossReference() *CrossReferenceBounces200ResponseCrossReference`

NewCrossReferenceBounces200ResponseCrossReference instantiates a new CrossReferenceBounces200ResponseCrossReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossReferenceBounces200ResponseCrossReferenceWithDefaults

`func NewCrossReferenceBounces200ResponseCrossReferenceWithDefaults() *CrossReferenceBounces200ResponseCrossReference`

NewCrossReferenceBounces200ResponseCrossReferenceWithDefaults instantiates a new CrossReferenceBounces200ResponseCrossReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalBounced

`func (o *CrossReferenceBounces200ResponseCrossReference) GetTotalBounced() int32`

GetTotalBounced returns the TotalBounced field if non-nil, zero value otherwise.

### GetTotalBouncedOk

`func (o *CrossReferenceBounces200ResponseCrossReference) GetTotalBouncedOk() (*int32, bool)`

GetTotalBouncedOk returns a tuple with the TotalBounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBounced

`func (o *CrossReferenceBounces200ResponseCrossReference) SetTotalBounced(v int32)`

SetTotalBounced sets TotalBounced field to given value.

### HasTotalBounced

`func (o *CrossReferenceBounces200ResponseCrossReference) HasTotalBounced() bool`

HasTotalBounced returns a boolean if a field has been set.

### GetMatched

`func (o *CrossReferenceBounces200ResponseCrossReference) GetMatched() int32`

GetMatched returns the Matched field if non-nil, zero value otherwise.

### GetMatchedOk

`func (o *CrossReferenceBounces200ResponseCrossReference) GetMatchedOk() (*int32, bool)`

GetMatchedOk returns a tuple with the Matched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatched

`func (o *CrossReferenceBounces200ResponseCrossReference) SetMatched(v int32)`

SetMatched sets Matched field to given value.

### HasMatched

`func (o *CrossReferenceBounces200ResponseCrossReference) HasMatched() bool`

HasMatched returns a boolean if a field has been set.

### GetUnmatched

`func (o *CrossReferenceBounces200ResponseCrossReference) GetUnmatched() int32`

GetUnmatched returns the Unmatched field if non-nil, zero value otherwise.

### GetUnmatchedOk

`func (o *CrossReferenceBounces200ResponseCrossReference) GetUnmatchedOk() (*int32, bool)`

GetUnmatchedOk returns a tuple with the Unmatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatched

`func (o *CrossReferenceBounces200ResponseCrossReference) SetUnmatched(v int32)`

SetUnmatched sets Unmatched field to given value.

### HasUnmatched

`func (o *CrossReferenceBounces200ResponseCrossReference) HasUnmatched() bool`

HasUnmatched returns a boolean if a field has been set.

### GetMatchRate

`func (o *CrossReferenceBounces200ResponseCrossReference) GetMatchRate() float32`

GetMatchRate returns the MatchRate field if non-nil, zero value otherwise.

### GetMatchRateOk

`func (o *CrossReferenceBounces200ResponseCrossReference) GetMatchRateOk() (*float32, bool)`

GetMatchRateOk returns a tuple with the MatchRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRate

`func (o *CrossReferenceBounces200ResponseCrossReference) SetMatchRate(v float32)`

SetMatchRate sets MatchRate field to given value.

### HasMatchRate

`func (o *CrossReferenceBounces200ResponseCrossReference) HasMatchRate() bool`

HasMatchRate returns a boolean if a field has been set.

### GetEntries

`func (o *CrossReferenceBounces200ResponseCrossReference) GetEntries() []CrossReferenceBounces200ResponseCrossReferenceEntriesInner`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *CrossReferenceBounces200ResponseCrossReference) GetEntriesOk() (*[]CrossReferenceBounces200ResponseCrossReferenceEntriesInner, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *CrossReferenceBounces200ResponseCrossReference) SetEntries(v []CrossReferenceBounces200ResponseCrossReferenceEntriesInner)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *CrossReferenceBounces200ResponseCrossReference) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


