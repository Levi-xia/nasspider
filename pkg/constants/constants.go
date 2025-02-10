package constants

type DownloaderType string
type TaskStatus int
type DownloaderName string
type ProviderName string
type ENVConfig string

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
	ENV_MYSQL_HOST     ENVConfig = "MYSQL_HOST"
	ENV_MYSQL_PORT     ENVConfig = "MYSQL_PORT"
	ENV_MYSQL_USER     ENVConfig = "MYSQL_USER"
	ENV_MYSQL_PASSWORD ENVConfig = "MYSQL_PASSWORD"
	ENV_SERVER_PORT    ENVConfig = "SERVER_PORT"
	ENV_THUNDER_HOST   ENVConfig = "THUNDER_HOST"
	ENV_THUNDER_PORT   ENVConfig = "THUNDER_PORT"
	ENV_ADMIN_USERNAME ENVConfig = "ADMIN_USERNAME"
	ENV_ADMIN_PASSWORD ENVConfig = "ADMIN_PASSWORD"
)

var (
	DownloaderThunder DownloaderName = "thunder"
	DownloaderDoMP4   ProviderName   = "domp4"
)
