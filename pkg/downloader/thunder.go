package downloader

import (
	"crypto/sha1"
	"encoding/base32"
	"encoding/json"
	"errors"
	"fmt"
	"nasspider/config"
	"nasspider/pkg/constants"
	"nasspider/pkg/logger"
	"nasspider/utils"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/zeebo/bencode"
)

type ThunderDownloader struct {
	deviceID      string
	tokenStr      string
	tokenTime     int64
}

type fileInfo struct {
	List list `json:"list"`
}

type list struct {
	Resources []resource `json:"resources"`
}

type resource struct {
	Name      string `json:"name"`
	FileSize  int    `json:"file_size"`
	FileCount int    `json:"file_count"`
	Dir       dir    `json:"dir"`
}

type dir struct {
	Resources []dirResource `json:"resources"`
}

type dirResource struct {
	FileIndex int `json:"file_index"`
}

func NewThunderDownloader() *ThunderDownloader {
	return &ThunderDownloader{}
}

// SendTask 发送任务
func (t *ThunderDownloader) SendTask(task Task) error {
	var (
		token    string
		deviceID string
		err      error
		fileInfo fileInfo
	)
	if token, err = t.getPanToken(); err != nil {
		return err
	}
	if deviceID, err = t.getDeviceID(); err != nil {
		return err
	}
	if task.Type == constants.Torrent {
		if task.URL, err = t.convertTorrentToMagnet(task.URL); err != nil {
			return err
		}
	}
	if fileInfo, err = t.ListFiles(token, task.URL); err != nil {
		return err
	}
	if len(fileInfo.List.Resources) == 0 {
		return errors.New("fileInfo is empty")
	}
	return t.doTask(token, deviceID, fileInfo, task.URL, task.Path)
}

// doTask 执行任务
func (t *ThunderDownloader) doTask(token, deviceID string, fileInfo fileInfo, url string, path string) error {
	resource := fileInfo.List.Resources[0]
	fileSize := int(resource.FileSize)
	fileCount := int(resource.FileCount)
	reqPayload := map[string]interface{}{
		"type":      "user#download-url",
		"name":      resource.Name,
		"file_name": resource.Name,
		"file_size": strconv.Itoa(fileSize),
		"space":     deviceID,
		"params": map[string]interface{}{
			"target":           deviceID,
			"url":              url,
			"total_file_count": strconv.Itoa(int(fileCount)),
			"sub_file_index":   t.getFileIndex(fileInfo),
			"file_id":          "",
			"parent_folder_id": "",
		},
	}
	var (
		resp []byte
		err  error
	)
	if resp, err = utils.HttpDo(
		fmt.Sprintf("%s:%d/webman/3rdparty/pan-xunlei-com/index.cgi/drive/v1/task?pan_auth=%s&device_space=", config.Conf.Downloader.Thunder.Host, config.Conf.Downloader.Thunder.Port, token),
		http.MethodPost,
		reqPayload,
		utils.WithHeaders(map[string]string{
			"pan-auth": token,
		}), utils.WithTimeout(time.Second*30)); err != nil {
		return err
	}
	logger.Logger.Infof("doTask resp:%v", string(resp))
	return nil
}

func (t *ThunderDownloader) ListFiles(token, url string) (fileInfo, error) {
	var (
		resp   []byte
		err    error
		result fileInfo
	)
	if resp, err = utils.HttpDo(
		fmt.Sprintf("%s:%d/webman/3rdparty/pan-xunlei-com/index.cgi/drive/v1/resource/list?pan_auth=%s&device_space=", config.Conf.Downloader.Thunder.Host, config.Conf.Downloader.Thunder.Port, token),
		http.MethodPost,
		map[string]interface{}{
			"urls": url,
		},
		utils.WithHeaders(map[string]string{
			"pan_auth": token,
		})); err != nil {
		return result, err
	}
	if err = json.Unmarshal(resp, &result); err != nil {
		return result, err
	}
	return result, nil
}

// getPanToken 获取pan token
func (t *ThunderDownloader) getPanToken() (version string, err error) {
	if version, err = t.getServerVersion(); err != nil {
		return
	}
	if !checkVersion(version, "3.21.0") {
		err = fmt.Errorf("version is not supported")
		return
	}
	if t.tokenStr != "" && t.tokenTime+600 > time.Now().Unix() {
		return t.tokenStr, nil
	}
	var resp []byte
	// 发起HTTP请求
	if resp, err = utils.HttpDo(
		fmt.Sprintf("%s:%d/webman/3rdparty/pan-xunlei-com/index.cgi/", config.Conf.Downloader.Thunder.Host, config.Conf.Downloader.Thunder.Port),
		string(http.MethodGet), nil); err != nil {
		return
	}
	re := regexp.MustCompile(`function uiauth\(value\){ return "(.*)" }`)
	if matches := re.FindStringSubmatch(string(resp)); len(matches) > 1 {
		t.tokenStr = matches[1]
		t.tokenTime = time.Now().Unix()
		return t.tokenStr, nil
	}
	return "", fmt.Errorf("get pan token failed")
}

// getDeviceID 获取设备ID
func (t *ThunderDownloader) getDeviceID() (deviceID string, err error) {
	if t.deviceID != "" {
		return t.deviceID, nil
	}
	var (
		token  string
		resp   []byte
		result map[string]interface{}
	)
	if token, err = t.getPanToken(); err != nil {
		return
	}
	if resp, err = utils.HttpDo(
		fmt.Sprintf("%s:%d/webman/3rdparty/pan-xunlei-com/index.cgi/device/info/watch", config.Conf.Downloader.Thunder.Host, config.Conf.Downloader.Thunder.Port),
		string(http.MethodPost), nil,
		utils.WithHeaders(map[string]string{
			"pan_auth": token,
		})); err != nil {
		return
	}
	if err = json.Unmarshal(resp, &result); err != nil {
		return
	}
	t.deviceID = result["target"].(string)
	return t.deviceID, nil
}

// getServerVersion 获取服务器版本
func (t *ThunderDownloader) getServerVersion() (string, error) {
	var (
		resp []byte
		err  error
	)
	if resp, err = utils.HttpDo(
		fmt.Sprintf("%s:%d/webman/3rdparty/pan-xunlei-com/index.cgi/launcher/status", config.Conf.Downloader.Thunder.Host, config.Conf.Downloader.Thunder.Port),
		string(http.MethodGet), nil); err != nil {
		return "", err
	}
	var result map[string]interface{}
	if err = json.Unmarshal(resp, &result); err != nil {
		return "", err
	}
	// 返回运行版本
	return result["running_version"].(string), nil
}

// convertTorrentToMagnet 将种子转换为磁力
func (t *ThunderDownloader) convertTorrentToMagnet(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read torrent file: %v", err)
	}
	var metadata map[string]interface{}
	if err := bencode.DecodeBytes(data, &metadata); err != nil {
		return "", fmt.Errorf("failed to decode torrent file: %v", err)
	}
	info, ok := metadata["info"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid torrent file: missing info section")
	}
	infoBytes, err := bencode.EncodeBytes(info)
	if err != nil {
		return "", fmt.Errorf("failed to encode info section: %v", err)
	}
	hash := sha1.Sum(infoBytes)
	b32hash := base32.StdEncoding.EncodeToString(hash[:])
	name, ok := info["name"].(string)
	if !ok {
		return "", fmt.Errorf("invalid torrent file: missing name in info section")
	}
	magnet := fmt.Sprintf("magnet:?xt=urn:btih:%s&dn=%s", b32hash, name)
	return magnet, nil
}

// getFileIndex 检查文件索引
func (t *ThunderDownloader) getFileIndex(fileInfo fileInfo) string {
	if fileInfo.List.Resources[0].FileCount == 1 {
		return "--1,"
	}
	maxSubFileIdx := 0
	dirResources := fileInfo.List.Resources[0].Dir.Resources
	if len(dirResources) > 0 {
		for _, resource := range dirResources {
			if resource.FileIndex > maxSubFileIdx {
				maxSubFileIdx = resource.FileIndex
			}
		}
	}
	return fmt.Sprintf("0-%d", maxSubFileIdx)
}

// CheckVersion 检查版本
func checkVersion(version, target string) bool {
	return version >= target
}
