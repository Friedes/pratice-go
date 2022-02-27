package main

import "testing"

func TestCase0FindResult(t *testing.T) {
	givenN := 5
	givenPattern := "BAACC"
	wantMaxScore := 3
	wantNames := "Bruno"

	getMaxScore, getNames := findResult(givenN, givenPattern)
	if wantMaxScore != getMaxScore && wantNames != getNames {
		t.Errorf("given %d, %s & want %d, %s but %d, %s\n", givenN, givenPattern, wantMaxScore, wantNames, getMaxScore, getNames)
	}
}

func TestCase1FindResult(t *testing.T) {
	givenN := 5
	givenPattern := "BACCB"
	wantMaxScore := 4
	wantNames := "Bruno"

	getMaxScore, getNames := findResult(givenN, givenPattern)
	if wantMaxScore != getMaxScore && wantNames != getNames {
		t.Errorf("given %d, %s & want %d, %s but %d, %s\n", givenN, givenPattern, wantMaxScore, wantNames, getMaxScore, getNames)
	}
}

func TestCase2FindResult(t *testing.T) {
	givenN := 7
	givenPattern := "CCABACB"
	wantMaxScore := 3
	wantNames := "Goran"

	getMaxScore, getNames := findResult(givenN, givenPattern)
	if wantMaxScore != getMaxScore && wantNames != getNames {
		t.Errorf("given %d, %s & want %d, %s but %d, %s\n", givenN, givenPattern, wantMaxScore, wantNames, getMaxScore, getNames)
	}
}

func TestCase3FindResult(t *testing.T) {
	givenN := 9
	givenPattern := "CBBBACBBC"
	wantMaxScore := 4
	wantNames := "Adrian"

	getMaxScore, getNames := findResult(givenN, givenPattern)
	if wantMaxScore != getMaxScore && wantNames != getNames {
		t.Errorf("given %d, %s & want %d, %s but %d, %s\n", givenN, givenPattern, wantMaxScore, wantNames, getMaxScore, getNames)
	}
}

func TestCase4FindResult(t *testing.T) {
	givenN := 10
	givenPattern := "AAABCBACCB"
	wantMaxScore := 3
	wantNames := "AdrianGoran"

	getMaxScore, getNames := findResult(givenN, givenPattern)
	if wantMaxScore != getMaxScore && wantNames != getNames {
		t.Errorf("given %d, %s & want %d, %s but %d, %s\n", givenN, givenPattern, wantMaxScore, wantNames, getMaxScore, getNames)
	}
}

func TestCase5FindResult(t *testing.T) {
	givenN := 15
	givenPattern := "AAACBCCAAABCAAA"
	wantMaxScore := 7
	wantNames := "AdrianBrunoGoran"

	getMaxScore, getNames := findResult(givenN, givenPattern)
	if wantMaxScore != getMaxScore && wantNames != getNames {
		t.Errorf("given %d, %s & want %d, %s but %d, %s\n", givenN, givenPattern, wantMaxScore, wantNames, getMaxScore, getNames)
	}
}

func TestCase6FindResult(t *testing.T) {
	givenN := 20
	givenPattern := "CCCACCACABBACACABABA"
	wantMaxScore := 9
	wantNames := "Goran"

	getMaxScore, getNames := findResult(givenN, givenPattern)
	if wantMaxScore != getMaxScore && wantNames != getNames {
		t.Errorf("given %d, %s & want %d, %s but %d, %s\n", givenN, givenPattern, wantMaxScore, wantNames, getMaxScore, getNames)
	}
}

func TestCase7FindResult(t *testing.T) {
	givenN := 38
	givenPattern := "AAAACCCCCAABACCBCCBBCBABCBCCBBCCCAACAC"
	wantMaxScore := 17
	wantNames := "Adrian"

	getMaxScore, getNames := findResult(givenN, givenPattern)
	if wantMaxScore != getMaxScore && wantNames != getNames {
		t.Errorf("given %d, %s & want %d, %s but %d, %s\n", givenN, givenPattern, wantMaxScore, wantNames, getMaxScore, getNames)
	}
}

func TestCase8FindResult(t *testing.T) {
	givenN := 69
	givenPattern := "BAABBBABCBBACABACAABABACBAACBAACAAAACCBBBBBBCCCBCBAACABAAAACCAAAAAABB"
	wantMaxScore := 25
	wantNames := "BrunoGoran"

	getMaxScore, getNames := findResult(givenN, givenPattern)
	if wantMaxScore != getMaxScore && wantNames != getNames {
		t.Errorf("given %d, %s & want %d, %s but %d, %s\n", givenN, givenPattern, wantMaxScore, wantNames, getMaxScore, getNames)
	}
}

func TestCase9FindResult(t *testing.T) {
	givenN := 99
	givenPattern := "AAACAAACAACCACBBBCAAABACBCCACBCCAAACBABBABBAACCABCAABBCAABBCACCBCCAACCCCCBAABACBAACBAAACCCBBCBBCCAC"
	wantMaxScore := 36
	wantNames := "Goran"

	getMaxScore, getNames := findResult(givenN, givenPattern)
	if wantMaxScore != getMaxScore && wantNames != getNames {
		t.Errorf("given %d, %s & want %d, %s but %d, %s\n", givenN, givenPattern, wantMaxScore, wantNames, getMaxScore, getNames)
	}
}

func TestCase10FindResult(t *testing.T) {
	givenN := 100
	givenPattern := "BBCCBAAAABCACCABCCACBBAABCACBAABABACACBBCBAACBAAACBAAAACACABBABAAAABCCCABBACBBBCACCCCACABCBBBBCCBBCB"
	wantMaxScore := 31
	wantNames := "AdrianBruno"

	getMaxScore, getNames := findResult(givenN, givenPattern)
	if wantMaxScore != getMaxScore && wantNames != getNames {
		t.Errorf("given %d, %s & want %d, %s but %d, %s\n", givenN, givenPattern, wantMaxScore, wantNames, getMaxScore, getNames)
	}
}
