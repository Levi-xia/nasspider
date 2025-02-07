package provider

import "nasspider/pkg/constants"

type Provider interface {
	ParseURLs(URL string, CurrentEp int) ([]string, int, error)
}

var ProviderMap = map[constants.ProviderName]Provider{
	constants.DownloaderDoMP4: &DoMP4Provider{},
}
