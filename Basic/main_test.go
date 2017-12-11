package main

import (
	"strings"
	"testing"
)

func TestTokenize1(t *testing.T) {
	s := "100 PRINT \"SIN(X)^2 = \", SIN(X) ^ 2"
	tokens := tokenize(s)

	if tokens == nil {
		t.Errorf("nil tokens")
	}

	if len(tokens) != 10 {
		t.Errorf("Expected 10 tokens but got: %v => %v", len(tokens), tokens)
	}

	if tokens[0] != "100" {
		t.Errorf("Expected 100	 but got: %v", tokens[0])
	}
}

func TestRemoveSpaces(t *testing.T) {
	s := removeSpaces(`10 GO TO 99`)
	if s != "10GOTO99" {
		t.Errorf("Expected 10GOTO99 but got => %v", s)
	}

	s = removeSpaces(`100 PRINT "HELLO WORLD", SIN(X) ^ 2`)
	exp := `100PRINT"HELLO WORLD",SIN(X)^2`

	if len(s) != len(exp) {
		t.Errorf("%s\n%s", s, exp)
	}
	if s != exp {
		t.Errorf("\n%v|\n%v|", exp, s)
	}

}
func TestTokenize2(t *testing.T) {
	s := "10 READ N"
	tokens := tokenize(s)
	expected := []string{
		"10", "READ", "N",
	}

	for i, k := range expected {
		if tokens[i] != k {
			t.Errorf("Expected *v\nbut got: %v", tokens)
		}
	}
}

func TestTokenizer(t *testing.T) {
	line := "X-1"
	tokens := tokenizer(line)

	if len(tokens) != 3 {
		t.Errorf("Expected 10 tokens but got: %v => %v", len(tokens), tokens)
	}

	if tokens[0] != "X" || tokens[1] != "-" || tokens[2] != "1" {
		t.Errorf("EXPECTED %v but got %v", line, tokens)
	}

	line = `PRINT "HELLO WORLD"`
	tokens = tokenizer(line)
	if len(tokens) != 2 {
		t.Errorf("Expected 2 tokens but got: %v => %v", len(tokens), tokens)
	}
	if tokens[0] != "PRINT" || tokens[1] != `"HELLO WORLD"` {
		t.Errorf("EXPECTED %v but got %v", line, tokens)
	}

	line = `100IFX1+123.4+E1-12.3E4 <> 1.2E-34*-12E34+1+"HI" THEN99`
	tokens = tokenizer(line)
	if len(tokens) != 20 {
		t.Errorf("Expected 20 tokens but got: %v => %v", len(tokens), tokens)
	}
	exp := []string{
		"100", "IF", "X1", "+", "123.4", "+", "E1", "-", "12.3E4", "<>",
		"1.2E-34", "*", "-", "12E34", "+", "1", "+", `"HI"`, "THEN", "99",
	}

	if strings.Join(exp, "") != strings.Join(tokens, "") {
		t.Errorf("slice we expected is not the same as the tokens parsed")
	}
}
