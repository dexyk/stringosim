package stringosim

import (
    "testing"
)

var testJaroOptions1 JaroSimilarityOptions = JaroSimilarityOptions{
    CaseInsensitive: true,
    Threshold:       0.7,
    PValue:          0.1,
    LValue:          4,
}

type JaroTest struct {
    src string
    trg string
    dis float64
    opt JaroSimilarityOptions
}

var JaroTests = []JaroTest{
    {"", "", 0.0, DefaultJaroSimilarityOptions},
    {"x", "", 0.0, DefaultJaroSimilarityOptions},
    {"xyzxyztt", "ttxyz", 0.6583333333333333, DefaultJaroSimilarityOptions},
    {"ttxyzxyz", "XYZtt", 0.5499999999999999, DefaultJaroSimilarityOptions},
    {"this is very long test case", "another long and also very case test", 0.699074074074074, DefaultJaroSimilarityOptions},
    {"stiohaotirtNOTsoRANDOMiodjofiosahfos", "ASJODJASOIDJOASnotSIADJsoSAIJDOSrandomSDIOJASOD", 0.39066193853427894, DefaultJaroSimilarityOptions},
    {"stiohaotirtNOTsoRANDOMiodjofiosahfos", "ASJODJASOIDJOASnotSIADJsoSAIJDOSrandomSDIOJASOD", 0.6602217719238995, testJaroOptions1},
    {"abracadabra", "baracadaba", 0.9363636363636364, DefaultJaroSimilarityOptions},
    {"abracadabra", "BARACADABA", 0.0, DefaultJaroSimilarityOptions},
    {"abracadabra", "BARACADABA", 0.9363636363636364, testJaroOptions1},
    {"human", "chimpanzee", 0.7333333333333334, DefaultJaroSimilarityOptions},
    {"nematode knowledge", "empty bottle", 0.5621693121693122, DefaultJaroSimilarityOptions},
    {"nEmAtOdE KnOwLeDgE", "eMpTy bOtTlE", 0.5621693121693122, testJaroOptions1},
}

var JaroWinklerTests = []JaroTest{
    {"", "", 0.0, DefaultJaroSimilarityOptions},
    {"x", "", 0.0, DefaultJaroSimilarityOptions},
    {"xyzxyztt", "ttxyz", 0.6583333333333333, DefaultJaroSimilarityOptions},
    {"ttxyzxyz", "XYZtt", 0.5499999999999999, DefaultJaroSimilarityOptions},
    {"this is very long test case", "another long and also very case test", 0.699074074074074, DefaultJaroSimilarityOptions},
    {"stiohaotirtNOTsoRANDOMiodjofiosahfos", "ASJODJASOIDJOASnotSIADJsoSAIJDOSrandomSDIOJASOD", 0.39066193853427894, DefaultJaroSimilarityOptions},
    {"stiohaotirtNOTsoRANDOMiodjofiosahfos", "ASJODJASOIDJOASnotSIADJsoSAIJDOSrandomSDIOJASOD", 0.6602217719238995, testJaroOptions1},
    {"abracadabra", "baracadaba", 0.9363636363636364, DefaultJaroSimilarityOptions},
    {"abracadabra", "BARACADABA", 0.0, DefaultJaroSimilarityOptions},
    {"abracadabra", "BARACADABA", 0.9363636363636364, testJaroOptions1},
    {"human", "chimpanzee", 0.7333333333333334, DefaultJaroSimilarityOptions},
    {"nematode knowledge", "empty bottle", 0.5621693121693122, DefaultJaroSimilarityOptions},
    {"nEmAtOdE KnOwLeDgE", "eMpTy bOtTlE", 0.5621693121693122, testJaroOptions1},
}

func TestJaro(t *testing.T) {
    for _, test := range JaroTests {
        dis := Jaro([]rune(test.src), []rune(test.trg), test.opt)
        if !EqualFloat64(dis, test.dis) {
            t.Log("Jaro distance between", test.src, "and", test.trg, "is", dis, "but should be", test.dis)
            t.Fail()
        }
    }
}

func TestJaroWinkler(t *testing.T) {
    for _, test := range JaroWinklerTests {
        dis := JaroWinkler([]rune(test.src), []rune(test.trg), test.opt)
        if !EqualFloat64(dis, test.dis) {
            t.Log("Jaro-Winkler distance between", test.src, "and", test.trg, "is", dis, "but should be", test.dis)
            t.Fail()
        }
    }
}
