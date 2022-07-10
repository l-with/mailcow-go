# CreateDomainPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | domain name to which policy is associated to | [optional] 
**ObjectFrom** | Pointer to **string** | exact address or use wildcard to match whole domain | [optional] 
**ObjectList** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateDomainPolicyRequest

`func NewCreateDomainPolicyRequest() *CreateDomainPolicyRequest`

NewCreateDomainPolicyRequest instantiates a new CreateDomainPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDomainPolicyRequestWithDefaults

`func NewCreateDomainPolicyRequestWithDefaults() *CreateDomainPolicyRequest`

NewCreateDomainPolicyRequestWithDefaults instantiates a new CreateDomainPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *CreateDomainPolicyRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateDomainPolicyRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateDomainPolicyRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CreateDomainPolicyRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetObjectFrom

`func (o *CreateDomainPolicyRequest) GetObjectFrom() string`

GetObjectFrom returns the ObjectFrom field if non-nil, zero value otherwise.

### GetObjectFromOk

`func (o *CreateDomainPolicyRequest) GetObjectFromOk() (*string, bool)`

GetObjectFromOk returns a tuple with the ObjectFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectFrom

`func (o *CreateDomainPolicyRequest) SetObjectFrom(v string)`

SetObjectFrom sets ObjectFrom field to given value.

### HasObjectFrom

`func (o *CreateDomainPolicyRequest) HasObjectFrom() bool`

HasObjectFrom returns a boolean if a field has been set.

### GetObjectList

`func (o *CreateDomainPolicyRequest) GetObjectList() string`

GetObjectList returns the ObjectList field if non-nil, zero value otherwise.

### GetObjectListOk

`func (o *CreateDomainPolicyRequest) GetObjectListOk() (*string, bool)`

GetObjectListOk returns a tuple with the ObjectList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectList

`func (o *CreateDomainPolicyRequest) SetObjectList(v string)`

SetObjectList sets ObjectList field to given value.

### HasObjectList

`func (o *CreateDomainPolicyRequest) HasObjectList() bool`

HasObjectList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


