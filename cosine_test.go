package stringosim

import (
    "testing"
)

var testCosineOptions1 CosineSimilarityOptions = CosineSimilarityOptions{
    CaseInsensitive: true,
    NGramLength:     3,
}

type CosineTest struct {
    src string
    trg string
    dis float64
    opt CosineSimilarityOptions
}

var CosineTests = []CosineTest{
    {"", "", 0.0, DefaultCosineSimilarityOptions},
    {"xxxyyy", "xxxyyy", 0.0, DefaultCosineSimilarityOptions},
    {"xxxyyy", "yyyxxx", 0.1111111111111, DefaultCosineSimilarityOptions},
    {"xxyzxyyzzy", "xyyxzyzxyzyx", 0.2364582844290667, DefaultCosineSimilarityOptions},
    {"xxyzxyyzzy", "XYYXZYZXYZYX", 0.5527864045000421, testCosineOptions1},
    {"xxyyzz", "xxxzzz", 0.40371520600005606, DefaultCosineSimilarityOptions},
    {"asdlkajsdlkasdkj", "fkdsjlkdf", 0.8825559560970593, DefaultCosineSimilarityOptions},
    {"STRING", "sting", 1.0, DefaultCosineSimilarityOptions},
    {"STRING", "sting", 0.7113248654051871, testCosineOptions1},
    {"comparing the string similarity", "this is compared using Cosine similarity", 0.3519132720916943, DefaultCosineSimilarityOptions},
    {"comparing the string similarity", "this is compared using Cosine similarity", 0.5291024827582208, testCosineOptions1},
}

func TestCosine(t *testing.T) {
    for _, test := range CosineTests {
        dis := Cosine([]rune(test.src), []rune(test.trg), test.opt)
        if !EqualFloat64(dis, test.dis) {
            t.Log("Cosine distance between", test.src, "and", test.trg, "is", dis, "but should be", test.dis)
            t.Fail()
        }
    }
}
