package gomamayo

import (
	"fmt"
	"testing"
)

type testcase struct {
	input  string
	expect *Result
}

func TestAnalyze(t *testing.T) {
	g, err := Init()
	if err != nil {
		panic(err)
	}

	trueGomamayo := []testcase{
		{
			input: "ゴママヨサラダ",
			expect: &Result{
				IsGomamayo: true,
				Ary:        1,
				Degree:     1,
			},
		},
		{
			input: "ゴママヨ",
			expect: &Result{
				IsGomamayo: true,
				Ary:        1,
				Degree:     1,
			},
		},
		{
			input: "サイレンススズカ",
			expect: &Result{
				IsGomamayo: true,
				Ary:        1,
				Degree:     1,
			},
		},
		{
			input: "ハースストーン",
			expect: &Result{
				IsGomamayo: true,
				Ary:        1,
				Degree:     1,
			},
		},
		{
			input: "部分分数分解",
			expect: &Result{
				IsGomamayo: true,
				Ary:        1,
				Degree:     2,
			},
		},
		{
			input: "オレンジレンジ",
			expect: &Result{
				IsGomamayo: true,
				Ary:        1,
				Degree:     3,
			},
		},
		{
			input: "太鼓公募募集終了",
			expect: &Result{
				IsGomamayo: true,
				Ary:        3,
				Degree:     2,
			},
		},
		{
			input: "多項高次ゴママヨ",
			expect: &Result{
				IsGomamayo: true,
				Degree:     2,
				Ary:        2,
			},
		},
		{
			input: "福山雅治",
			expect: &Result{
				IsGomamayo: true,
				Degree:     1,
				Ary:        1,
			},
		},
		{
			input: "長期金利",
			expect: &Result{
				IsGomamayo: true,
				Degree:     1,
				Ary:        1,
			},
		},
		{
			input: "株式公開買い付け",
			expect: &Result{
				IsGomamayo: true,
				Ary:        1,
				Degree:     2,
			},
		},
	}
	falseGomamayo := []string{
		"パパイヤ",
		"パパイヤジュース",
		"オレンジジュース",
	}

	for _, sample := range trueGomamayo {
		if r := g.Analyze(sample.input); !r.IsGomamayo {
			fmt.Printf("%#v\n", r)
			t.Errorf("%s is Gomamayo", sample.input)
		} else if r.Degree != sample.expect.Degree || r.Ary != sample.expect.Ary {
			fmt.Printf("%#v\n", r)
			t.Errorf("%s is %d ary %d degree Gomamayo", sample.input, sample.expect.Ary, sample.expect.Degree)
		}
	}

	for _, sample := range falseGomamayo {
		if r := g.Analyze(sample); r.IsGomamayo {
			fmt.Printf("%#v\n", r)
			t.Errorf("%s is not Gomamayo", sample)
		}
	}
}
