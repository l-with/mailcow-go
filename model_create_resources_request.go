/*
mailcow API

mailcow is complete e-mailing solution with advanced antispam, antivirus, nice UI and API.  In order to use this API you have to create a API key and add your IP address to the whitelist of allowed IPs this can be done by logging into the Mailcow UI using your admin account, then go to Configuration > Access > Edit administrator details > API. There you will find a collapsed API menu.  There are two types of API keys   - The read only key can only be used for all get endpoints   - The read write key can be used for all endpoints

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateResourcesRequest struct for CreateResourcesRequest
type CreateResourcesRequest struct {
	// 1 for a active transport map 0 for a disabled transport map
	Active *float32 `json:"active,omitempty"`
	// a description of the resource
	Description *string `json:"description,omitempty"`
	// the domain for which the resource should be
	Domain *string `json:"domain,omitempty"`
	// the kind of recouse
	Kind             *string `json:"kind,omitempty"`
	MultipleBookings *string `json:"multiple_bookings,omitempty"`
	// always empty
	MultipleBookingsCustom *float32 `json:"multiple_bookings_custom,omitempty"`
	MultipleBookingsSelect *string  `json:"multiple_bookings_select,omitempty"`
}

// NewCreateResourcesRequest instantiates a new CreateResourcesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateResourcesRequest() *CreateResourcesRequest {
	this := CreateResourcesRequest{}
	return &this
}

// NewCreateResourcesRequestWithDefaults instantiates a new CreateResourcesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateResourcesRequestWithDefaults() *CreateResourcesRequest {
	this := CreateResourcesRequest{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *CreateResourcesRequest) GetActive() float32 {
	if o == nil || o.Active == nil {
		var ret float32
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResourcesRequest) GetActiveOk() (*float32, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *CreateResourcesRequest) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given float32 and assigns it to the Active field.
func (o *CreateResourcesRequest) SetActive(v float32) {
	o.Active = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateResourcesRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResourcesRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateResourcesRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateResourcesRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *CreateResourcesRequest) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResourcesRequest) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *CreateResourcesRequest) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *CreateResourcesRequest) SetDomain(v string) {
	o.Domain = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *CreateResourcesRequest) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResourcesRequest) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *CreateResourcesRequest) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *CreateResourcesRequest) SetKind(v string) {
	o.Kind = &v
}

// GetMultipleBookings returns the MultipleBookings field value if set, zero value otherwise.
func (o *CreateResourcesRequest) GetMultipleBookings() string {
	if o == nil || o.MultipleBookings == nil {
		var ret string
		return ret
	}
	return *o.MultipleBookings
}

// GetMultipleBookingsOk returns a tuple with the MultipleBookings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResourcesRequest) GetMultipleBookingsOk() (*string, bool) {
	if o == nil || o.MultipleBookings == nil {
		return nil, false
	}
	return o.MultipleBookings, true
}

// HasMultipleBookings returns a boolean if a field has been set.
func (o *CreateResourcesRequest) HasMultipleBookings() bool {
	if o != nil && o.MultipleBookings != nil {
		return true
	}

	return false
}

// SetMultipleBookings gets a reference to the given string and assigns it to the MultipleBookings field.
func (o *CreateResourcesRequest) SetMultipleBookings(v string) {
	o.MultipleBookings = &v
}

// GetMultipleBookingsCustom returns the MultipleBookingsCustom field value if set, zero value otherwise.
func (o *CreateResourcesRequest) GetMultipleBookingsCustom() float32 {
	if o == nil || o.MultipleBookingsCustom == nil {
		var ret float32
		return ret
	}
	return *o.MultipleBookingsCustom
}

// GetMultipleBookingsCustomOk returns a tuple with the MultipleBookingsCustom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResourcesRequest) GetMultipleBookingsCustomOk() (*float32, bool) {
	if o == nil || o.MultipleBookingsCustom == nil {
		return nil, false
	}
	return o.MultipleBookingsCustom, true
}

// HasMultipleBookingsCustom returns a boolean if a field has been set.
func (o *CreateResourcesRequest) HasMultipleBookingsCustom() bool {
	if o != nil && o.MultipleBookingsCustom != nil {
		return true
	}

	return false
}

// SetMultipleBookingsCustom gets a reference to the given float32 and assigns it to the MultipleBookingsCustom field.
func (o *CreateResourcesRequest) SetMultipleBookingsCustom(v float32) {
	o.MultipleBookingsCustom = &v
}

// GetMultipleBookingsSelect returns the MultipleBookingsSelect field value if set, zero value otherwise.
func (o *CreateResourcesRequest) GetMultipleBookingsSelect() string {
	if o == nil || o.MultipleBookingsSelect == nil {
		var ret string
		return ret
	}
	return *o.MultipleBookingsSelect
}

// GetMultipleBookingsSelectOk returns a tuple with the MultipleBookingsSelect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResourcesRequest) GetMultipleBookingsSelectOk() (*string, bool) {
	if o == nil || o.MultipleBookingsSelect == nil {
		return nil, false
	}
	return o.MultipleBookingsSelect, true
}

// HasMultipleBookingsSelect returns a boolean if a field has been set.
func (o *CreateResourcesRequest) HasMultipleBookingsSelect() bool {
	if o != nil && o.MultipleBookingsSelect != nil {
		return true
	}

	return false
}

// SetMultipleBookingsSelect gets a reference to the given string and assigns it to the MultipleBookingsSelect field.
func (o *CreateResourcesRequest) SetMultipleBookingsSelect(v string) {
	o.MultipleBookingsSelect = &v
}

func (o CreateResourcesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.MultipleBookings != nil {
		toSerialize["multiple_bookings"] = o.MultipleBookings
	}
	if o.MultipleBookingsCustom != nil {
		toSerialize["multiple_bookings_custom"] = o.MultipleBookingsCustom
	}
	if o.MultipleBookingsSelect != nil {
		toSerialize["multiple_bookings_select"] = o.MultipleBookingsSelect
	}
	return json.Marshal(toSerialize)
}

type NullableCreateResourcesRequest struct {
	value *CreateResourcesRequest
	isSet bool
}

func (v NullableCreateResourcesRequest) Get() *CreateResourcesRequest {
	return v.value
}

func (v *NullableCreateResourcesRequest) Set(val *CreateResourcesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateResourcesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateResourcesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateResourcesRequest(val *CreateResourcesRequest) *NullableCreateResourcesRequest {
	return &NullableCreateResourcesRequest{value: val, isSet: true}
}

func (v NullableCreateResourcesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateResourcesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
