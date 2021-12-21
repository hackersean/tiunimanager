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

	// switchover
	TIEM_MASTER_SLAVE_SWITCHOVER_NOT_FOUND TIEM_ERROR_CODE = 21000
	TIEM_MASTER_SLAVE_SWITCHOVER_FAILED    TIEM_ERROR_CODE = 21001

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
	TIEM_RESOURCE_DECRYPT_PASSWD_ERROR             TIEM_ERROR_CODE = 530

	TIEM_DEFAULT_PARAM_GROUP_NOT_DEL                 TIEM_ERROR_CODE = 20500
	TIEM_DEFAULT_PARAM_GROUP_NOT_MODIFY              TIEM_ERROR_CODE = 20501
	TIEM_MODIFY_PARAM_FAILED                         TIEM_ERROR_CODE = 20502
	TIEM_CONVERT_OBJ_FAILED                          TIEM_ERROR_CODE = 20503
	TIEM_PARAMETER_GROUP_CREATE_ERROR                TIEM_ERROR_CODE = 20504
	TIEM_PARAMETER_GROUP_UPDATE_ERROR                TIEM_ERROR_CODE = 20505
	TIEM_PARAMETER_GROUP_DELETE_ERROR                TIEM_ERROR_CODE = 20506
	TIEM_PARAMETER_GROUP_QUERY_ERROR                 TIEM_ERROR_CODE = 20507
	TIEM_PARAMETER_GROUP_DETAIL_ERROR                TIEM_ERROR_CODE = 20508
	TIEM_PARAMETER_GROUP_COPY_ERROR                  TIEM_ERROR_CODE = 20509
	TIEM_PARAMETER_GROUP_APPLY_ERROR                 TIEM_ERROR_CODE = 20510
	TIEM_PARAMETER_GROUP_CREATE_RELATION_PARAM_ERROR TIEM_ERROR_CODE = 20511
	TIEM_PARAMETER_GROUP_DELETE_RELATION_PARAM_ERROR TIEM_ERROR_CODE = 20512
	TIEM_PARAMETER_GROUP_UPDATE_RELATION_PARAM_ERROR TIEM_ERROR_CODE = 20513
	TIEM_CLUSTER_PARAMETER_QUERY_ERROR               TIEM_ERROR_CODE = 20514
	TIEM_CLUSTER_PARAMETER_UPDATE_ERROR              TIEM_ERROR_CODE = 20515

	TIEM_PARAMETER_QUERY_ERROR  TIEM_ERROR_CODE = 20520
	TIEM_PARAMETER_CREATE_ERROR TIEM_ERROR_CODE = 20521
	TIEM_PARAMETER_DELETE_ERROR TIEM_ERROR_CODE = 20522
	TIEM_PARAMETER_DETAIL_ERROR TIEM_ERROR_CODE = 20523
	TIEM_PARAMETER_UPDATE_ERROR TIEM_ERROR_CODE = 20524

	// backup && restore
	TIEM_BACKUP_SYSTEM_CONFIG_NOT_FOUND TIEM_ERROR_CODE = 20600
	TIEM_BACKUP_SYSTEM_CONFIG_INVAILD   TIEM_ERROR_CODE = 20601
	TIEM_BACKUP_RECORD_CREATE_FAILED    TIEM_ERROR_CODE = 20602
	TIEM_BACKUP_RECORD_DELETE_FAILED    TIEM_ERROR_CODE = 20603
	TIEM_BACKUP_RECORD_QUERY_FAILED     TIEM_ERROR_CODE = 20604
	TIEM_BACKUP_STRATEGY_SAVE_FAILED    TIEM_ERROR_CODE = 20605
	TIEM_BACKUP_STRATEGY_QUERY_FAILED   TIEM_ERROR_CODE = 20606
	TIEM_BACKUP_STRATEGY_DELETE_FAILED  TIEM_ERROR_CODE = 20607
	TIEM_BACKUP_FILE_DELETE_FAILED      TIEM_ERROR_CODE = 20608
	TIEM_BACKUP_PATH_CREATE_FAILED      TIEM_ERROR_CODE = 20609

	// workflow
	TIEM_WORKFLOW_CREATE_FAILED    TIEM_ERROR_CODE = 40100
	TIEM_WORKFLOW_QUERY_FAILED     TIEM_ERROR_CODE = 40101
	TIEM_WORKFLOW_DETAIL_FAILED    TIEM_ERROR_CODE = 40102
	TIEM_WORKFLOW_START_FAILED     TIEM_ERROR_CODE = 40103
	TIEM_WORKFLOW_DEFINE_NOT_FOUND TIEM_ERROR_CODE = 40104

	// import && export
	TIEM_TRANSPORT_SYSTEM_CONFIG_NOT_FOUND TIEM_ERROR_CODE = 60100
	TIEM_TRANSPORT_SYSTEM_CONFIG_INVALID   TIEM_ERROR_CODE = 60101
	TIEM_TRANSPORT_RECORD_NOT_FOUND        TIEM_ERROR_CODE = 60102
	TIEM_TRANSPORT_RECORD_CREATE_FAILED    TIEM_ERROR_CODE = 60103
	TIEM_TRANSPORT_RECORD_DELETE_FAILED    TIEM_ERROR_CODE = 60104
	TIEM_TRANSPORT_RECORD_QUERY_FAILED     TIEM_ERROR_CODE = 60105
	TIEM_TRANSPORT_FILE_DELETE_FAILED      TIEM_ERROR_CODE = 60106
	TIEM_TRANSPORT_PATH_CREATE_FAILED      TIEM_ERROR_CODE = 60107
	TIEM_TRANSPORT_FILE_SIZE_INVALID       TIEM_ERROR_CODE = 60108
	TIEM_TRANSPORT_FILE_UPLOAD_FAILED      TIEM_ERROR_CODE = 60109
	TIEM_TRANSPORT_FILE_DOWNLOAD_FAILED    TIEM_ERROR_CODE = 60110
	TIEM_TRANSPORT_FILE_TRANSFER_LIMITED   TIEM_ERROR_CODE = 60111

	// dashboard && monitor
	TIEM_DASHBOARD_NOT_FOUND TIEM_ERROR_CODE = 80100

	TIEM_CHANGE_FEED_NOT_FOUND              TIEM_ERROR_CODE = 701
	TIEM_CHANGE_FEED_DUPLICATE_ID           TIEM_ERROR_CODE = 702
	TIEM_CHANGE_FEED_STATUS_CONFLICT        TIEM_ERROR_CODE = 703
	TIEM_CHANGE_FEED_LOCK_EXPIRED           TIEM_ERROR_CODE = 704
	TIEM_CHANGE_FEED_UNSUPPORTED_DOWNSTREAM TIEM_ERROR_CODE = 705
	TIEM_CHANGE_FEED_CREATE_ERROR           TIEM_ERROR_CODE = 706

	TIEM_INSTANCE_NOT_FOUND          TIEM_ERROR_CODE = 20801
	TIEM_DELETE_INSTANCE_ERROR       TIEM_ERROR_CODE = 20802
	TIEM_CHECK_PLACEMENT_RULES_ERROR TIEM_ERROR_CODE = 20803

	TIEM_CHECK_TIFLASH_MAX_REPLICAS_ERROR TIEM_ERROR_CODE = 20901
	TIEM_NOT_FOUND_TIDB_ERROR             TIEM_ERROR_CODE = 20902
	TIEM_CONNECT_DB_ERROR                 TIEM_ERROR_CODE = 20903
	TIEM_SCAN_MAX_REPLICA_COUNT_ERROR     TIEM_ERROR_CODE = 20904

	TIEM_CHECK_CLUSTER_VERSION_ERROR TIEM_ERROR_CODE = 21301

	TIEM_SECOND_PARTY_OPERATION_NOT_FOUND TIEM_ERROR_CODE = 1101

	TIEM_CLUSTER_LOG_QUERY_FAILED TIEM_ERROR_CODE = 80300
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
	TIEM_PARAMETER_INVALID:         {TIEM_PARAMETER_INVALID, "parameter is invalid", 400},
	TIEM_UNRECOGNIZED_ERROR:        {TIEM_UNRECOGNIZED_ERROR, "unrecognized error", 500},
	TIEM_MARSHAL_ERROR:             {TIEM_MARSHAL_ERROR, "marshal error", 500},
	TIEM_UNMARSHAL_ERROR:           {TIEM_UNMARSHAL_ERROR, "unmarshal error", 500},

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

	// dashboard && monitor
	TIEM_DASHBOARD_NOT_FOUND: {TIEM_DASHBOARD_NOT_FOUND, "dashboard is not found", 500},

	// workflow
	TIEM_WORKFLOW_CREATE_FAILED:    {TIEM_WORKFLOW_CREATE_FAILED, "workflow create failed", 500},
	TIEM_WORKFLOW_QUERY_FAILED:     {TIEM_WORKFLOW_QUERY_FAILED, "workflow query failed", 500},
	TIEM_WORKFLOW_DETAIL_FAILED:    {TIEM_WORKFLOW_DETAIL_FAILED, "workflow detail failed", 500},
	TIEM_WORKFLOW_START_FAILED:     {TIEM_WORKFLOW_START_FAILED, "workflow start failed", 500},
	TIEM_WORKFLOW_DEFINE_NOT_FOUND: {TIEM_WORKFLOW_DEFINE_NOT_FOUND, "workflow define not found", 404},

	// import && export
	TIEM_TRANSPORT_SYSTEM_CONFIG_NOT_FOUND: {TIEM_TRANSPORT_SYSTEM_CONFIG_NOT_FOUND, "data transport system config not found", 404},
	TIEM_TRANSPORT_SYSTEM_CONFIG_INVALID:   {TIEM_TRANSPORT_SYSTEM_CONFIG_INVALID, "data transport system config invalid", 400},
	TIEM_TRANSPORT_RECORD_NOT_FOUND:        {TIEM_TRANSPORT_RECORD_NOT_FOUND, "data transport record is not found", 404},
	TIEM_TRANSPORT_RECORD_CREATE_FAILED:    {TIEM_TRANSPORT_RECORD_CREATE_FAILED, "create data transport record failed", 500},
	TIEM_TRANSPORT_RECORD_DELETE_FAILED:    {TIEM_TRANSPORT_RECORD_DELETE_FAILED, "delete data transport record failed", 500},
	TIEM_TRANSPORT_RECORD_QUERY_FAILED:     {TIEM_TRANSPORT_RECORD_QUERY_FAILED, "query data transport record failed", 500},
	TIEM_TRANSPORT_FILE_DELETE_FAILED:      {TIEM_TRANSPORT_FILE_DELETE_FAILED, "remove transport file failed", 500},
	TIEM_TRANSPORT_PATH_CREATE_FAILED:      {TIEM_TRANSPORT_PATH_CREATE_FAILED, "data transport filepath create failed", 500},
	TIEM_TRANSPORT_FILE_SIZE_INVALID:       {TIEM_TRANSPORT_FILE_SIZE_INVALID, "data transport file size invalid", 400},
	TIEM_TRANSPORT_FILE_UPLOAD_FAILED:      {TIEM_TRANSPORT_FILE_UPLOAD_FAILED, "data transport file upload failed", 500},
	TIEM_TRANSPORT_FILE_DOWNLOAD_FAILED:    {TIEM_TRANSPORT_FILE_DOWNLOAD_FAILED, "data transport file download failed", 500},
	TIEM_TRANSPORT_FILE_TRANSFER_LIMITED:   {TIEM_TRANSPORT_FILE_TRANSFER_LIMITED, "exceed limit file transfer num", 400},

	// backup && restore
	TIEM_BACKUP_SYSTEM_CONFIG_NOT_FOUND: {TIEM_BACKUP_SYSTEM_CONFIG_NOT_FOUND, "backup system config not found", 404},
	TIEM_BACKUP_SYSTEM_CONFIG_INVAILD:   {TIEM_BACKUP_SYSTEM_CONFIG_INVAILD, "backup system config invalid", 400},
	TIEM_BACKUP_RECORD_CREATE_FAILED:    {TIEM_BACKUP_RECORD_CREATE_FAILED, "create backup record failed", 500},
	TIEM_BACKUP_RECORD_DELETE_FAILED:    {TIEM_BACKUP_RECORD_DELETE_FAILED, "delete backup record failed", 500},
	TIEM_BACKUP_RECORD_QUERY_FAILED:     {TIEM_BACKUP_RECORD_QUERY_FAILED, "query backup record failed", 500},
	TIEM_BACKUP_STRATEGY_SAVE_FAILED:    {TIEM_BACKUP_STRATEGY_SAVE_FAILED, "save backup strategy failed", 500},
	TIEM_BACKUP_STRATEGY_QUERY_FAILED:   {TIEM_BACKUP_STRATEGY_QUERY_FAILED, "query backup strategy failed", 500},
	TIEM_BACKUP_STRATEGY_DELETE_FAILED:  {TIEM_BACKUP_STRATEGY_DELETE_FAILED, "delete backup strategy failed", 500},
	TIEM_BACKUP_FILE_DELETE_FAILED:      {TIEM_BACKUP_FILE_DELETE_FAILED, "remove backup file failed", 500},
	TIEM_BACKUP_PATH_CREATE_FAILED:      {TIEM_BACKUP_PATH_CREATE_FAILED, "backup filepath create failed", 500},

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
	TIEM_RESOURCE_TEMPLATE_FILE_NOT_FOUND:          {TIEM_RESOURCE_TEMPLATE_FILE_NOT_FOUND, "resource template file is not existed", 500},

	// tenant
	TIEM_ACCOUNT_NOT_FOUND:       {TIEM_ACCOUNT_NOT_FOUND, "account not found", 500},
	TIEM_TENANT_NOT_FOUND:        {TIEM_TENANT_NOT_FOUND, "tenant not found", 500},
	TIEM_QUERY_PERMISSION_FAILED: {TIEM_QUERY_PERMISSION_FAILED, "query permission failed", 500},
	TIEM_ADD_TOKEN_FAILED:        {TIEM_ADD_TOKEN_FAILED, "add token failed", 500},
	TIEM_TOKEN_NOT_FOUND:         {TIEM_TOKEN_NOT_FOUND, "token not found", 500},

	// param group & cluster param
	TIEM_DEFAULT_PARAM_GROUP_NOT_DEL:                 {TIEM_DEFAULT_PARAM_GROUP_NOT_DEL, "Not allow to deleted the default parameter group", 409},
	TIEM_DEFAULT_PARAM_GROUP_NOT_MODIFY:              {TIEM_DEFAULT_PARAM_GROUP_NOT_MODIFY, "Not allow to modify the default parameter group", 409},
	TIEM_MODIFY_PARAM_FAILED:                         {TIEM_MODIFY_PARAM_FAILED, "Failed to modify parameter", 500},
	TIEM_CONVERT_OBJ_FAILED:                          {TIEM_CONVERT_OBJ_FAILED, "Failed to convert data type of parameter", 500},
	TIEM_PARAMETER_GROUP_CREATE_ERROR:                {TIEM_PARAMETER_GROUP_CREATE_ERROR, "Failed to create parameter group", 500},
	TIEM_PARAMETER_GROUP_UPDATE_ERROR:                {TIEM_PARAMETER_GROUP_UPDATE_ERROR, "Failed to update parameter group", 500},
	TIEM_PARAMETER_GROUP_DELETE_ERROR:                {TIEM_PARAMETER_GROUP_DELETE_ERROR, "Failed to delete parameter group", 500},
	TIEM_PARAMETER_GROUP_QUERY_ERROR:                 {TIEM_PARAMETER_GROUP_QUERY_ERROR, "Failed to query parameter group", 500},
	TIEM_PARAMETER_GROUP_DETAIL_ERROR:                {TIEM_PARAMETER_GROUP_DETAIL_ERROR, "Failed to get parameter group details", 500},
	TIEM_PARAMETER_GROUP_COPY_ERROR:                  {TIEM_PARAMETER_GROUP_COPY_ERROR, "Failed to copy parameter group", 500},
	TIEM_PARAMETER_GROUP_APPLY_ERROR:                 {TIEM_PARAMETER_GROUP_APPLY_ERROR, "Failed to apply parameter group", 500},
	TIEM_PARAMETER_GROUP_CREATE_RELATION_PARAM_ERROR: {TIEM_PARAMETER_GROUP_CREATE_RELATION_PARAM_ERROR, "Failed to create relation parameter in parameter group", 500},
	TIEM_PARAMETER_GROUP_DELETE_RELATION_PARAM_ERROR: {TIEM_PARAMETER_GROUP_DELETE_RELATION_PARAM_ERROR, "Failed to delete relation parameter in parameter group", 500},
	TIEM_PARAMETER_GROUP_UPDATE_RELATION_PARAM_ERROR: {TIEM_PARAMETER_GROUP_UPDATE_RELATION_PARAM_ERROR, "Failed to update relation parameter in parameter group", 500},
	TIEM_CLUSTER_PARAMETER_QUERY_ERROR:               {TIEM_CLUSTER_PARAMETER_QUERY_ERROR, "Failed to query cluster parameters", 500},
	TIEM_CLUSTER_PARAMETER_UPDATE_ERROR:              {TIEM_CLUSTER_PARAMETER_UPDATE_ERROR, "Failed to update cluster parameters", 500},

	TIEM_PARAMETER_QUERY_ERROR:  {TIEM_PARAMETER_QUERY_ERROR, "Failed to query parameter by parameter group id", 500},
	TIEM_PARAMETER_CREATE_ERROR: {TIEM_PARAMETER_CREATE_ERROR, "Failed to create parameter", 500},
	TIEM_PARAMETER_DELETE_ERROR: {TIEM_PARAMETER_DELETE_ERROR, "Failed to delete parameter", 500},
	TIEM_PARAMETER_DETAIL_ERROR: {TIEM_PARAMETER_DETAIL_ERROR, "Failed to detail parameter", 500},
	TIEM_PARAMETER_UPDATE_ERROR: {TIEM_PARAMETER_UPDATE_ERROR, "Failed to update parameter", 500},

	// change feed
	TIEM_CHANGE_FEED_NOT_FOUND:              {TIEM_CHANGE_FEED_NOT_FOUND, "change feed task not found", 404},
	TIEM_CHANGE_FEED_DUPLICATE_ID:           {TIEM_CHANGE_FEED_DUPLICATE_ID, "duplicate id", 500},
	TIEM_CHANGE_FEED_STATUS_CONFLICT:        {TIEM_CHANGE_FEED_STATUS_CONFLICT, "task status conflict", 409},
	TIEM_CHANGE_FEED_LOCK_EXPIRED:           {TIEM_CHANGE_FEED_LOCK_EXPIRED, "task status lock expired", 409},
	TIEM_CHANGE_FEED_UNSUPPORTED_DOWNSTREAM: {TIEM_CHANGE_FEED_UNSUPPORTED_DOWNSTREAM, "task downstream type not supported", 500},
	TIEM_CHANGE_FEED_CREATE_ERROR:           {TIEM_CHANGE_FEED_CREATE_ERROR, "failed to create change feed task", 500},

	// master/slave switchover
	TIEM_MASTER_SLAVE_SWITCHOVER_NOT_FOUND: {TIEM_MASTER_SLAVE_SWITCHOVER_NOT_FOUND, "master/slave relation not found", 404},
	TIEM_MASTER_SLAVE_SWITCHOVER_FAILED:    {TIEM_MASTER_SLAVE_SWITCHOVER_FAILED, "master/slave switchover failed", 500},

	TIEM_CLUSTER_LOG_QUERY_FAILED: {TIEM_CLUSTER_LOG_QUERY_FAILED, "Failed to query cluster log", 500},
}
