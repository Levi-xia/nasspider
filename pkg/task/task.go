package task

import (
	"fmt"
	"nasspider/pkg/bo"
	"nasspider/pkg/constants"
	"nasspider/pkg/downloader"
	"nasspider/pkg/logger"
	"nasspider/pkg/provider"
	"nasspider/pkg/service"
	"time"
)

func DoTask(tvTask bo.TVTask) error {
	logger.Logger.Infof("开始执行任务:%v", tvTask.Name)
	var (
		err       error
		URLs      []string
		currentEp int
	)
	defer func() {
		if err != nil {
			service.UpdateStatus(&bo.UpdateStatusRequest{ID: tvTask.ID, Status: int(constants.Error)})
			logger.Logger.Errorf("任务失败:%v", err)
		}
	}()
	if tvTask.Status != int(constants.Doing) {
		logger.Logger.Infof("任务状态不是正在执行, 跳过，等待下次执行")
		return nil
	}
	if tvTask.TotalEp != 0 && tvTask.CurrentEp >= tvTask.TotalEp {
		if _, err := service.UpdateStatus(&bo.UpdateStatusRequest{ID: tvTask.ID, Status: int(constants.Finish)}); err != nil {
			return err
		}
		return nil
	}
	p := provider.ProviderMap[constants.ProviderName(tvTask.Provider)]
	d := downloader.DownloaderMap[constants.DownloaderName(tvTask.Downloader)]

	if p == nil || d == nil {
		return fmt.Errorf("provider or downloader not found")
	}

	if URLs, currentEp, err = p.ParseURLs(tvTask.URL, tvTask.CurrentEp); err != nil {
		return err
	}
	if len(URLs) == 0 {
		logger.Logger.Info("未获取到更新的URLs, 跳过，等待下次执行")
		return nil
	}
	logger.Logger.Infof("获取到更新的URLs（共%d条）,当前已更新至%d集，开始下载...", len(URLs), currentEp)

	for index, URL := range URLs {
		// 发送下载任务
		if err = downloader.CommitDownloadTask(d, downloader.Task{
			URL:  URL,
			Type: constants.DownloaderType(tvTask.Type),
			Path: tvTask.DownloadPath,
		}); err != nil {
			return err
		}
		logger.Logger.Infof("任务%d发送成功", index)
		time.Sleep(time.Second * 1)
	}
	if _, err := service.UpdateCurrentEp(&bo.UpdateCurrentEpRequest{ID: tvTask.ID, CurrentEp: currentEp}); err != nil {
		return err
	}
	if currentEp == tvTask.TotalEp {
		if _, err := service.UpdateStatus(&bo.UpdateStatusRequest{ID: tvTask.ID, Status: int(constants.Finish)}); err != nil {
			return err
		}
	}
	return nil
}
