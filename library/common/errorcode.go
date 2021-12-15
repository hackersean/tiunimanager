/******************************************************************************
 * Copyright (c)  2021 PingCAP, Inc.                                          *
 * Licensed under the Apache License, Version 2.0 (the "License");            *
 * you may not use this file except in compliance with the License.           *
 * You may obtain a copy of the License at                                    *
 *                                                                            *
 * http://www.apache.org/licenses/LICENSE-2.0                                 *
 *                                                                            *
 * Unless required by applicable law or agreed to in writing, software        *
 * distributed under the License is distributed on an "AS IS" BASIS,          *
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.   *
 * See the License for the specific language governing permissions and        *
 * limitations under the License.                                             *
 *                                                                            *
 ******************************************************************************/

package common

type TIEM_ERROR_CODE int32

// all tiem error code
// todo standardize
const (
	TIEM_SUCCESS TIEM_ERROR_CODE = 0

	TIEM_METADB_SERVER_CALL_ERROR  TIEM_ERROR_CODE = 9998
	TIEM_CLUSTER_SERVER_CALL_ERROR TIEM_ERROR_CODE = 9999
	TIEM_TASK_TIMEOUT              TIEM_ERROR_CODE = 9997
	TIEM_FLOW_NOT_FOUND            TIEM_ERROR_CODE = 9996
	TIEM_TASK_FAILED               TIEM_ERROR_CODE = 9995
	TIEM_TASK_POLLING_TIME_OUT     TIEM_ERROR_CODE = 9994
	TIEM_TASK_CONFLICT             TIEM_ERROR_CODE = 9993
	TIEM_TASK_CANCELED             TIEM_ERROR_CODE = 9992

	TIEM_PARAMETER_INVALID  TIEM_ERROR_CODE = 1
	TIEM_UNRECOGNIZED_ERROR TIEM_ERROR_CODE = 2
	TIEM_MARSHAL_ERROR      TIEM_ERROR_CODE = 3
	TIEM_UNMARSHAL_ERROR    TIEM_ERROR_CODE = 4

	TIEM_TAKEOVER_SSH_CONNECT_ERROR     TIEM_ERROR_CODE = 20109
	TIEM_TAKEOVER_SFTP_ERROR            TIEM_ERROR_CODE = 20110
	TIEM_CLUSTER_NOT_FOUND              TIEM_ERROR_CODE = 20111
	TIEM_CLUSTER_RESOURCE_NOT_ENOUGH    TIEM_ERROR_CODE = 20112
	TIEM_CLUSTER_GET_CLUSTER_PORT_ERROR TIEM_ERROR_CODE = 20113
	TIEM_CLUSTER_MAINTENANCE_CONFLICT   TIEM_ERROR_CODE = 20114

	TIEM_ACCOUNT_NOT_FOUND       TIEM_ERROR_CODE = 100
	TIEM_TENANT_NOT_FOUND        TIEM_ERROR_CODE = 200
	TIEM_QUERY_PERMISSION_FAILED TIEM_ERROR_CODE = 300
	TIEM_ADD_TOKEN_FAILED        TIEM_ERROR_CODE = 400
	TIEM_TOKEN_NOT_FOUND         TIEM_ERROR_CODE = 401

	TIEM_RESOURCE_SQL_ERROR                        TIEM_ERROR_CODE = 500
	TIEM_RESOURCE_HOST_NOT_FOUND                   TIEM_ERROR_CODE = 501
	TIEM_RESOURCE_NO_ENOUGH_HOST                   TIEM_ERROR_CODE = 502
	TIEM_RESOURCE_NO_ENOUGH_DISK_AFTER_EXCLUDED    TIEM_ERROR_CODE = 503
	TIEM_RESOURCE_NO_ENOUGH_DISK_AFTER_DISK_FILTER TIEM_ERROR_CODE = 504
	TIEM_RESOURCE_NO_ENOUGH_DISK_AFTER_HOST_FILTER TIEM_ERROR_CODE = 505
	TIEM_RESOURCE_NO_ENOUGH_PORT                   TIEM_ERROR_CODE = 506
	TIEM_RESOURCE_NOT_ALL_SUCCEED                  TIEM_ERROR_CODE = 507
	TIEM_RESOURCE_INVALID_STRATEGY                 TIEM_ERROR_CODE = 508
	TIEM_RESOURCE_INVAILD_RECYCLE_TYPE             TIEM_ERROR_CODE = 509
	TIEM_UPDATE_HOST_STATUS_FAIL                   TIEM_ERROR_CODE = 510
	TIEM_RESERVE_HOST_FAIL                         TIEM_ERROR_CODE = 511
	TIEM_RESOURCE_NO_STOCK                         TIEM_ERROR_CODE = 512
	TIEM_RESOURCE_GET_DISK_ID_FAIL                 TIEM_ERROR_CODE = 513
	TIEM_RESOURCE_TRAIT_NOT_FOUND                  TIEM_ERROR_CODE = 514
	TIEM_RESOURCE_INVALID_LOCATION                 TIEM_ERROR_CODE = 515
	TIEM_RESOURCE_INVALID_ARCH                     TIEM_ERROR_CODE = 516
	TIEM_RESOURCE_ADD_TABLE_ERROR                  TIEM_ERROR_CODE = 517
	TIEM_RESOURCE_INIT_LABELS_ERROR                TIEM_ERROR_CODE = 518
	TIEM_RESOURCE_CREATE_HOST_ERROR                TIEM_ERROR_CODE = 519
	TIEM_RESOURCE_LOCK_TABLE_ERROR                 TIEM_ERROR_CODE = 520
	TIEM_RESOURCE_DELETE_HOST_ERROR                TIEM_ERROR_CODE = 521
	TIEM_RESOURCE_INVALID_PRODUCT_NAME             TIEM_ERROR_CODE = 522
	TIEM_RESOURCE_INVALID_PURPOSE                  TIEM_ERROR_CODE = 523
	TIEM_RESOURCE_INVALID_DISKTYPE                 TIEM_ERROR_CODE = 524
	TIEM_RESOURCE_HOST_ALREADY_EXIST               TIEM_ERROR_CODE = 525
	TIEM_RESOURCE_HOST_STILL_INUSED                TIEM_ERROR_CODE = 526
	TIEM_RESOURCE_CREATE_DISK_ERROR                TIEM_ERROR_CODE = 527
	TIEM_RESOURCE_INVALID_LABEL_NAEM               TIEM_ERROR_CODE = 528
	TIEM_RESOURCE_TEMPLATE_FILE_NOT_FOUND          TIEM_ERROR_CODE = 529

	TIEM_DASHBOARD_NOT_FOUND           TIEM_ERROR_CODE = 600
	TIEM_EXPORT_PARAM_INVALID          TIEM_ERROR_CODE = 601
	TIEM_EXPORT_PROCESS_FAILED         TIEM_ERROR_CODE = 602
	TIEM_IMPORT_PARAM_INVALID          TIEM_ERROR_CODE = 603
	TIEM_IMPORT_PROCESS_FAILED         TIEM_ERROR_CODE = 604
	TIEM_TRANSPORT_RECORD_NOT_FOUND    TIEM_ERROR_CODE = 605
	TIEM_BACKUP_PROCESS_FAILED         TIEM_ERROR_CODE = 606
	TIEM_RECOVER_PARAM_INVALID         TIEM_ERROR_CODE = 607
	TIEM_RECOVER_PROCESS_FAILED        TIEM_ERROR_CODE = 608
	TIEM_BACKUP_RECORD_DELETE_FAILED   TIEM_ERROR_CODE = 609
	TIEM_BACKUP_RECORD_QUERY_FAILED    TIEM_ERROR_CODE = 610
	TIEM_BACKUP_STRATEGY_PARAM_INVALID TIEM_ERROR_CODE = 611
	TIEM_BACKUP_STRATEGY_SAVE_FAILED   TIEM_ERROR_CODE = 612
	TIEM_BACKUP_STRATEGY_QUERY_FAILED  TIEM_ERROR_CODE = 613
	TIEM_MONITOR_NOT_FOUND             TIEM_ERROR_CODE = 614
	TIEM_TRANSPORT_RECORD_DEL_FAILED   TIEM_ERROR_CODE = 615
	TIEM_BACKUP_RECORD_NOT_FOUND       TIEM_ERROR_CODE = 616
	TIEM_BACKUP_STRATEGY_NOT_FOUND     TIEM_ERROR_CODE = 617
	TIEM_LIST_WORKFLOW_FAILED          TIEM_ERROR_CODE = 618
	TIEM_DETAIL_WORKFLOW_FAILED        TIEM_ERROR_CODE = 619
	TIEM_GET_CONFIG_FAILED             TIEM_ERROR_CODE = 620
	TIEM_TRANSPORT_RECORD_QUERY_FAIL   TIEM_ERROR_CODE = 621

	TIEM_DEFAULT_PARAM_GROUP_NOT_DEL                 TIEM_ERROR_CODE = 20500
	TIEM_MODIFY_PARAM_FAILED                         TIEM_ERROR_CODE = 20501
	TIEM_CONVERT_OBJ_FAILED                          TIEM_ERROR_CODE = 20502
	TIEM_PARAMETER_GROUP_CREATE_ERROR                TIEM_ERROR_CODE = 20503
	TIEM_PARAMETER_GROUP_UPDATE_ERROR                TIEM_ERROR_CODE = 20504
	TIEM_PARAMETER_GROUP_DELETE_ERROR                TIEM_ERROR_CODE = 20505
	TIEM_PARAMETER_GROUP_QUERY_ERROR                 TIEM_ERROR_CODE = 20506
	TIEM_PARAMETER_GROUP_DETAIL_ERROR                TIEM_ERROR_CODE = 20507
	TIEM_PARAMETER_GROUP_CREATE_RELATION_PARAM_ERROR TIEM_ERROR_CODE = 20508
	TIEM_PARAMETER_GROUP_DELETE_RELATION_PARAM_ERROR TIEM_ERROR_CODE = 20509
	TIEM_PARAMETER_GROUP_UPDATE_RELATION_PARAM_ERROR TIEM_ERROR_CODE = 20510
	TIEM_PARAMETER_QUERY_ERROR                       TIEM_ERROR_CODE = 20511
	TIEM_PARAMETER_CREATE_ERROR                      TIEM_ERROR_CODE = 20512
	TIEM_PARAMETER_DELETE_ERROR                      TIEM_ERROR_CODE = 20513
	TIEM_PARAMETER_DETAIL_ERROR                      TIEM_ERROR_CODE = 20514
	TIEM_PARAMETER_UPDATE_ERROR                      TIEM_ERROR_CODE = 20515
	TIEM_CLUSTER_PARAMETER_QUERY_ERROR               TIEM_ERROR_CODE = 20517
	TIEM_CLUSTER_PARAMETER_UPDATE_ERROR              TIEM_ERROR_CODE = 20518

	TIEM_CHANGE_FEED_NOT_FOUND              TIEM_ERROR_CODE = 701
	TIEM_CHANGE_FEED_DUPLICATE_ID           TIEM_ERROR_CODE = 702
	TIEM_CHANGE_FEED_STATUS_CONFLICT        TIEM_ERROR_CODE = 703
	TIEM_CHANGE_FEED_LOCK_EXPIRED           TIEM_ERROR_CODE = 704
	TIEM_CHANGE_FEED_UNSUPPORTED_DOWNSTREAM TIEM_ERROR_CODE = 705
	TIEM_CHANGE_FEED_CREATE_ERROR           TIEM_ERROR_CODE = 706

	TIEM_UPGRADE_QUERY_NOT_FOUND           TIEM_ERROR_CODE = 801
	TIEM_UPGRADE_PATH_DUPLICATED           TIEM_ERROR_CODE = 802
	TIEM_UPGRADE_QUERY_PATH_FAILED         TIEM_ERROR_CODE = 803
	TIEM_UPGRADE_QUERY_VERSION_DIFF_FAILED TIEM_ERROR_CODE = 804
	TIEM_UPGRADE_FAILED                    TIEM_ERROR_CODE = 805

	TIEM_INSTANCE_NOT_FOUND    TIEM_ERROR_CODE = 20801
	TIEM_DELETE_INSTANCE_ERROR TIEM_ERROR_CODE = 20802

	TIEM_SECOND_PARTY_OPERATION_NOT_FOUND TIEM_ERROR_CODE = 1101
)

type ErrorCodeExplanation struct {
	code        TIEM_ERROR_CODE
	explanation string
	httpCode    int
}

func (t TIEM_ERROR_CODE) GetHttpCode() int {
	return explanationContainer[t].httpCode
}

func (t TIEM_ERROR_CODE) Equal(code int32) bool {
	return code == int32(t)
}

func (t TIEM_ERROR_CODE) Explain() string {
	return explanationContainer[t].explanation
}

var explanationContainer = map[TIEM_ERROR_CODE]ErrorCodeExplanation{
	TIEM_SUCCESS: {code: TIEM_SUCCESS, explanation: "succeed", httpCode: 200},

	// system error
	TIEM_METADB_SERVER_CALL_ERROR:  {TIEM_METADB_SERVER_CALL_ERROR, "call metadb-Server failed", 500},
	TIEM_CLUSTER_SERVER_CALL_ERROR: {TIEM_CLUSTER_SERVER_CALL_ERROR, "call cluster-Server failed", 500},
	TIEM_PARAMETER_INVALID:         {TIEM_PARAMETER_INVALID, "parameter is invalid", 500},
	TIEM_UNRECOGNIZED_ERROR:        {TIEM_UNRECOGNIZED_ERROR, "unrecognized error", 500},
	TIEM_MARSHAL_ERROR:             {TIEM_MARSHAL_ERROR, "marshal error", 500},
	TIEM_UNMARSHAL_ERROR:           {TIEM_UNMARSHAL_ERROR, "UNmarshal error", 500},

	TIEM_TASK_TIMEOUT:          {TIEM_TASK_TIMEOUT, "task timeout", 500},
	TIEM_FLOW_NOT_FOUND:        {TIEM_FLOW_NOT_FOUND, "flow not found", 500},
	TIEM_TASK_FAILED:           {TIEM_TASK_FAILED, "task failed", 500},
	TIEM_TASK_CONFLICT:         {TIEM_TASK_CONFLICT, "task polling time out", 500},
	TIEM_TASK_CANCELED:         {TIEM_TASK_CONFLICT, "task canceled", 500},
	TIEM_TASK_POLLING_TIME_OUT: {TIEM_TASK_POLLING_TIME_OUT, "task polling time out", 500},

	// cluster management
	TIEM_CLUSTER_NOT_FOUND:            {TIEM_CLUSTER_NOT_FOUND, "cluster not found", 500},
	TIEM_TAKEOVER_SSH_CONNECT_ERROR:   {TIEM_TAKEOVER_SSH_CONNECT_ERROR, "ssh connect failed", 500},
	TIEM_TAKEOVER_SFTP_ERROR:          {TIEM_TAKEOVER_SFTP_ERROR, "sftp failed", 500},
	TIEM_CLUSTER_RESOURCE_NOT_ENOUGH:  {TIEM_CLUSTER_RESOURCE_NOT_ENOUGH, "no enough resource for cluster", 500},
	TIEM_CLUSTER_MAINTENANCE_CONFLICT: {TIEM_CLUSTER_MAINTENANCE_CONFLICT, "maintenance status conflict", 209},

	TIEM_DASHBOARD_NOT_FOUND: {TIEM_DASHBOARD_NOT_FOUND, "dashboard is not found", 500},

	// cluster import export
	TIEM_EXPORT_PARAM_INVALID:        {TIEM_EXPORT_PARAM_INVALID, "export data param invalid", 500},
	TIEM_EXPORT_PROCESS_FAILED:       {TIEM_EXPORT_PROCESS_FAILED, "export process failed", 500},
	TIEM_IMPORT_PARAM_INVALID:        {TIEM_IMPORT_PARAM_INVALID, "import data param invalid", 500},
	TIEM_IMPORT_PROCESS_FAILED:       {TIEM_IMPORT_PROCESS_FAILED, "import process failed", 500},
	TIEM_TRANSPORT_RECORD_DEL_FAILED: {TIEM_TRANSPORT_RECORD_DEL_FAILED, "delete data transport failed", 500},

	// cluster backup
	TIEM_TRANSPORT_RECORD_NOT_FOUND:    {TIEM_TRANSPORT_RECORD_NOT_FOUND, "transport record is not found", 500},
	TIEM_BACKUP_PROCESS_FAILED:         {TIEM_BACKUP_PROCESS_FAILED, "backup process failed", 500},
	TIEM_RECOVER_PARAM_INVALID:         {TIEM_RECOVER_PARAM_INVALID, "recover param invalid", 500},
	TIEM_RECOVER_PROCESS_FAILED:        {TIEM_RECOVER_PROCESS_FAILED, "recover process failed", 500},
	TIEM_BACKUP_RECORD_DELETE_FAILED:   {TIEM_BACKUP_RECORD_DELETE_FAILED, "delete backup record failed", 500},
	TIEM_BACKUP_RECORD_QUERY_FAILED:    {TIEM_BACKUP_RECORD_QUERY_FAILED, "query backup record failed", 500},
	TIEM_BACKUP_STRATEGY_PARAM_INVALID: {TIEM_BACKUP_STRATEGY_PARAM_INVALID, "backup strategy param invalid", 500},
	TIEM_BACKUP_STRATEGY_SAVE_FAILED:   {TIEM_BACKUP_STRATEGY_SAVE_FAILED, "save backup strategy failed", 500},
	TIEM_BACKUP_STRATEGY_QUERY_FAILED:  {TIEM_BACKUP_STRATEGY_QUERY_FAILED, "query backup strategy failed", 500},

	// resource
	TIEM_RESOURCE_SQL_ERROR:                        {TIEM_RESOURCE_SQL_ERROR, "resource sql error", 500},
	TIEM_RESOURCE_HOST_NOT_FOUND:                   {TIEM_RESOURCE_HOST_NOT_FOUND, "host not found", 500},
	TIEM_RESOURCE_NO_ENOUGH_HOST:                   {TIEM_RESOURCE_NO_ENOUGH_HOST, "no enough host resource", 500},
	TIEM_RESOURCE_NO_ENOUGH_DISK_AFTER_EXCLUDED:    {TIEM_RESOURCE_NO_ENOUGH_DISK_AFTER_EXCLUDED, "no enough disk resource after excluded", 500},
	TIEM_RESOURCE_NO_ENOUGH_DISK_AFTER_DISK_FILTER: {TIEM_RESOURCE_NO_ENOUGH_DISK_AFTER_DISK_FILTER, "no enough disk after disk filter", 500},
	TIEM_RESOURCE_NO_ENOUGH_DISK_AFTER_HOST_FILTER: {TIEM_RESOURCE_NO_ENOUGH_DISK_AFTER_HOST_FILTER, "no enough disk after host filter", 500},
	TIEM_RESOURCE_NO_ENOUGH_PORT:                   {TIEM_RESOURCE_NO_ENOUGH_PORT, "no enough port resource", 500},
	TIEM_RESOURCE_NOT_ALL_SUCCEED:                  {TIEM_RESOURCE_NOT_ALL_SUCCEED, "not all request succeed", 500},
	TIEM_RESOURCE_INVALID_STRATEGY:                 {TIEM_RESOURCE_INVALID_STRATEGY, "invalid alloc strategy", 500},
	TIEM_RESOURCE_TRAIT_NOT_FOUND:                  {TIEM_RESOURCE_TRAIT_NOT_FOUND, "trait not found by label name", 500},
	TIEM_RESOURCE_INVALID_LOCATION:                 {TIEM_RESOURCE_INVALID_LOCATION, "invalid location", 500},
	TIEM_RESOURCE_INVALID_ARCH:                     {TIEM_RESOURCE_INVALID_ARCH, "invalid arch", 500},

	// tenant
	TIEM_ACCOUNT_NOT_FOUND:       {TIEM_ACCOUNT_NOT_FOUND, "account not found", 500},
	TIEM_TENANT_NOT_FOUND:        {TIEM_TENANT_NOT_FOUND, "tenant not found", 500},
	TIEM_QUERY_PERMISSION_FAILED: {TIEM_QUERY_PERMISSION_FAILED, "query permission failed", 500},
	TIEM_ADD_TOKEN_FAILED:        {TIEM_ADD_TOKEN_FAILED, "add token failed", 500},
	TIEM_TOKEN_NOT_FOUND:         {TIEM_TOKEN_NOT_FOUND, "token not found", 500},

	// param group & cluster param
	TIEM_DEFAULT_PARAM_GROUP_NOT_DEL:                 {TIEM_DEFAULT_PARAM_GROUP_NOT_DEL, "The default param group cannot be deleted", 500},
	TIEM_MODIFY_PARAM_FAILED:                         {TIEM_MODIFY_PARAM_FAILED, "apply or modify parameters failed!", 500},
	TIEM_CONVERT_OBJ_FAILED:                          {TIEM_CONVERT_OBJ_FAILED, "convert obj failed!", 500},
	TIEM_PARAMETER_GROUP_CREATE_ERROR:                {TIEM_PARAMETER_GROUP_CREATE_ERROR, "failed to create parameter group", 500},
	TIEM_PARAMETER_GROUP_UPDATE_ERROR:                {TIEM_PARAMETER_GROUP_UPDATE_ERROR, "failed to update parameter group", 500},
	TIEM_PARAMETER_GROUP_DELETE_ERROR:                {TIEM_PARAMETER_GROUP_DELETE_ERROR, "failed to delete parameter group", 500},
	TIEM_PARAMETER_GROUP_QUERY_ERROR:                 {TIEM_PARAMETER_GROUP_QUERY_ERROR, "failed to query parameter group", 500},
	TIEM_PARAMETER_GROUP_DETAIL_ERROR:                {TIEM_PARAMETER_GROUP_DETAIL_ERROR, "failed to detail parameter group", 500},
	TIEM_PARAMETER_GROUP_CREATE_RELATION_PARAM_ERROR: {TIEM_PARAMETER_GROUP_CREATE_RELATION_PARAM_ERROR, "failed to create parameter group associated parameters", 500},
	TIEM_PARAMETER_GROUP_DELETE_RELATION_PARAM_ERROR: {TIEM_PARAMETER_GROUP_DELETE_RELATION_PARAM_ERROR, "failed to delete parameter group associated parameters", 500},
	TIEM_PARAMETER_GROUP_UPDATE_RELATION_PARAM_ERROR: {TIEM_PARAMETER_GROUP_UPDATE_RELATION_PARAM_ERROR, "failed to update parameter group associated parameters", 500},
	TIEM_PARAMETER_QUERY_ERROR:                       {TIEM_PARAMETER_QUERY_ERROR, "failed to query parameter by parameter group id", 500},
	TIEM_PARAMETER_CREATE_ERROR:                      {TIEM_PARAMETER_CREATE_ERROR, "failed to create parameter", 500},
	TIEM_PARAMETER_DELETE_ERROR:                      {TIEM_PARAMETER_DELETE_ERROR, "failed to delete parameter", 500},
	TIEM_PARAMETER_DETAIL_ERROR:                      {TIEM_PARAMETER_DETAIL_ERROR, "failed to detail parameter", 500},
	TIEM_PARAMETER_UPDATE_ERROR:                      {TIEM_PARAMETER_UPDATE_ERROR, "failed to update parameter", 500},
	TIEM_CLUSTER_PARAMETER_QUERY_ERROR:               {TIEM_CLUSTER_PARAMETER_QUERY_ERROR, "failed to query cluster parameter", 500},
	TIEM_CLUSTER_PARAMETER_UPDATE_ERROR:              {TIEM_CLUSTER_PARAMETER_UPDATE_ERROR, "failed to update cluster parameter", 500},

	// change feed
	TIEM_CHANGE_FEED_NOT_FOUND:              {TIEM_CHANGE_FEED_NOT_FOUND, "change feed task not found", 404},
	TIEM_CHANGE_FEED_DUPLICATE_ID:           {TIEM_CHANGE_FEED_DUPLICATE_ID, "duplicate id", 500},
	TIEM_CHANGE_FEED_STATUS_CONFLICT:        {TIEM_CHANGE_FEED_STATUS_CONFLICT, "task status conflict", 409},
	TIEM_CHANGE_FEED_LOCK_EXPIRED:           {TIEM_CHANGE_FEED_LOCK_EXPIRED, "task status lock expired", 409},
	TIEM_CHANGE_FEED_UNSUPPORTED_DOWNSTREAM: {TIEM_CHANGE_FEED_UNSUPPORTED_DOWNSTREAM, "task downstream type not supported", 500},
	TIEM_CHANGE_FEED_CREATE_ERROR:           {TIEM_CHANGE_FEED_CREATE_ERROR, "failed to create change feed task", 500},
}
