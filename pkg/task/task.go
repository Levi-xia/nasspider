package task

import (
	"nasspider/pkg/bo"
	"nasspider/pkg/constants"
	"nasspider/pkg/downloader"
	"nasspider/pkg/logger"
	"nasspider/pkg/provider"
	// "time"
)

func DoTask(p provider.Provider, d downloader.Downloader, tvTask bo.TVTask) error {
	var (
		err       error
		URLs      []string
		currentEp int
	)
	if tvTask.Status != int(constants.Doing) {
		return nil
	}
	if tvTask.TotalEp != 0 && tvTask.CurrentEp >= tvTask.TotalEp {
		// 更新status为finish Todo
		return nil
	}
	if URLs, currentEp, err = p.ParseURLs(tvTask.URL, tvTask.CurrentEp); err != nil {
		return err
	}
	if len(URLs) == 0 {
		logger.Logger.Info("未获取到更新的URLs, 跳过，等待下次执行")
		return nil
	}
	logger.Logger.Infof("URLs:%v\n", URLs)
	logger.Logger.Infof("currentEp:%d\n", currentEp)
	// for _, URL := range URLs {
	// 	// 发送下载任务
	// 	if err = downloader.CommitDownloadTask(d, downloader.Task{
	// 		URL:  URL,
	// 		Type: constants.DownloaderType(tvTask.Type),
	// 		Path: tvTask.DownloadPath,
	// 	}); err != nil {
	// 		return err
	// 	}
	// 	time.Sleep(time.Second * 1)
	// }

	// // 更新当前集数+追更状态
	// tvTask.CurrentEp = currentEp // Todo

	return nil
}
