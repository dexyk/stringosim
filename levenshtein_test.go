package stringosim

import (
    "testing"
)

var testLevenshteinOptions1 LevenshteinSimilarityOptions = LevenshteinSimilarityOptions{
    InsertCost:     2,
    DeleteCost:     3,
    SubstituteCost: 5,
}
var testLevenshteinOptions2 LevenshteinSimilarityOptions = LevenshteinSimilarityOptions{
    InsertCost:     5,
    DeleteCost:     2,
    SubstituteCost: 3,
}
var testLevenshteinOptions3 LevenshteinSimilarityOptions = LevenshteinSimilarityOptions{
    InsertCost:     3,
    DeleteCost:     5,
    SubstituteCost: 2,
}
var testLevenshteinOptions4 LevenshteinSimilarityOptions = LevenshteinSimilarityOptions{
    InsertCost:      DefaultLevenshteinSimilarityOptions.InsertCost,
    DeleteCost:      DefaultLevenshteinSimilarityOptions.DeleteCost,
    SubstituteCost:  DefaultLevenshteinSimilarityOptions.SubstituteCost,
    CaseInsensitive: true,
}

type LevenshteinTest struct {
    src string
    trg string
    dis int
    opt LevenshteinSimilarityOptions
}

var levenshteinTests = []LevenshteinTest{
    {"", "", 0, DefaultLevenshteinSimilarityOptions},
    {"", "", 0, testLevenshteinOptions1},
    {"", "", 0, testLevenshteinOptions2},
    {"", "", 0, testLevenshteinOptions3},
    {"x", "", 1, DefaultLevenshteinSimilarityOptions},
    {"x", "", 3, testLevenshteinOptions1},
    {"x", "", 2, testLevenshteinOptions2},
    {"x", "", 5, testLevenshteinOptions3},
    {"x", "x", 0, DefaultLevenshteinSimilarityOptions},
    {"x", "x", 0, testLevenshteinOptions1},
    {"x", "x", 0, testLevenshteinOptions2},
    {"x", "x", 0, testLevenshteinOptions3},
    {"xx", "xy", 1, DefaultLevenshteinSimilarityOptions},
    {"xx", "xy", 5, testLevenshteinOptions1},
    {"xx", "xy", 3, testLevenshteinOptions2},
    {"xx", "xy", 2, testLevenshteinOptions3},
    {"xxx", "xyx", 1, DefaultLevenshteinSimilarityOptions},
    {"xxx", "xyx", 5, testLevenshteinOptions1},
    {"xxx", "xyx", 3, testLevenshteinOptions2},
    {"xxx", "xyx", 2, testLevenshteinOptions3},
    {"xxyy", "xyz", 2, DefaultLevenshteinSimilarityOptions},
    {"xxyy", "xyz", 8, testLevenshteinOptions1},
    {"xxyy", "xyz", 5, testLevenshteinOptions2},
    {"xxyy", "xyz", 7, testLevenshteinOptions3},
    {"xxyyzz", "xxxzzz", 2, DefaultLevenshteinSimilarityOptions},
    {"xxyyzz", "xxxzzz", 10, testLevenshteinOptions1},
    {"xxyyzz", "xxxzzz", 6, testLevenshteinOptions2},
    {"xxyyzz", "xxxzzz", 4, testLevenshteinOptions3},
    {"asdlkajsdlkasdkj", "fkdsjlkdf", 11, DefaultLevenshteinSimilarityOptions},
    {"STRING", "string", 0, testLevenshteinOptions4},
    {"STRING", "Astring", 1, testLevenshteinOptions4},
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
