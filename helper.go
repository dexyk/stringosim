package stringosim

import (
    "math"
    "regexp"
)

type SimilarityOptions struct {
    InsertCost      int
    DeleteCost      int
    SubstituteCost  int
    CaseInsensitive bool
}

var DefaultSimilarityOptions SimilarityOptions = SimilarityOptions{
    InsertCost:     1,
    DeleteCost:     1,
    SubstituteCost: 1,
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

func GetNGram(s string, n ...int) map[string]int {
    nGram := 1
    if len(n) > 0 {
        for _, v := range n {
            nGram = v
            break
        }
    }

    regExp := regexp.MustCompile(`\s+`)
    t := regExp.ReplaceAll([]byte(s), []byte(" "))
    m := make(map[string]int)
    for i := 0; i <= len(t)-nGram; i++ {
        v := string(t[i:(i + nGram)])
        cnt, ok := m[v]
        if ok {
            m[v] = cnt + 1
        } else {
            m[v] = 1
        }
    }
    return m
}
