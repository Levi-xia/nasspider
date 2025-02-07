package provider

type DoMP4Provider struct{}

// ParseURLs 解析xpath获取URLs
func (d DoMP4Provider) ParseURLs(URL string, CurrentEp int, xPath string) ([]string, error) {
	return []string{}, nil
}
