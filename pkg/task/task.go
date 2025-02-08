package task

import (
	"nasspider/pkg/bo"
	"nasspider/pkg/constants"
	"nasspider/pkg/downloader"
	"nasspider/pkg/logger"
	"nasspider/pkg/provider"
	"time"
)

func DoTask(p provider.Provider, d downloader.Downloader, tvTask bo.TVTask) error {
	var (
		err       error
		URLs      []string
		currentEp int
	)
	defer func() {
		if err != nil {
			service.UpdateStatus(&bo.UpdateStatusRequest{
				ID:     tvTask.ID,
				Status: int(constants.Error),
			})
		}
	}()
	if tvTask.Status != int(constants.Doing) {
		return nil
	}
	if tvTask.TotalEp != 0 && tvTask.CurrentEp >= tvTask.TotalEp {
		if _, err := service.UpdateStatus(&bo.UpdateStatusRequest{
			ID:     tvTask.ID,
			Status: int(constants.Finish),
		}); err != nil {
			return err
		}
		return nil
	}
	if URLs, currentEp, err = p.ParseURLs(tvTask.URL, tvTask.CurrentEp); err != nil {
		return err
	}
	if len(URLs) == 0 {
		logger.Logger.Info("未获取到更新的URLs, 跳过，等待下次执行")
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
	if _, err := service.UpdateCurrentEp(&bo.UpdateCurrentEpRequest{
		ID:        tvTask.ID,
		CurrentEp: currentEp,
	}); err != nil {
		return err
	}
	if currentEp == tvTask.TotalEp {
		if _, err := service.UpdateStatus(&bo.UpdateStatusRequest{
			ID:     tvTask.ID,
			Status: constants.Finish,
		}); err != nil {
			return err
		}
	}
	return nil
}
