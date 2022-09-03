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

// OpenMetadataServerConnection struct for OpenMetadataServerConnection
type OpenMetadataServerConnection struct {
	ApiVersion                 *string                `json:"apiVersion,omitempty"`
	AuthProvider               *string                `json:"authProvider,omitempty"`
	EnableVersionValidation    *bool                  `json:"enableVersionValidation,omitempty"`
	HostPort                   string                 `json:"hostPort"`
	IncludeDashboards          *bool                  `json:"includeDashboards,omitempty"`
	IncludeDatabaseServices    *bool                  `json:"includeDatabaseServices,omitempty"`
	IncludeGlossaryTerms       *bool                  `json:"includeGlossaryTerms,omitempty"`
	IncludeMessagingServices   *bool                  `json:"includeMessagingServices,omitempty"`
	IncludeMlModels            *bool                  `json:"includeMlModels,omitempty"`
	IncludePipelineServices    *bool                  `json:"includePipelineServices,omitempty"`
	IncludePipelines           *bool                  `json:"includePipelines,omitempty"`
	IncludePolicy              *bool                  `json:"includePolicy,omitempty"`
	IncludeTables              *bool                  `json:"includeTables,omitempty"`
	IncludeTags                *bool                  `json:"includeTags,omitempty"`
	IncludeTeams               *bool                  `json:"includeTeams,omitempty"`
	IncludeTopics              *bool                  `json:"includeTopics,omitempty"`
	IncludeUsers               *bool                  `json:"includeUsers,omitempty"`
	LimitRecords               *int32                 `json:"limitRecords,omitempty"`
	SecurityConfig             map[string]interface{} `json:"securityConfig,omitempty"`
	SupportsMetadataExtraction *bool                  `json:"supportsMetadataExtraction,omitempty"`
	Type                       *string                `json:"type,omitempty"`
}

// NewOpenMetadataServerConnection instantiates a new OpenMetadataServerConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenMetadataServerConnection(hostPort string) *OpenMetadataServerConnection {
	this := OpenMetadataServerConnection{}
	this.HostPort = hostPort
	return &this
}

// NewOpenMetadataServerConnectionWithDefaults instantiates a new OpenMetadataServerConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenMetadataServerConnectionWithDefaults() *OpenMetadataServerConnection {
	this := OpenMetadataServerConnection{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *OpenMetadataServerConnection) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenMetadataServerConnection) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *OpenMetadataServerConnection) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *OpenMetadataServerConnection) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetAuthProvider returns the AuthProvider field value if set, zero value otherwise.
func (o *OpenMetadataServerConnection) GetAuthProvider() string {
	if o == nil || o.AuthProvider == nil {
		var ret string
		return ret
	}
	return *o.AuthProvider
}

// GetAuthProviderOk returns a tuple with the AuthProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenMetadataServerConnection) GetAuthProviderOk() (*string, bool) {
	if o == nil || o.AuthProvider == nil {
		return nil, false
	}
	return o.AuthProvider, true
}

// HasAuthProvider returns a boolean if a field has been set.
func (o *OpenMetadataServerConnection) HasAuthProvider() bool {
	if o != nil && o.AuthProvider != nil {
		return true
	}

	return false
}

// SetAuthProvider gets a reference to the given string and assigns it to the AuthProvider field.
func (o *OpenMetadataServerConnection) SetAuthProvider(v string) {
	o.AuthProvider = &v
}

// GetEnableVersionValidation returns the EnableVersionValidation field value if set, zero value otherwise.
func (o *OpenMetadataServerConnection) GetEnableVersionValidation() bool {
	if o == nil || o.EnableVersionValidation == nil {
		var ret bool
		return ret
	}
	return *o.EnableVersionValidation
}

// GetEnableVersionValidationOk returns a tuple with the EnableVersionValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenMetadataServerConnection) GetEnableVersionValidationOk() (*bool, bool) {
	if o == nil || o.EnableVersionValidation == nil {
		return nil, false
	}
	return o.EnableVersionValidation, true
}

// HasEnableVersionValidation returns a boolean if a field has been set.
func (o *OpenMetadataServerConnection) HasEnableVersionValidation() bool {
	if o != nil && o.EnableVersionValidation != nil {
		return true
	}

	return false
}

// SetEnableVersionValidation gets a reference to the given bool and assigns it to the EnableVersionValidation field.
func (o *OpenMetadataServerConnection) SetEnableVersionValidation(v bool) {
	o.EnableVersionValidation = &v
}

// GetHostPort returns the HostPort field value
func (o *OpenMetadataServerConnection) GetHostPort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostPort
}

// GetHostPortOk returns a tuple with the HostPort field value
// and a boolean to check if the value has been set.
func (o *OpenMetadataServerConnection) GetHostPortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostPort, true
}

// SetHostPort sets field value
func (o *OpenMetadataServerConnection) SetHostPort(v string) {
	o.HostPort = v
}

// GetIncludeDashboards returns the IncludeDashboards field value if set, zero value otherwise.
func (o *OpenMetadataServerConnection) GetIncludeDashboards() bool {
	if o == nil || o.IncludeDashboards == nil {
		var ret bool
		return ret
	}
	return *o.IncludeDashboards
}

// GetIncludeDashboardsOk returns a tuple with the IncludeDashboards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenMetadataServerConnection) GetIncludeDashboardsOk() (*bool, bool) {
	if o == nil || o.IncludeDashboards == nil {
		return nil, false
	}
	return o.IncludeDashboards, true
}

// HasIncludeDashboards returns a boolean if a field has been set.
func (o *OpenMetadataServerConnection) HasIncludeDashboards() bool {
	if o != nil && o.IncludeDashboards != nil {
		return true
	}

	return false
}

// SetIncludeDashboards gets a reference to the given bool and assigns it to the IncludeDashboards field.
func (o *OpenMetadataServerConnection) SetIncludeDashboards(v bool) {
	o.IncludeDashboards = &v
}

// GetIncludeDatabaseServices returns the IncludeDatabaseServices field value if set, zero value otherwise.
func (o *OpenMetadataServerConnection) GetIncludeDatabaseServices() bool {
	if o == nil || o.IncludeDatabaseServices == nil {
		var ret bool
		return ret
	}
	return *o.IncludeDatabaseServices
}

// GetIncludeDatabaseServicesOk returns a tuple with the IncludeDatabaseServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenMetadataServerConnection) GetIncludeDatabaseServicesOk() (*bool, bool) {
	if o == nil || o.IncludeDatabaseServices == nil {
		return nil, false
	}
	return o.IncludeDatabaseServices, true
}

// HasIncludeDatabaseServices returns a boolean if a field has been set.
func (o *OpenMetadataServerConnection) HasIncludeDatabaseServices() bool {
	if o != nil && o.IncludeDatabaseServices != nil {
		return true
	}

	return false
}

// SetIncludeDatabaseServices gets a reference to the given bool and assigns it to the IncludeDatabaseServices field.
func (o *OpenMetadataServerConnection) SetIncludeDatabaseServices(v bool) {
	o.IncludeDatabaseServices = &v
}

// GetIncludeGlossaryTerms returns the IncludeGlossaryTerms field value if set, zero value otherwise.
func (o *OpenMetadataServerConnection) GetIncludeGlossaryTerms() bool {
	if o == nil || o.IncludeGlossaryTerms == nil {
		var ret bool
		return ret
	}
	return *o.IncludeGlossaryTerms
}

// GetIncludeGlossaryTermsOk returns a tuple with the IncludeGlossaryTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenMetadataServerConnection) GetIncludeGlossaryTermsOk() (*bool, bool) {
	if o == nil || o.IncludeGlossaryTerms == nil {
		return nil, false
	}
	return o.IncludeGlossaryTerms, true
}

// HasIncludeGlossaryTerms returns a boolean if a field has been set.
func (o *OpenMetadataServerConnection) HasIncludeGlossaryTerms() bool {
	if o != nil && o.IncludeGlossaryTerms != nil {
		return true
	}

	return false
}

// SetIncludeGlossaryTerms gets a reference to the given bool and assigns it to the IncludeGlossaryTerms field.
func (o *OpenMetadataServerConnection) SetIncludeGlossaryTerms(v bool) {
	o.IncludeGlossaryTerms = &v
}

// GetIncludeMessagingServices returns the IncludeMessagingServices field value if set, zero value otherwise.
func (o *OpenMetadataServerConnection) GetIncludeMessagingServices() bool {
	if o == nil || o.IncludeMessagingServices == nil {
		var ret bool
		return ret
	}
	return *o.IncludeMessagingServices
}

// GetIncludeMessagingServicesOk returns a tuple with the IncludeMessagingServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenMetadataServerConnection) GetIncludeMessagingServicesOk() (*bool, bool) {
	if o == nil || o.IncludeMessagingServices == nil {
		return nil, false
	}
	return o.IncludeMessagingServices, true
}

// HasIncludeMessagingServices returns a boolean if a field has been set.
func (o *OpenMetadataServerConnection) HasIncludeMessagingServices() bool {
	if o != nil && o.IncludeMessagingServices != nil {
		return true
	}

	return false
}

// SetIncludeMessagingServices gets a reference to the given bool and assigns it to the IncludeMessagingServices field.
func (o *OpenMetadataServerConnection) SetIncludeMessagingServices(v bool) {
	o.IncludeMessagingServices = &v
}

// GetIncludeMlModels returns the IncludeMlModels field value if set, zero value otherwise.
func (o *OpenMetadataServerConnection) GetIncludeMlModels() bool {
	if o == nil || o.IncludeMlModels == nil {
		var ret bool
		return ret
	}
	return *o.IncludeMlModels
}

// GetIncludeMlModelsOk returns a tuple with the IncludeMlModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenMetadataServerConnection) GetIncludeMlModelsOk() (*bool, bool) {
	if o == nil || o.IncludeMlModels == nil {
		return nil, false
	}
	return o.IncludeMlModels, true
}

// HasIncludeMlModels returns a boolean if a field has been set.
func (o *OpenMetadataServerConnection) HasIncludeMlModels() bool {
	if o != nil && o.IncludeMlModels != nil {
		return true
	}

	return false
}

// SetIncludeMlModels gets a reference to the given bool and assigns it to the IncludeMlModels field.
func (o *OpenMetadataServerConnection) SetIncludeMlModels(v bool) {
	o.IncludeMlModels = &v
}

// GetIncludePipelineServices returns the IncludePipelineServices field value if set, zero value otherwise.
func (o *OpenMetadataServerConnection) GetIncludePipelineServices() bool {
	if o == nil || o.IncludePipelineServices == nil {
		var ret bool
		return ret
	}
	return *o.IncludePipelineServices
}

// GetIncludePipelineServicesOk returns a tuple with the IncludePipelineServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenMetadataServerConnection) GetIncludePipelineServicesOk() (*bool, bool) {
	if o == nil || o.IncludePipelineServices == nil {
		return nil, false
	}
	return o.IncludePipelineServices, true
}

// HasIncludePipelineServices returns a boolean if a field has been set.
func (o *OpenMetadataServerConnection) HasIncludePipelineServices() bool {
	if o != nil && o.IncludePipelineServices != nil {
		return true
	}

	return false
}

// SetIncludePipelineServices gets a reference to the given bool and assigns it to the IncludePipelineServices field.
func (o *OpenMetadataServerConnection) SetIncludePipelineServices(v bool) {
	o.IncludePipelineServices = &v
}

// GetIncludePipelines returns the IncludePipelines field value if set, zero value otherwise.
func (o *OpenMetadataServerConnection) GetIncludePipelines() bool {
	if o == nil || o.IncludePipelines == nil {
		var ret bool
		return ret
	}
	return *o.IncludePipelines
}

// GetIncludePipelinesOk returns a tuple with the IncludePipelines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenMetadataServerConnection) GetIncludePipelinesOk() (*bool, bool) {
	if o == nil || o.IncludePipelines == nil {
		return nil, false
	}
	return o.IncludePipelines, true
}

// HasIncludePipelines returns a boolean if a field has been set.
func (o *OpenMetadataServerConnection) HasIncludePipelines() bool {
	if o != nil && o.IncludePipelines != nil {
		return true
	}

	return false
}

// SetIncludePipelines gets a reference to the given bool and assigns it to the IncludePipelines field.
func (o *OpenMetadataServerConnection) SetIncludePipelines(v bool) {
	o.IncludePipelines = &v
}

// GetIncludePolicy returns the IncludePolicy field value if set, zero value otherwise.
func (o *OpenMetadataServerConnection) GetIncludePolicy() bool {
	if o == nil || o.IncludePolicy == nil {
		var ret bool
		return ret
	}
	return *o.IncludePolicy
}

// GetIncludePolicyOk returns a tuple with the IncludePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenMetadataServerConnection) GetIncludePolicyOk() (*bool, bool) {
	if o == nil || o.IncludePolicy == nil {
		return nil, false
	}
	return o.IncludePolicy, true
}

// HasIncludePolicy returns a boolean if a field has been set.
func (o *OpenMetadataServerConnection) HasIncludePolicy() bool {
	if o != nil && o.IncludePolicy != nil {
		return true
	}

	return false
}

// SetIncludePolicy gets a reference to the given bool and assigns it to the IncludePolicy field.
func (o *OpenMetadataServerConnection) SetIncludePolicy(v bool) {
	o.IncludePolicy = &v
}

// GetIncludeTables returns the IncludeTables field value if set, zero value otherwise.
func (o *OpenMetadataServerConnection) GetIncludeTables() bool {
	if o == nil || o.IncludeTables == nil {
		var ret bool
		return ret
	}
	return *o.IncludeTables
}

// GetIncludeTablesOk returns a tuple with the IncludeTables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenMetadataServerConnection) GetIncludeTablesOk() (*bool, bool) {
	if o == nil || o.IncludeTables == nil {
		return nil, false
	}
	return o.IncludeTables, true
}

// HasIncludeTables returns a boolean if a field has been set.
func (o *OpenMetadataServerConnection) HasIncludeTables() bool {
	if o != nil && o.IncludeTables != nil {
		return true
	}

	return false
}

// SetIncludeTables gets a reference to the given bool and assigns it to the IncludeTables field.
func (o *OpenMetadataServerConnection) SetIncludeTables(v bool) {
	o.IncludeTables = &v
}

// GetIncludeTags returns the IncludeTags field value if set, zero value otherwise.
func (o *OpenMetadataServerConnection) GetIncludeTags() bool {
	if o == nil || o.IncludeTags == nil {
		var ret bool
		return ret
	}
	return *o.IncludeTags
}

// GetIncludeTagsOk returns a tuple with the IncludeTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenMetadataServerConnection) GetIncludeTagsOk() (*bool, bool) {
	if o == nil || o.IncludeTags == nil {
		return nil, false
	}
	return o.IncludeTags, true
}

// HasIncludeTags returns a boolean if a field has been set.
func (o *OpenMetadataServerConnection) HasIncludeTags() bool {
	if o != nil && o.IncludeTags != nil {
		return true
	}

	return false
}

// SetIncludeTags gets a reference to the given bool and assigns it to the IncludeTags field.
func (o *OpenMetadataServerConnection) SetIncludeTags(v bool) {
	o.IncludeTags = &v
}

// GetIncludeTeams returns the IncludeTeams field value if set, zero value otherwise.
func (o *OpenMetadataServerConnection) GetIncludeTeams() bool {
	if o == nil || o.IncludeTeams == nil {
		var ret bool
		return ret
	}
	return *o.IncludeTeams
}

// GetIncludeTeamsOk returns a tuple with the IncludeTeams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenMetadataServerConnection) GetIncludeTeamsOk() (*bool, bool) {
	if o == nil || o.IncludeTeams == nil {
		return nil, false
	}
	return o.IncludeTeams, true
}

// HasIncludeTeams returns a boolean if a field has been set.
func (o *OpenMetadataServerConnection) HasIncludeTeams() bool {
	if o != nil && o.IncludeTeams != nil {
		return true
	}

	return false
}

// SetIncludeTeams gets a reference to the given bool and assigns it to the IncludeTeams field.
func (o *OpenMetadataServerConnection) SetIncludeTeams(v bool) {
	o.IncludeTeams = &v
}

// GetIncludeTopics returns the IncludeTopics field value if set, zero value otherwise.
func (o *OpenMetadataServerConnection) GetIncludeTopics() bool {
	if o == nil || o.IncludeTopics == nil {
		var ret bool
		return ret
	}
	return *o.IncludeTopics
}

// GetIncludeTopicsOk returns a tuple with the IncludeTopics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenMetadataServerConnection) GetIncludeTopicsOk() (*bool, bool) {
	if o == nil || o.IncludeTopics == nil {
		return nil, false
	}
	return o.IncludeTopics, true
}

// HasIncludeTopics returns a boolean if a field has been set.
func (o *OpenMetadataServerConnection) HasIncludeTopics() bool {
	if o != nil && o.IncludeTopics != nil {
		return true
	}

	return false
}

// SetIncludeTopics gets a reference to the given bool and assigns it to the IncludeTopics field.
func (o *OpenMetadataServerConnection) SetIncludeTopics(v bool) {
	o.IncludeTopics = &v
}

// GetIncludeUsers returns the IncludeUsers field value if set, zero value otherwise.
func (o *OpenMetadataServerConnection) GetIncludeUsers() bool {
	if o == nil || o.IncludeUsers == nil {
		var ret bool
		return ret
	}
	return *o.IncludeUsers
}

// GetIncludeUsersOk returns a tuple with the IncludeUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenMetadataServerConnection) GetIncludeUsersOk() (*bool, bool) {
	if o == nil || o.IncludeUsers == nil {
		return nil, false
	}
	return o.IncludeUsers, true
}

// HasIncludeUsers returns a boolean if a field has been set.
func (o *OpenMetadataServerConnection) HasIncludeUsers() bool {
	if o != nil && o.IncludeUsers != nil {
		return true
	}

	return false
}

// SetIncludeUsers gets a reference to the given bool and assigns it to the IncludeUsers field.
func (o *OpenMetadataServerConnection) SetIncludeUsers(v bool) {
	o.IncludeUsers = &v
}

// GetLimitRecords returns the LimitRecords field value if set, zero value otherwise.
func (o *OpenMetadataServerConnection) GetLimitRecords() int32 {
	if o == nil || o.LimitRecords == nil {
		var ret int32
		return ret
	}
	return *o.LimitRecords
}

// GetLimitRecordsOk returns a tuple with the LimitRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenMetadataServerConnection) GetLimitRecordsOk() (*int32, bool) {
	if o == nil || o.LimitRecords == nil {
		return nil, false
	}
	return o.LimitRecords, true
}

// HasLimitRecords returns a boolean if a field has been set.
func (o *OpenMetadataServerConnection) HasLimitRecords() bool {
	if o != nil && o.LimitRecords != nil {
		return true
	}

	return false
}

// SetLimitRecords gets a reference to the given int32 and assigns it to the LimitRecords field.
func (o *OpenMetadataServerConnection) SetLimitRecords(v int32) {
	o.LimitRecords = &v
}

// GetSecurityConfig returns the SecurityConfig field value if set, zero value otherwise.
func (o *OpenMetadataServerConnection) GetSecurityConfig() map[string]interface{} {
	if o == nil || o.SecurityConfig == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.SecurityConfig
}

// GetSecurityConfigOk returns a tuple with the SecurityConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenMetadataServerConnection) GetSecurityConfigOk() (map[string]interface{}, bool) {
	if o == nil || o.SecurityConfig == nil {
		return nil, false
	}
	return o.SecurityConfig, true
}

// HasSecurityConfig returns a boolean if a field has been set.
func (o *OpenMetadataServerConnection) HasSecurityConfig() bool {
	if o != nil && o.SecurityConfig != nil {
		return true
	}

	return false
}

// SetSecurityConfig gets a reference to the given map[string]interface{} and assigns it to the SecurityConfig field.
func (o *OpenMetadataServerConnection) SetSecurityConfig(v map[string]interface{}) {
	o.SecurityConfig = v
}

// GetSupportsMetadataExtraction returns the SupportsMetadataExtraction field value if set, zero value otherwise.
func (o *OpenMetadataServerConnection) GetSupportsMetadataExtraction() bool {
	if o == nil || o.SupportsMetadataExtraction == nil {
		var ret bool
		return ret
	}
	return *o.SupportsMetadataExtraction
}

// GetSupportsMetadataExtractionOk returns a tuple with the SupportsMetadataExtraction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenMetadataServerConnection) GetSupportsMetadataExtractionOk() (*bool, bool) {
	if o == nil || o.SupportsMetadataExtraction == nil {
		return nil, false
	}
	return o.SupportsMetadataExtraction, true
}

// HasSupportsMetadataExtraction returns a boolean if a field has been set.
func (o *OpenMetadataServerConnection) HasSupportsMetadataExtraction() bool {
	if o != nil && o.SupportsMetadataExtraction != nil {
		return true
	}

	return false
}

// SetSupportsMetadataExtraction gets a reference to the given bool and assigns it to the SupportsMetadataExtraction field.
func (o *OpenMetadataServerConnection) SetSupportsMetadataExtraction(v bool) {
	o.SupportsMetadataExtraction = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OpenMetadataServerConnection) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenMetadataServerConnection) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OpenMetadataServerConnection) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OpenMetadataServerConnection) SetType(v string) {
	o.Type = &v
}

func (o OpenMetadataServerConnection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if o.AuthProvider != nil {
		toSerialize["authProvider"] = o.AuthProvider
	}
	if o.EnableVersionValidation != nil {
		toSerialize["enableVersionValidation"] = o.EnableVersionValidation
	}
	if true {
		toSerialize["hostPort"] = o.HostPort
	}
	if o.IncludeDashboards != nil {
		toSerialize["includeDashboards"] = o.IncludeDashboards
	}
	if o.IncludeDatabaseServices != nil {
		toSerialize["includeDatabaseServices"] = o.IncludeDatabaseServices
	}
	if o.IncludeGlossaryTerms != nil {
		toSerialize["includeGlossaryTerms"] = o.IncludeGlossaryTerms
	}
	if o.IncludeMessagingServices != nil {
		toSerialize["includeMessagingServices"] = o.IncludeMessagingServices
	}
	if o.IncludeMlModels != nil {
		toSerialize["includeMlModels"] = o.IncludeMlModels
	}
	if o.IncludePipelineServices != nil {
		toSerialize["includePipelineServices"] = o.IncludePipelineServices
	}
	if o.IncludePipelines != nil {
		toSerialize["includePipelines"] = o.IncludePipelines
	}
	if o.IncludePolicy != nil {
		toSerialize["includePolicy"] = o.IncludePolicy
	}
	if o.IncludeTables != nil {
		toSerialize["includeTables"] = o.IncludeTables
	}
	if o.IncludeTags != nil {
		toSerialize["includeTags"] = o.IncludeTags
	}
	if o.IncludeTeams != nil {
		toSerialize["includeTeams"] = o.IncludeTeams
	}
	if o.IncludeTopics != nil {
		toSerialize["includeTopics"] = o.IncludeTopics
	}
	if o.IncludeUsers != nil {
		toSerialize["includeUsers"] = o.IncludeUsers
	}
	if o.LimitRecords != nil {
		toSerialize["limitRecords"] = o.LimitRecords
	}
	if o.SecurityConfig != nil {
		toSerialize["securityConfig"] = o.SecurityConfig
	}
	if o.SupportsMetadataExtraction != nil {
		toSerialize["supportsMetadataExtraction"] = o.SupportsMetadataExtraction
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableOpenMetadataServerConnection struct {
	value *OpenMetadataServerConnection
	isSet bool
}

func (v NullableOpenMetadataServerConnection) Get() *OpenMetadataServerConnection {
	return v.value
}

func (v *NullableOpenMetadataServerConnection) Set(val *OpenMetadataServerConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenMetadataServerConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenMetadataServerConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenMetadataServerConnection(val *OpenMetadataServerConnection) *NullableOpenMetadataServerConnection {
	return &NullableOpenMetadataServerConnection{value: val, isSet: true}
}

func (v NullableOpenMetadataServerConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenMetadataServerConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}