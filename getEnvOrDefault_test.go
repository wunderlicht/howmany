package main

import (
	"os"
	"testing"
)

func Test_getEnvOrDefaultString(t *testing.T) {
	const setEnvString = "Env String"
	type args struct {
		key      string
		fallback string
	}
	tests := []struct {
		name   string
		setEnv bool
		args   args
		want   string
	}{
		{"no env should return fallback", false,
			args{"SOME_KEY", "fb"}, "fb"},
		{"env set should return setEnvString", true,
			args{"SOME_KEY", "fb"}, setEnvString},
		// Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setEnv {
				_ = os.Setenv(tt.args.key, setEnvString) //maybe I should hard fail?
			}
			if got := getEnvOrDefaultString(tt.args.key, tt.args.fallback); got != tt.want {
				t.Errorf("getEnvOrDefaultString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEnvOrDefaultInt(t *testing.T) {
	const (
		setEnvInt = 1970
		setEnvStr = "1970"
	)

	type args struct {
		key      string
		fallback int
	}
	tests := []struct {
		name   string
		setEnv bool
		envVal string
		args   args
		want   int
	}{
		{"no env should return fallback",
			false, "",
			args{"SOME_KEY", 42}, 42},
		{"env set should return setEnvInt",
			true, setEnvStr,
			args{"SOME_KEY", 42}, setEnvInt},
		{"malformed env should return fallback",
			true, "hello, not a number",
			args{"SOME_KEY", 42}, 42},
		{"set but empty env should return fallback",
			true, "",
			args{"SOME_KEY", 42}, 42},
		// Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setEnv {
				_ = os.Setenv(tt.args.key, tt.envVal) //maybe I should hard fail?
			}
			if got := getEnvOrDefaultInt(tt.args.key, tt.args.fallback); got != tt.want {
				t.Errorf("getEnvOrDefaultInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEnvOrDefaultFloat(t *testing.T) {
	const (
		setEnvFloat = 99.9
		setEnvStr   = "99.9"
	)
	type args struct {
		key      string
		fallback float64
	}
	tests := []struct {
		name   string
		setEnv bool
		envVal string
		args   args
		want   float64
	}{
		{"no env should return fallback",
			false, "",
			args{"SOME_KEY", 42.0}, 42.0},
		{"env set should return setEnvFloat",
			true, setEnvStr,
			args{"SOME_KEY", 42.0}, setEnvFloat},
		{"malformed env should return fallback",
			true, "hello",
			args{"SOME_KEY", 42.0}, 42.0},
		{"empty env should return fallback",
			true, "",
			args{"SOME_KEY", 42.0}, 42.0},
		//Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setEnv {
				_ = os.Setenv(tt.args.key, tt.envVal) //maybe I should hard fail?
			}
			if got := getEnvOrDefaultFloat(tt.args.key, tt.args.fallback); got != tt.want {
				t.Errorf("getEnvOrDefaultFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEnvOrDefaultBool(t *testing.T) {
	type args struct {
		key      string
		fallback bool
	}
	tests := []struct {
		name   string
		setEnv bool
		envVal string
		args   args
		want   bool
	}{
		{"no env should return fallback",
			false, "",
			args{"SOME_KEY", false}, false},
		{"env set should return parsed env",
			true, "true",
			args{"SOME_KEY", false}, true},
		{"malformed env should return fallback",
			true, "hello",
			args{"SOME_KEY", true}, true},
		{"empty env should return fallback",
			true, "",
			args{"SOME_KEY", true}, true},
		// Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setEnv {
				_ = os.Setenv(tt.args.key, tt.envVal) //maybe I should hard fail?
			}
			if got := getEnvOrDefaultBool(tt.args.key, tt.args.fallback); got != tt.want {
				t.Errorf("getEnvOrDefaultBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
