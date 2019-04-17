// Code generated by thriftrw v1.18.0. DO NOT EDIT.
// @generated

package sqlblobs

import thriftreflect "go.uber.org/thriftrw/thriftreflect"

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "sqlblobs",
	Package:  "github.com/uber/cadence/.gen/go/sqlblobs",
	FilePath: "sqlblobs.thrift",
	SHA1:     "175e9a8bab33ee32eb59820194815d8808033bb3",
	Raw:      rawIDL,
}

const rawIDL = "// Copyright (c) 2017 Uber Technologies, Inc.\n//\n// Permission is hereby granted, free of charge, to any person obtaining a copy\n// of this software and associated documentation files (the \"Software\"), to deal\n// in the Software without restriction, including without limitation the rights\n// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell\n// copies of the Software, and to permit persons to whom the Software is\n// furnished to do so, subject to the following conditions:\n//\n// The above copyright notice and this permission notice shall be included in\n// all copies or substantial portions of the Software.\n//\n// THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\n// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\n// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\n// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\n// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\n// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN\n// THE SOFTWARE.\n\nnamespace java com.uber.cadence.sqlblobs\n\n\nstruct ShardInfo {\n\t10: optional i32 stolenSinceRenew\n\t12: optional i64 (js.type = \"Long\") updatedAtNanos\n\t14: optional i64 (js.type = \"Long\") replicationAckLevel\n\t16: optional i64 (js.type = \"Long\") transferAckLevel\n\t18: optional i64 (js.type = \"Long\") timerAckLevelNanos\n\t24: optional i64 (js.type = \"Long\") domainNotificationVersion\n\t34: optional binary clusterTransferAckLevel\n    36: optional binary clusterTimerAckLevel\n    38: optional string owner\n}\n\nstruct DomainInfo {\n  10: optional string name\n  12: optional string description\n  14: optional string owner\n  16: optional i32 status\n  18: optional i16 retentionDays\n  20: optional bool emitMetric\n  22: optional bool isGlobalDomain\n  24: optional string archivalBucket\n  26: optional i16 archivalStatus\n  28: optional i64 (js.type = \"Long\") configVersion\n  30: optional i64 (js.type = \"Long\") notificationVersion\n  32: optional i64 (js.type = \"Long\") failoverNotificationVersion\n  34: optional i64 (js.type = \"Long\") failoverVersion\n  36: optional string activeClusterName\n  38: optional binary clusters\n  40: optional binary data\n}\n\nstruct HistoryTreeInfo {\n\t10: optional bool inProgress    // For fork operation to prevent race condition with deleting history\n\t12: optional i64 (js.type = \"Long\") createdTimeNanos // For fork operation to prevent race condition of leaking event data when forking branches fail. Also can be used for clean up leaked data\n\t14: optional binary ancestors\n\t16: optional string info // For lookup back to workflow during debugging, also background cleanup when fork operation cannot finish self cleanup due to crash.\n}\n\nstruct CurrentExecutionInfo {\n  10: optional string createRequestID\n  12: optional i32 state\n  14: optional i32 closeStatus\n  16: optional i64 (js.type = \"Long\") startVersion\n  18: optional i64 (js.type = \"Long\") lastWriteVersion\n}\n\nstruct WorkflowExecutionInfo {\n\t10: optional binary parentDomainID\n\t12: optional string parentWorkflowID\n\t14: optional binary parentRunID\n\t16: optional i64 (js.type = \"Long\") initiatedID\n\t18: optional i64 (js.type = \"Long\") completionEventBatchID\n\t20: optional binary completionEvent\n\t22: optional string completionEventEncoding\n\t24: optional string taskList\n\t26: optional string workflowTypeName\n\t28: optional i64 (js.type = \"Long\") workflowTimeoutSeconds\n\t30: optional i64 (js.type = \"Long\") decisionTaskTimeoutMinutes\n\t32: optional binary executionContext\n\t34: optional i32 state\n\t36: optional i32 closeStatus\n\n    38: optional i64 (js.type = \"Long\") startVersion\n    40: optional i64 (js.type = \"Long\") currentVersion\n    42: optional i64 (js.type = \"Long\") lastWriteVersion\n    44: optional i64 (js.type = \"Long\") lastWriteEventID\n    46: optional binary lastReplicationInfo\n\n    48: optional i64 (js.type = \"Long\") lastEventTaskID\n\t50: optional i64 (js.type = \"Long\") lastFirstEventID\n\t52: optional i64 (js.type = \"Long\") lastProcessedEvent\n\t54: optional i64 (js.type = \"Long\") startTimeNanos\n\t56: optional i64 (js.type = \"Long\") lastUpdatedTimeNanos\n\t58: optional i64 (js.type = \"Long\") decisionVersion\n\t60: optional i64 (js.type = \"Long\") decisionScheduleID\n\t62: optional i64 (js.type = \"Long\") decisionStartedID\n\t64: optional i64 (js.type = \"Long\") decisionTimeout\n\t66: optional i64 (js.type = \"Long\") decisionAttempt\n\t68: optional i64 (js.type = \"Long\") decisionTimestampNanos\n\t70: optional i16 cancelRequested\n\t72: optional string createRequestID\n\t74: optional string decisionRequestID\n\t76: optional string cancelRequestID\n\t78: optional string stickyTaskList\n\t80: optional i64 (js.type = \"Long\") stickyScheduleToStartTimeout\n\n\t82: optional i64 (js.type = \"Long\") retryAttempt\n    84: optional i32 retryInitialIntervalSeconds\n    86: optional i32 retryMaximumIntervalSeconds\n    88: optional i32 retryMaximumAttempts\n    90: optional i32 retryExpirationSeconds\n    92: optional double retryBackoffCoefficient\n    94: optional i64 (js.type = \"Long\") retryExpirationTimeNanos\n    96: optional binary retryNonRetryableErrors\n\t98: optional bool hasRetryPolicy\n\n\t100: optional string cronSchedule\n\n    102: optional i32 eventStoreVersion\n    104: optional binary eventBranchToken\n\n    106: optional i64 (js.type = \"Long\") signalCount\n    108: optional i64 (js.type = \"Long\") historySize\n\n    110: optional string clientLibraryVersion\n    112: optional string clientFeatureVersion\n    114: optional string clientImpl\n}\n\nstruct ActivityInfo {\n    10: optional i64 (js.type = \"Long\") version\n    12: optional i64 (js.type = \"Long\") scheduledEventBatchID\n    14: optional binary scheduledEvent\n    16: optional string scheduledEventEncoding\n    18: optional i64 (js.type = \"Long\") scheduledTimeNanos\n    20: optional i64 (js.type = \"Long\") startedID\n    22: optional binary startedEvent\n    24: optional string startedEventEncoding\n    26: optional i64 (js.type = \"Long\") startedTimeNanos\n    28: optional string activityID\n    30: optional string requestID\n    32: optional i64 (js.type = \"Long\") scheduleToStartTimeoutSeconds\n    34: optional i64 (js.type = \"Long\") scheduleToCloseTimeoutSeconds\n    36: optional i64 (js.type = \"Long\") startToCloseTimeoutSeconds\n    38: optional i64 (js.type = \"Long\") heartbeatTimeoutSeconds\n    40: optional bool cancelRequested\n    42: optional i64 (js.type = \"Long\") cancelRequestID\n    44: optional i32 timerTaskStatus\n    46: optional i32 attempt\n    48: optional string taskList\n    50: optional string startedIdentity\n    52: optional bool hasRetryPolicy\n    54: optional i32 retryInitialIntervalSeconds\n    56: optional i32 retryMaximumIntervalSeconds\n    58: optional i32 retryMaximumAttempts\n    60: optional i64 (js.type = \"Long\") retryExpirationTimeNanos\n    62: optional double retryBackoffCoefficient\n    64: optional binary retryNonRetryableErrors\n}\n\nstruct ChildExecutionInfo {\n    10: optional i64 (js.type = \"Long\") version\n    12: optional i64 (js.type = \"Long\") initiatedEventBatchID\n    14: optional i64 (js.type = \"Long\") startedID\n    16: optional binary initiatedEvent\n    18: optional string initiatedEventEncoding\n    20: optional string startedWorkflowID\n    22: optional binary startedRunID\n    24: optional binary startedEvent\n    26: optional string startedEventEncoding\n    28: optional string createRequestID\n    30: optional string domainName\n    32: optional string workflowTypeName\n}\n\nstruct SignalInfo {\n    10: optional i64 (js.type = \"Long\") version\n    12: optional string requestID\n    14: optional string name\n    16: optional binary input\n    18: optional binary control\n}\n\nstruct RequestCancelInfo {\n    10: optional i64 (js.type = \"Long\") version\n    12: optional string cancelRequestID\n}\n\nstruct TimerInfo {\n    10: optional i64 (js.type = \"Long\") version\n    12: optional i64 (js.type = \"Long\") startedID\n    14: optional i64 (js.type = \"Long\") expiryTimeNanos\n    16: optional i64 (js.type = \"Long\") taskID\n}\n\nstruct BufferedReplicationTaskInfo {\n    10: optional i64 (js.type = \"Long\") version\n    12: optional i64 (js.type = \"Long\") nextEventID\n    14: optional binary history\n    16: optional string historyEncoding\n    18: optional binary newRunHistory\n    20: optional string newRunHistoryEncoding\n    22: optional i32 eventStoreVersion\n    24: optional i32 newRunEventStoreVersion\n}\n\nstruct TaskInfo {\n  10: optional string workflowID\n  12: optional binary runID\n  13: optional i64 (js.type = \"Long\") scheduleID\n  14: optional i64 (js.type = \"Long\") expiryTimeNanos\n}\n\nstruct TaskListInfo {\n  10: optional i16 kind // {Normal, Sticky}\n  12: optional i64 (js.type = \"Long\") ackLevel\n  14: optional i64 (js.type = \"Long\") expiryTimeNanos\n  16: optional i64 (js.type = \"Long\") lastUpdatedNanos\n}\n\nstruct TransferTaskInfo {\n\t10: optional binary domainID\n\t12: optional string workflowID\n\t14: optional binary runID\n\t16: optional i16 taskType\n    18: optional binary targetDomainID\n    20: optional string targetWorkflowID\n    22: optional binary targetRunID\n    24: optional string taskList\n\t26: optional i16 targetChildWorkflowOnly\n\t28: optional i64 (js.type = \"Long\") scheduleID\n\t30: optional i64 (js.type = \"Long\") version\n\t32: optional i64 (js.type = \"Long\") visibilityTimestampNanos\n}\n\nstruct TimerTaskInfo {\n\t10: optional binary domainID\n\t12: optional string workflowID\n\t14: optional binary runID\n\t16: optional i16 taskType\n    18: optional i16 timeoutType\n\t20: optional i64 (js.type = \"Long\") version\n\t22: optional i64 (js.type = \"Long\") scheduleAttempt\n\t24: optional i64 (js.type = \"Long\") eventID\n}\n\nstruct ReplicationTaskInfo {\n\t10: optional binary domainID\n\t12: optional string workflowID\n\t14: optional binary runID\n\t16: optional i16 taskType\n    18: optional i64 (js.type = \"Long\") version\n\t20: optional i64 (js.type = \"Long\") firstEventID\n\t22: optional i64 (js.type = \"Long\") nextEventID\n\t24: optional i64 (js.type = \"Long\") scheduledID\n\t26: optional i32 eventStoreVersion\n    28: optional i32 newRunEventStoreVersion\n\t30: optional binary branch_token\n\t32: optional binary lastReplicationInfo\n\t34: optional binary newRunBranchToken\n\t36: optional bool resetWorkflow\n}\n"
