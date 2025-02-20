package notification

import (
	"nasspider/pkg/constants"
)

type Notifier interface {
	Notify(title string, subTitle string, body string, options map[string]interface{}) error
}

var NotifierMap = map[constants.NotifierName]Notifier{
	constants.NotifierIosBark: &NotifierIosBark{},
}

func SendNotification(n Notifier, title string, subTitle string, body string, options map[string]interface{}) error {
	return n.Notify(title, subTitle, body, options)
}