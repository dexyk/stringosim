#stringosim

The plan for this package is to have Go implementation of different string distance/similarity functions, like Levenshtein (normalized, weighted, Damerau), Jaro-Winkler, Jaccard index, Euclidean distance, Hamming distance...

Currently it is empty. Work in progress...

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

Currently only Levenshtein and Jaccard string distances are implemented.

####Levenshtein

Levenshtein distance can be calculated with default parameters (use DefaultSimilarityOptions) where cost of insert, delete and substitute operation are 1. You can also use it with other parameters by using SimilarityOptions type. Setting CaseInsensitive to true in SimilarityOptions the comparison will be done without considering character cases. 

Example:
```go
fmt.Println(stringosim.Levenshtein([]rune("stringosim"), []rune("stingobim")))

fmt.Println(stringosim.Levenshtein([]rune("stringosim"), []rune("stingobim"),
    stringosim.SimilarityOptions{
        InsertCost:     3,
        DeleteCost:     5,
        SubstituteCost: 2,
}))

fmt.Println(stringosim.Levenshtein([]rune("stringosim"), []rune("STRINGOSIM"),
    stringosim.SimilarityOptions{
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




