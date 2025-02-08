package bo

type GetTaskListRequest struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	StatusList []int `json:"status_list"`
}

type GetTaskListResponse struct {
	List  []TVTask `json:"list"`
	Total int64    `json:"total"`
}

type CountTaskListRequest struct {
	StatusList []int `json:"status_list"`
}

type CountTaskListResponse struct {
	Count int64 `json:"count"`
}

type GetTaskRequest struct {
	ID int `json:"id"`
}

type GetTaskResponse struct {
	TVTask TVTask `json:"tv_task"`
}

type AddTaskRequest struct {
	URL          string `json:"url"`
	Name         string `json:"name"`
	TotalEp      int    `json:"total_ep"`
	CurrentEp    int    `json:"current_ep"`
	Status       int    `json:"status"`
	Provider     string `json:"provider"`
	Downloader   string `json:"downloader"`
	DownloadPath string `json:"download_path"`
	Type         string `json:"type"`
}

type AddTaskResponse struct {
	ID int `json:"id"`
}

type EditTaskRequest struct {
	ID           int    `json:"id"`
	URL          string `json:"url"`
	Name         string `json:"name"`
	TotalEp      int    `json:"total_ep"`
	CurrentEp    int    `json:"current_ep"`
	Status       int    `json:"status"`
	Provider     string `json:"provider"`
	Downloader   string `json:"downloader"`
	DownloadPath string `json:"download_path"`
	Type         string `json:"type"`
}

type EditTaskResponse struct {
	ID int `json:"id"`
}

type UpdateCurrentEpRequest struct {
	ID        int `json:"id"`
	CurrentEp int `json:"current_ep"`
}

type UpdateCurrentEpResponse struct {
	ID int `json:"id"`
}

type UpdateStatusRequest struct {
	ID     int `json:"id"`
	Status int `json:"status"`
}

type UpdateStatusResponse struct {
	ID int `json:"id"`
}
