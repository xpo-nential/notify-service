package notifyservice

type INotifyService interface {
	SendMessage() error
}
