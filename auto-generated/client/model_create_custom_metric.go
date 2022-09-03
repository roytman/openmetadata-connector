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

// CreateCustomMetric struct for CreateCustomMetric
type CreateCustomMetric struct {
	ColumnName  string           `json:"columnName"`
	Description *string          `json:"description,omitempty"`
	Expression  string           `json:"expression"`
	Name        string           `json:"name"`
	Owner       *EntityReference `json:"owner,omitempty"`
	UpdatedAt   *int64           `json:"updatedAt,omitempty"`
	UpdatedBy   *string          `json:"updatedBy,omitempty"`
}

// NewCreateCustomMetric instantiates a new CreateCustomMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCustomMetric(columnName string, expression string, name string) *CreateCustomMetric {
	this := CreateCustomMetric{}
	this.ColumnName = columnName
	this.Expression = expression
	this.Name = name
	return &this
}

// NewCreateCustomMetricWithDefaults instantiates a new CreateCustomMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCustomMetricWithDefaults() *CreateCustomMetric {
	this := CreateCustomMetric{}
	return &this
}

// GetColumnName returns the ColumnName field value
func (o *CreateCustomMetric) GetColumnName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value
// and a boolean to check if the value has been set.
func (o *CreateCustomMetric) GetColumnNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColumnName, true
}

// SetColumnName sets field value
func (o *CreateCustomMetric) SetColumnName(v string) {
	o.ColumnName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateCustomMetric) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomMetric) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateCustomMetric) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateCustomMetric) SetDescription(v string) {
	o.Description = &v
}

// GetExpression returns the Expression field value
func (o *CreateCustomMetric) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *CreateCustomMetric) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *CreateCustomMetric) SetExpression(v string) {
	o.Expression = v
}

// GetName returns the Name field value
func (o *CreateCustomMetric) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateCustomMetric) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateCustomMetric) SetName(v string) {
	o.Name = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *CreateCustomMetric) GetOwner() EntityReference {
	if o == nil || o.Owner == nil {
		var ret EntityReference
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomMetric) GetOwnerOk() (*EntityReference, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *CreateCustomMetric) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given EntityReference and assigns it to the Owner field.
func (o *CreateCustomMetric) SetOwner(v EntityReference) {
	o.Owner = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CreateCustomMetric) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomMetric) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CreateCustomMetric) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *CreateCustomMetric) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *CreateCustomMetric) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomMetric) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *CreateCustomMetric) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *CreateCustomMetric) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

func (o CreateCustomMetric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["columnName"] = o.ColumnName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["expression"] = o.Expression
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.UpdatedBy != nil {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCustomMetric struct {
	value *CreateCustomMetric
	isSet bool
}

func (v NullableCreateCustomMetric) Get() *CreateCustomMetric {
	return v.value
}

func (v *NullableCreateCustomMetric) Set(val *CreateCustomMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCustomMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCustomMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCustomMetric(val *CreateCustomMetric) *NullableCreateCustomMetric {
	return &NullableCreateCustomMetric{value: val, isSet: true}
}

func (v NullableCreateCustomMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCustomMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}