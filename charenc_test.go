package charenc

import (
	"os"
	"testing"
)

type EncodeTest struct {
	path  string
	fname string
	f     func([]byte) ([]byte, error)
}

// to ansi
var toAnsiTests = []EncodeTest{
	{"raw\\BOMUTF8.txt", "BomUTF8ToAnsi", BomUTF8ToAnsi},
	{"raw\\UTF8.txt", "UTF8ToAnsi", UTF8ToAnsi},
	{"raw\\UTF16BE.txt", "UTF16BEToAnsi", UTF16BEToAnsi},
	{"raw\\UTF16LE.txt", "UTF16LEToAnsi", UTF16LEToAnsi},
}

// to utf8
var toUTF8Tests = []EncodeTest{
	{"raw\\ANSI.txt", "AnsiToUTF8", AnsiToUTF8},
	{"raw\\BOMUTF8.txt", "BomUTF8ToUTF8", BomUTF8ToUTF8},
	{"raw\\UTF16BE.txt", "UTF16BEToUTF8", UTF16BEToUTF8},
	{"raw\\UTF16LE.txt", "UTF16LEToUTF8", UTF16LEToUTF8},
}

func runEncodeTests(t *testing.T, testCases []EncodeTest) {
	for _, test := range testCases {
		stream, err := os.ReadFile(test.path)
		if err != nil {
			t.Errorf("%v: %v", test.fname, err)
		}
		stream, err = test.f(stream)
		if err != nil {
			t.Errorf("%v: %v", test.fname, err)
		}
		err = os.WriteFile(test.path, stream, 0666)
		if err != nil {
			t.Errorf("%v: %v", test.fname, err)
		}
		t.Logf("%v: ok", test.fname)
	}
}

func TestToAnsi(t *testing.T) {
	runEncodeTests(t, toAnsiTests)
}

func TestToUTF8(t *testing.T) {
	runEncodeTests(t, toUTF8Tests)
}
