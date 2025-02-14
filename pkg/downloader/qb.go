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

func NewQBittorrentDownloader() *QBittorrentDownloader {
	return &QBittorrentDownloader{}
}

func (q *QBittorrentDownloader) SendTask(task Task) error {
	if err := q.getToken(); err != nil {
		return err
	}
	return nil
}

func (q *QBittorrentDownloader) doTask(URL string, path string) error {
	payload := map[string]interface{}{
		"urls":        []string{URL},
		"savepath":    path,
		"cookie":      q.tokenStr,
		"category":    "movies",
		"root_folder": true,
		"upLimit":     100,
		"ratioLimit":  0.1,
	}
	host := config.GetConf(config.Conf.Downloader.QBittorrent.Host, constants.ENV_QB_HOST)
	port := config.GetConf(config.Conf.Downloader.QBittorrent.Port, constants.ENV_QB_PORT)
	resp, err := utils.HttpDo(
		fmt.Sprintf("%s:%d/api/v2/torrents/add", host, port),
		http.MethodPost,
		payload,
		utils.WithHeaders(map[string]string{
			"Content-Type": "multipart/form-data",
		}),
	)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}

func (q *QBittorrentDownloader) getToken() error {
	if q.tokenStr != "" && q.tokenTime > time.Now().Unix() {
		return nil
	}
	host := config.GetConf(config.Conf.Downloader.QBittorrent.Host, constants.ENV_QB_HOST)
	port := config.GetConf(config.Conf.Downloader.QBittorrent.Port, constants.ENV_QB_PORT)
	username := config.GetConf(config.Conf.Downloader.QBittorrent.Username, constants.ENV_QB_USERNAME)
	password := config.GetConf(config.Conf.Downloader.QBittorrent.Password, constants.ENV_QB_PASSWORD)
	resp, err := utils.HttpDo(
		fmt.Sprintf("%s:%d/api/v2/auth/login", host, port),
		http.MethodPost,
		map[string]interface{}{
			"username": username,
			"password": password,
		},
	)
	if err != nil {
		return err
	}

	fmt.Println(string(resp))

	return nil
}
