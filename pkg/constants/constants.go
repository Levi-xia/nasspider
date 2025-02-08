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
	Doing  TaskStatus = 0
	Finish TaskStatus = 1
	Abort  TaskStatus = 2
	Error  TaskStatus = 3

	TaskStatusMap = map[TaskStatus]string{
		Doing:  "进行中",
		Finish: "已完成",
		Abort:  "已中止",
		Error:  "已出错",
	}
)

var (
	DownloaderThunder DownloaderName = "thunder"
	DownloaderDoMP4   ProviderName   = "domp4"
)
