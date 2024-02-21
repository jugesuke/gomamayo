package nihongo

import "testing"

type testcase struct {
	input  string
	expect string
}

func TestNeologdn(t *testing.T) {
	testcases := [...]testcase{
		{"０１２３４５６７８９", "0123456789"},
		{"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ", "ABCDEFGHIJKLMNOPQRSTUVWXYZ"},
		{"ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ", "abcdefghijklmnopqrstuvwxyz"},
		{"！”＃＄％＆’（）＊＋，－．／：；＜＞？＠［￥］＾＿｀｛｜｝", "!\"#$%&'()*+,-./:;<>?@[¥]^_`{|}"},

		{"｡､･=｢｣", "。、・＝「」"},
		{
			"ｱｲｳｴｵｶｷｸｹｺｻｼｽｾｿﾀﾁﾂﾃﾄﾅﾆﾇﾈﾉﾊﾋﾌﾍﾎﾏﾐﾑﾒﾓﾔﾕﾖﾗﾘﾙﾚﾛﾜｦﾝ",
			"アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲン",
		},
		{
			"ｶﾞｷﾞｸﾞｹﾞｺﾞｻﾞｼﾞｽﾞｾﾞｿﾞﾀﾞﾁﾞﾂﾞﾃﾞﾄﾞﾊﾞﾋﾞﾌﾞﾍﾞﾎﾞﾊﾟﾋﾟﾌﾟﾍﾟﾎﾟ",
			"ガギグゲゴザジズゼゾダヂヅデドバビブベボパピプペポ",
		},
		// replace characters that like two or more occurrences of Japanese Prolonged Sound Mark "ー"
		{"ずーーーーーーしーーーーーーほっきーーーーー", "ずーしーほっきー"},

		// Delete heading Hankaku Space and tailing Hankaku Space
		{"  はこだて", "はこだて"},
		{"はこだて  ", "はこだて"},

		// Delete Hankaku Space among Hiragana, Katakana, and Kanji
		{"情報 表現  入門 の 教科書 を 鍋敷き に する", "情報表現入門の教科書を鍋敷きにする"},
		{"Introduction to Information Expression", "Introduction to Information Expression"},

		// Delete Hankaku Space between Hiragana, Katakana, and Kanji, and Alphabet
		{"高度 ICT 演習", "高度ICT演習"},
		{"Advanced ICT Exercises", "Advanced ICT Exercises"},

		// general test
		{"公立 はこだて　未来  大学 Future　Ｕｎｉｖｅｒｓｉｔｙ Hakodate",
			"公立はこだて未来大学Future University Hakodate"},
		{"公立 はこだて　未来  大学-　Future*　Ｕｎｉｖｅｒｓｉｔｙ+　Hakodate",
			"公立はこだて未来大学-Future*University+Hakodate"},
	}
	for index, testcase := range testcases {
		if s, _ := WithNeologdn(testcase.input); s != testcase.expect {
			t.Errorf("failed: %d expects\"%s\" but output was \"%s\"", index, testcase.expect, s)
		}
	}
}
