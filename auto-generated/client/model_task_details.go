/*
OpenMetadata Apis

# Overview  OpenMetadata supports REST APIs for getting metadata in and out of metadata store. The API resources are grouped under following categories: - **Data assets** - includes resources for data entities, such as `databases`, `tables`, and `topics`. Resources for data assets created from data, such as `dashboards`, `reports`, `metrics`, and `ML Features` are part of this collection. `pipelines`, `notebooks`, etc. that are used for creating data assets are also available as resources as of this collection. - **Teams and Users** - includes `users`, `teams`, a special type of user called `bots` that performs many automated tasks such as ingestion. - **Services** - are services that OpenMetadata integrates with. Currently `databaseService` is the only service under this collection that represents data sources. In the future, services related to Dashboards, Reports, ETL pipelines will be added under this collection. - **Glossary** - OpenMetadata supports hierarchical tags that can be used to build business vocabulary to describe and classify data available under `tags` resource.

API version: 0.11.0
Contact: openmetadata-dev@googlegroups.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TaskDetails struct for TaskDetails
type TaskDetails struct {
	Assignees  []EntityReference `json:"assignees"`
	ClosedAt   *int64            `json:"closedAt,omitempty"`
	ClosedBy   *string           `json:"closedBy,omitempty"`
	Id         int32             `json:"id"`
	NewValue   *string           `json:"newValue,omitempty"`
	OldValue   *string           `json:"oldValue,omitempty"`
	Status     *string           `json:"status,omitempty"`
	Suggestion *string           `json:"suggestion,omitempty"`
	Type       string            `json:"type"`
}

// NewTaskDetails instantiates a new TaskDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskDetails(assignees []EntityReference, id int32, type_ string) *TaskDetails {
	this := TaskDetails{}
	this.Assignees = assignees
	this.Id = id
	this.Type = type_
	return &this
}

// NewTaskDetailsWithDefaults instantiates a new TaskDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskDetailsWithDefaults() *TaskDetails {
	this := TaskDetails{}
	return &this
}

// GetAssignees returns the Assignees field value
func (o *TaskDetails) GetAssignees() []EntityReference {
	if o == nil {
		var ret []EntityReference
		return ret
	}

	return o.Assignees
}

// GetAssigneesOk returns a tuple with the Assignees field value
// and a boolean to check if the value has been set.
func (o *TaskDetails) GetAssigneesOk() ([]EntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.Assignees, true
}

// SetAssignees sets field value
func (o *TaskDetails) SetAssignees(v []EntityReference) {
	o.Assignees = v
}

// GetClosedAt returns the ClosedAt field value if set, zero value otherwise.
func (o *TaskDetails) GetClosedAt() int64 {
	if o == nil || o.ClosedAt == nil {
		var ret int64
		return ret
	}
	return *o.ClosedAt
}

// GetClosedAtOk returns a tuple with the ClosedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDetails) GetClosedAtOk() (*int64, bool) {
	if o == nil || o.ClosedAt == nil {
		return nil, false
	}
	return o.ClosedAt, true
}

// HasClosedAt returns a boolean if a field has been set.
func (o *TaskDetails) HasClosedAt() bool {
	if o != nil && o.ClosedAt != nil {
		return true
	}

	return false
}

// SetClosedAt gets a reference to the given int64 and assigns it to the ClosedAt field.
func (o *TaskDetails) SetClosedAt(v int64) {
	o.ClosedAt = &v
}

// GetClosedBy returns the ClosedBy field value if set, zero value otherwise.
func (o *TaskDetails) GetClosedBy() string {
	if o == nil || o.ClosedBy == nil {
		var ret string
		return ret
	}
	return *o.ClosedBy
}

// GetClosedByOk returns a tuple with the ClosedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDetails) GetClosedByOk() (*string, bool) {
	if o == nil || o.ClosedBy == nil {
		return nil, false
	}
	return o.ClosedBy, true
}

// HasClosedBy returns a boolean if a field has been set.
func (o *TaskDetails) HasClosedBy() bool {
	if o != nil && o.ClosedBy != nil {
		return true
	}

	return false
}

// SetClosedBy gets a reference to the given string and assigns it to the ClosedBy field.
func (o *TaskDetails) SetClosedBy(v string) {
	o.ClosedBy = &v
}

// GetId returns the Id field value
func (o *TaskDetails) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TaskDetails) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TaskDetails) SetId(v int32) {
	o.Id = v
}

// GetNewValue returns the NewValue field value if set, zero value otherwise.
func (o *TaskDetails) GetNewValue() string {
	if o == nil || o.NewValue == nil {
		var ret string
		return ret
	}
	return *o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDetails) GetNewValueOk() (*string, bool) {
	if o == nil || o.NewValue == nil {
		return nil, false
	}
	return o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *TaskDetails) HasNewValue() bool {
	if o != nil && o.NewValue != nil {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given string and assigns it to the NewValue field.
func (o *TaskDetails) SetNewValue(v string) {
	o.NewValue = &v
}

// GetOldValue returns the OldValue field value if set, zero value otherwise.
func (o *TaskDetails) GetOldValue() string {
	if o == nil || o.OldValue == nil {
		var ret string
		return ret
	}
	return *o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDetails) GetOldValueOk() (*string, bool) {
	if o == nil || o.OldValue == nil {
		return nil, false
	}
	return o.OldValue, true
}

// HasOldValue returns a boolean if a field has been set.
func (o *TaskDetails) HasOldValue() bool {
	if o != nil && o.OldValue != nil {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given string and assigns it to the OldValue field.
func (o *TaskDetails) SetOldValue(v string) {
	o.OldValue = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TaskDetails) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDetails) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TaskDetails) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TaskDetails) SetStatus(v string) {
	o.Status = &v
}

// GetSuggestion returns the Suggestion field value if set, zero value otherwise.
func (o *TaskDetails) GetSuggestion() string {
	if o == nil || o.Suggestion == nil {
		var ret string
		return ret
	}
	return *o.Suggestion
}

// GetSuggestionOk returns a tuple with the Suggestion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDetails) GetSuggestionOk() (*string, bool) {
	if o == nil || o.Suggestion == nil {
		return nil, false
	}
	return o.Suggestion, true
}

// HasSuggestion returns a boolean if a field has been set.
func (o *TaskDetails) HasSuggestion() bool {
	if o != nil && o.Suggestion != nil {
		return true
	}

	return false
}

// SetSuggestion gets a reference to the given string and assigns it to the Suggestion field.
func (o *TaskDetails) SetSuggestion(v string) {
	o.Suggestion = &v
}

// GetType returns the Type field value
func (o *TaskDetails) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TaskDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TaskDetails) SetType(v string) {
	o.Type = v
}

func (o TaskDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["assignees"] = o.Assignees
	}
	if o.ClosedAt != nil {
		toSerialize["closedAt"] = o.ClosedAt
	}
	if o.ClosedBy != nil {
		toSerialize["closedBy"] = o.ClosedBy
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.NewValue != nil {
		toSerialize["newValue"] = o.NewValue
	}
	if o.OldValue != nil {
		toSerialize["oldValue"] = o.OldValue
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Suggestion != nil {
		toSerialize["suggestion"] = o.Suggestion
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTaskDetails struct {
	value *TaskDetails
	isSet bool
}

func (v NullableTaskDetails) Get() *TaskDetails {
	return v.value
}

func (v *NullableTaskDetails) Set(val *TaskDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskDetails(val *TaskDetails) *NullableTaskDetails {
	return &NullableTaskDetails{value: val, isSet: true}
}

func (v NullableTaskDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}