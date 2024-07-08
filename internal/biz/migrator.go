package biz

import (
	"context"
	"fmt"
	"github.com/BitofferHub/pkg/middlewares/log"
	"github.com/BitofferHub/xtimer/internal/conf"
	"github.com/BitofferHub/xtimer/internal/constant"
	"github.com/BitofferHub/xtimer/internal/utils"
	"time"
)

// xtimerUseCase is a User usecase.
type MigratorUseCase struct {
	confData  *conf.Data
	timerRepo TimerRepo
	taskRepo  TimerTaskRepo
	taskCache TaskCache
}

// NewUserUseCase new a User usecase.
func NewMigratorUseCase(confData *conf.Data, timerRepo TimerRepo, taskRepo TimerTaskRepo, taskCache TaskCache) *MigratorUseCase {
	return &MigratorUseCase{
		confData:  confData,
		timerRepo: timerRepo,
		taskRepo:  taskRepo,
		taskCache: taskCache,
	}
}

func (uc *MigratorUseCase) BatchMigratorTimer(ctx context.Context) error {
	timers, err := uc.timerRepo.FindByStatus(ctx, constant.Enabled.ToInt())
	if err != nil {
		log.ErrorContextf(ctx, "批量迁移Timer失败，查询数据库失败，err:: %v", err)
		return err
	}
	for _, timer := range timers {
		err = uc.MigratorTimer(ctx, timer)
		if err != nil {
			log.ErrorContextf(ctx, "批量迁移，迁移单个Timer失败，timerId:%s", timer.TimerId)
		}
		time.Sleep(5 * time.Second)
	}
	return nil
}

func (uc *MigratorUseCase) MigratorTimer(ctx context.Context, timer *Timer) error {
	// 校验状态
	if timer.Status != constant.Enabled.ToInt() {
		return fmt.Errorf("Timer非Unable状态，迁移失败，timerId:: %d", timer.TimerId)
	}

	// 3.取得批量的执行时机
	start := time.Now()
	end := start.Add(2 * time.Duration(uc.confData.GetMigrator().MigrateStepMinutes) * time.Minute)
	executeTimes, err := utils.NextsBefore(timer.Cron, end)
	if err != nil {
		log.ErrorContextf(ctx, "get executeTimes failed, err: %v", err)
		return err
	}

	// 执行时机加入数据库
	tasks := timer.BatchTasksFromTimer(executeTimes)
	// 基于 timer_id + run_timer 唯一键，保证任务不被重复插入
	if err := uc.taskRepo.BatchSave(ctx, tasks); err != nil {
		log.ErrorContextf(ctx, "DB存储tasks失败: %v", err)
		return err
	}

	// 执行时机加入 redis 跳表
	if err := uc.taskCache.BatchCreateTasks(ctx, tasks); err != nil {
		log.ErrorContextf(ctx, "Zset存储tasks失败: %v", err)
		return err
	}
	return nil
}
