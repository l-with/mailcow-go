# DuplicateDKIMKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FronDomain** | Pointer to **string** | the domain where the dkim key should be copied from | [optional] 
**ToDomain** | Pointer to **string** | the domain where the dkim key should be copied to | [optional] 

## Methods

### NewDuplicateDKIMKeyRequest

`func NewDuplicateDKIMKeyRequest() *DuplicateDKIMKeyRequest`

NewDuplicateDKIMKeyRequest instantiates a new DuplicateDKIMKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDuplicateDKIMKeyRequestWithDefaults

`func NewDuplicateDKIMKeyRequestWithDefaults() *DuplicateDKIMKeyRequest`

NewDuplicateDKIMKeyRequestWithDefaults instantiates a new DuplicateDKIMKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFronDomain

`func (o *DuplicateDKIMKeyRequest) GetFronDomain() string`

GetFronDomain returns the FronDomain field if non-nil, zero value otherwise.

### GetFronDomainOk

`func (o *DuplicateDKIMKeyRequest) GetFronDomainOk() (*string, bool)`

GetFronDomainOk returns a tuple with the FronDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFronDomain

`func (o *DuplicateDKIMKeyRequest) SetFronDomain(v string)`

SetFronDomain sets FronDomain field to given value.

### HasFronDomain

`func (o *DuplicateDKIMKeyRequest) HasFronDomain() bool`

HasFronDomain returns a boolean if a field has been set.

### GetToDomain

`func (o *DuplicateDKIMKeyRequest) GetToDomain() string`

GetToDomain returns the ToDomain field if non-nil, zero value otherwise.

### GetToDomainOk

`func (o *DuplicateDKIMKeyRequest) GetToDomainOk() (*string, bool)`

GetToDomainOk returns a tuple with the ToDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDomain

`func (o *DuplicateDKIMKeyRequest) SetToDomain(v string)`

SetToDomain sets ToDomain field to given value.

### HasToDomain

`func (o *DuplicateDKIMKeyRequest) HasToDomain() bool`

HasToDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


