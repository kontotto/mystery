package caesar

import (
	"testing"
	"unicode/utf8"

	"github.com/google/go-cmp/cmp"
)

func TestClassify(t *testing.T) {
	t.Run("should classify for alphabet", func(t *testing.T) {
		want := AlphabetType
		got := classify('a')

		if !cmp.Equal(want, got) {
			t.Fatalf("want: %#v, got: %#v", want, got)
		}
	})

	t.Run("should not for kanji", func(t *testing.T) {
		want := Undefined
		got := classify('心')

		if !cmp.Equal(want, got) {
			t.Fatalf("want: %#v, got: %#v", want, got)
		}
	})
}

func TestCaesarOne(t *testing.T) {
	t.Run("should 1 forwarder", func(t *testing.T) {
		want := 'a'
		got, err := caesarOne(Alphabet.Chars, 'z', 1)

		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(want, got) {
			t.Fatalf("want: %#v, got: %#v", want, got)
		}
	})

	t.Run("should 1 backwarder", func(t *testing.T) {
		want := 'z'
		got, err := caesarOne(Alphabet.Chars, 'a', -1)

		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(want, got) {
			t.Fatalf("want: %#v, got: %#v", want, got)
		}
	})

	t.Run("should fail", func(t *testing.T) {
		want := utf8.RuneError
		got, err := caesarOne(Alphabet.Chars, 'A', -1)

		if err == nil {
			t.Fatalf("want: %#v, got: %#v", want, got)
		}
	})
}

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
