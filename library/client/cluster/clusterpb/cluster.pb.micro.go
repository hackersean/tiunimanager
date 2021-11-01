// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: cluster.proto

package clusterpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for ClusterService service

func NewClusterServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ClusterService service

type ClusterService interface {
	// Cluster manager module
	CreateCluster(ctx context.Context, in *ClusterCreateReqDTO, opts ...client.CallOption) (*ClusterCreateRespDTO, error)
	QueryCluster(ctx context.Context, in *ClusterQueryReqDTO, opts ...client.CallOption) (*ClusterQueryRespDTO, error)
	DeleteCluster(ctx context.Context, in *ClusterDeleteReqDTO, opts ...client.CallOption) (*ClusterDeleteRespDTO, error)
	DetailCluster(ctx context.Context, in *ClusterDetailReqDTO, opts ...client.CallOption) (*ClusterDetailRespDTO, error)
	RestartCluster(ctx context.Context, in *ClusterRestartReqDTO, opts ...client.CallOption) (*ClusterRestartRespDTO, error)
	StopCluster(ctx context.Context, in *ClusterStopReqDTO, opts ...client.CallOption) (*ClusterStopRespDTO, error)
	TakeoverClusters(ctx context.Context, in *ClusterTakeoverReqDTO, opts ...client.CallOption) (*ClusterTakeoverRespDTO, error)
	ImportData(ctx context.Context, in *DataImportRequest, opts ...client.CallOption) (*DataImportResponse, error)
	ExportData(ctx context.Context, in *DataExportRequest, opts ...client.CallOption) (*DataExportResponse, error)
	DescribeDataTransport(ctx context.Context, in *DataTransportQueryRequest, opts ...client.CallOption) (*DataTransportQueryResponse, error)
	QueryBackupRecord(ctx context.Context, in *QueryBackupRequest, opts ...client.CallOption) (*QueryBackupResponse, error)
	CreateBackup(ctx context.Context, in *CreateBackupRequest, opts ...client.CallOption) (*CreateBackupResponse, error)
	RecoverCluster(ctx context.Context, in *RecoverRequest, opts ...client.CallOption) (*RecoverResponse, error)
	DeleteBackupRecord(ctx context.Context, in *DeleteBackupRequest, opts ...client.CallOption) (*DeleteBackupResponse, error)
	SaveBackupStrategy(ctx context.Context, in *SaveBackupStrategyRequest, opts ...client.CallOption) (*SaveBackupStrategyResponse, error)
	GetBackupStrategy(ctx context.Context, in *GetBackupStrategyRequest, opts ...client.CallOption) (*GetBackupStrategyResponse, error)
	QueryParameters(ctx context.Context, in *QueryClusterParametersRequest, opts ...client.CallOption) (*QueryClusterParametersResponse, error)
	SaveParameters(ctx context.Context, in *SaveClusterParametersRequest, opts ...client.CallOption) (*SaveClusterParametersResponse, error)
	DescribeDashboard(ctx context.Context, in *DescribeDashboardRequest, opts ...client.CallOption) (*DescribeDashboardResponse, error)
	DescribeMonitor(ctx context.Context, in *DescribeMonitorRequest, opts ...client.CallOption) (*DescribeMonitorResponse, error)
	// Auth manager module
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...client.CallOption) (*LogoutResponse, error)
	VerifyIdentity(ctx context.Context, in *VerifyIdentityRequest, opts ...client.CallOption) (*VerifyIdentityResponse, error)
	// host manager module
	ImportHost(ctx context.Context, in *ImportHostRequest, opts ...client.CallOption) (*ImportHostResponse, error)
	ImportHostsInBatch(ctx context.Context, in *ImportHostsInBatchRequest, opts ...client.CallOption) (*ImportHostsInBatchResponse, error)
	RemoveHost(ctx context.Context, in *RemoveHostRequest, opts ...client.CallOption) (*RemoveHostResponse, error)
	RemoveHostsInBatch(ctx context.Context, in *RemoveHostsInBatchRequest, opts ...client.CallOption) (*RemoveHostsInBatchResponse, error)
	ListHost(ctx context.Context, in *ListHostsRequest, opts ...client.CallOption) (*ListHostsResponse, error)
	CheckDetails(ctx context.Context, in *CheckDetailsRequest, opts ...client.CallOption) (*CheckDetailsResponse, error)
	AllocHosts(ctx context.Context, in *AllocHostsRequest, opts ...client.CallOption) (*AllocHostResponse, error)
	GetFailureDomain(ctx context.Context, in *GetFailureDomainRequest, opts ...client.CallOption) (*GetFailureDomainResponse, error)
	AllocResourcesInBatch(ctx context.Context, in *BatchAllocRequest, opts ...client.CallOption) (*BatchAllocResponse, error)
	RecycleResources(ctx context.Context, in *RecycleRequest, opts ...client.CallOption) (*RecycleResponse, error)
	UpdateHostStatus(ctx context.Context, in *UpdateHostStatusRequest, opts ...client.CallOption) (*UpdateHostStatusResponse, error)
	ReserveHost(ctx context.Context, in *ReserveHostRequest, opts ...client.CallOption) (*ReserveHostResponse, error)
	// task manager
	ListFlows(ctx context.Context, in *ListFlowsRequest, opts ...client.CallOption) (*ListFlowsResponse, error)
}

type clusterService struct {
	c    client.Client
	name string
}

func NewClusterService(name string, c client.Client) ClusterService {
	return &clusterService{
		c:    c,
		name: name,
	}
}

func (c *clusterService) CreateCluster(ctx context.Context, in *ClusterCreateReqDTO, opts ...client.CallOption) (*ClusterCreateRespDTO, error) {
	req := c.c.NewRequest(c.name, "ClusterService.CreateCluster", in)
	out := new(ClusterCreateRespDTO)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) QueryCluster(ctx context.Context, in *ClusterQueryReqDTO, opts ...client.CallOption) (*ClusterQueryRespDTO, error) {
	req := c.c.NewRequest(c.name, "ClusterService.QueryCluster", in)
	out := new(ClusterQueryRespDTO)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) DeleteCluster(ctx context.Context, in *ClusterDeleteReqDTO, opts ...client.CallOption) (*ClusterDeleteRespDTO, error) {
	req := c.c.NewRequest(c.name, "ClusterService.DeleteCluster", in)
	out := new(ClusterDeleteRespDTO)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) DetailCluster(ctx context.Context, in *ClusterDetailReqDTO, opts ...client.CallOption) (*ClusterDetailRespDTO, error) {
	req := c.c.NewRequest(c.name, "ClusterService.DetailCluster", in)
	out := new(ClusterDetailRespDTO)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) RestartCluster(ctx context.Context, in *ClusterRestartReqDTO, opts ...client.CallOption) (*ClusterRestartRespDTO, error) {
	req := c.c.NewRequest(c.name, "ClusterService.RestartCluster", in)
	out := new(ClusterRestartRespDTO)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) StopCluster(ctx context.Context, in *ClusterStopReqDTO, opts ...client.CallOption) (*ClusterStopRespDTO, error) {
	req := c.c.NewRequest(c.name, "ClusterService.StopCluster", in)
	out := new(ClusterStopRespDTO)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) TakeoverClusters(ctx context.Context, in *ClusterTakeoverReqDTO, opts ...client.CallOption) (*ClusterTakeoverRespDTO, error) {
	req := c.c.NewRequest(c.name, "ClusterService.TakeoverClusters", in)
	out := new(ClusterTakeoverRespDTO)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) ImportData(ctx context.Context, in *DataImportRequest, opts ...client.CallOption) (*DataImportResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.ImportData", in)
	out := new(DataImportResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) ExportData(ctx context.Context, in *DataExportRequest, opts ...client.CallOption) (*DataExportResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.ExportData", in)
	out := new(DataExportResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) DescribeDataTransport(ctx context.Context, in *DataTransportQueryRequest, opts ...client.CallOption) (*DataTransportQueryResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.DescribeDataTransport", in)
	out := new(DataTransportQueryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) QueryBackupRecord(ctx context.Context, in *QueryBackupRequest, opts ...client.CallOption) (*QueryBackupResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.QueryBackupRecord", in)
	out := new(QueryBackupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) CreateBackup(ctx context.Context, in *CreateBackupRequest, opts ...client.CallOption) (*CreateBackupResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.CreateBackup", in)
	out := new(CreateBackupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) RecoverCluster(ctx context.Context, in *RecoverRequest, opts ...client.CallOption) (*RecoverResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.RecoverCluster", in)
	out := new(RecoverResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) DeleteBackupRecord(ctx context.Context, in *DeleteBackupRequest, opts ...client.CallOption) (*DeleteBackupResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.DeleteBackupRecord", in)
	out := new(DeleteBackupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) SaveBackupStrategy(ctx context.Context, in *SaveBackupStrategyRequest, opts ...client.CallOption) (*SaveBackupStrategyResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.SaveBackupStrategy", in)
	out := new(SaveBackupStrategyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) GetBackupStrategy(ctx context.Context, in *GetBackupStrategyRequest, opts ...client.CallOption) (*GetBackupStrategyResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.GetBackupStrategy", in)
	out := new(GetBackupStrategyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) QueryParameters(ctx context.Context, in *QueryClusterParametersRequest, opts ...client.CallOption) (*QueryClusterParametersResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.QueryParameters", in)
	out := new(QueryClusterParametersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) SaveParameters(ctx context.Context, in *SaveClusterParametersRequest, opts ...client.CallOption) (*SaveClusterParametersResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.SaveParameters", in)
	out := new(SaveClusterParametersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) DescribeDashboard(ctx context.Context, in *DescribeDashboardRequest, opts ...client.CallOption) (*DescribeDashboardResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.DescribeDashboard", in)
	out := new(DescribeDashboardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) DescribeMonitor(ctx context.Context, in *DescribeMonitorRequest, opts ...client.CallOption) (*DescribeMonitorResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.DescribeMonitor", in)
	out := new(DescribeMonitorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) Logout(ctx context.Context, in *LogoutRequest, opts ...client.CallOption) (*LogoutResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.Logout", in)
	out := new(LogoutResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) VerifyIdentity(ctx context.Context, in *VerifyIdentityRequest, opts ...client.CallOption) (*VerifyIdentityResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.VerifyIdentity", in)
	out := new(VerifyIdentityResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) ImportHost(ctx context.Context, in *ImportHostRequest, opts ...client.CallOption) (*ImportHostResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.ImportHost", in)
	out := new(ImportHostResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) ImportHostsInBatch(ctx context.Context, in *ImportHostsInBatchRequest, opts ...client.CallOption) (*ImportHostsInBatchResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.ImportHostsInBatch", in)
	out := new(ImportHostsInBatchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) RemoveHost(ctx context.Context, in *RemoveHostRequest, opts ...client.CallOption) (*RemoveHostResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.RemoveHost", in)
	out := new(RemoveHostResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) RemoveHostsInBatch(ctx context.Context, in *RemoveHostsInBatchRequest, opts ...client.CallOption) (*RemoveHostsInBatchResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.RemoveHostsInBatch", in)
	out := new(RemoveHostsInBatchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) ListHost(ctx context.Context, in *ListHostsRequest, opts ...client.CallOption) (*ListHostsResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.ListHost", in)
	out := new(ListHostsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) CheckDetails(ctx context.Context, in *CheckDetailsRequest, opts ...client.CallOption) (*CheckDetailsResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.CheckDetails", in)
	out := new(CheckDetailsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) AllocHosts(ctx context.Context, in *AllocHostsRequest, opts ...client.CallOption) (*AllocHostResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.AllocHosts", in)
	out := new(AllocHostResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) GetFailureDomain(ctx context.Context, in *GetFailureDomainRequest, opts ...client.CallOption) (*GetFailureDomainResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.GetFailureDomain", in)
	out := new(GetFailureDomainResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) AllocResourcesInBatch(ctx context.Context, in *BatchAllocRequest, opts ...client.CallOption) (*BatchAllocResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.AllocResourcesInBatch", in)
	out := new(BatchAllocResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) RecycleResources(ctx context.Context, in *RecycleRequest, opts ...client.CallOption) (*RecycleResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.RecycleResources", in)
	out := new(RecycleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) UpdateHostStatus(ctx context.Context, in *UpdateHostStatusRequest, opts ...client.CallOption) (*UpdateHostStatusResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.UpdateHostStatus", in)
	out := new(UpdateHostStatusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) ReserveHost(ctx context.Context, in *ReserveHostRequest, opts ...client.CallOption) (*ReserveHostResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.ReserveHost", in)
	out := new(ReserveHostResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterService) ListFlows(ctx context.Context, in *ListFlowsRequest, opts ...client.CallOption) (*ListFlowsResponse, error) {
	req := c.c.NewRequest(c.name, "ClusterService.ListFlows", in)
	out := new(ListFlowsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ClusterService service

type ClusterServiceHandler interface {
	// Cluster manager module
	CreateCluster(context.Context, *ClusterCreateReqDTO, *ClusterCreateRespDTO) error
	QueryCluster(context.Context, *ClusterQueryReqDTO, *ClusterQueryRespDTO) error
	DeleteCluster(context.Context, *ClusterDeleteReqDTO, *ClusterDeleteRespDTO) error
	DetailCluster(context.Context, *ClusterDetailReqDTO, *ClusterDetailRespDTO) error
	RestartCluster(context.Context, *ClusterRestartReqDTO, *ClusterRestartRespDTO) error
	StopCluster(context.Context, *ClusterStopReqDTO, *ClusterStopRespDTO) error
	TakeoverClusters(context.Context, *ClusterTakeoverReqDTO, *ClusterTakeoverRespDTO) error
	ImportData(context.Context, *DataImportRequest, *DataImportResponse) error
	ExportData(context.Context, *DataExportRequest, *DataExportResponse) error
	DescribeDataTransport(context.Context, *DataTransportQueryRequest, *DataTransportQueryResponse) error
	QueryBackupRecord(context.Context, *QueryBackupRequest, *QueryBackupResponse) error
	CreateBackup(context.Context, *CreateBackupRequest, *CreateBackupResponse) error
	RecoverCluster(context.Context, *RecoverRequest, *RecoverResponse) error
	DeleteBackupRecord(context.Context, *DeleteBackupRequest, *DeleteBackupResponse) error
	SaveBackupStrategy(context.Context, *SaveBackupStrategyRequest, *SaveBackupStrategyResponse) error
	GetBackupStrategy(context.Context, *GetBackupStrategyRequest, *GetBackupStrategyResponse) error
	QueryParameters(context.Context, *QueryClusterParametersRequest, *QueryClusterParametersResponse) error
	SaveParameters(context.Context, *SaveClusterParametersRequest, *SaveClusterParametersResponse) error
	DescribeDashboard(context.Context, *DescribeDashboardRequest, *DescribeDashboardResponse) error
	DescribeMonitor(context.Context, *DescribeMonitorRequest, *DescribeMonitorResponse) error
	// Auth manager module
	Login(context.Context, *LoginRequest, *LoginResponse) error
	Logout(context.Context, *LogoutRequest, *LogoutResponse) error
	VerifyIdentity(context.Context, *VerifyIdentityRequest, *VerifyIdentityResponse) error
	// host manager module
	ImportHost(context.Context, *ImportHostRequest, *ImportHostResponse) error
	ImportHostsInBatch(context.Context, *ImportHostsInBatchRequest, *ImportHostsInBatchResponse) error
	RemoveHost(context.Context, *RemoveHostRequest, *RemoveHostResponse) error
	RemoveHostsInBatch(context.Context, *RemoveHostsInBatchRequest, *RemoveHostsInBatchResponse) error
	ListHost(context.Context, *ListHostsRequest, *ListHostsResponse) error
	CheckDetails(context.Context, *CheckDetailsRequest, *CheckDetailsResponse) error
	AllocHosts(context.Context, *AllocHostsRequest, *AllocHostResponse) error
	GetFailureDomain(context.Context, *GetFailureDomainRequest, *GetFailureDomainResponse) error
	AllocResourcesInBatch(context.Context, *BatchAllocRequest, *BatchAllocResponse) error
	RecycleResources(context.Context, *RecycleRequest, *RecycleResponse) error
	UpdateHostStatus(context.Context, *UpdateHostStatusRequest, *UpdateHostStatusResponse) error
	ReserveHost(context.Context, *ReserveHostRequest, *ReserveHostResponse) error
	// task manager
	ListFlows(context.Context, *ListFlowsRequest, *ListFlowsResponse) error
}

func RegisterClusterServiceHandler(s server.Server, hdlr ClusterServiceHandler, opts ...server.HandlerOption) error {
	type clusterService interface {
		CreateCluster(ctx context.Context, in *ClusterCreateReqDTO, out *ClusterCreateRespDTO) error
		QueryCluster(ctx context.Context, in *ClusterQueryReqDTO, out *ClusterQueryRespDTO) error
		DeleteCluster(ctx context.Context, in *ClusterDeleteReqDTO, out *ClusterDeleteRespDTO) error
		DetailCluster(ctx context.Context, in *ClusterDetailReqDTO, out *ClusterDetailRespDTO) error
		RestartCluster(ctx context.Context, in *ClusterRestartReqDTO, out *ClusterRestartRespDTO) error
		StopCluster(ctx context.Context, in *ClusterStopReqDTO, out *ClusterStopRespDTO) error
		TakeoverClusters(ctx context.Context, in *ClusterTakeoverReqDTO, out *ClusterTakeoverRespDTO) error
		ImportData(ctx context.Context, in *DataImportRequest, out *DataImportResponse) error
		ExportData(ctx context.Context, in *DataExportRequest, out *DataExportResponse) error
		DescribeDataTransport(ctx context.Context, in *DataTransportQueryRequest, out *DataTransportQueryResponse) error
		QueryBackupRecord(ctx context.Context, in *QueryBackupRequest, out *QueryBackupResponse) error
		CreateBackup(ctx context.Context, in *CreateBackupRequest, out *CreateBackupResponse) error
		RecoverCluster(ctx context.Context, in *RecoverRequest, out *RecoverResponse) error
		DeleteBackupRecord(ctx context.Context, in *DeleteBackupRequest, out *DeleteBackupResponse) error
		SaveBackupStrategy(ctx context.Context, in *SaveBackupStrategyRequest, out *SaveBackupStrategyResponse) error
		GetBackupStrategy(ctx context.Context, in *GetBackupStrategyRequest, out *GetBackupStrategyResponse) error
		QueryParameters(ctx context.Context, in *QueryClusterParametersRequest, out *QueryClusterParametersResponse) error
		SaveParameters(ctx context.Context, in *SaveClusterParametersRequest, out *SaveClusterParametersResponse) error
		DescribeDashboard(ctx context.Context, in *DescribeDashboardRequest, out *DescribeDashboardResponse) error
		DescribeMonitor(ctx context.Context, in *DescribeMonitorRequest, out *DescribeMonitorResponse) error
		Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error
		Logout(ctx context.Context, in *LogoutRequest, out *LogoutResponse) error
		VerifyIdentity(ctx context.Context, in *VerifyIdentityRequest, out *VerifyIdentityResponse) error
		ImportHost(ctx context.Context, in *ImportHostRequest, out *ImportHostResponse) error
		ImportHostsInBatch(ctx context.Context, in *ImportHostsInBatchRequest, out *ImportHostsInBatchResponse) error
		RemoveHost(ctx context.Context, in *RemoveHostRequest, out *RemoveHostResponse) error
		RemoveHostsInBatch(ctx context.Context, in *RemoveHostsInBatchRequest, out *RemoveHostsInBatchResponse) error
		ListHost(ctx context.Context, in *ListHostsRequest, out *ListHostsResponse) error
		CheckDetails(ctx context.Context, in *CheckDetailsRequest, out *CheckDetailsResponse) error
		AllocHosts(ctx context.Context, in *AllocHostsRequest, out *AllocHostResponse) error
		GetFailureDomain(ctx context.Context, in *GetFailureDomainRequest, out *GetFailureDomainResponse) error
		AllocResourcesInBatch(ctx context.Context, in *BatchAllocRequest, out *BatchAllocResponse) error
		RecycleResources(ctx context.Context, in *RecycleRequest, out *RecycleResponse) error
		UpdateHostStatus(ctx context.Context, in *UpdateHostStatusRequest, out *UpdateHostStatusResponse) error
		ReserveHost(ctx context.Context, in *ReserveHostRequest, out *ReserveHostResponse) error
		ListFlows(ctx context.Context, in *ListFlowsRequest, out *ListFlowsResponse) error
	}
	type ClusterService struct {
		clusterService
	}
	h := &clusterServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ClusterService{h}, opts...))
}

type clusterServiceHandler struct {
	ClusterServiceHandler
}

func (h *clusterServiceHandler) CreateCluster(ctx context.Context, in *ClusterCreateReqDTO, out *ClusterCreateRespDTO) error {
	return h.ClusterServiceHandler.CreateCluster(ctx, in, out)
}

func (h *clusterServiceHandler) QueryCluster(ctx context.Context, in *ClusterQueryReqDTO, out *ClusterQueryRespDTO) error {
	return h.ClusterServiceHandler.QueryCluster(ctx, in, out)
}

func (h *clusterServiceHandler) DeleteCluster(ctx context.Context, in *ClusterDeleteReqDTO, out *ClusterDeleteRespDTO) error {
	return h.ClusterServiceHandler.DeleteCluster(ctx, in, out)
}

func (h *clusterServiceHandler) DetailCluster(ctx context.Context, in *ClusterDetailReqDTO, out *ClusterDetailRespDTO) error {
	return h.ClusterServiceHandler.DetailCluster(ctx, in, out)
}

func (h *clusterServiceHandler) RestartCluster(ctx context.Context, in *ClusterRestartReqDTO, out *ClusterRestartRespDTO) error {
	return h.ClusterServiceHandler.RestartCluster(ctx, in, out)
}

func (h *clusterServiceHandler) StopCluster(ctx context.Context, in *ClusterStopReqDTO, out *ClusterStopRespDTO) error {
	return h.ClusterServiceHandler.StopCluster(ctx, in, out)
}

func (h *clusterServiceHandler) TakeoverClusters(ctx context.Context, in *ClusterTakeoverReqDTO, out *ClusterTakeoverRespDTO) error {
	return h.ClusterServiceHandler.TakeoverClusters(ctx, in, out)
}

func (h *clusterServiceHandler) ImportData(ctx context.Context, in *DataImportRequest, out *DataImportResponse) error {
	return h.ClusterServiceHandler.ImportData(ctx, in, out)
}

func (h *clusterServiceHandler) ExportData(ctx context.Context, in *DataExportRequest, out *DataExportResponse) error {
	return h.ClusterServiceHandler.ExportData(ctx, in, out)
}

func (h *clusterServiceHandler) DescribeDataTransport(ctx context.Context, in *DataTransportQueryRequest, out *DataTransportQueryResponse) error {
	return h.ClusterServiceHandler.DescribeDataTransport(ctx, in, out)
}

func (h *clusterServiceHandler) QueryBackupRecord(ctx context.Context, in *QueryBackupRequest, out *QueryBackupResponse) error {
	return h.ClusterServiceHandler.QueryBackupRecord(ctx, in, out)
}

func (h *clusterServiceHandler) CreateBackup(ctx context.Context, in *CreateBackupRequest, out *CreateBackupResponse) error {
	return h.ClusterServiceHandler.CreateBackup(ctx, in, out)
}

func (h *clusterServiceHandler) RecoverCluster(ctx context.Context, in *RecoverRequest, out *RecoverResponse) error {
	return h.ClusterServiceHandler.RecoverCluster(ctx, in, out)
}

func (h *clusterServiceHandler) DeleteBackupRecord(ctx context.Context, in *DeleteBackupRequest, out *DeleteBackupResponse) error {
	return h.ClusterServiceHandler.DeleteBackupRecord(ctx, in, out)
}

func (h *clusterServiceHandler) SaveBackupStrategy(ctx context.Context, in *SaveBackupStrategyRequest, out *SaveBackupStrategyResponse) error {
	return h.ClusterServiceHandler.SaveBackupStrategy(ctx, in, out)
}

func (h *clusterServiceHandler) GetBackupStrategy(ctx context.Context, in *GetBackupStrategyRequest, out *GetBackupStrategyResponse) error {
	return h.ClusterServiceHandler.GetBackupStrategy(ctx, in, out)
}

func (h *clusterServiceHandler) QueryParameters(ctx context.Context, in *QueryClusterParametersRequest, out *QueryClusterParametersResponse) error {
	return h.ClusterServiceHandler.QueryParameters(ctx, in, out)
}

func (h *clusterServiceHandler) SaveParameters(ctx context.Context, in *SaveClusterParametersRequest, out *SaveClusterParametersResponse) error {
	return h.ClusterServiceHandler.SaveParameters(ctx, in, out)
}

func (h *clusterServiceHandler) DescribeDashboard(ctx context.Context, in *DescribeDashboardRequest, out *DescribeDashboardResponse) error {
	return h.ClusterServiceHandler.DescribeDashboard(ctx, in, out)
}

func (h *clusterServiceHandler) DescribeMonitor(ctx context.Context, in *DescribeMonitorRequest, out *DescribeMonitorResponse) error {
	return h.ClusterServiceHandler.DescribeMonitor(ctx, in, out)
}

func (h *clusterServiceHandler) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.ClusterServiceHandler.Login(ctx, in, out)
}

func (h *clusterServiceHandler) Logout(ctx context.Context, in *LogoutRequest, out *LogoutResponse) error {
	return h.ClusterServiceHandler.Logout(ctx, in, out)
}

func (h *clusterServiceHandler) VerifyIdentity(ctx context.Context, in *VerifyIdentityRequest, out *VerifyIdentityResponse) error {
	return h.ClusterServiceHandler.VerifyIdentity(ctx, in, out)
}

func (h *clusterServiceHandler) ImportHost(ctx context.Context, in *ImportHostRequest, out *ImportHostResponse) error {
	return h.ClusterServiceHandler.ImportHost(ctx, in, out)
}

func (h *clusterServiceHandler) ImportHostsInBatch(ctx context.Context, in *ImportHostsInBatchRequest, out *ImportHostsInBatchResponse) error {
	return h.ClusterServiceHandler.ImportHostsInBatch(ctx, in, out)
}

func (h *clusterServiceHandler) RemoveHost(ctx context.Context, in *RemoveHostRequest, out *RemoveHostResponse) error {
	return h.ClusterServiceHandler.RemoveHost(ctx, in, out)
}

func (h *clusterServiceHandler) RemoveHostsInBatch(ctx context.Context, in *RemoveHostsInBatchRequest, out *RemoveHostsInBatchResponse) error {
	return h.ClusterServiceHandler.RemoveHostsInBatch(ctx, in, out)
}

func (h *clusterServiceHandler) ListHost(ctx context.Context, in *ListHostsRequest, out *ListHostsResponse) error {
	return h.ClusterServiceHandler.ListHost(ctx, in, out)
}

func (h *clusterServiceHandler) CheckDetails(ctx context.Context, in *CheckDetailsRequest, out *CheckDetailsResponse) error {
	return h.ClusterServiceHandler.CheckDetails(ctx, in, out)
}

func (h *clusterServiceHandler) AllocHosts(ctx context.Context, in *AllocHostsRequest, out *AllocHostResponse) error {
	return h.ClusterServiceHandler.AllocHosts(ctx, in, out)
}

func (h *clusterServiceHandler) GetFailureDomain(ctx context.Context, in *GetFailureDomainRequest, out *GetFailureDomainResponse) error {
	return h.ClusterServiceHandler.GetFailureDomain(ctx, in, out)
}

func (h *clusterServiceHandler) AllocResourcesInBatch(ctx context.Context, in *BatchAllocRequest, out *BatchAllocResponse) error {
	return h.ClusterServiceHandler.AllocResourcesInBatch(ctx, in, out)
}

func (h *clusterServiceHandler) RecycleResources(ctx context.Context, in *RecycleRequest, out *RecycleResponse) error {
	return h.ClusterServiceHandler.RecycleResources(ctx, in, out)
}

func (h *clusterServiceHandler) UpdateHostStatus(ctx context.Context, in *UpdateHostStatusRequest, out *UpdateHostStatusResponse) error {
	return h.ClusterServiceHandler.UpdateHostStatus(ctx, in, out)
}

func (h *clusterServiceHandler) ReserveHost(ctx context.Context, in *ReserveHostRequest, out *ReserveHostResponse) error {
	return h.ClusterServiceHandler.ReserveHost(ctx, in, out)
}

func (h *clusterServiceHandler) ListFlows(ctx context.Context, in *ListFlowsRequest, out *ListFlowsResponse) error {
	return h.ClusterServiceHandler.ListFlows(ctx, in, out)
}
