package stringosim

import ()

type JaroSimilarityOptions struct {
    Threshold       float64
    PValue          float64
    LValue          float64
    CaseInsensitive bool
}

var DefaultJaroSimilarityOptions = JaroSimilarityOptions{
    CaseInsensitive: false,
    Threshold:       0.7,
    PValue:          0.1,
    LValue:          4,
}

func Jaro(s []rune, t []rune, options ...JaroSimilarityOptions) float64 {
    opt := DefaultJaroSimilarityOptions
    for _, option := range options {
        opt = option
    }
    lenMatched, numT, _ := jaroHelper(s, t, opt)
    if lenMatched == 0 {
        return 0.0
    }
    lenS := len(s)
    lenT := len(t)
    return jaroFormula(lenMatched, numT, lenS, lenT)
}

func JaroWinkler(s []rune, t []rune, options ...JaroSimilarityOptions) float64 {
    opt := DefaultJaroSimilarityOptions
    for _, option := range options {
        opt = option
    }
    lenMatched, numT, prefixLen := jaroHelper(s, t, opt)
    if lenMatched == 0 {
        return 0.0
    }
    lenS := len(s)
    lenT := len(t)
    jaroDis := jaroFormula(lenMatched, numT, lenS, lenT)
    if jaroDis < opt.Threshold {
        return jaroDis
    }
    p := opt.PValue
    if p*float64(prefixLen) > 1.0 {
        p = 1.0 / float64(prefixLen)
    }
    return jaroDis + (1.0-jaroDis)*p*float64(prefixLen)
}

func jaroHelper(s []rune, t []rune, option JaroSimilarityOptions) (int, int, int) {
    lenS := len(s)
    lenT := len(t)
    checkS := make([]rune, 0, len(s))
    checkT := make([]rune, 0, len(t))
    matchedT := make([]bool, len(t))
    maxDis := Max(lenS, lenT) / 2
    for is, cs := range s {
        for it := Max(0, is-maxDis); it <= Min(lenT-1, is+maxDis); it++ {
            if !matchedT[it] && SameRune(cs, t[it], option.CaseInsensitive) {
                matchedT[it] = true
                checkS = append(checkS, cs)
                break
            }
        }
    }
    for it, ct := range t {
        if matchedT[it] {
            checkT = append(checkT, ct)
        }
    }
    minLen := Min(lenS, lenT)
    prefixLen := 0
    for i := 0; i < minLen; i++ {
        if !SameRune(s[i], t[i], option.CaseInsensitive) {
            prefixLen = i
            break
        }
    }

    numTranspositions := 0
    for i, cs := range checkS {
        if !SameRune(cs, checkT[i], option.CaseInsensitive) {
            numTranspositions++
        }
    }
    if len(checkS) == 0 {
        return 0, 0, 0
    }
    return len(checkS), numTranspositions / 2, prefixLen
}

func jaroFormula(m, t, s1, s2 int) float64 {
    return (float64(m)/float64(s1) + float64(m)/float64(s2) + float64(m-t)/float64(m)) / 3.0
}
