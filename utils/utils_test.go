package utils

import (
	"strings"
	"testing"
	
)

// test  for unsupported and unallowed characters
func Test_IsValidInput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    bool
		wantErr string
	}{
		{
			name:    "No unsupported characters",
			input:   "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789???",
			want:    true,
			wantErr: "",
		},
		{
			name:    "One unsupported character",
			input:   "abcde\tfghij",
			want:    false,
			wantErr: "Error: input contains non-printable character: %q\n",
		},
		{
			name:    "Multiple unsupported characters",
			input:   "abcde\a\bfghi$j",
			want:    false,
			wantErr: "Error: input contains non-printable character: %q\n",
		},
		{
			name:  "Empty string",
			input: "",
			want:  true,
			wantErr: `Usage: go run . [OPTION] [STRING] [BANNER]

			Example: go run . --output=right something standard`,
		},
		{
			name:    "String with only unsupported characters",
			input:   "\a\b\t\v",
			want:    false,
			wantErr: "Error: input contains non-printable character: %q\n",
		},
		{
			name:    "String with special characters",
			input:   "Hello, world!How are you?",
			want:    true,
			wantErr: "",
		},

		{
			name:    "string with unicode characters",
			input:   "こんにちは、世界！",
			want:    false,
			wantErr: "Error: input contains unallowed character: %q\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, errmsg := IsValidInput(tt.input)
			if got != tt.want {
				t.Errorf("ContainsUnsupportedCharacters() got = %v, want %v", got, tt.want)
			}
			if got && !strings.Contains(errmsg, "Error: input contains non-printable/ unallowed character") != tt.want {
				t.Errorf("ContainsUnsupportedCharacters() returned unexpected error message: %s,", errmsg)
			}
		})
	}
}

// test for valid banner
func Test_ValidBanner(t *testing.T) {
	tests := []struct {
		name   string
		banner string
		want   bool
	}{
		{
			name:   "Valid banner 'standard'",
			banner: "standard",
			want:   true,
		},
		{
			name:   "Valid banner 'shadow'",
			banner: "shadow",
			want:   true,
		},
		{
			name:   "Valid banner 'thinkertoy'",
			banner: "thinkertoy",
			want:   true,
		},
		{
			name:   "Invalid banner 'other'",
			banner: "other",
			want:   false,
		},
		{
			name:   "Empty banner",
			banner: "",
			want:   false,
		},
		{
			name:   "Uppercase banner 'Standard'",
			banner: "Standard",
			want:   false,
		},
		{
			name:   "Mixed case banner 'ThinKerToy'",
			banner: "ThinKerToy",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidBanner(tt.banner); got != tt.want {
				t.Errorf("ValidBanner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckFlag(t *testing.T) {
	

	tests := []struct {
		name           string
		args           []string
		expectedAlign  string
		expectedOutput string
		expectedText   string
		expectedBanner string
	}{
		{
			name:           "Valid input with align and output",
			args:           []string{"program", "--align=center", "--output=file.txt", "Hello", "standard"},
			expectedAlign:  "center",
			expectedOutput: "file.txt",
			expectedText:   "Hello",
			expectedBanner: "standard",
		},
		{
			name:           "Valid input with only text",
			args:           []string{"program", "Hello"},
			expectedAlign:  "",
			expectedOutput: "",
			expectedText:   "Hello",
			expectedBanner: "",
		},
		
	
	
	}

	for _, tt := range tests {
		if _, _,got,_ := CheckFlag(tt.args); got != tt.expectedText {
			t.Errorf("CheckFlag() = %v, want %v", got, tt.expectedText)
		}
	}
}