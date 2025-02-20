package provider

import (
	"nasspider/pkg/constants"
)

var ProviderMap = map[constants.ProviderName]Provider{
	constants.ProviderDoMP4: &DoMP4Provider{},
}

var SearchProviderMap = map[constants.ProviderName]Provider{
	constants.ProviderDoMP4: &DoMP4Provider{},
}

type Provider interface {
	ParseURLs(URL string, currentEp int) ([]string, int, error)
	Search(content string) ([]SearchSet, error)
}

type SearchResult struct {
	Provider   string      `json:"provider"`
	SearchSets []SearchSet `json:"search_sets"`
	Cost       int64       `json:"cost"`
}

type SearchSet struct {
	URL       string   `json:"url"`
	Title     string   `json:"title"`
	UpdatedEp int      `json:"current_ep"`
	TotalEp   int      `json:"total_ep"`
	Resources []string `json:"resources"`
	CanChased bool     `json:"can_chased"` // 是否可以追更
}
