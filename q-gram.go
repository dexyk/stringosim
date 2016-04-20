package stringosim

import (
    "strings"
)

type QGramSimilarityOptions struct {
    CaseInsensitive bool
    NGramLength     int
}

var DefaultQGramSimilarityOptions = QGramSimilarityOptions{
    CaseInsensitive: false,
    NGramLength:     2,
}

func QGram(s []rune, t []rune, options ...QGramSimilarityOptions) int {
    opt := DefaultQGramSimilarityOptions
    for _, option := range options {
        opt = option
        break
    }
    var sGrams, tGrams map[string]int
    if opt.CaseInsensitive {
        sGrams = GetNGram(strings.ToLower(string(s)), opt.NGramLength)
        tGrams = GetNGram(strings.ToLower(string(t)), opt.NGramLength)
    } else {
        sGrams = GetNGram(string(s), opt.NGramLength)
        tGrams = GetNGram(string(t), opt.NGramLength)
    }
    dis := 0
    for k, vs := range sGrams {
        vt, ok := tGrams[k]
        if ok {
            dis += AbsInt(vs - vt)
        } else {
            dis += vs
        }
    }
    for k, vt := range tGrams {
        _, ok := sGrams[k]
        if !ok {
            dis += vt
        }
    }
    return dis
}
