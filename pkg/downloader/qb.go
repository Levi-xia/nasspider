package downloader

import (
	"fmt"
	"nasspider/config"
	"nasspider/pkg/constants"
	"nasspider/utils"
	"net/http"
	"time"
)

type QBittorrentDownloader struct {
	tokenStr  string
	tokenTime int64
}

func (q *QBittorrentDownloader) SendTask(task Task) error {
	return nil
}

func (q *QBittorrentDownloader) login() error {
	if q.tokenStr != "" && q.tokenTime > time.Now().Unix() {
		return nil
	}
	host := config.GetConf(config.Conf.Downloader.QBittorrent.Host, constants.ENV_QB_HOST)
	port := config.GetConf(config.Conf.Downloader.QBittorrent.Port, constants.ENV_QB_PORT)

	utils.HttpDo(
		fmt.Sprintf("%s:%d/api/v2/auth/login", host, port),
		http.MethodPost,
		map[string]interface{}{
			"username": "",
			"password": "",
		},
	)
	return nil
}
