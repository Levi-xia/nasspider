package service

import (
	"nasspider/pkg/bo"
	"nasspider/pkg/common"
	"nasspider/utils"
)

type TvTask struct {
	ID           common.ID
	Name         string `gorm:"column:name;type:varchar(1024);not null;default:'';comment:名称"`
	URL          string `gorm:"column:url;type:varchar(1024);not null;default:'';comment:链接"`
	TotalEp      int    `gorm:"column:total_ep;type:int(11);not null;default:0;comment:总集数"`
	CurrentEp    int    `gorm:"column:current_ep;type:int(11);not null;default:0;comment:当前集数"`
	Status       int    `gorm:"column:status;type:int(11);not null;default:0;comment:状态"`
	DownloadPath string `gorm:"column:download_path;type:varchar(1024);not null;default:'';comment:下载路径"`
	Type         string `gorm:"column:type;type:varchar(32);not null;default:'';comment:类型"`
	Provider     string `gorm:"column:provider;type:varchar(32);not null;default:'';comment:提供商"`
	Downloader   string `gorm:"column:downloader;type:varchar(32);not null;default:'';comment:下载器"`
	common.Timestamps
	common.SoftDeletes
}

func (TvTask) TableName() string {
	return "tv_task"
}

func GetTaskList(req *bo.GetTaskListRequest) (*bo.GetTaskListResponse, error) {
	var (
		tasks   []TvTask
		tvTasks []bo.TVTask
	)
	if err := serctx.SerCtx.Db.Order("id desc").Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Find(&tasks).Error; err != nil {
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

func AddTask(req *bo.AddTaskRequest) (*bo.AddTaskResponse, error) {
	task := TvTask{
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
		ID: task.ID,
	}, nil
}

func EditTask(req *bo.EditTaskRequest) (*bo.EditTaskResponse, error) {
	task := TvTask{
		ID:           req.ID,
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
	if err := serctx.SerCtx.Db.Save(&task).Error; err!= nil {
		return nil, err
	}
	return &bo.EditTaskResponse{
		ID: req.ID
	}, nil
}

func UpdateCurrentEp(req *bo.UpdateCurrentEpRequest) (*bo.UpdateCurrentEpResponse, error) {
	if err := db.Model(&User{}).Where("id = ?", req.ID).Update("current_ep", req.CurrentEp).Error; err != nil {
		return nil, err
	}
	return &bo.UpdateCurrentEpResponse{
		ID: req.ID
	}, nil
}

func UpdateStatus(*bo.UpdateStatusRequest) (*bo.UpdateStatusResponse, error) {
	if err := db.Model(&User{}).Where("id =?", req.ID).Update("status", req.Status).Error; err!= nil {
		return nil, err
	}
	return &bo.UpdateStatusResponse{
		ID: req.ID
	}, nil
}

func mo2Bo(m TvTask) (bo.TVTask, error) {
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
