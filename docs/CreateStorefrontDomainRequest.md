# CreateStorefrontDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | **string** | Fully qualified domain name | 
**StoreId** | **string** | Store connection ID | 

## Methods

### NewCreateStorefrontDomainRequest

`func NewCreateStorefrontDomainRequest(fqdn string, storeId string, ) *CreateStorefrontDomainRequest`

NewCreateStorefrontDomainRequest instantiates a new CreateStorefrontDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStorefrontDomainRequestWithDefaults

`func NewCreateStorefrontDomainRequestWithDefaults() *CreateStorefrontDomainRequest`

NewCreateStorefrontDomainRequestWithDefaults instantiates a new CreateStorefrontDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *CreateStorefrontDomainRequest) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *CreateStorefrontDomainRequest) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *CreateStorefrontDomainRequest) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetStoreId

`func (o *CreateStorefrontDomainRequest) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *CreateStorefrontDomainRequest) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *CreateStorefrontDomainRequest) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


