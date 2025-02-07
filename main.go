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
	tvTask := bo.TVTask{
		ID:           1,
		URL:          "https://www.ddmp4.cc/html/KisTCQJJJJJQ.html",
		Type:         "magnet",
		DownloadPath: "/downloads/",
		Status:       0,
		CurrentEp:    18,
		TotalEp:      38,
		Provider:     "domp4",
		Downloader:   "thunder",
	}
	provider := provider.ProviderMap[constants.ProviderName(tvTask.Provider)]
	downloader := downloader.DownloaderMap[constants.DownloaderName(tvTask.Downloader)]

	err = task.DoTask(provider, downloader, tvTask)
	if err != nil {
		fmt.Printf("%v", err)
	}
}
