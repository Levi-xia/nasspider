package provider

type Provider interface {
	ParseURLs(webUrl string) ([]string, error)
}

func DoTask(p Provider) error {
	// 获取链接
	// 获取当前已更新剧集
	// 获取总集数
	// 获取存储位置
	// 获取xpath impl
	// 获取下载链接列表
	// 获取downloader
	// 触发下载
	// 通知
	// 发起异步检查，下载完成后移动到存储位置
	return nil
}