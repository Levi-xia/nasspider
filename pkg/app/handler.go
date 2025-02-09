package app

import (
	"nasspider/config"
	"nasspider/pkg/bo"
	"nasspider/pkg/common"
	"nasspider/pkg/constants"
	"nasspider/pkg/dto"
	"nasspider/pkg/logger"
	"nasspider/pkg/service"
	"nasspider/pkg/task"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// GetTaskList 获取任务列表
func GetTaskList(c *gin.Context) {
	resp := &common.Result{}
	req := &dto.GetTaskListRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusOK, resp.Error(common.ParamError, common.GetErrorMsg(req, err)))
		return
	}
	taskListResp, err := service.GetTaskList(&bo.GetTaskListRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		c.JSON(http.StatusOK, resp.Error(common.BusinessError, err.Error()))
		return
	}
	totalResp, err := service.CountTaskList(&bo.CountTaskListRequest{})
	if err != nil {
		c.JSON(http.StatusOK, resp.Error(common.BusinessError, err.Error()))
		return
	}

	var dtoTVTasks []dto.TVTask
	for _, tvTask := range taskListResp.List {
		dtoTVTasks = append(dtoTVTasks, dto.TVTask{
			ID:           tvTask.ID,
			Name:         tvTask.Name,
			URL:          tvTask.URL,
			Provider:     tvTask.Provider,
			Downloader:   tvTask.Downloader,
			DownloadPath: tvTask.DownloadPath,
			Type:         tvTask.Type,
			TotalEp:      tvTask.TotalEp,
			CurrentEp:    tvTask.CurrentEp,
			Status:       tvTask.Status,
			CreatedAt:   tvTask.CreatedAt,
			UpdatedAt:   tvTask.UpdatedAt,
			StatusDesc:  constants.TaskStatusMap[constants.TaskStatus(tvTask.Status)],
		})
	}
	c.JSON(http.StatusOK, resp.Success(&dto.GetTaskListResponse{
		List:  dtoTVTasks,
		Total: totalResp.Count,
	}))
}

// Login 用户登录
func SubmitLogin(c *gin.Context) {
	resp := &common.Result{}
	req := &dto.LoginRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusOK, resp.Error(common.ParamError, common.GetErrorMsg(req, err)))
		return
	}
	passportConf := config.Conf.Passport

	if req.Username != passportConf.Username || req.Password != passportConf.Password {
		c.JSON(http.StatusOK, resp.Error(common.BusinessError, "用户名或密码错误"))
		return
	}
	token, err := (&common.JwtService{}).CreateToken(common.AppGuardName, "1")
	if err != nil {
		c.JSON(http.StatusOK, resp.Error(common.BusinessError, err.Error()))
		return
	}
	// 设置cookies
	c.SetCookie("AT", token.AccessToken, int(token.ExpiresIn), "/", "", false, true)

	c.JSON(http.StatusOK, resp.Success(&dto.LoginResponse{
		Token:     token.AccessToken,
		ExpiresAt: token.ExpiresIn,
	}))
}

// EditTaskRequest 发起修改任务
func EditTask(c *gin.Context) {
	resp := &common.Result{}
	req := &dto.EditTaskRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusOK, resp.Error(common.ParamError, common.GetErrorMsg(req, err)))
		return
	}
	response, err := service.EditTask(&bo.EditTaskRequest{
		ID:           req.ID,
		Name:         req.Name,
		URL:          req.URL,
		Provider:     req.Provider,
		Downloader:   req.Downloader,
		DownloadPath: req.DownloadPath,
		Type:         req.Type,
		TotalEp:      req.TotalEp,
		CurrentEp:    req.CurrentEp,
		Status:       req.Status,
	})
	if err != nil {
		c.JSON(http.StatusOK, resp.Error(common.BusinessError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(&dto.EditTaskResponse{
		ID: response.ID,
	}))
}

// AddTaskRequest 发起添加任务
func AddTask(c *gin.Context) {
	resp := &common.Result{}
	req := &dto.AddTaskRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusOK, resp.Error(common.ParamError, common.GetErrorMsg(req, err)))
		return
	}
	response, err := service.AddTask(&bo.AddTaskRequest{
		Name:         req.Name,
		URL:          req.URL,
		Provider:     req.Provider,
		Downloader:   req.Downloader,
		DownloadPath: req.DownloadPath,
		Type:         req.Type,
		TotalEp:      req.TotalEp,
		CurrentEp:    req.CurrentEp,
	})
	if err != nil {
		c.JSON(http.StatusOK, resp.Error(common.BusinessError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(&dto.AddTaskResponse{
		ID: response.ID,
	}))
}

// 触发任务
func TriggerTask(c *gin.Context) {
	resp := &common.Result{}
	req := &dto.TriggerTaskRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusOK, resp.Error(common.ParamError, common.GetErrorMsg(req, err)))
		return
	}
	taskResp, err := service.GetTask(&bo.GetTaskRequest{
		ID: req.ID,
	})
	if err != nil {
		c.JSON(http.StatusOK, resp.Error(common.BusinessError, err.Error()))
		return
	}
	tvTask := taskResp.TVTask

	go func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Logger.Errorf("任务失败:%v", err)
			}
		}()
		task.DoTask(tvTask, false)
	}()

	c.JSON(http.StatusOK, resp.Success(&dto.TriggerTaskResponse{
		ID: tvTask.ID,
	}))
}
