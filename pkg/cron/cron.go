package cron

import (
	"github.com/robfig/cron/v3"
	"nasspider/config"
	"nasspider/pkg/constants"
	"nasspider/pkg/logger"
)

type crontab struct {
	spec   string
	cmd    func()
	enable bool
}

func getCronTasks() []crontab {
	return []crontab{
		{
			enable: config.GetConf(config.Conf.Cron.TvTask.Enabled, constants.ENV_CRON_TV_TASK_ENABLED),
			spec:   config.GetConf(config.Conf.Cron.TvTask.Spec, constants.ENV_CRON_TV_TASK_SPEC),
			cmd:    CronCommitExecuteTvTask,
		},
	}
}

func InitCron() {
	c := cron.New()
	cronTabs := getCronTasks()
	for _, tab := range cronTabs {
		if !tab.enable {
			continue
		}
		entryId, err := c.AddFunc(tab.spec, tab.cmd)
		if err != nil {
			logger.Logger.Fatalf("add cron func err: %v", err)
		}
		logger.Logger.Infof("add cron func success entryId: %d", entryId)
	}
	c.Start()
}
