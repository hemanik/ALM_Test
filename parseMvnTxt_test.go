package main

import (
	"reflect"
	"testing"
)

func TestParseTxt(t *testing.T) {

	sampleFile := "test.txt"

	m1, _ := parseTxt(sampleFile)
	m2, _ := parseTxt(sampleFile)

	if reflect.DeepEqual(m1, m2) == false {
		t.Fail()
	}
}
