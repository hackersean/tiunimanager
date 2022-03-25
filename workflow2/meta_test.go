/******************************************************************************
 * Copyright (c)  2022 PingCAP, Inc.                                          *
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

package workflow2

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/pingcap-inc/tiem/common/constants"
	"github.com/pingcap-inc/tiem/library/framework"
	"github.com/pingcap-inc/tiem/models"
	"github.com/pingcap-inc/tiem/models/common"
	"github.com/pingcap-inc/tiem/models/workflow"
	"github.com/pingcap-inc/tiem/test/mockmodels/mockworkflow"
	"github.com/stretchr/testify/assert"
	"testing"
)

var doNode = func(node *workflow.WorkFlowNode, context *FlowContext) error {
	fmt.Println("doNodeName1")
	return nil
}

func TestWorkFlowMeta_GetData(t *testing.T) {
	meta := &WorkFlowMeta{
		Flow: &workflow.WorkFlow{
			Entity: common.Entity{
				ID:     "test",
				Status: constants.WorkFlowStatusInitializing,
			},
			Name: "test",
		},
		CurrentNode: &workflow.WorkFlowNode{
			Entity: common.Entity{
				ID:     "test",
				Status: constants.WorkFlowStatusInitializing,
			},
			Name: "test",
		},
		CurrentNodeDefine: &NodeDefine{
			FailEvent:  "",
			Executor:   doNode,
			ReturnType: SyncFuncNode,
		},
		Context: NewFlowContext(context.Background(), make(map[string]string)),
	}
	errSet := meta.Context.SetData("key", "value")
	var result string
	errGet := meta.Context.GetData("key", &result)
	assert.Nil(t, errSet)
	assert.Nil(t, errGet)
	assert.Equal(t, result, "value")
}

func TestWorkFlowMeta_Fail(t *testing.T) {
	meta := &WorkFlowMeta{
		Flow: &workflow.WorkFlow{},
		CurrentNodeDefine: &NodeDefine{
			FailEvent: "pause",
		},
	}
	meta.Fail()
	meta2 := &WorkFlowMeta{
		Flow: &workflow.WorkFlow{},
		CurrentNodeDefine: &NodeDefine{
			FailEvent: "",
		},
	}
	meta2.Fail()
}

func TestWorkFlowMeta_Execute_case1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFlowRW := mockworkflow.NewMockReaderWriter(ctrl)
	mockFlowRW.EXPECT().CreateWorkFlowNode(gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()
	mockFlowRW.EXPECT().UpdateWorkFlowDetail(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	mockFlowRW.EXPECT().UpdateWorkFlow(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	mockFlowRW.EXPECT().GetWorkFlow(gomock.Any(), gomock.Any()).Return(&workflow.WorkFlow{
		Entity: common.Entity{
			Status:   constants.WorkFlowStatusInitializing,
			TenantId: framework.GetTenantIDFromContext(context.TODO()),
			ID:       "testflowId",
		},
	}, nil).AnyTimes()
	models.SetWorkFlowReaderWriter(mockFlowRW)

	meta := &WorkFlowMeta{
		Flow: &workflow.WorkFlow{
			Entity: common.Entity{
				ID:     "test",
				Status: constants.WorkFlowStatusInitializing,
			},
			Name: "test",
		},
		CurrentNode: &workflow.WorkFlowNode{
			Entity: common.Entity{
				ID:     "test",
				Status: constants.WorkFlowStatusInitializing,
			},
			Name: "test",
		},
		CurrentNodeDefine: &NodeDefine{
			FailEvent:  "",
			Executor:   doNode,
			ReturnType: SyncFuncNode,
		},
		Context: NewFlowContext(context.Background(), make(map[string]string)),
	}
	meta.Execute()
}
