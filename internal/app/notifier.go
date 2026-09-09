package app

type Notifier interface {
	Notify(message string) error
}
