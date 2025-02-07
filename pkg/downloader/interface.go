package downloader

import (
	"log"
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
}

func CommitDownloadTask(d Downloader, task Task) error {
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
		}
	}()
	if err = d.SendTask(task); err != nil {
		return err
	}
	return nil
}
