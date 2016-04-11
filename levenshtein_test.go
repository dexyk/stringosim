package stringosim

import (
    "testing"
)

var testSimilarityOptions1 SimilarityOptions = SimilarityOptions{
    InsertCost:     2,
    DeleteCost:     3,
    SubstituteCost: 5,
}
var testSimilarityOptions2 SimilarityOptions = SimilarityOptions{
    InsertCost:     5,
    DeleteCost:     2,
    SubstituteCost: 3,
}
var testSimilarityOptions3 SimilarityOptions = SimilarityOptions{
    InsertCost:     3,
    DeleteCost:     5,
    SubstituteCost: 2,
}
var testSimilarityOptions4 SimilarityOptions = SimilarityOptions{
    InsertCost:      DefaultSimilarityOptions.InsertCost,
    DeleteCost:      DefaultSimilarityOptions.DeleteCost,
    SubstituteCost:  DefaultSimilarityOptions.SubstituteCost,
    CaseInsensitive: true,
}

type LevenshteinTest struct {
    src string
    trg string
    dis int
    opt SimilarityOptions
}

var levenshteinTests = []LevenshteinTest{
    {"", "", 0, DefaultSimilarityOptions},
    {"", "", 0, testSimilarityOptions1},
    {"", "", 0, testSimilarityOptions2},
    {"", "", 0, testSimilarityOptions3},
    {"x", "", 1, DefaultSimilarityOptions},
    {"x", "", 3, testSimilarityOptions1},
    {"x", "", 2, testSimilarityOptions2},
    {"x", "", 5, testSimilarityOptions3},
    {"x", "x", 0, DefaultSimilarityOptions},
    {"x", "x", 0, testSimilarityOptions1},
    {"x", "x", 0, testSimilarityOptions2},
    {"x", "x", 0, testSimilarityOptions3},
    {"xx", "xy", 1, DefaultSimilarityOptions},
    {"xx", "xy", 5, testSimilarityOptions1},
    {"xx", "xy", 3, testSimilarityOptions2},
    {"xx", "xy", 2, testSimilarityOptions3},
    {"xxx", "xyx", 1, DefaultSimilarityOptions},
    {"xxx", "xyx", 5, testSimilarityOptions1},
    {"xxx", "xyx", 3, testSimilarityOptions2},
    {"xxx", "xyx", 2, testSimilarityOptions3},
    {"xxyy", "xyz", 2, DefaultSimilarityOptions},
    {"xxyy", "xyz", 8, testSimilarityOptions1},
    {"xxyy", "xyz", 5, testSimilarityOptions2},
    {"xxyy", "xyz", 7, testSimilarityOptions3},
    {"xxyyzz", "xxxzzz", 2, DefaultSimilarityOptions},
    {"xxyyzz", "xxxzzz", 10, testSimilarityOptions1},
    {"xxyyzz", "xxxzzz", 6, testSimilarityOptions2},
    {"xxyyzz", "xxxzzz", 4, testSimilarityOptions3},
    {"asdlkajsdlkasdkj", "fkdsjlkdf", 11, DefaultSimilarityOptions},
    {"STRING", "string", 0, testSimilarityOptions4},
    {"STRING", "Astring", 1, testSimilarityOptions4},
}

func TestLevenshtein(t *testing.T) {
    for _, test := range levenshteinTests {
        dis := Levenshtein([]rune(test.src), []rune(test.trg), test.opt)

        if dis != test.dis {
            t.Log("Levenshtein distance between", test.src, "and", test.trg, "is", dis, "but should be", test.dis)
            t.Fail()
        }
    }
}
