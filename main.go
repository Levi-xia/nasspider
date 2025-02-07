package main

import (
	"fmt"
	"nasspider/config"
	"nasspider/pkg/bo"
	"nasspider/pkg/constants"
	"nasspider/pkg/downloader"
	"nasspider/pkg/logger"
	"nasspider/pkg/provider"
	"nasspider/pkg/task"
)

func main() {

	// 初始化配置
	config.InitConfig()
	// 初始化日志
	logger.InitLog()

	var err error
	tvTasks := []bo.TVTask{
		{
			ID:           1,
			Name:         "白色橄榄树",
			URL:          "https://www.ddmp4.cc/html/KisTCQJJJJJQ.html",
			Type:         "magnet",
			DownloadPath: "/downloads/",
			Status:       0,
			CurrentEp:    18,
			TotalEp:      38,
			Provider:     "domp4",
			Downloader:   "thunder",
		},
		{
			ID:           2,
			Name:         "五福临门",
			URL:          "https://www.ddmp4.cc/html/p0KStZMMMMMZ.html",
			Type:         "magnet",
			DownloadPath: "/downloads/",
			Status:       0,
			CurrentEp:    23,
			TotalEp:      24,
			Provider:     "domp4",
			Downloader:   "thunder",
		},
		{
			ID:           3,
			Name:         "无所畏惧-2",
			URL:          "https://www.ddmp4.cc/html/SI4Qy4222224.html",
			Type:         "magnet",
			DownloadPath: "/downloads/",
			Status:       0,
			CurrentEp:    22,
			TotalEp:      40,
			Provider:     "domp4",
			Downloader:   "thunder",
		},
	}
	for _, tvTask := range tvTasks {
		logger.Logger.Infof("开始任务:%s\n", tvTask.Name)
		provider := provider.ProviderMap[constants.ProviderName(tvTask.Provider)]
		downloader := downloader.DownloaderMap[constants.DownloaderName(tvTask.Downloader)]

		err = task.DoTask(provider, downloader, tvTask)
		if err != nil {
			fmt.Printf("%v", err)
		}
	}
}
