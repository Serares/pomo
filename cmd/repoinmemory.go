//go:build inmemory
// +build inmemory

package cmd

import (
	"github.com/Serares/pomo/pomodoro"
	"github.com/Serares/pomo/pomodoro/repository"
)

func getRepo() (pomodoro.Repository, error) {
	return repository.NewInMemoryRepo(), nil
}
