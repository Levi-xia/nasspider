package cron

import (
	"log"

	"github.com/robfig/cron/v3"
)

type crontab struct {
	spec   string
	cmd    func()
	enable bool
}

// 在这里加cron任务
var cronTabs = []crontab{
	// 每2小时整点执行一次
	{spec: "0 */2 * * *", cmd: CronCommitExecuteTvTask, enable: true},
}

func InitCron() {
	c := cron.New()
	for _, tab := range cronTabs {
		if !tab.enable {
			continue
		}
		entryId, err := c.AddFunc(tab.spec, tab.cmd)
		if err != nil {
			log.Fatalf("add cron func err: %v", err)
		}
		log.Printf("add cron func success entryId: %d", entryId)
	}
	c.Start()
}
