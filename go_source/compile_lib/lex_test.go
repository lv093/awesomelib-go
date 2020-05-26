package compile_lib

import "testing"

//test lex
func TestLex(test *testing.T) {
	type lexTestPair struct {
		name   string //for debug
		errs   []string
		input  string
		output []int
	}

	var lexTestPairs = []lexTestPair{
		{
			"Normal Test 1",
			[]string{},
			`tree:
	-> black
		-> red
		-> red`,
			[]int{kTREE, tCOLON, tTAG, tRIGHT_ARROW, tIDENT,
				tTAG, tTAG, tRIGHT_ARROW, tIDENT,
				tTAG, tTAG, tRIGHT_ARROW, tIDENT}},
		{"Number test 1",
			[]string{}, `-120.100`, []int{tNUM},
		},
		{
			"Exception Test 1",
			[]string{"unexpected '<' after '-' at line 2,column 4"},
			`tree:
	a-<`,
			[]int{kTREE, tCOLON, tTAG, tIDENT},
		},
		{
			"Exception Test 2",
			[]string{"unkown char '*' at line 2,column 1"},
			`tree
*-> black`,
			[]int{kTREE},
		}, {
			"Exception Test 3",
			[]string{"unkown char '\\' at line 5,column 3"},
			`
	->black
		-> red
		\`,
			[]int{tTAG, tRIGHT_ARROW, tIDENT, tTAG, tTAG, tRIGHT_ARROW, tIDENT, tTAG, tTAG},
		},
	}

	for _, ltp := range lexTestPairs {
		lex := newLexer(ltp.input)
		i := 0
		for t := lex.token(); t.typ != tEOF; t = lex.token() {
			if ltp.output[i] != t.typ {
				test.Fatalf("%s(wanted) not equal %s in %s\n", tokens[ltp.output[i]], tokens[t.typ], ltp.name)
			}
			i++
		}
		for i := 0; i < len(ltp.errs); i++ {
			if i > len(lex.errors)-1 {
				test.Fatalf("%s(wanted error) equal nothing in %s", ltp.errs[i], ltp.name)
				break
			}
			if ltp.errs[i] != lex.errors[i] {
				test.Fatalf("%s(wanted error) not equal %s in %s\n", ltp.errs[i], lex.errors[i], ltp.name)
			}

		}
	}
}