package solution

import (
	"github.com/kyokomi/emoji"
	"testing"
)

func TestGetMessage(t *testing.T) {

	got := GetMessage()
	want := emoji.Sprint("Hello :world_map:!")

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	} else {
		t.Logf("got %q", got)
	}
}
