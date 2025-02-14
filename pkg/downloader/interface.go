package downloader

import (
	"nasspider/pkg/constants"
)

type Task struct {
	URL  string
	Path string
	Type constants.DownloaderType
}

type Downloader interface {
	SendTask(task Task) error
}

var DownloaderMap = map[constants.DownloaderName]Downloader{
	constants.DownloaderThunder: NewThunderDownloader(),
	constants.DownloaderQBittorrent: NewQBittorrentDownloader(),
}

func CommitDownloadTask(d Downloader, task Task) error {
	if err := d.SendTask(task); err != nil {
		return err
	}
	return nil
}
