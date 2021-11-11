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

package secondparty

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime/debug"
)

func assert(b bool) {
	if b {
	} else {
		logger.Error("unexpected panic with stack trace:", string(debug.Stack()))
		panic("unexpected")
	}
}

func myPanic(v interface{}) {
	s := fmt.Sprint(v)
	logger.Errorf("panic: %s, with stack trace: %s", s, string(debug.Stack()))
	panic("unexpected" + s)
}

func newTmpFileWithContent(content []byte) (fileName string, err error) {
	tmpfile, err := ioutil.TempFile("", "tiem-topology-*.yaml")
	if err != nil {
		err = fmt.Errorf("fail to create temp file err: %s", err)
		return "", err
	}
	fileName = tmpfile.Name()
	var ct int
	ct, err = tmpfile.Write(content)
	if err != nil || ct != len(content) {
		tmpfile.Close()
		os.Remove(fileName)
		err = fmt.Errorf(fmt.Sprint("fail to write content to temp file ", fileName, "err:", err, "length of content:", "writed:", ct))
		return "", err
	}
	if err := tmpfile.Close(); err != nil {
		myPanic(fmt.Sprintln("fail to close temp file ", fileName))
	}
	return fileName, nil
}

func jsonMustMarshal(v interface{}) []byte {
	bs, err := json.Marshal(v)
	assert(err == nil)
	return bs
}