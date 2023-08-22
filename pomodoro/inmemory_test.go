//go:build inmemory
// +build inmemory

package pomodoro_test

import (
	"testing"

	"github.com/Serares/pomo/pomodoro"
	"github.com/Serares/pomo/pomodoro/repository"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()
	// return an empty cleanup function, because in-memory doesn't requre a cleanup functio
	return repository.NewInMemoryRepo(), func() {}
}
