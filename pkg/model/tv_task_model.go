package model

import "nasspider/pkg/common"

type TvTask struct {
	common.ID    `gorm:"primary_key;auto_increment:10000000;column:id;comment:id"`
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