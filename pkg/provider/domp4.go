package provider

import (
	"nasspider/config"
	"nasspider/pkg/logger"
	"nasspider/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
)

type DoMP4Provider struct{}

var doMp4Config = config.Conf.Provider.DoMP4

// ParseURLs 解析xpath获取URLs,当前集
func (d DoMP4Provider) ParseURLs(URL string, CurrentEp int) ([]string, int, error) {
	html, err := getHtml(URL)
	if err != nil {
		return nil, 0, err
	}
	root, _ := htmlquery.Parse(strings.NewReader(html))
	numsPath := htmlquery.FindOne(root, doMp4Config.Xpath)
	numsStr := htmlquery.InnerText(numsPath)
	numsInt, _ := strconv.Atoi(numsStr)
	logger.Logger.Infof("共：%d条记录\n", numsInt)

	return []string{}, 0, nil

	// 集数与URLs各种一致性判断 。。。
	// if currentEp > tvTask.TotalEp {
	// 	return fmt.Errorf("异常:解析集数=%d, 总集数=%d", currentEp, tvTask.TotalEp)
	// }

}

func getHtml(url string) (string, error) {
	resp, err := utils.HttpDo(
		url,
		string(http.MethodGet),
		nil,
		utils.WithHeaders(map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3776.0 Safari/537.36",
		}),
	)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}
