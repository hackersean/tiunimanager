package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/pingcap-inc/tiem/library/framework"
	"github.com/pingcap-inc/tiem/library/knowledge"

	"github.com/pingcap-inc/tiem/micro-metadb/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	dbPb "github.com/pingcap-inc/tiem/micro-metadb/proto"
)

type FailureDomain int32

const (
	ROOT FailureDomain = iota
	DATACENTER
	ZONE
	RACK
	HOST
	DISK
)

func genDomainCodeByName(pre string, name string) string {
	return fmt.Sprintf("%s,%s", pre, name)
}

func GetDomainNameFromCode(failureDomain string) string {
	pos := strings.LastIndex(failureDomain, ",")
	return failureDomain[pos+1:]
}

func copyHostInfoFromReq(src *dbPb.DBHostInfoDTO, dst *models.Host) {
	dst.HostName = src.HostName
	dst.IP = src.Ip
	dst.UserName = src.UserName
	dst.Passwd = src.Passwd
	dst.OS = src.Os
	dst.Kernel = src.Kernel
	dst.CpuCores = int(src.CpuCores)
	dst.Memory = int(src.Memory)
	dst.Spec = src.Spec
	dst.Nic = src.Nic
	dst.DC = src.Dc
	dst.AZ = genDomainCodeByName(dst.DC, src.Az)
	dst.Rack = genDomainCodeByName(dst.AZ, src.Rack)
	dst.Status = int32(src.Status)
	dst.Purpose = src.Purpose
	for _, disk := range src.Disks {
		dst.Disks = append(dst.Disks, models.Disk{
			Name:     disk.Name,
			Path:     disk.Path,
			Status:   int32(disk.Status),
			Capacity: disk.Capacity,
		})
	}
}

func (handler *DBServiceHandler) AddHost(ctx context.Context, req *dbPb.DBAddHostRequest, rsp *dbPb.DBAddHostResponse) error {
	db := handler.Dao().Db()
	login := framework.Log()
	var host models.Host
	copyHostInfoFromReq(req.Host, &host)

	hostId, err := models.CreateHost(db, &host)
	rsp.Rs = new(dbPb.DBHostResponseStatus)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			rsp.Rs.Code = int32(st.Code())
			rsp.Rs.Message = st.Message()
		} else {
			rsp.Rs.Code = int32(codes.Internal)
			rsp.Rs.Message = fmt.Sprintf("failed to import host(%s) %s, %v", host.HostName, host.IP, err)
		}
		login.Warnln(rsp.Rs.Message)

		// return nil to use rsp
		return nil
	}
	rsp.HostId = hostId
	rsp.Rs.Code = int32(codes.OK)
	return nil
}

func (handler *DBServiceHandler) AddHostsInBatch(ctx context.Context, req *dbPb.DBAddHostsInBatchRequest, rsp *dbPb.DBAddHostsInBatchResponse) error {
	db := handler.Dao().Db()
	login := framework.Log()
	var hosts []*models.Host
	for _, v := range req.Hosts {
		var host models.Host
		copyHostInfoFromReq(v, &host)
		hosts = append(hosts, &host)
	}
	hostIds, err := models.CreateHostsInBatch(db, hosts)
	rsp.Rs = new(dbPb.DBHostResponseStatus)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			rsp.Rs.Code = int32(st.Code())
			rsp.Rs.Message = st.Message()
		} else {
			rsp.Rs.Code = int32(codes.Internal)
			rsp.Rs.Message = fmt.Sprintf("failed to import hosts, %v", err)
		}
		login.Warnln(rsp.Rs.Message)

		// return nil to use rsp
		return nil
	}
	rsp.HostIds = hostIds
	rsp.Rs.Code = int32(codes.OK)
	return nil
}

func (handler *DBServiceHandler) RemoveHost(ctx context.Context, req *dbPb.DBRemoveHostRequest, rsp *dbPb.DBRemoveHostResponse) error {
	db := handler.Dao().Db()
	login := framework.Log()
	hostId := req.HostId
	err := models.DeleteHost(db, hostId)
	rsp.Rs = new(dbPb.DBHostResponseStatus)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			rsp.Rs.Code = int32(st.Code())
			rsp.Rs.Message = st.Message()
		} else {
			rsp.Rs.Code = int32(codes.Internal)
			rsp.Rs.Message = fmt.Sprintf("failed to delete host(%s), %v", hostId, err)
		}
		login.Warnln(rsp.Rs.Message)

		// return nil to use rsp
		return nil
	}
	rsp.Rs.Code = int32(codes.OK)
	return nil
}

func (handler *DBServiceHandler) RemoveHostsInBatch(ctx context.Context, req *dbPb.DBRemoveHostsInBatchRequest, rsp *dbPb.DBRemoveHostsInBatchResponse) error {
	db := handler.Dao().Db()
	login := framework.Log()
	err := models.DeleteHostsInBatch(db, req.HostIds)
	rsp.Rs = new(dbPb.DBHostResponseStatus)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			rsp.Rs.Code = int32(st.Code())
			rsp.Rs.Message = st.Message()
		} else {
			rsp.Rs.Code = int32(codes.Internal)
			rsp.Rs.Message = fmt.Sprintf("failed to delete host in batch, %v", err)
		}
		login.Warnln(rsp.Rs.Message)

		// return nil to use rsp
		return nil
	}
	rsp.Rs.Code = int32(codes.OK)
	return nil
}

func copyHostInfoToRsp(src *models.Host, dst *dbPb.DBHostInfoDTO) {
	dst.HostId = src.ID
	dst.HostName = src.HostName
	dst.Ip = src.IP
	dst.Os = src.OS
	dst.Kernel = src.Kernel
	dst.CpuCores = int32(src.CpuCores)
	dst.Memory = int32(src.Memory)
	dst.Spec = src.Spec
	dst.Nic = src.Nic
	dst.Dc = src.DC
	dst.Az = GetDomainNameFromCode(src.AZ)
	dst.Rack = GetDomainNameFromCode(src.Rack)
	dst.Status = src.Status
	dst.Purpose = src.Purpose
	dst.CreateAt = src.CreatedAt.Unix()
	for _, disk := range src.Disks {
		dst.Disks = append(dst.Disks, &dbPb.DBDiskDTO{
			DiskId:   disk.ID,
			Name:     disk.Name,
			Path:     disk.Path,
			Capacity: disk.Capacity,
			Status:   disk.Status,
		})
	}
}

func (handler *DBServiceHandler) ListHost(ctx context.Context, req *dbPb.DBListHostsRequest, rsp *dbPb.DBListHostsResponse) error {
	db := handler.Dao().Db()
	login := framework.Log()
	var hostReq models.ListHostReq
	hostReq.Purpose = req.Purpose
	hostReq.Status = models.HostStatus(req.Status)
	hostReq.Limit = int(req.Page.PageSize)
	if req.Page.Page >= 1 {
		hostReq.Offset = (int(req.Page.Page) - 1) * int(req.Page.PageSize)
	} else {
		hostReq.Offset = 0
	}
	hosts, err := models.ListHosts(db, hostReq)
	rsp.Rs = new(dbPb.DBHostResponseStatus)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			rsp.Rs.Code = int32(st.Code())
			rsp.Rs.Message = st.Message()
		} else {
			rsp.Rs.Code = int32(codes.Internal)
			rsp.Rs.Message = fmt.Sprintf("failed to list hosts, %v", err)
		}
		login.Warnln(rsp.Rs.Message)

		// return nil to use rsp
		return nil
	}
	for _, v := range hosts {
		var host dbPb.DBHostInfoDTO
		copyHostInfoToRsp(&v, &host)
		rsp.HostList = append(rsp.HostList, &host)
	}
	rsp.Page = new(dbPb.DBHostPageDTO)
	rsp.Page.Page = req.Page.Page
	rsp.Page.PageSize = req.Page.PageSize
	rsp.Page.Total = int32(len(hosts))
	rsp.Rs.Code = int32(codes.OK)
	return nil
}
func (handler *DBServiceHandler) CheckDetails(ctx context.Context, req *dbPb.DBCheckDetailsRequest, rsp *dbPb.DBCheckDetailsResponse) error {
	db := handler.Dao().Db()
	login := framework.Log()
	host, err := models.FindHostById(db, req.HostId)
	rsp.Rs = new(dbPb.DBHostResponseStatus)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			rsp.Rs.Code = int32(st.Code())
			rsp.Rs.Message = st.Message()
		} else {
			rsp.Rs.Code = int32(codes.Internal)
			rsp.Rs.Message = fmt.Sprintf("failed to list hosts %s, %v", req.HostId, err)
		}
		login.Warnln(rsp.Rs.Message)

		// return nil to use rsp
		return nil
	}

	rsp.Details = new(dbPb.DBHostInfoDTO)
	copyHostInfoToRsp(host, rsp.Details)
	rsp.Rs.Code = int32(codes.OK)
	return nil
}

func copyAllocReq(component string, req models.AllocReqs, in []*dbPb.DBAllocationReq) {
	for _, eachReq := range in {
		if eachReq.Count == 0 {
			continue
		}
		req[component] = append(req[component], &models.HostAllocReq{
			FailureDomain: eachReq.FailureDomain,
			CpuCores:      int(eachReq.CpuCores),
			Memory:        int(eachReq.Memory),
			Count:         int(eachReq.Count),
		})
	}
}

func buildAllocRsp(componet string, req models.AllocRsps, out *[]*dbPb.DBAllocHostDTO) {
	for _, result := range req[componet] {
		*out = append(*out, &dbPb.DBAllocHostDTO{
			HostName: result.HostName,
			Ip:       result.Ip,
			UserName: result.UserName,
			Passwd:   result.Passwd,
			CpuCores: int32(result.CpuCores),
			Memory:   int32(result.Memory),
			Disk: &dbPb.DBDiskDTO{
				DiskId:   result.DiskId,
				Name:     result.DiskName,
				Path:     result.Path,
				Capacity: int32(result.Capacity),
				Status:   int32(models.DISK_AVAILABLE),
			},
		})
	}
}

func (handler *DBServiceHandler) AllocHosts(ctx context.Context, in *dbPb.DBAllocHostsRequest, out *dbPb.DBAllocHostsResponse) error {
	db := handler.Dao().Db()
	log := framework.Log()
	// Build up allocHosts request for model
	req := make(models.AllocReqs)
	copyAllocReq("PD", req, in.PdReq)
	copyAllocReq("TiDB", req, in.TidbReq)
	copyAllocReq("TiKV", req, in.TikvReq)

	resources, err := models.AllocHosts(db, req)
	out.Rs = new(dbPb.DBHostResponseStatus)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			out.Rs.Code = int32(st.Code())
			out.Rs.Message = st.Message()
		} else {
			out.Rs.Code = int32(codes.Internal)
			out.Rs.Message = fmt.Sprintf("db service receive alloc hosts error, %v", err)
		}
		log.Warnln(out.Rs.Message)

		// return nil to use rsp
		return nil
	}

	buildAllocRsp("PD", resources, &out.PdHosts)
	buildAllocRsp("TiDB", resources, &out.TidbHosts)
	buildAllocRsp("TiKV", resources, &out.TikvHosts)

	out.Rs.Code = int32(codes.OK)
	return nil
}

/*
func (handler *DBServiceHandler) PreAllocHosts(ctx context.Context, req *dbPb.DBPreAllocHostsRequest, rsp *dbPb.DBPreAllocHostsResponse) error {
	db := handler.Dao().Db()
	login := framework.GetRootLogger()
	login.Infof("db service receive alloc host in %s for %d x (%du%dg)", req.Req.FailureDomain, req.Req.Count, req.Req.CpuCores, req.Req.Memory)
	resources, err := models.PreAllocHosts(db, req.Req.FailureDomain, int(req.Req.Count), int(req.Req.CpuCores), int(req.Req.Memory))
	rsp.Rs = new(dbPb.DBHostResponseStatus)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			rsp.Rs.Code = int32(st.Code())
			rsp.Rs.Message = st.Message()
		} else {
			rsp.Rs.Code = int32(codes.Internal)
			rsp.Rs.Message = fmt.Sprintf("db service receive alloc host in %s for %d x (%du%dg) error, %v",
				req.Req.FailureDomain, req.Req.Count, req.Req.CpuCores, req.Req.Memory, err)
		}
		login.Warnln(rsp.Rs.Message)

		// return nil to use rsp
		return nil
	}

	if len(resources) < int(req.Req.Count) {
		errMsg := fmt.Sprintf("no enough host resources(%d/%d) in %s", len(resources), req.Req.Count, req.Req.FailureDomain)
		login.Errorln(errMsg)
		rsp.Rs.Code = int32(codes.ResourceExhausted)
		rsp.Rs.Message = errMsg
		return nil
	}
	for _, v := range resources {
		rsp.Results = append(rsp.Results, &dbPb.DBPreAllocation{
			FailureDomain: req.Req.FailureDomain,
			HostId:        v.HostId,
			DiskId:        v.Id,
			OriginCores:   int32(v.CpuCores),
			OriginMem:     int32(v.Memory),
			RequestCores:  req.Req.CpuCores,
			RequestMem:    req.Req.Memory,
			HostName:      v.HostName,
			Ip:            v.Ip,
			UserName:      v.UserName,
			Passwd:        v.Passwd,
			DiskName:      v.Name,
			DiskPath:      v.Path,
			DiskCap:       int32(v.Capacity),
		})
	}
	return nil
}

func (handler *DBServiceHandler) LockHosts(ctx context.Context, req *dbPb.DBLockHostsRequest, rsp *dbPb.DBLockHostsResponse) error {
	db := handler.Dao().Db()
	login := framework.GetRootLogger()
	var resources []models.ResourceLock
	for _, v := range req.Req {
		resources = append(resources, models.ResourceLock{
			HostId:       v.HostId,
			DiskId:       v.DiskId,
			OriginCores:  int(v.OriginCores),
			OriginMem:    int(v.OriginMem),
			RequestCores: int(v.RequestCores),
			RequestMem:   int(v.RequestMem),
		})
	}
	rsp.Rs = new(dbPb.DBHostResponseStatus)
	err := models.LockHosts(db, resources)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			rsp.Rs.Code = int32(st.Code())
			rsp.Rs.Message = st.Message()
		} else {
			rsp.Rs.Code = int32(codes.Internal)
			rsp.Rs.Message = fmt.Sprintf("lock hosts failed, err: %v", err)
		}
		login.Warnln(rsp.Rs.Message)

		// return nil to use rsp
		return nil
	}
	return nil
}
*/

func getFailureDomainByType(fd FailureDomain) (domain string, err error) {
	switch fd {
	case DATACENTER:
		domain = "dc"
	case ZONE:
		domain = "az"
	case RACK:
		domain = "rack"
	default:
		err = status.Errorf(codes.InvalidArgument, "%d is invalid domain type(1-DataCenter, 2-Zone, 3-Rack)", fd)
	}
	return
}

func (handler *DBServiceHandler) GetFailureDomain(ctx context.Context, req *dbPb.DBGetFailureDomainRequest, rsp *dbPb.DBGetFailureDomainResponse) error {

	login := framework.Log()
	domainType := req.FailureDomainType
	domain, err := getFailureDomainByType(FailureDomain(domainType))
	rsp.Rs = new(dbPb.DBHostResponseStatus)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			rsp.Rs.Code = int32(st.Code())
			rsp.Rs.Message = st.Message()
		} else {
			rsp.Rs.Code = int32(codes.Internal)
			rsp.Rs.Message = fmt.Sprintf("get failure domain resources failed, err: %v", err)
		}
		login.Warnln(rsp.Rs.Message)

		// return nil to use rsp
		return nil
	}

	db := handler.Dao().Db()
	resources, err := models.GetFailureDomain(db, domain)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			rsp.Rs.Code = int32(st.Code())
			rsp.Rs.Message = st.Message()
		} else {
			rsp.Rs.Code = int32(codes.Internal)
			rsp.Rs.Message = fmt.Sprintf("get failure domain resources failed, err: %v", err)
		}
		login.Warnln(rsp.Rs.Message)

		// return nil to use rsp
		return nil
	}
	for _, v := range resources {
		rsp.FdList = append(rsp.FdList, &dbPb.DBFailureDomainResource{
			FailureDomain: v.FailureDomain,
			Purpose:       v.Purpose,
			Spec:          knowledge.GenSpecCode(int32(v.CpuCores), int32(v.Memory)),
			Count:         int32(v.Count),
		})
	}
	return nil
}
