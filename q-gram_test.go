package stringosim

import (
	"testing"
)

var testQGramOptions1 QGramSimilarityOptions = QGramSimilarityOptions{
	CaseInsensitive: true,
	NGramSizes:      []int{3},
}

type QGramTest struct {
	src string
	trg string
	dis int
	opt QGramSimilarityOptions
}

var qgramTests = []QGramTest{
	{"", "", 0, DefaultQGramSimilarityOptions},
	{"xxxyyy", "xxxyyy", 0, DefaultQGramSimilarityOptions},
	{"xxxyyy", "yyyxxx", 2, DefaultQGramSimilarityOptions},
	{"xxyzxyyzzy", "xyyxzyzxyzyx", 6, DefaultQGramSimilarityOptions},
	{"xxyzxyyzzy", "XYYXZYZXYZYX", 10, testQGramOptions1},
	{"xxyyzz", "xxxzzz", 6, DefaultQGramSimilarityOptions},
	{"asdlkajsdlkasdkj", "fkdsjlkdf", 21, DefaultQGramSimilarityOptions},
	{"STRING", "sting", 9, DefaultQGramSimilarityOptions},
	{"STRING", "sting", 5, testQGramOptions1},
	{"comparing the string similarity", "this is compared using qgram similarity", 30, DefaultQGramSimilarityOptions},
	{"comparing the string similarity", "this is compared using qgram similarity", 36, testQGramOptions1},
}

func TestQGram(t *testing.T) {
	for _, test := range qgramTests {
		dis := QGram([]rune(test.src), []rune(test.trg), test.opt)
		if dis != test.dis {
			t.Log("Hamming distance between", test.src, "and", test.trg, "is", dis, "but should be", test.dis)
			t.Fail()
		}
	}
}
