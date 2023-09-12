package main

import (
	"reflect"
	"testing"
)

func Test_splitString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Without Chars",
			args: args{
				input: "",
			},
			want: []string{},
		},
		{
			name: "One Char",
			args: args{
				input: "C",
			},
			want: []string{"C"},
		},
		{
			name: "Multiple Chars",
			args: args{
				input: "COME",
			},
			want: []string{"C", "O", "M", "E"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitString(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid Without Chars",
			args: args{
				input: "",
			},
			want: true,
		},
		{
			name: "Invalid Only Open",
			args: args{
				input: "[",
			},
			want: false,
		},
		{
			name: "Invalid Only Close",
			args: args{
				input: "[",
			},
			want: false,
		},
		{
			name: "Invalid Close",
			args: args{
				input: "{]",
			},
			want: false,
		},
		{
			name: "Valid Single Brackets",
			args: args{
				input: "[]",
			},
			want: true,
		},
		{
			name: "Valid Single Braces",
			args: args{
				input: "{}",
			},
			want: true,
		},
		{
			name: "Valid Single Parenthesis",
			args: args{
				input: "()",
			},
			want: true,
		},
		{
			name: "Valid Multiple",
			args: args{
				input: "([{}])",
			},
			want: true,
		},
		{
			name: "Invalid Multiple",
			args: args{
				input: "({[}])",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateString(tt.args.input); got != tt.want {
				t.Errorf("validateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
