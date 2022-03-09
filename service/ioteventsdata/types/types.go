// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Contains the configuration information of an acknowledge action.
type AcknowledgeActionConfiguration struct {

	// The note that you can leave when you acknowledge the alarm.
	Note *string

	noSmithyDocumentSerde
}

// Information needed to acknowledge the alarm.
type AcknowledgeAlarmActionRequest struct {

	// The name of the alarm model.
	//
	// This member is required.
	AlarmModelName *string

	// The request ID. Each ID must be unique within each batch.
	//
	// This member is required.
	RequestId *string

	// The value of the key used as a filter to select only the alarms associated with
	// the key
	// (https://docs.aws.amazon.com/iotevents/latest/apireference/API_CreateAlarmModel.html#iotevents-CreateAlarmModel-request-key).
	KeyValue *string

	// The note that you can leave when you acknowledge the alarm.
	Note *string

	noSmithyDocumentSerde
}

// Contains information about an alarm.
type Alarm struct {

	// The name of the alarm model.
	AlarmModelName *string

	// The version of the alarm model.
	AlarmModelVersion *string

	// Contains information about the current state of the alarm.
	AlarmState *AlarmState

	// The time the alarm was created, in the Unix epoch format.
	CreationTime *time.Time

	// The value of the key used as a filter to select only the alarms associated with
	// the key
	// (https://docs.aws.amazon.com/iotevents/latest/apireference/API_CreateAlarmModel.html#iotevents-CreateAlarmModel-request-key).
	KeyValue *string

	// The time the alarm was last updated, in the Unix epoch format.
	LastUpdateTime *time.Time

	// A non-negative integer that reflects the severity level of the alarm.
	Severity *int32

	noSmithyDocumentSerde
}

// Contains information about the current state of the alarm.
type AlarmState struct {

	// Contains information about the action that you can take to respond to the alarm.
	CustomerAction *CustomerAction

	// Information needed to evaluate data.
	RuleEvaluation *RuleEvaluation

	// The name of the alarm state. The state name can be one of the following
	// values:
	//
	// * DISABLED - When the alarm is in the DISABLED state, it isn't ready to
	// evaluate data. To enable the alarm, you must change the alarm to the NORMAL
	// state.
	//
	// * NORMAL - When the alarm is in the NORMAL state, it's ready to evaluate
	// data.
	//
	// * ACTIVE - If the alarm is in the ACTIVE state, the alarm is invoked.
	//
	// *
	// ACKNOWLEDGED - When the alarm is in the ACKNOWLEDGED state, the alarm was
	// invoked and you acknowledged the alarm.
	//
	// * SNOOZE_DISABLED - When the alarm is
	// in the SNOOZE_DISABLED state, the alarm is disabled for a specified period of
	// time. After the snooze time, the alarm automatically changes to the NORMAL
	// state.
	//
	// * LATCHED - When the alarm is in the LATCHED state, the alarm was
	// invoked. However, the data that the alarm is currently evaluating is within the
	// specified range. To change the alarm to the NORMAL state, you must acknowledge
	// the alarm.
	StateName AlarmStateName

	// Contains information about alarm state changes.
	SystemEvent *SystemEvent

	noSmithyDocumentSerde
}

// Contains a summary of an alarm.
type AlarmSummary struct {

	// The name of the alarm model.
	AlarmModelName *string

	// The version of the alarm model.
	AlarmModelVersion *string

	// The time the alarm was created, in the Unix epoch format.
	CreationTime *time.Time

	// The value of the key used as a filter to select only the alarms associated with
	// the key
	// (https://docs.aws.amazon.com/iotevents/latest/apireference/API_CreateAlarmModel.html#iotevents-CreateAlarmModel-request-key).
	KeyValue *string

	// The time the alarm was last updated, in the Unix epoch format.
	LastUpdateTime *time.Time

	// The name of the alarm state. The state name can be one of the following
	// values:
	//
	// * DISABLED - When the alarm is in the DISABLED state, it isn't ready to
	// evaluate data. To enable the alarm, you must change the alarm to the NORMAL
	// state.
	//
	// * NORMAL - When the alarm is in the NORMAL state, it's ready to evaluate
	// data.
	//
	// * ACTIVE - If the alarm is in the ACTIVE state, the alarm is invoked.
	//
	// *
	// ACKNOWLEDGED - When the alarm is in the ACKNOWLEDGED state, the alarm was
	// invoked and you acknowledged the alarm.
	//
	// * SNOOZE_DISABLED - When the alarm is
	// in the SNOOZE_DISABLED state, the alarm is disabled for a specified period of
	// time. After the snooze time, the alarm automatically changes to the NORMAL
	// state.
	//
	// * LATCHED - When the alarm is in the LATCHED state, the alarm was
	// invoked. However, the data that the alarm is currently evaluating is within the
	// specified range. To change the alarm to the NORMAL state, you must acknowledge
	// the alarm.
	StateName AlarmStateName

	noSmithyDocumentSerde
}

// Contains error messages associated with one of the following requests:
//
// *
// BatchAcknowledgeAlarm
// (https://docs.aws.amazon.com/iotevents/latest/apireference/API_iotevents-data_BatchAcknowledgeAlarm.html)
//
// *
// BatchDisableAlarm
// (https://docs.aws.amazon.com/iotevents/latest/apireference/API_iotevents-data_BatchDisableAlarm.html)
//
// *
// BatchEnableAlarm
// (https://docs.aws.amazon.com/iotevents/latest/apireference/API_iotevents-data_BatchEnableAlarm.html)
//
// *
// BatchResetAlarm
// (https://docs.aws.amazon.com/iotevents/latest/apireference/API_iotevents-data_BatchResetAlarm.html)
//
// *
// BatchSnoozeAlarm
// (https://docs.aws.amazon.com/iotevents/latest/apireference/API_iotevents-data_BatchSnoozeAlarm.html)
type BatchAlarmActionErrorEntry struct {

	// The error code.
	ErrorCode ErrorCode

	// A message that describes the error.
	ErrorMessage *string

	// The request ID. Each ID must be unique within each batch.
	RequestId *string

	noSmithyDocumentSerde
}

// Contains information about the errors encountered.
type BatchPutMessageErrorEntry struct {

	// The error code.
	ErrorCode ErrorCode

	// A message that describes the error.
	ErrorMessage *string

	// The ID of the message that caused the error. (See the value corresponding to the
	// "messageId" key in the "message" object.)
	MessageId *string

	noSmithyDocumentSerde
}

// Information about the error that occurred when attempting to update a detector.
type BatchUpdateDetectorErrorEntry struct {

	// The error code.
	ErrorCode ErrorCode

	// A message that describes the error.
	ErrorMessage *string

	// The "messageId" of the update request that caused the error. (The value of the
	// "messageId" in the update request "Detector" object.)
	MessageId *string

	noSmithyDocumentSerde
}

// Contains information about the action that you can take to respond to the alarm.
type CustomerAction struct {

	// Contains the configuration information of an acknowledge action.
	AcknowledgeActionConfiguration *AcknowledgeActionConfiguration

	// The name of the action. The action name can be one of the following values:
	//
	// *
	// SNOOZE - When you snooze the alarm, the alarm state changes to
	// SNOOZE_DISABLED.
	//
	// * ENABLE - When you enable the alarm, the alarm state changes
	// to NORMAL.
	//
	// * DISABLE - When you disable the alarm, the alarm state changes to
	// DISABLED.
	//
	// * ACKNOWLEDGE - When you acknowledge the alarm, the alarm state
	// changes to ACKNOWLEDGED.
	//
	// * RESET - When you reset the alarm, the alarm state
	// changes to NORMAL.
	//
	// For more information, see the AlarmState
	// (https://docs.aws.amazon.com/iotevents/latest/apireference/API_iotevents-data_AlarmState.html)
	// API.
	ActionName CustomerActionName

	// Contains the configuration information of a disable action.
	DisableActionConfiguration *DisableActionConfiguration

	// Contains the configuration information of an enable action.
	EnableActionConfiguration *EnableActionConfiguration

	// Contains the configuration information of a reset action.
	ResetActionConfiguration *ResetActionConfiguration

	// Contains the configuration information of a snooze action.
	SnoozeActionConfiguration *SnoozeActionConfiguration

	noSmithyDocumentSerde
}

// Information about the detector (instance).
type Detector struct {

	// The time the detector (instance) was created.
	CreationTime *time.Time

	// The name of the detector model that created this detector (instance).
	DetectorModelName *string

	// The version of the detector model that created this detector (instance).
	DetectorModelVersion *string

	// The value of the key (identifying the device or system) that caused the creation
	// of this detector (instance).
	KeyValue *string

	// The time the detector (instance) was last updated.
	LastUpdateTime *time.Time

	// The current state of the detector (instance).
	State *DetectorState

	noSmithyDocumentSerde
}

// Information about the current state of the detector instance.
type DetectorState struct {

	// The name of the state.
	//
	// This member is required.
	StateName *string

	// The current state of the detector's timers.
	//
	// This member is required.
	Timers []Timer

	// The current values of the detector's variables.
	//
	// This member is required.
	Variables []Variable

	noSmithyDocumentSerde
}

// The new state, variable values, and timer settings of the detector (instance).
type DetectorStateDefinition struct {

	// The name of the new state of the detector (instance).
	//
	// This member is required.
	StateName *string

	// The new values of the detector's timers. Any timer whose value isn't specified
	// is cleared, and its timeout event won't occur.
	//
	// This member is required.
	Timers []TimerDefinition

	// The new values of the detector's variables. Any variable whose value isn't
	// specified is cleared.
	//
	// This member is required.
	Variables []VariableDefinition

	noSmithyDocumentSerde
}

// Information about the detector state.
type DetectorStateSummary struct {

	// The name of the state.
	StateName *string

	noSmithyDocumentSerde
}

// Information about the detector (instance).
type DetectorSummary struct {

	// The time the detector (instance) was created.
	CreationTime *time.Time

	// The name of the detector model that created this detector (instance).
	DetectorModelName *string

	// The version of the detector model that created this detector (instance).
	DetectorModelVersion *string

	// The value of the key (identifying the device or system) that caused the creation
	// of this detector (instance).
	KeyValue *string

	// The time the detector (instance) was last updated.
	LastUpdateTime *time.Time

	// The current state of the detector (instance).
	State *DetectorStateSummary

	noSmithyDocumentSerde
}

// Contains the configuration information of a disable action.
type DisableActionConfiguration struct {

	// The note that you can leave when you disable the alarm.
	Note *string

	noSmithyDocumentSerde
}

// Information used to disable the alarm.
type DisableAlarmActionRequest struct {

	// The name of the alarm model.
	//
	// This member is required.
	AlarmModelName *string

	// The request ID. Each ID must be unique within each batch.
	//
	// This member is required.
	RequestId *string

	// The value of the key used as a filter to select only the alarms associated with
	// the key
	// (https://docs.aws.amazon.com/iotevents/latest/apireference/API_CreateAlarmModel.html#iotevents-CreateAlarmModel-request-key).
	KeyValue *string

	// The note that you can leave when you disable the alarm.
	Note *string

	noSmithyDocumentSerde
}

// Contains the configuration information of an enable action.
type EnableActionConfiguration struct {

	// The note that you can leave when you enable the alarm.
	Note *string

	noSmithyDocumentSerde
}

// Information needed to enable the alarm.
type EnableAlarmActionRequest struct {

	// The name of the alarm model.
	//
	// This member is required.
	AlarmModelName *string

	// The request ID. Each ID must be unique within each batch.
	//
	// This member is required.
	RequestId *string

	// The value of the key used as a filter to select only the alarms associated with
	// the key
	// (https://docs.aws.amazon.com/iotevents/latest/apireference/API_CreateAlarmModel.html#iotevents-CreateAlarmModel-request-key).
	KeyValue *string

	// The note that you can leave when you enable the alarm.
	Note *string

	noSmithyDocumentSerde
}

// Information about a message.
type Message struct {

	// The name of the input into which the message payload is transformed.
	//
	// This member is required.
	InputName *string

	// The ID to assign to the message. Within each batch sent, each "messageId" must
	// be unique.
	//
	// This member is required.
	MessageId *string

	// The payload of the message. This can be a JSON string or a Base-64-encoded
	// string representing binary data (in which case you must decode it).
	//
	// This member is required.
	Payload []byte

	// The timestamp associated with the message.
	Timestamp *TimestampValue

	noSmithyDocumentSerde
}

// Contains the configuration information of a reset action.
type ResetActionConfiguration struct {

	// The note that you can leave when you reset the alarm.
	Note *string

	noSmithyDocumentSerde
}

// Information needed to reset the alarm.
type ResetAlarmActionRequest struct {

	// The name of the alarm model.
	//
	// This member is required.
	AlarmModelName *string

	// The request ID. Each ID must be unique within each batch.
	//
	// This member is required.
	RequestId *string

	// The value of the key used as a filter to select only the alarms associated with
	// the key
	// (https://docs.aws.amazon.com/iotevents/latest/apireference/API_CreateAlarmModel.html#iotevents-CreateAlarmModel-request-key).
	KeyValue *string

	// The note that you can leave when you reset the alarm.
	Note *string

	noSmithyDocumentSerde
}

// Information needed to evaluate data.
type RuleEvaluation struct {

	// Information needed to compare two values with a comparison operator.
	SimpleRuleEvaluation *SimpleRuleEvaluation

	noSmithyDocumentSerde
}

// Information needed to compare two values with a comparison operator.
type SimpleRuleEvaluation struct {

	// The value of the input property, on the left side of the comparison operator.
	InputPropertyValue *string

	// The comparison operator.
	Operator ComparisonOperator

	// The threshold value, on the right side of the comparison operator.
	ThresholdValue *string

	noSmithyDocumentSerde
}

// Contains the configuration information of a snooze action.
type SnoozeActionConfiguration struct {

	// The note that you can leave when you snooze the alarm.
	Note *string

	// The snooze time in seconds. The alarm automatically changes to the NORMAL state
	// after this duration.
	SnoozeDuration *int32

	noSmithyDocumentSerde
}

// Information needed to snooze the alarm.
type SnoozeAlarmActionRequest struct {

	// The name of the alarm model.
	//
	// This member is required.
	AlarmModelName *string

	// The request ID. Each ID must be unique within each batch.
	//
	// This member is required.
	RequestId *string

	// The snooze time in seconds. The alarm automatically changes to the NORMAL state
	// after this duration.
	//
	// This member is required.
	SnoozeDuration *int32

	// The value of the key used as a filter to select only the alarms associated with
	// the key
	// (https://docs.aws.amazon.com/iotevents/latest/apireference/API_CreateAlarmModel.html#iotevents-CreateAlarmModel-request-key).
	KeyValue *string

	// The note that you can leave when you snooze the alarm.
	Note *string

	noSmithyDocumentSerde
}

// Contains the configuration information of alarm state changes.
type StateChangeConfiguration struct {

	// The trigger type. If the value is SNOOZE_TIMEOUT, the snooze duration ends and
	// the alarm automatically changes to the NORMAL state.
	TriggerType TriggerType

	noSmithyDocumentSerde
}

// Contains information about alarm state changes.
type SystemEvent struct {

	// The event type. If the value is STATE_CHANGE, the event contains information
	// about alarm state changes.
	EventType EventType

	// Contains the configuration information of alarm state changes.
	StateChangeConfiguration *StateChangeConfiguration

	noSmithyDocumentSerde
}

// The current state of a timer.
type Timer struct {

	// The name of the timer.
	//
	// This member is required.
	Name *string

	// The expiration time for the timer.
	//
	// This member is required.
	Timestamp *time.Time

	noSmithyDocumentSerde
}

// The new setting of a timer.
type TimerDefinition struct {

	// The name of the timer.
	//
	// This member is required.
	Name *string

	// The new setting of the timer (the number of seconds before the timer elapses).
	//
	// This member is required.
	Seconds *int32

	noSmithyDocumentSerde
}

// Contains information about a timestamp.
type TimestampValue struct {

	// The value of the timestamp, in the Unix epoch format.
	TimeInMillis *int64

	noSmithyDocumentSerde
}

// Information used to update the detector (instance).
type UpdateDetectorRequest struct {

	// The name of the detector model that created the detectors (instances).
	//
	// This member is required.
	DetectorModelName *string

	// The ID to assign to the detector update "message". Each "messageId" must be
	// unique within each batch sent.
	//
	// This member is required.
	MessageId *string

	// The new state, variable values, and timer settings of the detector (instance).
	//
	// This member is required.
	State *DetectorStateDefinition

	// The value of the input key attribute (identifying the device or system) that
	// caused the creation of this detector (instance).
	KeyValue *string

	noSmithyDocumentSerde
}

// The current state of the variable.
type Variable struct {

	// The name of the variable.
	//
	// This member is required.
	Name *string

	// The current value of the variable.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// The new value of the variable.
type VariableDefinition struct {

	// The name of the variable.
	//
	// This member is required.
	Name *string

	// The new value of the variable.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde