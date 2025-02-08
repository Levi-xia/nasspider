package app

import "github.com/gin-gonic/gin"

// Index 首页列表
func Index(c *gin.Context) {
}


// EditTaskRequest 发起修改任务
func EditTask(c *gin.Context) {
	
}

// AddTaskRequest 发起添加任务
func AddTask(c *gin.Context) {
	resp := &common.Result{}
	req := &dto.AddTaskRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		c.JSON(http.StatusOK, resp.Error(common.ParamError, bootstrap.GetErrorMsg(req, err)))
	}
	resp, err := service.AddTask(&bo.AddTaskRequest{
		Name:         req.Name,
		URL:          req.URL,
	})

	c.JSON(http.StatusOK, resp.Success(&dto.AddTaskResponse{
		ID: resp.ID.ID,
	}))
}

// GetTaskRecords 获取任务记录列表
func GetTaskRecords(c *gin.Context) {

}

// 触发任务
func TriggerTask(c *gin.Context) {
	
}