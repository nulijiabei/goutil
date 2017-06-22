package goutil

import (
	"testing"
	"unicode/utf8"
)

func TestString2Float(t *testing.T) {
	v := String2Float("3.1415926", -1)
	if v != 3.1415926 {
		t.Fail()
	}
}

func TestString2Int(t *testing.T) {
	v := String2Int("123", -1)
	if v != 123 {
		t.Fail()
	}
}

func TestString2Int64(t *testing.T) {
	v := String2Int64("123456", -1)
	if v != 123456 {
		t.Fail()
	}
}

func TestString2Utf8(t *testing.T) {
	v := String2Utf8("你好")
	if !utf8.Valid([]byte(v)) {
		t.Fail()
	}
}

func TestIsSpace(t *testing.T) {
	if IsSpace('a') {
		t.Fail()
	}
	if !IsSpace(' ') {
		t.Fail()
	}
	if !IsSpace('\n') {
		t.Fail()
	}
	if !IsSpace('\t') {
		t.Fail()
	}
	if !IsSpace('\r') {
		t.Fail()
	}
}

func TestIsBlank(t *testing.T) {
	if !IsBlank(" ") {
		t.Fail()
	}
	if !IsBlank("\t\n\r") {
		t.Fail()
	}
	if IsBlank("abc") {
		t.Fail()
	}
}
