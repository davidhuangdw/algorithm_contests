package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestCFXXX(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{"1", args{strings.NewReader(`0 1`)}, ``},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := &bytes.Buffer{}
			CFXXX(tt.args.input, output)
			if gotOutput := output.String(); gotOutput != tt.wantOutput {
				t.Errorf("CFXXX() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
