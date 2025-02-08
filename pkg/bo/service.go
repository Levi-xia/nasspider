package bo

type GetTaskListRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type GetTaskListResponse struct {
	List []bo.TvTask `json:"list"`
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
	ID common.ID `json:"id"`
}

type EditTaskRequest struct {
	ID           common.ID `json:"id"`
	URL          string    `json:"url"`
	Name         string    `json:"name"`
	TotalEp      int       `json:"total_ep"`
	CurrentEp    int       `json:"current_ep"`
	Status       int       `json:"status"`
	Provider     string    `json:"provider"`
	Downloader   string    `json:"downloader"`
	DownloadPath string    `json:"download_path"`
	Type         string    `json:"type"`
}

type EditTaskResponse struct {
	ID common.ID `json:"id"`
}

type UpdateCurrentEpRequest struct {
	ID        common.ID `json:"id"`
	CurrentEp int       `json:"current_ep"`
}

type UpdateCurrentEpResponse struct {
	ID common.ID `json:"id"`
}

type UpdateStatusRequest struct {
	ID     common.ID `json:"id"`
	Status int       `json:"status"`
}

type UpdateStatusResponse struct {
	ID common.ID `json:"id"`
}