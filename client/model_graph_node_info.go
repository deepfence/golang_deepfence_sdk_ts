/*
Deepfence ThreatStryker

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.5.6
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GraphNodeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphNodeInfo{}

// GraphNodeInfo struct for GraphNodeInfo
type GraphNodeInfo struct {
	AlertsCount int32 `json:"alerts_count"`
	CloudComplianceCount int32 `json:"cloud_compliance_count"`
	CloudWarnAlarmCount int32 `json:"cloud_warn_alarm_count"`
	ComplianceCount int32 `json:"compliance_count"`
	ExploitableSecretsCount int32 `json:"exploitable_secrets_count"`
	ExploitableVulnerabilitiesCount int32 `json:"exploitable_vulnerabilities_count"`
	Name string `json:"name"`
	NodeId string `json:"node_id"`
	SecretsCount int32 `json:"secrets_count"`
	VulnerabilityCount int32 `json:"vulnerability_count"`
	WarnAlarmCount int32 `json:"warn_alarm_count"`
}

type _GraphNodeInfo GraphNodeInfo

// NewGraphNodeInfo instantiates a new GraphNodeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphNodeInfo(alertsCount int32, cloudComplianceCount int32, cloudWarnAlarmCount int32, complianceCount int32, exploitableSecretsCount int32, exploitableVulnerabilitiesCount int32, name string, nodeId string, secretsCount int32, vulnerabilityCount int32, warnAlarmCount int32) *GraphNodeInfo {
	this := GraphNodeInfo{}
	this.AlertsCount = alertsCount
	this.CloudComplianceCount = cloudComplianceCount
	this.CloudWarnAlarmCount = cloudWarnAlarmCount
	this.ComplianceCount = complianceCount
	this.ExploitableSecretsCount = exploitableSecretsCount
	this.ExploitableVulnerabilitiesCount = exploitableVulnerabilitiesCount
	this.Name = name
	this.NodeId = nodeId
	this.SecretsCount = secretsCount
	this.VulnerabilityCount = vulnerabilityCount
	this.WarnAlarmCount = warnAlarmCount
	return &this
}

// NewGraphNodeInfoWithDefaults instantiates a new GraphNodeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphNodeInfoWithDefaults() *GraphNodeInfo {
	this := GraphNodeInfo{}
	return &this
}

// GetAlertsCount returns the AlertsCount field value
func (o *GraphNodeInfo) GetAlertsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AlertsCount
}

// GetAlertsCountOk returns a tuple with the AlertsCount field value
// and a boolean to check if the value has been set.
func (o *GraphNodeInfo) GetAlertsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertsCount, true
}

// SetAlertsCount sets field value
func (o *GraphNodeInfo) SetAlertsCount(v int32) {
	o.AlertsCount = v
}

// GetCloudComplianceCount returns the CloudComplianceCount field value
func (o *GraphNodeInfo) GetCloudComplianceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CloudComplianceCount
}

// GetCloudComplianceCountOk returns a tuple with the CloudComplianceCount field value
// and a boolean to check if the value has been set.
func (o *GraphNodeInfo) GetCloudComplianceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudComplianceCount, true
}

// SetCloudComplianceCount sets field value
func (o *GraphNodeInfo) SetCloudComplianceCount(v int32) {
	o.CloudComplianceCount = v
}

// GetCloudWarnAlarmCount returns the CloudWarnAlarmCount field value
func (o *GraphNodeInfo) GetCloudWarnAlarmCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CloudWarnAlarmCount
}

// GetCloudWarnAlarmCountOk returns a tuple with the CloudWarnAlarmCount field value
// and a boolean to check if the value has been set.
func (o *GraphNodeInfo) GetCloudWarnAlarmCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudWarnAlarmCount, true
}

// SetCloudWarnAlarmCount sets field value
func (o *GraphNodeInfo) SetCloudWarnAlarmCount(v int32) {
	o.CloudWarnAlarmCount = v
}

// GetComplianceCount returns the ComplianceCount field value
func (o *GraphNodeInfo) GetComplianceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ComplianceCount
}

// GetComplianceCountOk returns a tuple with the ComplianceCount field value
// and a boolean to check if the value has been set.
func (o *GraphNodeInfo) GetComplianceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComplianceCount, true
}

// SetComplianceCount sets field value
func (o *GraphNodeInfo) SetComplianceCount(v int32) {
	o.ComplianceCount = v
}

// GetExploitableSecretsCount returns the ExploitableSecretsCount field value
func (o *GraphNodeInfo) GetExploitableSecretsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExploitableSecretsCount
}

// GetExploitableSecretsCountOk returns a tuple with the ExploitableSecretsCount field value
// and a boolean to check if the value has been set.
func (o *GraphNodeInfo) GetExploitableSecretsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExploitableSecretsCount, true
}

// SetExploitableSecretsCount sets field value
func (o *GraphNodeInfo) SetExploitableSecretsCount(v int32) {
	o.ExploitableSecretsCount = v
}

// GetExploitableVulnerabilitiesCount returns the ExploitableVulnerabilitiesCount field value
func (o *GraphNodeInfo) GetExploitableVulnerabilitiesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExploitableVulnerabilitiesCount
}

// GetExploitableVulnerabilitiesCountOk returns a tuple with the ExploitableVulnerabilitiesCount field value
// and a boolean to check if the value has been set.
func (o *GraphNodeInfo) GetExploitableVulnerabilitiesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExploitableVulnerabilitiesCount, true
}

// SetExploitableVulnerabilitiesCount sets field value
func (o *GraphNodeInfo) SetExploitableVulnerabilitiesCount(v int32) {
	o.ExploitableVulnerabilitiesCount = v
}

// GetName returns the Name field value
func (o *GraphNodeInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GraphNodeInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GraphNodeInfo) SetName(v string) {
	o.Name = v
}

// GetNodeId returns the NodeId field value
func (o *GraphNodeInfo) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *GraphNodeInfo) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *GraphNodeInfo) SetNodeId(v string) {
	o.NodeId = v
}

// GetSecretsCount returns the SecretsCount field value
func (o *GraphNodeInfo) GetSecretsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SecretsCount
}

// GetSecretsCountOk returns a tuple with the SecretsCount field value
// and a boolean to check if the value has been set.
func (o *GraphNodeInfo) GetSecretsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretsCount, true
}

// SetSecretsCount sets field value
func (o *GraphNodeInfo) SetSecretsCount(v int32) {
	o.SecretsCount = v
}

// GetVulnerabilityCount returns the VulnerabilityCount field value
func (o *GraphNodeInfo) GetVulnerabilityCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VulnerabilityCount
}

// GetVulnerabilityCountOk returns a tuple with the VulnerabilityCount field value
// and a boolean to check if the value has been set.
func (o *GraphNodeInfo) GetVulnerabilityCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VulnerabilityCount, true
}

// SetVulnerabilityCount sets field value
func (o *GraphNodeInfo) SetVulnerabilityCount(v int32) {
	o.VulnerabilityCount = v
}

// GetWarnAlarmCount returns the WarnAlarmCount field value
func (o *GraphNodeInfo) GetWarnAlarmCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WarnAlarmCount
}

// GetWarnAlarmCountOk returns a tuple with the WarnAlarmCount field value
// and a boolean to check if the value has been set.
func (o *GraphNodeInfo) GetWarnAlarmCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WarnAlarmCount, true
}

// SetWarnAlarmCount sets field value
func (o *GraphNodeInfo) SetWarnAlarmCount(v int32) {
	o.WarnAlarmCount = v
}

func (o GraphNodeInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphNodeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alerts_count"] = o.AlertsCount
	toSerialize["cloud_compliance_count"] = o.CloudComplianceCount
	toSerialize["cloud_warn_alarm_count"] = o.CloudWarnAlarmCount
	toSerialize["compliance_count"] = o.ComplianceCount
	toSerialize["exploitable_secrets_count"] = o.ExploitableSecretsCount
	toSerialize["exploitable_vulnerabilities_count"] = o.ExploitableVulnerabilitiesCount
	toSerialize["name"] = o.Name
	toSerialize["node_id"] = o.NodeId
	toSerialize["secrets_count"] = o.SecretsCount
	toSerialize["vulnerability_count"] = o.VulnerabilityCount
	toSerialize["warn_alarm_count"] = o.WarnAlarmCount
	return toSerialize, nil
}

func (o *GraphNodeInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alerts_count",
		"cloud_compliance_count",
		"cloud_warn_alarm_count",
		"compliance_count",
		"exploitable_secrets_count",
		"exploitable_vulnerabilities_count",
		"name",
		"node_id",
		"secrets_count",
		"vulnerability_count",
		"warn_alarm_count",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGraphNodeInfo := _GraphNodeInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGraphNodeInfo)

	if err != nil {
		return err
	}

	*o = GraphNodeInfo(varGraphNodeInfo)

	return err
}

type NullableGraphNodeInfo struct {
	value *GraphNodeInfo
	isSet bool
}

func (v NullableGraphNodeInfo) Get() *GraphNodeInfo {
	return v.value
}

func (v *NullableGraphNodeInfo) Set(val *GraphNodeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphNodeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphNodeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphNodeInfo(val *GraphNodeInfo) *NullableGraphNodeInfo {
	return &NullableGraphNodeInfo{value: val, isSet: true}
}

func (v NullableGraphNodeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphNodeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


