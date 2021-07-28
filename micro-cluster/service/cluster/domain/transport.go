package domain

import (
	"fmt"
	"github.com/BurntSushi/toml"
	proto "github.com/pingcap/ticp/micro-cluster/proto"
	log "github.com/sirupsen/logrus"
	"time"
)

type TransportType uint32
const (
	TransportTypeExport TransportType = 0
	TransportTypeImport TransportType = 1
)

type TransportRecord struct {
	Id 				uint
	ClusterId 		string
	StartTime 		time.Time
	EndTime 		time.Time
	Operator 		Operator
	TransportType   TransportType
	FilePath 		string
}

type ImportInfo struct {
	ClusterId 		string
	UserName		string
	Password 		string
	FilePath 		string
}

type ExportInfo struct {
	ClusterId 		string
	UserName		string
	Password 		string
	FileType 		string
}

/*
	data import toml config for lighting
	https://docs.pingcap.com/zh/tidb/dev/deploy-tidb-lightning
 */
type DataImportConfig struct {
	Lighting 		LightingCfg			`toml:"lighting"`
	TikvImporter	TikvImporterCfg 	`toml:"tikv-importer"`
	MyDumper		MyDumperCfg 		`toml:"mydumper"`
	Tidb 			TidbCfg 			`toml:"tidb"`
}

type LightingCfg struct {
	Level			string	`toml:"level"`	//lighting log level
	File 			string 	`toml:"file"`	//lighting log path
}

type TikvImporterCfg struct {
	Backend 		string 	`toml:"backend"`		//backend mode: local/normal
	SortedKvDir		string 	`toml:"sorted-kv-dir"`	//temp store path
}

type MyDumperCfg struct {
	DataSourceDir	string	`toml:"data-source-dir"`	//import data filepath
}

type TidbCfg struct {
	Host 			string 	`toml:"host"`
	Port 			uint32	`toml:"port"`
	User 			string 	`toml:"user"`
	Password		string 	`toml:"password"`
	StatusPort		uint32 	`toml:"status-port"`	//table infomation from tidb status port
	PdAddr 			string 	`toml:"pd-addr"`
}

var contextDataTransportKey = "dataTransportInfo"

func ExportData(ope *proto.OperatorDTO, clusterId string, userName string, password string, fileType string) (uint32, error) {
	log.Infof("[domain] begin exportdata clusterId: %s, userName: %s, password: %s, fileType: %s", clusterId, userName, password, fileType)
	defer log.Infof("[domain] end exportdata")
	//todo: check operator
	operator := parseOperatorFromDTO(ope)
	log.Info(operator)
	clusterAggregation, err := ClusterRepo.Load(clusterId)

	info := &ExportInfo{
		ClusterId: clusterId,
		UserName: userName,
		Password: password,//todo: need encrypt
		FileType: fileType,
	}

	// Start the workflow
	flow, err := CreateFlowWork(clusterId, FlowExportData)
	if err != nil {
		return 0, nil
	}
	flow.AddContext(contextClusterKey, clusterAggregation)
	flow.AddContext(contextDataTransportKey, info)
	flow.Start()

	clusterAggregation.CurrentWorkFlow = flow.FlowWork
	ClusterRepo.Persist(clusterAggregation)
	return uint32(flow.FlowWork.Id), nil
}

func ImportData(ope *proto.OperatorDTO, clusterId string, userName string, password string, filepath string) (uint32, error) {
	log.Infof("[domain] begin importdata clusterId: %s, userName: %s, password: %s, datadIR: %s", clusterId, userName, password, filepath)
	defer log.Infof("[domain] end importdata")
	//todo: check operator
	operator := parseOperatorFromDTO(ope)
	log.Info(operator)
	clusterAggregation, err := ClusterRepo.Load(clusterId)

	info := &ImportInfo{
		ClusterId: clusterId,
		UserName: userName,
		Password: password,//todo: need encrypt
		FilePath: filepath,
	}

	// Start the workflow
	flow, err := CreateFlowWork(clusterId, FlowImportData)
	if err != nil {
		return 0, nil
	}
	flow.AddContext(contextClusterKey, clusterAggregation)
	flow.AddContext(contextDataTransportKey, info)
	flow.Start()

	clusterAggregation.CurrentWorkFlow = flow.FlowWork
	ClusterRepo.Persist(clusterAggregation)
	return uint32(flow.FlowWork.Id), nil
}

func convertTomlConfig(cluster *Cluster, info *ImportInfo) *DataImportConfig {
	//todo: replace config item
	config := &DataImportConfig{
		Lighting: LightingCfg{
			Level: "info",
			File: fmt.Sprintf("/user/local/tiem/datatransport/%s/import/tidb-lighting.log", cluster.Id),
		},
		TikvImporter: TikvImporterCfg{
			Backend: "local",
			SortedKvDir: "/mnt/ssd/sorted-kv-dir",
		},
		MyDumper: MyDumperCfg{
			DataSourceDir: info.FilePath,
		},
		Tidb: TidbCfg{
			Host: "127.0.0.1",
			Port: 4000,
			User: info.UserName,
			Password: info.Password,
			StatusPort: 10080,
			PdAddr: "127.0.0.1:2379",
		},
	}
	return config
}

func buildDataImportConfig(task *TaskEntity, context *FlowContext) bool {
	clusterAggregation := context.value(contextClusterKey).(ClusterAggregation)
	info := context.value(contextDataTransportKey).(ImportInfo)

	config := convertTomlConfig(clusterAggregation.Cluster, &info)
	filePath := fmt.Sprintf("/user/local/tiem/datatransport/%s/import/tidb-lighting.toml", clusterAggregation.Cluster.Id)
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		log.Errorf("[domain] decode data import toml config failed, %s", err.Error())
		return false
	}
	return true
}

func importDataToCluster(task *TaskEntity, context *FlowContext) bool {
	//todo: implement
	return true
}

func updateDataImportRecord(task *TaskEntity, context *FlowContext) bool {
	//todo: implement
	return true
}

func cleanImportTempfile(task *TaskEntity, context *FlowContext) bool {
	//todo: implement
	return true
}

func exportDataFromCluster(task *TaskEntity, context *FlowContext) bool {
	//todo: implement
	return true
}

func updateDataExportRecord(task *TaskEntity, context *FlowContext) bool {
	//todo: implement
	return true
}

func cleanDataImportRecord(task *TaskEntity, context *FlowContext) bool {
	//todo: implement
	return true
}

func cleanExportTempfile(task *TaskEntity, context *FlowContext) bool {
	//todo: implement
	return true
}