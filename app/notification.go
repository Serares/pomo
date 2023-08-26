//go:build !containers && !disable_notificaitons
// +build !containers,!disable_notificaitons

package app

import "github.com/Serares/notify"

func send_notification(msg string) {
	n := notify.New("Pomodoro", msg, notify.SeverityNormal)

	n.Send()
}
