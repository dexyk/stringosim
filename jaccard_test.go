package stringosim

import (
    "testing"
)

type JaccardTest struct {
    src   string
    trg   string
    nGram int
    dis   float64
}

var jaccardTests = []JaccardTest{
    {"abracadabra", "baccarda", 1, 0.0},
    {"abracadabra", "baccarda", 2, 0.7272727272727273},
    {"abracadabra", "baccarda", 3, 1.0},
    {"this    is     space    test", "this is space test", 3, 0.0},
    {"just another test of jaccard", "i will test jaccard", 1, 0.47058823529411764},
    {"just another test of jaccard", "i will test jaccard", 2, 0.6129032258064516},
    {"just another test of jaccard", "i will test jaccard", 3, 0.6875},
    {"book is on the shelv", "buk is on the shelf", 1, 0.2142857142857143},
    {"book is on the shelv", "buk is on the shelf", 2, 0.33333333333333337},
    {"book is on the shelv", "buk is on the shelf", 3, 0.33333333333333337},
    {"cardiogram", "krdiogram", 1, 0.2222222222222222},
    {"cardiogram", "krdiogram", 2, 0.30000000000000004},
    {"cardiogram", "krdiogram", 3, 0.33333333333333337},
}

func TestJaccard(t *testing.T) {
    for _, test := range jaccardTests {
        dis := Jaccard([]rune(test.src), []rune(test.trg), test.nGram)
        if !EqualFloat64(dis, test.dis) {
            t.Log("Jaccard distance between", test.src, "and", test.trg, "is", dis, "but should be", test.dis)
            t.Fail()
        }

    }
}
