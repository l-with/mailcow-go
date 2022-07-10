# GenerateDKIMKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DkimSelector** | Pointer to **string** | the DKIM selector default dkim | [optional] 
**Domains** | Pointer to **string** | a list of domains for which a dkim key should be generated | [optional] 
**KeySize** | Pointer to **float32** | the key size (1024 or 2048) | [optional] 

## Methods

### NewGenerateDKIMKeyRequest

`func NewGenerateDKIMKeyRequest() *GenerateDKIMKeyRequest`

NewGenerateDKIMKeyRequest instantiates a new GenerateDKIMKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateDKIMKeyRequestWithDefaults

`func NewGenerateDKIMKeyRequestWithDefaults() *GenerateDKIMKeyRequest`

NewGenerateDKIMKeyRequestWithDefaults instantiates a new GenerateDKIMKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDkimSelector

`func (o *GenerateDKIMKeyRequest) GetDkimSelector() string`

GetDkimSelector returns the DkimSelector field if non-nil, zero value otherwise.

### GetDkimSelectorOk

`func (o *GenerateDKIMKeyRequest) GetDkimSelectorOk() (*string, bool)`

GetDkimSelectorOk returns a tuple with the DkimSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkimSelector

`func (o *GenerateDKIMKeyRequest) SetDkimSelector(v string)`

SetDkimSelector sets DkimSelector field to given value.

### HasDkimSelector

`func (o *GenerateDKIMKeyRequest) HasDkimSelector() bool`

HasDkimSelector returns a boolean if a field has been set.

### GetDomains

`func (o *GenerateDKIMKeyRequest) GetDomains() string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *GenerateDKIMKeyRequest) GetDomainsOk() (*string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *GenerateDKIMKeyRequest) SetDomains(v string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *GenerateDKIMKeyRequest) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetKeySize

`func (o *GenerateDKIMKeyRequest) GetKeySize() float32`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *GenerateDKIMKeyRequest) GetKeySizeOk() (*float32, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *GenerateDKIMKeyRequest) SetKeySize(v float32)`

SetKeySize sets KeySize field to given value.

### HasKeySize

`func (o *GenerateDKIMKeyRequest) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


