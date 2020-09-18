// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by model-api-gen. DO NOT EDIT.

package monitor

import (
	time "time"

	"yunion.io/x/onecloud/pkg/apis"
)

// SAlert is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SAlert.
type SAlert struct {
	// db.SVirtualResourceBase
	apis.SEnabledResourceBase
	apis.SStatusStandaloneResourceBase
	apis.SScopedResourceBase
	// Frequency is evaluate period
	Frequency int64       `json:"frequency"`
	Settings  interface{} `json:"settings"`
	Level     string      `json:"level"`
	Message   string      `json:"message"`
	UsedBy    string      `json:"used_by"`
	// Silenced       bool
	ExecutionError string `json:"execution_error"`
	// If an alert rule has a configured `For` and the query violates the configured threshold
	// it will first go from `OK` to `Pending`. Going from `OK` to `Pending` will not send any
	// notifications. Once the alert rule has been firing for more than `For` duration, it will
	// change to `Alerting` and send alert notifications.
	For                 int64       `json:"for"`
	EvalData            interface{} `json:"eval_data"`
	State               string      `json:"state"`
	NoDataState         string      `json:"no_data_state"`
	ExecutionErrorState string      `json:"execution_error_state"`
	LastStateChange     time.Time   `json:"last_state_change"`
	StateChanges        int         `json:"state_changes"`
}

// SAlertDashBoard is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SAlertDashBoard.
type SAlertDashBoard struct {
	// db.SVirtualResourceBase
	apis.SEnabledResourceBase
	apis.SStatusStandaloneResourceBase
	apis.SScopedResourceBase
	Refresh  string      `json:"refresh"`
	Settings interface{} `json:"settings"`
	Message  string      `json:"message"`
}

// SAlertJointsBase is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SAlertJointsBase.
type SAlertJointsBase struct {
	apis.SVirtualJointResourceBase
	AlertId string `json:"alert_id"`
}

// SAlertRecord is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SAlertRecord.
type SAlertRecord struct {
	// db.SVirtualResourceBase
	apis.SEnabledResourceBase
	apis.SStatusStandaloneResourceBase
	apis.SScopedResourceBase
	AlertId  string      `json:"alert_id"`
	Level    string      `json:"level"`
	State    string      `json:"state"`
	EvalData interface{} `json:"eval_data"`
}

// SAlertnotification is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SAlertnotification.
type SAlertnotification struct {
	SAlertJointsBase
	NotificationId string      `json:"notification_id"`
	State          string      `json:"state"`
	Index          byte        `json:"index"`
	UsedBy         string      `json:"used_by"`
	Params         interface{} `json:"params"`
}

// SDataSource is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SDataSource.
type SDataSource struct {
	apis.SStandaloneResourceBase
	Type      string `json:"type"`
	Url       string `json:"url"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Database  string `json:"database"`
	IsDefault *bool  `json:"is_default,omitempty"`
}

// SMetric is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SMetric.
type SMetric struct {
	apis.SVirtualJointResourceBase
	MeasurementId string `json:"measurement_id"`
	FieldId       string `json:"field_id"`
}

// SMetricField is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SMetricField.
type SMetricField struct {
	// db.SVirtualResourceBase
	apis.SEnabledResourceBase
	apis.SStatusStandaloneResourceBase
	apis.SScopedResourceBase
	DisplayName string `json:"display_name"`
	Unit        string `json:"unit"`
	ValueType   string `json:"value_type"`
}

// SMetricMeasurement is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SMetricMeasurement.
type SMetricMeasurement struct {
	// db.SVirtualResourceBase
	apis.SEnabledResourceBase
	apis.SStatusStandaloneResourceBase
	apis.SScopedResourceBase
	ResType     string `json:"res_type"`
	Database    string `json:"database"`
	DisplayName string `json:"display_name"`
}

// SNotification is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SNotification.
type SNotification struct {
	apis.SVirtualResourceBase
	Type                  string      `json:"type"`
	IsDefault             bool        `json:"is_default"`
	SendReminder          bool        `json:"send_reminder"`
	DisableResolveMessage bool        `json:"disable_resolve_message"`
	Frequency             int64       `json:"frequency"`
	Settings              interface{} `json:"settings"`
}

// SSuggestSysAlert is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SSuggestSysAlert.
type SSuggestSysAlert struct {
	apis.SVirtualResourceBase
	apis.SEnabledResourceBase
	// 监控规则对应的json对象
	RuleName      string      `json:"rule_name"`
	MonitorConfig interface{} `json:"monitor_config"`
	// 监控规则type：Rule Type
	Type    string      `json:"type"`
	ResMeta interface{} `json:"res_meta"`
	// problem may have more than one ,so the problem instanceof jsonutils.jsonArray
	Problem interface{} `json:"problem"`
	// Suggest string               `width:"256"  list:"user" update:"user"`
	Action       string `json:"action"`
	ResId        string `json:"res_id"`
	CloudEnv     string `json:"cloud_env"`
	Provider     string `json:"provider"`
	Project      string `json:"project"`
	Cloudaccount string `json:"cloudaccount"`
	// 费用
	Amount float64 `json:"amount"`
	// 币种
	Currency string `json:"currency"`
}

// SSuggestSysRule is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SSuggestSysRule.
type SSuggestSysRule struct {
	apis.SStandaloneResourceBase
	apis.SEnabledResourceBase
	Type     string      `json:"type"`
	Period   string      `json:"period"`
	TimeFrom string      `json:"time_from"`
	Setting  interface{} `json:"setting"`
	ExecTime time.Time   `json:"exec_time"`
}

// SSuggestSysRuleConfig is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SSuggestSysRuleConfig.
type SSuggestSysRuleConfig struct {
	apis.SStandaloneResourceBase
	apis.SScopedResourceBase
	// RuleId is SSuggestSysRule model object id
	// RuleId string `width:"36" charset:"ascii" nullable:"true" list:"user" create:"optional"`
	// Type is suggestsysrules driver type
	Type string `json:"type"`
	// ResourceType is suggestsysrules driver resource type
	ResourceType string `json:"resource_type"`
	// ResourceId is suggest alert result resource id
	ResourceId string `json:"resource_id"`
	// IgnoreAlert means whether or not show SSuggestSysAlert results for current scope
	IgnoreAlert bool `json:"ignore_alert"`
}
