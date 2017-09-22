package stringosim

import (
	"math"
	"regexp"
	"unicode"
)

type SimilarityOptions struct {
	CaseInsensitive bool
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

const EPS = 0.000000001

func EqualFloat64(x, y float64) bool {
	return math.Abs(x-y) < EPS
}

func GetNGram(s string, NGramSizes []int) map[string]int {
	regExp := regexp.MustCompile(`\s+`)
	t := regExp.ReplaceAllString(s, " ")
	m := make(map[string]int)
	for _, nGram := range NGramSizes {
		for i := 0; i <= len(t)-nGram; i++ {
			v := string(t[i:(i + nGram)])
			cnt, ok := m[v]
			if ok {
				m[v] = cnt + 1
			} else {
				m[v] = 1
			}
		}
	}
	return m
}

func CompareErrors(e1 error, e2 error) bool {
	if e1 == nil {
		if e2 == nil {
			return true
		} else {
			return false
		}
	} else {
		if e2 == nil {
			return false
		} else {
			return e1.Error() == e2.Error()
		}
	}
}

func SameRune(a rune, b rune, caseInsensitive bool) bool {
	return a == b || (caseInsensitive && unicode.ToLower(a) == unicode.ToLower(b))
}

func AbsInt(v int) int {
	if v > 0 {
		return v
	} else {
		return -v
	}
}

func DotProductNGrams(m1, m2 map[string]int) int {
	ret := 0
	for k1, v1 := range m1 {
		v2, ok := m2[k1]
		if ok {
			ret += v1 * v2
		}
	}
	return ret
}

func NormNGram(m map[string]int) float64 {
	ret := 0.0
	for _, v := range m {
		ret += float64(v) * float64(v)
	}
	return math.Sqrt(ret)
}
