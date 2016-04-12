#stringosim

The plan for this package is to have Go implementation of different string distance/similarity functions, like Levenshtein (normalized, weighted, Damerau), Jaro-Winkler, Jaccard index, Euclidean distance, Hamming distance...

Currently it has implemented Levenshtein, Jaccard and Hamming distance functions. Work in progress...

##Import and installation

To get the library just run:
```
go get github.com/dexyk/stringosim
```

To use the library just import it in your code:

```go
import "github.com/dexyk/stringosim"
```

##Usage

Currently only Levenshtein, Jaccard and Hamming string distances are implemented.

####Levenshtein

Levenshtein distance can be calculated with default parameters (use DefaultSimilarityOptions) where cost of insert, delete and substitute operation are 1. You can also use it with other parameters by using SimilarityOptions type. Setting CaseInsensitive to true in SimilarityOptions the comparison will be done without considering character cases. 

Example:
```go
fmt.Println(stringosim.Levenshtein([]rune("stringosim"), []rune("stingobim")))

fmt.Println(stringosim.Levenshtein([]rune("stringosim"), []rune("stingobim"),
    stringosim.LevenshteinSimilarityOptions{
        InsertCost:     3,
        DeleteCost:     5,
        SubstituteCost: 2,
}))

fmt.Println(stringosim.Levenshtein([]rune("stringosim"), []rune("STRINGOSIM"),
    stringosim.LevenshteinSimilarityOptions{
        InsertCost:      3,
        DeleteCost:      4,
        SubstituteCost:  5,
        CaseInsensitive: true,
}))
```

####Jaccard

Jaccard distance can be calculated by setting the size of the n-gram which will be used for comparison. If the size is omitted the default value of 1 will be used.

Example:
```go
fmt.Println(stringosim.Jaccard([]rune("stringosim"), []rune("stingobim")))

fmt.Println(stringosim.Jaccard([]rune("stringosim"), []rune("stingobim"), 2))

fmt.Println(stringosim.Jaccard([]rune("stringosim"), []rune("stingobim"), 3))
```

####Hamming

Hamming distance can be calculated with options. Default function will calculate standard hamming distance with case sensitive option. It can be also used without case sensitive option.

If the strings to compare have different lengths, the error will be returned.

Example:
```go
dis, _ := stringosim.Hamming([]rune("testing"), []rune("restink"))
fmt.Println(dis)

dis, _ = stringosim.Hamming([]rune("testing"), []rune("FESTING"), stringosim.HammingSimilarityOptions{
    CaseInsensitive: true,
})
fmt.Println(dis)

_, err := stringosim.Hamming([]rune("testing"), []rune("testin"))
fmt.Println(err)
```



