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

package errors

type EM_ERROR_CODE int32

// all tiem error code
// todo standardize
const (
	TIEM_SUCCESS EM_ERROR_CODE = 0

	TIEM_METADB_SERVER_CALL_ERROR  EM_ERROR_CODE = 9998
	TIEM_CLUSTER_SERVER_CALL_ERROR EM_ERROR_CODE = 9999

	TIEM_PARAMETER_INVALID  EM_ERROR_CODE = 1
	TIEM_UNRECOGNIZED_ERROR EM_ERROR_CODE = 2

	TIEM_TAKEOVER_SSH_CONNECT_ERROR     EM_ERROR_CODE = 20109
	TIEM_TAKEOVER_SFTP_ERROR            EM_ERROR_CODE = 20110
	TIEM_CLUSTER_NOT_FOUND              EM_ERROR_CODE = 20111
	TIEM_CLUSTER_RESOURCE_NOT_ENOUGH    EM_ERROR_CODE = 20112
	TIEM_CLUSTER_GET_CLUSTER_PORT_ERROR EM_ERROR_CODE = 20113

	// backup && restore
	TIEM_BACKUP_SYSTEM_CONFIG_NOT_FOUND EM_ERROR_CODE = 20600
	TIEM_BACKUP_SYSTEM_CONFIG_INVAILD   EM_ERROR_CODE = 20601
	TIEM_BACKUP_RECORD_CREATE_FAILED    EM_ERROR_CODE = 20602
	TIEM_BACKUP_RECORD_DELETE_FAILED    EM_ERROR_CODE = 20603
	TIEM_BACKUP_RECORD_QUERY_FAILED     EM_ERROR_CODE = 20604
	TIEM_BACKUP_STRATEGY_SAVE_FAILED    EM_ERROR_CODE = 20605
	TIEM_BACKUP_STRATEGY_QUERY_FAILED   EM_ERROR_CODE = 20606
	TIEM_BACKUP_STRATEGY_DELETE_FAILED  EM_ERROR_CODE = 20607
	TIEM_BACKUP_FILE_DELETE_FAILED      EM_ERROR_CODE = 20608
	TIEM_BACKUP_PATH_CREATE_FAILED      EM_ERROR_CODE = 20609

	// workflow
	TIEM_WORKFLOW_CREATE_FAILED EM_ERROR_CODE = 40100
	TIEM_WORKFLOW_QUERY_FAILED  EM_ERROR_CODE = 40101
	TIEM_WORKFLOW_DETAIL_FAILED EM_ERROR_CODE = 40102
	TIEM_WORKFLOW_START_FAILED  EM_ERROR_CODE = 40103

	// import && export
	TIEM_TRANSPORT_SYSTEM_CONFIG_NOT_FOUND EM_ERROR_CODE = 60100
	TIEM_TRANSPORT_SYSTEM_CONFIG_INVALID   EM_ERROR_CODE = 60101
	TIEM_TRANSPORT_RECORD_NOT_FOUND        EM_ERROR_CODE = 60102
	TIEM_TRANSPORT_RECORD_CREATE_FAILED    EM_ERROR_CODE = 60103
	TIEM_TRANSPORT_RECORD_DELETE_FAILED    EM_ERROR_CODE = 60104
	TIEM_TRANSPORT_RECORD_QUERY_FAILED     EM_ERROR_CODE = 60105
	TIEM_TRANSPORT_FILE_DELETE_FAILED      EM_ERROR_CODE = 60106
	TIEM_TRANSPORT_PATH_CREATE_FAILED      EM_ERROR_CODE = 60107
	TIEM_TRANSPORT_FILE_SIZE_INVALID       EM_ERROR_CODE = 60108
	TIEM_TRANSPORT_FILE_UPLOAD_FAILED      EM_ERROR_CODE = 60109
	TIEM_TRANSPORT_FILE_DOWNLOAD_FAILED    EM_ERROR_CODE = 60110
	TIEM_TRANSPORT_FILE_TRANSFER_LIMITED   EM_ERROR_CODE = 60111

	// dashboard && monitor
	TIEM_DASHBOARD_NOT_FOUND EM_ERROR_CODE = 80100

	TIEM_ACCOUNT_NOT_FOUND       EM_ERROR_CODE = 100
	TIEM_TENANT_NOT_FOUND        EM_ERROR_CODE = 200
	TIEM_QUERY_PERMISSION_FAILED EM_ERROR_CODE = 300
	TIEM_ADD_TOKEN_FAILED        EM_ERROR_CODE = 400
	TIEM_TOKEN_NOT_FOUND         EM_ERROR_CODE = 401

	TIEM_RESOURCE_SQL_ERROR                        EM_ERROR_CODE = 500
	TIEM_RESOURCE_HOST_NOT_FOUND                   EM_ERROR_CODE = 501
	TIEM_RESOURCE_NO_ENOUGH_HOST                   EM_ERROR_CODE = 502
	TIEM_RESOURCE_NO_ENOUGH_DISK_AFTER_EXCLUDED    EM_ERROR_CODE = 503
	TIEM_RESOURCE_NO_ENOUGH_DISK_AFTER_DISK_FILTER EM_ERROR_CODE = 504
	TIEM_RESOURCE_NO_ENOUGH_DISK_AFTER_HOST_FILTER EM_ERROR_CODE = 505
	TIEM_RESOURCE_NO_ENOUGH_PORT                   EM_ERROR_CODE = 506
	TIEM_RESOURCE_NOT_ALL_SUCCEED                  EM_ERROR_CODE = 507
	TIEM_RESOURCE_INVALID_STRATEGY                 EM_ERROR_CODE = 508
	TIEM_RESOURCE_INVAILD_RECYCLE_TYPE             EM_ERROR_CODE = 509
	TIEM_UPDATE_HOST_STATUS_FAIL                   EM_ERROR_CODE = 510
	TIEM_RESERVE_HOST_FAIL                         EM_ERROR_CODE = 511
	TIEM_RESOURCE_NO_STOCK                         EM_ERROR_CODE = 512
	TIEM_RESOURCE_GET_DISK_ID_FAIL                 EM_ERROR_CODE = 513
	TIEM_RESOURCE_TRAIT_NOT_FOUND                  EM_ERROR_CODE = 514
	TIEM_RESOURCE_INVALID_LOCATION                 EM_ERROR_CODE = 515
	TIEM_RESOURCE_INVALID_ARCH                     EM_ERROR_CODE = 516

	TIEM_MONITOR_NOT_FOUND EM_ERROR_CODE = 614

	TIEM_DEFAULT_PARAM_GROUP_NOT_DEL EM_ERROR_CODE = 20500

	TIEM_CHANGE_FEED_NOT_FOUND EM_ERROR_CODE = 701

	TIEM_CHANGE_FEED_DUPLICATE_ID           EM_ERROR_CODE = 702
	TIEM_CHANGE_FEED_STATUS_CONFLICT        EM_ERROR_CODE = 703
	TIEM_CHANGE_FEED_LOCK_EXPIRED           EM_ERROR_CODE = 704
	TIEM_CHANGE_FEED_UNSUPPORTED_DOWNSTREAM EM_ERROR_CODE = 705

	QueryProductComponentProperty = 706 //TODO
	QueryProductsScanRowError     = 707 //TODO
	QueryZoneScanRowError         = 708 //TODO
)

type ErrorCodeExplanation struct {
	code        EM_ERROR_CODE
	explanation string
	httpCode    int
}

func (t EM_ERROR_CODE) GetHttpCode() int {
	return explanationContainer[t].httpCode
}

func (t EM_ERROR_CODE) Equal(code int32) bool {
	return code == int32(t)
}

func (t EM_ERROR_CODE) Explain() string {
	return explanationContainer[t].explanation
}

var explanationContainer = map[EM_ERROR_CODE]ErrorCodeExplanation{
	TIEM_SUCCESS: {code: TIEM_SUCCESS, explanation: "succeed", httpCode: 200},

	// system error
	TIEM_METADB_SERVER_CALL_ERROR:  {TIEM_METADB_SERVER_CALL_ERROR, "call metadb-Server failed", 500},
	TIEM_CLUSTER_SERVER_CALL_ERROR: {TIEM_CLUSTER_SERVER_CALL_ERROR, "call cluster-Server failed", 500},
	TIEM_PARAMETER_INVALID:         {TIEM_PARAMETER_INVALID, "parameter is invalid", 400},
	TIEM_UNRECOGNIZED_ERROR:        {TIEM_UNRECOGNIZED_ERROR, "unrecognized error", 500},

	// cluster management
	TIEM_CLUSTER_NOT_FOUND:           {TIEM_CLUSTER_NOT_FOUND, "cluster not found", 404},
	TIEM_TAKEOVER_SSH_CONNECT_ERROR:  {TIEM_TAKEOVER_SSH_CONNECT_ERROR, "ssh connect failed", 500},
	TIEM_TAKEOVER_SFTP_ERROR:         {TIEM_TAKEOVER_SFTP_ERROR, "sftp failed", 500},
	TIEM_CLUSTER_RESOURCE_NOT_ENOUGH: {TIEM_CLUSTER_RESOURCE_NOT_ENOUGH, "no enough resource for cluster", 500},

	// dashboard && monitor
	TIEM_DASHBOARD_NOT_FOUND: {TIEM_DASHBOARD_NOT_FOUND, "dashboard is not found", 500},

	// workflow
	TIEM_WORKFLOW_CREATE_FAILED: {TIEM_WORKFLOW_CREATE_FAILED, "workflow create failed", 500},
	TIEM_WORKFLOW_QUERY_FAILED:  {TIEM_WORKFLOW_QUERY_FAILED, "workflow query failed", 500},
	TIEM_WORKFLOW_DETAIL_FAILED: {TIEM_WORKFLOW_DETAIL_FAILED, "workflow detail failed", 500},
	TIEM_WORKFLOW_START_FAILED:  {TIEM_WORKFLOW_START_FAILED, "workflow start failed", 500},

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

	// tenant
	TIEM_ACCOUNT_NOT_FOUND:       {TIEM_ACCOUNT_NOT_FOUND, "account not found", 500},
	TIEM_TENANT_NOT_FOUND:        {TIEM_TENANT_NOT_FOUND, "tenant not found", 500},
	TIEM_QUERY_PERMISSION_FAILED: {TIEM_QUERY_PERMISSION_FAILED, "query permission failed", 500},
	TIEM_ADD_TOKEN_FAILED:        {TIEM_ADD_TOKEN_FAILED, "add token failed", 500},
	TIEM_TOKEN_NOT_FOUND:         {TIEM_TOKEN_NOT_FOUND, "token not found", 500},

	// param group & cluster param
	TIEM_DEFAULT_PARAM_GROUP_NOT_DEL: {TIEM_DEFAULT_PARAM_GROUP_NOT_DEL, "The default param group cannot be deleted", 500},

	// change feed
	TIEM_CHANGE_FEED_NOT_FOUND:              {TIEM_CHANGE_FEED_NOT_FOUND, "change feed task not found", 404},
	TIEM_CHANGE_FEED_DUPLICATE_ID:           {TIEM_CHANGE_FEED_DUPLICATE_ID, "duplicate id", 500},
	TIEM_CHANGE_FEED_STATUS_CONFLICT:        {TIEM_CHANGE_FEED_STATUS_CONFLICT, "task status conflict", 409},
	TIEM_CHANGE_FEED_LOCK_EXPIRED:           {TIEM_CHANGE_FEED_LOCK_EXPIRED, "task status lock expired", 409},
	TIEM_CHANGE_FEED_UNSUPPORTED_DOWNSTREAM: {TIEM_CHANGE_FEED_UNSUPPORTED_DOWNSTREAM, "task downstream type not supported", 500},
}
