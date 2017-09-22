package stringosim

func Jaccard(s []rune, t []rune, NGramSizes []int) float64 {
	sGrams := GetNGram(string(s), NGramSizes)
	tGrams := GetNGram(string(t), NGramSizes)

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
