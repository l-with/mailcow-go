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

// UpdateSyncJobRequest struct for UpdateSyncJobRequest
type UpdateSyncJobRequest struct {
	Attr *UpdateSyncJobRequestAttr `json:"attr,omitempty"`
	// contains list of aliases you want update
	Items []string `json:"items,omitempty"`
}

// NewUpdateSyncJobRequest instantiates a new UpdateSyncJobRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSyncJobRequest() *UpdateSyncJobRequest {
	this := UpdateSyncJobRequest{}
	return &this
}

// NewUpdateSyncJobRequestWithDefaults instantiates a new UpdateSyncJobRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSyncJobRequestWithDefaults() *UpdateSyncJobRequest {
	this := UpdateSyncJobRequest{}
	return &this
}

// GetAttr returns the Attr field value if set, zero value otherwise.
func (o *UpdateSyncJobRequest) GetAttr() UpdateSyncJobRequestAttr {
	if o == nil || o.Attr == nil {
		var ret UpdateSyncJobRequestAttr
		return ret
	}
	return *o.Attr
}

// GetAttrOk returns a tuple with the Attr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSyncJobRequest) GetAttrOk() (*UpdateSyncJobRequestAttr, bool) {
	if o == nil || o.Attr == nil {
		return nil, false
	}
	return o.Attr, true
}

// HasAttr returns a boolean if a field has been set.
func (o *UpdateSyncJobRequest) HasAttr() bool {
	if o != nil && o.Attr != nil {
		return true
	}

	return false
}

// SetAttr gets a reference to the given UpdateSyncJobRequestAttr and assigns it to the Attr field.
func (o *UpdateSyncJobRequest) SetAttr(v UpdateSyncJobRequestAttr) {
	o.Attr = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *UpdateSyncJobRequest) GetItems() []string {
	if o == nil || o.Items == nil {
		var ret []string
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSyncJobRequest) GetItemsOk() ([]string, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *UpdateSyncJobRequest) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []string and assigns it to the Items field.
func (o *UpdateSyncJobRequest) SetItems(v []string) {
	o.Items = v
}

func (o UpdateSyncJobRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attr != nil {
		toSerialize["attr"] = o.Attr
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSyncJobRequest struct {
	value *UpdateSyncJobRequest
	isSet bool
}

func (v NullableUpdateSyncJobRequest) Get() *UpdateSyncJobRequest {
	return v.value
}

func (v *NullableUpdateSyncJobRequest) Set(val *UpdateSyncJobRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSyncJobRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSyncJobRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSyncJobRequest(val *UpdateSyncJobRequest) *NullableUpdateSyncJobRequest {
	return &NullableUpdateSyncJobRequest{value: val, isSet: true}
}

func (v NullableUpdateSyncJobRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSyncJobRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
