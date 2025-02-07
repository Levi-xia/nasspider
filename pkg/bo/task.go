package bo

type TVTask struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	URL          string `json:"url"`
	TotalEp      int    `json:"total_ep"`
	CurrentEp    int    `json:"current_ep"`
	Status       int `json:"status"`
	DownloadPath string `json:"download_path"`
	Type         string `json:"type"`
	Provider     string `json:"provider"`
	Downloader   string `json:"downloader"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type TVTaskRecord struct {
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
