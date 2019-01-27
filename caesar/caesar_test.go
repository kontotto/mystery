package caesar

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNormalize(t *testing.T) {
	t.Run("should upper hiragana and katakana", func(t *testing.T) {
		want := "あおつやよアオツヤヨ"
		got := normalize("ぁぉっゃょァォッャョ")

		if !cmp.Equal(want, got) {
			t.Fatalf("want: %#v, got: %#v", want, got)
		}
	})

	t.Run("should not change", func(t *testing.T) {
		want := "azAZあんアンがぼぱぽガボパポ"
		got := normalize(want)

		if !cmp.Equal(want, got) {
			t.Fatalf("want: %#v, got: %#v", want, got)
		}
	})
}

func TestValidate(t *testing.T) {
	t.Run("should validate literal", func(t *testing.T) {
		err := validate("azAZあんアンがぼぱぽガボパポぁぉっゃょ")

		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("should not validate literal", func(t *testing.T) {
		err := validate("漢字")

		if err == nil {
			t.Fail()
		}
	})
}
