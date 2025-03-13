package app

import (
	"nasspider/config"
	"nasspider/pkg/bo"
	"nasspider/pkg/common"
	"nasspider/pkg/constants"
	"nasspider/pkg/downloader"
	"nasspider/pkg/dto"
	"nasspider/pkg/logger"
	"nasspider/pkg/service"
	"nasspider/pkg/task"
	"net/http"
	"strings"
	"time"

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
	username := config.GetConf(passportConf.Username, constants.ENV_ADMIN_USERNAME)
	password := config.GetConf(passportConf.Password, constants.ENV_ADMIN_PASSWORD)

	if req.Username != username || req.Password != password {
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

// AddDownloadTask 添加下载任务
func AddDownloadTask(c *gin.Context) {
	resp := &common.Result{}
	req := &dto.AddDownloadTaskRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusOK, resp.Error(common.ParamError, common.GetErrorMsg(req, err)))
		return
	}
	// URLs按照换行切分成string数组
	URLs := strings.Split(req.URLs, "\n")
	if len(URLs) == 0 {
		c.JSON(http.StatusOK, resp.Error(common.ParamError, "URLs不能为空"))
		return
	}
	d := downloader.DownloaderMap[constants.DownloaderName(req.Downloader)]
	if d == nil {
		c.JSON(http.StatusOK, resp.Error(common.BusinessError, "下载器不存在"))
		return
	}
	go func() {
		for index, URL := range URLs {
			if err := downloader.CommitDownloadTask(d, downloader.Task{
				URL:  URL,
				Type: constants.DownloaderType(req.Type),
				Path: req.DownloadPath,
			}); err != nil {
				logger.Logger.Errorf("任务%d发送失败:%v", index, err)
			}
			logger.Logger.Infof("任务%d发送成功", index)
			time.Sleep(time.Second * 1)
		}
	}()
	c.JSON(http.StatusOK, resp.Success(&dto.AddDownloadTaskResponse{}))
}