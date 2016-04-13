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
