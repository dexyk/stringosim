package stringosim

import (
    "testing"
)

var testHammingOptions1 HammingSimilarityOptions = HammingSimilarityOptions{
    CaseInsensitive: true,
}

type HammingTest struct {
    src string
    trg string
    dis int
    opt HammingSimilarityOptions
    err error
}

var hammingTests = []HammingTest{
    {"", "", 0, DefaultHammingSimilarityOptions, nil},
    {"", "", 0, testHammingOptions1, nil},
    {"x", "", -1, DefaultHammingSimilarityOptions, HAMMING_ERROR_DIFFERENT_LENGTH},
    {"x", "", -1, testHammingOptions1, HAMMING_ERROR_DIFFERENT_LENGTH},
    {"x", "x", 0, DefaultHammingSimilarityOptions, nil},
    {"x", "x", 0, testHammingOptions1, nil},
    {"xx", "xy", 1, DefaultHammingSimilarityOptions, nil},
    {"xx", "xy", 1, testHammingOptions1, nil},
    {"xxx", "XYX", 3, DefaultHammingSimilarityOptions, nil},
    {"xxx", "xyx", 1, testHammingOptions1, nil},
    {"xxyy", "xyz", -1, DefaultHammingSimilarityOptions, HAMMING_ERROR_DIFFERENT_LENGTH},
    {"xxyy", "xyz", -1, testHammingOptions1, HAMMING_ERROR_DIFFERENT_LENGTH},
    {"xxyyzz", "xxxzzz", 2, DefaultHammingSimilarityOptions, nil},
    {"xxyyzz", "xXxzZz", 2, testHammingOptions1, nil},
    {"asdlkajsdlkasdkj", "fkdsjlkdf", -1, DefaultHammingSimilarityOptions, HAMMING_ERROR_DIFFERENT_LENGTH},
    {"STRING", "string", 6, DefaultHammingSimilarityOptions, nil},
    {"STRING", "string", 0, testHammingOptions1, nil},
}

func TestHamming(t *testing.T) {
    for _, test := range hammingTests {
        dis, err := Hamming([]rune(test.src), []rune(test.trg), test.opt)
        if err != nil || test.err != nil {
            if !CompareErrors(err, test.err) {
                t.Log("Error received", err, "instead of", test.err)
                t.Fail()
            }
        } else {
            if dis != test.dis {
                t.Log("Hamming distance between", test.src, "and", test.trg, "is", dis, "but should be", test.dis)
                t.Fail()
            }
        }
    }
}
