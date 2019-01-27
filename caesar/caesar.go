package caesar

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	Number             = "0123456789"
	Alphabet           = "abcdefghijklmnopqrstuvwxyz"
	EmAlphabet         = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Hiragana           = "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをん"
	HiraganaDakuten    = "がぎぐげござじずぜぞだぢづでどばびぶべぼ"
	HiraganaHandakuten = "ぱぴぷぺぽ"
	Katakana           = "アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲン"
	KatakanaDakuten    = "ガギグゲゴザジズゼゾダヂヅデドバビブベボ"
	KatakanaHandakuten = "パピプペポ"
)

const Komoji = "ぁぃぅぇぉっゃゅょァィゥェォッャュョ"

//func caesarOne(s string, substr rune, offset int) (rune, error){
//	index := strings.IndexRune(s, substr)
//	if index == -1 {
//		return utf8.RuneError, fmt.Errorf("%c is not found in %v", substr, s)
//	}
//	index = (index + offset) % utf8.RuneCountInString(s)
//	return []rune(s)[index], nil
//}

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
			Alphabet, Number, EmAlphabet,
			Hiragana, HiraganaDakuten, HiraganaHandakuten,
			Katakana, KatakanaDakuten, KatakanaHandakuten,
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
