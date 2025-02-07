package provider

import (
	"nasspider/pkg/constants"
	"nasspider/utils"
	"net/http"
)

type Provider interface {
	ParseURLs(URL string, CurrentEp int) ([]string, int, error)
}

var ProviderMap = map[constants.ProviderName]Provider{
	constants.DownloaderDoMP4: &DoMP4Provider{},
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