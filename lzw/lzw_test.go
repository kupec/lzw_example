package lzw

import (
	"bytes"
	"fmt"
	"testing"
)

var tests = []struct {
	input  []byte
	output []byte
}{
	{[]byte{}, []byte{}},
	{[]byte{1}, []byte{1, 0}},
	{[]byte{1, 1}, []byte{1, 0, 1, 0}},
	{[]byte{1, 1, 1}, []byte{1, 0, 0, 1}},
	{[]byte{1, 1, 1, 1}, []byte{1, 0, 0, 1, 1, 0}},
	{[]byte{1, 1, 1, 1, 1}, []byte{1, 0, 0, 1, 0, 1}},
	{[]byte{1, 1, 1, 1, 1, 1}, []byte{1, 0, 0, 1, 1, 1}},
	{[]byte{1, 2}, []byte{1, 0, 2, 0}},
	{[]byte{1, 2, 1, 2}, []byte{1, 0, 2, 0, 0, 1}},
}

func TestCompress(t *testing.T) {
	for i, test := range tests {
		testName := fmt.Sprintf("%d: %q -> %q", i, test.input, test.output)
		t.Run(testName, func(t *testing.T) {
			r := bytes.NewBuffer(test.input)
			var w bytes.Buffer
			if err := Compress(r, &w); err != nil {
				t.Fatalf("Compress(%q) return error: %q", test.input, err)
			}

			actual := w.Bytes()
			if !bytes.Equal(actual, test.output) {
				t.Fatalf("Compress(%q) = %q, want %q", test.input, actual, test.output)
			}
		})
	}
}
