package service

import (
	pb "github.com/BitofferHub/proto_center/api/xtimer/v1"
	"github.com/BitofferHub/xtimer/internal/biz"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewXTimerService)

type XTimerService struct {
	pb.UnimplementedXTimerServer
	timerUC     *biz.XtimerUseCase
	schedulerUC *biz.SchedulerUseCase
	migratorUC  *biz.MigratorUseCase
}

func NewXTimerService(timerUC *biz.XtimerUseCase, schedulerUC *biz.SchedulerUseCase, migratorUC *biz.MigratorUseCase) *XTimerService {
	return &XTimerService{timerUC: timerUC, schedulerUC: schedulerUC, migratorUC: migratorUC}
}
