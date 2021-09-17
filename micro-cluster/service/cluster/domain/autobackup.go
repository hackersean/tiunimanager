package domain

import (
	ctx "context"
	"github.com/pingcap-inc/tiem/library/client"
	proto "github.com/pingcap-inc/tiem/micro-cluster/proto"
	db "github.com/pingcap-inc/tiem/micro-metadb/proto"
	"github.com/robfig/cron"
	"time"
)

type AutoBackupManager struct {
	JobCron *cron.Cron
	JobSpec string
}

type autoBackupHandler struct {
}

func InitAutoBackupCronJob() {
	mgr := NewAutoBackupManager()
	go mgr.Start()
}

func NewAutoBackupManager() *AutoBackupManager{
	mgr := &AutoBackupManager{
		JobCron: cron.New(),
		JobSpec: "0 0 * * * *", // every integer hour
	}
	err := mgr.JobCron.AddJob(mgr.JobSpec, &autoBackupHandler{})
	if err != nil {
		getLogger().Fatalf("add auto backup cron job failed, %s", err.Error())
		return nil
	}

	return mgr
}

func (mgr *AutoBackupManager) Start() {
	time.Sleep(5 * time.Second) //wait db client ready
	mgr.JobCron.Start()
	defer mgr.JobCron.Stop()

	select {}
}

func (auto *autoBackupHandler) Run() {
	curWeekDay := time.Now().Weekday().String()
	curHour := time.Now().Hour()

	getLogger().Infof("begin AutoBackupHandler Run at WeekDay: %s, Hour: %d", curWeekDay, curHour)
	defer getLogger().Infof("end AutoBackupHandler Run")

	resp, err := client.DBClient.QueryBackupStrategyByTime(ctx.TODO(), &db.DBQueryBackupStrategyByTimeRequest{
		Weekday: curWeekDay,
		StartHour: uint32(curHour),
	})
	if err != nil {
		getLogger().Errorf("query backup strategy by weekday %s, hour: %d failed, %s", curWeekDay, curHour, err.Error())
		return
	}

	getLogger().Infof("WeekDay %s, Hour: %d need do auto backup for %d clusters", curWeekDay, curHour, len(resp.GetStrategys()))
	for _, strategy := range resp.GetStrategys() {
		go auto.doBackup(strategy)
	}
}

func (auto *autoBackupHandler) doBackup(straegy *db.DBBackupStrategyDTO) {
	getLogger().Infof("begin do auto backup for cluster %s", straegy.GetClusterId())
	defer getLogger().Infof("end do auto backup for cluster %s", straegy.GetClusterId())

	ope := &proto.OperatorDTO{
		Id: straegy.GetOperatorId(),
		TenantId: straegy.GetTenantId(),
	}
	_, err := Backup(ope, straegy.GetClusterId(), straegy.GetBackupRange(), straegy.GetBackupType(), BackupModeAuto, "")
	if err != nil {
		getLogger().Errorf("do backup for cluster %s failed, %s", straegy.GetClusterId(), err.Error())
		return
	}
}