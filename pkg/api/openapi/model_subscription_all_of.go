/*
rh-trex Service API

rh-trex Service API

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SubscriptionAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionAllOf{}

// SubscriptionAllOf struct for SubscriptionAllOf
type SubscriptionAllOf struct {
	Foo *string `json:"foo,omitempty"`
}

// NewSubscriptionAllOf instantiates a new SubscriptionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionAllOf() *SubscriptionAllOf {
	this := SubscriptionAllOf{}
	return &this
}

// NewSubscriptionAllOfWithDefaults instantiates a new SubscriptionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionAllOfWithDefaults() *SubscriptionAllOf {
	this := SubscriptionAllOf{}
	return &this
}

// GetFoo returns the Foo field value if set, zero value otherwise.
func (o *SubscriptionAllOf) GetFoo() string {
	if o == nil || IsNil(o.Foo) {
		var ret string
		return ret
	}
	return *o.Foo
}

// GetFooOk returns a tuple with the Foo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAllOf) GetFooOk() (*string, bool) {
	if o == nil || IsNil(o.Foo) {
		return nil, false
	}
	return o.Foo, true
}

// HasFoo returns a boolean if a field has been set.
func (o *SubscriptionAllOf) HasFoo() bool {
	if o != nil && !IsNil(o.Foo) {
		return true
	}

	return false
}

// SetFoo gets a reference to the given string and assigns it to the Foo field.
func (o *SubscriptionAllOf) SetFoo(v string) {
	o.Foo = &v
}

func (o SubscriptionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Foo) {
		toSerialize["foo"] = o.Foo
	}
	return toSerialize, nil
}

type NullableSubscriptionAllOf struct {
	value *SubscriptionAllOf
	isSet bool
}

func (v NullableSubscriptionAllOf) Get() *SubscriptionAllOf {
	return v.value
}

func (v *NullableSubscriptionAllOf) Set(val *SubscriptionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionAllOf(val *SubscriptionAllOf) *NullableSubscriptionAllOf {
	return &NullableSubscriptionAllOf{value: val, isSet: true}
}

func (v NullableSubscriptionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
