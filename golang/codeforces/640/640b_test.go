package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestCF640B(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{"1", args{strings.NewReader(`
			8
			10 3
			100 4
			8 7
			97 2
			8 8
			3 10
			5 3
			1000000000 9
		`)}, `YES
6 2 2
YES
97 1 1 1
NO
NO
YES
1 1 1 1 1 1 1 1
NO
YES
3 1 1
YES
999999984 2 2 2 2 2 2 2 2
`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := &bytes.Buffer{}
			CF640B(tt.args.input, output)
			if gotOutput := output.String(); gotOutput != tt.wantOutput {
				t.Errorf("CF640B() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
