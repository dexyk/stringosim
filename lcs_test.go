package stringosim

import (
    "testing"
)

var testLCSOptions1 LCSSimilarityOptions = LCSSimilarityOptions{
    CaseInsensitive: true,
}

type LCSTest struct {
    src string
    trg string
    dis int
    opt LCSSimilarityOptions
}

var LCSTests = []LCSTest{
    {"", "", 0, DefaultLCSSimilarityOptions},
    {"x", "", 0, DefaultLCSSimilarityOptions},
    {"xyzxyztt", "ttxyz", 3, DefaultLCSSimilarityOptions},
    {"ttxyzxyz", "XYZtt", 3, testLCSOptions1},
    {"this is very long test case", "another long and also very case test", 16, DefaultLCSSimilarityOptions},
    {"stiohaotirtNOTsoRANDOMiodjofiosahfos", "ASJODJASOIDJOASnotSIADJsoSAIJDOSrandomSDIOJASOD", 9, DefaultLCSSimilarityOptions},
    {"stiohaotirtNOTsoRANDOMiodjofiosahfos", "ASJODJASOIDJOASnotSIADJsoSAIJDOSrandomSDIOJASOD", 21, testLCSOptions1},
    {"abracadabra", "baracadaba", 9, DefaultLCSSimilarityOptions},
    {"ABRACADABRA", "BARACADABA", 9, testLCSOptions1},
    {"human", "chimpanzee", 4, DefaultLCSSimilarityOptions},
    {"nematode knowledge", "empty bottle", 7, DefaultLCSSimilarityOptions},
    {"nEmAtOdE KnOwLeDgE", "eMpTy bOtTlE", 7, testLCSOptions1},
}

func TestLCS(t *testing.T) {
    for _, test := range LCSTests {
        dis := LCS([]rune(test.src), []rune(test.trg), test.opt)
        if dis != test.dis {
            t.Log("LCS distance between", test.src, "and", test.trg, "is", dis, "but should be", test.dis)
            t.Fail()
        }
    }
}
