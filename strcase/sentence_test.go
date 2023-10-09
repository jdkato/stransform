package strcase_test

import (
	"testing"

	"github.com/jdkato/stransform/strcase"
)

var cases = []testCase{
	{"", ""},
	{"1. An important heading", "1. An important heading"},
	{"getting started with vale server", "Getting started with vale server"},
	{"Lession 1: getting started with vale server", "Lession 1: Getting started with vale server"},
}

var vocabCases = []testCase{
	{"Getting started with vale server", "Getting started with Vale Server"},
}

func TestSentence(t *testing.T) {
	tc := strcase.NewSentenceConverter()
	for _, test := range cases {
		sent := tc.Sentence(test.Input)
		if test.Expect != sent {
			t.Fatalf("Got '%s'; expected '%s'", sent, test.Expect)
		}
	}
}

func TestVocab(t *testing.T) {
	tc := strcase.NewSentenceConverter(strcase.UsingVocab([]string{
		"Vale Server",
	}))

	for _, test := range vocabCases {
		sent := tc.Sentence(test.Input)
		if test.Expect != sent {
			t.Fatalf("Got '%s'; expected '%s'", sent, test.Expect)
		}
	}
}