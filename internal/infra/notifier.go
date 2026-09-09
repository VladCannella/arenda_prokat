package infra

import (
	"fmt"
	"rental/internal/app"
)

type ConsoleNotifier struct{}

var _ app.Notifier = &ConsoleNotifier{}

func (c *ConsoleNotifier) Notify(message string) error {
	_, err := fmt.Println(message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

type NoopNotifier struct{}

var _ app.Notifier = &NoopNotifier{}

func (n *NoopNotifier) Notify(message string) error {
	return nil
}
