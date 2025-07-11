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

// checks if the GraphThreatNodeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphThreatNodeInfo{}

// GraphThreatNodeInfo struct for GraphThreatNodeInfo
type GraphThreatNodeInfo struct {
	AlertsCount int32 `json:"alerts_count"`
	AttackPath [][]string `json:"attack_path"`
	CloudComplianceCount int32 `json:"cloud_compliance_count"`
	CloudWarnAlarmCount int32 `json:"cloud_warn_alarm_count"`
	ComplianceCount int32 `json:"compliance_count"`
	Count int32 `json:"count"`
	ExploitableSecretsCount int32 `json:"exploitable_secrets_count"`
	ExploitableVulnerabilitiesCount int32 `json:"exploitable_vulnerabilities_count"`
	Id string `json:"id"`
	Label string `json:"label"`
	NodeType string `json:"node_type"`
	Nodes map[string]GraphNodeInfo `json:"nodes"`
	SecretsCount int32 `json:"secrets_count"`
	VulnerabilityCount int32 `json:"vulnerability_count"`
	WarnAlarmCount int32 `json:"warn_alarm_count"`
}

type _GraphThreatNodeInfo GraphThreatNodeInfo

// NewGraphThreatNodeInfo instantiates a new GraphThreatNodeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphThreatNodeInfo(alertsCount int32, attackPath [][]string, cloudComplianceCount int32, cloudWarnAlarmCount int32, complianceCount int32, count int32, exploitableSecretsCount int32, exploitableVulnerabilitiesCount int32, id string, label string, nodeType string, nodes map[string]GraphNodeInfo, secretsCount int32, vulnerabilityCount int32, warnAlarmCount int32) *GraphThreatNodeInfo {
	this := GraphThreatNodeInfo{}
	this.AlertsCount = alertsCount
	this.AttackPath = attackPath
	this.CloudComplianceCount = cloudComplianceCount
	this.CloudWarnAlarmCount = cloudWarnAlarmCount
	this.ComplianceCount = complianceCount
	this.Count = count
	this.ExploitableSecretsCount = exploitableSecretsCount
	this.ExploitableVulnerabilitiesCount = exploitableVulnerabilitiesCount
	this.Id = id
	this.Label = label
	this.NodeType = nodeType
	this.Nodes = nodes
	this.SecretsCount = secretsCount
	this.VulnerabilityCount = vulnerabilityCount
	this.WarnAlarmCount = warnAlarmCount
	return &this
}

// NewGraphThreatNodeInfoWithDefaults instantiates a new GraphThreatNodeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphThreatNodeInfoWithDefaults() *GraphThreatNodeInfo {
	this := GraphThreatNodeInfo{}
	return &this
}

// GetAlertsCount returns the AlertsCount field value
func (o *GraphThreatNodeInfo) GetAlertsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AlertsCount
}

// GetAlertsCountOk returns a tuple with the AlertsCount field value
// and a boolean to check if the value has been set.
func (o *GraphThreatNodeInfo) GetAlertsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertsCount, true
}

// SetAlertsCount sets field value
func (o *GraphThreatNodeInfo) SetAlertsCount(v int32) {
	o.AlertsCount = v
}

// GetAttackPath returns the AttackPath field value
// If the value is explicit nil, the zero value for [][]string will be returned
func (o *GraphThreatNodeInfo) GetAttackPath() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}

	return o.AttackPath
}

// GetAttackPathOk returns a tuple with the AttackPath field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphThreatNodeInfo) GetAttackPathOk() ([][]string, bool) {
	if o == nil || IsNil(o.AttackPath) {
		return nil, false
	}
	return o.AttackPath, true
}

// SetAttackPath sets field value
func (o *GraphThreatNodeInfo) SetAttackPath(v [][]string) {
	o.AttackPath = v
}

// GetCloudComplianceCount returns the CloudComplianceCount field value
func (o *GraphThreatNodeInfo) GetCloudComplianceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CloudComplianceCount
}

// GetCloudComplianceCountOk returns a tuple with the CloudComplianceCount field value
// and a boolean to check if the value has been set.
func (o *GraphThreatNodeInfo) GetCloudComplianceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudComplianceCount, true
}

// SetCloudComplianceCount sets field value
func (o *GraphThreatNodeInfo) SetCloudComplianceCount(v int32) {
	o.CloudComplianceCount = v
}

// GetCloudWarnAlarmCount returns the CloudWarnAlarmCount field value
func (o *GraphThreatNodeInfo) GetCloudWarnAlarmCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CloudWarnAlarmCount
}

// GetCloudWarnAlarmCountOk returns a tuple with the CloudWarnAlarmCount field value
// and a boolean to check if the value has been set.
func (o *GraphThreatNodeInfo) GetCloudWarnAlarmCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudWarnAlarmCount, true
}

// SetCloudWarnAlarmCount sets field value
func (o *GraphThreatNodeInfo) SetCloudWarnAlarmCount(v int32) {
	o.CloudWarnAlarmCount = v
}

// GetComplianceCount returns the ComplianceCount field value
func (o *GraphThreatNodeInfo) GetComplianceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ComplianceCount
}

// GetComplianceCountOk returns a tuple with the ComplianceCount field value
// and a boolean to check if the value has been set.
func (o *GraphThreatNodeInfo) GetComplianceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComplianceCount, true
}

// SetComplianceCount sets field value
func (o *GraphThreatNodeInfo) SetComplianceCount(v int32) {
	o.ComplianceCount = v
}

// GetCount returns the Count field value
func (o *GraphThreatNodeInfo) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *GraphThreatNodeInfo) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *GraphThreatNodeInfo) SetCount(v int32) {
	o.Count = v
}

// GetExploitableSecretsCount returns the ExploitableSecretsCount field value
func (o *GraphThreatNodeInfo) GetExploitableSecretsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExploitableSecretsCount
}

// GetExploitableSecretsCountOk returns a tuple with the ExploitableSecretsCount field value
// and a boolean to check if the value has been set.
func (o *GraphThreatNodeInfo) GetExploitableSecretsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExploitableSecretsCount, true
}

// SetExploitableSecretsCount sets field value
func (o *GraphThreatNodeInfo) SetExploitableSecretsCount(v int32) {
	o.ExploitableSecretsCount = v
}

// GetExploitableVulnerabilitiesCount returns the ExploitableVulnerabilitiesCount field value
func (o *GraphThreatNodeInfo) GetExploitableVulnerabilitiesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExploitableVulnerabilitiesCount
}

// GetExploitableVulnerabilitiesCountOk returns a tuple with the ExploitableVulnerabilitiesCount field value
// and a boolean to check if the value has been set.
func (o *GraphThreatNodeInfo) GetExploitableVulnerabilitiesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExploitableVulnerabilitiesCount, true
}

// SetExploitableVulnerabilitiesCount sets field value
func (o *GraphThreatNodeInfo) SetExploitableVulnerabilitiesCount(v int32) {
	o.ExploitableVulnerabilitiesCount = v
}

// GetId returns the Id field value
func (o *GraphThreatNodeInfo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GraphThreatNodeInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GraphThreatNodeInfo) SetId(v string) {
	o.Id = v
}

// GetLabel returns the Label field value
func (o *GraphThreatNodeInfo) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *GraphThreatNodeInfo) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *GraphThreatNodeInfo) SetLabel(v string) {
	o.Label = v
}

// GetNodeType returns the NodeType field value
func (o *GraphThreatNodeInfo) GetNodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *GraphThreatNodeInfo) GetNodeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *GraphThreatNodeInfo) SetNodeType(v string) {
	o.NodeType = v
}

// GetNodes returns the Nodes field value
// If the value is explicit nil, the zero value for map[string]GraphNodeInfo will be returned
func (o *GraphThreatNodeInfo) GetNodes() map[string]GraphNodeInfo {
	if o == nil {
		var ret map[string]GraphNodeInfo
		return ret
	}

	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphThreatNodeInfo) GetNodesOk() (*map[string]GraphNodeInfo, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return &o.Nodes, true
}

// SetNodes sets field value
func (o *GraphThreatNodeInfo) SetNodes(v map[string]GraphNodeInfo) {
	o.Nodes = v
}

// GetSecretsCount returns the SecretsCount field value
func (o *GraphThreatNodeInfo) GetSecretsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SecretsCount
}

// GetSecretsCountOk returns a tuple with the SecretsCount field value
// and a boolean to check if the value has been set.
func (o *GraphThreatNodeInfo) GetSecretsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretsCount, true
}

// SetSecretsCount sets field value
func (o *GraphThreatNodeInfo) SetSecretsCount(v int32) {
	o.SecretsCount = v
}

// GetVulnerabilityCount returns the VulnerabilityCount field value
func (o *GraphThreatNodeInfo) GetVulnerabilityCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VulnerabilityCount
}

// GetVulnerabilityCountOk returns a tuple with the VulnerabilityCount field value
// and a boolean to check if the value has been set.
func (o *GraphThreatNodeInfo) GetVulnerabilityCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VulnerabilityCount, true
}

// SetVulnerabilityCount sets field value
func (o *GraphThreatNodeInfo) SetVulnerabilityCount(v int32) {
	o.VulnerabilityCount = v
}

// GetWarnAlarmCount returns the WarnAlarmCount field value
func (o *GraphThreatNodeInfo) GetWarnAlarmCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WarnAlarmCount
}

// GetWarnAlarmCountOk returns a tuple with the WarnAlarmCount field value
// and a boolean to check if the value has been set.
func (o *GraphThreatNodeInfo) GetWarnAlarmCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WarnAlarmCount, true
}

// SetWarnAlarmCount sets field value
func (o *GraphThreatNodeInfo) SetWarnAlarmCount(v int32) {
	o.WarnAlarmCount = v
}

func (o GraphThreatNodeInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphThreatNodeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alerts_count"] = o.AlertsCount
	if o.AttackPath != nil {
		toSerialize["attack_path"] = o.AttackPath
	}
	toSerialize["cloud_compliance_count"] = o.CloudComplianceCount
	toSerialize["cloud_warn_alarm_count"] = o.CloudWarnAlarmCount
	toSerialize["compliance_count"] = o.ComplianceCount
	toSerialize["count"] = o.Count
	toSerialize["exploitable_secrets_count"] = o.ExploitableSecretsCount
	toSerialize["exploitable_vulnerabilities_count"] = o.ExploitableVulnerabilitiesCount
	toSerialize["id"] = o.Id
	toSerialize["label"] = o.Label
	toSerialize["node_type"] = o.NodeType
	if o.Nodes != nil {
		toSerialize["nodes"] = o.Nodes
	}
	toSerialize["secrets_count"] = o.SecretsCount
	toSerialize["vulnerability_count"] = o.VulnerabilityCount
	toSerialize["warn_alarm_count"] = o.WarnAlarmCount
	return toSerialize, nil
}

func (o *GraphThreatNodeInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alerts_count",
		"attack_path",
		"cloud_compliance_count",
		"cloud_warn_alarm_count",
		"compliance_count",
		"count",
		"exploitable_secrets_count",
		"exploitable_vulnerabilities_count",
		"id",
		"label",
		"node_type",
		"nodes",
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

	varGraphThreatNodeInfo := _GraphThreatNodeInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGraphThreatNodeInfo)

	if err != nil {
		return err
	}

	*o = GraphThreatNodeInfo(varGraphThreatNodeInfo)

	return err
}

type NullableGraphThreatNodeInfo struct {
	value *GraphThreatNodeInfo
	isSet bool
}

func (v NullableGraphThreatNodeInfo) Get() *GraphThreatNodeInfo {
	return v.value
}

func (v *NullableGraphThreatNodeInfo) Set(val *GraphThreatNodeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphThreatNodeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphThreatNodeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphThreatNodeInfo(val *GraphThreatNodeInfo) *NullableGraphThreatNodeInfo {
	return &NullableGraphThreatNodeInfo{value: val, isSet: true}
}

func (v NullableGraphThreatNodeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphThreatNodeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


