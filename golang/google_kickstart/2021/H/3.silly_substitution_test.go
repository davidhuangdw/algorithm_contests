package main

import (
	"testing"
)

func Test_sillySub_small(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"0101010101"}, "22222"},
		{"1", args{"789012345678901234567890123456789"}, "99246802468024680"},
		{"1", args{"780412475609033434434122364570902"}, "33"},
		{"1", args{"6778989745900144589212552123456673212458090131223774556634678788956656455672330135621274512232107805"}, "214149"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sillySub_small(tt.args.S); got != tt.want {
				t.Errorf("sillySub() = %v, want %v", got, tt.want)
			}
		})
	}
}
