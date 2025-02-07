package downloader

import "log"

type DownloaderType string

var (
	Torrent DownloaderType = "torrent"
	Magnet  DownloaderType = "magnet"
)

type Task struct {
	URL  string
	Path string
	Type DownloaderType
}

type Downloader interface {
	SendTask(task Task) error
}

func Do(f Downloader, task Task) error {
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
		}
	}()
	if err = f.SendTask(task); err != nil {
		return err
	}
	return nil
}
