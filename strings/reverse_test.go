package strings

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReverse(t *testing.T) {
	want := "あいうえお"
	got := Reverse("おえういあ")

	if !cmp.Equal(want, got) {
		t.Fatalf("want: %#v, got: %#v", want, got)
	}
}
