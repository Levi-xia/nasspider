package provider

type DoMP4Provider struct{}

// ParseURLs 解析xpath获取URLs,当前集
func (d DoMP4Provider) ParseURLs(URL string, CurrentEp int) ([]string, int, error) {
	return []string{}, 0, nil


	// 集数与URLs各种一致性判断 。。。
	// if currentEp > tvTask.TotalEp {
	// 	return fmt.Errorf("异常:解析集数=%d, 总集数=%d", currentEp, tvTask.TotalEp)
	// }

}
