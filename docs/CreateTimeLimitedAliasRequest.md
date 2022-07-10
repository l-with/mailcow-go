# CreateTimeLimitedAliasRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | the mailbox an alias should be created for | [optional] 
**Domain** | Pointer to **string** | the domain | [optional] 

## Methods

### NewCreateTimeLimitedAliasRequest

`func NewCreateTimeLimitedAliasRequest() *CreateTimeLimitedAliasRequest`

NewCreateTimeLimitedAliasRequest instantiates a new CreateTimeLimitedAliasRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTimeLimitedAliasRequestWithDefaults

`func NewCreateTimeLimitedAliasRequestWithDefaults() *CreateTimeLimitedAliasRequest`

NewCreateTimeLimitedAliasRequestWithDefaults instantiates a new CreateTimeLimitedAliasRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CreateTimeLimitedAliasRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateTimeLimitedAliasRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateTimeLimitedAliasRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateTimeLimitedAliasRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDomain

`func (o *CreateTimeLimitedAliasRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateTimeLimitedAliasRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateTimeLimitedAliasRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CreateTimeLimitedAliasRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


