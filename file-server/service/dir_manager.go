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

package service

import (
	"fmt"
	"github.com/pingcap-inc/tiem/library/common"
	"path/filepath"
)

var DirMgr DirManager

type DirManager struct {
}

func InitDirManager() *DirManager {
	DirMgr = DirManager{}
	return &DirMgr
}

func (mgr *DirManager) GetImportPath(clusterId string) (string, error) {
	importAbsDir, err := filepath.Abs(common.DefaultImportDir) //todo: get from config center
	if err != nil {
		getLogger().Errorf("import dir %s is not vaild", common.DefaultImportDir)
		return "", fmt.Errorf("import dir %s is not vaild", common.DefaultImportDir)
	}
	return fmt.Sprintf("%s/%s/temp", importAbsDir, clusterId), nil
}