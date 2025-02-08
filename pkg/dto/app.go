package dto

import "nasspider/pkg/common"

type AddTaskRequest struct {
	URL          string `json:"url" form:"url" binding:"required"`
	Name         string `json:"name" form:"name" binding:"required"`
	TotalEp      int    `json:"total_ep" form:"total_ep" binding:"required"`
	CurrentEp    int    `json:"current_ep" form:"current_ep"`
	DownloadPath string `json:"download_path" form:"download_path" binding:"required"`
	Type         string `json:"type" form:"type" binding:"required"`
	Downloader   string `json:"downloader" form:"downloader" binding:"required"`
	Provider     string `json:"provider" form:"provider" binding:"required"`
}

type AddTaskResponse struct {
	ID int `json:"id"`
}

func (AddTaskRequest) GetMessages() common.ValidatorMessages {
	return common.ValidatorMessages{
		"url.required":           "URL不能为空",
		"name.required":          "名称不能为空",
		"total_ep.required":      "总集数不能为空",
		"download_path.required": "下载路径不能为空",
		"type.required":          "类型不能为空",
		"downloader.required":    "下载器不能为空",
		"provider.required":      "提供商不能为空",
	}
}

type EditTaskRequest struct {
	ID           int    `json:"id" form:"id" binding:"required"`
	URL          string `json:"url" form:"url" binding:"required"`
	Name         string `json:"name" form:"name" binding:"required"`
	TotalEp      int    `json:"total_ep" form:"total_ep" binding:"required"`
	CurrentEp    int    `json:"current_ep" form:"current_ep"`
	DownloadPath string `json:"download_path" form:"download_path" binding:"required"`
	Type         string `json:"type" form:"type" binding:"required"`
	Downloader   string `json:"downloader" form:"downloader" binding:"required"`
	Provider     string `json:"provider" form:"provider" binding:"required"`
	Status       int    `json:"status" form:"status"`
}

func (EditTaskRequest) GetMessages() common.ValidatorMessages {
	return common.ValidatorMessages{
		"id.required":           "ID不能为空",
		"url.required":          "URL不能为空",
		"name.required":         "名称不能为空",
		"total_ep.required":     "总集数不能为空",
		"download_path.required": "下载路径不能为空",
		"type.required":         "类型不能为空",
		"downloader.required":   "下载器不能为空",
		"provider.required":     "提供商不能为空",
	}
}

type EditTaskResponse struct {
	ID int `json:"id"`
}

type TriggerTaskRequest struct {
	ID int `json:"id" form:"id" binding:"required"`
}

type TriggerTaskResponse struct {
	ID int `json:"id"`
}

func (TriggerTaskRequest) GetMessages() common.ValidatorMessages {
	return common.ValidatorMessages{
		"id.required": "ID不能为空",
	}
}


type LoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (LoginRequest) GetMessages() common.ValidatorMessages {
	return common.ValidatorMessages{
		"username.required": "用户名不能为空",
		"password.required": "密码不能为空",
	}
}

type LoginResponse struct {
	Token string `json:"token"`
	ExpiresAt int `json:"expires_at"`
}