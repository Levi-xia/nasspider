package main

import (
	"nasspider/config"
	"nasspider/pkg/downloader"
)

func main() {

	// 初始化配置
	config.InitConfig()

	thunder := downloader.NewThunderDownloader()
	downloader.Do(thunder, downloader.Task{
		URL: "magnet:?xt=urn:btih:8fcaa1ea861922f6952fcd7b41f1a3296cc9ee50&dn=[www.domp4.cc]冬至.EP36.HD1080p.mp4&tr=https://sparkle.ghostchu-services.top:443/announce&tr=https://1337.abcvg.info:443/announce&tr=https://p2p.azu.red:443/announce&tr=https://pybittrack.retiolus.net:443/announce&tr=https://torrent.tracker.durukanbal.com:443/announce&tr=https://tracker-zhuqiy.dgj055.icu:443/announce&tr=https://tracker.cloudit.top:443/announce&tr=https://tracker.gcrenwp.top:443/announce&tr=https://tracker.lilithraws.org:443/announce&tr=https://tracker.pmman.tech:443/announce&tr=https://tracker.tamersunion.org:443/announce&tr=https://tracker1.520.jp:443/announce&tr=https://trackers.mlsub.net:443/announce&tr=https://www.peckservers.com:9443/announce&tr=https://btn-prod.ghostchu-services.top/tracker/announce&tr=http://1337.abcvg.info:80/announce&tr=http://ipv6.rer.lol:6969/announce&tr=http://public.tracker.vraphim.com:6969/announce&tr=http://taciturn-shadow.spb.ru:6969/announce&tr=http://tracker.bt4g.com:2095/announce",
		Path: "/downloads/",
		Type: downloader.Magnet,
	})
}