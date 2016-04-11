package stringosim

import (
    "unicode"
)

func Levenshtein(s []rune, t []rune, options ...SimilarityOptions) int {
    changeCost := DefaultSimilarityOptions.SubstituteCost
    deleteCost := DefaultSimilarityOptions.DeleteCost
    insertCost := DefaultSimilarityOptions.InsertCost
    caseInsensitive := DefaultSimilarityOptions.CaseInsensitive
    if len(options) > 0 {
        for _, option := range options {
            changeCost = option.SubstituteCost
            insertCost = option.InsertCost
            deleteCost = option.DeleteCost
            caseInsensitive = option.CaseInsensitive
            break
        }
    }
    d := make([]int, len(t)+1)
    for i := 0; i <= len(t); i++ {
        d[i] = i * insertCost
    }

    for is, cs := range s {
        tmpD := d[0]
        d[0] = (is + 1) * deleteCost
        for it, ct := range t {
            curChangeCost := changeCost
            if cs == ct || (caseInsensitive && unicode.ToLower(cs) == unicode.ToLower(ct)) {
                curChangeCost = 0
            }
            nextD := Min(Min(d[it+1]+deleteCost, d[it]+insertCost), tmpD+curChangeCost)
            tmpD, d[it+1] = d[it+1], nextD
        }
    }
    return d[len(t)]
}
