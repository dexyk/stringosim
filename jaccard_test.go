package stringosim

import (
	"testing"
)

type JaccardTest struct {
	src        string
	trg        string
	NGramSizes []int
	dis        float64
}

var jaccardTests = []JaccardTest{
	{"abracadabra", "baccarda", []int{1}, 0.0},
	{"abracadabra", "baccarda", []int{2}, 0.7272727272727273},
	{"abracadabra", "baccarda", []int{3}, 1.0},
	{"this    is     space    test", "this is space test", []int{3}, 0.0},
	{"just another test of jaccard", "i will test jaccard", []int{1}, 0.47058823529411764},
	{"just another test of jaccard", "i will test jaccard", []int{2}, 0.6129032258064516},
	{"just another test of jaccard", "i will test jaccard", []int{3}, 0.6875},
	{"book is on the shelv", "buk is on the shelf", []int{1}, 0.2142857142857143},
	{"book is on the shelv", "buk is on the shelf", []int{2}, 0.33333333333333337},
	{"book is on the shelv", "buk is on the shelf", []int{3}, 0.33333333333333337},
	{"cardiogram", "krdiogram", []int{1}, 0.2222222222222222},
	{"cardiogram", "krdiogram", []int{2}, 0.30000000000000004},
	{"cardiogram", "krdiogram", []int{3}, 0.33333333333333337},
}

func TestJaccard(t *testing.T) {
	for _, test := range jaccardTests {
		dis := Jaccard([]rune(test.src), []rune(test.trg), test.NGramSizes)
		if !EqualFloat64(dis, test.dis) {
			t.Log("Jaccard distance between", test.src, "and", test.trg, "is", dis, "but should be", test.dis)
			t.Fail()
		}

	}
}
