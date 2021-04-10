package guess

import (
	"testing"
)

func TestFetchGameInfo(t *testing.T) {
	instance := New("74089117-d46c-49cc-b6dc-1336004d27aa")
	t.Error(instance.QueryCorrectGuesses())
}
