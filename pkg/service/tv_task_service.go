package service

import (
	"nasspider/pkg/bo"
	"nasspider/pkg/model"
	"nasspider/pkg/serctx"
	"nasspider/utils"
)


func GetTaskList(req *bo.GetTaskListRequest) (*bo.GetTaskListResponse, error) {
	var (
		tasks   []model.TvTask
		tvTasks []bo.TVTask
	)
	db := serctx.SerCtx.Db

	if len(req.StatusList) > 0 {
		db = db.Where("status in ?", req.StatusList)
	}
	if err := db.Order("id desc").Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Find(&tasks).Error; err != nil {
		return nil, err
	}
	for _, task := range tasks {
		if tvTask, err := mo2Bo(task); err != nil {
			return nil, err
		} else {
			tvTasks = append(tvTasks, tvTask)
		}
	}
	return &bo.GetTaskListResponse{
		List: tvTasks,
	}, nil
}

func CountTaskList(req *bo.CountTaskListRequest) (*bo.CountTaskListResponse, error) {
	var count int64
	db := serctx.SerCtx.Db
	if len(req.StatusList) > 0 {
		db = db.Where("status in?", req.StatusList)
	}
	if err := db.Model(&model.TvTask{}).Count(&count).Error; err != nil {
		return nil, err
	}
	return &bo.CountTaskListResponse{
		Count: count,
	}, nil
}
func GetTask(req *bo.GetTaskRequest) (*bo.GetTaskResponse, error) {
	var task model.TvTask
	if err := serctx.SerCtx.Db.Where("id = ?", req.ID).First(&task).Error; err != nil {
		return nil, err
	}
	if tvTask, err := mo2Bo(task); err != nil {
		return nil, err
	} else {
		return &bo.GetTaskResponse{
			TVTask: tvTask,
		}, nil
	}
}

func AddTask(req *bo.AddTaskRequest) (*bo.AddTaskResponse, error) {
	task := model.TvTask{
		Name:         req.Name,
		URL:          req.URL,
		TotalEp:      req.TotalEp,
		CurrentEp:    req.CurrentEp,
		Status:       req.Status,
		DownloadPath: req.DownloadPath,
		Type:         req.Type,
		Downloader:   req.Downloader,
		Provider:     req.Provider,
	}
	if err := serctx.SerCtx.Db.Create(&task).Error; err != nil {
		return nil, err
	}
	return &bo.AddTaskResponse{
		ID: task.ID.ID,
	}, nil
}

func EditTask(req *bo.EditTaskRequest) (*bo.EditTaskResponse, error) {
	if err := serctx.SerCtx.Db.Model(&model.TvTask{}).Where("id = ?", req.ID).Updates(map[string]interface{}{
		"name":          req.Name,
		"url":           req.URL,
		"total_ep":      req.TotalEp,
		"current_ep":    req.CurrentEp,
		"status":        req.Status,
		"download_path": req.DownloadPath,
		"type":          req.Type,
		"downloader":    req.Downloader,
		"provider":      req.Provider,
	}).Error; err != nil {
		return nil, err
	}
	return &bo.EditTaskResponse{
		ID: req.ID,
	}, nil
}

func UpdateCurrentEp(req *bo.UpdateCurrentEpRequest) (*bo.UpdateCurrentEpResponse, error) {
	if err := serctx.SerCtx.Db.Model(&model.TvTask{}).Where("id = ?", req.ID).Update("current_ep", req.CurrentEp).Error; err != nil {
		return nil, err
	}
	return &bo.UpdateCurrentEpResponse{
		ID: req.ID,
	}, nil
}

func UpdateStatus(req *bo.UpdateStatusRequest) (*bo.UpdateStatusResponse, error) {
	if err := serctx.SerCtx.Db.Model(&model.TvTask{}).Where("id =?", req.ID).Update("status", req.Status).Error; err != nil {
		return nil, err
	}
	return &bo.UpdateStatusResponse{
		ID: req.ID,
	}, nil
}

func mo2Bo(m model.TvTask) (bo.TVTask, error) {
	return bo.TVTask{
		ID:           m.ID.ID,
		URL:          m.URL,
		Name:         m.Name,
		Type:         m.Type,
		DownloadPath: m.DownloadPath,
		Status:       m.Status,
		CurrentEp:    m.CurrentEp,
		TotalEp:      m.TotalEp,
		Provider:     m.Provider,
		Downloader:   m.Downloader,
		CreatedAt:    utils.FormatTime(m.CreatedAt),
		UpdatedAt:    utils.FormatTime(m.UpdatedAt),
	}, nil
}
