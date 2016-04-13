package stringosim

import (
    "errors"
)

var HAMMING_ERROR_DIFFERENT_LENGTH = errors.New("Can't compare strings of different lengths")

type HammingSimilarityOptions struct {
    CaseInsensitive bool
}

var DefaultHammingSimilarityOptions = HammingSimilarityOptions{
    CaseInsensitive: false,
}

func Hamming(s []rune, t []rune, options ...HammingSimilarityOptions) (int, error) {
    if len(s) != len(t) {
        return -1, HAMMING_ERROR_DIFFERENT_LENGTH
    }
    opt := DefaultHammingSimilarityOptions
    for _, option := range options {
        opt = option
        break
    }
    ret := 0
    for i, cs := range s {
        if !SameRune(cs, t[i], opt.CaseInsensitive) {
            ret++
        }
    }
    return ret, nil
}
