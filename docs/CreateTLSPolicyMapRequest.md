# CreateTLSPolicyMapRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | Pointer to **string** | custom parameters you find out more about them [here](http://www.postfix.org/postconf.5.html#smtp_tls_policy_maps) | [optional] 
**Active** | Pointer to **float32** | 1 for a active user account 0 for a disabled user account | [optional] 
**Dest** | Pointer to **string** | the target domain or email address | [optional] 
**Policy** | Pointer to **string** | the policy | [optional] 

## Methods

### NewCreateTLSPolicyMapRequest

`func NewCreateTLSPolicyMapRequest() *CreateTLSPolicyMapRequest`

NewCreateTLSPolicyMapRequest instantiates a new CreateTLSPolicyMapRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTLSPolicyMapRequestWithDefaults

`func NewCreateTLSPolicyMapRequestWithDefaults() *CreateTLSPolicyMapRequest`

NewCreateTLSPolicyMapRequestWithDefaults instantiates a new CreateTLSPolicyMapRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *CreateTLSPolicyMapRequest) GetParameters() string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CreateTLSPolicyMapRequest) GetParametersOk() (*string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CreateTLSPolicyMapRequest) SetParameters(v string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *CreateTLSPolicyMapRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetActive

`func (o *CreateTLSPolicyMapRequest) GetActive() float32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateTLSPolicyMapRequest) GetActiveOk() (*float32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateTLSPolicyMapRequest) SetActive(v float32)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateTLSPolicyMapRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDest

`func (o *CreateTLSPolicyMapRequest) GetDest() string`

GetDest returns the Dest field if non-nil, zero value otherwise.

### GetDestOk

`func (o *CreateTLSPolicyMapRequest) GetDestOk() (*string, bool)`

GetDestOk returns a tuple with the Dest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDest

`func (o *CreateTLSPolicyMapRequest) SetDest(v string)`

SetDest sets Dest field to given value.

### HasDest

`func (o *CreateTLSPolicyMapRequest) HasDest() bool`

HasDest returns a boolean if a field has been set.

### GetPolicy

`func (o *CreateTLSPolicyMapRequest) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *CreateTLSPolicyMapRequest) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *CreateTLSPolicyMapRequest) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *CreateTLSPolicyMapRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


