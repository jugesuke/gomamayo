package gomamayo

import (
	"github.com/ikawaha/kagome-dict/uni"
	"github.com/ikawaha/kagome/v2/tokenizer"
	"github.com/jugesuke/gomamayo/pkg/nihongo"
)

type (
	// Result has a result of IsDajare function.
	Result struct {
		// If it is Dajare, this field is True, else False.
		IsGomamayo bool

		Ary    int // n項
		Degree int // n次

		Sentence string
		Reading  string

		Tokens []string
	}

	GoMamayo struct {
		*tokenizer.Tokenizer

		readPos int
	}
)

// Set dictionary. You must do this first.
// Set dictionary you like. You can use a Kagome Dictionary.
// https://github.com/ikawaha/kagome#dictionaries
func Init() (*GoMamayo, error) {
	g := new(GoMamayo)
	g.readPos = 6
	dic := uni.Dict()

	if t, err := tokenizer.New(dic, tokenizer.OmitBosEos()); err != nil {
		return nil, err
	} else {
		g.Tokenizer = t
		// g.Analyze("")
		return g, nil
	}
}

// Analyze checks if a sentence is Dajare.
func (g *GoMamayo) Analyze(s string) *Result {
	s, _ = nihongo.Normalize(s, nihongo.WithNeologdn)
	tokens := g.Tokenize(s)

	r := new(Result)

	for _, t := range tokens {
		f := t.Features()
		if len(f) <= g.readPos {
			return r
		}

		r.Tokens = append(r.Tokens, f[g.readPos])
	}

	for position := 0; position < len(r.Tokens)-1; position += 1 {
		former := r.Tokens[position]
		formerMora := nihongo.DivideMora(former)
		formerLen := len(formerMora)

		later := r.Tokens[position+1]
		laterMora := nihongo.DivideMora(later)
		laterLen := len(laterMora)

		deg := 0
		isGomamayo := false

		for maxDeg := 0; maxDeg <= min(formerLen, laterLen); maxDeg += 1 {
			flag := true
			for searchPos := 0; searchPos < maxDeg; searchPos += 1 {
				if formerMora[formerLen-maxDeg+searchPos] == laterMora[searchPos] {
					deg += 1
				} else {
					flag = false
					deg = 0
					break
				}
			}

			if flag && deg != 0 {
				isGomamayo = true
				r.Degree = max(deg, r.Degree)
			}
			deg = 0
		}

		if isGomamayo {
			r.Ary += 1
		}
	}

	if r.Ary > 0 {
		r.IsGomamayo = true
	}

	return r
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
