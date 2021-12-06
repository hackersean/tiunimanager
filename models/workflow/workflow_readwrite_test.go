/******************************************************************************
 * Copyright (c)  2021 PingCAP, Inc.                                          *
 * Licensed under the Apache License, Version 2.0 (the "License");            *
 * you may not use this file except in compliance with the License.           *
 * You may obtain a copy of the License at                                    *
 *                                                                            *
 * http://www.apache.org/licenses/LICENSE-2.0                                 *
 *                                                                            *
 *  Unless required by applicable law or agreed to in writing, software       *
 *  distributed under the License is distributed on an "AS IS" BASIS,         *
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  *
 *  See the License for the specific language governing permissions and       *
 *  limitations under the License.                                            *
 ******************************************************************************/

package workflow

import (
	"context"
	"github.com/pingcap-inc/tiem/models/common"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var rw *WorkFlowReadWrite

func TestFlowReadWrite_CreateWorkFlow(t *testing.T) {
	flow := &WorkFlow{
		Entities: common.Entities{
			TenantId: "tenantId",
			Status:   "FlowInitStatus",
		},
		Name:  "flowName",
		BizID: "clusterId",
	}
	_, err := rw.CreateWorkFlow(context.TODO(), flow)
	assert.NoError(t, err)
}

func TestFlowReadWrite_GetWorkFlow(t *testing.T) {
	flow := &WorkFlow{
		Entities: common.Entities{
			TenantId: "tenantId",
			Status:   "FlowInitStatus",
		},
		Name:  "flowName",
		BizID: "clusterId",
	}
	flowCreate, errCreate := rw.CreateWorkFlow(context.TODO(), flow)
	assert.NoError(t, errCreate)

	flowGet, errGet := rw.GetWorkFlow(context.TODO(), flowCreate.ID)
	assert.Equal(t, flowCreate.ID, flowGet.ID)
	assert.NoError(t, errGet)
}

func TestFlowReadWrite_QueryWorkFlows(t *testing.T) {
	flow := &WorkFlow{
		Entities: common.Entities{
			TenantId: "tenantId",
			Status:   "FlowInitStatus",
		},
		Name:  "flowNameQuery",
		BizID: "clusterId",
	}
	flowCreate, errCreate := rw.CreateWorkFlow(context.TODO(), flow)
	assert.NoError(t, errCreate)

	flowQuery, total, errQuery := rw.QueryWorkFlows(context.TODO(), "", "flowNameQuery", "", 0, 10)
	assert.Equal(t, int64(1), total)
	assert.Equal(t, flowCreate.ID, flowQuery[0].ID)
	assert.NoError(t, errQuery)
}

func TestFlowReadWrite_UpdateWorkFlowStatus(t *testing.T) {
	flow := &WorkFlow{
		Entities: common.Entities{
			TenantId: "tenantId",
			Status:   "FlowInitStatus",
		},
		Name:  "flowName",
		BizID: "clusterId",
	}
	flowCreate, errCreate := rw.CreateWorkFlow(context.TODO(), flow)
	assert.NoError(t, errCreate)

	errUpdate := rw.UpdateWorkFlowStatus(context.TODO(), flowCreate.ID, "FlowEndStatus")
	assert.NoError(t, errUpdate)

	flowGet, errGet := rw.GetWorkFlow(context.TODO(), flowCreate.ID)
	assert.Equal(t, flowCreate.ID, flowGet.ID)
	assert.Equal(t, flowGet.Status, "FlowEndStatus")
	assert.NoError(t, errGet)
}

func TestFlowReadWrite_DetailWorkFlow(t *testing.T) {
	flow := &WorkFlow{
		Entities: common.Entities{
			TenantId: "tenantId",
			Status:   "FlowInitStatus",
		},
		Name:  "flowName",
		BizID: "clusterId",
	}
	flowCreate, errFlowCreate := rw.CreateWorkFlow(context.TODO(), flow)
	assert.NoError(t, errFlowCreate)

	node := &WorkFlowNode{
		Entities: common.Entities{
			TenantId: "tenantId",
			Status:   "NodeInitStatus",
		},
		Parameters: "nodeParam",
		ParentID:   flowCreate.ID,
		Name:       "nodeName",
		ReturnType: string(SyncFuncNode),
		Result:     "success",
		StartTime:  time.Now(),
		EndTime:    time.Now(),
	}
	nodeCreate, errNodeCreate := rw.CreateWorkFlowNode(context.TODO(), node)
	assert.NoError(t, errNodeCreate)

	flowDetail, nodeDetail, errDetail := rw.QueryDetailWorkFlow(context.TODO(), flowCreate.ID)
	assert.NoError(t, errDetail)
	assert.Equal(t, flowCreate.ID, flowDetail.ID)
	assert.Equal(t, nodeCreate.ID, nodeDetail[0].ID)
}

func TestFlowReadWrite_CreateWorkFlowNode(t *testing.T) {
	node := &WorkFlowNode{
		Entities: common.Entities{
			TenantId: "tenantId",
			Status:   "NodeInitStatus",
		},
		Parameters: "nodeParam",
		ParentID:   "flowId",
		Name:       "nodeName",
		ReturnType: string(SyncFuncNode),
		Result:     "success",
		StartTime:  time.Now(),
		EndTime:    time.Now(),
	}
	_, err := rw.CreateWorkFlowNode(context.TODO(), node)
	assert.NoError(t, err)
}

func TestFlowReadWrite_UpdateWorkFlowNode(t *testing.T) {
	node := &WorkFlowNode{
		Entities: common.Entities{
			TenantId: "tenantId",
			Status:   "NodeInitStatus",
		},
		Parameters: "nodeParam",
		ParentID:   "flowId",
		Name:       "nodeName",
		ReturnType: string(SyncFuncNode),
		Result:     "init",
		StartTime:  time.Now(),
		EndTime:    time.Now(),
	}
	nodeCreate, errCreate := rw.CreateWorkFlowNode(context.TODO(), node)
	assert.NoError(t, errCreate)

	node = &WorkFlowNode{
		Entities: common.Entities{
			ID:       nodeCreate.ID,
			TenantId: "tenantId",
			Status:   "NodeEndStatus",
		},
		Parameters: "nodeParam",
		ParentID:   "flowId",
		Name:       "nodeName",
		ReturnType: string(SyncFuncNode),
		Result:     "success",
		StartTime:  time.Now(),
		EndTime:    time.Now(),
	}
	errUpdate := rw.UpdateWorkFlowNode(context.TODO(), node)
	assert.NoError(t, errUpdate)

	nodeQuery, errQuery := rw.GetWorkFlowNode(context.TODO(), nodeCreate.ID)
	assert.NoError(t, errQuery)
	assert.Equal(t, nodeQuery.ID, nodeCreate.ID)
	assert.Equal(t, "NodeEndStatus", nodeQuery.Status)
	assert.Equal(t, "success", nodeQuery.Result)
}

func TestFlowReadWrite_UpdateWorkFlowDetail(t *testing.T) {
	flow := &WorkFlow{
		Entities: common.Entities{
			TenantId: "tenantId",
			Status:   "FlowInitStatus",
		},
		Name:  "flowName",
		BizID: "clusterId",
	}
	flowCreate, errFlowCreate := rw.CreateWorkFlow(context.TODO(), flow)
	assert.NoError(t, errFlowCreate)

	node := &WorkFlowNode{
		Entities: common.Entities{
			TenantId: "tenantId",
			Status:   "NodeInitStatus",
		},
		Parameters: "nodeParam",
		ParentID:   flowCreate.ID,
		Name:       "nodeName",
		ReturnType: string(SyncFuncNode),
		Result:     "init",
		StartTime:  time.Now(),
		EndTime:    time.Now(),
	}
	nodeCreate, errNodeCreate := rw.CreateWorkFlowNode(context.TODO(), node)
	assert.NoError(t, errNodeCreate)

	nodes := make([]*WorkFlowNode, 1)
	nodeCreate.Result = "success"
	nodeCreate.Status = "NodeEndStatus"
	nodes[0] = nodeCreate
	flowCreate.Status = "FlowEndStatus"
	errUpdate := rw.UpdateWorkFlowDetail(context.TODO(), flowCreate, nodes)
	assert.NoError(t, errUpdate)

	flowQuery, nodeQuery, errQuery := rw.QueryDetailWorkFlow(context.TODO(), flowCreate.ID)
	assert.NoError(t, errQuery)
	assert.Equal(t, "FlowEndStatus", flowQuery.Status)
	assert.Equal(t, "NodeEndStatus", nodeQuery[0].Status)
	assert.Equal(t, "success", nodeQuery[0].Result)
}
