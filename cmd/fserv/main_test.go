package main

import (
	"testing"
)

func TestIsDirExist(t *testing.T) {
	tests := []struct {
		path    string
		isExist bool
	}{
		{path: "./", isExist: true},
		{path: "../", isExist: true},
		{path: "hoge", isExist: false},
	}

	for n, tt := range tests {
		actual := IsDirExist(tt.path)
		if actual != tt.isExist {
			t.Errorf("#%d: got %v, want %v", n, actual, tt.isExist)
		}
	}
}
