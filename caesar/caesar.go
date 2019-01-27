package caesar

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Charset struct {
	Chars string
	Type  CharType
}

var Number = Charset{
	"0123456789",
	NumberType,
}

var Alphabet = Charset{
	"abcdefghijklmnopqrstuvwxyz",
	AlphabetType,
}

var EmAlphabet = Charset{
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	EmAlphabetType,
}

var Hiragana = Charset{
	"あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをん",
	HiraganaType,
}

var HiraganaDakuten = Charset{
	"がぎぐげござじずぜぞだぢづでどばびぶべぼ",
	HiraganaDakutenType,
}

var HiraganaHandakuten = Charset{
	"ぱぴぷぺぽ",
	HiraganaHandakutenType,
}

var Katakana = Charset{
	"アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲン",
	KatakanaType,
}

var KatakanaDakuten = Charset{
	"ガギグゲゴザジズゼゾダヂヅデドバビブベボ",
	KatakanaDakutenType,
}

var KatakanaHandakuten = Charset{
	"パピプペポ",
	KatakanaHandakutenType,
}

var Undefined = Charset{
	"",
	UndefinedType,
}

type CharType int

const (
	UndefinedType CharType = iota
	NumberType
	AlphabetType
	EmAlphabetType
	HiraganaType
	HiraganaDakutenType
	HiraganaHandakutenType
	KatakanaType
	KatakanaDakutenType
	KatakanaHandakutenType
)

const Komoji = "ぁぃぅぇぉっゃゅょァィゥェォッャュョ"

func Caesar(s string, offset int) (string, error) {
	normal := normalize(s)
	result := make([]rune, utf8.RuneCountInString(normal))
	var err error

	i := 0
	for _, r := range normal {
		charset := classify(r)
		result[i], err = caesarOne(charset.Chars, r, offset)
		if err != nil {
			return "", err
		}
		i++
	}

	return string(result), nil
}

func caesarOne(s string, substr rune, offset int) (rune, error) {
	index := runeIndex(s, substr)

	if index == -1 {
		return utf8.RuneError, fmt.Errorf("%c is not found in %v", substr, s)
	}
	length := utf8.RuneCountInString(s)
	index = (index + offset) % length
	index = (index + length) % length
	return []rune(s)[index], nil
}

func classify(r rune) Charset {
	charsets := []Charset{
		Alphabet, Number, EmAlphabet,
		Hiragana, HiraganaDakuten, HiraganaHandakuten,
		Katakana, KatakanaDakuten, KatakanaHandakuten,
	}
	for _, c := range charsets {
		if strings.ContainsRune(c.Chars, r) {
			return c
		}
	}
	return Undefined
}

func normalize(src string) string {
	normal := make([]rune, utf8.RuneCountInString(src))

	i := 0
	for _, r := range src {
		if strings.ContainsRune(Komoji, r) {
			// https://unicodemap.org/range/62/Hiragana/
			// https://unicodemap.org/range/63/Katakana/
			normal[i] = []rune(src)[i] + 1
		} else {
			normal[i] = []rune(src)[i]
		}
		i++
	}
	return string(normal)
}

func validate(str string) error {
	all := strings.Join(
		[]string{
			Alphabet.Chars, Number.Chars, EmAlphabet.Chars,
			Hiragana.Chars, HiraganaDakuten.Chars, HiraganaHandakuten.Chars,
			Katakana.Chars, KatakanaDakuten.Chars, KatakanaHandakuten.Chars,
		},
		"",
	)
	tmp := normalize(str)
	for len(tmp) > 0 {
		r, size := utf8.DecodeRuneInString(tmp)
		if strings.Contains(all, tmp[:size]) {
			tmp = tmp[size:]
			continue
		}
		return fmt.Errorf("not supported literal: %c", r)
	}
	return nil
}

func runeIndex(s string, substr rune) int {
	runes := []rune(s)
	for i, r := range runes {
		if r == substr {
			return i
		}
	}
	return -1
}
