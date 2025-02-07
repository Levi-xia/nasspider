package service

import (
	"nasspider/pkg/bo"
	"nasspider/pkg/common"
	"nasspider/utils"
)

type TvTask struct {
	ID           common.ID
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

func GetTaskList() ([]bo.TVTask, error) {
	return []bo.TVTask{}, nil
}

func AddTask() error {
	return nil
}

func EditTask() error {
	return nil
}

func UpdateCurrentEp() error {
	return nil
}

func UpdateStatus() error {
	return nil
}

func mo2Bo(m TvTask) (bo.TVTask, error) {
	return bo.TVTask{
		ID:           m.ID.ID,
		URL:          m.URL,
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
