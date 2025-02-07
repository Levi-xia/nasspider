package task

import (
	"nasspider/pkg/bo"
	"nasspider/pkg/constants"
	"nasspider/pkg/downloader"
	"nasspider/pkg/provider"
	"time"
)

func DoTask(p provider.Provider, d downloader.Downloader, tvTask bo.TVTask) error {
	var (
		err  error
		URLs []string
	)
	if tvTask.Status != int(constants.Doing) {
		return nil
	}
	if tvTask.TotalEp != 0 && tvTask.CurrentEp >= tvTask.TotalEp {
		// 更新status为finish Todo
		return nil
	}
	if URLs, err = p.ParseURLs(tvTask.URL, tvTask.CurrentEp, tvTask.Xpath); err != nil {
		return err
	}
	if len(URLs) == 0 {
		return nil
	}
	for _, URL := range URLs {
		// 发送下载任务
		if err = downloader.CommitDownloadTask(d, downloader.Task{
			URL:  URL,
			Type: constants.DownloaderType(tvTask.Type),
			Path: tvTask.DownloadPath,
		}); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}
