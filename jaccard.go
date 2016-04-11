package stringosim

import ()

func Jaccard(s []rune, t []rune, n ...int) float64 {
    nGram := 1
    if len(n) > 0 {
        for _, v := range n {
            nGram = v
            break
        }
    }
    sGrams := GetNGram(string(s), nGram)
    tGrams := GetNGram(string(t), nGram)

    total := len(sGrams) + len(tGrams)
    intersection := 0
    for k, _ := range sGrams {
        _, ok := tGrams[k]
        if ok {
            intersection++
        }
    }
    return 1.0 - float64(intersection)/float64(total-intersection)
}
