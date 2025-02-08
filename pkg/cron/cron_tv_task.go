package cron

import (
	"nasspider/pkg/bo"
	"nasspider/pkg/constants"
	"nasspider/pkg/logger"
	"nasspider/pkg/service"
	"nasspider/pkg/task"
	"nasspider/utils"
	"sync"
	"time"
)

// CronCommitExecuteTvTask 定时执行tv追更任务
var running bool = false

func CronCommitExecuteTvTask() {
	if running {
		logger.Logger.Infof("定时任务正在执行，跳过本次执行")
		return
	}
	running = true
	defer func() {
		running = false
	}()

	logger.Logger.Infof("%s 定时追更任务开始执行....", utils.FormatTime(time.Now()))

	page := 1
	pageSize := 100
	for {
		response, err := service.GetTaskList(&bo.GetTaskListRequest{
			Page:     page,
			PageSize: pageSize,
			StatusList: []int{
				int(constants.Doing),
			},
		})
		if err != nil {
			logger.Logger.Errorf("定时任务执行失败，获取任务列表失败:err: %v", err)
			break
		}
		tvTasks := response.List

		if len(tvTasks) == 0 {
			logger.Logger.Infof("无待执行的任务，等待下一次执行")
			break
		}

		var wg sync.WaitGroup
		for _, tvTask := range tvTasks {
			wg.Add(1)
			go func(t bo.TVTask) {
				defer wg.Done()
				defer func() {
					if err := recover(); err != nil {
						logger.Logger.Errorf("任务执行失败:%v", err)
					}
				}()
				task.DoTask(t)
			}(tvTask)
		}
		wg.Wait() // 等待所有任务完成

		if len(tvTasks) < pageSize {
			break
		}
		page++
	}
}
