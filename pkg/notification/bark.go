package notification

import (
	"nasspider/config"
	"nasspider/pkg/constants"
)

type NotifierIosBark struct {}


func(n *NotifierIosBark) Notify(title string, subTitle string, body string, options map[string]interface{}) error {

	host := config.GetConf(config.Conf.Notification.Bark.Host, constants.ENV_NOTIFY_BARK_HOST)
	key := config.GetConf(config.Conf.Notification.Bark.Key, constants.ENV_NOTIFY_BARK_KEY)

	if host == "" || key == "" {
		return nil
	}
	// Todo
	return nil
}