package main

import (
	"testing"
)

func TestGitURL(t *testing.T) {
	if parseGit("git@github.com:jhaals/yopass.git") != "github.com/jhaals/yopass" {
		t.Error("failed")
	}
}

func TestHTTPURL(t *testing.T) {
	if parseHTTP("https://github.com/jhaals/yopass.git") != "github.com/jhaals/yopass" {
		t.Error("failed")
	}
}

func TestParse(t *testing.T) {
	testCases := []struct {
		desc    string
		input   string
		want    string
		wantErr bool
	}{
		{
			desc:    "git url",
			input:   "git@github.com:jhaals/yopass.git",
			want:    "github.com/jhaals/yopass",
			wantErr: false,
		},
		{
			desc:    "http url",
			input:   "https://github.com/jhaals/yopass.git",
			want:    "github.com/jhaals/yopass",
			wantErr: false,
		},
		{
			desc:    "No input",
			input:   "",
			want:    "",
			wantErr: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			path, err := ParseRepository(tC.input)
			if err != nil && !tC.wantErr {
				t.Errorf("Did not expect error")
			}
			if tC.want != path {
				t.Errorf("Wanted %s, got %s", tC.want, path)
			}
		})
	}
}
