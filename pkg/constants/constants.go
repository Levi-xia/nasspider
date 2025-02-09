package constants

type DownloaderType string
type TaskStatus int
type DownloaderName string
type ProviderName string

var (
	Torrent DownloaderType = "torrent"
	Magnet  DownloaderType = "magnet"
)

var (
	Waiting TaskStatus = 0
	Doing   TaskStatus = 1
	Finish  TaskStatus = 2
	Abort   TaskStatus = 3
	Error   TaskStatus = 4

	TaskStatusMap = map[TaskStatus]string{
		Waiting: "等待中",
		Doing:   "追更中",
		Finish:  "已完成",
		Abort:   "已中止",
		Error:   "已出错",
	}
)

var (
	DownloaderThunder DownloaderName = "thunder"
	DownloaderDoMP4   ProviderName   = "domp4"
)
