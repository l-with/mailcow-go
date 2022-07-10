# EditMailboxRatelimitsRequestAttr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RlFrame** | Pointer to **string** | contains the frame for the ratelimit h,s,m | [optional] 
**RlValue** | Pointer to **float32** | contains the rate for the ratelimit 10,20,50,1 | [optional] 

## Methods

### NewEditMailboxRatelimitsRequestAttr

`func NewEditMailboxRatelimitsRequestAttr() *EditMailboxRatelimitsRequestAttr`

NewEditMailboxRatelimitsRequestAttr instantiates a new EditMailboxRatelimitsRequestAttr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditMailboxRatelimitsRequestAttrWithDefaults

`func NewEditMailboxRatelimitsRequestAttrWithDefaults() *EditMailboxRatelimitsRequestAttr`

NewEditMailboxRatelimitsRequestAttrWithDefaults instantiates a new EditMailboxRatelimitsRequestAttr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRlFrame

`func (o *EditMailboxRatelimitsRequestAttr) GetRlFrame() string`

GetRlFrame returns the RlFrame field if non-nil, zero value otherwise.

### GetRlFrameOk

`func (o *EditMailboxRatelimitsRequestAttr) GetRlFrameOk() (*string, bool)`

GetRlFrameOk returns a tuple with the RlFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlFrame

`func (o *EditMailboxRatelimitsRequestAttr) SetRlFrame(v string)`

SetRlFrame sets RlFrame field to given value.

### HasRlFrame

`func (o *EditMailboxRatelimitsRequestAttr) HasRlFrame() bool`

HasRlFrame returns a boolean if a field has been set.

### GetRlValue

`func (o *EditMailboxRatelimitsRequestAttr) GetRlValue() float32`

GetRlValue returns the RlValue field if non-nil, zero value otherwise.

### GetRlValueOk

`func (o *EditMailboxRatelimitsRequestAttr) GetRlValueOk() (*float32, bool)`

GetRlValueOk returns a tuple with the RlValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlValue

`func (o *EditMailboxRatelimitsRequestAttr) SetRlValue(v float32)`

SetRlValue sets RlValue field to given value.

### HasRlValue

`func (o *EditMailboxRatelimitsRequestAttr) HasRlValue() bool`

HasRlValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


