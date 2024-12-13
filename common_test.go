package directive

import (
	"testing"
)

func Test_isNameMatch(t *testing.T) {
	type args struct {
		typeName string
		pattern  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "same type",
			args: args{
				pattern:  "ID",
				typeName: "ID",
			},
			want: true,
		},
		{
			name: "not same type",
			args: args{
				pattern:  "ID",
				typeName: "String",
			},
			want: false,
		},
		{
			name: "wildcard type",
			args: args{
				pattern:  ".+",
				typeName: "ID",
			},
			want: true,
		},
		{
			name: "same list type",
			args: args{
				pattern:  "[ID]",
				typeName: "[ID]",
			},
			want: true,
		},
		{
			name: "same type to list type",
			args: args{
				pattern:  "ID",
				typeName: "[ID]",
			},
			want: true,
		},
		{
			name: "not same list type",
			args: args{
				pattern:  "[ID]",
				typeName: "[String]",
			},
			want: false,
		},
		{
			name: "wildcard list type",
			args: args{
				pattern:  `\[.+\]`,
				typeName: "[ID]",
			},
			want: true,
		},
		{
			name: "not same wildcard list type",
			args: args{
				pattern:  `\[.+\]`,
				typeName: "ID",
			},
			want: false,
		},
		{
			name: "same wildcard list type",
			args: args{
				pattern:  ".+",
				typeName: "[ID]",
			},
			want: true,
		},
		{
			name: "same float type",
			args: args{
				pattern:  `^\[?Float\]?$`,
				typeName: "Float",
			},
			want: true,
		},
		{
			name: "same float list type",
			args: args{
				pattern:  `^\[?Float\]?$`,
				typeName: "[Float]",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.pattern, tt.args.typeName); got != tt.want {
				t.Errorf("isTypeNameMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
