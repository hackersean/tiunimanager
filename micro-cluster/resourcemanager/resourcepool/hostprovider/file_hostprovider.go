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
 ******************************************************************************/

package hostprovider

import (
	"context"

	"github.com/pingcap-inc/tiem/common/structs"

	"github.com/pingcap-inc/tiem/models/resource"
	"github.com/pingcap-inc/tiem/models/resource/resourcepool"
)

type FileHostProvider struct {
	rw resource.ResourceReaderWriter
}

func GetFileHostProvider() HostProvider {
	hostProvider := new(FileHostProvider)
	hostProvider.rw = resource.NewGormChangeFeedReadWrite()
	return hostProvider
}

func (p *FileHostProvider) SetResourceReaderWriter(rw resource.ResourceReaderWriter) {
	p.rw = rw
}

func (p *FileHostProvider) ImportHosts(ctx context.Context, hosts []structs.HostInfo) (hostIds []string, err error) {
	var dbModelHosts []resourcepool.Host
	for _, host := range hosts {
		var dbHost resourcepool.Host
		err = dbHost.ConstructFromHostInfo(&host)
		if err != nil {
			return nil, err
		}
		dbModelHosts = append(dbModelHosts, dbHost)
	}
	return p.rw.Create(ctx, dbModelHosts)
}

func (p *FileHostProvider) DeleteHosts(ctx context.Context, hostIds []string) (err error) {
	return nil
}

func (p *FileHostProvider) Query(ctx context.Context, filter structs.HostFilter) (hosts []*structs.HostInfo, err error) {
	return nil, nil
}

func (p *FileHostProvider) UpdateHostStatus(ctx context.Context, hostId []string, status string) (err error) {
	return nil
}

func (p *FileHostProvider) UpdateHostReserved(ctx context.Context, hostId []string, reserved bool) (err error) {
	return nil
}

func (p *FileHostProvider) GetHierarchy(ctx context.Context, filter structs.HostFilter, level int32, depth int32) (root *structs.HierarchyTreeNode, err error) {
	return nil, nil
}

func (p *FileHostProvider) GetStocks(ctx context.Context, location structs.Location, hostFilter structs.HostFilter, diskFilter structs.DiskFilter) (stocks *structs.Stocks, err error) {
	return nil, nil
}
