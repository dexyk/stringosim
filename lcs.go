package stringosim

import ()

type LCSSimilarityOptions struct {
    CaseInsensitive bool
}

var DefaultLCSSimilarityOptions = LCSSimilarityOptions{
    CaseInsensitive: false,
}

func LCS(s []rune, t []rune, options ...LCSSimilarityOptions) int {
    opt := DefaultLCSSimilarityOptions
    for _, option := range options {
        opt = option
        break
    }
    d := make([][]int, len(s)+1)
    for i := 0; i <= len(s); i++ {
        d[i] = make([]int, len(t)+1)
    }
    for j := 0; j <= len(t); j++ {
        d[0][j] = 0
    }
    for i := 1; i <= len(s); i++ {
        d[i][0] = 0
        for j := 1; j <= len(t); j++ {
            if SameRune(s[i-1], t[j-1], opt.CaseInsensitive) {
                d[i][j] = d[i-1][j-1] + 1
            } else {
                d[i][j] = Max(d[i-1][j], d[i][j-1])
            }
        }
    }

    return d[len(s)][len(t)]
}
