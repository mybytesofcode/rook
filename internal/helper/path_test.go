package helper

import (
	"fmt"
	"os"
	"os/user"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func GetPWD() string {
	d, err := os.Getwd()
	if err != nil {
		// TODO: log
		return ""
	}

	return d
}

func GetHomeDir() string {
	u, err := user.Current()
	if err != nil {
		// TODO: log
		return ""
	}

	return u.HomeDir
}
func TestPathNormalize(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty path",
			args: args{
				path: "",
			},
			want: "",
		},
		{
			name: "Space path",
			args: args{
				path: " ",
			},
			want: "",
		},
		{
			name: "Absolute path",
			args: args{
				path: "/",
			},
			want: "/",
		},
		{
			name: "Relative path",
			args: args{
				path: "path",
			},
			want: fmt.Sprintf("%s/%s", GetPWD(), "path"),
		},
		{
			name: "Tilda path",
			args: args{
				path: "~",
			},
			want: GetHomeDir(),
		},
		{
			name: "Home path",
			args: args{
				path: "~/path",
			},
			want: fmt.Sprintf("%s/%s", GetHomeDir(), "path"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PathNormalize(tt.args.path); got != tt.want {
				t.Errorf("NormalizePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
